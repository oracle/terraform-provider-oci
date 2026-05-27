// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateSchedulingPlanDetails Request to create Scheduling Plan.
type CreateSchedulingPlanDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Scheduling Policy.
	SchedulingPolicyId *string `mandatory:"true" json:"schedulingPolicyId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource.
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// The service type of the Scheduling Plan.
	ServiceType CreateSchedulingPlanDetailsServiceTypeEnum `mandatory:"true" json:"serviceType"`

	// If true, recommended scheduled actions will be generated for the scheduling plan.
	IsUsingRecommendedScheduledActions *bool `mandatory:"false" json:"isUsingRecommendedScheduledActions"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The current intent of the Scheduling Plan. Valid states are EXADATA_INFRASTRUCTURE_FULL_SOFTWARE_UPDATE and EXADATA_INFRASTRUCTURE_SECURITY_UPDATE.
	PlanIntent CreateSchedulingPlanDetailsPlanIntentEnum `mandatory:"false" json:"planIntent,omitempty"`
}

func (m CreateSchedulingPlanDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateSchedulingPlanDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateSchedulingPlanDetailsServiceTypeEnum(string(m.ServiceType)); !ok && m.ServiceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ServiceType: %s. Supported values are: %s.", m.ServiceType, strings.Join(GetCreateSchedulingPlanDetailsServiceTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingCreateSchedulingPlanDetailsPlanIntentEnum(string(m.PlanIntent)); !ok && m.PlanIntent != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlanIntent: %s. Supported values are: %s.", m.PlanIntent, strings.Join(GetCreateSchedulingPlanDetailsPlanIntentEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateSchedulingPlanDetailsServiceTypeEnum Enum with underlying type: string
type CreateSchedulingPlanDetailsServiceTypeEnum string

// Set of constants representing the allowable values for CreateSchedulingPlanDetailsServiceTypeEnum
const (
	CreateSchedulingPlanDetailsServiceTypeExacc  CreateSchedulingPlanDetailsServiceTypeEnum = "EXACC"
	CreateSchedulingPlanDetailsServiceTypeExacs  CreateSchedulingPlanDetailsServiceTypeEnum = "EXACS"
	CreateSchedulingPlanDetailsServiceTypeFpppcs CreateSchedulingPlanDetailsServiceTypeEnum = "FPPPCS"
)

var mappingCreateSchedulingPlanDetailsServiceTypeEnum = map[string]CreateSchedulingPlanDetailsServiceTypeEnum{
	"EXACC":  CreateSchedulingPlanDetailsServiceTypeExacc,
	"EXACS":  CreateSchedulingPlanDetailsServiceTypeExacs,
	"FPPPCS": CreateSchedulingPlanDetailsServiceTypeFpppcs,
}

var mappingCreateSchedulingPlanDetailsServiceTypeEnumLowerCase = map[string]CreateSchedulingPlanDetailsServiceTypeEnum{
	"exacc":  CreateSchedulingPlanDetailsServiceTypeExacc,
	"exacs":  CreateSchedulingPlanDetailsServiceTypeExacs,
	"fpppcs": CreateSchedulingPlanDetailsServiceTypeFpppcs,
}

// GetCreateSchedulingPlanDetailsServiceTypeEnumValues Enumerates the set of values for CreateSchedulingPlanDetailsServiceTypeEnum
func GetCreateSchedulingPlanDetailsServiceTypeEnumValues() []CreateSchedulingPlanDetailsServiceTypeEnum {
	values := make([]CreateSchedulingPlanDetailsServiceTypeEnum, 0)
	for _, v := range mappingCreateSchedulingPlanDetailsServiceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateSchedulingPlanDetailsServiceTypeEnumStringValues Enumerates the set of values in String for CreateSchedulingPlanDetailsServiceTypeEnum
func GetCreateSchedulingPlanDetailsServiceTypeEnumStringValues() []string {
	return []string{
		"EXACC",
		"EXACS",
		"FPPPCS",
	}
}

// GetMappingCreateSchedulingPlanDetailsServiceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateSchedulingPlanDetailsServiceTypeEnum(val string) (CreateSchedulingPlanDetailsServiceTypeEnum, bool) {
	enum, ok := mappingCreateSchedulingPlanDetailsServiceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateSchedulingPlanDetailsPlanIntentEnum Enum with underlying type: string
type CreateSchedulingPlanDetailsPlanIntentEnum string

// Set of constants representing the allowable values for CreateSchedulingPlanDetailsPlanIntentEnum
const (
	CreateSchedulingPlanDetailsPlanIntentFullSoftwareUpdate CreateSchedulingPlanDetailsPlanIntentEnum = "EXADATA_INFRASTRUCTURE_FULL_SOFTWARE_UPDATE"
	CreateSchedulingPlanDetailsPlanIntentSecurityUpdate     CreateSchedulingPlanDetailsPlanIntentEnum = "EXADATA_INFRASTRUCTURE_SECURITY_UPDATE"
)

var mappingCreateSchedulingPlanDetailsPlanIntentEnum = map[string]CreateSchedulingPlanDetailsPlanIntentEnum{
	"EXADATA_INFRASTRUCTURE_FULL_SOFTWARE_UPDATE": CreateSchedulingPlanDetailsPlanIntentFullSoftwareUpdate,
	"EXADATA_INFRASTRUCTURE_SECURITY_UPDATE":      CreateSchedulingPlanDetailsPlanIntentSecurityUpdate,
}

var mappingCreateSchedulingPlanDetailsPlanIntentEnumLowerCase = map[string]CreateSchedulingPlanDetailsPlanIntentEnum{
	"exadata_infrastructure_full_software_update": CreateSchedulingPlanDetailsPlanIntentFullSoftwareUpdate,
	"exadata_infrastructure_security_update":      CreateSchedulingPlanDetailsPlanIntentSecurityUpdate,
}

// GetCreateSchedulingPlanDetailsPlanIntentEnumValues Enumerates the set of values for CreateSchedulingPlanDetailsPlanIntentEnum
func GetCreateSchedulingPlanDetailsPlanIntentEnumValues() []CreateSchedulingPlanDetailsPlanIntentEnum {
	values := make([]CreateSchedulingPlanDetailsPlanIntentEnum, 0)
	for _, v := range mappingCreateSchedulingPlanDetailsPlanIntentEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateSchedulingPlanDetailsPlanIntentEnumStringValues Enumerates the set of values in String for CreateSchedulingPlanDetailsPlanIntentEnum
func GetCreateSchedulingPlanDetailsPlanIntentEnumStringValues() []string {
	return []string{
		"EXADATA_INFRASTRUCTURE_FULL_SOFTWARE_UPDATE",
		"EXADATA_INFRASTRUCTURE_SECURITY_UPDATE",
	}
}

// GetMappingCreateSchedulingPlanDetailsPlanIntentEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateSchedulingPlanDetailsPlanIntentEnum(val string) (CreateSchedulingPlanDetailsPlanIntentEnum, bool) {
	enum, ok := mappingCreateSchedulingPlanDetailsPlanIntentEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
