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

// HostCpuRecommendations Contains CPU recommendation.
type HostCpuRecommendations struct {

	// Show if OPSI recommends to change the shape of an instance and show recommended shape based on CPU utilization.
	Shape *string `mandatory:"false" json:"shape"`

	// Identify if an instance is abandoned.
	IsAbandonedInstance *bool `mandatory:"false" json:"isAbandonedInstance"`

	// Show if OPSI recommends to convert an instance to a burstable instance and show recommended cpu baseline if positive recommendation.
	Burstable HostCpuRecommendationsBurstableEnum `mandatory:"false" json:"burstable,omitempty"`

	// Identify unused instances based on cpu, memory and network metrics.
	UnusedInstance HostCpuRecommendationsUnusedInstanceEnum `mandatory:"false" json:"unusedInstance,omitempty"`
}

func (m HostCpuRecommendations) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HostCpuRecommendations) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingHostCpuRecommendationsBurstableEnum(string(m.Burstable)); !ok && m.Burstable != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Burstable: %s. Supported values are: %s.", m.Burstable, strings.Join(GetHostCpuRecommendationsBurstableEnumStringValues(), ",")))
	}
	if _, ok := GetMappingHostCpuRecommendationsUnusedInstanceEnum(string(m.UnusedInstance)); !ok && m.UnusedInstance != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UnusedInstance: %s. Supported values are: %s.", m.UnusedInstance, strings.Join(GetHostCpuRecommendationsUnusedInstanceEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m HostCpuRecommendations) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeHostCpuRecommendations HostCpuRecommendations
	s := struct {
		DiscriminatorParam string `json:"metricRecommendationName"`
		MarshalTypeHostCpuRecommendations
	}{
		"HOST_CPU_RECOMMENDATIONS",
		(MarshalTypeHostCpuRecommendations)(m),
	}

	return json.Marshal(&s)
}

// HostCpuRecommendationsBurstableEnum Enum with underlying type: string
type HostCpuRecommendationsBurstableEnum string

// Set of constants representing the allowable values for HostCpuRecommendationsBurstableEnum
const (
	HostCpuRecommendationsBurstableBaseline18       HostCpuRecommendationsBurstableEnum = "BASELINE_1_8"
	HostCpuRecommendationsBurstableBaseline12       HostCpuRecommendationsBurstableEnum = "BASELINE_1_2"
	HostCpuRecommendationsBurstableNoRecommendation HostCpuRecommendationsBurstableEnum = "NO_RECOMMENDATION"
	HostCpuRecommendationsBurstableDisableBurstable HostCpuRecommendationsBurstableEnum = "DISABLE_BURSTABLE"
)

var mappingHostCpuRecommendationsBurstableEnum = map[string]HostCpuRecommendationsBurstableEnum{
	"BASELINE_1_8":      HostCpuRecommendationsBurstableBaseline18,
	"BASELINE_1_2":      HostCpuRecommendationsBurstableBaseline12,
	"NO_RECOMMENDATION": HostCpuRecommendationsBurstableNoRecommendation,
	"DISABLE_BURSTABLE": HostCpuRecommendationsBurstableDisableBurstable,
}

var mappingHostCpuRecommendationsBurstableEnumLowerCase = map[string]HostCpuRecommendationsBurstableEnum{
	"baseline_1_8":      HostCpuRecommendationsBurstableBaseline18,
	"baseline_1_2":      HostCpuRecommendationsBurstableBaseline12,
	"no_recommendation": HostCpuRecommendationsBurstableNoRecommendation,
	"disable_burstable": HostCpuRecommendationsBurstableDisableBurstable,
}

// GetHostCpuRecommendationsBurstableEnumValues Enumerates the set of values for HostCpuRecommendationsBurstableEnum
func GetHostCpuRecommendationsBurstableEnumValues() []HostCpuRecommendationsBurstableEnum {
	values := make([]HostCpuRecommendationsBurstableEnum, 0)
	for _, v := range mappingHostCpuRecommendationsBurstableEnum {
		values = append(values, v)
	}
	return values
}

// GetHostCpuRecommendationsBurstableEnumStringValues Enumerates the set of values in String for HostCpuRecommendationsBurstableEnum
func GetHostCpuRecommendationsBurstableEnumStringValues() []string {
	return []string{
		"BASELINE_1_8",
		"BASELINE_1_2",
		"NO_RECOMMENDATION",
		"DISABLE_BURSTABLE",
	}
}

// GetMappingHostCpuRecommendationsBurstableEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHostCpuRecommendationsBurstableEnum(val string) (HostCpuRecommendationsBurstableEnum, bool) {
	enum, ok := mappingHostCpuRecommendationsBurstableEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// HostCpuRecommendationsUnusedInstanceEnum Enum with underlying type: string
type HostCpuRecommendationsUnusedInstanceEnum string

// Set of constants representing the allowable values for HostCpuRecommendationsUnusedInstanceEnum
const (
	HostCpuRecommendationsUnusedInstanceInUse           HostCpuRecommendationsUnusedInstanceEnum = "IN_USE"
	HostCpuRecommendationsUnusedInstanceNotInUse        HostCpuRecommendationsUnusedInstanceEnum = "NOT_IN_USE"
	HostCpuRecommendationsUnusedInstanceIsNotDetermined HostCpuRecommendationsUnusedInstanceEnum = "IS_NOT_DETERMINED"
)

var mappingHostCpuRecommendationsUnusedInstanceEnum = map[string]HostCpuRecommendationsUnusedInstanceEnum{
	"IN_USE":            HostCpuRecommendationsUnusedInstanceInUse,
	"NOT_IN_USE":        HostCpuRecommendationsUnusedInstanceNotInUse,
	"IS_NOT_DETERMINED": HostCpuRecommendationsUnusedInstanceIsNotDetermined,
}

var mappingHostCpuRecommendationsUnusedInstanceEnumLowerCase = map[string]HostCpuRecommendationsUnusedInstanceEnum{
	"in_use":            HostCpuRecommendationsUnusedInstanceInUse,
	"not_in_use":        HostCpuRecommendationsUnusedInstanceNotInUse,
	"is_not_determined": HostCpuRecommendationsUnusedInstanceIsNotDetermined,
}

// GetHostCpuRecommendationsUnusedInstanceEnumValues Enumerates the set of values for HostCpuRecommendationsUnusedInstanceEnum
func GetHostCpuRecommendationsUnusedInstanceEnumValues() []HostCpuRecommendationsUnusedInstanceEnum {
	values := make([]HostCpuRecommendationsUnusedInstanceEnum, 0)
	for _, v := range mappingHostCpuRecommendationsUnusedInstanceEnum {
		values = append(values, v)
	}
	return values
}

// GetHostCpuRecommendationsUnusedInstanceEnumStringValues Enumerates the set of values in String for HostCpuRecommendationsUnusedInstanceEnum
func GetHostCpuRecommendationsUnusedInstanceEnumStringValues() []string {
	return []string{
		"IN_USE",
		"NOT_IN_USE",
		"IS_NOT_DETERMINED",
	}
}

// GetMappingHostCpuRecommendationsUnusedInstanceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHostCpuRecommendationsUnusedInstanceEnum(val string) (HostCpuRecommendationsUnusedInstanceEnum, bool) {
	enum, ok := mappingHostCpuRecommendationsUnusedInstanceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
