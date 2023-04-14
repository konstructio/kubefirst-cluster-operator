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

// ManagementClusterSpec defines the desired state of ManagementCluster
type ManagementClusterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// CloudProvider specifies the cloud on which the cluster will be created
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	CloudProvider string `json:"cloud_provider"`
	// CloudRegion specifies the region within the cloud provider in which the cluster will be created
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	CloudRegion string `json:"cloud_region"`
	// DomainName specifies the domain name for DNS
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	DomainName string `json:"domain_name"`
	// AlertsEmail specifies the contact email for the cluster
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	AlertsEmail string `json:"alerts_email"`
}

// ManagementClusterStatus defines the observed state of ManagementCluster
type ManagementClusterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Conditions store the status conditions of the cluster instances
	// +operator-sdk:csv:customresourcedefinitions:type=status
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ManagementCluster is the Schema for the managementclusters API
// +kubebuilder:subresource:status
type ManagementCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ManagementClusterSpec   `json:"spec,omitempty"`
	Status ManagementClusterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ManagementClusterList contains a list of ManagementCluster
type ManagementClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagementCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ManagementCluster{}, &ManagementClusterList{})
}
