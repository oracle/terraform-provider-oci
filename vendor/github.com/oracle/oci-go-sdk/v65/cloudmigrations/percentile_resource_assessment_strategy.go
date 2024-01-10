// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PercentileResourceAssessmentStrategy The strategy based on percentile usage.
type PercentileResourceAssessmentStrategy struct {

	// The real resource usage is multiplied to this number before making any recommendation.
	AdjustmentMultiplier *float32 `mandatory:"false" json:"adjustmentMultiplier"`

	// Percentile value
	Percentile PercentileResourceAssessmentStrategyPercentileEnum `mandatory:"true" json:"percentile"`

	// The type of resource.
	ResourceType ResourceAssessmentStrategyResourceTypeEnum `mandatory:"true" json:"resourceType"`

	// The current state of the migration plan.
	MetricTimeWindow MetricTimeWindowEnum `mandatory:"false" json:"metricTimeWindow,omitempty"`
}

// GetResourceType returns ResourceType
func (m PercentileResourceAssessmentStrategy) GetResourceType() ResourceAssessmentStrategyResourceTypeEnum {
	return m.ResourceType
}

func (m PercentileResourceAssessmentStrategy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PercentileResourceAssessmentStrategy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPercentileResourceAssessmentStrategyPercentileEnum(string(m.Percentile)); !ok && m.Percentile != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Percentile: %s. Supported values are: %s.", m.Percentile, strings.Join(GetPercentileResourceAssessmentStrategyPercentileEnumStringValues(), ",")))
	}

	if _, ok := GetMappingResourceAssessmentStrategyResourceTypeEnum(string(m.ResourceType)); !ok && m.ResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceType: %s. Supported values are: %s.", m.ResourceType, strings.Join(GetResourceAssessmentStrategyResourceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMetricTimeWindowEnum(string(m.MetricTimeWindow)); !ok && m.MetricTimeWindow != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MetricTimeWindow: %s. Supported values are: %s.", m.MetricTimeWindow, strings.Join(GetMetricTimeWindowEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PercentileResourceAssessmentStrategy) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePercentileResourceAssessmentStrategy PercentileResourceAssessmentStrategy
	s := struct {
		DiscriminatorParam string `json:"strategyType"`
		MarshalTypePercentileResourceAssessmentStrategy
	}{
		"PERCENTILE",
		(MarshalTypePercentileResourceAssessmentStrategy)(m),
	}

	return json.Marshal(&s)
}

// PercentileResourceAssessmentStrategyPercentileEnum Enum with underlying type: string
type PercentileResourceAssessmentStrategyPercentileEnum string

// Set of constants representing the allowable values for PercentileResourceAssessmentStrategyPercentileEnum
const (
	PercentileResourceAssessmentStrategyPercentileP50 PercentileResourceAssessmentStrategyPercentileEnum = "P50"
	PercentileResourceAssessmentStrategyPercentileP90 PercentileResourceAssessmentStrategyPercentileEnum = "P90"
	PercentileResourceAssessmentStrategyPercentileP95 PercentileResourceAssessmentStrategyPercentileEnum = "P95"
	PercentileResourceAssessmentStrategyPercentileP99 PercentileResourceAssessmentStrategyPercentileEnum = "P99"
)

var mappingPercentileResourceAssessmentStrategyPercentileEnum = map[string]PercentileResourceAssessmentStrategyPercentileEnum{
	"P50": PercentileResourceAssessmentStrategyPercentileP50,
	"P90": PercentileResourceAssessmentStrategyPercentileP90,
	"P95": PercentileResourceAssessmentStrategyPercentileP95,
	"P99": PercentileResourceAssessmentStrategyPercentileP99,
}

var mappingPercentileResourceAssessmentStrategyPercentileEnumLowerCase = map[string]PercentileResourceAssessmentStrategyPercentileEnum{
	"p50": PercentileResourceAssessmentStrategyPercentileP50,
	"p90": PercentileResourceAssessmentStrategyPercentileP90,
	"p95": PercentileResourceAssessmentStrategyPercentileP95,
	"p99": PercentileResourceAssessmentStrategyPercentileP99,
}

// GetPercentileResourceAssessmentStrategyPercentileEnumValues Enumerates the set of values for PercentileResourceAssessmentStrategyPercentileEnum
func GetPercentileResourceAssessmentStrategyPercentileEnumValues() []PercentileResourceAssessmentStrategyPercentileEnum {
	values := make([]PercentileResourceAssessmentStrategyPercentileEnum, 0)
	for _, v := range mappingPercentileResourceAssessmentStrategyPercentileEnum {
		values = append(values, v)
	}
	return values
}

// GetPercentileResourceAssessmentStrategyPercentileEnumStringValues Enumerates the set of values in String for PercentileResourceAssessmentStrategyPercentileEnum
func GetPercentileResourceAssessmentStrategyPercentileEnumStringValues() []string {
	return []string{
		"P50",
		"P90",
		"P95",
		"P99",
	}
}

// GetMappingPercentileResourceAssessmentStrategyPercentileEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPercentileResourceAssessmentStrategyPercentileEnum(val string) (PercentileResourceAssessmentStrategyPercentileEnum, bool) {
	enum, ok := mappingPercentileResourceAssessmentStrategyPercentileEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
