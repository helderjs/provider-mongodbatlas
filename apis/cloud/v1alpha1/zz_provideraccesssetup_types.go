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

type AwsConfigInitParameters struct {
}

type AwsConfigObservation struct {
	AtlasAssumedRoleExternalID *string `json:"atlasAssumedRoleExternalId,omitempty" tf:"atlas_assumed_role_external_id,omitempty"`

	AtlasAwsAccountArn *string `json:"atlasAwsAccountArn,omitempty" tf:"atlas_aws_account_arn,omitempty"`
}

type AwsConfigParameters struct {
}

type ProviderAccessSetupInitParameters struct {

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-mongodbatlas/apis/mongodbatlas/v1alpha1.Project
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-mongodbatlas/config/common.ExtractResourceID()
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a Project in mongodbatlas to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a Project in mongodbatlas to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`

	ProviderName *string `json:"providerName,omitempty" tf:"provider_name,omitempty"`
}

type ProviderAccessSetupObservation struct {

	// +mapType=granular
	Aws map[string]*string `json:"aws,omitempty" tf:"aws,omitempty"`

	AwsConfig []AwsConfigObservation `json:"awsConfig,omitempty" tf:"aws_config,omitempty"`

	CreatedDate *string `json:"createdDate,omitempty" tf:"created_date,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	ProviderName *string `json:"providerName,omitempty" tf:"provider_name,omitempty"`

	RoleID *string `json:"roleId,omitempty" tf:"role_id,omitempty"`
}

type ProviderAccessSetupParameters struct {

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-mongodbatlas/apis/mongodbatlas/v1alpha1.Project
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-mongodbatlas/config/common.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a Project in mongodbatlas to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a Project in mongodbatlas to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ProviderName *string `json:"providerName,omitempty" tf:"provider_name,omitempty"`
}

// ProviderAccessSetupSpec defines the desired state of ProviderAccessSetup
type ProviderAccessSetupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProviderAccessSetupParameters `json:"forProvider"`
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
	InitProvider ProviderAccessSetupInitParameters `json:"initProvider,omitempty"`
}

// ProviderAccessSetupStatus defines the observed state of ProviderAccessSetup.
type ProviderAccessSetupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProviderAccessSetupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ProviderAccessSetup is the Schema for the ProviderAccessSetups API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,mongodbatlas}
type ProviderAccessSetup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.providerName) || (has(self.initProvider) && has(self.initProvider.providerName))",message="spec.forProvider.providerName is a required parameter"
	Spec   ProviderAccessSetupSpec   `json:"spec"`
	Status ProviderAccessSetupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProviderAccessSetupList contains a list of ProviderAccessSetups
type ProviderAccessSetupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProviderAccessSetup `json:"items"`
}

// Repository type metadata.
var (
	ProviderAccessSetup_Kind             = "ProviderAccessSetup"
	ProviderAccessSetup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProviderAccessSetup_Kind}.String()
	ProviderAccessSetup_KindAPIVersion   = ProviderAccessSetup_Kind + "." + CRDGroupVersion.String()
	ProviderAccessSetup_GroupVersionKind = CRDGroupVersion.WithKind(ProviderAccessSetup_Kind)
)

func init() {
	SchemeBuilder.Register(&ProviderAccessSetup{}, &ProviderAccessSetupList{})
}
