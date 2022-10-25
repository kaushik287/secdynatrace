/*
Copyright 2020 The Crossplane Authors.

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
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// MyTypeParameters are the configurable fields of a MyType.
type ManagementzoneParameters struct {
	ConfigurableField string `json:"configurableField"`
}

// MyTypeObservation are the observable fields of a MyType.
type ManagementzoneObservation struct {
	ObservableField string `json:"observableField,omitempty"`
}

// A MyTypeSpec defines the desired state of a MyType.
//crossplane runtime package
//forProvider is sent to provider
type ManagementzoneSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ManagementzoneParameters `json:"forProvider"`
	//forprovider is actual struct sent to dynatrace provider
}

// A MyTypeStatus represents the observed state of a MyType.
type ManagementzoneStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ManagementzoneObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A MyType is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type Managementzone struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ManagementzoneSpec `json:"spec"`
	Status Managementzone     `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MyTypeList contains a list of MyType
type ManagementzoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Managementzone `json:"items"`
}

// MyType type metadata.
var (
	ManagementzoneKind             = reflect.TypeOf(Managementzone{}).Name()
	ManagementzoneGroupKind        = schema.GroupKind{Group: Group, Kind: ManagementzoneKind}.String()
	ManagementzoneAPIVersion       = ManagementzoneKind + "." + SchemeGroupVersion.String()
	ManagementzoneGroupVersionKind = SchemeGroupVersion.WithKind(ManagementzoneKind)
)

func init() {
	SchemeBuilder.Register(&Managementzone{}, &Managementzone{})
}
