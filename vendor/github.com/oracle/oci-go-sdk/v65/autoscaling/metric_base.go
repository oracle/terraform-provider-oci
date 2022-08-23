// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Autoscaling API
//
// Use the Autoscaling API to dynamically scale compute resources to meet application requirements. For more information about
// autoscaling, see Autoscaling (https://docs.cloud.oracle.com/Content/Compute/Tasks/autoscalinginstancepools.htm). For information about the
// Compute service, see Compute (https://docs.cloud.oracle.com/Content/Compute/home.htm).
//

package autoscaling

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MetricBase The representation of MetricBase
type MetricBase interface {

	// The period of time that the condition defined in the alarm must persist before the alarm state
	// changes from "OK" to "FIRING" or vice versa. For example, a value of 5 minutes means that the
	// alarm must persist in breaching the condition for five minutes before the alarm updates its
	// state to "FIRING"; likewise, the alarm must persist in not breaching the condition for five
	// minutes before the alarm updates its state to "OK."
	// The duration is specified as a string in ISO 8601 format (`PT10M` for ten minutes or `PT1H`
	// for one hour). Minimum: PT3M. Maximum: PT1H. Default: PT3M.
	GetPendingDuration() *string
}

type metricbase struct {
	JsonData        []byte
	PendingDuration *string `mandatory:"false" json:"pendingDuration"`
	MetricSource    string  `json:"metricSource"`
}

// UnmarshalJSON unmarshals json
func (m *metricbase) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalermetricbase metricbase
	s := struct {
		Model Unmarshalermetricbase
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.PendingDuration = s.Model.PendingDuration
	m.MetricSource = s.Model.MetricSource

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *metricbase) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.MetricSource {
	case "CUSTOM_QUERY":
		mm := CustomMetric{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COMPUTE_AGENT":
		mm := Metric{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for MetricBase: %s.", m.MetricSource)
		return *m, nil
	}
}

//GetPendingDuration returns PendingDuration
func (m metricbase) GetPendingDuration() *string {
	return m.PendingDuration
}

func (m metricbase) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m metricbase) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MetricBaseMetricSourceEnum Enum with underlying type: string
type MetricBaseMetricSourceEnum string

// Set of constants representing the allowable values for MetricBaseMetricSourceEnum
const (
	MetricBaseMetricSourceComputeAgent MetricBaseMetricSourceEnum = "COMPUTE_AGENT"
	MetricBaseMetricSourceCustomQuery  MetricBaseMetricSourceEnum = "CUSTOM_QUERY"
)

var mappingMetricBaseMetricSourceEnum = map[string]MetricBaseMetricSourceEnum{
	"COMPUTE_AGENT": MetricBaseMetricSourceComputeAgent,
	"CUSTOM_QUERY":  MetricBaseMetricSourceCustomQuery,
}

var mappingMetricBaseMetricSourceEnumLowerCase = map[string]MetricBaseMetricSourceEnum{
	"compute_agent": MetricBaseMetricSourceComputeAgent,
	"custom_query":  MetricBaseMetricSourceCustomQuery,
}

// GetMetricBaseMetricSourceEnumValues Enumerates the set of values for MetricBaseMetricSourceEnum
func GetMetricBaseMetricSourceEnumValues() []MetricBaseMetricSourceEnum {
	values := make([]MetricBaseMetricSourceEnum, 0)
	for _, v := range mappingMetricBaseMetricSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetMetricBaseMetricSourceEnumStringValues Enumerates the set of values in String for MetricBaseMetricSourceEnum
func GetMetricBaseMetricSourceEnumStringValues() []string {
	return []string{
		"COMPUTE_AGENT",
		"CUSTOM_QUERY",
	}
}

// GetMappingMetricBaseMetricSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMetricBaseMetricSourceEnum(val string) (MetricBaseMetricSourceEnum, bool) {
	enum, ok := mappingMetricBaseMetricSourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
