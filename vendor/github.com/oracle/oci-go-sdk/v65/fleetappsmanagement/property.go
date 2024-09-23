// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management Service API. Use this API to for all FAMS related activities.
// To manage fleets,view complaince report for the Fleet,scedule patches and other lifecycle activities
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Property Taxonomy metadata aka Property .
type Property struct {

	// The OCID of the resource.
	Id *string `mandatory:"true" json:"id"`

	// Tenancy OCID
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My new resource`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The time this resource was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Associated region
	ResourceRegion *string `mandatory:"true" json:"resourceRegion"`

	// The current state of the Property.
	LifecycleState PropertyLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time this resource was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Text selection of the category
	Selection SelectionEnum `mandatory:"false" json:"selection,omitempty"`

	// Format of the value
	ValueType ValueTypeEnum `mandatory:"false" json:"valueType,omitempty"`

	// Values of the property (must be a single value if selection = 'single choice')
	Values []string `mandatory:"false" json:"values"`

	// The scope of the property
	Scope ScopeEnum `mandatory:"false" json:"scope,omitempty"`

	// The type of the property.
	Type PropertyTypeEnum `mandatory:"false" json:"type,omitempty"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m Property) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Property) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPropertyLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPropertyLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingSelectionEnum(string(m.Selection)); !ok && m.Selection != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Selection: %s. Supported values are: %s.", m.Selection, strings.Join(GetSelectionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingValueTypeEnum(string(m.ValueType)); !ok && m.ValueType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ValueType: %s. Supported values are: %s.", m.ValueType, strings.Join(GetValueTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingScopeEnum(string(m.Scope)); !ok && m.Scope != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Scope: %s. Supported values are: %s.", m.Scope, strings.Join(GetScopeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPropertyTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetPropertyTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PropertyLifecycleStateEnum Enum with underlying type: string
type PropertyLifecycleStateEnum string

// Set of constants representing the allowable values for PropertyLifecycleStateEnum
const (
	PropertyLifecycleStateActive  PropertyLifecycleStateEnum = "ACTIVE"
	PropertyLifecycleStateDeleted PropertyLifecycleStateEnum = "DELETED"
	PropertyLifecycleStateFailed  PropertyLifecycleStateEnum = "FAILED"
)

var mappingPropertyLifecycleStateEnum = map[string]PropertyLifecycleStateEnum{
	"ACTIVE":  PropertyLifecycleStateActive,
	"DELETED": PropertyLifecycleStateDeleted,
	"FAILED":  PropertyLifecycleStateFailed,
}

var mappingPropertyLifecycleStateEnumLowerCase = map[string]PropertyLifecycleStateEnum{
	"active":  PropertyLifecycleStateActive,
	"deleted": PropertyLifecycleStateDeleted,
	"failed":  PropertyLifecycleStateFailed,
}

// GetPropertyLifecycleStateEnumValues Enumerates the set of values for PropertyLifecycleStateEnum
func GetPropertyLifecycleStateEnumValues() []PropertyLifecycleStateEnum {
	values := make([]PropertyLifecycleStateEnum, 0)
	for _, v := range mappingPropertyLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPropertyLifecycleStateEnumStringValues Enumerates the set of values in String for PropertyLifecycleStateEnum
func GetPropertyLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
		"FAILED",
	}
}

// GetMappingPropertyLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPropertyLifecycleStateEnum(val string) (PropertyLifecycleStateEnum, bool) {
	enum, ok := mappingPropertyLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// PropertyTypeEnum Enum with underlying type: string
type PropertyTypeEnum string

// Set of constants representing the allowable values for PropertyTypeEnum
const (
	PropertyTypeUserDefined   PropertyTypeEnum = "USER_DEFINED"
	PropertyTypeOracleDefined PropertyTypeEnum = "ORACLE_DEFINED"
	PropertyTypeSystemDefined PropertyTypeEnum = "SYSTEM_DEFINED"
)

var mappingPropertyTypeEnum = map[string]PropertyTypeEnum{
	"USER_DEFINED":   PropertyTypeUserDefined,
	"ORACLE_DEFINED": PropertyTypeOracleDefined,
	"SYSTEM_DEFINED": PropertyTypeSystemDefined,
}

var mappingPropertyTypeEnumLowerCase = map[string]PropertyTypeEnum{
	"user_defined":   PropertyTypeUserDefined,
	"oracle_defined": PropertyTypeOracleDefined,
	"system_defined": PropertyTypeSystemDefined,
}

// GetPropertyTypeEnumValues Enumerates the set of values for PropertyTypeEnum
func GetPropertyTypeEnumValues() []PropertyTypeEnum {
	values := make([]PropertyTypeEnum, 0)
	for _, v := range mappingPropertyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPropertyTypeEnumStringValues Enumerates the set of values in String for PropertyTypeEnum
func GetPropertyTypeEnumStringValues() []string {
	return []string{
		"USER_DEFINED",
		"ORACLE_DEFINED",
		"SYSTEM_DEFINED",
	}
}

// GetMappingPropertyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPropertyTypeEnum(val string) (PropertyTypeEnum, bool) {
	enum, ok := mappingPropertyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
