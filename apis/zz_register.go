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

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha1 "github.com/crossplane-contrib/provider-jet-mongodbatlas/apis/access/v1alpha1"
	v1alpha1alert "github.com/crossplane-contrib/provider-jet-mongodbatlas/apis/alert/v1alpha1"
	v1alpha1api "github.com/crossplane-contrib/provider-jet-mongodbatlas/apis/api/v1alpha1"
	v1alpha1backup "github.com/crossplane-contrib/provider-jet-mongodbatlas/apis/backup/v1alpha1"
	v1alpha1cloud "github.com/crossplane-contrib/provider-jet-mongodbatlas/apis/cloud/v1alpha1"
	v1alpha1custom "github.com/crossplane-contrib/provider-jet-mongodbatlas/apis/custom/v1alpha1"
	v1alpha1data "github.com/crossplane-contrib/provider-jet-mongodbatlas/apis/data/v1alpha1"
	v1alpha2 "github.com/crossplane-contrib/provider-jet-mongodbatlas/apis/database/v1alpha2"
	v1alpha1event "github.com/crossplane-contrib/provider-jet-mongodbatlas/apis/event/v1alpha1"
	v1alpha1federated "github.com/crossplane-contrib/provider-jet-mongodbatlas/apis/federated/v1alpha1"
	v1alpha1global "github.com/crossplane-contrib/provider-jet-mongodbatlas/apis/global/v1alpha1"
	v1alpha1ldap "github.com/crossplane-contrib/provider-jet-mongodbatlas/apis/ldap/v1alpha1"
	v1alpha1maintenance "github.com/crossplane-contrib/provider-jet-mongodbatlas/apis/maintenance/v1alpha1"
	v1alpha1mongodbatlas "github.com/crossplane-contrib/provider-jet-mongodbatlas/apis/mongodbatlas/v1alpha1"
	v1alpha2mongodbatlas "github.com/crossplane-contrib/provider-jet-mongodbatlas/apis/mongodbatlas/v1alpha2"
	v1alpha1network "github.com/crossplane-contrib/provider-jet-mongodbatlas/apis/network/v1alpha1"
	v1alpha1online "github.com/crossplane-contrib/provider-jet-mongodbatlas/apis/online/v1alpha1"
	v1alpha1org "github.com/crossplane-contrib/provider-jet-mongodbatlas/apis/org/v1alpha1"
	v1alpha1private "github.com/crossplane-contrib/provider-jet-mongodbatlas/apis/private/v1alpha1"
	v1alpha1privatelink "github.com/crossplane-contrib/provider-jet-mongodbatlas/apis/privatelink/v1alpha1"
	v1alpha1project "github.com/crossplane-contrib/provider-jet-mongodbatlas/apis/project/v1alpha1"
	v1alpha1search "github.com/crossplane-contrib/provider-jet-mongodbatlas/apis/search/v1alpha1"
	v1alpha1serverless "github.com/crossplane-contrib/provider-jet-mongodbatlas/apis/serverless/v1alpha1"
	v1alpha1third "github.com/crossplane-contrib/provider-jet-mongodbatlas/apis/third/v1alpha1"
	v1alpha1apis "github.com/crossplane-contrib/provider-jet-mongodbatlas/apis/v1alpha1"
	v1beta1 "github.com/crossplane-contrib/provider-jet-mongodbatlas/apis/v1beta1"
	v1alpha1x509 "github.com/crossplane-contrib/provider-jet-mongodbatlas/apis/x509/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1alpha1alert.SchemeBuilder.AddToScheme,
		v1alpha1api.SchemeBuilder.AddToScheme,
		v1alpha1backup.SchemeBuilder.AddToScheme,
		v1alpha1cloud.SchemeBuilder.AddToScheme,
		v1alpha1custom.SchemeBuilder.AddToScheme,
		v1alpha1data.SchemeBuilder.AddToScheme,
		v1alpha2.SchemeBuilder.AddToScheme,
		v1alpha1event.SchemeBuilder.AddToScheme,
		v1alpha1federated.SchemeBuilder.AddToScheme,
		v1alpha1global.SchemeBuilder.AddToScheme,
		v1alpha1ldap.SchemeBuilder.AddToScheme,
		v1alpha1maintenance.SchemeBuilder.AddToScheme,
		v1alpha1mongodbatlas.SchemeBuilder.AddToScheme,
		v1alpha2mongodbatlas.SchemeBuilder.AddToScheme,
		v1alpha1network.SchemeBuilder.AddToScheme,
		v1alpha1online.SchemeBuilder.AddToScheme,
		v1alpha1org.SchemeBuilder.AddToScheme,
		v1alpha1private.SchemeBuilder.AddToScheme,
		v1alpha1privatelink.SchemeBuilder.AddToScheme,
		v1alpha1project.SchemeBuilder.AddToScheme,
		v1alpha1search.SchemeBuilder.AddToScheme,
		v1alpha1serverless.SchemeBuilder.AddToScheme,
		v1alpha1third.SchemeBuilder.AddToScheme,
		v1alpha1apis.SchemeBuilder.AddToScheme,
		v1beta1.SchemeBuilder.AddToScheme,
		v1alpha1x509.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
