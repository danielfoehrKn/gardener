// Copyright (c) 2021 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
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

package kubelet

import (
	"fmt"
	"math"
	"time"

	gardencorev1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	"github.com/gardener/gardener/pkg/operation/botanist/component/extensions/operatingsystemconfig/original/components"
	"k8s.io/apimachinery/pkg/api/resource"

	"github.com/Masterminds/semver"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubeletconfigv1beta1 "k8s.io/kubelet/config/v1beta1"
	"k8s.io/utils/pointer"
)

// Config returns a kubelet config based on the provided parameters and for the provided Kubernetes version.
func Config(kubernetesVersion *semver.Version, clusterDNSAddress, clusterDomain string, machineType *gardencorev1beta1.MachineType, rootVolume *gardencorev1beta1.Volume, params components.ConfigurableKubeletConfigParameters) (*kubeletconfigv1beta1.KubeletConfiguration, error) {
	// set default values for the kube-reserved resources and eviction thresholds
	err := setConfigDefaults(machineType, rootVolume, &params)
	if err != nil {
		return nil, err
	}

	config := &kubeletconfigv1beta1.KubeletConfiguration{
		Authentication: kubeletconfigv1beta1.KubeletAuthentication{
			Anonymous: kubeletconfigv1beta1.KubeletAnonymousAuthentication{
				Enabled: pointer.BoolPtr(false),
			},
			X509: kubeletconfigv1beta1.KubeletX509Authentication{
				ClientCAFile: PathKubeletCACert,
			},
			Webhook: kubeletconfigv1beta1.KubeletWebhookAuthentication{
				Enabled:  pointer.BoolPtr(true),
				CacheTTL: metav1.Duration{Duration: 2 * time.Minute},
			},
		},
		Authorization: kubeletconfigv1beta1.KubeletAuthorization{
			Mode: kubeletconfigv1beta1.KubeletAuthorizationModeWebhook,
			Webhook: kubeletconfigv1beta1.KubeletWebhookAuthorization{
				CacheAuthorizedTTL:   metav1.Duration{Duration: 5 * time.Minute},
				CacheUnauthorizedTTL: metav1.Duration{Duration: 30 * time.Second},
			},
		},
		CgroupDriver:                     "cgroupfs",
		CgroupRoot:                       "/",
		CgroupsPerQOS:                    pointer.BoolPtr(true),
		ClusterDNS:                       []string{clusterDNSAddress},
		ClusterDomain:                    clusterDomain,
		CPUCFSQuota:                      params.CpuCFSQuota,
		CPUManagerPolicy:                 *params.CpuManagerPolicy,
		CPUManagerReconcilePeriod:        metav1.Duration{Duration: 10 * time.Second},
		EnableControllerAttachDetach:     pointer.BoolPtr(true),
		EnableDebuggingHandlers:          pointer.BoolPtr(true),
		EnableServer:                     pointer.BoolPtr(true),
		EnforceNodeAllocatable:           []string{"pods"},
		EventBurst:                       50,
		EventRecordQPS:                   pointer.Int32Ptr(50),
		EvictionHard:                     params.EvictionHard,
		EvictionMinimumReclaim:           params.EvictionMinimumReclaim,
		EvictionSoft:                     params.EvictionSoft,
		EvictionSoftGracePeriod:          params.EvictionSoftGracePeriod,
		EvictionPressureTransitionPeriod: *params.EvictionPressureTransitionPeriod,
		EvictionMaxPodGracePeriod:        *params.EvictionMaxPodGracePeriod,
		FailSwapOn:                       params.FailSwapOn,
		FeatureGates:                     params.FeatureGates,
		FileCheckFrequency:               metav1.Duration{Duration: 20 * time.Second},
		HairpinMode:                      kubeletconfigv1beta1.PromiscuousBridge,
		HTTPCheckFrequency:               metav1.Duration{Duration: 20 * time.Second},
		ImageGCHighThresholdPercent:      pointer.Int32Ptr(50),
		ImageGCLowThresholdPercent:       pointer.Int32Ptr(40),
		ImageMinimumGCAge:                metav1.Duration{Duration: 2 * time.Minute},
		KubeAPIBurst:                     50,
		KubeAPIQPS:                       pointer.Int32Ptr(50),
		KubeReserved:                     params.KubeReserved,
		MaxOpenFiles:                     1000000,
		MaxPods:                          *params.MaxPods,
		NodeStatusUpdateFrequency:        metav1.Duration{Duration: 10 * time.Second},
		PodsPerCore:                      0,
		PodPidsLimit:                     params.PodPidsLimit,
		ReadOnlyPort:                     0,
		RegistryBurst:                    10,
		RegistryPullQPS:                  pointer.Int32Ptr(5),
		ResolverConfig:                   "/etc/resolv.conf",
		RotateCertificates:               true,
		RuntimeRequestTimeout:            metav1.Duration{Duration: 2 * time.Minute},
		SerializeImagePulls:              pointer.BoolPtr(true),
		SyncFrequency:                    metav1.Duration{Duration: time.Minute},
		SystemReserved:                   params.SystemReserved,
		VolumeStatsAggPeriod:             metav1.Duration{Duration: time.Minute},
	}

	if versionConstraintK8sGreaterEqual119.Check(kubernetesVersion) {
		config.VolumePluginDir = pathVolumePluginDirectory
	}

	return config, nil
}

