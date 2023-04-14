/*
Copyright (C) 2021-2023, Kubefirst

This program is licensed under MIT.
See the LICENSE file for more details.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// WorkloadClusterSpec defines the desired state of WorkloadCluster
type WorkloadClusterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of WorkloadCluster. Edit workloadcluster_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// WorkloadClusterStatus defines the observed state of WorkloadCluster
type WorkloadClusterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// WorkloadCluster is the Schema for the workloadclusters API
type WorkloadCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WorkloadClusterSpec   `json:"spec,omitempty"`
	Status WorkloadClusterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// WorkloadClusterList contains a list of WorkloadCluster
type WorkloadClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorkloadCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&WorkloadCluster{}, &WorkloadClusterList{})
}
