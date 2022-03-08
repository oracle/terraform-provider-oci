// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Connectivity Management API
//
// Use the DCMS APIs to perform Metadata/Data operations.
//

package dataconnectivity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v61/common"
	"strings"
)

// ProfileConfig Profiling configuration.
type ProfileConfig struct {

	// Array of column names to profile. If empty all columns in the entity are profiled.
	Attributes []string `mandatory:"false" json:"attributes"`

	// Array of enum Strings basically what all profile functions to run. If empty, all supported functions are run.
	Functions []ProfileConfigFunctionsEnum `mandatory:"false" json:"functions,omitempty"`

	// The maximum number of value frequencies to return per column. The VFs are sorted descending on frequency and ascending on value and then topN are returned and rest discarded.
	TopNValFreq *int `mandatory:"false" json:"topNValFreq"`

	// A pattern has to qualify minumum this percentage threshold to be considered a legitimate pattern on its own. All patterns which does not qualify this will be clubbed together into a single 'Others' pattern.
	PatternThreshold *int `mandatory:"false" json:"patternThreshold"`

	// A data type has to qualify minimum this percentage threshold to be considered an infrred data type for a column.
	DataTypeThreshold *int `mandatory:"false" json:"dataTypeThreshold"`
}

func (m ProfileConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ProfileConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.Functions {
		if _, ok := GetMappingProfileConfigFunctionsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Functions: %s. Supported values are: %s.", val, strings.Join(GetProfileConfigFunctionsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ProfileConfigFunctionsEnum Enum with underlying type: string
type ProfileConfigFunctionsEnum string

// Set of constants representing the allowable values for ProfileConfigFunctionsEnum
const (
	ProfileConfigFunctionsAttributeCount    ProfileConfigFunctionsEnum = "ATTRIBUTE_COUNT"
	ProfileConfigFunctionsRowCount          ProfileConfigFunctionsEnum = "ROW_COUNT"
	ProfileConfigFunctionsDataType          ProfileConfigFunctionsEnum = "DATA_TYPE"
	ProfileConfigFunctionsDistinctCount     ProfileConfigFunctionsEnum = "DISTINCT_COUNT"
	ProfileConfigFunctionsDuplicateCount    ProfileConfigFunctionsEnum = "DUPLICATE_COUNT"
	ProfileConfigFunctionsHistogram         ProfileConfigFunctionsEnum = "HISTOGRAM"
	ProfileConfigFunctionsMax               ProfileConfigFunctionsEnum = "MAX"
	ProfileConfigFunctionsMaxLength         ProfileConfigFunctionsEnum = "MAX_LENGTH"
	ProfileConfigFunctionsMean              ProfileConfigFunctionsEnum = "MEAN"
	ProfileConfigFunctionsMeanLength        ProfileConfigFunctionsEnum = "MEAN_LENGTH"
	ProfileConfigFunctionsMedian            ProfileConfigFunctionsEnum = "MEDIAN"
	ProfileConfigFunctionsMin               ProfileConfigFunctionsEnum = "MIN"
	ProfileConfigFunctionsMinLength         ProfileConfigFunctionsEnum = "MIN_LENGTH"
	ProfileConfigFunctionsNullCount         ProfileConfigFunctionsEnum = "NULL_COUNT"
	ProfileConfigFunctionsOutlier           ProfileConfigFunctionsEnum = "OUTLIER"
	ProfileConfigFunctionsPattern           ProfileConfigFunctionsEnum = "PATTERN"
	ProfileConfigFunctionsStandardDeviation ProfileConfigFunctionsEnum = "STANDARD_DEVIATION"
	ProfileConfigFunctionsUniqueCount       ProfileConfigFunctionsEnum = "UNIQUE_COUNT"
	ProfileConfigFunctionsVariance          ProfileConfigFunctionsEnum = "VARIANCE"
	ProfileConfigFunctionsValueFrequency    ProfileConfigFunctionsEnum = "VALUE_FREQUENCY"
)

var mappingProfileConfigFunctionsEnum = map[string]ProfileConfigFunctionsEnum{
	"ATTRIBUTE_COUNT":    ProfileConfigFunctionsAttributeCount,
	"ROW_COUNT":          ProfileConfigFunctionsRowCount,
	"DATA_TYPE":          ProfileConfigFunctionsDataType,
	"DISTINCT_COUNT":     ProfileConfigFunctionsDistinctCount,
	"DUPLICATE_COUNT":    ProfileConfigFunctionsDuplicateCount,
	"HISTOGRAM":          ProfileConfigFunctionsHistogram,
	"MAX":                ProfileConfigFunctionsMax,
	"MAX_LENGTH":         ProfileConfigFunctionsMaxLength,
	"MEAN":               ProfileConfigFunctionsMean,
	"MEAN_LENGTH":        ProfileConfigFunctionsMeanLength,
	"MEDIAN":             ProfileConfigFunctionsMedian,
	"MIN":                ProfileConfigFunctionsMin,
	"MIN_LENGTH":         ProfileConfigFunctionsMinLength,
	"NULL_COUNT":         ProfileConfigFunctionsNullCount,
	"OUTLIER":            ProfileConfigFunctionsOutlier,
	"PATTERN":            ProfileConfigFunctionsPattern,
	"STANDARD_DEVIATION": ProfileConfigFunctionsStandardDeviation,
	"UNIQUE_COUNT":       ProfileConfigFunctionsUniqueCount,
	"VARIANCE":           ProfileConfigFunctionsVariance,
	"VALUE_FREQUENCY":    ProfileConfigFunctionsValueFrequency,
}

var mappingProfileConfigFunctionsEnumLowerCase = map[string]ProfileConfigFunctionsEnum{
	"attribute_count":    ProfileConfigFunctionsAttributeCount,
	"row_count":          ProfileConfigFunctionsRowCount,
	"data_type":          ProfileConfigFunctionsDataType,
	"distinct_count":     ProfileConfigFunctionsDistinctCount,
	"duplicate_count":    ProfileConfigFunctionsDuplicateCount,
	"histogram":          ProfileConfigFunctionsHistogram,
	"max":                ProfileConfigFunctionsMax,
	"max_length":         ProfileConfigFunctionsMaxLength,
	"mean":               ProfileConfigFunctionsMean,
	"mean_length":        ProfileConfigFunctionsMeanLength,
	"median":             ProfileConfigFunctionsMedian,
	"min":                ProfileConfigFunctionsMin,
	"min_length":         ProfileConfigFunctionsMinLength,
	"null_count":         ProfileConfigFunctionsNullCount,
	"outlier":            ProfileConfigFunctionsOutlier,
	"pattern":            ProfileConfigFunctionsPattern,
	"standard_deviation": ProfileConfigFunctionsStandardDeviation,
	"unique_count":       ProfileConfigFunctionsUniqueCount,
	"variance":           ProfileConfigFunctionsVariance,
	"value_frequency":    ProfileConfigFunctionsValueFrequency,
}

// GetProfileConfigFunctionsEnumValues Enumerates the set of values for ProfileConfigFunctionsEnum
func GetProfileConfigFunctionsEnumValues() []ProfileConfigFunctionsEnum {
	values := make([]ProfileConfigFunctionsEnum, 0)
	for _, v := range mappingProfileConfigFunctionsEnum {
		values = append(values, v)
	}
	return values
}

// GetProfileConfigFunctionsEnumStringValues Enumerates the set of values in String for ProfileConfigFunctionsEnum
func GetProfileConfigFunctionsEnumStringValues() []string {
	return []string{
		"ATTRIBUTE_COUNT",
		"ROW_COUNT",
		"DATA_TYPE",
		"DISTINCT_COUNT",
		"DUPLICATE_COUNT",
		"HISTOGRAM",
		"MAX",
		"MAX_LENGTH",
		"MEAN",
		"MEAN_LENGTH",
		"MEDIAN",
		"MIN",
		"MIN_LENGTH",
		"NULL_COUNT",
		"OUTLIER",
		"PATTERN",
		"STANDARD_DEVIATION",
		"UNIQUE_COUNT",
		"VARIANCE",
		"VALUE_FREQUENCY",
	}
}

// GetMappingProfileConfigFunctionsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingProfileConfigFunctionsEnum(val string) (ProfileConfigFunctionsEnum, bool) {
	enum, ok := mappingProfileConfigFunctionsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