var (
	evictionHardDefaults = map[string]string{
		components.MemoryAvailable:   "100Mi",
		components.ImageFSAvailable:  "15%",
		components.ImageFSInodesFree: "5%",
		components.NodeFSAvailable:   "10%",
		components.NodeFSInodesFree:  "5%",
		components.PIDAvailable:      "10%",
	}
)

// TODO: remove soft eviction thresholds and their options
// Reason:
//    - they are not a default configuration of the kubelet
//    - soft eviction has the risk of running into Node condition oscillation
//      User can still configure their own soft-evictions https://kubernetes.io/docs/concepts/scheduling-eviction/node-pressure-eviction/#node-condition-oscillation
//    - Workload might run slightly below the hard eviction threshold stable without any problems (no need for soft eviction) --> workload specific
//    - more simple configuration
//    - more intelligent kube-reserved should lower the risk for processes outside the kubepods cgroup
//     --> hopefully no need to have even earlier evictions

func setConfigDefaults(machineType *gardencorev1beta1.MachineType, rootVolume *gardencorev1beta1.Volume, c *components.ConfigurableKubeletConfigParameters) error {
	if c.KubeReserved == nil {
		c.KubeReserved = make(map[string]string, 2)
	}

	kubeReservedDefaults, err := calculateKubeReserved(machineType, rootVolume)
	if err != nil {
		return err
	}

	// set default kube-reserved values if not already specified
	for k, v := range kubeReservedDefaults {
		if c.KubeReserved[k] == "" {
			c.KubeReserved[k] = v
		}
	}

	if c.EvictionHard == nil {
		c.EvictionHard = make(map[string]string, 5)
	}

	// set default hard-eviction values if not already specified
	for k, v := range evictionHardDefaults {
		if c.EvictionHard[k] == "" {
			c.EvictionHard[k] = v
		}
	}

	if c.CpuCFSQuota == nil {
		c.CpuCFSQuota = pointer.BoolPtr(true)
	}

	if c.CpuManagerPolicy == nil {
		c.CpuManagerPolicy = pointer.StringPtr(kubeletconfigv1beta1.NoneTopologyManagerPolicy)
	}

	if c.EvictionPressureTransitionPeriod == nil {
		c.EvictionPressureTransitionPeriod = &metav1.Duration{Duration: 4 * time.Minute}
	}

	if c.EvictionMaxPodGracePeriod == nil {
		c.EvictionMaxPodGracePeriod = pointer.Int32Ptr(90)
	}

	if c.FailSwapOn == nil {
		c.FailSwapOn = pointer.BoolPtr(true)
	}

	if c.MaxPods == nil {
		c.MaxPods = pointer.Int32Ptr(110)
	}

	return nil
}

