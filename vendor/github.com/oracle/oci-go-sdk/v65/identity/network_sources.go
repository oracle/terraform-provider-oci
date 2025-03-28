// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// Use the Identity and Access Management Service API to manage users, groups, identity domains, compartments, policies, tagging, and limits. For information about managing users, groups, compartments, and policies, see Identity and Access Management (without identity domains) (https://docs.oracle.com/iaas/Content/Identity/Concepts/overview.htm). For information about tagging and service limits, see Tagging (https://docs.oracle.com/iaas/Content/Tagging/Concepts/taggingoverview.htm) and Service Limits (https://docs.oracle.com/iaas/Content/General/Concepts/servicelimits.htm). For information about creating, modifying, and deleting identity domains, see Identity and Access Management (with identity domains) (https://docs.oracle.com/iaas/Content/Identity/home.htm).
//

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NetworkSources A network source specifies a list of source IP addresses that are allowed to make authorization requests.
// Use the network source in policy statements to restrict access to only requests that come from the specified IPs.
// For more information, see Managing Network Sources (https://docs.oracle.com/iaas/Content/Identity/Tasks/managingnetworksources.htm).
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

	// Date and time the network source was created, in the format defined by RFC3339.
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

	// -- The services attribute has no effect and is reserved for use by Oracle. --
	Services []string `mandatory:"false" json:"services"`

	// The detailed status of INACTIVE lifecycleState.
	InactiveStatus *int64 `mandatory:"false" json:"inactiveStatus"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m NetworkSources) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NetworkSources) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingNetworkSourcesLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetNetworkSourcesLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingNetworkSourcesLifecycleStateEnum = map[string]NetworkSourcesLifecycleStateEnum{
	"CREATING": NetworkSourcesLifecycleStateCreating,
	"ACTIVE":   NetworkSourcesLifecycleStateActive,
	"INACTIVE": NetworkSourcesLifecycleStateInactive,
	"DELETING": NetworkSourcesLifecycleStateDeleting,
	"DELETED":  NetworkSourcesLifecycleStateDeleted,
}

var mappingNetworkSourcesLifecycleStateEnumLowerCase = map[string]NetworkSourcesLifecycleStateEnum{
	"creating": NetworkSourcesLifecycleStateCreating,
	"active":   NetworkSourcesLifecycleStateActive,
	"inactive": NetworkSourcesLifecycleStateInactive,
	"deleting": NetworkSourcesLifecycleStateDeleting,
	"deleted":  NetworkSourcesLifecycleStateDeleted,
}

// GetNetworkSourcesLifecycleStateEnumValues Enumerates the set of values for NetworkSourcesLifecycleStateEnum
func GetNetworkSourcesLifecycleStateEnumValues() []NetworkSourcesLifecycleStateEnum {
	values := make([]NetworkSourcesLifecycleStateEnum, 0)
	for _, v := range mappingNetworkSourcesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetNetworkSourcesLifecycleStateEnumStringValues Enumerates the set of values in String for NetworkSourcesLifecycleStateEnum
func GetNetworkSourcesLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
	}
}

// GetMappingNetworkSourcesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNetworkSourcesLifecycleStateEnum(val string) (NetworkSourcesLifecycleStateEnum, bool) {
	enum, ok := mappingNetworkSourcesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
