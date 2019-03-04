// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Auto Scaling API
//
// Auto Scaling API spec
//

package autoscaling

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Metric Metric threshold details
type Metric struct {
	MetricType MetricMetricTypeEnum `mandatory:"true" json:"metricType"`

	Threshold *Threshold `mandatory:"true" json:"threshold"`
}

func (m Metric) String() string {
	return common.PointerString(m)
}

// MetricMetricTypeEnum Enum with underlying type: string
type MetricMetricTypeEnum string

// Set of constants representing the allowable values for MetricMetricTypeEnum
const (
	MetricMetricTypeCpuUtilization    MetricMetricTypeEnum = "CPU_UTILIZATION"
	MetricMetricTypeMemoryUtilization MetricMetricTypeEnum = "MEMORY_UTILIZATION"
)

var mappingMetricMetricType = map[string]MetricMetricTypeEnum{
	"CPU_UTILIZATION":    MetricMetricTypeCpuUtilization,
	"MEMORY_UTILIZATION": MetricMetricTypeMemoryUtilization,
}

// GetMetricMetricTypeEnumValues Enumerates the set of values for MetricMetricTypeEnum
func GetMetricMetricTypeEnumValues() []MetricMetricTypeEnum {
	values := make([]MetricMetricTypeEnum, 0)
	for _, v := range mappingMetricMetricType {
		values = append(values, v)
	}
	return values
}
