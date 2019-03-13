// Copyright (c) 2018 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package plant

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/controller-runtime/pkg/client"

	gardencorev1alpha1 "github.com/gardener/gardener/pkg/apis/core/v1alpha1"
	"github.com/gardener/gardener/pkg/apis/core/v1alpha1/helper"
	gardencorelisters "github.com/gardener/gardener/pkg/client/core/listers/core/v1alpha1"
	"github.com/gardener/gardener/pkg/client/kubernetes"
	"github.com/gardener/gardener/pkg/controllermanager/apis/config"
	"github.com/gardener/gardener/pkg/logger"
	"github.com/gardener/gardener/pkg/utils"
	kutil "github.com/gardener/gardener/pkg/utils/kubernetes"
	kubernetesclientset "k8s.io/client-go/kubernetes"
	kubecorev1listers "k8s.io/client-go/listers/core/v1"

	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/equality"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/retry"
)

const kubeconfigChecksumAnnotationKey = "kubeconfig.secret.checksum/value"

// plantAdd adds the plant resource
func (c *Controller) updateClients(obj interface{}) {
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err != nil {
		logger.Logger.Errorf("Couldn't get key for object %+v: %v", obj, err)
		return
	}
	c.plantQueue.Add(key)
}

// plantAdd adds the plant resource
func (c *Controller) plantAdd(obj interface{}) {
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err != nil {
		logger.Logger.Errorf("Couldn't get key for object %+v: %v", obj, err)
		return
	}
	c.plantQueue.Add(key)
}

// plantUpdate updates the plant resource
func (c *Controller) plantUpdate(oldObj, newObj interface{}) {
	_, ok1 := oldObj.(*gardencorev1alpha1.Plant)
	_, ok2 := newObj.(*gardencorev1alpha1.Plant)
	if !ok1 || !ok2 {
		return
	}

	c.plantAdd(newObj)
}

func (c *Controller) plantDelete(obj interface{}) {
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
	if err != nil {
		logger.Logger.Errorf("Couldn't get key for object %+v: %v", obj, err)
		return
	}
	c.plantQueue.Add(key)
}

func (c *Controller) reconcilePlantKey(ctx context.Context, key string) error {
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return err
	}

	plant, err := c.plantLister.Plants(namespace).Get(name)
	if apierrors.IsNotFound(err) {
		logger.Logger.Debugf("[PLANT RECONCILE] %s - skipping because Plant has been deleted", key)
		return nil
	}
	if err != nil {
		logger.Logger.Infof("[PLANT RECONCILE] %s - unable to retrieve object from store: %v", key, err)
		return err
	}
	if err := c.plantControl.Reconcile(ctx, plant, key); err != nil {
		return err
	}

	c.plantQueue.AddAfter(key, c.config.Controllers.Plant.SyncPeriod.Duration)

	return nil

}

// ControlInterface implements the control logic for updating Plants. It is implemented as an interface to allow
// for extensions that provide different semantics. Currently, there is only one implementation.
type ControlInterface interface {
	Reconcile(context.Context, *gardencorev1alpha1.Plant, string) error
}

// NewDefaultPlantControl returns a new instance of the default implementation ControlInterface that
// implements the documented semantics for Plants. updater is the UpdaterInterface used
// to update the status of Plants.
func NewDefaultPlantControl(k8sGardenClient kubernetes.Interface, recorder record.EventRecorder, config *config.ControllerManagerConfiguration, plantsLister gardencorelisters.PlantLister, secretLister kubecorev1listers.SecretLister) ControlInterface {
	return &defaultPlantControl{
		k8sGardenClient: k8sGardenClient,
		plantClient:     make(map[string]client.Client),
		discoveryClient: make(map[string]*kubernetesclientset.Clientset),
		healthChecker:   make(map[string]*HealthChecker),
		plantLister:     plantsLister,
		secretsLister:   secretLister,
		recorder:        recorder,
		config:          config,
	}
}

func (c *defaultPlantControl) Reconcile(ctx context.Context, obj *gardencorev1alpha1.Plant, key string) error {
	var (
		plant  = obj.DeepCopy()
		logger = logger.NewFieldLogger(logger.Logger, "plant", plant.Name)
	)

	if plant.DeletionTimestamp != nil {
		return c.delete(plant, logger)
	}

	return c.reconcile(ctx, plant, key, logger)
}

func (c *defaultPlantControl) updatePlantConditions(plant *gardencorev1alpha1.Plant, conditions ...gardencorev1alpha1.Condition) (*gardencorev1alpha1.Plant, error) {
	newPlant, err := kutil.TryUpdatePlantConditions(c.k8sGardenClient.GardenCore(), retry.DefaultBackoff, plant.ObjectMeta,
		func(plant *gardencorev1alpha1.Plant) (*gardencorev1alpha1.Plant, error) {
			plant.Status.Conditions = conditions
			return plant, nil
		})

	return newPlant, err
}

