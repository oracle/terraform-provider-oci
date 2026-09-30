// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// SubsettingErrorSummary Summary of a subsetting error. A Subsetting error is an error seen during the subsetting run.
type SubsettingErrorSummary struct {

	// The stepName of the subsetting error.
	StepName SubsettingErrorSummaryStepNameEnum `mandatory:"true" json:"stepName"`

	// The text of the subsetting error.
	Error *string `mandatory:"true" json:"error"`

	// The date and time the error entry was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The statement resulting into the error.
	FailedStatement *string `mandatory:"false" json:"failedStatement"`
}

func (m SubsettingErrorSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SubsettingErrorSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSubsettingErrorSummaryStepNameEnum(string(m.StepName)); !ok && m.StepName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for StepName: %s. Supported values are: %s.", m.StepName, strings.Join(GetSubsettingErrorSummaryStepNameEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SubsettingErrorSummaryStepNameEnum Enum with underlying type: string
type SubsettingErrorSummaryStepNameEnum string

// Set of constants representing the allowable values for SubsettingErrorSummaryStepNameEnum
const (
	SubsettingErrorSummaryStepNameValidate              SubsettingErrorSummaryStepNameEnum = "VALIDATE"
	SubsettingErrorSummaryStepNameIdentifyRowsForSubset SubsettingErrorSummaryStepNameEnum = "IDENTIFY_ROWS_FOR_SUBSET"
	SubsettingErrorSummaryStepNameGenerateScript        SubsettingErrorSummaryStepNameEnum = "GENERATE_SCRIPT"
	SubsettingErrorSummaryStepNameExecuteSubsetting     SubsettingErrorSummaryStepNameEnum = "EXECUTE_SUBSETTING"
	SubsettingErrorSummaryStepNamePreSubsetting         SubsettingErrorSummaryStepNameEnum = "PRE_SUBSETTING"
	SubsettingErrorSummaryStepNamePostSubsetting        SubsettingErrorSummaryStepNameEnum = "POST_SUBSETTING"
)

var mappingSubsettingErrorSummaryStepNameEnum = map[string]SubsettingErrorSummaryStepNameEnum{
	"VALIDATE":                 SubsettingErrorSummaryStepNameValidate,
	"IDENTIFY_ROWS_FOR_SUBSET": SubsettingErrorSummaryStepNameIdentifyRowsForSubset,
	"GENERATE_SCRIPT":          SubsettingErrorSummaryStepNameGenerateScript,
	"EXECUTE_SUBSETTING":       SubsettingErrorSummaryStepNameExecuteSubsetting,
	"PRE_SUBSETTING":           SubsettingErrorSummaryStepNamePreSubsetting,
	"POST_SUBSETTING":          SubsettingErrorSummaryStepNamePostSubsetting,
}

var mappingSubsettingErrorSummaryStepNameEnumLowerCase = map[string]SubsettingErrorSummaryStepNameEnum{
	"validate":                 SubsettingErrorSummaryStepNameValidate,
	"identify_rows_for_subset": SubsettingErrorSummaryStepNameIdentifyRowsForSubset,
	"generate_script":          SubsettingErrorSummaryStepNameGenerateScript,
	"execute_subsetting":       SubsettingErrorSummaryStepNameExecuteSubsetting,
	"pre_subsetting":           SubsettingErrorSummaryStepNamePreSubsetting,
	"post_subsetting":          SubsettingErrorSummaryStepNamePostSubsetting,
}

// GetSubsettingErrorSummaryStepNameEnumValues Enumerates the set of values for SubsettingErrorSummaryStepNameEnum
func GetSubsettingErrorSummaryStepNameEnumValues() []SubsettingErrorSummaryStepNameEnum {
	values := make([]SubsettingErrorSummaryStepNameEnum, 0)
	for _, v := range mappingSubsettingErrorSummaryStepNameEnum {
		values = append(values, v)
	}
	return values
}

// GetSubsettingErrorSummaryStepNameEnumStringValues Enumerates the set of values in String for SubsettingErrorSummaryStepNameEnum
func GetSubsettingErrorSummaryStepNameEnumStringValues() []string {
	return []string{
		"VALIDATE",
		"IDENTIFY_ROWS_FOR_SUBSET",
		"GENERATE_SCRIPT",
		"EXECUTE_SUBSETTING",
		"PRE_SUBSETTING",
		"POST_SUBSETTING",
	}
}

// GetMappingSubsettingErrorSummaryStepNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSubsettingErrorSummaryStepNameEnum(val string) (SubsettingErrorSummaryStepNameEnum, bool) {
	enum, ok := mappingSubsettingErrorSummaryStepNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
