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

// MaskingErrorSummary Summary of a masking error. A Masking error is an error seen during the masking run.
type MaskingErrorSummary struct {

	// The stepName of the masking error.
	StepName MaskingErrorSummaryStepNameEnum `mandatory:"true" json:"stepName"`

	// The text of the masking error.
	Error *string `mandatory:"true" json:"error"`

	// The date and time the error entry was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The statement resulting into the error.
	FailedStatement *string `mandatory:"false" json:"failedStatement"`
}

func (m MaskingErrorSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MaskingErrorSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMaskingErrorSummaryStepNameEnum(string(m.StepName)); !ok && m.StepName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for StepName: %s. Supported values are: %s.", m.StepName, strings.Join(GetMaskingErrorSummaryStepNameEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MaskingErrorSummaryStepNameEnum Enum with underlying type: string
type MaskingErrorSummaryStepNameEnum string

// Set of constants representing the allowable values for MaskingErrorSummaryStepNameEnum
const (
	MaskingErrorSummaryStepNameValidate       MaskingErrorSummaryStepNameEnum = "VALIDATE"
	MaskingErrorSummaryStepNameGenerateScript MaskingErrorSummaryStepNameEnum = "GENERATE_SCRIPT"
	MaskingErrorSummaryStepNameExecuteMasking MaskingErrorSummaryStepNameEnum = "EXECUTE_MASKING"
	MaskingErrorSummaryStepNamePreMasking     MaskingErrorSummaryStepNameEnum = "PRE_MASKING"
	MaskingErrorSummaryStepNamePostMasking    MaskingErrorSummaryStepNameEnum = "POST_MASKING"
)

var mappingMaskingErrorSummaryStepNameEnum = map[string]MaskingErrorSummaryStepNameEnum{
	"VALIDATE":        MaskingErrorSummaryStepNameValidate,
	"GENERATE_SCRIPT": MaskingErrorSummaryStepNameGenerateScript,
	"EXECUTE_MASKING": MaskingErrorSummaryStepNameExecuteMasking,
	"PRE_MASKING":     MaskingErrorSummaryStepNamePreMasking,
	"POST_MASKING":    MaskingErrorSummaryStepNamePostMasking,
}

var mappingMaskingErrorSummaryStepNameEnumLowerCase = map[string]MaskingErrorSummaryStepNameEnum{
	"validate":        MaskingErrorSummaryStepNameValidate,
	"generate_script": MaskingErrorSummaryStepNameGenerateScript,
	"execute_masking": MaskingErrorSummaryStepNameExecuteMasking,
	"pre_masking":     MaskingErrorSummaryStepNamePreMasking,
	"post_masking":    MaskingErrorSummaryStepNamePostMasking,
}

// GetMaskingErrorSummaryStepNameEnumValues Enumerates the set of values for MaskingErrorSummaryStepNameEnum
func GetMaskingErrorSummaryStepNameEnumValues() []MaskingErrorSummaryStepNameEnum {
	values := make([]MaskingErrorSummaryStepNameEnum, 0)
	for _, v := range mappingMaskingErrorSummaryStepNameEnum {
		values = append(values, v)
	}
	return values
}

// GetMaskingErrorSummaryStepNameEnumStringValues Enumerates the set of values in String for MaskingErrorSummaryStepNameEnum
func GetMaskingErrorSummaryStepNameEnumStringValues() []string {
	return []string{
		"VALIDATE",
		"GENERATE_SCRIPT",
		"EXECUTE_MASKING",
		"PRE_MASKING",
		"POST_MASKING",
	}
}

// GetMappingMaskingErrorSummaryStepNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMaskingErrorSummaryStepNameEnum(val string) (MaskingErrorSummaryStepNameEnum, bool) {
	enum, ok := mappingMaskingErrorSummaryStepNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
