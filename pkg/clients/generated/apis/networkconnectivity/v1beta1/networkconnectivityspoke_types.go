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

type SpokeInstances struct {
	/* Immutable. The IP address on the VM to use for peering. */
	// +optional
	IpAddress *string `json:"ipAddress,omitempty"`

	/* Immutable. */
	// +optional
	VirtualMachineRef *v1alpha1.ResourceRef `json:"virtualMachineRef,omitempty"`
}

type SpokeLinkedInterconnectAttachments struct {
	/* Immutable. A value that controls whether site-to-site data transfer is enabled for these resources. Note that data transfer is available only in supported locations. */
	SiteToSiteDataTransfer bool `json:"siteToSiteDataTransfer"`

	/* Immutable. */
	Uris []v1alpha1.ResourceRef `json:"uris"`
}

type SpokeLinkedRouterApplianceInstances struct {
	/* Immutable. The list of router appliance instances */
	Instances []SpokeInstances `json:"instances"`

	/* Immutable. A value that controls whether site-to-site data transfer is enabled for these resources. Note that data transfer is available only in supported locations. */
	SiteToSiteDataTransfer bool `json:"siteToSiteDataTransfer"`
}

type SpokeLinkedVpnTunnels struct {
	/* Immutable. A value that controls whether site-to-site data transfer is enabled for these resources. Note that data transfer is available only in supported locations. */
	SiteToSiteDataTransfer bool `json:"siteToSiteDataTransfer"`

	/* Immutable. */
	Uris []v1alpha1.ResourceRef `json:"uris"`
}

type NetworkConnectivitySpokeSpec struct {
	/* An optional description of the spoke. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Immutable. */
	HubRef v1alpha1.ResourceRef `json:"hubRef"`

	/* Immutable. A collection of VLAN attachment resources. These resources should be redundant attachments that all advertise the same prefixes to Google Cloud. Alternatively, in active/passive configurations, all attachments should be capable of advertising the same prefixes. */
	// +optional
	LinkedInterconnectAttachments *SpokeLinkedInterconnectAttachments `json:"linkedInterconnectAttachments,omitempty"`

	/* Immutable. The URIs of linked Router appliance resources */
	// +optional
	LinkedRouterApplianceInstances *SpokeLinkedRouterApplianceInstances `json:"linkedRouterApplianceInstances,omitempty"`

	/* Immutable. The URIs of linked VPN tunnel resources */
	// +optional
	LinkedVpnTunnels *SpokeLinkedVpnTunnels `json:"linkedVpnTunnels,omitempty"`

	/* Immutable. The location for the resource */
	Location string `json:"location"`

	/* Immutable. The Project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

type NetworkConnectivitySpokeStatus struct {
	/* Conditions represent the latest available observations of the
	   NetworkConnectivitySpoke's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Output only. The time the spoke was created. */
	CreateTime string `json:"createTime,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	ObservedGeneration int `json:"observedGeneration,omitempty"`
	/* Output only. The current lifecycle state of this spoke. Possible values: STATE_UNSPECIFIED, CREATING, ACTIVE, DELETING */
	State string `json:"state,omitempty"`
	/* Output only. The Google-generated UUID for the spoke. This value is unique across all spoke resources. If a spoke is deleted and another with the same name is created, the new spoke is assigned a different unique_id. */
	UniqueId string `json:"uniqueId,omitempty"`
	/* Output only. The time the spoke was last updated. */
	UpdateTime string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetworkConnectivitySpoke is the Schema for the networkconnectivity API
// +k8s:openapi-gen=true
type NetworkConnectivitySpoke struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkConnectivitySpokeSpec   `json:"spec,omitempty"`
	Status NetworkConnectivitySpokeStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetworkConnectivitySpokeList contains a list of NetworkConnectivitySpoke
type NetworkConnectivitySpokeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkConnectivitySpoke `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkConnectivitySpoke{}, &NetworkConnectivitySpokeList{})
}