// calculateKubeReserved calculates kube-reserved based on the machine type
// As larger machine types tend to run more containers (and by extension, more Pods), the amount of resources
// that Kubernetes system pods (e.g, due to VPA), the container runtime (e.g more containerd-shim processes)
// and the kubelet (more pods to handle) requires is also larger.
func calculateKubeReserved(machineType *gardencorev1beta1.MachineType, rootVolume *gardencorev1beta1.Volume) (map[string]string, error) {
	i := 2048
	kubeReserved := map[string]string{
		"pid": fmt.Sprintf("%", i),
	}

	var bootDiskSize *resource.Quantity
	if machineType != nil && machineType.Storage != nil && machineType.Storage.StorageSize != nil {
		bootDiskSize = machineType.Storage.StorageSize
	}

	if rootVolume != nil {
		volSize, err := resource.ParseQuantity(rootVolume.VolumeSize)
		if err != nil {
			return nil, fmt.Errorf("failed to parse worker volume size %q: %w", rootVolume.VolumeSize, err)
		}
		bootDiskSize = &volSize
	}

	if bootDiskSize != nil {
		kubeReserved[string(corev1.ResourceEphemeralStorage)] = CalculateReservedEphemeralStorage(*bootDiskSize)
	}

	// currently it is possible to delete a machine type from the cloud profile that is in use by Shoots
	// return default values
	if machineType == nil {
		kubeReserved[string(corev1.ResourceCPU)] = "80m"
		kubeReserved[string(corev1.ResourceMemory)] = "1Gi"
		return kubeReserved, nil
	}

	kubeReserved[string(corev1.ResourceCPU)] = CalculateReservedCPU(machineType.CPU)
	kubeReserved[string(corev1.ResourceMemory)] = CalculateReservedMemory(machineType.Memory)

	return kubeReserved, nil
}

var (
	oneCore  = resource.MustParse("1")
	twoCores = resource.MustParse("2")
)

// Azure
// CPU cores on host			1	2	4	8	16	32	64
// Kube-reserved (millicores)	60	100	140	180	260	420	740
// TODO: what is more realistic? GKE or Azure? Do I need to adjust the formula

// CalculateReservedCPU calculates a regressive amount of cpu reservations
// 6% of the first core
// 1% of the next core (up to 2 cores)
// 0.5% of the next 2 cores (up to 4 cores)
// 0.25% of any cores above 4 cores
func CalculateReservedCPU(cpu resource.Quantity) string {
	if oneCore.Cmp(cpu) >= 0 {
		cpu.SetMilli(int64(float64(cpu.MilliValue()) * 0.06))
		return cpu.String()
	}

	cpu.Sub(oneCore)

	// 1% of the next core (up to 2 cores)
	if oneCore.Cmp(cpu) >= 0 {
		// 6 % of one core
		// + 1% of the next core
		cpu.SetMilli(int64(60 + float64(cpu.MilliValue())*0.01))
		return cpu.String()
	}

	cpu.Sub(oneCore)

	// 0.5% of the next 2 cores (up to 4 cores)
	if twoCores.Cmp(cpu) >= 0 {
		// 6 % of one core
		// + 1% of the next core
		// + 0.5% of the next 2 cores
		cpu.SetMilli(int64(60 + 10 + float64(cpu.MilliValue())*0.005))
		return cpu.String()
	}

	cpu.Sub(twoCores)

	// 6 % of one core
	// + 1% of the next core
	// + 0.5% of the next 2 cores
	// + 0.25% of any cores above 4 cores
	cpu.SetMilli(int64(60 + 10 + 10 + float64(cpu.MilliValue())*0.0025))
	return cpu.String()
}

var (
	oneGi                   = resource.MustParse("1Gi")
	fourGi                  = resource.MustParse("4Gi")
	eightGi                 = resource.MustParse("8Gi")
	oneTwelveGi             = resource.MustParse("112Gi")
	twentyFivePercentFourGi = float64(fourGi.Value()) * 0.25
	twentyPercentFourGi     = float64(fourGi.Value()) * 0.20
	tenPercentEightGi       = float64(eightGi.Value()) * 0.10
	sixPercentOneTwelveGi   = float64(oneTwelveGi.Value()) * 0.06
)

const (
	KiB      = 1024
	Mebibyte = KiB * 1024
	KB       = 1000
	Megabyte = KB * 1000
)

// TODO: check against real memory consumption. CHECK ON RANDOM NODE FIRST, AND THEN SIMULATE 20,50,80,110 Pods
// use this formula, but also take the maxPods into consideration for memory consumption!
// check for how many pods those memory settings are reproducable on the node (and if they make even sense)
// then check how much more the container runtime and the kublet consume for memory if there are more pods deployed

