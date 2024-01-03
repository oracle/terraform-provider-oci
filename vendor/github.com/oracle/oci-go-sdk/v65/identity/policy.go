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

// Policy A document that specifies the type of access a group has to the resources in a compartment. For information about
// policies and other IAM Service components, see
// Overview of IAM (https://docs.cloud.oracle.com/Content/Identity/getstarted/identity-domains.htm). If you're new to policies, see
// Get Started with Policies (https://docs.cloud.oracle.com/Content/Identity/policiesgs/get-started-with-policies.htm).
// The word "policy" is used by people in different ways:
//   - An individual statement written in the policy language
//   - A collection of statements in a single, named "policy" document (which has an Oracle Cloud ID (OCID) assigned to it)
//   - The overall body of policies your organization uses to control access to resources
//
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values
// using the API.
type Policy struct {

	// The OCID of the policy.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment containing the policy (either the tenancy or another compartment).
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name you assign to the policy during creation. The name must be unique across all policies
	// in the tenancy and cannot be changed.
	Name *string `mandatory:"true" json:"name"`

	// An array of one or more policy statements written in the policy language.
	Statements []string `mandatory:"true" json:"statements"`

	// The description you assign to the policy. Does not have to be unique, and it's changeable.
	Description *string `mandatory:"true" json:"description"`

	// Date and time the policy was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The policy's current state. After creating a policy, make sure its `lifecycleState` changes from CREATING to
	// ACTIVE before using it.
	LifecycleState PolicyLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The detailed status of INACTIVE lifecycleState.
	InactiveStatus *int64 `mandatory:"false" json:"inactiveStatus"`

	// The version of the policy. If null or set to an empty string, when a request comes in for authorization, the
	// policy will be evaluated according to the current behavior of the services at that moment. If set to a particular
	// date (YYYY-MM-DD), the policy will be evaluated according to the behavior of the services on that date.
	VersionDate *common.SDKDate `mandatory:"false" json:"versionDate"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m Policy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Policy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPolicyLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPolicyLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PolicyLifecycleStateEnum Enum with underlying type: string
type PolicyLifecycleStateEnum string

// Set of constants representing the allowable values for PolicyLifecycleStateEnum
const (
	PolicyLifecycleStateCreating PolicyLifecycleStateEnum = "CREATING"
	PolicyLifecycleStateActive   PolicyLifecycleStateEnum = "ACTIVE"
	PolicyLifecycleStateInactive PolicyLifecycleStateEnum = "INACTIVE"
	PolicyLifecycleStateDeleting PolicyLifecycleStateEnum = "DELETING"
	PolicyLifecycleStateDeleted  PolicyLifecycleStateEnum = "DELETED"
)

var mappingPolicyLifecycleStateEnum = map[string]PolicyLifecycleStateEnum{
	"CREATING": PolicyLifecycleStateCreating,
	"ACTIVE":   PolicyLifecycleStateActive,
	"INACTIVE": PolicyLifecycleStateInactive,
	"DELETING": PolicyLifecycleStateDeleting,
	"DELETED":  PolicyLifecycleStateDeleted,
}

var mappingPolicyLifecycleStateEnumLowerCase = map[string]PolicyLifecycleStateEnum{
	"creating": PolicyLifecycleStateCreating,
	"active":   PolicyLifecycleStateActive,
	"inactive": PolicyLifecycleStateInactive,
	"deleting": PolicyLifecycleStateDeleting,
	"deleted":  PolicyLifecycleStateDeleted,
}

// GetPolicyLifecycleStateEnumValues Enumerates the set of values for PolicyLifecycleStateEnum
func GetPolicyLifecycleStateEnumValues() []PolicyLifecycleStateEnum {
	values := make([]PolicyLifecycleStateEnum, 0)
	for _, v := range mappingPolicyLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPolicyLifecycleStateEnumStringValues Enumerates the set of values in String for PolicyLifecycleStateEnum
func GetPolicyLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
	}
}

// GetMappingPolicyLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPolicyLifecycleStateEnum(val string) (PolicyLifecycleStateEnum, bool) {
	enum, ok := mappingPolicyLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
