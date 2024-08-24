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

// SchedulingPlanSummary Details of a Scheduling Plan.
type SchedulingPlanSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Scheduling Plan.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Scheduling Policy.
	SchedulingPolicyId *string `mandatory:"true" json:"schedulingPolicyId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the resource.
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// The current state of the Scheduling Plan. Valid states are CREATING, NEEDS_ATTENTION, AVAILABLE, UPDATING, FAILED, DELETING and DELETED.
	LifecycleState SchedulingPlanSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The service type of the Scheduling Plan.
	ServiceType SchedulingPlanSummaryServiceTypeEnum `mandatory:"true" json:"serviceType"`

	// The date and time the Scheduling Plan Resource was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The display name of the Scheduling Plan.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// If true, recommended scheduled actions will be generated for the scheduling plan.
	IsUsingRecommendedScheduledActions *bool `mandatory:"false" json:"isUsingRecommendedScheduledActions"`

	// The current intent the Scheduling Plan. Valid states is EXADATA_INFRASTRUCTURE_FULL_SOFTWARE_UPDATE.
	PlanIntent SchedulingPlanSummaryPlanIntentEnum `mandatory:"false" json:"planIntent,omitempty"`

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

func (m SchedulingPlanSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SchedulingPlanSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSchedulingPlanSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSchedulingPlanSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSchedulingPlanSummaryServiceTypeEnum(string(m.ServiceType)); !ok && m.ServiceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ServiceType: %s. Supported values are: %s.", m.ServiceType, strings.Join(GetSchedulingPlanSummaryServiceTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingSchedulingPlanSummaryPlanIntentEnum(string(m.PlanIntent)); !ok && m.PlanIntent != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlanIntent: %s. Supported values are: %s.", m.PlanIntent, strings.Join(GetSchedulingPlanSummaryPlanIntentEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SchedulingPlanSummaryLifecycleStateEnum Enum with underlying type: string
type SchedulingPlanSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for SchedulingPlanSummaryLifecycleStateEnum
const (
	SchedulingPlanSummaryLifecycleStateCreating       SchedulingPlanSummaryLifecycleStateEnum = "CREATING"
	SchedulingPlanSummaryLifecycleStateNeedsAttention SchedulingPlanSummaryLifecycleStateEnum = "NEEDS_ATTENTION"
	SchedulingPlanSummaryLifecycleStateAvailable      SchedulingPlanSummaryLifecycleStateEnum = "AVAILABLE"
	SchedulingPlanSummaryLifecycleStateUpdating       SchedulingPlanSummaryLifecycleStateEnum = "UPDATING"
	SchedulingPlanSummaryLifecycleStateFailed         SchedulingPlanSummaryLifecycleStateEnum = "FAILED"
	SchedulingPlanSummaryLifecycleStateDeleting       SchedulingPlanSummaryLifecycleStateEnum = "DELETING"
	SchedulingPlanSummaryLifecycleStateDeleted        SchedulingPlanSummaryLifecycleStateEnum = "DELETED"
)

var mappingSchedulingPlanSummaryLifecycleStateEnum = map[string]SchedulingPlanSummaryLifecycleStateEnum{
	"CREATING":        SchedulingPlanSummaryLifecycleStateCreating,
	"NEEDS_ATTENTION": SchedulingPlanSummaryLifecycleStateNeedsAttention,
	"AVAILABLE":       SchedulingPlanSummaryLifecycleStateAvailable,
	"UPDATING":        SchedulingPlanSummaryLifecycleStateUpdating,
	"FAILED":          SchedulingPlanSummaryLifecycleStateFailed,
	"DELETING":        SchedulingPlanSummaryLifecycleStateDeleting,
	"DELETED":         SchedulingPlanSummaryLifecycleStateDeleted,
}

var mappingSchedulingPlanSummaryLifecycleStateEnumLowerCase = map[string]SchedulingPlanSummaryLifecycleStateEnum{
	"creating":        SchedulingPlanSummaryLifecycleStateCreating,
	"needs_attention": SchedulingPlanSummaryLifecycleStateNeedsAttention,
	"available":       SchedulingPlanSummaryLifecycleStateAvailable,
	"updating":        SchedulingPlanSummaryLifecycleStateUpdating,
	"failed":          SchedulingPlanSummaryLifecycleStateFailed,
	"deleting":        SchedulingPlanSummaryLifecycleStateDeleting,
	"deleted":         SchedulingPlanSummaryLifecycleStateDeleted,
}

// GetSchedulingPlanSummaryLifecycleStateEnumValues Enumerates the set of values for SchedulingPlanSummaryLifecycleStateEnum
func GetSchedulingPlanSummaryLifecycleStateEnumValues() []SchedulingPlanSummaryLifecycleStateEnum {
	values := make([]SchedulingPlanSummaryLifecycleStateEnum, 0)
	for _, v := range mappingSchedulingPlanSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSchedulingPlanSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for SchedulingPlanSummaryLifecycleStateEnum
func GetSchedulingPlanSummaryLifecycleStateEnumStringValues() []string {
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

// GetMappingSchedulingPlanSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSchedulingPlanSummaryLifecycleStateEnum(val string) (SchedulingPlanSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingSchedulingPlanSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SchedulingPlanSummaryPlanIntentEnum Enum with underlying type: string
type SchedulingPlanSummaryPlanIntentEnum string

// Set of constants representing the allowable values for SchedulingPlanSummaryPlanIntentEnum
const (
	SchedulingPlanSummaryPlanIntentExadataInfrastructureFullSoftwareUpdate SchedulingPlanSummaryPlanIntentEnum = "EXADATA_INFRASTRUCTURE_FULL_SOFTWARE_UPDATE"
)

var mappingSchedulingPlanSummaryPlanIntentEnum = map[string]SchedulingPlanSummaryPlanIntentEnum{
	"EXADATA_INFRASTRUCTURE_FULL_SOFTWARE_UPDATE": SchedulingPlanSummaryPlanIntentExadataInfrastructureFullSoftwareUpdate,
}

var mappingSchedulingPlanSummaryPlanIntentEnumLowerCase = map[string]SchedulingPlanSummaryPlanIntentEnum{
	"exadata_infrastructure_full_software_update": SchedulingPlanSummaryPlanIntentExadataInfrastructureFullSoftwareUpdate,
}

// GetSchedulingPlanSummaryPlanIntentEnumValues Enumerates the set of values for SchedulingPlanSummaryPlanIntentEnum
func GetSchedulingPlanSummaryPlanIntentEnumValues() []SchedulingPlanSummaryPlanIntentEnum {
	values := make([]SchedulingPlanSummaryPlanIntentEnum, 0)
	for _, v := range mappingSchedulingPlanSummaryPlanIntentEnum {
		values = append(values, v)
	}
	return values
}

// GetSchedulingPlanSummaryPlanIntentEnumStringValues Enumerates the set of values in String for SchedulingPlanSummaryPlanIntentEnum
func GetSchedulingPlanSummaryPlanIntentEnumStringValues() []string {
	return []string{
		"EXADATA_INFRASTRUCTURE_FULL_SOFTWARE_UPDATE",
	}
}

// GetMappingSchedulingPlanSummaryPlanIntentEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSchedulingPlanSummaryPlanIntentEnum(val string) (SchedulingPlanSummaryPlanIntentEnum, bool) {
	enum, ok := mappingSchedulingPlanSummaryPlanIntentEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SchedulingPlanSummaryServiceTypeEnum Enum with underlying type: string
type SchedulingPlanSummaryServiceTypeEnum string

// Set of constants representing the allowable values for SchedulingPlanSummaryServiceTypeEnum
const (
	SchedulingPlanSummaryServiceTypeExacc  SchedulingPlanSummaryServiceTypeEnum = "EXACC"
	SchedulingPlanSummaryServiceTypeExacs  SchedulingPlanSummaryServiceTypeEnum = "EXACS"
	SchedulingPlanSummaryServiceTypeFpppcs SchedulingPlanSummaryServiceTypeEnum = "FPPPCS"
)

var mappingSchedulingPlanSummaryServiceTypeEnum = map[string]SchedulingPlanSummaryServiceTypeEnum{
	"EXACC":  SchedulingPlanSummaryServiceTypeExacc,
	"EXACS":  SchedulingPlanSummaryServiceTypeExacs,
	"FPPPCS": SchedulingPlanSummaryServiceTypeFpppcs,
}

var mappingSchedulingPlanSummaryServiceTypeEnumLowerCase = map[string]SchedulingPlanSummaryServiceTypeEnum{
	"exacc":  SchedulingPlanSummaryServiceTypeExacc,
	"exacs":  SchedulingPlanSummaryServiceTypeExacs,
	"fpppcs": SchedulingPlanSummaryServiceTypeFpppcs,
}

// GetSchedulingPlanSummaryServiceTypeEnumValues Enumerates the set of values for SchedulingPlanSummaryServiceTypeEnum
func GetSchedulingPlanSummaryServiceTypeEnumValues() []SchedulingPlanSummaryServiceTypeEnum {
	values := make([]SchedulingPlanSummaryServiceTypeEnum, 0)
	for _, v := range mappingSchedulingPlanSummaryServiceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSchedulingPlanSummaryServiceTypeEnumStringValues Enumerates the set of values in String for SchedulingPlanSummaryServiceTypeEnum
func GetSchedulingPlanSummaryServiceTypeEnumStringValues() []string {
	return []string{
		"EXACC",
		"EXACS",
		"FPPPCS",
	}
}

// GetMappingSchedulingPlanSummaryServiceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSchedulingPlanSummaryServiceTypeEnum(val string) (SchedulingPlanSummaryServiceTypeEnum, bool) {
	enum, ok := mappingSchedulingPlanSummaryServiceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