func (c *defaultPlantControl) reconcile(ctx context.Context, plant *gardencorev1alpha1.Plant, key string, logger logrus.FieldLogger) error {
	_, err := kutil.TryUpdatePlantStatusWithEqualFunc(c.k8sGardenClient.GardenCore(), retry.DefaultBackoff, plant.ObjectMeta,
		func(p *gardencorev1alpha1.Plant) (*gardencorev1alpha1.Plant, error) {
			if finalizers := sets.NewString(p.Finalizers...); !finalizers.Has(FinalizerName) {
				finalizers.Insert(FinalizerName)
				p.Finalizers = finalizers.UnsortedList()
			}
			return p, nil
		}, func(cur, updated *gardencorev1alpha1.Plant) bool {
			return sets.NewString(cur.Finalizers...).Has(FinalizerName)
		})
	if err != nil {
		return err
	}

	var (
		newConditions               = helper.MergeConditions(plant.Status.Conditions, helper.InitCondition(gardencorev1alpha1.PlantAPIServerAvailable), helper.InitCondition(gardencorev1alpha1.PlantEveryNodeReady))
		conditionAPIServerAvailable = newConditions[0]
		conditionEveryNodeReady     = newConditions[1]
	)

	defer func() {
		if _, err := c.updateConditions(plant, conditionAPIServerAvailable, conditionEveryNodeReady); err != nil {
			logger.Errorf("Failed to update the conditions : %+v", err)
		}
	}()

	kubeconfigSecret, err := c.secretsLister.Secrets(plant.Spec.SecretRef.Namespace).Get(plant.Spec.SecretRef.Name)
	if err != nil {
		return err
	}
	kubeconfig, ok := kubeconfigSecret.Data["kubeconfig"]
	if !ok {
		message := "Plant secret needs to contain a kubeconfig key."
		conditionAPIServerAvailable = helper.UpdatedCondition(conditionAPIServerAvailable, corev1.ConditionFalse, "APIServerDown", message)
		conditionEveryNodeReady = helper.UpdatedCondition(conditionEveryNodeReady, corev1.ConditionFalse, "Nodes not reachable", message)
		resetClients(c, key)

		return c.updateStatus(plant, &StatusCloudInfo{}, conditionAPIServerAvailable, conditionEveryNodeReady)
	}

	secretChecksum := utils.ComputeSHA256Hex([]byte(strings.TrimSpace(string(kubeconfig))))
	_, err = kutil.TryUpdatePlantWithEqualFunc(c.k8sGardenClient.GardenCore(), retry.DefaultBackoff, plant.ObjectMeta,
		func(p *gardencorev1alpha1.Plant) (*gardencorev1alpha1.Plant, error) {
			if !metav1.HasAnnotation(p.ObjectMeta, kubeconfigChecksumAnnotationKey) || p.Annotations[kubeconfigChecksumAnnotationKey] != secretChecksum {
				metav1.SetMetaDataAnnotation(&p.ObjectMeta, kubeconfigChecksumAnnotationKey, secretChecksum)
			}
			return p, nil
		}, func(cur, updated *gardencorev1alpha1.Plant) bool {
			return equality.Semantic.DeepEqual(cur, updated)
		})
	if err != nil {
		return err
	}

	// only initialize / re-initialize the clients in-case the kubeconfig for the Plant cluster changes
	if err := c.intializeClientsWithUpdateFunc(plant, key, kubeconfig, func() bool {
		return plant.Annotations[kubeconfigChecksumAnnotationKey] != utils.ComputeSHA256Hex([]byte(strings.TrimSpace(string(kubeconfig))))
	}); err != nil {
		message := fmt.Sprintf("Could not initialize Plant clients for health check: %+v", err)
		conditionAPIServerAvailable = helper.UpdatedCondition(conditionAPIServerAvailable, corev1.ConditionFalse, "APIServerDown", "Could not reach API server during client initialization.")
		conditionEveryNodeReady = helper.UpdatedConditionUnknownErrorMessage(conditionEveryNodeReady, message)
		resetClients(c, key)

		return fmt.Errorf("%v:%v", c.updateStatus(plant, &StatusCloudInfo{}, conditionAPIServerAvailable, conditionEveryNodeReady), err)
	}

	// Trigger health check
	conditionAPIServerAvailable, conditionEveryNodeReady = c.healthChecks(ctx, key, logger, conditionAPIServerAvailable, conditionEveryNodeReady)

	cloudInfo, err := FetchCloudInfo(ctx, c.plantClient[key], c.discoveryClient[key], logger)
	if err != nil {
		return err
	}

	return c.updateStatus(plant, cloudInfo, conditionAPIServerAvailable, conditionEveryNodeReady)
}

