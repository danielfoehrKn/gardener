package plant

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	kutil "github.com/gardener/gardener/pkg/utils/kubernetes"
	"k8s.io/apimachinery/pkg/labels"

	"github.com/sirupsen/logrus"

	gardencorev1alpha1 "github.com/gardener/gardener/pkg/apis/core/v1alpha1"
	"github.com/gardener/gardener/pkg/utils/kubernetes/health"
	corev1 "k8s.io/api/core/v1"
	kubernetesclientset "k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/gardener/gardener/pkg/apis/core/v1alpha1/helper"
)

// NewPlantHealthChecker creates a new health checker.
func NewPlantHealthChecker(plantClient client.Client, discoveryClient *kubernetesclientset.Clientset) *HealthChecker {
	return &HealthChecker{
		plantClient:     plantClient,
		discoveryClient: discoveryClient,
	}
}

// HealthChecker contains the condition thresholds.
type HealthChecker struct {
	plantClient     client.Client
	discoveryClient *kubernetesclientset.Clientset
}

// CheckPlantClusterNodes checks whether cluster nodes in the given listers are complete and healthy.
func (c *defaultPlantControl) CheckPlantClusterNodes(condition *gardencorev1alpha1.Condition, nodeLister kutil.NodeLister) (*gardencorev1alpha1.Condition, error) {
	nodeList, err := nodeLister.List(labels.Everything())
	if err != nil {
		return nil, err
	}

	if exitCondition := c.checkNodes(*condition, nodeList); exitCondition != nil {
		return exitCondition, nil
	}

	updatedCondition := helper.UpdatedCondition(*condition, corev1.ConditionTrue, "EveryNodeReady", "Every node registered to the cluster is ready")
	return &updatedCondition, nil
}

func (c *defaultPlantControl) checkNodes(condition gardencorev1alpha1.Condition, objects []*corev1.Node) *gardencorev1alpha1.Condition {
	for _, object := range objects {
		if err := health.CheckNode(object); err != nil {
			failedCondition := helper.UpdatedCondition(condition, corev1.ConditionFalse, "NodesUnhealthy", fmt.Sprintf("Node %s is unhealthy: %v", object.Name, err))
			return &failedCondition
		}
	}
	return nil
}

// checkAPIServerAvailability checks if the API server of a Shoot cluster is reachable and measure the response time.
func (c *defaultPlantControl) checkAPIServerAvailability(key string, condition gardencorev1alpha1.Condition) gardencorev1alpha1.Condition {
	// Try to reach the Shoot API server and measure the response time.
	now := time.Now()
	response := c.discoveryClient[key].RESTClient().Get().AbsPath("/healthz").Do()
	responseDurationText := fmt.Sprintf("[response_time:%dms]", time.Now().Sub(now).Nanoseconds()/time.Millisecond.Nanoseconds())
	if response.Error() != nil {
		message := fmt.Sprintf("Request to Shoot API server /healthz endpoint failed. %s (%s)", responseDurationText, response.Error().Error())
		return helper.UpdatedCondition(condition, corev1.ConditionFalse, "HealthzRequestFailed", message)
	}

	// Determine the status code of the response.
	var statusCode int
	response.StatusCode(&statusCode)

	if statusCode != http.StatusOK {
		var body string
		bodyRaw, err := response.Raw()
		if err != nil {
			body = fmt.Sprintf("Could not parse response body: %s", err.Error())
		} else {
			body = string(bodyRaw)
		}
		message := fmt.Sprintf("Shoot API server /healthz endpoint endpoint check returned a non ok status code %d. %s (%s)", statusCode, responseDurationText, body)
		return helper.UpdatedCondition(condition, corev1.ConditionFalse, "HealthzRequestError", message)
	}

	message := fmt.Sprintf("Shoot API server /healthz endpoint responded with success status code. %s", responseDurationText)
	return helper.UpdatedCondition(condition, corev1.ConditionTrue, "HealthzRequestSucceeded", message)
}

func (c *defaultPlantControl) healthChecks(ctx context.Context, key string, logger logrus.FieldLogger, apiserverAvailability, nodes gardencorev1alpha1.Condition) (gardencorev1alpha1.Condition, gardencorev1alpha1.Condition) {
	var (
		wg sync.WaitGroup
	)

	wg.Add(2)
	go func() {
		defer wg.Done()
		apiserverAvailability = c.checkAPIServerAvailability(key, apiserverAvailability)
	}()
	go func() {
		defer wg.Done()
		newNodes, err := c.CheckPlantClusterNodes(&nodes, c.makePlantNodeLister(ctx, key, &client.ListOptions{}))
		nodes = newConditionOrError(nodes, *newNodes, err)
	}()

	wg.Wait()

	return apiserverAvailability, nodes
}

func newConditionOrError(oldCondition, newCondition gardencorev1alpha1.Condition, err error) gardencorev1alpha1.Condition {
	if err != nil {
		return helper.UpdatedConditionUnknownError(oldCondition, err)
	}
	return newCondition
}

func (c *defaultPlantControl) makePlantNodeLister(ctx context.Context, key string, options *client.ListOptions) kutil.NodeLister {
	var (
		once  sync.Once
		items []*corev1.Node
		err   error

		onceBody = func() {
			nodeList := &corev1.NodeList{}
			err = c.plantClient[key].List(ctx, options, nodeList)
			if err != nil {
				return
			}

			for _, item := range nodeList.Items {
				it := item
				items = append(items, &it)
			}
		}
	)

	return kutil.NewNodeLister(func() ([]*corev1.Node, error) {
		once.Do(onceBody)
		return items, err
	})
}
