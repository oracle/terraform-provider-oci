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

// DynamicGroup A dynamic group defines a matching rule. Every bare metal or virtual machine instance is deployed with an instance certificate.
// The certificate contains metadata about the instance. This includes the instance OCID and the compartment OCID, along
// with a few other optional properties. When an API call is made using this instance certificate as the authenticator,
// the certificate can be matched to one or multiple dynamic groups. The instance can then get access to the API
// based on the permissions granted in policies written for the dynamic groups.
// This works like regular user/group membership. But in that case, the membership is a static relationship, whereas
// in a dynamic group, the membership of an instance certificate to a dynamic group is determined during runtime.
// For more information, see Managing Dynamic Groups (https://docs.oracle.com/iaas/Content/Identity/dynamicgroups/managingdynamicgroups.htm).
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using
// the API.
type DynamicGroup struct {

	// The OCID of the group.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the tenancy containing the group.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name you assign to the group during creation. The name must be unique across all groups in
	// the tenancy and cannot be changed.
	Name *string `mandatory:"true" json:"name"`

	// The description you assign to the group. Does not have to be unique, and it's changeable.
	// (For tenancies that support identity domains) You can have an empty description.
	Description *string `mandatory:"true" json:"description"`

	// A rule string that defines which instance certificates will be matched.
	// For syntax, see Managing Dynamic Groups (https://docs.oracle.com/iaas/Content/Identity/dynamicgroups/managingdynamicgroups.htm).
	MatchingRule *string `mandatory:"true" json:"matchingRule"`

	// Date and time the group was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The group's current state. After creating a group, make sure its `lifecycleState` changes from CREATING to
	// ACTIVE before using it.
	LifecycleState DynamicGroupLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

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

func (m DynamicGroup) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DynamicGroup) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDynamicGroupLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDynamicGroupLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DynamicGroupLifecycleStateEnum Enum with underlying type: string
type DynamicGroupLifecycleStateEnum string

// Set of constants representing the allowable values for DynamicGroupLifecycleStateEnum
const (
	DynamicGroupLifecycleStateCreating DynamicGroupLifecycleStateEnum = "CREATING"
	DynamicGroupLifecycleStateActive   DynamicGroupLifecycleStateEnum = "ACTIVE"
	DynamicGroupLifecycleStateInactive DynamicGroupLifecycleStateEnum = "INACTIVE"
	DynamicGroupLifecycleStateDeleting DynamicGroupLifecycleStateEnum = "DELETING"
	DynamicGroupLifecycleStateDeleted  DynamicGroupLifecycleStateEnum = "DELETED"
)

var mappingDynamicGroupLifecycleStateEnum = map[string]DynamicGroupLifecycleStateEnum{
	"CREATING": DynamicGroupLifecycleStateCreating,
	"ACTIVE":   DynamicGroupLifecycleStateActive,
	"INACTIVE": DynamicGroupLifecycleStateInactive,
	"DELETING": DynamicGroupLifecycleStateDeleting,
	"DELETED":  DynamicGroupLifecycleStateDeleted,
}

var mappingDynamicGroupLifecycleStateEnumLowerCase = map[string]DynamicGroupLifecycleStateEnum{
	"creating": DynamicGroupLifecycleStateCreating,
	"active":   DynamicGroupLifecycleStateActive,
	"inactive": DynamicGroupLifecycleStateInactive,
	"deleting": DynamicGroupLifecycleStateDeleting,
	"deleted":  DynamicGroupLifecycleStateDeleted,
}

// GetDynamicGroupLifecycleStateEnumValues Enumerates the set of values for DynamicGroupLifecycleStateEnum
func GetDynamicGroupLifecycleStateEnumValues() []DynamicGroupLifecycleStateEnum {
	values := make([]DynamicGroupLifecycleStateEnum, 0)
	for _, v := range mappingDynamicGroupLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDynamicGroupLifecycleStateEnumStringValues Enumerates the set of values in String for DynamicGroupLifecycleStateEnum
func GetDynamicGroupLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
	}
}

// GetMappingDynamicGroupLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDynamicGroupLifecycleStateEnum(val string) (DynamicGroupLifecycleStateEnum, bool) {
	enum, ok := mappingDynamicGroupLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
