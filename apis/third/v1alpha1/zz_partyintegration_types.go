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

type PartyIntegrationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PartyIntegrationParameters struct {

	// +kubebuilder:validation:Optional
	APIKeySecretRef *v1.SecretKeySelector `json:"apiKeySecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	APITokenSecretRef *v1.SecretKeySelector `json:"apiTokenSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// +kubebuilder:validation:Optional
	ChannelName *string `json:"channelName,omitempty" tf:"channel_name,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	FlowName *string `json:"flowName,omitempty" tf:"flow_name,omitempty"`

	// +kubebuilder:validation:Optional
	LicenseKeySecretRef *v1.SecretKeySelector `json:"licenseKeySecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	MicrosoftTeamsWebhookURLSecretRef *v1.SecretKeySelector `json:"microsoftTeamsWebhookUrlSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	OrgName *string `json:"orgName,omitempty" tf:"org_name,omitempty"`

	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-mongodbatlas/apis/mongodbatlas/v1alpha1.Project
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-jet-mongodbatlas/config/common.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a Project in mongodbatlas to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a Project in mongodbatlas to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ReadTokenSecretRef *v1.SecretKeySelector `json:"readTokenSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Optional
	RoutingKeySecretRef *v1.SecretKeySelector `json:"routingKeySecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Scheme *string `json:"scheme,omitempty" tf:"scheme,omitempty"`

	// +kubebuilder:validation:Optional
	SecretSecretRef *v1.SecretKeySelector `json:"secretSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ServiceDiscoverySecretRef *v1.SecretKeySelector `json:"serviceDiscoverySecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ServiceKeySecretRef *v1.SecretKeySelector `json:"serviceKeySecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	TeamName *string `json:"teamName,omitempty" tf:"team_name,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// +kubebuilder:validation:Optional
	UserNameSecretRef *v1.SecretKeySelector `json:"userNameSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	WriteTokenSecretRef *v1.SecretKeySelector `json:"writeTokenSecretRef,omitempty" tf:"-"`
}

// PartyIntegrationSpec defines the desired state of PartyIntegration
type PartyIntegrationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PartyIntegrationParameters `json:"forProvider"`
}

// PartyIntegrationStatus defines the observed state of PartyIntegration.
type PartyIntegrationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PartyIntegrationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PartyIntegration is the Schema for the PartyIntegrations API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,mongodbatlas}
type PartyIntegration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PartyIntegrationSpec   `json:"spec"`
	Status            PartyIntegrationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PartyIntegrationList contains a list of PartyIntegrations
type PartyIntegrationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PartyIntegration `json:"items"`
}

// Repository type metadata.
var (
	PartyIntegration_Kind             = "PartyIntegration"
	PartyIntegration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PartyIntegration_Kind}.String()
	PartyIntegration_KindAPIVersion   = PartyIntegration_Kind + "." + CRDGroupVersion.String()
	PartyIntegration_GroupVersionKind = CRDGroupVersion.WithKind(PartyIntegration_Kind)
)

func init() {
	SchemeBuilder.Register(&PartyIntegration{}, &PartyIntegrationList{})
}
