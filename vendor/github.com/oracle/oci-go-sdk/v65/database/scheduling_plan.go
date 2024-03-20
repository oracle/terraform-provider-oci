// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SchedulingPlan Details of a Scheduling Plan.
type SchedulingPlan struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Scheduling Plan.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Scheduling Policy.
	SchedulingPolicyId *string `mandatory:"true" json:"schedulingPolicyId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the resource.
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// The current state of the Scheduling Plan. Valid states are CREATING, NEEDS_ATTENTION, AVAILABLE, UPDATING, FAILED, DELETING and DELETED.
	LifecycleState SchedulingPlanLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The service type of the Scheduling Plan.
	ServiceType SchedulingPlanServiceTypeEnum `mandatory:"true" json:"serviceType"`

	// The date and time the Scheduling Plan Resource was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The display name of the Scheduling Plan.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// If true, recommended scheduled actions will be generated for the scheduling plan.
	IsUsingRecommendedScheduledActions *bool `mandatory:"false" json:"isUsingRecommendedScheduledActions"`

	// The current intent the Scheduling Plan. Valid states is EXADATA_INFRASTRUCTURE_FULL_SOFTWARE_UPDATE.
	PlanIntent SchedulingPlanPlanIntentEnum `mandatory:"false" json:"planIntent,omitempty"`

	// The estimated time for the Scheduling Plan.
	EstimatedTimeInMins *int `mandatory:"false" json:"estimatedTimeInMins"`

	// The date and time the Scheduling Plan Resource was updated.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m SchedulingPlan) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SchedulingPlan) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSchedulingPlanLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSchedulingPlanLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSchedulingPlanServiceTypeEnum(string(m.ServiceType)); !ok && m.ServiceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ServiceType: %s. Supported values are: %s.", m.ServiceType, strings.Join(GetSchedulingPlanServiceTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingSchedulingPlanPlanIntentEnum(string(m.PlanIntent)); !ok && m.PlanIntent != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlanIntent: %s. Supported values are: %s.", m.PlanIntent, strings.Join(GetSchedulingPlanPlanIntentEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SchedulingPlanLifecycleStateEnum Enum with underlying type: string
type SchedulingPlanLifecycleStateEnum string

// Set of constants representing the allowable values for SchedulingPlanLifecycleStateEnum
const (
	SchedulingPlanLifecycleStateCreating       SchedulingPlanLifecycleStateEnum = "CREATING"
	SchedulingPlanLifecycleStateNeedsAttention SchedulingPlanLifecycleStateEnum = "NEEDS_ATTENTION"
	SchedulingPlanLifecycleStateAvailable      SchedulingPlanLifecycleStateEnum = "AVAILABLE"
	SchedulingPlanLifecycleStateUpdating       SchedulingPlanLifecycleStateEnum = "UPDATING"
	SchedulingPlanLifecycleStateFailed         SchedulingPlanLifecycleStateEnum = "FAILED"
	SchedulingPlanLifecycleStateDeleting       SchedulingPlanLifecycleStateEnum = "DELETING"
	SchedulingPlanLifecycleStateDeleted        SchedulingPlanLifecycleStateEnum = "DELETED"
)

var mappingSchedulingPlanLifecycleStateEnum = map[string]SchedulingPlanLifecycleStateEnum{
	"CREATING":        SchedulingPlanLifecycleStateCreating,
	"NEEDS_ATTENTION": SchedulingPlanLifecycleStateNeedsAttention,
	"AVAILABLE":       SchedulingPlanLifecycleStateAvailable,
	"UPDATING":        SchedulingPlanLifecycleStateUpdating,
	"FAILED":          SchedulingPlanLifecycleStateFailed,
	"DELETING":        SchedulingPlanLifecycleStateDeleting,
	"DELETED":         SchedulingPlanLifecycleStateDeleted,
}

var mappingSchedulingPlanLifecycleStateEnumLowerCase = map[string]SchedulingPlanLifecycleStateEnum{
	"creating":        SchedulingPlanLifecycleStateCreating,
	"needs_attention": SchedulingPlanLifecycleStateNeedsAttention,
	"available":       SchedulingPlanLifecycleStateAvailable,
	"updating":        SchedulingPlanLifecycleStateUpdating,
	"failed":          SchedulingPlanLifecycleStateFailed,
	"deleting":        SchedulingPlanLifecycleStateDeleting,
	"deleted":         SchedulingPlanLifecycleStateDeleted,
}

// GetSchedulingPlanLifecycleStateEnumValues Enumerates the set of values for SchedulingPlanLifecycleStateEnum
func GetSchedulingPlanLifecycleStateEnumValues() []SchedulingPlanLifecycleStateEnum {
	values := make([]SchedulingPlanLifecycleStateEnum, 0)
	for _, v := range mappingSchedulingPlanLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSchedulingPlanLifecycleStateEnumStringValues Enumerates the set of values in String for SchedulingPlanLifecycleStateEnum
func GetSchedulingPlanLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"NEEDS_ATTENTION",
		"AVAILABLE",
		"UPDATING",
		"FAILED",
		"DELETING",
		"DELETED",
	}
}

// GetMappingSchedulingPlanLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSchedulingPlanLifecycleStateEnum(val string) (SchedulingPlanLifecycleStateEnum, bool) {
	enum, ok := mappingSchedulingPlanLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SchedulingPlanPlanIntentEnum Enum with underlying type: string
type SchedulingPlanPlanIntentEnum string

// Set of constants representing the allowable values for SchedulingPlanPlanIntentEnum
const (
	SchedulingPlanPlanIntentExadataInfrastructureFullSoftwareUpdate SchedulingPlanPlanIntentEnum = "EXADATA_INFRASTRUCTURE_FULL_SOFTWARE_UPDATE"
)

var mappingSchedulingPlanPlanIntentEnum = map[string]SchedulingPlanPlanIntentEnum{
	"EXADATA_INFRASTRUCTURE_FULL_SOFTWARE_UPDATE": SchedulingPlanPlanIntentExadataInfrastructureFullSoftwareUpdate,
}

var mappingSchedulingPlanPlanIntentEnumLowerCase = map[string]SchedulingPlanPlanIntentEnum{
	"exadata_infrastructure_full_software_update": SchedulingPlanPlanIntentExadataInfrastructureFullSoftwareUpdate,
}

// GetSchedulingPlanPlanIntentEnumValues Enumerates the set of values for SchedulingPlanPlanIntentEnum
func GetSchedulingPlanPlanIntentEnumValues() []SchedulingPlanPlanIntentEnum {
	values := make([]SchedulingPlanPlanIntentEnum, 0)
	for _, v := range mappingSchedulingPlanPlanIntentEnum {
		values = append(values, v)
	}
	return values
}

// GetSchedulingPlanPlanIntentEnumStringValues Enumerates the set of values in String for SchedulingPlanPlanIntentEnum
func GetSchedulingPlanPlanIntentEnumStringValues() []string {
	return []string{
		"EXADATA_INFRASTRUCTURE_FULL_SOFTWARE_UPDATE",
	}
}

// GetMappingSchedulingPlanPlanIntentEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSchedulingPlanPlanIntentEnum(val string) (SchedulingPlanPlanIntentEnum, bool) {
	enum, ok := mappingSchedulingPlanPlanIntentEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SchedulingPlanServiceTypeEnum Enum with underlying type: string
type SchedulingPlanServiceTypeEnum string

// Set of constants representing the allowable values for SchedulingPlanServiceTypeEnum
const (
	SchedulingPlanServiceTypeExacc  SchedulingPlanServiceTypeEnum = "EXACC"
	SchedulingPlanServiceTypeExacs  SchedulingPlanServiceTypeEnum = "EXACS"
	SchedulingPlanServiceTypeFpppcs SchedulingPlanServiceTypeEnum = "FPPPCS"
)

var mappingSchedulingPlanServiceTypeEnum = map[string]SchedulingPlanServiceTypeEnum{
	"EXACC":  SchedulingPlanServiceTypeExacc,
	"EXACS":  SchedulingPlanServiceTypeExacs,
	"FPPPCS": SchedulingPlanServiceTypeFpppcs,
}

var mappingSchedulingPlanServiceTypeEnumLowerCase = map[string]SchedulingPlanServiceTypeEnum{
	"exacc":  SchedulingPlanServiceTypeExacc,
	"exacs":  SchedulingPlanServiceTypeExacs,
	"fpppcs": SchedulingPlanServiceTypeFpppcs,
}

// GetSchedulingPlanServiceTypeEnumValues Enumerates the set of values for SchedulingPlanServiceTypeEnum
func GetSchedulingPlanServiceTypeEnumValues() []SchedulingPlanServiceTypeEnum {
	values := make([]SchedulingPlanServiceTypeEnum, 0)
	for _, v := range mappingSchedulingPlanServiceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSchedulingPlanServiceTypeEnumStringValues Enumerates the set of values in String for SchedulingPlanServiceTypeEnum
func GetSchedulingPlanServiceTypeEnumStringValues() []string {
	return []string{
		"EXACC",
		"EXACS",
		"FPPPCS",
	}
}

// GetMappingSchedulingPlanServiceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSchedulingPlanServiceTypeEnum(val string) (SchedulingPlanServiceTypeEnum, bool) {
	enum, ok := mappingSchedulingPlanServiceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
