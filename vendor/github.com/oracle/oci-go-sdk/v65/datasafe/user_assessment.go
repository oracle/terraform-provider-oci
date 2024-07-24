// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UserAssessment The details of the user assessment, which includes statistics related to target database users.
type UserAssessment struct {

	// The OCID of the compartment that contains the user assessment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of the user assessment.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the user assessment.
	Id *string `mandatory:"true" json:"id"`

	// The current state of the user assessment.
	LifecycleState UserAssessmentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the user assessment was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the user assessment was last updated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The type of the user assessment. The possible types are:
	// LATEST: The latest assessment that was executed for a target. It can either be system generated as part of the scheduled assessments or user driven by refreshing the latest assessment.
	// SAVED: A saved user assessment. All user assessments are saved in the user assessment history.
	// SAVE_SCHEDULE: The schedule to periodically save the LATEST assessment of a target database.
	// COMPARTMENT: An automatic managed assessment type that stores all details of the targets in one compartment. This will keep an up-to-date status of all potential risks identified in the compartment.
	//        It also keeps track of user count and target count for each profile available on the targets in a given compartment.
	//        It is automatically updated once the latest assessment or refresh action is executed, as well as when a target is deleted or moved to a different compartment.
	Type UserAssessmentTypeEnum `mandatory:"true" json:"type"`

	// The description of the user assessment.
	Description *string `mandatory:"false" json:"description"`

	// List containing maps as values.
	// Example: `{"Operations": [ {"CostCenter": "42"} ] }`
	IgnoredTargets []interface{} `mandatory:"false" json:"ignoredTargets"`

	// List containing maps as values.
	// Example: `{"Operations": [ {"CostCenter": "42"} ] }`
	IgnoredAssessmentIds []interface{} `mandatory:"false" json:"ignoredAssessmentIds"`

	// Indicates if the user assessment is set as a baseline. This is applicable only to saved user assessments.
	IsBaseline *bool `mandatory:"false" json:"isBaseline"`

	// Indicates if the user assessment deviates from the baseline.
	IsDeviatedFromBaseline *bool `mandatory:"false" json:"isDeviatedFromBaseline"`

	// The OCID of the last user assessment baseline against which the latest assessment was compared.
	LastComparedBaselineId *string `mandatory:"false" json:"lastComparedBaselineId"`

	// Details about the current state of the user assessment.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The OCID of the user assessment that is responsible for creating this scheduled save assessment.
	ScheduleAssessmentId *string `mandatory:"false" json:"scheduleAssessmentId"`

	// Indicates whether the assessment is scheduled to run.
	IsAssessmentScheduled *bool `mandatory:"false" json:"isAssessmentScheduled"`

	// Schedule of the assessment that runs periodically in this specified format:
	//   <version-string>;<version-specific-schedule>
	//   Allowed version strings - "v1"
	//   v1's version specific schedule -<ss> <mm> <hh> <day-of-week> <day-of-month>
	//   Each of the above fields potentially introduce constraints. A workrequest is created only
	//   when clock time satisfies all the constraints. Constraints introduced:
	//   1. seconds = <ss> (So, the allowed range for <ss> is [0, 59])
	//   2. minutes = <mm> (So, the allowed range for <mm> is [0, 59])
	//   3. hours = <hh> (So, the allowed range for <hh> is [0, 23])
	//   <day-of-week> can be either '*' (without quotes or a number between 1(Monday) and 7(Sunday))
	//   4. No constraint introduced when it is '*'. When not, day of week must equal the given value
	//   <day-of-month> can be either '*' (without quotes or a number between 1 and 28)
	//   5. No constraint introduced when it is '*'. When not, day of month must equal the given value
	Schedule *string `mandatory:"false" json:"schedule"`

	// Map that contains maps of values.
	//  Example: `{"Operations": {"CostCenter": "42"}}`
	Statistics map[string]map[string]interface{} `mandatory:"false" json:"statistics"`

	// Array of database target OCIDs.
	TargetIds []string `mandatory:"false" json:"targetIds"`

	// The date and time the user assessment was last executed, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeLastAssessed *common.SDKTime `mandatory:"false" json:"timeLastAssessed"`

	// Indicates whether the user assessment was created by the system or the user.
	TriggeredBy UserAssessmentTriggeredByEnum `mandatory:"false" json:"triggeredBy,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m UserAssessment) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UserAssessment) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUserAssessmentLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetUserAssessmentLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUserAssessmentTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetUserAssessmentTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingUserAssessmentTriggeredByEnum(string(m.TriggeredBy)); !ok && m.TriggeredBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TriggeredBy: %s. Supported values are: %s.", m.TriggeredBy, strings.Join(GetUserAssessmentTriggeredByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UserAssessmentTriggeredByEnum Enum with underlying type: string
type UserAssessmentTriggeredByEnum string

// Set of constants representing the allowable values for UserAssessmentTriggeredByEnum
const (
	UserAssessmentTriggeredByUser   UserAssessmentTriggeredByEnum = "USER"
	UserAssessmentTriggeredBySystem UserAssessmentTriggeredByEnum = "SYSTEM"
)

var mappingUserAssessmentTriggeredByEnum = map[string]UserAssessmentTriggeredByEnum{
	"USER":   UserAssessmentTriggeredByUser,
	"SYSTEM": UserAssessmentTriggeredBySystem,
}

var mappingUserAssessmentTriggeredByEnumLowerCase = map[string]UserAssessmentTriggeredByEnum{
	"user":   UserAssessmentTriggeredByUser,
	"system": UserAssessmentTriggeredBySystem,
}

// GetUserAssessmentTriggeredByEnumValues Enumerates the set of values for UserAssessmentTriggeredByEnum
func GetUserAssessmentTriggeredByEnumValues() []UserAssessmentTriggeredByEnum {
	values := make([]UserAssessmentTriggeredByEnum, 0)
	for _, v := range mappingUserAssessmentTriggeredByEnum {
		values = append(values, v)
	}
	return values
}

// GetUserAssessmentTriggeredByEnumStringValues Enumerates the set of values in String for UserAssessmentTriggeredByEnum
func GetUserAssessmentTriggeredByEnumStringValues() []string {
	return []string{
		"USER",
		"SYSTEM",
	}
}

// GetMappingUserAssessmentTriggeredByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUserAssessmentTriggeredByEnum(val string) (UserAssessmentTriggeredByEnum, bool) {
	enum, ok := mappingUserAssessmentTriggeredByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UserAssessmentTypeEnum Enum with underlying type: string
type UserAssessmentTypeEnum string

// Set of constants representing the allowable values for UserAssessmentTypeEnum
const (
	UserAssessmentTypeLatest       UserAssessmentTypeEnum = "LATEST"
	UserAssessmentTypeSaved        UserAssessmentTypeEnum = "SAVED"
	UserAssessmentTypeSaveSchedule UserAssessmentTypeEnum = "SAVE_SCHEDULE"
	UserAssessmentTypeCompartment  UserAssessmentTypeEnum = "COMPARTMENT"
)

var mappingUserAssessmentTypeEnum = map[string]UserAssessmentTypeEnum{
	"LATEST":        UserAssessmentTypeLatest,
	"SAVED":         UserAssessmentTypeSaved,
	"SAVE_SCHEDULE": UserAssessmentTypeSaveSchedule,
	"COMPARTMENT":   UserAssessmentTypeCompartment,
}

var mappingUserAssessmentTypeEnumLowerCase = map[string]UserAssessmentTypeEnum{
	"latest":        UserAssessmentTypeLatest,
	"saved":         UserAssessmentTypeSaved,
	"save_schedule": UserAssessmentTypeSaveSchedule,
	"compartment":   UserAssessmentTypeCompartment,
}

// GetUserAssessmentTypeEnumValues Enumerates the set of values for UserAssessmentTypeEnum
func GetUserAssessmentTypeEnumValues() []UserAssessmentTypeEnum {
	values := make([]UserAssessmentTypeEnum, 0)
	for _, v := range mappingUserAssessmentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUserAssessmentTypeEnumStringValues Enumerates the set of values in String for UserAssessmentTypeEnum
func GetUserAssessmentTypeEnumStringValues() []string {
	return []string{
		"LATEST",
		"SAVED",
		"SAVE_SCHEDULE",
		"COMPARTMENT",
	}
}

// GetMappingUserAssessmentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUserAssessmentTypeEnum(val string) (UserAssessmentTypeEnum, bool) {
	enum, ok := mappingUserAssessmentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
