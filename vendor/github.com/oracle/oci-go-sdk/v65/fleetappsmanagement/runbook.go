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

// Runbook Runbook definition.
type Runbook struct {

	// The OCID of the resource.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My new resource`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The type of the runbook.
	Type RunbookTypeEnum `mandatory:"true" json:"type"`

	// Type of runbook structure.
	RunbookRelevance RunbookRunbookRelevanceEnum `mandatory:"true" json:"runbookRelevance"`

	// The lifecycle operation performed by the task.
	Operation *string `mandatory:"true" json:"operation"`

	// The OS type for the runbook.
	OsType OsTypeEnum `mandatory:"true" json:"osType"`

	// The platform of the runbook.
	Platform *string `mandatory:"true" json:"platform"`

	// Is the runbook default?
	IsDefault *bool `mandatory:"true" json:"isDefault"`

	// The current state of the Runbook.
	LifecycleState RunbookLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time this resource was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time this resource was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// OCID of the compartment to which the resource belongs to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly description. To provide some insight about the resource.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// Estimated time to successfully complete the runbook execution
	EstimatedTime *string `mandatory:"false" json:"estimatedTime"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	Associations *Associations `mandatory:"false" json:"associations"`

	// Associated region
	ResourceRegion *string `mandatory:"false" json:"resourceRegion"`

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

func (m Runbook) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Runbook) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRunbookTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetRunbookTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRunbookRunbookRelevanceEnum(string(m.RunbookRelevance)); !ok && m.RunbookRelevance != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RunbookRelevance: %s. Supported values are: %s.", m.RunbookRelevance, strings.Join(GetRunbookRunbookRelevanceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOsTypeEnum(string(m.OsType)); !ok && m.OsType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OsType: %s. Supported values are: %s.", m.OsType, strings.Join(GetOsTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRunbookLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetRunbookLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RunbookTypeEnum Enum with underlying type: string
type RunbookTypeEnum string

// Set of constants representing the allowable values for RunbookTypeEnum
const (
	RunbookTypeUserDefined   RunbookTypeEnum = "USER_DEFINED"
	RunbookTypeOracleDefined RunbookTypeEnum = "ORACLE_DEFINED"
	RunbookTypeSystemDefined RunbookTypeEnum = "SYSTEM_DEFINED"
)

var mappingRunbookTypeEnum = map[string]RunbookTypeEnum{
	"USER_DEFINED":   RunbookTypeUserDefined,
	"ORACLE_DEFINED": RunbookTypeOracleDefined,
	"SYSTEM_DEFINED": RunbookTypeSystemDefined,
}

var mappingRunbookTypeEnumLowerCase = map[string]RunbookTypeEnum{
	"user_defined":   RunbookTypeUserDefined,
	"oracle_defined": RunbookTypeOracleDefined,
	"system_defined": RunbookTypeSystemDefined,
}

// GetRunbookTypeEnumValues Enumerates the set of values for RunbookTypeEnum
func GetRunbookTypeEnumValues() []RunbookTypeEnum {
	values := make([]RunbookTypeEnum, 0)
	for _, v := range mappingRunbookTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRunbookTypeEnumStringValues Enumerates the set of values in String for RunbookTypeEnum
func GetRunbookTypeEnumStringValues() []string {
	return []string{
		"USER_DEFINED",
		"ORACLE_DEFINED",
		"SYSTEM_DEFINED",
	}
}

// GetMappingRunbookTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRunbookTypeEnum(val string) (RunbookTypeEnum, bool) {
	enum, ok := mappingRunbookTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// RunbookRunbookRelevanceEnum Enum with underlying type: string
type RunbookRunbookRelevanceEnum string

// Set of constants representing the allowable values for RunbookRunbookRelevanceEnum
const (
	RunbookRunbookRelevanceProductGroup RunbookRunbookRelevanceEnum = "PRODUCT_GROUP"
	RunbookRunbookRelevanceProduct      RunbookRunbookRelevanceEnum = "PRODUCT"
)

var mappingRunbookRunbookRelevanceEnum = map[string]RunbookRunbookRelevanceEnum{
	"PRODUCT_GROUP": RunbookRunbookRelevanceProductGroup,
	"PRODUCT":       RunbookRunbookRelevanceProduct,
}

var mappingRunbookRunbookRelevanceEnumLowerCase = map[string]RunbookRunbookRelevanceEnum{
	"product_group": RunbookRunbookRelevanceProductGroup,
	"product":       RunbookRunbookRelevanceProduct,
}

// GetRunbookRunbookRelevanceEnumValues Enumerates the set of values for RunbookRunbookRelevanceEnum
func GetRunbookRunbookRelevanceEnumValues() []RunbookRunbookRelevanceEnum {
	values := make([]RunbookRunbookRelevanceEnum, 0)
	for _, v := range mappingRunbookRunbookRelevanceEnum {
		values = append(values, v)
	}
	return values
}

// GetRunbookRunbookRelevanceEnumStringValues Enumerates the set of values in String for RunbookRunbookRelevanceEnum
func GetRunbookRunbookRelevanceEnumStringValues() []string {
	return []string{
		"PRODUCT_GROUP",
		"PRODUCT",
	}
}

// GetMappingRunbookRunbookRelevanceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRunbookRunbookRelevanceEnum(val string) (RunbookRunbookRelevanceEnum, bool) {
	enum, ok := mappingRunbookRunbookRelevanceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// RunbookLifecycleStateEnum Enum with underlying type: string
type RunbookLifecycleStateEnum string

// Set of constants representing the allowable values for RunbookLifecycleStateEnum
const (
	RunbookLifecycleStateActive  RunbookLifecycleStateEnum = "ACTIVE"
	RunbookLifecycleStateDeleted RunbookLifecycleStateEnum = "DELETED"
	RunbookLifecycleStateFailed  RunbookLifecycleStateEnum = "FAILED"
)

var mappingRunbookLifecycleStateEnum = map[string]RunbookLifecycleStateEnum{
	"ACTIVE":  RunbookLifecycleStateActive,
	"DELETED": RunbookLifecycleStateDeleted,
	"FAILED":  RunbookLifecycleStateFailed,
}

var mappingRunbookLifecycleStateEnumLowerCase = map[string]RunbookLifecycleStateEnum{
	"active":  RunbookLifecycleStateActive,
	"deleted": RunbookLifecycleStateDeleted,
	"failed":  RunbookLifecycleStateFailed,
}

// GetRunbookLifecycleStateEnumValues Enumerates the set of values for RunbookLifecycleStateEnum
func GetRunbookLifecycleStateEnumValues() []RunbookLifecycleStateEnum {
	values := make([]RunbookLifecycleStateEnum, 0)
	for _, v := range mappingRunbookLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetRunbookLifecycleStateEnumStringValues Enumerates the set of values in String for RunbookLifecycleStateEnum
func GetRunbookLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
		"FAILED",
	}
}

// GetMappingRunbookLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRunbookLifecycleStateEnum(val string) (RunbookLifecycleStateEnum, bool) {
	enum, ok := mappingRunbookLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
