// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, policies, and identity domains.
//

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Group A collection of users who all need the same type of access to a particular set of resources or compartment.
// For conceptual information about groups and other IAM Service components, see
// Overview of IAM (https://docs.cloud.oracle.com/Content/Identity/getstarted/identity-domains.htm).
// If you're federating with an identity provider (IdP), you need to create mappings between the groups
// defined in the IdP and groups you define in the IAM service. For more information, see
// Identity Providers and Federation (https://docs.cloud.oracle.com/Content/Identity/Concepts/federation.htm). Also see
// IdentityProvider and
// IdpGroupMapping.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access,
// see Get Started with Policies (https://docs.cloud.oracle.com/Content/Identity/policiesgs/get-started-with-policies.htm).
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values
// using the API.
type Group struct {

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

	// Date and time the group was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The group's current state. After creating a group, make sure its `lifecycleState` changes from CREATING to
	// ACTIVE before using it.
	LifecycleState GroupLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

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

func (m Group) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Group) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGroupLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetGroupLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GroupLifecycleStateEnum Enum with underlying type: string
type GroupLifecycleStateEnum string

// Set of constants representing the allowable values for GroupLifecycleStateEnum
const (
	GroupLifecycleStateCreating GroupLifecycleStateEnum = "CREATING"
	GroupLifecycleStateActive   GroupLifecycleStateEnum = "ACTIVE"
	GroupLifecycleStateInactive GroupLifecycleStateEnum = "INACTIVE"
	GroupLifecycleStateDeleting GroupLifecycleStateEnum = "DELETING"
	GroupLifecycleStateDeleted  GroupLifecycleStateEnum = "DELETED"
)

var mappingGroupLifecycleStateEnum = map[string]GroupLifecycleStateEnum{
	"CREATING": GroupLifecycleStateCreating,
	"ACTIVE":   GroupLifecycleStateActive,
	"INACTIVE": GroupLifecycleStateInactive,
	"DELETING": GroupLifecycleStateDeleting,
	"DELETED":  GroupLifecycleStateDeleted,
}

var mappingGroupLifecycleStateEnumLowerCase = map[string]GroupLifecycleStateEnum{
	"creating": GroupLifecycleStateCreating,
	"active":   GroupLifecycleStateActive,
	"inactive": GroupLifecycleStateInactive,
	"deleting": GroupLifecycleStateDeleting,
	"deleted":  GroupLifecycleStateDeleted,
}

// GetGroupLifecycleStateEnumValues Enumerates the set of values for GroupLifecycleStateEnum
func GetGroupLifecycleStateEnumValues() []GroupLifecycleStateEnum {
	values := make([]GroupLifecycleStateEnum, 0)
	for _, v := range mappingGroupLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetGroupLifecycleStateEnumStringValues Enumerates the set of values in String for GroupLifecycleStateEnum
func GetGroupLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
	}
}

// GetMappingGroupLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGroupLifecycleStateEnum(val string) (GroupLifecycleStateEnum, bool) {
	enum, ok := mappingGroupLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
