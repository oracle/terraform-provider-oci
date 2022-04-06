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

// TagRule A document that specifies one or more actions on tag key(s) and/or value(s) for one or more resource types in a
// compartment.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator.
type TagRule struct {

	// The OCID of the tag rule.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment containing the tag rule (either the tenancy or another compartment).
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name you assign to the tag rule during creation. The name must be unique across all tag rules
	// in the tenancy and cannot be changed.
	Name *string `mandatory:"true" json:"name"`

	// A statement that is written in the tag rule language.
	RuleText *string `mandatory:"true" json:"ruleText"`

	// Date and time the tag rule was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the tag rule. After creating a tag rule, make sure its `lifecycleState` changes from
	// CREATING to ACTIVE before using it.
	LifecycleState TagRuleLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Returned only if the user's `lifecycleState` is INACTIVE. A 16-bit value showing the reason why the user
	// is inactive:
	// - bit 0: SUSPENDED (reserved for future use)
	// - bit 1: DISABLED (reserved for future use)
	// - bit 2: BLOCKED (the user has exceeded the maximum number of failed login attempts for the Console)
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

func (m TagRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TagRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTagRuleLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetTagRuleLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TagRuleLifecycleStateEnum Enum with underlying type: string
type TagRuleLifecycleStateEnum string

// Set of constants representing the allowable values for TagRuleLifecycleStateEnum
const (
	TagRuleLifecycleStateCreating TagRuleLifecycleStateEnum = "CREATING"
	TagRuleLifecycleStateActive   TagRuleLifecycleStateEnum = "ACTIVE"
	TagRuleLifecycleStateInactive TagRuleLifecycleStateEnum = "INACTIVE"
	TagRuleLifecycleStateDeleting TagRuleLifecycleStateEnum = "DELETING"
	TagRuleLifecycleStateDeleted  TagRuleLifecycleStateEnum = "DELETED"
)

var mappingTagRuleLifecycleStateEnum = map[string]TagRuleLifecycleStateEnum{
	"CREATING": TagRuleLifecycleStateCreating,
	"ACTIVE":   TagRuleLifecycleStateActive,
	"INACTIVE": TagRuleLifecycleStateInactive,
	"DELETING": TagRuleLifecycleStateDeleting,
	"DELETED":  TagRuleLifecycleStateDeleted,
}

var mappingTagRuleLifecycleStateEnumLowerCase = map[string]TagRuleLifecycleStateEnum{
	"creating": TagRuleLifecycleStateCreating,
	"active":   TagRuleLifecycleStateActive,
	"inactive": TagRuleLifecycleStateInactive,
	"deleting": TagRuleLifecycleStateDeleting,
	"deleted":  TagRuleLifecycleStateDeleted,
}

// GetTagRuleLifecycleStateEnumValues Enumerates the set of values for TagRuleLifecycleStateEnum
func GetTagRuleLifecycleStateEnumValues() []TagRuleLifecycleStateEnum {
	values := make([]TagRuleLifecycleStateEnum, 0)
	for _, v := range mappingTagRuleLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetTagRuleLifecycleStateEnumStringValues Enumerates the set of values in String for TagRuleLifecycleStateEnum
func GetTagRuleLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
	}
}

// GetMappingTagRuleLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTagRuleLifecycleStateEnum(val string) (TagRuleLifecycleStateEnum, bool) {
	enum, ok := mappingTagRuleLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
