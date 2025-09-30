// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// RunbookImportDependencyFinderRequiredValue Runbook import required value.
type RunbookImportDependencyFinderRequiredValue struct {

	// Field for required value.
	Field *string `mandatory:"true" json:"field"`

	// Rendering component value for required value.
	RenderingComponent RunbookImportDependencyFinderRequiredValueRenderingComponentEnum `mandatory:"false" json:"renderingComponent,omitempty"`

	// Label for required value.
	Label *string `mandatory:"false" json:"label"`

	// Step name for required value.
	StepName *string `mandatory:"false" json:"stepName"`

	// Placeholder assignment for required value.
	PlaceholderAssignment *string `mandatory:"false" json:"placeholderAssignment"`

	// Expected assignment.
	ExpectedAssignment RunbookImportDependencyFinderRequiredValueExpectedAssignmentEnum `mandatory:"false" json:"expectedAssignment,omitempty"`
}

func (m RunbookImportDependencyFinderRequiredValue) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RunbookImportDependencyFinderRequiredValue) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRunbookImportDependencyFinderRequiredValueRenderingComponentEnum(string(m.RenderingComponent)); !ok && m.RenderingComponent != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RenderingComponent: %s. Supported values are: %s.", m.RenderingComponent, strings.Join(GetRunbookImportDependencyFinderRequiredValueRenderingComponentEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRunbookImportDependencyFinderRequiredValueExpectedAssignmentEnum(string(m.ExpectedAssignment)); !ok && m.ExpectedAssignment != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExpectedAssignment: %s. Supported values are: %s.", m.ExpectedAssignment, strings.Join(GetRunbookImportDependencyFinderRequiredValueExpectedAssignmentEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RunbookImportDependencyFinderRequiredValueRenderingComponentEnum Enum with underlying type: string
type RunbookImportDependencyFinderRequiredValueRenderingComponentEnum string

// Set of constants representing the allowable values for RunbookImportDependencyFinderRequiredValueRenderingComponentEnum
const (
	RunbookImportDependencyFinderRequiredValueRenderingComponentSingleChoice RunbookImportDependencyFinderRequiredValueRenderingComponentEnum = "SINGLE_CHOICE"
	RunbookImportDependencyFinderRequiredValueRenderingComponentMultiChoice  RunbookImportDependencyFinderRequiredValueRenderingComponentEnum = "MULTI_CHOICE"
	RunbookImportDependencyFinderRequiredValueRenderingComponentInputText    RunbookImportDependencyFinderRequiredValueRenderingComponentEnum = "INPUT_TEXT"
)

var mappingRunbookImportDependencyFinderRequiredValueRenderingComponentEnum = map[string]RunbookImportDependencyFinderRequiredValueRenderingComponentEnum{
	"SINGLE_CHOICE": RunbookImportDependencyFinderRequiredValueRenderingComponentSingleChoice,
	"MULTI_CHOICE":  RunbookImportDependencyFinderRequiredValueRenderingComponentMultiChoice,
	"INPUT_TEXT":    RunbookImportDependencyFinderRequiredValueRenderingComponentInputText,
}

var mappingRunbookImportDependencyFinderRequiredValueRenderingComponentEnumLowerCase = map[string]RunbookImportDependencyFinderRequiredValueRenderingComponentEnum{
	"single_choice": RunbookImportDependencyFinderRequiredValueRenderingComponentSingleChoice,
	"multi_choice":  RunbookImportDependencyFinderRequiredValueRenderingComponentMultiChoice,
	"input_text":    RunbookImportDependencyFinderRequiredValueRenderingComponentInputText,
}

// GetRunbookImportDependencyFinderRequiredValueRenderingComponentEnumValues Enumerates the set of values for RunbookImportDependencyFinderRequiredValueRenderingComponentEnum
func GetRunbookImportDependencyFinderRequiredValueRenderingComponentEnumValues() []RunbookImportDependencyFinderRequiredValueRenderingComponentEnum {
	values := make([]RunbookImportDependencyFinderRequiredValueRenderingComponentEnum, 0)
	for _, v := range mappingRunbookImportDependencyFinderRequiredValueRenderingComponentEnum {
		values = append(values, v)
	}
	return values
}

// GetRunbookImportDependencyFinderRequiredValueRenderingComponentEnumStringValues Enumerates the set of values in String for RunbookImportDependencyFinderRequiredValueRenderingComponentEnum
func GetRunbookImportDependencyFinderRequiredValueRenderingComponentEnumStringValues() []string {
	return []string{
		"SINGLE_CHOICE",
		"MULTI_CHOICE",
		"INPUT_TEXT",
	}
}

// GetMappingRunbookImportDependencyFinderRequiredValueRenderingComponentEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRunbookImportDependencyFinderRequiredValueRenderingComponentEnum(val string) (RunbookImportDependencyFinderRequiredValueRenderingComponentEnum, bool) {
	enum, ok := mappingRunbookImportDependencyFinderRequiredValueRenderingComponentEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// RunbookImportDependencyFinderRequiredValueExpectedAssignmentEnum Enum with underlying type: string
type RunbookImportDependencyFinderRequiredValueExpectedAssignmentEnum string

// Set of constants representing the allowable values for RunbookImportDependencyFinderRequiredValueExpectedAssignmentEnum
const (
	RunbookImportDependencyFinderRequiredValueExpectedAssignmentTfPackage    RunbookImportDependencyFinderRequiredValueExpectedAssignmentEnum = "TF_PACKAGE"
	RunbookImportDependencyFinderRequiredValueExpectedAssignmentNonTfPackage RunbookImportDependencyFinderRequiredValueExpectedAssignmentEnum = "NON_TF_PACKAGE"
	RunbookImportDependencyFinderRequiredValueExpectedAssignmentConfigFile   RunbookImportDependencyFinderRequiredValueExpectedAssignmentEnum = "CONFIG_FILE"
	RunbookImportDependencyFinderRequiredValueExpectedAssignmentNull         RunbookImportDependencyFinderRequiredValueExpectedAssignmentEnum = "NULL"
)

var mappingRunbookImportDependencyFinderRequiredValueExpectedAssignmentEnum = map[string]RunbookImportDependencyFinderRequiredValueExpectedAssignmentEnum{
	"TF_PACKAGE":     RunbookImportDependencyFinderRequiredValueExpectedAssignmentTfPackage,
	"NON_TF_PACKAGE": RunbookImportDependencyFinderRequiredValueExpectedAssignmentNonTfPackage,
	"CONFIG_FILE":    RunbookImportDependencyFinderRequiredValueExpectedAssignmentConfigFile,
	"NULL":           RunbookImportDependencyFinderRequiredValueExpectedAssignmentNull,
}

var mappingRunbookImportDependencyFinderRequiredValueExpectedAssignmentEnumLowerCase = map[string]RunbookImportDependencyFinderRequiredValueExpectedAssignmentEnum{
	"tf_package":     RunbookImportDependencyFinderRequiredValueExpectedAssignmentTfPackage,
	"non_tf_package": RunbookImportDependencyFinderRequiredValueExpectedAssignmentNonTfPackage,
	"config_file":    RunbookImportDependencyFinderRequiredValueExpectedAssignmentConfigFile,
	"null":           RunbookImportDependencyFinderRequiredValueExpectedAssignmentNull,
}

// GetRunbookImportDependencyFinderRequiredValueExpectedAssignmentEnumValues Enumerates the set of values for RunbookImportDependencyFinderRequiredValueExpectedAssignmentEnum
func GetRunbookImportDependencyFinderRequiredValueExpectedAssignmentEnumValues() []RunbookImportDependencyFinderRequiredValueExpectedAssignmentEnum {
	values := make([]RunbookImportDependencyFinderRequiredValueExpectedAssignmentEnum, 0)
	for _, v := range mappingRunbookImportDependencyFinderRequiredValueExpectedAssignmentEnum {
		values = append(values, v)
	}
	return values
}

// GetRunbookImportDependencyFinderRequiredValueExpectedAssignmentEnumStringValues Enumerates the set of values in String for RunbookImportDependencyFinderRequiredValueExpectedAssignmentEnum
func GetRunbookImportDependencyFinderRequiredValueExpectedAssignmentEnumStringValues() []string {
	return []string{
		"TF_PACKAGE",
		"NON_TF_PACKAGE",
		"CONFIG_FILE",
		"NULL",
	}
}

// GetMappingRunbookImportDependencyFinderRequiredValueExpectedAssignmentEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRunbookImportDependencyFinderRequiredValueExpectedAssignmentEnum(val string) (RunbookImportDependencyFinderRequiredValueExpectedAssignmentEnum, bool) {
	enum, ok := mappingRunbookImportDependencyFinderRequiredValueExpectedAssignmentEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
