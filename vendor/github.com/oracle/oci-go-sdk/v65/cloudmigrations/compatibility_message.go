// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CompatibilityMessage Information about shape compatibility with the client's current resource configuration.
type CompatibilityMessage struct {

	// Severity level of the compatibility issue.
	Severity CompatibilityMessageSeverityEnum `mandatory:"false" json:"severity,omitempty"`

	// Name of the compatibility issue.
	Name CompatibilityMessageNameEnum `mandatory:"false" json:"name,omitempty"`

	// Detailed description of the compatibility issue.
	Message *string `mandatory:"false" json:"message"`
}

func (m CompatibilityMessage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CompatibilityMessage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCompatibilityMessageSeverityEnum(string(m.Severity)); !ok && m.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", m.Severity, strings.Join(GetCompatibilityMessageSeverityEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCompatibilityMessageNameEnum(string(m.Name)); !ok && m.Name != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Name: %s. Supported values are: %s.", m.Name, strings.Join(GetCompatibilityMessageNameEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CompatibilityMessageSeverityEnum Enum with underlying type: string
type CompatibilityMessageSeverityEnum string

// Set of constants representing the allowable values for CompatibilityMessageSeverityEnum
const (
	CompatibilityMessageSeverityError   CompatibilityMessageSeverityEnum = "ERROR"
	CompatibilityMessageSeverityWarning CompatibilityMessageSeverityEnum = "WARNING"
	CompatibilityMessageSeverityInfo    CompatibilityMessageSeverityEnum = "INFO"
)

var mappingCompatibilityMessageSeverityEnum = map[string]CompatibilityMessageSeverityEnum{
	"ERROR":   CompatibilityMessageSeverityError,
	"WARNING": CompatibilityMessageSeverityWarning,
	"INFO":    CompatibilityMessageSeverityInfo,
}

var mappingCompatibilityMessageSeverityEnumLowerCase = map[string]CompatibilityMessageSeverityEnum{
	"error":   CompatibilityMessageSeverityError,
	"warning": CompatibilityMessageSeverityWarning,
	"info":    CompatibilityMessageSeverityInfo,
}

// GetCompatibilityMessageSeverityEnumValues Enumerates the set of values for CompatibilityMessageSeverityEnum
func GetCompatibilityMessageSeverityEnumValues() []CompatibilityMessageSeverityEnum {
	values := make([]CompatibilityMessageSeverityEnum, 0)
	for _, v := range mappingCompatibilityMessageSeverityEnum {
		values = append(values, v)
	}
	return values
}

// GetCompatibilityMessageSeverityEnumStringValues Enumerates the set of values in String for CompatibilityMessageSeverityEnum
func GetCompatibilityMessageSeverityEnumStringValues() []string {
	return []string{
		"ERROR",
		"WARNING",
		"INFO",
	}
}

// GetMappingCompatibilityMessageSeverityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCompatibilityMessageSeverityEnum(val string) (CompatibilityMessageSeverityEnum, bool) {
	enum, ok := mappingCompatibilityMessageSeverityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CompatibilityMessageNameEnum Enum with underlying type: string
type CompatibilityMessageNameEnum string

// Set of constants representing the allowable values for CompatibilityMessageNameEnum
const (
	CompatibilityMessageNameNotEnoughData                 CompatibilityMessageNameEnum = "NOT_ENOUGH_DATA"
	CompatibilityMessageNameInvalidData                   CompatibilityMessageNameEnum = "INVALID_DATA"
	CompatibilityMessageNameCpuCompatibilityWarning       CompatibilityMessageNameEnum = "CPU_COMPATIBILITY_WARNING"
	CompatibilityMessageNameCpuMetricInfo                 CompatibilityMessageNameEnum = "CPU_METRIC_INFO"
	CompatibilityMessageNameMemoryCompatibilityWarning    CompatibilityMessageNameEnum = "MEMORY_COMPATIBILITY_WARNING"
	CompatibilityMessageNameMemoryMetricInfo              CompatibilityMessageNameEnum = "MEMORY_METRIC_INFO"
	CompatibilityMessageNameVnicsCompatibilityWarning     CompatibilityMessageNameEnum = "VNICS_COMPATIBILITY_WARNING"
	CompatibilityMessageNameBandwidthCompatibilityWarning CompatibilityMessageNameEnum = "BANDWIDTH_COMPATIBILITY_WARNING"
	CompatibilityMessageNameGpuCompatibilityWarning       CompatibilityMessageNameEnum = "GPU_COMPATIBILITY_WARNING"
	CompatibilityMessageNameOsWarning                     CompatibilityMessageNameEnum = "OS_WARNING"
)

var mappingCompatibilityMessageNameEnum = map[string]CompatibilityMessageNameEnum{
	"NOT_ENOUGH_DATA":                 CompatibilityMessageNameNotEnoughData,
	"INVALID_DATA":                    CompatibilityMessageNameInvalidData,
	"CPU_COMPATIBILITY_WARNING":       CompatibilityMessageNameCpuCompatibilityWarning,
	"CPU_METRIC_INFO":                 CompatibilityMessageNameCpuMetricInfo,
	"MEMORY_COMPATIBILITY_WARNING":    CompatibilityMessageNameMemoryCompatibilityWarning,
	"MEMORY_METRIC_INFO":              CompatibilityMessageNameMemoryMetricInfo,
	"VNICS_COMPATIBILITY_WARNING":     CompatibilityMessageNameVnicsCompatibilityWarning,
	"BANDWIDTH_COMPATIBILITY_WARNING": CompatibilityMessageNameBandwidthCompatibilityWarning,
	"GPU_COMPATIBILITY_WARNING":       CompatibilityMessageNameGpuCompatibilityWarning,
	"OS_WARNING":                      CompatibilityMessageNameOsWarning,
}

var mappingCompatibilityMessageNameEnumLowerCase = map[string]CompatibilityMessageNameEnum{
	"not_enough_data":                 CompatibilityMessageNameNotEnoughData,
	"invalid_data":                    CompatibilityMessageNameInvalidData,
	"cpu_compatibility_warning":       CompatibilityMessageNameCpuCompatibilityWarning,
	"cpu_metric_info":                 CompatibilityMessageNameCpuMetricInfo,
	"memory_compatibility_warning":    CompatibilityMessageNameMemoryCompatibilityWarning,
	"memory_metric_info":              CompatibilityMessageNameMemoryMetricInfo,
	"vnics_compatibility_warning":     CompatibilityMessageNameVnicsCompatibilityWarning,
	"bandwidth_compatibility_warning": CompatibilityMessageNameBandwidthCompatibilityWarning,
	"gpu_compatibility_warning":       CompatibilityMessageNameGpuCompatibilityWarning,
	"os_warning":                      CompatibilityMessageNameOsWarning,
}

// GetCompatibilityMessageNameEnumValues Enumerates the set of values for CompatibilityMessageNameEnum
func GetCompatibilityMessageNameEnumValues() []CompatibilityMessageNameEnum {
	values := make([]CompatibilityMessageNameEnum, 0)
	for _, v := range mappingCompatibilityMessageNameEnum {
		values = append(values, v)
	}
	return values
}

// GetCompatibilityMessageNameEnumStringValues Enumerates the set of values in String for CompatibilityMessageNameEnum
func GetCompatibilityMessageNameEnumStringValues() []string {
	return []string{
		"NOT_ENOUGH_DATA",
		"INVALID_DATA",
		"CPU_COMPATIBILITY_WARNING",
		"CPU_METRIC_INFO",
		"MEMORY_COMPATIBILITY_WARNING",
		"MEMORY_METRIC_INFO",
		"VNICS_COMPATIBILITY_WARNING",
		"BANDWIDTH_COMPATIBILITY_WARNING",
		"GPU_COMPATIBILITY_WARNING",
		"OS_WARNING",
	}
}

// GetMappingCompatibilityMessageNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCompatibilityMessageNameEnum(val string) (CompatibilityMessageNameEnum, bool) {
	enum, ok := mappingCompatibilityMessageNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
