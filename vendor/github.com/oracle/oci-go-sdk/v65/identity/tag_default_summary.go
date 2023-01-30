// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// TagDefaultSummary Summary information for the specified tag default.
type TagDefaultSummary struct {

	// The OCID of the tag default.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment. The tag default will apply to all new resources that are created in the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the tag namespace that contains the tag definition.
	TagNamespaceId *string `mandatory:"true" json:"tagNamespaceId"`

	// The OCID of the tag definition. The tag default will always assign a default value for this tag definition.
	TagDefinitionId *string `mandatory:"true" json:"tagDefinitionId"`

	// The name used in the tag definition. This field is informational in the context of the tag default.
	TagDefinitionName *string `mandatory:"true" json:"tagDefinitionName"`

	// The default value for the tag definition. This will be applied to all new resources created in the compartment.
	Value *string `mandatory:"true" json:"value"`

	// Date and time the `TagDefault` object was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// If you specify that a value is required, a value is set during resource creation (either by
	// the user creating the resource or another tag defualt). If no value is set, resource
	// creation is blocked.
	// * If the `isRequired` flag is set to "true", the value is set during resource creation.
	// * If the `isRequired` flag is set to "false", the value you enter is set during resource creation.
	// Example: `false`
	IsRequired *bool `mandatory:"true" json:"isRequired"`

	// The tag default's current state. After creating a `TagDefault`, make sure its `lifecycleState` is ACTIVE before using it.
	LifecycleState TagDefaultSummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`
}

func (m TagDefaultSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TagDefaultSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingTagDefaultSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetTagDefaultSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TagDefaultSummaryLifecycleStateEnum Enum with underlying type: string
type TagDefaultSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for TagDefaultSummaryLifecycleStateEnum
const (
	TagDefaultSummaryLifecycleStateActive TagDefaultSummaryLifecycleStateEnum = "ACTIVE"
)

var mappingTagDefaultSummaryLifecycleStateEnum = map[string]TagDefaultSummaryLifecycleStateEnum{
	"ACTIVE": TagDefaultSummaryLifecycleStateActive,
}

var mappingTagDefaultSummaryLifecycleStateEnumLowerCase = map[string]TagDefaultSummaryLifecycleStateEnum{
	"active": TagDefaultSummaryLifecycleStateActive,
}

// GetTagDefaultSummaryLifecycleStateEnumValues Enumerates the set of values for TagDefaultSummaryLifecycleStateEnum
func GetTagDefaultSummaryLifecycleStateEnumValues() []TagDefaultSummaryLifecycleStateEnum {
	values := make([]TagDefaultSummaryLifecycleStateEnum, 0)
	for _, v := range mappingTagDefaultSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetTagDefaultSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for TagDefaultSummaryLifecycleStateEnum
func GetTagDefaultSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
	}
}

// GetMappingTagDefaultSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTagDefaultSummaryLifecycleStateEnum(val string) (TagDefaultSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingTagDefaultSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
