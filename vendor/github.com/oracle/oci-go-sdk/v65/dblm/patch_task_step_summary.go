// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Dblm API
//
// A description of the Dblm API
//

package dblm

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PatchTaskStepSummary Summary of a patch task step.
type PatchTaskStepSummary struct {

	// Unique identifier that is system generated
	Key *int64 `mandatory:"true" json:"key"`

	// Unique patch task Identifier
	TaskKey *int64 `mandatory:"true" json:"taskKey"`

	// PatchOperation Identifier which contains this patch task step
	OperationId *string `mandatory:"true" json:"operationId"`

	// The type of the PatchTaskStep
	Type PatchTaskStepTypeEnum `mandatory:"true" json:"type"`

	// The work state of the PatchTaskStep.
	WorkState PatchTaskStepWorkStateEnum `mandatory:"true" json:"workState"`

	// The status of the PatchTaskStep.
	Status PatchTaskStepStatusEnum `mandatory:"true" json:"status"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// The Name of the PatchTaskStep
	Name *string `mandatory:"false" json:"name"`

	// Detailed description of step
	Description *string `mandatory:"false" json:"description"`

	// The time the PatchTaskStep actually started. An RFC3339 formatted datetime string
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The time the PatchTaskStep actually completed. An RFC3339 formatted datetime string
	TimeCompleted *common.SDKTime `mandatory:"false" json:"timeCompleted"`

	// The elapsed time for the PatchTaskStep in seconds.
	TimeElapsedInSeconds *int64 `mandatory:"false" json:"timeElapsedInSeconds"`

	// If true, the patchTaskStep is part of ROLLBACK/CLEANUP tasks
	IsRollback *bool `mandatory:"false" json:"isRollback"`

	Properties *PatchTaskStepPropertiesList `mandatory:"false" json:"properties"`

	// The type of the PatchTaskStep
	OutputType PatchTaskStepSummaryOutputTypeEnum `mandatory:"false" json:"outputType,omitempty"`

	Output *PatchTaskStepOutputList `mandatory:"false" json:"output"`

	ValidationOutput *PatchTaskStepValidationMessages `mandatory:"false" json:"validationOutput"`

	// Location of log file for the step
	LogFileLocation *string `mandatory:"false" json:"logFileLocation"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m PatchTaskStepSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PatchTaskStepSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPatchTaskStepTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetPatchTaskStepTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPatchTaskStepWorkStateEnum(string(m.WorkState)); !ok && m.WorkState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for WorkState: %s. Supported values are: %s.", m.WorkState, strings.Join(GetPatchTaskStepWorkStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPatchTaskStepStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetPatchTaskStepStatusEnumStringValues(), ",")))
	}

	if _, ok := GetMappingPatchTaskStepSummaryOutputTypeEnum(string(m.OutputType)); !ok && m.OutputType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OutputType: %s. Supported values are: %s.", m.OutputType, strings.Join(GetPatchTaskStepSummaryOutputTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PatchTaskStepSummaryOutputTypeEnum Enum with underlying type: string
type PatchTaskStepSummaryOutputTypeEnum string

// Set of constants representing the allowable values for PatchTaskStepSummaryOutputTypeEnum
const (
	PatchTaskStepSummaryOutputTypeValidation PatchTaskStepSummaryOutputTypeEnum = "VALIDATION"
	PatchTaskStepSummaryOutputTypeMessages   PatchTaskStepSummaryOutputTypeEnum = "MESSAGES"
)

var mappingPatchTaskStepSummaryOutputTypeEnum = map[string]PatchTaskStepSummaryOutputTypeEnum{
	"VALIDATION": PatchTaskStepSummaryOutputTypeValidation,
	"MESSAGES":   PatchTaskStepSummaryOutputTypeMessages,
}

var mappingPatchTaskStepSummaryOutputTypeEnumLowerCase = map[string]PatchTaskStepSummaryOutputTypeEnum{
	"validation": PatchTaskStepSummaryOutputTypeValidation,
	"messages":   PatchTaskStepSummaryOutputTypeMessages,
}

// GetPatchTaskStepSummaryOutputTypeEnumValues Enumerates the set of values for PatchTaskStepSummaryOutputTypeEnum
func GetPatchTaskStepSummaryOutputTypeEnumValues() []PatchTaskStepSummaryOutputTypeEnum {
	values := make([]PatchTaskStepSummaryOutputTypeEnum, 0)
	for _, v := range mappingPatchTaskStepSummaryOutputTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchTaskStepSummaryOutputTypeEnumStringValues Enumerates the set of values in String for PatchTaskStepSummaryOutputTypeEnum
func GetPatchTaskStepSummaryOutputTypeEnumStringValues() []string {
	return []string{
		"VALIDATION",
		"MESSAGES",
	}
}

// GetMappingPatchTaskStepSummaryOutputTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchTaskStepSummaryOutputTypeEnum(val string) (PatchTaskStepSummaryOutputTypeEnum, bool) {
	enum, ok := mappingPatchTaskStepSummaryOutputTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
