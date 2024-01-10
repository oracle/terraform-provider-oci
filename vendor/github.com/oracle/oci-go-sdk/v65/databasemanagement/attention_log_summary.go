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

// AttentionLogSummary The details for one attention log entry.
type AttentionLogSummary struct {

	// The urgency of the attention log.
	MessageUrgency AttentionLogSummaryMessageUrgencyEnum `mandatory:"true" json:"messageUrgency"`

	// The type of attention log message.
	MessageType AttentionLogSummaryMessageTypeEnum `mandatory:"true" json:"messageType"`

	// The contents of the attention log message.
	MessageContent *string `mandatory:"false" json:"messageContent"`

	// The date and time the attention log was created.
	Timestamp *common.SDKTime `mandatory:"false" json:"timestamp"`

	// The database scope for the attention log.
	Scope *string `mandatory:"false" json:"scope"`

	// The user who must act on the attention log message.
	TargetUser *string `mandatory:"false" json:"targetUser"`

	// The cause of the attention log.
	Cause *string `mandatory:"false" json:"cause"`

	// The recommended action to handle the attention log.
	Action *string `mandatory:"false" json:"action"`

	// The supplemental details of the attention log.
	SupplementalDetail *string `mandatory:"false" json:"supplementalDetail"`

	// The attention log file location.
	FileLocation *string `mandatory:"false" json:"fileLocation"`
}

