// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.oracle.com/iaas/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// HostInsightHostRecommendations Contains recommendations depending of resource metric received.
type HostInsightHostRecommendations interface {
}

type hostinsighthostrecommendations struct {
	JsonData                 []byte
	MetricRecommendationName string `json:"metricRecommendationName"`
}

// UnmarshalJSON unmarshals json
func (m *hostinsighthostrecommendations) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerhostinsighthostrecommendations hostinsighthostrecommendations
	s := struct {
		Model Unmarshalerhostinsighthostrecommendations
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.MetricRecommendationName = s.Model.MetricRecommendationName

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *hostinsighthostrecommendations) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.MetricRecommendationName {
	case "HOST_CPU_RECOMMENDATIONS":
		mm := HostCpuRecommendations{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HOST_NETWORK_RECOMMENDATIONS":
		mm := HostNetworkRecommendations{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HOST_MEMORY_RECOMMENDATIONS":
		mm := HostMemoryRecommendations{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HOST_STORAGE_RECOMMENDATIONS":
		mm := HostStorageRecommendations{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for HostInsightHostRecommendations: %s.", m.MetricRecommendationName)
		return *m, nil
	}
}

func (m hostinsighthostrecommendations) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m hostinsighthostrecommendations) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// HostInsightHostRecommendationsMetricRecommendationNameEnum Enum with underlying type: string
type HostInsightHostRecommendationsMetricRecommendationNameEnum string

// Set of constants representing the allowable values for HostInsightHostRecommendationsMetricRecommendationNameEnum
const (
	HostInsightHostRecommendationsMetricRecommendationNameCpuRecommendations     HostInsightHostRecommendationsMetricRecommendationNameEnum = "HOST_CPU_RECOMMENDATIONS"
	HostInsightHostRecommendationsMetricRecommendationNameMemoryRecommendations  HostInsightHostRecommendationsMetricRecommendationNameEnum = "HOST_MEMORY_RECOMMENDATIONS"
	HostInsightHostRecommendationsMetricRecommendationNameNetworkRecommendations HostInsightHostRecommendationsMetricRecommendationNameEnum = "HOST_NETWORK_RECOMMENDATIONS"
	HostInsightHostRecommendationsMetricRecommendationNameStorageRecommendations HostInsightHostRecommendationsMetricRecommendationNameEnum = "HOST_STORAGE_RECOMMENDATIONS"
)

var mappingHostInsightHostRecommendationsMetricRecommendationNameEnum = map[string]HostInsightHostRecommendationsMetricRecommendationNameEnum{
	"HOST_CPU_RECOMMENDATIONS":     HostInsightHostRecommendationsMetricRecommendationNameCpuRecommendations,
	"HOST_MEMORY_RECOMMENDATIONS":  HostInsightHostRecommendationsMetricRecommendationNameMemoryRecommendations,
	"HOST_NETWORK_RECOMMENDATIONS": HostInsightHostRecommendationsMetricRecommendationNameNetworkRecommendations,
	"HOST_STORAGE_RECOMMENDATIONS": HostInsightHostRecommendationsMetricRecommendationNameStorageRecommendations,
}

var mappingHostInsightHostRecommendationsMetricRecommendationNameEnumLowerCase = map[string]HostInsightHostRecommendationsMetricRecommendationNameEnum{
	"host_cpu_recommendations":     HostInsightHostRecommendationsMetricRecommendationNameCpuRecommendations,
	"host_memory_recommendations":  HostInsightHostRecommendationsMetricRecommendationNameMemoryRecommendations,
	"host_network_recommendations": HostInsightHostRecommendationsMetricRecommendationNameNetworkRecommendations,
	"host_storage_recommendations": HostInsightHostRecommendationsMetricRecommendationNameStorageRecommendations,
}

// GetHostInsightHostRecommendationsMetricRecommendationNameEnumValues Enumerates the set of values for HostInsightHostRecommendationsMetricRecommendationNameEnum
func GetHostInsightHostRecommendationsMetricRecommendationNameEnumValues() []HostInsightHostRecommendationsMetricRecommendationNameEnum {
	values := make([]HostInsightHostRecommendationsMetricRecommendationNameEnum, 0)
	for _, v := range mappingHostInsightHostRecommendationsMetricRecommendationNameEnum {
		values = append(values, v)
	}
	return values
}

// GetHostInsightHostRecommendationsMetricRecommendationNameEnumStringValues Enumerates the set of values in String for HostInsightHostRecommendationsMetricRecommendationNameEnum
func GetHostInsightHostRecommendationsMetricRecommendationNameEnumStringValues() []string {
	return []string{
		"HOST_CPU_RECOMMENDATIONS",
		"HOST_MEMORY_RECOMMENDATIONS",
		"HOST_NETWORK_RECOMMENDATIONS",
		"HOST_STORAGE_RECOMMENDATIONS",
	}
}

// GetMappingHostInsightHostRecommendationsMetricRecommendationNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHostInsightHostRecommendationsMetricRecommendationNameEnum(val string) (HostInsightHostRecommendationsMetricRecommendationNameEnum, bool) {
	enum, ok := mappingHostInsightHostRecommendationsMetricRecommendationNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
