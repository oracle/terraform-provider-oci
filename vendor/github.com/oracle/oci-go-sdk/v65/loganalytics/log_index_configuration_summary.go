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

// LogIndexConfigurationSummary Configuration that controls how a particular data type uses log indexes.
type LogIndexConfigurationSummary struct {

	// The data type this configuration applies to (e.g. LOG, APM, …).
	DataType StorageDataTypeEnum `mandatory:"true" json:"dataType"`

	// Indicates whether the log‑index feature is enabled for this data type.
	IsLogIndexEnabled *bool `mandatory:"true" json:"isLogIndexEnabled"`

	// The field name in the event/schema that is used as logical log set for this data type (e.g. "logset", "domain").
	LogSetMappedToField *string `mandatory:"false" json:"logSetMappedToField"`

	// Number of log indexes to allocate.  Default = 1.
	// Must be ≤ the system‑wide limit (default 5).
	LogIndexCount *int `mandatory:"false" json:"logIndexCount"`

	// How a log set is mapped to a log index.
	// RANDOM – system picks a log index automatically.
	// SPECIFIC – user maps log sets explicitly.
	MappingMechanism LogIndexConfigurationSummaryMappingMechanismEnum `mandatory:"false" json:"mappingMechanism,omitempty"`
}

func (m LogIndexConfigurationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LogIndexConfigurationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingStorageDataTypeEnum(string(m.DataType)); !ok && m.DataType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataType: %s. Supported values are: %s.", m.DataType, strings.Join(GetStorageDataTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingLogIndexConfigurationSummaryMappingMechanismEnum(string(m.MappingMechanism)); !ok && m.MappingMechanism != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MappingMechanism: %s. Supported values are: %s.", m.MappingMechanism, strings.Join(GetLogIndexConfigurationSummaryMappingMechanismEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LogIndexConfigurationSummaryMappingMechanismEnum Enum with underlying type: string
type LogIndexConfigurationSummaryMappingMechanismEnum string

// Set of constants representing the allowable values for LogIndexConfigurationSummaryMappingMechanismEnum
const (
	LogIndexConfigurationSummaryMappingMechanismRandom   LogIndexConfigurationSummaryMappingMechanismEnum = "RANDOM"
	LogIndexConfigurationSummaryMappingMechanismSpecific LogIndexConfigurationSummaryMappingMechanismEnum = "SPECIFIC"
)

var mappingLogIndexConfigurationSummaryMappingMechanismEnum = map[string]LogIndexConfigurationSummaryMappingMechanismEnum{
	"RANDOM":   LogIndexConfigurationSummaryMappingMechanismRandom,
	"SPECIFIC": LogIndexConfigurationSummaryMappingMechanismSpecific,
}

var mappingLogIndexConfigurationSummaryMappingMechanismEnumLowerCase = map[string]LogIndexConfigurationSummaryMappingMechanismEnum{
	"random":   LogIndexConfigurationSummaryMappingMechanismRandom,
	"specific": LogIndexConfigurationSummaryMappingMechanismSpecific,
}

// GetLogIndexConfigurationSummaryMappingMechanismEnumValues Enumerates the set of values for LogIndexConfigurationSummaryMappingMechanismEnum
func GetLogIndexConfigurationSummaryMappingMechanismEnumValues() []LogIndexConfigurationSummaryMappingMechanismEnum {
	values := make([]LogIndexConfigurationSummaryMappingMechanismEnum, 0)
	for _, v := range mappingLogIndexConfigurationSummaryMappingMechanismEnum {
		values = append(values, v)
	}
	return values
}

// GetLogIndexConfigurationSummaryMappingMechanismEnumStringValues Enumerates the set of values in String for LogIndexConfigurationSummaryMappingMechanismEnum
func GetLogIndexConfigurationSummaryMappingMechanismEnumStringValues() []string {
	return []string{
		"RANDOM",
		"SPECIFIC",
	}
}

// GetMappingLogIndexConfigurationSummaryMappingMechanismEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLogIndexConfigurationSummaryMappingMechanismEnum(val string) (LogIndexConfigurationSummaryMappingMechanismEnum, bool) {
	enum, ok := mappingLogIndexConfigurationSummaryMappingMechanismEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
