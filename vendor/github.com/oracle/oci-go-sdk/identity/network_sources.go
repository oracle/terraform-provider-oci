// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"github.com/oracle/oci-go-sdk/common"
)

// NetworkSources A network source specifies a list of source IP addresses that are allowed to make authorization requests.
// Use the network source in policy statements to restrict access to only requests that come from the specified IPs.
// For more information, see Managing Network Sources (https://docs.cloud.oracle.com/Content/Identity/Tasks/managingnetworksources.htm).
type NetworkSources struct {

	// The OCID of the network source.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the tenancy containing the network source. The tenancy is the root compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name you assign to the network source during creation. The name must be unique across
	// the tenancy and cannot be changed.
	Name *string `mandatory:"true" json:"name"`

	// The description you assign to the network source. Does not have to be unique, and it's changeable.
	Description *string `mandatory:"true" json:"description"`

	// Date and time the group was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The network source object's current state. After creating a network source, make sure its `lifecycleState` changes from CREATING to
	// ACTIVE before using it.
	LifecycleState NetworkSourcesLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A list of allowed public IPs and CIDR ranges.
	PublicSourceList []string `mandatory:"false" json:"publicSourceList"`

	// A list of allowed VCN OCID and IP range pairs.
	// Example:`"vcnId": "ocid1.vcn.oc1.iad.aaaaaaaaexampleuniqueID", "ipRanges": [ "129.213.39.0/24" ]`
	VirtualSourceList []NetworkSourcesVirtualSourceList `mandatory:"false" json:"virtualSourceList"`

	// A list of services allowed to make on-behalf-of requests. These requests can have different source IPs than
	// those specified in the network source.
	// Currently, only `all` and `none` are supported. The default is `all`.
	Services []string `mandatory:"false" json:"services"`

	// The detailed status of INACTIVE lifecycleState.
	InactiveStatus *int64 `mandatory:"false" json:"inactiveStatus"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m NetworkSources) String() string {
	return common.PointerString(m)
}

// NetworkSourcesLifecycleStateEnum Enum with underlying type: string
type NetworkSourcesLifecycleStateEnum string

// Set of constants representing the allowable values for NetworkSourcesLifecycleStateEnum
const (
	NetworkSourcesLifecycleStateCreating NetworkSourcesLifecycleStateEnum = "CREATING"
	NetworkSourcesLifecycleStateActive   NetworkSourcesLifecycleStateEnum = "ACTIVE"
	NetworkSourcesLifecycleStateInactive NetworkSourcesLifecycleStateEnum = "INACTIVE"
	NetworkSourcesLifecycleStateDeleting NetworkSourcesLifecycleStateEnum = "DELETING"
	NetworkSourcesLifecycleStateDeleted  NetworkSourcesLifecycleStateEnum = "DELETED"
)

var mappingNetworkSourcesLifecycleState = map[string]NetworkSourcesLifecycleStateEnum{
	"CREATING": NetworkSourcesLifecycleStateCreating,
	"ACTIVE":   NetworkSourcesLifecycleStateActive,
	"INACTIVE": NetworkSourcesLifecycleStateInactive,
	"DELETING": NetworkSourcesLifecycleStateDeleting,
	"DELETED":  NetworkSourcesLifecycleStateDeleted,
}

// GetNetworkSourcesLifecycleStateEnumValues Enumerates the set of values for NetworkSourcesLifecycleStateEnum
func GetNetworkSourcesLifecycleStateEnumValues() []NetworkSourcesLifecycleStateEnum {
	values := make([]NetworkSourcesLifecycleStateEnum, 0)
	for _, v := range mappingNetworkSourcesLifecycleState {
		values = append(values, v)
	}
	return values
}
