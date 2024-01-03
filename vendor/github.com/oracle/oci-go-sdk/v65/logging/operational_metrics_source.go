// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Management API
//
// Use the Logging Management API to create, read, list, update, move and delete
// log groups, log objects, log saved searches, and agent configurations.
// For more information, see Logging Overview (https://docs.cloud.oracle.com/iaas/Content/Logging/Concepts/loggingoverview.htm).
//

package logging

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OperationalMetricsSource Unified monitoring agent operational metrics source object.
type OperationalMetricsSource struct {

	// Type of the unified monitoring agent operational metrics source object.
	Type OperationalMetricsSourceTypeEnum `mandatory:"true" json:"type"`

	RecordInput *OperationalMetricsRecordInput `mandatory:"true" json:"recordInput"`

	// List of unified monitoring agent operational metrics.
	Metrics []string `mandatory:"false" json:"metrics"`
}

func (m OperationalMetricsSource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OperationalMetricsSource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOperationalMetricsSourceTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetOperationalMetricsSourceTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OperationalMetricsSourceTypeEnum Enum with underlying type: string
type OperationalMetricsSourceTypeEnum string

// Set of constants representing the allowable values for OperationalMetricsSourceTypeEnum
const (
	OperationalMetricsSourceTypeUmaMetrics OperationalMetricsSourceTypeEnum = "UMA_METRICS"
)

var mappingOperationalMetricsSourceTypeEnum = map[string]OperationalMetricsSourceTypeEnum{
	"UMA_METRICS": OperationalMetricsSourceTypeUmaMetrics,
}

var mappingOperationalMetricsSourceTypeEnumLowerCase = map[string]OperationalMetricsSourceTypeEnum{
	"uma_metrics": OperationalMetricsSourceTypeUmaMetrics,
}

// GetOperationalMetricsSourceTypeEnumValues Enumerates the set of values for OperationalMetricsSourceTypeEnum
func GetOperationalMetricsSourceTypeEnumValues() []OperationalMetricsSourceTypeEnum {
	values := make([]OperationalMetricsSourceTypeEnum, 0)
	for _, v := range mappingOperationalMetricsSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationalMetricsSourceTypeEnumStringValues Enumerates the set of values in String for OperationalMetricsSourceTypeEnum
func GetOperationalMetricsSourceTypeEnumStringValues() []string {
	return []string{
		"UMA_METRICS",
	}
}

// GetMappingOperationalMetricsSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationalMetricsSourceTypeEnum(val string) (OperationalMetricsSourceTypeEnum, bool) {
	enum, ok := mappingOperationalMetricsSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
