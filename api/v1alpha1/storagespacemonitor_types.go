/*
Copyright 2025.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// StorageSpaceMonitorSpec defines the desired state of StorageSpaceMonitor.
type StorageSpaceMonitorSpec struct {
	// Pools is a list of storage pool names to monitor
	Pools []string `json:"pools,omitempty"`
}

// StorageSpaceMonitorStatus defines the observed state of StorageSpaceMonitor.
// PoolStatus represents the health status of a storage pool
type PoolStatus struct {
	FriendlyName string `json:"friendlyName"`
	HealthStatus string `json:"healthStatus"`
}

// StorageSpaceMonitorStatus defines the observed state of StorageSpaceMonitor.
type StorageSpaceMonitorStatus struct {
	// PoolStatuses contains the health status of monitored pools
	PoolStatuses []PoolStatus `json:"poolStatuses,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// StorageSpaceMonitor is the Schema for the storagespacemonitors API.
type StorageSpaceMonitor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StorageSpaceMonitorSpec   `json:"spec,omitempty"`
	Status StorageSpaceMonitorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StorageSpaceMonitorList contains a list of StorageSpaceMonitor.
type StorageSpaceMonitorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageSpaceMonitor `json:"items"`
}

func init() {
	SchemeBuilder.Register(&StorageSpaceMonitor{}, &StorageSpaceMonitorList{})
}