func (m AttentionLogSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AttentionLogSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAttentionLogSummaryMessageUrgencyEnum(string(m.MessageUrgency)); !ok && m.MessageUrgency != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MessageUrgency: %s. Supported values are: %s.", m.MessageUrgency, strings.Join(GetAttentionLogSummaryMessageUrgencyEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAttentionLogSummaryMessageTypeEnum(string(m.MessageType)); !ok && m.MessageType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MessageType: %s. Supported values are: %s.", m.MessageType, strings.Join(GetAttentionLogSummaryMessageTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AttentionLogSummaryMessageUrgencyEnum Enum with underlying type: string
type AttentionLogSummaryMessageUrgencyEnum string

// Set of constants representing the allowable values for AttentionLogSummaryMessageUrgencyEnum
const (
	AttentionLogSummaryMessageUrgencyImmediate  AttentionLogSummaryMessageUrgencyEnum = "IMMEDIATE"
	AttentionLogSummaryMessageUrgencySoon       AttentionLogSummaryMessageUrgencyEnum = "SOON"
	AttentionLogSummaryMessageUrgencyDeferrable AttentionLogSummaryMessageUrgencyEnum = "DEFERRABLE"
	AttentionLogSummaryMessageUrgencyInfo       AttentionLogSummaryMessageUrgencyEnum = "INFO"
)

var mappingAttentionLogSummaryMessageUrgencyEnum = map[string]AttentionLogSummaryMessageUrgencyEnum{
	"IMMEDIATE":  AttentionLogSummaryMessageUrgencyImmediate,
	"SOON":       AttentionLogSummaryMessageUrgencySoon,
	"DEFERRABLE": AttentionLogSummaryMessageUrgencyDeferrable,
	"INFO":       AttentionLogSummaryMessageUrgencyInfo,
}

var mappingAttentionLogSummaryMessageUrgencyEnumLowerCase = map[string]AttentionLogSummaryMessageUrgencyEnum{
	"immediate":  AttentionLogSummaryMessageUrgencyImmediate,
	"soon":       AttentionLogSummaryMessageUrgencySoon,
	"deferrable": AttentionLogSummaryMessageUrgencyDeferrable,
	"info":       AttentionLogSummaryMessageUrgencyInfo,
}

// GetAttentionLogSummaryMessageUrgencyEnumValues Enumerates the set of values for AttentionLogSummaryMessageUrgencyEnum
func GetAttentionLogSummaryMessageUrgencyEnumValues() []AttentionLogSummaryMessageUrgencyEnum {
	values := make([]AttentionLogSummaryMessageUrgencyEnum, 0)
	for _, v := range mappingAttentionLogSummaryMessageUrgencyEnum {
		values = append(values, v)
	}
	return values
}

// GetAttentionLogSummaryMessageUrgencyEnumStringValues Enumerates the set of values in String for AttentionLogSummaryMessageUrgencyEnum
func GetAttentionLogSummaryMessageUrgencyEnumStringValues() []string {
	return []string{
		"IMMEDIATE",
		"SOON",
		"DEFERRABLE",
		"INFO",
	}
}

// GetMappingAttentionLogSummaryMessageUrgencyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAttentionLogSummaryMessageUrgencyEnum(val string) (AttentionLogSummaryMessageUrgencyEnum, bool) {
	enum, ok := mappingAttentionLogSummaryMessageUrgencyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AttentionLogSummaryMessageTypeEnum Enum with underlying type: string
type AttentionLogSummaryMessageTypeEnum string

// Set of constants representing the allowable values for AttentionLogSummaryMessageTypeEnum
const (
	AttentionLogSummaryMessageTypeUnknown       AttentionLogSummaryMessageTypeEnum = "UNKNOWN"
	AttentionLogSummaryMessageTypeIncidentError AttentionLogSummaryMessageTypeEnum = "INCIDENT_ERROR"
	AttentionLogSummaryMessageTypeError         AttentionLogSummaryMessageTypeEnum = "ERROR"
	AttentionLogSummaryMessageTypeWarning       AttentionLogSummaryMessageTypeEnum = "WARNING"
	AttentionLogSummaryMessageTypeNotification  AttentionLogSummaryMessageTypeEnum = "NOTIFICATION"
	AttentionLogSummaryMessageTypeTrace         AttentionLogSummaryMessageTypeEnum = "TRACE"
)

var mappingAttentionLogSummaryMessageTypeEnum = map[string]AttentionLogSummaryMessageTypeEnum{
	"UNKNOWN":        AttentionLogSummaryMessageTypeUnknown,
	"INCIDENT_ERROR": AttentionLogSummaryMessageTypeIncidentError,
	"ERROR":          AttentionLogSummaryMessageTypeError,
	"WARNING":        AttentionLogSummaryMessageTypeWarning,
	"NOTIFICATION":   AttentionLogSummaryMessageTypeNotification,
	"TRACE":          AttentionLogSummaryMessageTypeTrace,
}

var mappingAttentionLogSummaryMessageTypeEnumLowerCase = map[string]AttentionLogSummaryMessageTypeEnum{
	"unknown":        AttentionLogSummaryMessageTypeUnknown,
	"incident_error": AttentionLogSummaryMessageTypeIncidentError,
	"error":          AttentionLogSummaryMessageTypeError,
	"warning":        AttentionLogSummaryMessageTypeWarning,
	"notification":   AttentionLogSummaryMessageTypeNotification,
	"trace":          AttentionLogSummaryMessageTypeTrace,
}

// GetAttentionLogSummaryMessageTypeEnumValues Enumerates the set of values for AttentionLogSummaryMessageTypeEnum
func GetAttentionLogSummaryMessageTypeEnumValues() []AttentionLogSummaryMessageTypeEnum {
	values := make([]AttentionLogSummaryMessageTypeEnum, 0)
	for _, v := range mappingAttentionLogSummaryMessageTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAttentionLogSummaryMessageTypeEnumStringValues Enumerates the set of values in String for AttentionLogSummaryMessageTypeEnum
func GetAttentionLogSummaryMessageTypeEnumStringValues() []string {
	return []string{
		"UNKNOWN",
		"INCIDENT_ERROR",
		"ERROR",
		"WARNING",
		"NOTIFICATION",
		"TRACE",
	}
}

// GetMappingAttentionLogSummaryMessageTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAttentionLogSummaryMessageTypeEnum(val string) (AttentionLogSummaryMessageTypeEnum, bool) {
	enum, ok := mappingAttentionLogSummaryMessageTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
