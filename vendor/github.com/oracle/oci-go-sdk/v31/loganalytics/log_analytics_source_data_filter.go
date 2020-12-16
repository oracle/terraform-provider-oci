// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v31/common"
)

// LogAnalyticsSourceDataFilter LogAnalyticsSourceDataFilter
type LogAnalyticsSourceDataFilter struct {

	// description
	Description *string `mandatory:"false" json:"description"`

	// display name
	DisplayName *string `mandatory:"false" json:"displayName"`

	// edit version
	EditVersion *int64 `mandatory:"false" json:"editVersion"`

	// enabled
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// field internal name
	FieldName *string `mandatory:"false" json:"fieldName"`

	// hash type
	HashType *int `mandatory:"false" json:"hashType"`

	// filter Id
	DataFilterId *int64 `mandatory:"false" json:"dataFilterId"`

	// is system flag
	IsSystem *bool `mandatory:"false" json:"isSystem"`

	// regular expression match
	MatchRegularExpression *string `mandatory:"false" json:"matchRegularExpression"`

	// order
	Order *int64 `mandatory:"false" json:"order"`

	// path
	Path *string `mandatory:"false" json:"path"`

	// replacement string
	ReplacementString *string `mandatory:"false" json:"replacementString"`

	// source Id
	SourceId *int64 `mandatory:"false" json:"sourceId"`

	// filterType
	FilterType LogAnalyticsSourceDataFilterFilterTypeEnum `mandatory:"false" json:"filterType,omitempty"`
}

func (m LogAnalyticsSourceDataFilter) String() string {
	return common.PointerString(m)
}

// LogAnalyticsSourceDataFilterFilterTypeEnum Enum with underlying type: string
type LogAnalyticsSourceDataFilterFilterTypeEnum string

// Set of constants representing the allowable values for LogAnalyticsSourceDataFilterFilterTypeEnum
const (
	LogAnalyticsSourceDataFilterFilterTypeMask         LogAnalyticsSourceDataFilterFilterTypeEnum = "MASK"
	LogAnalyticsSourceDataFilterFilterTypeHashMask     LogAnalyticsSourceDataFilterFilterTypeEnum = "HASH_MASK"
	LogAnalyticsSourceDataFilterFilterTypeDropLogEntry LogAnalyticsSourceDataFilterFilterTypeEnum = "DROP_LOG_ENTRY"
	LogAnalyticsSourceDataFilterFilterTypeDropString   LogAnalyticsSourceDataFilterFilterTypeEnum = "DROP_STRING"
)

var mappingLogAnalyticsSourceDataFilterFilterType = map[string]LogAnalyticsSourceDataFilterFilterTypeEnum{
	"MASK":           LogAnalyticsSourceDataFilterFilterTypeMask,
	"HASH_MASK":      LogAnalyticsSourceDataFilterFilterTypeHashMask,
	"DROP_LOG_ENTRY": LogAnalyticsSourceDataFilterFilterTypeDropLogEntry,
	"DROP_STRING":    LogAnalyticsSourceDataFilterFilterTypeDropString,
}

// GetLogAnalyticsSourceDataFilterFilterTypeEnumValues Enumerates the set of values for LogAnalyticsSourceDataFilterFilterTypeEnum
func GetLogAnalyticsSourceDataFilterFilterTypeEnumValues() []LogAnalyticsSourceDataFilterFilterTypeEnum {
	values := make([]LogAnalyticsSourceDataFilterFilterTypeEnum, 0)
	for _, v := range mappingLogAnalyticsSourceDataFilterFilterType {
		values = append(values, v)
	}
	return values
}
