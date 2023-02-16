// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// Use the Identity and Access Management Service API to manage users, groups, identity domains, compartments, policies, tagging, and limits. For information about managing users, groups, compartments, and policies, see Identity and Access Management (without identity domains) (https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm). For information about tagging and service limits, see Tagging (https://docs.cloud.oracle.com/iaas/Content/Tagging/Concepts/taggingoverview.htm) and Service Limits (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/servicelimits.htm). For information about creating, modifying, and deleting identity domains, see Identity and Access Management (with identity domains) (https://docs.cloud.oracle.com/iaas/Content/Identity/home.htm).
//

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NetworkLocationSummary Network location.
type NetworkLocationSummary struct {

	// The OCID of the network location.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the tenancy (root compartment) containing the network location.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name you assign to the network location during creation. The name must be unique across
	// the tenancy and cannot be changed.
	Name *string `mandatory:"true" json:"name"`

	// The description you assign to the network location. Does not have to be unique, and it's changeable.
	Description *string `mandatory:"true" json:"description"`

	// Date and time the network location was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Date and time the network location was last updated, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The network location object's current state. After creating a network location, its `lifecycleState` always returns ACTIVE.
	LifecycleState NetworkLocationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A list of virtual cloud network OCIDs. Includes all service gateways and private endpoints attached to the virtual cloud networks.
	VcnIds []string `mandatory:"false" json:"vcnIds"`

	// A list of service gateway OCIDs.
	ServiceGatewayIds []string `mandatory:"false" json:"serviceGatewayIds"`

	// A list of private endpoint OCIDs.
	PrivateEndpointIds []string `mandatory:"false" json:"privateEndpointIds"`

	// A list of IP addresses, CIDR blocks, or IPv6 prefixes.
	// Example: `[ "192.0.2.0/24", "::1/48" ]`
	IpAddresses []string `mandatory:"false" json:"ipAddresses"`

	// A list of header names and values.
	// Example: `[{"headerName": "x-mynetwork-header", "headerValue": "acme"}]`
	CorporateHeaders []CorporateHeader `mandatory:"false" json:"corporateHeaders"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m NetworkLocationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NetworkLocationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingNetworkLocationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetNetworkLocationLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
