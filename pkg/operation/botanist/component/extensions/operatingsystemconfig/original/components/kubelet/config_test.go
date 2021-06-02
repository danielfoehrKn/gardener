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

package kubelet_test

import (
	"time"

	gardencorev1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	"github.com/gardener/gardener/pkg/operation/botanist/component/extensions/operatingsystemconfig/original/components"
	"github.com/gardener/gardener/pkg/operation/botanist/component/extensions/operatingsystemconfig/original/components/kubelet"
	"github.com/gardener/gardener/pkg/utils"
	"k8s.io/apimachinery/pkg/api/resource"

	"github.com/Masterminds/semver"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubeletconfigv1beta1 "k8s.io/kubelet/config/v1beta1"
	"k8s.io/utils/pointer"
)

var _ = Describe("Config", func() {
	var (
		clusterDNSAddress = "foo"
		clusterDomain     = "bar"
		params            = components.ConfigurableKubeletConfigParameters{
			CpuCFSQuota:                      pointer.BoolPtr(false),
			CpuManagerPolicy:                 pointer.StringPtr("policy"),
			EvictionHard:                     map[string]string{"memory.available": "123Mi"},
			EvictionPressureTransitionPeriod: &metav1.Duration{Duration: 42 * time.Minute},
			EvictionMaxPodGracePeriod:        pointer.Int32Ptr(120),
			FailSwapOn:                       pointer.BoolPtr(false),
			FeatureGates:                     map[string]bool{"Foo": false},
			KubeReserved:                     map[string]string{"cpu": "123"},
			MaxPods:                          pointer.Int32Ptr(24),
			PodPidsLimit:                     pointer.Int64Ptr(101),
			SystemReserved:                   map[string]string{"memory": "321"},
		}

		kubeletConfigWithDefaults = &kubeletconfigv1beta1.KubeletConfiguration{
			Authentication: kubeletconfigv1beta1.KubeletAuthentication{
				Anonymous: kubeletconfigv1beta1.KubeletAnonymousAuthentication{
					Enabled: pointer.BoolPtr(false),
				},
				X509: kubeletconfigv1beta1.KubeletX509Authentication{
					ClientCAFile: "/var/lib/kubelet/ca.crt",
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
			CgroupDriver:                 "cgroupfs",
			CgroupRoot:                   "/",
			CgroupsPerQOS:                pointer.BoolPtr(true),
			ClusterDNS:                   []string{clusterDNSAddress},
			ClusterDomain:                clusterDomain,
			CPUCFSQuota:                  pointer.BoolPtr(true),
			CPUManagerPolicy:             "none",
			CPUManagerReconcilePeriod:    metav1.Duration{Duration: 10 * time.Second},
			EnableControllerAttachDetach: pointer.BoolPtr(true),
			EnableDebuggingHandlers:      pointer.BoolPtr(true),
			EnableServer:                 pointer.BoolPtr(true),
			EnforceNodeAllocatable:       []string{"pods"},
			EventBurst:                   50,
			EventRecordQPS:               pointer.Int32Ptr(50),
			EvictionHard: map[string]string{
				"memory.available":   "100Mi",
				"imagefs.available":  "15%",
				"imagefs.inodesFree": "5%",
				"nodefs.available":   "10%",
				"nodefs.inodesFree":  "5%",
				"pid.available":  "10%",
			},
			EvictionPressureTransitionPeriod: metav1.Duration{Duration: 4 * time.Minute},
			EvictionMaxPodGracePeriod:        90,
			FailSwapOn:                       pointer.BoolPtr(true),
			FileCheckFrequency:               metav1.Duration{Duration: 20 * time.Second},
			HairpinMode:                      kubeletconfigv1beta1.PromiscuousBridge,
			HTTPCheckFrequency:               metav1.Duration{Duration: 20 * time.Second},
			ImageGCHighThresholdPercent:      pointer.Int32Ptr(50),
			ImageGCLowThresholdPercent:       pointer.Int32Ptr(40),
			ImageMinimumGCAge:                metav1.Duration{Duration: 2 * time.Minute},
			KubeAPIBurst:                     50,
			KubeAPIQPS:                       pointer.Int32Ptr(50),
			// default kube-reserved if both the machine type and the root volume are not set
			KubeReserved: map[string]string{
				"cpu":    "80m",
				"memory": "1Gi",
				"pid": "2048",
			},
			MaxOpenFiles:              1000000,
			MaxPods:                   110,
			NodeStatusUpdateFrequency: metav1.Duration{Duration: 10 * time.Second},
			PodsPerCore:               0,
			ReadOnlyPort:              0,
			RegistryBurst:             10,
			RegistryPullQPS:           pointer.Int32Ptr(5),
			ResolverConfig:            "/etc/resolv.conf",
			RuntimeRequestTimeout:     metav1.Duration{Duration: 2 * time.Minute},
			SerializeImagePulls:       pointer.BoolPtr(true),
			SyncFrequency:             metav1.Duration{Duration: time.Minute},
			VolumeStatsAggPeriod:      metav1.Duration{Duration: time.Minute},
		}

		kubeletConfigWithParams = &kubeletconfigv1beta1.KubeletConfiguration{
			Authentication: kubeletconfigv1beta1.KubeletAuthentication{
				Anonymous: kubeletconfigv1beta1.KubeletAnonymousAuthentication{
					Enabled: pointer.BoolPtr(false),
				},
				X509: kubeletconfigv1beta1.KubeletX509Authentication{
					ClientCAFile: "/var/lib/kubelet/ca.crt",
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
			CgroupDriver:                 "cgroupfs",
			CgroupRoot:                   "/",
			CgroupsPerQOS:                pointer.BoolPtr(true),
			ClusterDomain:                clusterDomain,
			ClusterDNS:                   []string{clusterDNSAddress},
			CPUCFSQuota:                  params.CpuCFSQuota,
			CPUManagerPolicy:             *params.CpuManagerPolicy,
			CPUManagerReconcilePeriod:    metav1.Duration{Duration: 10 * time.Second},
			EnableControllerAttachDetach: pointer.BoolPtr(true),
			EnableDebuggingHandlers:      pointer.BoolPtr(true),
			EnableServer:                 pointer.BoolPtr(true),
			EnforceNodeAllocatable:       []string{"pods"},
			EventBurst:                   50,
			EventRecordQPS:               pointer.Int32Ptr(50),
			EvictionHard: utils.MergeStringMaps(params.EvictionHard, map[string]string{
				"imagefs.available":  "15%",
				"imagefs.inodesFree": "5%",
				"nodefs.available":   "10%",
				"nodefs.inodesFree":  "5%",
				"pid.available":  "10%",
			}),
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
			KubeReserved:                     utils.MergeStringMaps(params.KubeReserved,
												map[string]string{"memory": "1Gi", "pid": "2048"}),
			MaxOpenFiles:                     1000000,
			MaxPods:                          *params.MaxPods,
			NodeStatusUpdateFrequency:        metav1.Duration{Duration: 10 * time.Second},
			PodsPerCore:                      0,
			PodPidsLimit:                     params.PodPidsLimit,
			ReadOnlyPort:                     0,
			RegistryBurst:                    10,
			RegistryPullQPS:                  pointer.Int32Ptr(5),
			ResolverConfig:                   "/etc/resolv.conf",
			RuntimeRequestTimeout:            metav1.Duration{Duration: 2 * time.Minute},
			SerializeImagePulls:              pointer.BoolPtr(true),
			SyncFrequency:                    metav1.Duration{Duration: time.Minute},
			SystemReserved:                   params.SystemReserved,
			VolumeStatsAggPeriod:             metav1.Duration{Duration: time.Minute},
		}
	)

	fiftyGi := resource.MustParse("50Gi")
	twentyGi := resource.MustParse("20Gi")

	DescribeTable("#Config",
		func(kubernetesVersion string, clusterDNSAddress, clusterDomain string, machineType *gardencorev1beta1.MachineType, volume *gardencorev1beta1.Volume,params components.ConfigurableKubeletConfigParameters, expectedConfig *kubeletconfigv1beta1.KubeletConfiguration, mutateExpectConfigFn func(*kubeletconfigv1beta1.KubeletConfiguration)) {
			expectation := expectedConfig.DeepCopy()
			if mutateExpectConfigFn != nil {
				mutateExpectConfigFn(expectation)
			}

			Expect(kubelet.Config(semver.MustParse(kubernetesVersion), clusterDNSAddress, clusterDomain, machineType, volume, params)).To(Equal(expectation))
		},

		Entry(
			"kube-reserved based on the machine type - root disk size from volume",
			"1.15.1",
			clusterDNSAddress,
			clusterDomain,
			&gardencorev1beta1.MachineType{
				CPU:     resource.MustParse("64"),
				Memory:  resource.MustParse("128Gi"),
			},
			&gardencorev1beta1.Volume{
				VolumeSize: "50Gi",
			},
			components.ConfigurableKubeletConfigParameters{},
			kubeletConfigWithDefaults,
			func(cfg *kubeletconfigv1beta1.KubeletConfiguration) {
				cfg.RotateCertificates = true
				cfg.KubeReserved = map[string]string{
				"memory": "9543Mi",
				"pid": "2048",
				"ephemeral-storage": "24064Mi",
				"cpu":    "230m",
			} },
		),
		Entry(
			"kube-reserved based on the machine type - no root disk size given",
			"1.15.1",
			clusterDNSAddress,
			clusterDomain,
			&gardencorev1beta1.MachineType{
				CPU:     resource.MustParse("64"),
				Memory:  resource.MustParse("128Gi"),
			},
			nil,
			components.ConfigurableKubeletConfigParameters{},
			kubeletConfigWithDefaults,
			func(cfg *kubeletconfigv1beta1.KubeletConfiguration) {
				cfg.RotateCertificates = true
				cfg.KubeReserved = map[string]string{
					"memory": "9543Mi",
					"pid": "2048",
					"cpu":    "230m",
				} },
		),
		Entry(
			"kube-reserved based on the machine type - root disk size from machine type",
			"1.15.1",
			clusterDNSAddress,
			clusterDomain,
			&gardencorev1beta1.MachineType{
				CPU:     resource.MustParse("64"),
				Memory:  resource.MustParse("128Gi"),
				Storage: &gardencorev1beta1.MachineTypeStorage{
					StorageSize: &fiftyGi,
				},
			},
			nil,
			components.ConfigurableKubeletConfigParameters{},
			kubeletConfigWithDefaults,
			func(cfg *kubeletconfigv1beta1.KubeletConfiguration) {
				cfg.RotateCertificates = true
				cfg.KubeReserved = map[string]string{
					"memory": "9543Mi",
					"pid": "2048",
					"ephemeral-storage": "24064Mi",
					"cpu":    "230m",
				} },
		),
		Entry(
			"kube-reserved based on the machine type - root disk size from worker volume overwrites the disk size from the machine type",
			"1.15.1",
			clusterDNSAddress,
			clusterDomain,
			&gardencorev1beta1.MachineType{
				CPU:     resource.MustParse("64"),
				Memory:  resource.MustParse("128Gi"),
				Storage: &gardencorev1beta1.MachineTypeStorage{
					StorageSize: &twentyGi,
				},
			},
			&gardencorev1beta1.Volume{
				VolumeSize: "50Gi",
			},
			components.ConfigurableKubeletConfigParameters{},
			kubeletConfigWithDefaults,
			func(cfg *kubeletconfigv1beta1.KubeletConfiguration) {
				cfg.RotateCertificates = true
				cfg.KubeReserved = map[string]string{
					"memory": "9543Mi",
					"pid": "2048",
					"ephemeral-storage": "24064Mi",
					"cpu":    "230m",
				} },
		),
		Entry(
			"kubernetes 1.15 w/o defaults",
			"1.15.1",
			clusterDNSAddress,
			clusterDomain,
			nil,
			nil,
			components.ConfigurableKubeletConfigParameters{},
			kubeletConfigWithDefaults,
			func(cfg *kubeletconfigv1beta1.KubeletConfiguration) { cfg.RotateCertificates = true },
		),
		Entry(
			"kubernetes 1.15 w/ defaults",
			"1.15.1",
			clusterDNSAddress,
			clusterDomain,
			nil,
			nil,
			params,
			kubeletConfigWithParams,
			func(cfg *kubeletconfigv1beta1.KubeletConfiguration) { cfg.RotateCertificates = true },
		),

		Entry(
			"kubernetes 1.16 w/o defaults",
			"1.16.1",
			clusterDNSAddress,
			clusterDomain,
			nil,
			nil,
			components.ConfigurableKubeletConfigParameters{},
			kubeletConfigWithDefaults,
			func(cfg *kubeletconfigv1beta1.KubeletConfiguration) { cfg.RotateCertificates = true },
		),
		Entry(
			"kubernetes 1.16 w/ defaults",
			"1.16.1",
			clusterDNSAddress,
			clusterDomain,
			nil,
			nil,
			params,
			kubeletConfigWithParams,
			func(cfg *kubeletconfigv1beta1.KubeletConfiguration) { cfg.RotateCertificates = true },
		),

		Entry(
			"kubernetes 1.17 w/o defaults",
			"1.17.1",
			clusterDNSAddress,
			clusterDomain,
			nil,
			nil,
			components.ConfigurableKubeletConfigParameters{},
			kubeletConfigWithDefaults,
			func(cfg *kubeletconfigv1beta1.KubeletConfiguration) { cfg.RotateCertificates = true },
		),
		Entry(
			"kubernetes 1.17 w/ defaults",
			"1.17.1",
			clusterDNSAddress,
			clusterDomain,
			nil,
			nil,
			params,
			kubeletConfigWithParams,
			func(cfg *kubeletconfigv1beta1.KubeletConfiguration) { cfg.RotateCertificates = true },
		),

		Entry(
			"kubernetes 1.18 w/o defaults",
			"1.18.1",
			clusterDNSAddress,
			clusterDomain,
			nil,
			nil,
			components.ConfigurableKubeletConfigParameters{},
			kubeletConfigWithDefaults,
			func(cfg *kubeletconfigv1beta1.KubeletConfiguration) { cfg.RotateCertificates = true },
		),
		Entry(
			"kubernetes 1.18 w/ defaults",
			"1.18.1",
			clusterDNSAddress,
			clusterDomain,
			nil,
			nil,
			params,
			kubeletConfigWithParams,
			func(cfg *kubeletconfigv1beta1.KubeletConfiguration) { cfg.RotateCertificates = true },
		),

		Entry(
			"kubernetes 1.19 w/o defaults",
			"1.19.1",
			clusterDNSAddress,
			clusterDomain,
			nil,
			nil,
			components.ConfigurableKubeletConfigParameters{},
			kubeletConfigWithDefaults,
			func(cfg *kubeletconfigv1beta1.KubeletConfiguration) {
				cfg.RotateCertificates = true
				cfg.VolumePluginDir = "/var/lib/kubelet/volumeplugins"
			},
		),
		Entry(
			"kubernetes 1.19 w/ defaults",
			"1.19.1",
			clusterDNSAddress,
			clusterDomain,
			nil,
			nil,
			params,
			kubeletConfigWithParams,
			func(cfg *kubeletconfigv1beta1.KubeletConfiguration) {
				cfg.RotateCertificates = true
				cfg.VolumePluginDir = "/var/lib/kubelet/volumeplugins"
			},
		),

		Entry(
			"kubernetes 1.20 w/o defaults",
			"1.20.1",
			clusterDNSAddress,
			clusterDomain,
			nil,
			nil,
			components.ConfigurableKubeletConfigParameters{},
			kubeletConfigWithDefaults,
			func(cfg *kubeletconfigv1beta1.KubeletConfiguration) {
				cfg.RotateCertificates = true
				cfg.VolumePluginDir = "/var/lib/kubelet/volumeplugins"
			},
		),
		Entry(
			"kubernetes 1.20 w/ defaults",
			"1.20.1",
			clusterDNSAddress,
			clusterDomain,
			nil,
			nil,
			params,
			kubeletConfigWithParams,
			func(cfg *kubeletconfigv1beta1.KubeletConfiguration) {
				cfg.RotateCertificates = true
				cfg.VolumePluginDir = "/var/lib/kubelet/volumeplugins"
			},
		),
	)

	DescribeTable("#CalculateReservedEphemeralStorage",
		func(value, expected string) {
			res := resource.MustParse(value)
			result := kubelet.CalculateReservedEphemeralStorage(res)
			Expect(result).To(Equal(expected))
		},
		Entry(
			"should reserve 50% * BOOT-DISK-CAPACITY",
			"20Gi",
			"10Gi",
		),
		Entry(
			"should reserve 6Gi + 35% * BOOT-DISK-CAPACITY",
			"50Gi",
			// 35% of 50Gi = 17920Mi
			// 6Gi = 6144Mi
			"24064Mi",
		),
		Entry(
			"should reserve 100Gi",
			"1000Gi",
			"100Gi",
		),
	)


	DescribeTable("#CalculateReservedCPU",
		func(value, expected string) {
			res := resource.MustParse(value)
			result := kubelet.CalculateReservedCPU(res)
			Expect(result).To(Equal(expected))
		},
		Entry(
			"should reserve 6% of the first core - full cores",
			"1",
			"60m",
		),
		Entry(
			"should reserve 6% of the first core - fractional",
			"0.5",
			"30m",
		),
		Entry(
			"should reserve 6% of the first core - high precision",
			"0.5532",
			"33m",
		),
		Entry(
			"should reserve 6% of the first core - millicpu",
			"800m",
			"48m",
		),
		Entry(
			"should reserve 1% of the next core (up to 2 cores) - full cores",
			"2",
			"70m",
		),
		Entry(
			"should reserve 1% of the next core (up to 2 cores) - fractional",
			"1.5",
			"65m",
		),
		Entry(
			"should reserve 1% of the next core (up to 2 cores) - high precision",
			"1.614",
			"66m",
		),
		Entry(
			"should reserve 1% of the next core (up to 2 cores) - millicpu",
			"1800m",
			"68m",
		),
		Entry(
			"should reserve 0.5% of the next 2 cores (up to 4 cores) - full cores",
			"4",
			"80m",
		),
		Entry(
			"should reserve 0.5% of the next 2 cores (up to 4 cores) - fractional",
			"3.5",
			"77m",
		),
		Entry(
			"should reserve 0.5% of the next 2 cores (up to 4 cores) - high precision",
			"3.614",
			"78m",
		),
		Entry(
			"should reserve 0.5% of the next 2 cores (up to 4 cores) - millicpu",
			"3900m",
			"79m",
		),
		// 0.25% of any cores above 4 core
		Entry(
			"should reserve 0.25% of any cores above 4 core - full cores (8)",
			"8",
			// 60m + (1000m * 0.01) + (2000m * 0.005)  + (4000m * 0.0025)
			// 60m + 10m + 10m + 10 m = 90m
			"90m",
		),
		Entry(
			"should reserve 0.25% of any cores above 4 core - full cores (32)",
			"32",
			"150m",
		),
		Entry(
			"should reserve 0.25% of any cores above 4 core - full cores (64)",
			"64",
			// 60m + (1000m * 0.01) + (2000m * 0.005)  + (60000m * 0.0025)
			// 60m + 10m + 10m + 150 m = 230m
			"230m",
		),
		Entry(
			"should reserve 0.25% of any cores above 4 core - fractional",
			"60.5",
			"221m",
		),
		Entry(
			"should reserve 0.25% of any cores above 4 core - high precision",
			"8.614",
			"91m",
		),
		Entry(
			"should reserve 0.25% of any cores above 4 core - millicpu",
			"16900m",
			"112m",
		),
		Entry(
			"should reserve 0.25% of any cores above 4 core - high cores",
			"1000",
			"2570m",
		),
		)


	// 1 Mi = 1048576 bytes (megabinary): 2 to the power of 20
	// 1 M = 1000000 bytes (megabyte): 10 to the power of 6
	DescribeTable("#CalculateReservedMemory",
		func(value, expected string) {
			res := resource.MustParse(value)
			result := kubelet.CalculateReservedMemory(res)
			Expect(result).To(Equal(expected))
		},
		Entry(
			"should reserve 255 MiB of memory for machines with less than 1 GiB of memory - BinarySI",
			"500Mi",
			"255Mi",
		),
		Entry(
			"should reserve 255 MiB of memory for machines with less than 1 GiB of memory - BinarySI (1Gi)",
			"1Gi",
			"255Mi",
		),
		Entry(
			"should reserve 255 MiB of memory for machines with less than 1 GiB of memory - DecimalSI",
			"1G",
			"255Mi",
		),
		Entry(
			"should reserve 255 MiB of memory for machines with less than 1 GiB of memory - DecimalExponent",
			"500e6",
			"255Mi",
		),
		Entry(
			"should reserve 25% of the first 4 GiB of memory - BinarySI (4 Gi)",
			"4Gi",
			"1Gi",
		),
		Entry(
			"should reserve 25% of the first 4 GiB of memory - BinarySI (2 Gi)",
			"2Gi",
			"512Mi",
		),
		Entry(
			"should reserve 25% of the first 4 GiB of memory - DecimalSI",
			"2G",
			"500M",
		),
		Entry(
			"should reserve 25% of the first 4 GiB of memory - BinarySI - should round to full Mebibyte",
			"1481Mi",
			"370Mi",
		),
		Entry(
			"should reserve 25% of the first 4 GiB of memory - DecimalSI - should round to full Megabytes",
			"1481M",
			"370M",
		),
		Entry(
			"should reserve 25% of the first 4 GiB of memory - Decimal exponential",
			"2e9",
			"500e6",
		),
		Entry(
			"should reserve 20% of the next 4 GiB of memory (up to 8 GiB) - BinarySI",
			"8Gi",
			// 25 % of 4 Gi = 1024 Mi
			// 20 % of 4 Gi = 819.2 Mi
			"1843Mi",
		),
		Entry(
			"should reserve 20% of the next 4 GiB of memory (up to 8 GiB) - DecimalSI",
			"8G",
			// First 4 Gi = 4294.97 M
			// 25 % of 4 Gi = 1024 Mi = 1073.74 M
			// 8 GB - 4 Gi are left = 3705 M
			// 20 % of 3705 M = 741 M
			// Total: 1073.74 M + 741 M = 1814.7 M
			"1814M",
		),
		Entry(
			"should reserve 20% of the next 4 GiB of memory (up to 8 GiB) - Decimal exponential",
			"8e9",
			"1814e6",
		),
		Entry(
			"should reserve 10% of the next 8 GiB of memory (up to 16 GiB) - BinarySI",
			"16Gi",
			// 25 % of 4 Gi = 1024 Mi
			// 20 % of 4 Gi = 819.2 Mi
			// 10% of 8 Gi = 819 Mi
			"2662Mi",
		),
		Entry(
			"should reserve 10% of the next 8 GiB of memory (up to 16 GiB) - BinarySI (1350Mi)",
			"14200Mi",
			// 25 % of 4 Gi = 1024 Mi
			// 20 % of 4 Gi = 819.2 Mi
			// remaining: 14200Mi - 8192Mi = 6008 Mi
			// 10% of 6008 Mi = 600.8 Mi
			"2444Mi",
		),
		Entry(
			"should reserve 10% of the next 8 GiB of memory (up to 16 GiB) - DecimalSI",
			"16G",
			// 25 % of 4 Gi = 1024 Mi = 1073.74 M
			// 20 % of 4 Gi = 819.2 Mi = 859 M
			// 16 Gb - 8 Gi (8590 M) = 7410 M left
			// 10% of 7410 M = 741 M
			"2673M",
		),
		Entry(
			"should reserve 10% of the next 8 GiB of memory (up to 16 GiB) - Decimal exponential",
			"16e9",
			"2673e6",
		),
		Entry(
			"should reserve 6% of the next 112 GiB of memory (up to 128 GiB) - BinarySI",
			"128Gi",
			// 25 % of 4 Gi = 1024 Mi
			// 20 % of 4 Gi = 819.2 Mi
			// 10% of 8 Gi = 819 Mi
			// 6% of 112 Gi = 6881 Mi
			"9543Mi",
		),
		Entry(
			"should reserve 6% of the next 112 GiB of memory (up to 128 GiB) - DecimalSI",
			"128G",
			"9440M",
		),
		Entry(
			"should reserve 6% of the next 112 GiB of memory (up to 128 GiB) - Decimal exponential",
			"128e9",
			"9440e6",
		),
		Entry(
			"2% of any memory above 128 GiB - BinarySI",
			"512Gi",
			// 25 % of 4 Gi = 1024 Mi
			// 20 % of 4 Gi = 819.2 Mi
			// 10% of 8 Gi = 819 Mi
			// 6% of 112 Gi = 6881 Mi
			// 2% of remaining 384Gi (393216 Mi) = 7864Mi
			// Total: 17407 == 17 Gi
			"17Gi",
		),
		Entry(
			"2% of any memory above 128 GiB - DecimalSI",
			"512G",
			"17498M",
		),
		Entry(
			"2% of any memory above 128 GiB - Decimal exponential",
			"512e9",
			"17498e6",
		),
	)
})
