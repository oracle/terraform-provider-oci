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

// AverageResourceAssessmentStrategy The strategy based on average usage.
type AverageResourceAssessmentStrategy struct {

	// The real resource usage is multiplied to this number before making any recommendation.
	AdjustmentMultiplier *float32 `mandatory:"false" json:"adjustmentMultiplier"`

	// The type of resource.
	ResourceType ResourceAssessmentStrategyResourceTypeEnum `mandatory:"true" json:"resourceType"`

	// The current state of the migration plan.
	MetricType MetricTypeEnum `mandatory:"false" json:"metricType,omitempty"`

	// The current state of the migration plan.
	MetricTimeWindow MetricTimeWindowEnum `mandatory:"false" json:"metricTimeWindow,omitempty"`
}

// GetResourceType returns ResourceType
func (m AverageResourceAssessmentStrategy) GetResourceType() ResourceAssessmentStrategyResourceTypeEnum {
	return m.ResourceType
}

func (m AverageResourceAssessmentStrategy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AverageResourceAssessmentStrategy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingResourceAssessmentStrategyResourceTypeEnum(string(m.ResourceType)); !ok && m.ResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceType: %s. Supported values are: %s.", m.ResourceType, strings.Join(GetResourceAssessmentStrategyResourceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMetricTypeEnum(string(m.MetricType)); !ok && m.MetricType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MetricType: %s. Supported values are: %s.", m.MetricType, strings.Join(GetMetricTypeEnumStringValues(), ",")))
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
func (m AverageResourceAssessmentStrategy) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAverageResourceAssessmentStrategy AverageResourceAssessmentStrategy
	s := struct {
		DiscriminatorParam string `json:"strategyType"`
		MarshalTypeAverageResourceAssessmentStrategy
	}{
		"AVERAGE",
		(MarshalTypeAverageResourceAssessmentStrategy)(m),
	}

	return json.Marshal(&s)
}
