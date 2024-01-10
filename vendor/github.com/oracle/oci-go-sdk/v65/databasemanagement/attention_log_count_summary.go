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

// AttentionLogCountSummary The details for one attention log count entry.
type AttentionLogCountSummary struct {

	// The category of different attention logs.
	Category AttentionLogCountSummaryCategoryEnum `mandatory:"true" json:"category"`

	// The count of attention logs with specific category.
	Count *int `mandatory:"true" json:"count"`
}

func (m AttentionLogCountSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AttentionLogCountSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAttentionLogCountSummaryCategoryEnum(string(m.Category)); !ok && m.Category != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Category: %s. Supported values are: %s.", m.Category, strings.Join(GetAttentionLogCountSummaryCategoryEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AttentionLogCountSummaryCategoryEnum Enum with underlying type: string
type AttentionLogCountSummaryCategoryEnum string

// Set of constants representing the allowable values for AttentionLogCountSummaryCategoryEnum
const (
	AttentionLogCountSummaryCategoryUnknown       AttentionLogCountSummaryCategoryEnum = "UNKNOWN"
	AttentionLogCountSummaryCategoryIncidentError AttentionLogCountSummaryCategoryEnum = "INCIDENT_ERROR"
	AttentionLogCountSummaryCategoryError         AttentionLogCountSummaryCategoryEnum = "ERROR"
	AttentionLogCountSummaryCategoryWarning       AttentionLogCountSummaryCategoryEnum = "WARNING"
	AttentionLogCountSummaryCategoryNotification  AttentionLogCountSummaryCategoryEnum = "NOTIFICATION"
	AttentionLogCountSummaryCategoryTrace         AttentionLogCountSummaryCategoryEnum = "TRACE"
	AttentionLogCountSummaryCategoryImmediate     AttentionLogCountSummaryCategoryEnum = "IMMEDIATE"
	AttentionLogCountSummaryCategorySoon          AttentionLogCountSummaryCategoryEnum = "SOON"
	AttentionLogCountSummaryCategoryDeferrable    AttentionLogCountSummaryCategoryEnum = "DEFERRABLE"
	AttentionLogCountSummaryCategoryInfo          AttentionLogCountSummaryCategoryEnum = "INFO"
	AttentionLogCountSummaryCategoryOther         AttentionLogCountSummaryCategoryEnum = "OTHER"
)

var mappingAttentionLogCountSummaryCategoryEnum = map[string]AttentionLogCountSummaryCategoryEnum{
	"UNKNOWN":        AttentionLogCountSummaryCategoryUnknown,
	"INCIDENT_ERROR": AttentionLogCountSummaryCategoryIncidentError,
	"ERROR":          AttentionLogCountSummaryCategoryError,
	"WARNING":        AttentionLogCountSummaryCategoryWarning,
	"NOTIFICATION":   AttentionLogCountSummaryCategoryNotification,
	"TRACE":          AttentionLogCountSummaryCategoryTrace,
	"IMMEDIATE":      AttentionLogCountSummaryCategoryImmediate,
	"SOON":           AttentionLogCountSummaryCategorySoon,
	"DEFERRABLE":     AttentionLogCountSummaryCategoryDeferrable,
	"INFO":           AttentionLogCountSummaryCategoryInfo,
	"OTHER":          AttentionLogCountSummaryCategoryOther,
}

var mappingAttentionLogCountSummaryCategoryEnumLowerCase = map[string]AttentionLogCountSummaryCategoryEnum{
	"unknown":        AttentionLogCountSummaryCategoryUnknown,
	"incident_error": AttentionLogCountSummaryCategoryIncidentError,
	"error":          AttentionLogCountSummaryCategoryError,
	"warning":        AttentionLogCountSummaryCategoryWarning,
	"notification":   AttentionLogCountSummaryCategoryNotification,
	"trace":          AttentionLogCountSummaryCategoryTrace,
	"immediate":      AttentionLogCountSummaryCategoryImmediate,
	"soon":           AttentionLogCountSummaryCategorySoon,
	"deferrable":     AttentionLogCountSummaryCategoryDeferrable,
	"info":           AttentionLogCountSummaryCategoryInfo,
	"other":          AttentionLogCountSummaryCategoryOther,
}

// GetAttentionLogCountSummaryCategoryEnumValues Enumerates the set of values for AttentionLogCountSummaryCategoryEnum
func GetAttentionLogCountSummaryCategoryEnumValues() []AttentionLogCountSummaryCategoryEnum {
	values := make([]AttentionLogCountSummaryCategoryEnum, 0)
	for _, v := range mappingAttentionLogCountSummaryCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetAttentionLogCountSummaryCategoryEnumStringValues Enumerates the set of values in String for AttentionLogCountSummaryCategoryEnum
func GetAttentionLogCountSummaryCategoryEnumStringValues() []string {
	return []string{
		"UNKNOWN",
		"INCIDENT_ERROR",
		"ERROR",
		"WARNING",
		"NOTIFICATION",
		"TRACE",
		"IMMEDIATE",
		"SOON",
		"DEFERRABLE",
		"INFO",
		"OTHER",
	}
}

// GetMappingAttentionLogCountSummaryCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAttentionLogCountSummaryCategoryEnum(val string) (AttentionLogCountSummaryCategoryEnum, bool) {
	enum, ok := mappingAttentionLogCountSummaryCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
