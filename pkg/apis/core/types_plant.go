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

package v1alpha1

import (
	"github.com/gardener/gardener/pkg/apis/garden/v1beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Plant ...
type Plant struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object metadata.
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Spec contains the specification of this installation.
	Spec PlantSpec `json:"spec,omitempty"`
	// Status contains the status of this installation.
	Status PlantStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ControllerInstallationList is a collection of ControllerInstallations.
type PlantList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list object metadata.
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is the list of ControllerInstallations.
	Items []Plant `json:"items"`
}

// ControllerInstallationSpec is the specification of a ControllerInstallation.

/*
---
apiVersion: v1
kind: Secret
metadata:
  name: my-external-cluster
  namespace: garden-my-project
type: Opaque
data:
  kubeconfig: base64(kubeconfig-for-external-cluster)

---
apiVersion: core.gardener.cloud/v1alpha1
kind: Plant
metadata:
  name: my-external-cluster
  namespace: garden-my-project
spec:
  secretRef:
    name: my-external-cluster
  # optional fields
  monitoring:
    endpoints:
    - name: Kubernetes Dashboard
      url: https://...
    - name: Prometheus
      url: https://...
  logging:
    endpoints:
    - name: Kibana
      url: https://...
 */
type PlantSpec struct {
	// SecretRef is a reference to a Secret object containing the Kubeconfig and the cloud provider credentials for
	// the non-gardener kubernetes cluster.
	SecretRef corev1.SecretReference `json:"secretRef"`

	// +optional
	Monitoring Endpoints `json:"monitoring,omitempty"`

	// +optional
	Logging Endpoints `json:"logging,omitempty"`
}

type Endpoints struct {
	Endpoints []Endpoint `json:"endpoints"`
}

type Endpoint struct{
	// Name is unique within a namespace to reference a secret resource.
	// +optional
	Name string `json:"name"`
	// Namespace defines the space within which the secret name must be unique.
	// +optional
	Url string `json:"url"`
}



/*
status:
  conditions:
    - type: APIServerAvailable
      status: 'True'
      lastTransitionTime: '2019-03-02T13:00:52Z'
      lastUpdateTime: '2019-03-02T14:45:28Z'
      reason: HealthzRequestSucceeded
      message: >-
        API server /healthz endpoint responded with success status code. [response_time:92ms]
    - type: EveryNodeReady
      status: 'True'
      lastTransitionTime: '2019-03-02T13:35:32Z'
      lastUpdateTime: '2019-03-02T14:45:28Z'
      reason: EveryNodeReady
      message: Every node registered to the cluster is ready
  observedGeneration: 4
  clusterInfo:
    cloud:
      type: aws|azure|gcp|openstack|alicloud|unknown
      region: eu-central-1
    kubernetes:
      version: 1.12.3
 */
type PlantStatus struct {
	// Conditions represents the latest available observations of a ControllerInstallations's current state.
	// +optional
	Conditions []Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the most recent generation observed for this Plant. It corresponds to the
	// Plant's generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	ClusterInfo ClusterInfo `json:"clusterInfo"`
}

type ClusterInfo struct {
	Cloud Cloud `json:"cloud"`
	// Kubernetes contains the version and configuration settings of the control plane components.
	Kubernetes v1beta1.Kubernetes `json:"kubernetes"`
}

type Cloud struct {
	// e.g v1beta1.CloudProviderAWS
	Type *string `json:"type"`
	Region *string `json:"region"`
}