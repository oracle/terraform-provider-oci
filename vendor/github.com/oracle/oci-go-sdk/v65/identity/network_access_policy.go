// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// NetworkAccessPolicy A network access policy.
type NetworkAccessPolicy struct {

	// The OCID of the network access policy.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the tenancy containing the network access policy.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name you assign to the network access policy during creation. The name must be unique across all network access policies
	// in the tenancy and cannot be changed.
	Name *string `mandatory:"true" json:"name"`

	// An array of one or more network access policy statements written in the policy language.
	Statements []string `mandatory:"true" json:"statements"`

	// The description you assign to the network access policy. Does not have to be unique, and it's changeable.
	Description *string `mandatory:"true" json:"description"`

	// Date and time the network access policy was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Date and time the network access policy was last updated, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The network access policy's current state. After creating a network access policy, its `lifecycleState` always returns ACTIVE.
	LifecycleState NetworkAccessPolicyLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m NetworkAccessPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NetworkAccessPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingNetworkAccessPolicyLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetNetworkAccessPolicyLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// NetworkAccessPolicyLifecycleStateEnum Enum with underlying type: string
type NetworkAccessPolicyLifecycleStateEnum string

// Set of constants representing the allowable values for NetworkAccessPolicyLifecycleStateEnum
const (
	NetworkAccessPolicyLifecycleStateActive NetworkAccessPolicyLifecycleStateEnum = "ACTIVE"
)

var mappingNetworkAccessPolicyLifecycleStateEnum = map[string]NetworkAccessPolicyLifecycleStateEnum{
	"ACTIVE": NetworkAccessPolicyLifecycleStateActive,
}

var mappingNetworkAccessPolicyLifecycleStateEnumLowerCase = map[string]NetworkAccessPolicyLifecycleStateEnum{
	"active": NetworkAccessPolicyLifecycleStateActive,
}

// GetNetworkAccessPolicyLifecycleStateEnumValues Enumerates the set of values for NetworkAccessPolicyLifecycleStateEnum
func GetNetworkAccessPolicyLifecycleStateEnumValues() []NetworkAccessPolicyLifecycleStateEnum {
	values := make([]NetworkAccessPolicyLifecycleStateEnum, 0)
	for _, v := range mappingNetworkAccessPolicyLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetNetworkAccessPolicyLifecycleStateEnumStringValues Enumerates the set of values in String for NetworkAccessPolicyLifecycleStateEnum
func GetNetworkAccessPolicyLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
	}
}

// GetMappingNetworkAccessPolicyLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNetworkAccessPolicyLifecycleStateEnum(val string) (NetworkAccessPolicyLifecycleStateEnum, bool) {
	enum, ok := mappingNetworkAccessPolicyLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
