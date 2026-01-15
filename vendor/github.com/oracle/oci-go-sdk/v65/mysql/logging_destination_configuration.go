// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LoggingDestinationConfiguration Configuration parameters for a given destination.
type LoggingDestinationConfiguration struct {

	// Type of destination where MySQL telemetry is exposed to.
	Destination LoggingDestinationConfigurationDestinationEnum `mandatory:"true" json:"destination"`

	// List of configuration variables for a given destination type.
	DestinationConfigurations []DestinationConfiguration `mandatory:"true" json:"destinationConfigurations"`

	// List of MySQL telemetry types that can be exposed on a telemetry destination
	LogTypes []LoggingDestinationConfigurationLogTypesEnum `mandatory:"true" json:"logTypes"`
}

func (m LoggingDestinationConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LoggingDestinationConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLoggingDestinationConfigurationDestinationEnum(string(m.Destination)); !ok && m.Destination != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Destination: %s. Supported values are: %s.", m.Destination, strings.Join(GetLoggingDestinationConfigurationDestinationEnumStringValues(), ",")))
	}
	for _, val := range m.LogTypes {
		if _, ok := GetMappingLoggingDestinationConfigurationLogTypesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LogTypes: %s. Supported values are: %s.", val, strings.Join(GetLoggingDestinationConfigurationLogTypesEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LoggingDestinationConfigurationDestinationEnum Enum with underlying type: string
type LoggingDestinationConfigurationDestinationEnum string

// Set of constants representing the allowable values for LoggingDestinationConfigurationDestinationEnum
const (
	LoggingDestinationConfigurationDestinationLogAnalytics  LoggingDestinationConfigurationDestinationEnum = "LOG_ANALYTICS"
	LoggingDestinationConfigurationDestinationOpenTelemetry LoggingDestinationConfigurationDestinationEnum = "OPEN_TELEMETRY"
)

var mappingLoggingDestinationConfigurationDestinationEnum = map[string]LoggingDestinationConfigurationDestinationEnum{
	"LOG_ANALYTICS":  LoggingDestinationConfigurationDestinationLogAnalytics,
	"OPEN_TELEMETRY": LoggingDestinationConfigurationDestinationOpenTelemetry,
}

var mappingLoggingDestinationConfigurationDestinationEnumLowerCase = map[string]LoggingDestinationConfigurationDestinationEnum{
	"log_analytics":  LoggingDestinationConfigurationDestinationLogAnalytics,
	"open_telemetry": LoggingDestinationConfigurationDestinationOpenTelemetry,
}

// GetLoggingDestinationConfigurationDestinationEnumValues Enumerates the set of values for LoggingDestinationConfigurationDestinationEnum
func GetLoggingDestinationConfigurationDestinationEnumValues() []LoggingDestinationConfigurationDestinationEnum {
	values := make([]LoggingDestinationConfigurationDestinationEnum, 0)
	for _, v := range mappingLoggingDestinationConfigurationDestinationEnum {
		values = append(values, v)
	}
	return values
}

// GetLoggingDestinationConfigurationDestinationEnumStringValues Enumerates the set of values in String for LoggingDestinationConfigurationDestinationEnum
func GetLoggingDestinationConfigurationDestinationEnumStringValues() []string {
	return []string{
		"LOG_ANALYTICS",
		"OPEN_TELEMETRY",
	}
}

// GetMappingLoggingDestinationConfigurationDestinationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLoggingDestinationConfigurationDestinationEnum(val string) (LoggingDestinationConfigurationDestinationEnum, bool) {
	enum, ok := mappingLoggingDestinationConfigurationDestinationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// LoggingDestinationConfigurationLogTypesEnum Enum with underlying type: string
type LoggingDestinationConfigurationLogTypesEnum string

// Set of constants representing the allowable values for LoggingDestinationConfigurationLogTypesEnum
const (
	LoggingDestinationConfigurationLogTypesErrorLog     LoggingDestinationConfigurationLogTypesEnum = "ERROR_LOG"
	LoggingDestinationConfigurationLogTypesGeneralLog   LoggingDestinationConfigurationLogTypesEnum = "GENERAL_LOG"
	LoggingDestinationConfigurationLogTypesSlowQueryLog LoggingDestinationConfigurationLogTypesEnum = "SLOW_QUERY_LOG"
	LoggingDestinationConfigurationLogTypesAuditLog     LoggingDestinationConfigurationLogTypesEnum = "AUDIT_LOG"
)

var mappingLoggingDestinationConfigurationLogTypesEnum = map[string]LoggingDestinationConfigurationLogTypesEnum{
	"ERROR_LOG":      LoggingDestinationConfigurationLogTypesErrorLog,
	"GENERAL_LOG":    LoggingDestinationConfigurationLogTypesGeneralLog,
	"SLOW_QUERY_LOG": LoggingDestinationConfigurationLogTypesSlowQueryLog,
	"AUDIT_LOG":      LoggingDestinationConfigurationLogTypesAuditLog,
}

var mappingLoggingDestinationConfigurationLogTypesEnumLowerCase = map[string]LoggingDestinationConfigurationLogTypesEnum{
	"error_log":      LoggingDestinationConfigurationLogTypesErrorLog,
	"general_log":    LoggingDestinationConfigurationLogTypesGeneralLog,
	"slow_query_log": LoggingDestinationConfigurationLogTypesSlowQueryLog,
	"audit_log":      LoggingDestinationConfigurationLogTypesAuditLog,
}

// GetLoggingDestinationConfigurationLogTypesEnumValues Enumerates the set of values for LoggingDestinationConfigurationLogTypesEnum
func GetLoggingDestinationConfigurationLogTypesEnumValues() []LoggingDestinationConfigurationLogTypesEnum {
	values := make([]LoggingDestinationConfigurationLogTypesEnum, 0)
	for _, v := range mappingLoggingDestinationConfigurationLogTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetLoggingDestinationConfigurationLogTypesEnumStringValues Enumerates the set of values in String for LoggingDestinationConfigurationLogTypesEnum
func GetLoggingDestinationConfigurationLogTypesEnumStringValues() []string {
	return []string{
		"ERROR_LOG",
		"GENERAL_LOG",
		"SLOW_QUERY_LOG",
		"AUDIT_LOG",
	}
}

// GetMappingLoggingDestinationConfigurationLogTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLoggingDestinationConfigurationLogTypesEnum(val string) (LoggingDestinationConfigurationLogTypesEnum, bool) {
	enum, ok := mappingLoggingDestinationConfigurationLogTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
