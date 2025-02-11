// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type RepositoryPubsubConfigs struct {
	/* The format of the Cloud Pub/Sub messages.
	- PROTOBUF: The message payload is a serialized protocol buffer of SourceRepoEvent.
	- JSON: The message payload is a JSON string of SourceRepoEvent. Possible values: ["PROTOBUF", "JSON"]. */
	MessageFormat string `json:"messageFormat"`

	/* Service account used for publishing Cloud Pub/Sub messages. This
	service account needs to be in the same project as the
	pubsubConfig. When added, the caller needs to have
	iam.serviceAccounts.actAs permission on this service account. If
	unspecified, it defaults to the compute engine default service
	account. */
	// +optional
	ServiceAccountRef *v1alpha1.ResourceRef `json:"serviceAccountRef,omitempty"`

	/*  */
	TopicRef v1alpha1.ResourceRef `json:"topicRef"`
}

type SourceRepoRepositorySpec struct {
	/* How this repository publishes a change in the repository through Cloud Pub/Sub.
	Keyed by the topic names. */
	// +optional
	PubsubConfigs []RepositoryPubsubConfigs `json:"pubsubConfigs,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

type SourceRepoRepositoryStatus struct {
	/* Conditions represent the latest available observations of the
	   SourceRepoRepository's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	ObservedGeneration int `json:"observedGeneration,omitempty"`
	/* The disk usage of the repo, in bytes. */
	Size int `json:"size,omitempty"`
	/* URL to clone the repository from Google Cloud Source Repositories. */
	Url string `json:"url,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SourceRepoRepository is the Schema for the sourcerepo API
// +k8s:openapi-gen=true
type SourceRepoRepository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SourceRepoRepositorySpec   `json:"spec,omitempty"`
	Status SourceRepoRepositoryStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SourceRepoRepositoryList contains a list of SourceRepoRepository
type SourceRepoRepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SourceRepoRepository `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SourceRepoRepository{}, &SourceRepoRepositoryList{})
}
