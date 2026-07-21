// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ValidateDataFilterDetails Details of the data filter to be validated.
type ValidateDataFilterDetails struct {

	// The filter type.
	FilterType ValidateDataFilterDetailsFilterTypeEnum `mandatory:"false" json:"filterType,omitempty"`

	// The example content on which the data filter will be applied.
	ExampleContent *string `mandatory:"false" json:"exampleContent"`

	// The regular expression the data filter will use for matching.
	MatchRegularExpression *string `mandatory:"false" json:"matchRegularExpression"`

	// The replacement string.
	ReplacementString *string `mandatory:"false" json:"replacementString"`

	// The hash type (0 for numeric, 1 for text). Only applicable for HASH_MASK filter type
	HashType *int `mandatory:"false" json:"hashType"`
}

func (m ValidateDataFilterDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ValidateDataFilterDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingValidateDataFilterDetailsFilterTypeEnum(string(m.FilterType)); !ok && m.FilterType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FilterType: %s. Supported values are: %s.", m.FilterType, strings.Join(GetValidateDataFilterDetailsFilterTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ValidateDataFilterDetailsFilterTypeEnum Enum with underlying type: string
type ValidateDataFilterDetailsFilterTypeEnum string

// Set of constants representing the allowable values for ValidateDataFilterDetailsFilterTypeEnum
const (
	ValidateDataFilterDetailsFilterTypeHashMask     ValidateDataFilterDetailsFilterTypeEnum = "HASH_MASK"
	ValidateDataFilterDetailsFilterTypeMask         ValidateDataFilterDetailsFilterTypeEnum = "MASK"
	ValidateDataFilterDetailsFilterTypeDropLogEntry ValidateDataFilterDetailsFilterTypeEnum = "DROP_LOG_ENTRY"
	ValidateDataFilterDetailsFilterTypeDropString   ValidateDataFilterDetailsFilterTypeEnum = "DROP_STRING"
)

var mappingValidateDataFilterDetailsFilterTypeEnum = map[string]ValidateDataFilterDetailsFilterTypeEnum{
	"HASH_MASK":      ValidateDataFilterDetailsFilterTypeHashMask,
	"MASK":           ValidateDataFilterDetailsFilterTypeMask,
	"DROP_LOG_ENTRY": ValidateDataFilterDetailsFilterTypeDropLogEntry,
	"DROP_STRING":    ValidateDataFilterDetailsFilterTypeDropString,
}

var mappingValidateDataFilterDetailsFilterTypeEnumLowerCase = map[string]ValidateDataFilterDetailsFilterTypeEnum{
	"hash_mask":      ValidateDataFilterDetailsFilterTypeHashMask,
	"mask":           ValidateDataFilterDetailsFilterTypeMask,
	"drop_log_entry": ValidateDataFilterDetailsFilterTypeDropLogEntry,
	"drop_string":    ValidateDataFilterDetailsFilterTypeDropString,
}

// GetValidateDataFilterDetailsFilterTypeEnumValues Enumerates the set of values for ValidateDataFilterDetailsFilterTypeEnum
func GetValidateDataFilterDetailsFilterTypeEnumValues() []ValidateDataFilterDetailsFilterTypeEnum {
	values := make([]ValidateDataFilterDetailsFilterTypeEnum, 0)
	for _, v := range mappingValidateDataFilterDetailsFilterTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetValidateDataFilterDetailsFilterTypeEnumStringValues Enumerates the set of values in String for ValidateDataFilterDetailsFilterTypeEnum
func GetValidateDataFilterDetailsFilterTypeEnumStringValues() []string {
	return []string{
		"HASH_MASK",
		"MASK",
		"DROP_LOG_ENTRY",
		"DROP_STRING",
	}
}

// GetMappingValidateDataFilterDetailsFilterTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingValidateDataFilterDetailsFilterTypeEnum(val string) (ValidateDataFilterDetailsFilterTypeEnum, bool) {
	enum, ok := mappingValidateDataFilterDetailsFilterTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
