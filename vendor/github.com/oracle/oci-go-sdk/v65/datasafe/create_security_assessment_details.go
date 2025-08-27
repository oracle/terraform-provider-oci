// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// CreateSecurityAssessmentDetails The details used to save a security assessment.
type CreateSecurityAssessmentDetails struct {

	// The OCID of the compartment that contains the security assessment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of the security assessment.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of the security assessment.
	Description *string `mandatory:"false" json:"description"`

	// The OCID of the target database or target database group on which security assessment is to be run.
	TargetId *string `mandatory:"false" json:"targetId"`

	// The type of security assessment resource whether it is individual or group resource. For individual target use type TARGET_DATABASE and for group resource use type TARGET_DATABASE_GROUP. If not provided, TARGET_DATABASE would be used as default value.
	TargetType SecurityAssessmentTargetTypeEnum `mandatory:"false" json:"targetType,omitempty"`

	// The type of the security assessment
	Type CreateSecurityAssessmentDetailsTypeEnum `mandatory:"false" json:"type,omitempty"`

	// The OCID of the template assessment. It will be required while creating the template baseline assessment.
	TemplateAssessmentId *string `mandatory:"false" json:"templateAssessmentId"`

	// The OCID of the security assessment. The assessment should be of type SAVED.
	// It will be required while creating the template baseline assessment for individual targets to fetch the detailed information from an existing security assessment.
	BaseSecurityAssessmentId *string `mandatory:"false" json:"baseSecurityAssessmentId"`

	// Indicates whether the assessment is scheduled to run.
	IsAssessmentScheduled *bool `mandatory:"false" json:"isAssessmentScheduled"`

	// To schedule the assessment for running periodically, specify the schedule in this attribute.
	// Create or schedule one assessment per compartment. If not defined, the assessment runs immediately.
	// Format -
	// <version-string>;<version-specific-schedule>
	// Allowed version strings - "v1"
	// v1's version specific schedule -<ss> <mm> <hh> <day-of-week> <day-of-month>
	// Each of the above fields potentially introduce constraints. A workrequest is created only
	// when clock time satisfies all the constraints. Constraints introduced:
	// 1. seconds = <ss> (So, the allowed range for <ss> is [0, 59])
	// 2. minutes = <mm> (So, the allowed range for <mm> is [0, 59])
	// 3. hours = <hh> (So, the allowed range for <hh> is [0, 23])
	// <day-of-week> can be either '*' (without quotes or a number between 1(Monday) and 7(Sunday))
	// 4. No constraint introduced when it is '*'. When not, day of week must equal the given value
	// <day-of-month> can be either '*' (without quotes or a number between 1 and 28)
	// 5. No constraint introduced when it is '*'. When not, day of month must equal the given value
	Schedule *string `mandatory:"false" json:"schedule"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateSecurityAssessmentDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateSecurityAssessmentDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSecurityAssessmentTargetTypeEnum(string(m.TargetType)); !ok && m.TargetType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TargetType: %s. Supported values are: %s.", m.TargetType, strings.Join(GetSecurityAssessmentTargetTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateSecurityAssessmentDetailsTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetCreateSecurityAssessmentDetailsTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateSecurityAssessmentDetailsTypeEnum Enum with underlying type: string
type CreateSecurityAssessmentDetailsTypeEnum string

// Set of constants representing the allowable values for CreateSecurityAssessmentDetailsTypeEnum
const (
	CreateSecurityAssessmentDetailsTypeLatest           CreateSecurityAssessmentDetailsTypeEnum = "LATEST"
	CreateSecurityAssessmentDetailsTypeSaved            CreateSecurityAssessmentDetailsTypeEnum = "SAVED"
	CreateSecurityAssessmentDetailsTypeSaveSchedule     CreateSecurityAssessmentDetailsTypeEnum = "SAVE_SCHEDULE"
	CreateSecurityAssessmentDetailsTypeCompartment      CreateSecurityAssessmentDetailsTypeEnum = "COMPARTMENT"
	CreateSecurityAssessmentDetailsTypeTemplate         CreateSecurityAssessmentDetailsTypeEnum = "TEMPLATE"
	CreateSecurityAssessmentDetailsTypeTemplateBaseline CreateSecurityAssessmentDetailsTypeEnum = "TEMPLATE_BASELINE"
)

var mappingCreateSecurityAssessmentDetailsTypeEnum = map[string]CreateSecurityAssessmentDetailsTypeEnum{
	"LATEST":            CreateSecurityAssessmentDetailsTypeLatest,
	"SAVED":             CreateSecurityAssessmentDetailsTypeSaved,
	"SAVE_SCHEDULE":     CreateSecurityAssessmentDetailsTypeSaveSchedule,
	"COMPARTMENT":       CreateSecurityAssessmentDetailsTypeCompartment,
	"TEMPLATE":          CreateSecurityAssessmentDetailsTypeTemplate,
	"TEMPLATE_BASELINE": CreateSecurityAssessmentDetailsTypeTemplateBaseline,
}

var mappingCreateSecurityAssessmentDetailsTypeEnumLowerCase = map[string]CreateSecurityAssessmentDetailsTypeEnum{
	"latest":            CreateSecurityAssessmentDetailsTypeLatest,
	"saved":             CreateSecurityAssessmentDetailsTypeSaved,
	"save_schedule":     CreateSecurityAssessmentDetailsTypeSaveSchedule,
	"compartment":       CreateSecurityAssessmentDetailsTypeCompartment,
	"template":          CreateSecurityAssessmentDetailsTypeTemplate,
	"template_baseline": CreateSecurityAssessmentDetailsTypeTemplateBaseline,
}

// GetCreateSecurityAssessmentDetailsTypeEnumValues Enumerates the set of values for CreateSecurityAssessmentDetailsTypeEnum
func GetCreateSecurityAssessmentDetailsTypeEnumValues() []CreateSecurityAssessmentDetailsTypeEnum {
	values := make([]CreateSecurityAssessmentDetailsTypeEnum, 0)
	for _, v := range mappingCreateSecurityAssessmentDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateSecurityAssessmentDetailsTypeEnumStringValues Enumerates the set of values in String for CreateSecurityAssessmentDetailsTypeEnum
func GetCreateSecurityAssessmentDetailsTypeEnumStringValues() []string {
	return []string{
		"LATEST",
		"SAVED",
		"SAVE_SCHEDULE",
		"COMPARTMENT",
		"TEMPLATE",
		"TEMPLATE_BASELINE",
	}
}

// GetMappingCreateSecurityAssessmentDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateSecurityAssessmentDetailsTypeEnum(val string) (CreateSecurityAssessmentDetailsTypeEnum, bool) {
	enum, ok := mappingCreateSecurityAssessmentDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
