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

// QueryUsageDetails Input to get storage usage grouped by data type
type QueryUsageDetails struct {

	// The type of data whose usage is requested. Use LOG, APM, or ALL for both.
	DataType QueryUsageDetailsDataTypeEnum `mandatory:"true" json:"dataType"`

	// Optional array of logIndex numbers to filter by.
	// 0 for _default LogIndex (logIndex holding data not associated with any logSet), -2 for _unmapped LogIndex
	// (logIndex holding data having logSet defined but without logIndex mapped when the "mappingMechanism" is
	// "specific").
	LogIndexes []int `mandatory:"false" json:"logIndexes"`
}

func (m QueryUsageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m QueryUsageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingQueryUsageDetailsDataTypeEnum(string(m.DataType)); !ok && m.DataType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataType: %s. Supported values are: %s.", m.DataType, strings.Join(GetQueryUsageDetailsDataTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// QueryUsageDetailsDataTypeEnum Enum with underlying type: string
type QueryUsageDetailsDataTypeEnum string

// Set of constants representing the allowable values for QueryUsageDetailsDataTypeEnum
const (
	QueryUsageDetailsDataTypeLog QueryUsageDetailsDataTypeEnum = "LOG"
	QueryUsageDetailsDataTypeApm QueryUsageDetailsDataTypeEnum = "APM"
	QueryUsageDetailsDataTypeAll QueryUsageDetailsDataTypeEnum = "ALL"
)

var mappingQueryUsageDetailsDataTypeEnum = map[string]QueryUsageDetailsDataTypeEnum{
	"LOG": QueryUsageDetailsDataTypeLog,
	"APM": QueryUsageDetailsDataTypeApm,
	"ALL": QueryUsageDetailsDataTypeAll,
}

var mappingQueryUsageDetailsDataTypeEnumLowerCase = map[string]QueryUsageDetailsDataTypeEnum{
	"log": QueryUsageDetailsDataTypeLog,
	"apm": QueryUsageDetailsDataTypeApm,
	"all": QueryUsageDetailsDataTypeAll,
}

// GetQueryUsageDetailsDataTypeEnumValues Enumerates the set of values for QueryUsageDetailsDataTypeEnum
func GetQueryUsageDetailsDataTypeEnumValues() []QueryUsageDetailsDataTypeEnum {
	values := make([]QueryUsageDetailsDataTypeEnum, 0)
	for _, v := range mappingQueryUsageDetailsDataTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetQueryUsageDetailsDataTypeEnumStringValues Enumerates the set of values in String for QueryUsageDetailsDataTypeEnum
func GetQueryUsageDetailsDataTypeEnumStringValues() []string {
	return []string{
		"LOG",
		"APM",
		"ALL",
	}
}

// GetMappingQueryUsageDetailsDataTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingQueryUsageDetailsDataTypeEnum(val string) (QueryUsageDetailsDataTypeEnum, bool) {
	enum, ok := mappingQueryUsageDetailsDataTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
