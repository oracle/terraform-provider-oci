// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AlertLogSummary The detail for one alert log entry.
type AlertLogSummary struct {

	// The level of the alert log.
	MessageLevel AlertLogSummaryMessageLevelEnum `mandatory:"true" json:"messageLevel"`

	// The type of alert log message.
	MessageType AlertLogSummaryMessageTypeEnum `mandatory:"true" json:"messageType"`

	// The contents of the alert log message.
	MessageContent *string `mandatory:"false" json:"messageContent"`

	// The date and time the alert log was created.
	Timestamp *common.SDKTime `mandatory:"false" json:"timestamp"`

	// The supplemental details of the alert log.
	SupplementalDetail *string `mandatory:"false" json:"supplementalDetail"`

	// The alert log file location.
	FileLocation *string `mandatory:"false" json:"fileLocation"`
}

func (m AlertLogSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AlertLogSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAlertLogSummaryMessageLevelEnum(string(m.MessageLevel)); !ok && m.MessageLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MessageLevel: %s. Supported values are: %s.", m.MessageLevel, strings.Join(GetAlertLogSummaryMessageLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAlertLogSummaryMessageTypeEnum(string(m.MessageType)); !ok && m.MessageType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MessageType: %s. Supported values are: %s.", m.MessageType, strings.Join(GetAlertLogSummaryMessageTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AlertLogSummaryMessageLevelEnum Enum with underlying type: string
type AlertLogSummaryMessageLevelEnum string

// Set of constants representing the allowable values for AlertLogSummaryMessageLevelEnum
const (
	AlertLogSummaryMessageLevelCritical  AlertLogSummaryMessageLevelEnum = "CRITICAL"
	AlertLogSummaryMessageLevelSevere    AlertLogSummaryMessageLevelEnum = "SEVERE"
	AlertLogSummaryMessageLevelImportant AlertLogSummaryMessageLevelEnum = "IMPORTANT"
	AlertLogSummaryMessageLevelNormal    AlertLogSummaryMessageLevelEnum = "NORMAL"
)

var mappingAlertLogSummaryMessageLevelEnum = map[string]AlertLogSummaryMessageLevelEnum{
	"CRITICAL":  AlertLogSummaryMessageLevelCritical,
	"SEVERE":    AlertLogSummaryMessageLevelSevere,
	"IMPORTANT": AlertLogSummaryMessageLevelImportant,
	"NORMAL":    AlertLogSummaryMessageLevelNormal,
}

var mappingAlertLogSummaryMessageLevelEnumLowerCase = map[string]AlertLogSummaryMessageLevelEnum{
	"critical":  AlertLogSummaryMessageLevelCritical,
	"severe":    AlertLogSummaryMessageLevelSevere,
	"important": AlertLogSummaryMessageLevelImportant,
	"normal":    AlertLogSummaryMessageLevelNormal,
}

// GetAlertLogSummaryMessageLevelEnumValues Enumerates the set of values for AlertLogSummaryMessageLevelEnum
func GetAlertLogSummaryMessageLevelEnumValues() []AlertLogSummaryMessageLevelEnum {
	values := make([]AlertLogSummaryMessageLevelEnum, 0)
	for _, v := range mappingAlertLogSummaryMessageLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetAlertLogSummaryMessageLevelEnumStringValues Enumerates the set of values in String for AlertLogSummaryMessageLevelEnum
func GetAlertLogSummaryMessageLevelEnumStringValues() []string {
	return []string{
		"CRITICAL",
		"SEVERE",
		"IMPORTANT",
		"NORMAL",
	}
}

// GetMappingAlertLogSummaryMessageLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAlertLogSummaryMessageLevelEnum(val string) (AlertLogSummaryMessageLevelEnum, bool) {
	enum, ok := mappingAlertLogSummaryMessageLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AlertLogSummaryMessageTypeEnum Enum with underlying type: string
type AlertLogSummaryMessageTypeEnum string

// Set of constants representing the allowable values for AlertLogSummaryMessageTypeEnum
const (
	AlertLogSummaryMessageTypeUnknown       AlertLogSummaryMessageTypeEnum = "UNKNOWN"
	AlertLogSummaryMessageTypeIncidentError AlertLogSummaryMessageTypeEnum = "INCIDENT_ERROR"
	AlertLogSummaryMessageTypeError         AlertLogSummaryMessageTypeEnum = "ERROR"
	AlertLogSummaryMessageTypeWarning       AlertLogSummaryMessageTypeEnum = "WARNING"
	AlertLogSummaryMessageTypeNotification  AlertLogSummaryMessageTypeEnum = "NOTIFICATION"
	AlertLogSummaryMessageTypeTrace         AlertLogSummaryMessageTypeEnum = "TRACE"
)

var mappingAlertLogSummaryMessageTypeEnum = map[string]AlertLogSummaryMessageTypeEnum{
	"UNKNOWN":        AlertLogSummaryMessageTypeUnknown,
	"INCIDENT_ERROR": AlertLogSummaryMessageTypeIncidentError,
	"ERROR":          AlertLogSummaryMessageTypeError,
	"WARNING":        AlertLogSummaryMessageTypeWarning,
	"NOTIFICATION":   AlertLogSummaryMessageTypeNotification,
	"TRACE":          AlertLogSummaryMessageTypeTrace,
}

var mappingAlertLogSummaryMessageTypeEnumLowerCase = map[string]AlertLogSummaryMessageTypeEnum{
	"unknown":        AlertLogSummaryMessageTypeUnknown,
	"incident_error": AlertLogSummaryMessageTypeIncidentError,
	"error":          AlertLogSummaryMessageTypeError,
	"warning":        AlertLogSummaryMessageTypeWarning,
	"notification":   AlertLogSummaryMessageTypeNotification,
	"trace":          AlertLogSummaryMessageTypeTrace,
}

// GetAlertLogSummaryMessageTypeEnumValues Enumerates the set of values for AlertLogSummaryMessageTypeEnum
func GetAlertLogSummaryMessageTypeEnumValues() []AlertLogSummaryMessageTypeEnum {
	values := make([]AlertLogSummaryMessageTypeEnum, 0)
	for _, v := range mappingAlertLogSummaryMessageTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAlertLogSummaryMessageTypeEnumStringValues Enumerates the set of values in String for AlertLogSummaryMessageTypeEnum
func GetAlertLogSummaryMessageTypeEnumStringValues() []string {
	return []string{
		"UNKNOWN",
		"INCIDENT_ERROR",
		"ERROR",
		"WARNING",
		"NOTIFICATION",
		"TRACE",
	}
}

// GetMappingAlertLogSummaryMessageTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAlertLogSummaryMessageTypeEnum(val string) (AlertLogSummaryMessageTypeEnum, bool) {
	enum, ok := mappingAlertLogSummaryMessageTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
