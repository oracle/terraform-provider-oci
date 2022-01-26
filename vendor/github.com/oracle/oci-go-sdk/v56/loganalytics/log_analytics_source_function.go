// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// LogAnalyticsSourceFunction LogAnalyticsSourceFunction
type LogAnalyticsSourceFunction struct {

	// The function argument.
	Arguments []LogAnalyticsMetaFunctionArgument `mandatory:"false" json:"arguments"`

	// A flag inidcating whether or not the source function is enabled.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	Function *LogAnalyticsMetaFunction `mandatory:"false" json:"function"`

	// The source function name
	FunctionName LogAnalyticsSourceFunctionFunctionNameEnum `mandatory:"false" json:"functionName,omitempty"`

	// The source function unique identifier as a string.
	FunctionReference *string `mandatory:"false" json:"functionReference"`

	// The source unique identifier as a string.
	SourceReference *string `mandatory:"false" json:"sourceReference"`

	// Features of the source function to use for enrichment.
	Features []string `mandatory:"false" json:"features"`

	// The source function unique identifier.
	FunctionId *int64 `mandatory:"false" json:"functionId"`

	// The source function order.
	Order *int64 `mandatory:"false" json:"order"`

	// The system flag.  A value of false denotes a custom, or user
	// defined object.  A value of true denotes a built in object.
	IsSystem *bool `mandatory:"false" json:"isSystem"`

	// The lookup column.
	LookupColumn *string `mandatory:"false" json:"lookupColumn"`

	// The lookup column position.
	LookupColumnPosition *int64 `mandatory:"false" json:"lookupColumnPosition"`

	// The lookup display name.
	LookupDisplayName *string `mandatory:"false" json:"lookupDisplayName"`

	// The lookup  mode.
	LookupMode *int64 `mandatory:"false" json:"lookupMode"`

	// The lookup table.
	LookupTable *string `mandatory:"false" json:"lookupTable"`

	// The source unique identifier.
	SourceId *int64 `mandatory:"false" json:"sourceId"`
}

func (m LogAnalyticsSourceFunction) String() string {
	return common.PointerString(m)
}

// LogAnalyticsSourceFunctionFunctionNameEnum Enum with underlying type: string
type LogAnalyticsSourceFunctionFunctionNameEnum string

// Set of constants representing the allowable values for LogAnalyticsSourceFunctionFunctionNameEnum
const (
	LogAnalyticsSourceFunctionFunctionNameGeolocation LogAnalyticsSourceFunctionFunctionNameEnum = "GEOLOCATION"
	LogAnalyticsSourceFunctionFunctionNameLookup      LogAnalyticsSourceFunctionFunctionNameEnum = "LOOKUP"
)

var mappingLogAnalyticsSourceFunctionFunctionName = map[string]LogAnalyticsSourceFunctionFunctionNameEnum{
	"GEOLOCATION": LogAnalyticsSourceFunctionFunctionNameGeolocation,
	"LOOKUP":      LogAnalyticsSourceFunctionFunctionNameLookup,
}

// GetLogAnalyticsSourceFunctionFunctionNameEnumValues Enumerates the set of values for LogAnalyticsSourceFunctionFunctionNameEnum
func GetLogAnalyticsSourceFunctionFunctionNameEnumValues() []LogAnalyticsSourceFunctionFunctionNameEnum {
	values := make([]LogAnalyticsSourceFunctionFunctionNameEnum, 0)
	for _, v := range mappingLogAnalyticsSourceFunctionFunctionName {
		values = append(values, v)
	}
	return values
}
