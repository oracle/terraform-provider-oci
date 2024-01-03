// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AlertLogCountSummary The details for one alert log count entry.
type AlertLogCountSummary struct {

	// The category of different alert logs.
	Category AlertLogCountSummaryCategoryEnum `mandatory:"true" json:"category"`

	// The count of alert logs with specific category.
	Count *int `mandatory:"true" json:"count"`
}

func (m AlertLogCountSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AlertLogCountSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAlertLogCountSummaryCategoryEnum(string(m.Category)); !ok && m.Category != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Category: %s. Supported values are: %s.", m.Category, strings.Join(GetAlertLogCountSummaryCategoryEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AlertLogCountSummaryCategoryEnum Enum with underlying type: string
type AlertLogCountSummaryCategoryEnum string

// Set of constants representing the allowable values for AlertLogCountSummaryCategoryEnum
const (
	AlertLogCountSummaryCategoryUnknown       AlertLogCountSummaryCategoryEnum = "UNKNOWN"
	AlertLogCountSummaryCategoryIncidentError AlertLogCountSummaryCategoryEnum = "INCIDENT_ERROR"
	AlertLogCountSummaryCategoryError         AlertLogCountSummaryCategoryEnum = "ERROR"
	AlertLogCountSummaryCategoryWarning       AlertLogCountSummaryCategoryEnum = "WARNING"
	AlertLogCountSummaryCategoryNotification  AlertLogCountSummaryCategoryEnum = "NOTIFICATION"
	AlertLogCountSummaryCategoryTrace         AlertLogCountSummaryCategoryEnum = "TRACE"
	AlertLogCountSummaryCategoryCritical      AlertLogCountSummaryCategoryEnum = "CRITICAL"
	AlertLogCountSummaryCategorySevere        AlertLogCountSummaryCategoryEnum = "SEVERE"
	AlertLogCountSummaryCategoryImportant     AlertLogCountSummaryCategoryEnum = "IMPORTANT"
	AlertLogCountSummaryCategoryNormal        AlertLogCountSummaryCategoryEnum = "NORMAL"
	AlertLogCountSummaryCategoryOther         AlertLogCountSummaryCategoryEnum = "OTHER"
)

var mappingAlertLogCountSummaryCategoryEnum = map[string]AlertLogCountSummaryCategoryEnum{
	"UNKNOWN":        AlertLogCountSummaryCategoryUnknown,
	"INCIDENT_ERROR": AlertLogCountSummaryCategoryIncidentError,
	"ERROR":          AlertLogCountSummaryCategoryError,
	"WARNING":        AlertLogCountSummaryCategoryWarning,
	"NOTIFICATION":   AlertLogCountSummaryCategoryNotification,
	"TRACE":          AlertLogCountSummaryCategoryTrace,
	"CRITICAL":       AlertLogCountSummaryCategoryCritical,
	"SEVERE":         AlertLogCountSummaryCategorySevere,
	"IMPORTANT":      AlertLogCountSummaryCategoryImportant,
	"NORMAL":         AlertLogCountSummaryCategoryNormal,
	"OTHER":          AlertLogCountSummaryCategoryOther,
}

var mappingAlertLogCountSummaryCategoryEnumLowerCase = map[string]AlertLogCountSummaryCategoryEnum{
	"unknown":        AlertLogCountSummaryCategoryUnknown,
	"incident_error": AlertLogCountSummaryCategoryIncidentError,
	"error":          AlertLogCountSummaryCategoryError,
	"warning":        AlertLogCountSummaryCategoryWarning,
	"notification":   AlertLogCountSummaryCategoryNotification,
	"trace":          AlertLogCountSummaryCategoryTrace,
	"critical":       AlertLogCountSummaryCategoryCritical,
	"severe":         AlertLogCountSummaryCategorySevere,
	"important":      AlertLogCountSummaryCategoryImportant,
	"normal":         AlertLogCountSummaryCategoryNormal,
	"other":          AlertLogCountSummaryCategoryOther,
}

// GetAlertLogCountSummaryCategoryEnumValues Enumerates the set of values for AlertLogCountSummaryCategoryEnum
func GetAlertLogCountSummaryCategoryEnumValues() []AlertLogCountSummaryCategoryEnum {
	values := make([]AlertLogCountSummaryCategoryEnum, 0)
	for _, v := range mappingAlertLogCountSummaryCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetAlertLogCountSummaryCategoryEnumStringValues Enumerates the set of values in String for AlertLogCountSummaryCategoryEnum
func GetAlertLogCountSummaryCategoryEnumStringValues() []string {
	return []string{
		"UNKNOWN",
		"INCIDENT_ERROR",
		"ERROR",
		"WARNING",
		"NOTIFICATION",
		"TRACE",
		"CRITICAL",
		"SEVERE",
		"IMPORTANT",
		"NORMAL",
		"OTHER",
	}
}

// GetMappingAlertLogCountSummaryCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAlertLogCountSummaryCategoryEnum(val string) (AlertLogCountSummaryCategoryEnum, bool) {
	enum, ok := mappingAlertLogCountSummaryCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
