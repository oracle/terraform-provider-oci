// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Fleet A fleet is a collection or grouping of resources based on criteria.
type Fleet struct {

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

	// Type of the Fleet.
	// PRODUCT - A fleet of product-specific resources for a product type.
	// ENVIRONMENT - A fleet of environment-specific resources for a product stack.
	// GROUP - A fleet of a fleet of either environment or product fleets.
	// GENERIC - A fleet of resources selected dynamically or manually for reporting purposes
	FleetType FleetFleetTypeEnum `mandatory:"true" json:"fleetType"`

	// The lifecycle state of the Fleet.
	LifecycleState FleetLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// Associated region
	ResourceRegion *string `mandatory:"false" json:"resourceRegion"`

	// A user-friendly description. To provide some insight about the resource.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// The time this resource was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Products associated with the Fleet.
	Products []string `mandatory:"false" json:"products"`

	// Product stack associated with the Fleet.
	// Applicable for ENVIRONMENT fleet types.
	ApplicationType *string `mandatory:"false" json:"applicationType"`

	// Environment Type associated with the Fleet.
	// Applicable for ENVIRONMENT fleet types.
	EnvironmentType *string `mandatory:"false" json:"environmentType"`

	// Group Type associated with Group Fleet.
	// Applicable for GROUP fleet types.
	GroupType FleetGroupTypeEnum `mandatory:"false" json:"groupType,omitempty"`

	// Type of resource selection in a Fleet.
	// Select resources manually or select resources based on rules.
	ResourceSelectionType FleetResourceSelectionTypeEnum `mandatory:"false" json:"resourceSelectionType,omitempty"`

	RuleSelectionCriteria *SelectionCriteria `mandatory:"false" json:"ruleSelectionCriteria"`

	NotificationPreferences *NotificationPreferences `mandatory:"false" json:"notificationPreferences"`

	// Resources associated with the Fleet if resourceSelectionType is MANUAL.
	Resources []AssociatedFleetResourceDetails `mandatory:"false" json:"resources"`

	// Properties associated with the Fleet.
	Properties []AssociatedFleetPropertyDetails `mandatory:"false" json:"properties"`

	// Credentials associated with the Fleet.
	Credentials []AssociatedFleetCredentialDetails `mandatory:"false" json:"credentials"`

	// A value that represents if auto-confirming of the targets can be enabled.
	// This will allow targets to be auto-confirmed in the fleet without manual intervention.
	IsTargetAutoConfirm *bool `mandatory:"false" json:"isTargetAutoConfirm"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m Fleet) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Fleet) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingFleetFleetTypeEnum(string(m.FleetType)); !ok && m.FleetType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FleetType: %s. Supported values are: %s.", m.FleetType, strings.Join(GetFleetFleetTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingFleetLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetFleetLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingFleetGroupTypeEnum(string(m.GroupType)); !ok && m.GroupType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GroupType: %s. Supported values are: %s.", m.GroupType, strings.Join(GetFleetGroupTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingFleetResourceSelectionTypeEnum(string(m.ResourceSelectionType)); !ok && m.ResourceSelectionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceSelectionType: %s. Supported values are: %s.", m.ResourceSelectionType, strings.Join(GetFleetResourceSelectionTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FleetFleetTypeEnum Enum with underlying type: string
type FleetFleetTypeEnum string

// Set of constants representing the allowable values for FleetFleetTypeEnum
const (
	FleetFleetTypeProduct     FleetFleetTypeEnum = "PRODUCT"
	FleetFleetTypeEnvironment FleetFleetTypeEnum = "ENVIRONMENT"
	FleetFleetTypeGeneric     FleetFleetTypeEnum = "GENERIC"
	FleetFleetTypeGroup       FleetFleetTypeEnum = "GROUP"
)

var mappingFleetFleetTypeEnum = map[string]FleetFleetTypeEnum{
	"PRODUCT":     FleetFleetTypeProduct,
	"ENVIRONMENT": FleetFleetTypeEnvironment,
	"GENERIC":     FleetFleetTypeGeneric,
	"GROUP":       FleetFleetTypeGroup,
}

var mappingFleetFleetTypeEnumLowerCase = map[string]FleetFleetTypeEnum{
	"product":     FleetFleetTypeProduct,
	"environment": FleetFleetTypeEnvironment,
	"generic":     FleetFleetTypeGeneric,
	"group":       FleetFleetTypeGroup,
}

// GetFleetFleetTypeEnumValues Enumerates the set of values for FleetFleetTypeEnum
func GetFleetFleetTypeEnumValues() []FleetFleetTypeEnum {
	values := make([]FleetFleetTypeEnum, 0)
	for _, v := range mappingFleetFleetTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetFleetFleetTypeEnumStringValues Enumerates the set of values in String for FleetFleetTypeEnum
func GetFleetFleetTypeEnumStringValues() []string {
	return []string{
		"PRODUCT",
		"ENVIRONMENT",
		"GENERIC",
		"GROUP",
	}
}

// GetMappingFleetFleetTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFleetFleetTypeEnum(val string) (FleetFleetTypeEnum, bool) {
	enum, ok := mappingFleetFleetTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// FleetGroupTypeEnum Enum with underlying type: string
type FleetGroupTypeEnum string

// Set of constants representing the allowable values for FleetGroupTypeEnum
const (
	FleetGroupTypeEnvironment FleetGroupTypeEnum = "ENVIRONMENT"
	FleetGroupTypeProduct     FleetGroupTypeEnum = "PRODUCT"
)

var mappingFleetGroupTypeEnum = map[string]FleetGroupTypeEnum{
	"ENVIRONMENT": FleetGroupTypeEnvironment,
	"PRODUCT":     FleetGroupTypeProduct,
}

var mappingFleetGroupTypeEnumLowerCase = map[string]FleetGroupTypeEnum{
	"environment": FleetGroupTypeEnvironment,
	"product":     FleetGroupTypeProduct,
}

// GetFleetGroupTypeEnumValues Enumerates the set of values for FleetGroupTypeEnum
func GetFleetGroupTypeEnumValues() []FleetGroupTypeEnum {
	values := make([]FleetGroupTypeEnum, 0)
	for _, v := range mappingFleetGroupTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetFleetGroupTypeEnumStringValues Enumerates the set of values in String for FleetGroupTypeEnum
func GetFleetGroupTypeEnumStringValues() []string {
	return []string{
		"ENVIRONMENT",
		"PRODUCT",
	}
}

// GetMappingFleetGroupTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFleetGroupTypeEnum(val string) (FleetGroupTypeEnum, bool) {
	enum, ok := mappingFleetGroupTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// FleetResourceSelectionTypeEnum Enum with underlying type: string
type FleetResourceSelectionTypeEnum string

// Set of constants representing the allowable values for FleetResourceSelectionTypeEnum
const (
	FleetResourceSelectionTypeDynamic FleetResourceSelectionTypeEnum = "DYNAMIC"
	FleetResourceSelectionTypeManual  FleetResourceSelectionTypeEnum = "MANUAL"
)

var mappingFleetResourceSelectionTypeEnum = map[string]FleetResourceSelectionTypeEnum{
	"DYNAMIC": FleetResourceSelectionTypeDynamic,
	"MANUAL":  FleetResourceSelectionTypeManual,
}

var mappingFleetResourceSelectionTypeEnumLowerCase = map[string]FleetResourceSelectionTypeEnum{
	"dynamic": FleetResourceSelectionTypeDynamic,
	"manual":  FleetResourceSelectionTypeManual,
}

// GetFleetResourceSelectionTypeEnumValues Enumerates the set of values for FleetResourceSelectionTypeEnum
func GetFleetResourceSelectionTypeEnumValues() []FleetResourceSelectionTypeEnum {
	values := make([]FleetResourceSelectionTypeEnum, 0)
	for _, v := range mappingFleetResourceSelectionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetFleetResourceSelectionTypeEnumStringValues Enumerates the set of values in String for FleetResourceSelectionTypeEnum
func GetFleetResourceSelectionTypeEnumStringValues() []string {
	return []string{
		"DYNAMIC",
		"MANUAL",
	}
}

// GetMappingFleetResourceSelectionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFleetResourceSelectionTypeEnum(val string) (FleetResourceSelectionTypeEnum, bool) {
	enum, ok := mappingFleetResourceSelectionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// FleetLifecycleStateEnum Enum with underlying type: string
type FleetLifecycleStateEnum string

// Set of constants representing the allowable values for FleetLifecycleStateEnum
const (
	FleetLifecycleStateActive         FleetLifecycleStateEnum = "ACTIVE"
	FleetLifecycleStateInactive       FleetLifecycleStateEnum = "INACTIVE"
	FleetLifecycleStateCreating       FleetLifecycleStateEnum = "CREATING"
	FleetLifecycleStateDeleted        FleetLifecycleStateEnum = "DELETED"
	FleetLifecycleStateDeleting       FleetLifecycleStateEnum = "DELETING"
	FleetLifecycleStateFailed         FleetLifecycleStateEnum = "FAILED"
	FleetLifecycleStateUpdating       FleetLifecycleStateEnum = "UPDATING"
	FleetLifecycleStateNeedsAttention FleetLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingFleetLifecycleStateEnum = map[string]FleetLifecycleStateEnum{
	"ACTIVE":          FleetLifecycleStateActive,
	"INACTIVE":        FleetLifecycleStateInactive,
	"CREATING":        FleetLifecycleStateCreating,
	"DELETED":         FleetLifecycleStateDeleted,
	"DELETING":        FleetLifecycleStateDeleting,
	"FAILED":          FleetLifecycleStateFailed,
	"UPDATING":        FleetLifecycleStateUpdating,
	"NEEDS_ATTENTION": FleetLifecycleStateNeedsAttention,
}

var mappingFleetLifecycleStateEnumLowerCase = map[string]FleetLifecycleStateEnum{
	"active":          FleetLifecycleStateActive,
	"inactive":        FleetLifecycleStateInactive,
	"creating":        FleetLifecycleStateCreating,
	"deleted":         FleetLifecycleStateDeleted,
	"deleting":        FleetLifecycleStateDeleting,
	"failed":          FleetLifecycleStateFailed,
	"updating":        FleetLifecycleStateUpdating,
	"needs_attention": FleetLifecycleStateNeedsAttention,
}

// GetFleetLifecycleStateEnumValues Enumerates the set of values for FleetLifecycleStateEnum
func GetFleetLifecycleStateEnumValues() []FleetLifecycleStateEnum {
	values := make([]FleetLifecycleStateEnum, 0)
	for _, v := range mappingFleetLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetFleetLifecycleStateEnumStringValues Enumerates the set of values in String for FleetLifecycleStateEnum
func GetFleetLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
		"CREATING",
		"DELETED",
		"DELETING",
		"FAILED",
		"UPDATING",
		"NEEDS_ATTENTION",
	}
}

// GetMappingFleetLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFleetLifecycleStateEnum(val string) (FleetLifecycleStateEnum, bool) {
	enum, ok := mappingFleetLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
