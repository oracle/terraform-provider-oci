// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Manager API
//
// Use the Resource Manager API to automate deployment and operations for all Oracle Cloud Infrastructure resources.
// Using the infrastructure-as-code (IaC) model, the service is based on Terraform, an open source industry standard that lets DevOps engineers develop and deploy their infrastructure anywhere.
// For more information, see
// the Resource Manager documentation (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/home.htm).
//

package resourcemanager

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// LogEntry Log entry for an operation resulting from a job's execution.
type LogEntry struct {

	// Specifies the log type for the log entry.
	Type LogEntryTypeEnum `mandatory:"false" json:"type,omitempty"`

	// Specifies the severity level of the log entry.
	Level LogEntryLevelEnum `mandatory:"false" json:"level,omitempty"`

	// The date and time of the log entry.
	// Format is defined by RFC3339.
	// Example: `2020-01-25T21:10:29.600Z`
	Timestamp *common.SDKTime `mandatory:"false" json:"timestamp"`

	// The log entry value.
	Message *string `mandatory:"false" json:"message"`
}

func (m LogEntry) String() string {
	return common.PointerString(m)
}

// LogEntryTypeEnum Enum with underlying type: string
type LogEntryTypeEnum string

// Set of constants representing the allowable values for LogEntryTypeEnum
const (
	LogEntryTypeTerraformConsole LogEntryTypeEnum = "TERRAFORM_CONSOLE"
)

var mappingLogEntryType = map[string]LogEntryTypeEnum{
	"TERRAFORM_CONSOLE": LogEntryTypeTerraformConsole,
}

// GetLogEntryTypeEnumValues Enumerates the set of values for LogEntryTypeEnum
func GetLogEntryTypeEnumValues() []LogEntryTypeEnum {
	values := make([]LogEntryTypeEnum, 0)
	for _, v := range mappingLogEntryType {
		values = append(values, v)
	}
	return values
}

// LogEntryLevelEnum Enum with underlying type: string
type LogEntryLevelEnum string

// Set of constants representing the allowable values for LogEntryLevelEnum
const (
	LogEntryLevelTrace LogEntryLevelEnum = "TRACE"
	LogEntryLevelDebug LogEntryLevelEnum = "DEBUG"
	LogEntryLevelInfo  LogEntryLevelEnum = "INFO"
	LogEntryLevelWarn  LogEntryLevelEnum = "WARN"
	LogEntryLevelError LogEntryLevelEnum = "ERROR"
	LogEntryLevelFatal LogEntryLevelEnum = "FATAL"
)

var mappingLogEntryLevel = map[string]LogEntryLevelEnum{
	"TRACE": LogEntryLevelTrace,
	"DEBUG": LogEntryLevelDebug,
	"INFO":  LogEntryLevelInfo,
	"WARN":  LogEntryLevelWarn,
	"ERROR": LogEntryLevelError,
	"FATAL": LogEntryLevelFatal,
}

// GetLogEntryLevelEnumValues Enumerates the set of values for LogEntryLevelEnum
func GetLogEntryLevelEnumValues() []LogEntryLevelEnum {
	values := make([]LogEntryLevelEnum, 0)
	for _, v := range mappingLogEntryLevel {
		values = append(values, v)
	}
	return values
}
