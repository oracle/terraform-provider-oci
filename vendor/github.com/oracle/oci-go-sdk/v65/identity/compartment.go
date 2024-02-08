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

// Compartment A collection of related resources. Compartments are a fundamental component of Oracle Cloud Infrastructure
// for organizing and isolating your cloud resources. You use them to clearly separate resources for the purposes
// of measuring usage and billing, access (through the use of IAM Service policies), and isolation (separating the
// resources for one project or business unit from another). A common approach is to create a compartment for each
// major part of your organization. For more information, see
// Overview of IAM (https://docs.cloud.oracle.com//Content/Identity/getstarted/identity-domains.htm) and also
// Setting Up Your Tenancy (https://docs.cloud.oracle.com/Content/GSG/Concepts/settinguptenancy.htm).
// To place a resource in a compartment, simply specify the compartment ID in the "Create" request object when
// initially creating the resource. For example, to launch an instance into a particular compartment, specify
// that compartment's OCID in the `LaunchInstance` request. You can't move an existing resource from one
// compartment to another.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access,
// see Get Started with Policies (https://docs.cloud.oracle.com/Content/Identity/policiesgs/get-started-with-policies.htm).
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values
// using the API.
type Compartment struct {

	// The OCID of the compartment.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the parent compartment containing the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name you assign to the compartment during creation. The name must be unique across all
	// compartments in the parent. Avoid entering confidential information.
	Name *string `mandatory:"true" json:"name"`

	// The description you assign to the compartment. Does not have to be unique, and it's changeable.
	Description *string `mandatory:"true" json:"description"`

	// Date and time the compartment was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The compartment's current state. After creating a compartment, make sure its `lifecycleState` changes from
	// CREATING to ACTIVE before using it.
	LifecycleState CompartmentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The detailed status of INACTIVE lifecycleState.
	InactiveStatus *int64 `mandatory:"false" json:"inactiveStatus"`

	// Indicates whether or not the compartment is accessible for the user making the request.
	// Returns true when the user has INSPECT permissions directly on a resource in the
	// compartment or indirectly (permissions can be on a resource in a subcompartment).
	IsAccessible *bool `mandatory:"false" json:"isAccessible"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m Compartment) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Compartment) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCompartmentLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCompartmentLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CompartmentLifecycleStateEnum Enum with underlying type: string
type CompartmentLifecycleStateEnum string

// Set of constants representing the allowable values for CompartmentLifecycleStateEnum
const (
	CompartmentLifecycleStateCreating CompartmentLifecycleStateEnum = "CREATING"
	CompartmentLifecycleStateActive   CompartmentLifecycleStateEnum = "ACTIVE"
	CompartmentLifecycleStateInactive CompartmentLifecycleStateEnum = "INACTIVE"
	CompartmentLifecycleStateDeleting CompartmentLifecycleStateEnum = "DELETING"
	CompartmentLifecycleStateDeleted  CompartmentLifecycleStateEnum = "DELETED"
)

var mappingCompartmentLifecycleStateEnum = map[string]CompartmentLifecycleStateEnum{
	"CREATING": CompartmentLifecycleStateCreating,
	"ACTIVE":   CompartmentLifecycleStateActive,
	"INACTIVE": CompartmentLifecycleStateInactive,
	"DELETING": CompartmentLifecycleStateDeleting,
	"DELETED":  CompartmentLifecycleStateDeleted,
}

var mappingCompartmentLifecycleStateEnumLowerCase = map[string]CompartmentLifecycleStateEnum{
	"creating": CompartmentLifecycleStateCreating,
	"active":   CompartmentLifecycleStateActive,
	"inactive": CompartmentLifecycleStateInactive,
	"deleting": CompartmentLifecycleStateDeleting,
	"deleted":  CompartmentLifecycleStateDeleted,
}

// GetCompartmentLifecycleStateEnumValues Enumerates the set of values for CompartmentLifecycleStateEnum
func GetCompartmentLifecycleStateEnumValues() []CompartmentLifecycleStateEnum {
	values := make([]CompartmentLifecycleStateEnum, 0)
	for _, v := range mappingCompartmentLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetCompartmentLifecycleStateEnumStringValues Enumerates the set of values in String for CompartmentLifecycleStateEnum
func GetCompartmentLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
	}
}

// GetMappingCompartmentLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCompartmentLifecycleStateEnum(val string) (CompartmentLifecycleStateEnum, bool) {
	enum, ok := mappingCompartmentLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