func (c *defaultPlantControl) updateStatus(plant *gardencorev1alpha1.Plant, cloudInfo *StatusCloudInfo, conditions ...gardencorev1alpha1.Condition) error {
	_, err := kutil.TryUpdatePlantStatusWithEqualFunc(c.k8sGardenClient.GardenCore(), retry.DefaultBackoff, plant.ObjectMeta,
		func(p *gardencorev1alpha1.Plant) (*gardencorev1alpha1.Plant, error) {
			if p.Status.ClusterInfo == nil {
				p.Status.ClusterInfo = &gardencorev1alpha1.ClusterInfo{}
			}

			p.Status.ClusterInfo.Cloud.Type = cloudInfo.CloudType
			p.Status.ClusterInfo.Cloud.Region = cloudInfo.Region
			p.Status.ClusterInfo.Kubernetes.Version = cloudInfo.K8sVersion
			p.Status.Conditions = conditions
			return p, nil
		}, func(cur, updated *gardencorev1alpha1.Plant) bool {
			return equality.Semantic.DeepEqual(cur, updated)
		})
	return err
}

func (c *defaultPlantControl) delete(plant *gardencorev1alpha1.Plant, logger logrus.FieldLogger) error {
	_, err := kutil.TryUpdatePlantStatusWithEqualFunc(c.k8sGardenClient.GardenCore(), retry.DefaultBackoff, plant.ObjectMeta, func(c *gardencorev1alpha1.Plant) (*gardencorev1alpha1.Plant, error) {
		finalizers := sets.NewString(c.Finalizers...)
		finalizers.Delete(FinalizerName)
		c.Finalizers = finalizers.UnsortedList()
		return c, nil
	}, func(cur, updated *gardencorev1alpha1.Plant) bool {
		return !sets.NewString(cur.Finalizers...).Has(FinalizerName)
	})
	return err
}

func (c *defaultPlantControl) updateConditions(plant *gardencorev1alpha1.Plant, conditions ...gardencorev1alpha1.Condition) (*gardencorev1alpha1.Plant, error) {
	return kutil.TryUpdatePlantStatusWithEqualFunc(c.k8sGardenClient.GardenCore(), retry.DefaultBackoff, plant.ObjectMeta,
		func(plant *gardencorev1alpha1.Plant) (*gardencorev1alpha1.Plant, error) {
			plant.Status.Conditions = conditions
			return plant, nil
		}, func(cur, updated *gardencorev1alpha1.Plant) bool {
			return equality.Semantic.DeepEqual(cur.Status.Conditions, updated.Status.Conditions)
		},
	)
}

func (c *defaultPlantControl) intializeClientsWithUpdateFunc(plant *gardencorev1alpha1.Plant, key string, kubeconfig []byte, needsClientsUpdate func() bool) error {
	if c.discoveryClient[key] == nil || c.plantClient[key] == nil || needsClientsUpdate() {
		if err := c.initializePlantClients(plant, key, kubeconfig); err != nil {
			return err
		}
		c.initializeHealthChecker(key)
		return nil
	}
	return nil
}

func (c *defaultPlantControl) initializeHealthChecker(key string) {
	c.healthChecker[key] = NewHealthCheker(c.plantClient[key], c.discoveryClient[key])
}

func (c *defaultPlantControl) initializePlantClients(plant *gardencorev1alpha1.Plant, key string, kubeconfigSecretValue []byte) error {
	config, err := clientcmd.RESTConfigFromKubeConfig(kubeconfigSecretValue)
	if err != nil {
		return fmt.Errorf("%v:%v", "invalid kubconfig supplied resulted in: ", err)
	}
	plantClusterClient, err := kubernetes.NewRuntimeClientForConfig(config, client.Options{
		Scheme: kubernetes.PlantScheme,
	})

	discoveryClient, err := kubernetesclientset.NewForConfig(config)
	if err != nil {
		return err
	}

	c.plantClient[key] = plantClusterClient
	c.discoveryClient[key] = discoveryClient

	return nil
}

func (c *defaultPlantControl) healthChecks(ctx context.Context, key string, logger logrus.FieldLogger, apiserverAvailability, healthyNodes gardencorev1alpha1.Condition) (gardencorev1alpha1.Condition, gardencorev1alpha1.Condition) {
	var (
		wg      sync.WaitGroup
		checker = c.healthChecker[key]
	)

	wg.Add(2)
	go func() {
		defer wg.Done()
		apiserverAvailability = checker.CheckAPIServerAvailability(apiserverAvailability)
	}()
	go func() {
		defer wg.Done()
		newNodes, err := checker.CheckPlantClusterNodes(&healthyNodes, checker.makePlantNodeLister(ctx, &client.ListOptions{}))
		healthyNodes = newConditionOrError(healthyNodes, *newNodes, err)
	}()

	wg.Wait()

	return apiserverAvailability, healthyNodes
}
