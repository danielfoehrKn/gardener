package plant

import (
	"context"
	"fmt"
	"strings"

	gardencorev1alpha1 "github.com/gardener/gardener/pkg/apis/core/v1alpha1"
	"github.com/gardener/gardener/pkg/client/kubernetes"
	kutil "github.com/gardener/gardener/pkg/utils/kubernetes"
	"k8s.io/apimachinery/pkg/api/equality"
	kubernetesclientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/retry"

	"github.com/pkg/errors"

	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// Following labels come from k8s.io/kubernetes/pkg/kubelet/apis

	// LabelZoneFailureDomain zone failure domain label
	LabelZoneFailureDomain = "failure-domain.beta.kubernetes.io/zone"
	// LabelZoneRegion zone region label
	LabelZoneRegion = "failure-domain.beta.kubernetes.io/region"
)

// getClusterInfo gets the kubernetes cluster zones and region by inspecting labels on nodes in the cluster.
func getClusterInfo(ctx context.Context, cl client.Client, logger logrus.FieldLogger) (*plantStatusInfo, error) {
	var nodes = &corev1.NodeList{}

	err := cl.List(ctx, &client.ListOptions{Namespace: "garden"}, nodes)
	if err != nil {
		logger.Errorf("Failed to list nodes while getting zone names: %v", err)
		return nil, err
	}

	if len(nodes.Items) == 0 {
		return nil, fmt.Errorf("there are no nodes available in this cluster to retrieve zones and regions from")
	}

	// we are only taking the first node because all nodes that
	firstNode := nodes.Items[0]
	region, err := getRegionNameForNode(firstNode)
	if err != nil {
		return nil, err
	}

	provider := getCloudProviderForNode(firstNode.Spec.ProviderID)

	return &plantStatusInfo{
		region:    region,
		cloudType: provider,
	}, nil
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

func (c *defaultPlantControl) fetchCloudInfo(ctx context.Context, plant *gardencorev1alpha1.Plant, logger logrus.FieldLogger) (*plantStatusInfo, error) {
	if c.plantClient == nil || c.discoveryClient == nil {
		return nil, fmt.Errorf("plant clients need to be initialized first")
	}

	cloudInfo, err := getClusterInfo(ctx, c.plantClient, logger)
	if err != nil {
		return nil, err
	}

	kubernetesVersionInfo, err := c.discoveryClient.ServerVersion()
	if err != nil {
		return nil, err
	}

	cloudInfo.k8sVersion = kubernetesVersionInfo.String()

	return cloudInfo, nil
}

func (c *defaultPlantControl) intializeClientsWithUpdateFunc(plant *gardencorev1alpha1.Plant, kubeconfig []byte, needsClientUpdate func() bool) error {
	if c.discoveryClient == nil || c.plantClient == nil || needsClientUpdate() {
		fmt.Println("NEEDS CLIENT UPDATE")
		return c.initializePlantClients(plant, kubeconfig)
	}
	return nil
}

func (c *defaultPlantControl) initializePlantClients(plant *gardencorev1alpha1.Plant, kubeconfigSecretValue []byte) error {
	config, err := clientcmd.RESTConfigFromKubeConfig(kubeconfigSecretValue)
	if err != nil {
		return err
	}
	plantClusterClient, err := kubernetes.NewRuntimeClientForConfig(config, client.Options{
		Scheme: kubernetes.PlantScheme,
	})

	discoveryClient, err := kubernetesclientset.NewForConfig(config)
	if err != nil {
		return err
	}

	c.plantClient = plantClusterClient
	c.discoveryClient = discoveryClient

	return nil
}

func getCloudProviderForNode(providerID string) string {

	provider := strings.Split(providerID, "://")
	if len(provider) == 0 {
		return "<unknown>"
	}

	return provider[0]
}

// getRegionNameForNode Finds the name of the region in which a Node is running.
func getRegionNameForNode(node corev1.Node) (string, error) {
	for key, value := range node.Labels {
		if key == LabelZoneRegion {
			return value, nil
		}
	}
	return "", errors.Errorf("Region name for node %s not found. No label with key %s",
		node.Name, LabelZoneRegion)
}
