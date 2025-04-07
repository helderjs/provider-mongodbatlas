/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ListAPIKeyInitParameters struct {
	APIKeyID *string `json:"apiKeyId,omitempty" tf:"api_key_id,omitempty"`

	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`
}

type ListAPIKeyObservation struct {
	APIKeyID *string `json:"apiKeyId,omitempty" tf:"api_key_id,omitempty"`

	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`
}

type ListAPIKeyParameters struct {

	// +kubebuilder:validation:Optional
	APIKeyID *string `json:"apiKeyId,omitempty" tf:"api_key_id,omitempty"`

	// +kubebuilder:validation:Optional
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// +kubebuilder:validation:Optional
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`
}

// ListAPIKeySpec defines the desired state of ListAPIKey
type ListAPIKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ListAPIKeyParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ListAPIKeyInitParameters `json:"initProvider,omitempty"`
}

// ListAPIKeyStatus defines the observed state of ListAPIKey.
type ListAPIKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ListAPIKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ListAPIKey is the Schema for the ListAPIKeys API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,mongodbatlas}
type ListAPIKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.apiKeyId) || (has(self.initProvider) && has(self.initProvider.apiKeyId))",message="spec.forProvider.apiKeyId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.orgId) || (has(self.initProvider) && has(self.initProvider.orgId))",message="spec.forProvider.orgId is a required parameter"
	Spec   ListAPIKeySpec   `json:"spec"`
	Status ListAPIKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ListAPIKeyList contains a list of ListAPIKeys
type ListAPIKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ListAPIKey `json:"items"`
}

// Repository type metadata.
var (
	ListAPIKey_Kind             = "ListAPIKey"
	ListAPIKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ListAPIKey_Kind}.String()
	ListAPIKey_KindAPIVersion   = ListAPIKey_Kind + "." + CRDGroupVersion.String()
	ListAPIKey_GroupVersionKind = CRDGroupVersion.WithKind(ListAPIKey_Kind)
)

func init() {
	SchemeBuilder.Register(&ListAPIKey{}, &ListAPIKeyList{})
}
