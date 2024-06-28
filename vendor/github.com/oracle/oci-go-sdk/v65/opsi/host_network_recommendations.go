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

// HostNetworkRecommendations Contains network recommendation.
type HostNetworkRecommendations struct {

	// Identify if an instance is abandoned.
	IsAbandonedInstance *bool `mandatory:"false" json:"isAbandonedInstance"`

	// Identify unused instances based on cpu, memory and network metrics.
	UnusedInstance HostNetworkRecommendationsUnusedInstanceEnum `mandatory:"false" json:"unusedInstance,omitempty"`
}

func (m HostNetworkRecommendations) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HostNetworkRecommendations) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingHostNetworkRecommendationsUnusedInstanceEnum(string(m.UnusedInstance)); !ok && m.UnusedInstance != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UnusedInstance: %s. Supported values are: %s.", m.UnusedInstance, strings.Join(GetHostNetworkRecommendationsUnusedInstanceEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m HostNetworkRecommendations) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeHostNetworkRecommendations HostNetworkRecommendations
	s := struct {
		DiscriminatorParam string `json:"metricRecommendationName"`
		MarshalTypeHostNetworkRecommendations
	}{
		"HOST_NETWORK_RECOMMENDATIONS",
		(MarshalTypeHostNetworkRecommendations)(m),
	}

	return json.Marshal(&s)
}

// HostNetworkRecommendationsUnusedInstanceEnum Enum with underlying type: string
type HostNetworkRecommendationsUnusedInstanceEnum string

// Set of constants representing the allowable values for HostNetworkRecommendationsUnusedInstanceEnum
const (
	HostNetworkRecommendationsUnusedInstanceInUse           HostNetworkRecommendationsUnusedInstanceEnum = "IN_USE"
	HostNetworkRecommendationsUnusedInstanceNotInUse        HostNetworkRecommendationsUnusedInstanceEnum = "NOT_IN_USE"
	HostNetworkRecommendationsUnusedInstanceIsNotDetermined HostNetworkRecommendationsUnusedInstanceEnum = "IS_NOT_DETERMINED"
)

var mappingHostNetworkRecommendationsUnusedInstanceEnum = map[string]HostNetworkRecommendationsUnusedInstanceEnum{
	"IN_USE":            HostNetworkRecommendationsUnusedInstanceInUse,
	"NOT_IN_USE":        HostNetworkRecommendationsUnusedInstanceNotInUse,
	"IS_NOT_DETERMINED": HostNetworkRecommendationsUnusedInstanceIsNotDetermined,
}

var mappingHostNetworkRecommendationsUnusedInstanceEnumLowerCase = map[string]HostNetworkRecommendationsUnusedInstanceEnum{
	"in_use":            HostNetworkRecommendationsUnusedInstanceInUse,
	"not_in_use":        HostNetworkRecommendationsUnusedInstanceNotInUse,
	"is_not_determined": HostNetworkRecommendationsUnusedInstanceIsNotDetermined,
}

// GetHostNetworkRecommendationsUnusedInstanceEnumValues Enumerates the set of values for HostNetworkRecommendationsUnusedInstanceEnum
func GetHostNetworkRecommendationsUnusedInstanceEnumValues() []HostNetworkRecommendationsUnusedInstanceEnum {
	values := make([]HostNetworkRecommendationsUnusedInstanceEnum, 0)
	for _, v := range mappingHostNetworkRecommendationsUnusedInstanceEnum {
		values = append(values, v)
	}
	return values
}

// GetHostNetworkRecommendationsUnusedInstanceEnumStringValues Enumerates the set of values in String for HostNetworkRecommendationsUnusedInstanceEnum
func GetHostNetworkRecommendationsUnusedInstanceEnumStringValues() []string {
	return []string{
		"IN_USE",
		"NOT_IN_USE",
		"IS_NOT_DETERMINED",
	}
}

// GetMappingHostNetworkRecommendationsUnusedInstanceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHostNetworkRecommendationsUnusedInstanceEnum(val string) (HostNetworkRecommendationsUnusedInstanceEnum, bool) {
	enum, ok := mappingHostNetworkRecommendationsUnusedInstanceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