// CalculateReservedMemory calculates a regressive rate of memory reservations
// 255 MiB of memory for machines with less than 1 GiB of memory
// 25% of the first 4 GiB of memory
// 20% of the next 4 GiB of memory (up to 8 GiB)
// 10% of the next 8 GiB of memory (up to 16 GiB)
// 6% of the next 112 GiB of memory (up to 128 GiB)
// 2% of any memory above 128 GiB
// 0.75 (memory.eviction - is 100 Mi for us) + (0.25*4) + (0.20*3) = 0.75GB + 1GB + 0.6GB = 2.35GB / 7GB = 33.57% reserved
// Also see: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#meaning-of-memory
func CalculateReservedMemory(memory resource.Quantity) string {
	// 255 MiB of memory for machines with less than 1 GiB of memory
	if oneGi.Cmp(memory) >= 0 {
		return "255Mi"
	}

	// 25% of the first 4 GiB of memory
	if fourGi.Cmp(memory) >= 0 {
		// equal or less than 8 Gi
		memory.Set(roundMemoryResource(int64(float64(memory.Value())*0.25), memory.Format))
		return memory.String()
	}

	memory.Sub(fourGi)

	// 20% of the next 4 GiB of memory (up to 8 GiB)
	if fourGi.Cmp(memory) >= 0 {
		// equal or less than 8 Gi
		// result is 25 % of 4 Gi + 20% of rest
		memory.Set(roundMemoryResource(int64(twentyFivePercentFourGi+float64(memory.Value())*0.20), memory.Format))
		return memory.String()
	}

	memory.Sub(fourGi)

	// 10% of the next 8 GiB of memory (up to 16 GiB)
	if eightGi.Cmp(memory) >= 0 {
		// equal or less than 16 Gi
		memory.Set(roundMemoryResource(int64(twentyFivePercentFourGi+twentyPercentFourGi+float64(memory.Value())*0.10), memory.Format))
		return memory.String()
	}

	memory.Sub(eightGi)

	// 6% of the next 112 GiB of memory (up to 128 GiB)
	if oneTwelveGi.Cmp(memory) >= 0 {
		// equal or less than 128 Gi
		memory.Set(roundMemoryResource(int64(twentyFivePercentFourGi+twentyPercentFourGi+tenPercentEightGi+float64(memory.Value())*0.06), memory.Format))
		return memory.String()
	}

	// 2% of any memory above 128 GiB
	memory.Sub(oneTwelveGi)
	memory.Set(roundMemoryResource(int64(twentyFivePercentFourGi+twentyPercentFourGi+tenPercentEightGi+sixPercentOneTwelveGi+float64(memory.Value())*0.02), memory.Format))
	return memory.String()
}

// roundMemoryResource rounds (floor) the given resource quantity to multiples of the resource.Format
// this is done for better readability and simplicity
// Example: 711.5Mi -> 711Mi.
// Without rounding, the kubelet config would contain instead: 746061 ki
func roundMemoryResource(v int64, format resource.Format) int64 {
	var remainder int64
	switch format {
	case resource.BinarySI:
		remainder = v % Mebibyte
		break
	case resource.DecimalSI, resource.DecimalExponent:
		remainder = v % Megabyte
		break
	default:
		// bytes
		return v
	}
	return v - remainder
}

var (
	sixGi        = resource.MustParse("6Gi")
	oneHundredGi = resource.MustParse("100Gi")
)

// CalculateReservedEphemeralStorage calculates ephemeral storage reservations based on the formula
// Min(50% * BOOT-DISK-CAPACITY, 6Gi + 35% * BOOT-DISK-CAPACITY, 100 Gi)
func CalculateReservedEphemeralStorage(bootDiskSize resource.Quantity) string {
	halfBootDiskSize := float64(bootDiskSize.Value()) * 0.5
	thirtyFivePercentBootDiskSize := float64(sixGi.Value()) + float64(bootDiskSize.Value())*0.35
	reserved := math.Min(math.Min(halfBootDiskSize, thirtyFivePercentBootDiskSize), float64(oneHundredGi.Value()))
	roundedReserved := roundMemoryResource(int64(reserved), bootDiskSize.Format)
	bootDiskSize.Set(roundedReserved)
	return bootDiskSize.String()
}
