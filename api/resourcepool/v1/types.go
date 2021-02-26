/*
Copyright 2020 Netflix, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +resourceName=resourcepoolconfigs

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ResourcePoolConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResourcePoolSpec `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ResourcePoolConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourcePoolConfig `json:"items"`
}

type ResourceShape struct {
	ComputeResource `json:",inline"`

	Labels map[string]string `json:"labels,omitempty"`
}

type ResourceDemandStatus struct {
	// Resource provisioning status (Pending, Rejected, Fulfilled)
	Status string `json:"status"`

	// Human readable message.
	StatusMessage string `json:"statusMessage"`

	// Time at which the report was made.
	ReportedAt int64 `json:"reportedAt"`
}

type ResourcePoolScalingRules struct {
	// Minimum idle capacity measured in number of shapes.
	MinIdle int64 `json:"minIdle"`

	// Maximum idle capacity measured in number of shapes.
	MaxIdle int64 `json:"maxIdle"`

	// Resource pool minimum size limit (minimum number of shapes that can be ever allocated).
	MinSize int64 `json:"minSize"`

	// Resource pool maximum size limit (maximum number of shapes that can be ever allocated).
	MaxSize int64 `json:"maxSize"`

	// Set to true to enable auto scaling of this resource pool.
	AutoScalingEnabled bool `json:"autoScalingEnabled"`
}

type ResourcePoolSpec struct {
	// Resource pool name (for example 'critical', 'elastic', 'gpu.p4', etc).
	Name string `json:"name"`

	// Preferential machine shape (minimum size and resource ratio). A total amount of requested resource is
	// defined as resource_shape * resource_count.
	ResourceShape ResourceShape `json:"resourceShape"`

	// Resource pool scaling rules.
	ScalingRules ResourcePoolScalingRules `json:"scalingRules"`

	// Number of resource shapes requested. A total amount of requested resource is
	// defined as resource_shape * resource_count.
	ResourceCount int64 `json:"resourceCount"`

	// Time at which the last resource request was made.
	RequestedAt int64 `json:"requestedAt"`

	// Resource demand fulfillment report.
	Status ResourceDemandStatus `json:"resourceDemandStatus"`
}
