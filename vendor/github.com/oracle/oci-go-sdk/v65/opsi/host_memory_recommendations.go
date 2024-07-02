// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// HostMemoryRecommendations Contains memory recommendation.
type HostMemoryRecommendations struct {

	// Identify if an instance is abandoned.
	IsAbandonedInstance *bool `mandatory:"false" json:"isAbandonedInstance"`

	// Show if OPSI recommends to change memory capacity based on Memory utilization and current shape.
	MemoryOptimization *string `mandatory:"false" json:"memoryOptimization"`

	// Identify unused instances based on cpu, memory and network metrics.
	UnusedInstance HostMemoryRecommendationsUnusedInstanceEnum `mandatory:"false" json:"unusedInstance,omitempty"`
}

func (m HostMemoryRecommendations) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HostMemoryRecommendations) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingHostMemoryRecommendationsUnusedInstanceEnum(string(m.UnusedInstance)); !ok && m.UnusedInstance != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UnusedInstance: %s. Supported values are: %s.", m.UnusedInstance, strings.Join(GetHostMemoryRecommendationsUnusedInstanceEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m HostMemoryRecommendations) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeHostMemoryRecommendations HostMemoryRecommendations
	s := struct {
		DiscriminatorParam string `json:"metricRecommendationName"`
		MarshalTypeHostMemoryRecommendations
	}{
		"HOST_MEMORY_RECOMMENDATIONS",
		(MarshalTypeHostMemoryRecommendations)(m),
	}

	return json.Marshal(&s)
}

// HostMemoryRecommendationsUnusedInstanceEnum Enum with underlying type: string
type HostMemoryRecommendationsUnusedInstanceEnum string

// Set of constants representing the allowable values for HostMemoryRecommendationsUnusedInstanceEnum
const (
	HostMemoryRecommendationsUnusedInstanceInUse           HostMemoryRecommendationsUnusedInstanceEnum = "IN_USE"
	HostMemoryRecommendationsUnusedInstanceNotInUse        HostMemoryRecommendationsUnusedInstanceEnum = "NOT_IN_USE"
	HostMemoryRecommendationsUnusedInstanceIsNotDetermined HostMemoryRecommendationsUnusedInstanceEnum = "IS_NOT_DETERMINED"
)

var mappingHostMemoryRecommendationsUnusedInstanceEnum = map[string]HostMemoryRecommendationsUnusedInstanceEnum{
	"IN_USE":            HostMemoryRecommendationsUnusedInstanceInUse,
	"NOT_IN_USE":        HostMemoryRecommendationsUnusedInstanceNotInUse,
	"IS_NOT_DETERMINED": HostMemoryRecommendationsUnusedInstanceIsNotDetermined,
}

var mappingHostMemoryRecommendationsUnusedInstanceEnumLowerCase = map[string]HostMemoryRecommendationsUnusedInstanceEnum{
	"in_use":            HostMemoryRecommendationsUnusedInstanceInUse,
	"not_in_use":        HostMemoryRecommendationsUnusedInstanceNotInUse,
	"is_not_determined": HostMemoryRecommendationsUnusedInstanceIsNotDetermined,
}

// GetHostMemoryRecommendationsUnusedInstanceEnumValues Enumerates the set of values for HostMemoryRecommendationsUnusedInstanceEnum
func GetHostMemoryRecommendationsUnusedInstanceEnumValues() []HostMemoryRecommendationsUnusedInstanceEnum {
	values := make([]HostMemoryRecommendationsUnusedInstanceEnum, 0)
	for _, v := range mappingHostMemoryRecommendationsUnusedInstanceEnum {
		values = append(values, v)
	}
	return values
}

// GetHostMemoryRecommendationsUnusedInstanceEnumStringValues Enumerates the set of values in String for HostMemoryRecommendationsUnusedInstanceEnum
func GetHostMemoryRecommendationsUnusedInstanceEnumStringValues() []string {
	return []string{
		"IN_USE",
		"NOT_IN_USE",
		"IS_NOT_DETERMINED",
	}
}

// GetMappingHostMemoryRecommendationsUnusedInstanceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHostMemoryRecommendationsUnusedInstanceEnum(val string) (HostMemoryRecommendationsUnusedInstanceEnum, bool) {
	enum, ok := mappingHostMemoryRecommendationsUnusedInstanceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
