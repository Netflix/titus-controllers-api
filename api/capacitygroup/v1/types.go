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

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	poolV1 "github.com/Netflix/titus-controllers-api/api/resourcepool/v1"
)

// +genclient
// +resourceName=capacitygroups

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type CapacityGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec CapacityGroupSpec `json:"spec,omitempty"`
}

type CapacityGroupSpec struct {
	CapacityGroupName      string `json:"capacityGroupName"`
	ResourcePoolName       string `json:"resourcePoolName"`
	SchedulerName          string `json:"schedulerName"`
	CreatedBy              string `json:"createdBy"`
	InstanceCount          uint32 `json:"instanceCount"`
	poolV1.ComputeResource `json:"resourceDimensions"`
}

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type CapacityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CapacityGroup `json:"items"`
}
