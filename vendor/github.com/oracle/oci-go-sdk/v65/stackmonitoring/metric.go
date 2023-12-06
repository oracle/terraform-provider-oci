// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Metric Details of a metric which is part of this metric extension
type Metric struct {

	// Name of the metric.
	Name *string `mandatory:"true" json:"name"`

	// Data type of value of this metric
	DataType MetricDataTypeEnum `mandatory:"true" json:"dataType"`

	// Display name of the metric.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Current metric need to be included as dimension or not
	IsDimension *bool `mandatory:"false" json:"isDimension"`

	// Compute Expression to calculate the value of this metric
	ComputeExpression *string `mandatory:"false" json:"computeExpression"`

	// Flag to marks whether a metric has to be uploaded or not. When isHidden = false -> Metric is uploaded, isHidden = true -> Metric is NOT uploaded
	IsHidden *bool `mandatory:"false" json:"isHidden"`

	// Metric category
	MetricCategory MetricMetricCategoryEnum `mandatory:"false" json:"metricCategory,omitempty"`

	// Unit of metric value
	Unit *string `mandatory:"false" json:"unit"`
}

func (m Metric) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Metric) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMetricDataTypeEnum(string(m.DataType)); !ok && m.DataType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataType: %s. Supported values are: %s.", m.DataType, strings.Join(GetMetricDataTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingMetricMetricCategoryEnum(string(m.MetricCategory)); !ok && m.MetricCategory != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MetricCategory: %s. Supported values are: %s.", m.MetricCategory, strings.Join(GetMetricMetricCategoryEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MetricDataTypeEnum Enum with underlying type: string
type MetricDataTypeEnum string

// Set of constants representing the allowable values for MetricDataTypeEnum
const (
	MetricDataTypeString MetricDataTypeEnum = "STRING"
	MetricDataTypeNumber MetricDataTypeEnum = "NUMBER"
)

var mappingMetricDataTypeEnum = map[string]MetricDataTypeEnum{
	"STRING": MetricDataTypeString,
	"NUMBER": MetricDataTypeNumber,
}

var mappingMetricDataTypeEnumLowerCase = map[string]MetricDataTypeEnum{
	"string": MetricDataTypeString,
	"number": MetricDataTypeNumber,
}

// GetMetricDataTypeEnumValues Enumerates the set of values for MetricDataTypeEnum
func GetMetricDataTypeEnumValues() []MetricDataTypeEnum {
	values := make([]MetricDataTypeEnum, 0)
	for _, v := range mappingMetricDataTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMetricDataTypeEnumStringValues Enumerates the set of values in String for MetricDataTypeEnum
func GetMetricDataTypeEnumStringValues() []string {
	return []string{
		"STRING",
		"NUMBER",
	}
}

// GetMappingMetricDataTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMetricDataTypeEnum(val string) (MetricDataTypeEnum, bool) {
	enum, ok := mappingMetricDataTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// MetricMetricCategoryEnum Enum with underlying type: string
type MetricMetricCategoryEnum string

// Set of constants representing the allowable values for MetricMetricCategoryEnum
const (
	MetricMetricCategoryLoad         MetricMetricCategoryEnum = "LOAD"
	MetricMetricCategoryUtilization  MetricMetricCategoryEnum = "UTILIZATION"
	MetricMetricCategoryCapacity     MetricMetricCategoryEnum = "CAPACITY"
	MetricMetricCategoryAvailability MetricMetricCategoryEnum = "AVAILABILITY"
)

var mappingMetricMetricCategoryEnum = map[string]MetricMetricCategoryEnum{
	"LOAD":         MetricMetricCategoryLoad,
	"UTILIZATION":  MetricMetricCategoryUtilization,
	"CAPACITY":     MetricMetricCategoryCapacity,
	"AVAILABILITY": MetricMetricCategoryAvailability,
}

var mappingMetricMetricCategoryEnumLowerCase = map[string]MetricMetricCategoryEnum{
	"load":         MetricMetricCategoryLoad,
	"utilization":  MetricMetricCategoryUtilization,
	"capacity":     MetricMetricCategoryCapacity,
	"availability": MetricMetricCategoryAvailability,
}

// GetMetricMetricCategoryEnumValues Enumerates the set of values for MetricMetricCategoryEnum
func GetMetricMetricCategoryEnumValues() []MetricMetricCategoryEnum {
	values := make([]MetricMetricCategoryEnum, 0)
	for _, v := range mappingMetricMetricCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetMetricMetricCategoryEnumStringValues Enumerates the set of values in String for MetricMetricCategoryEnum
func GetMetricMetricCategoryEnumStringValues() []string {
	return []string{
		"LOAD",
		"UTILIZATION",
		"CAPACITY",
		"AVAILABILITY",
	}
}

// GetMappingMetricMetricCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMetricMetricCategoryEnum(val string) (MetricMetricCategoryEnum, bool) {
	enum, ok := mappingMetricMetricCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
