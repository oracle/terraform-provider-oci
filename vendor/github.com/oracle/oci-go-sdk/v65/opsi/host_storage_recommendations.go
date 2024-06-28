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

// HostStorageRecommendations Contains storage recommendation.
type HostStorageRecommendations struct {

	// Identify if an instance is abandoned.
	IsAbandonedInstance *bool `mandatory:"false" json:"isAbandonedInstance"`

	// Identify unused instances based on cpu, memory and network metrics.
	UnusedInstance HostStorageRecommendationsUnusedInstanceEnum `mandatory:"false" json:"unusedInstance,omitempty"`
}

func (m HostStorageRecommendations) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HostStorageRecommendations) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingHostStorageRecommendationsUnusedInstanceEnum(string(m.UnusedInstance)); !ok && m.UnusedInstance != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UnusedInstance: %s. Supported values are: %s.", m.UnusedInstance, strings.Join(GetHostStorageRecommendationsUnusedInstanceEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m HostStorageRecommendations) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeHostStorageRecommendations HostStorageRecommendations
	s := struct {
		DiscriminatorParam string `json:"metricRecommendationName"`
		MarshalTypeHostStorageRecommendations
	}{
		"HOST_STORAGE_RECOMMENDATIONS",
		(MarshalTypeHostStorageRecommendations)(m),
	}

	return json.Marshal(&s)
}

// HostStorageRecommendationsUnusedInstanceEnum Enum with underlying type: string
type HostStorageRecommendationsUnusedInstanceEnum string

// Set of constants representing the allowable values for HostStorageRecommendationsUnusedInstanceEnum
const (
	HostStorageRecommendationsUnusedInstanceInUse           HostStorageRecommendationsUnusedInstanceEnum = "IN_USE"
	HostStorageRecommendationsUnusedInstanceNotInUse        HostStorageRecommendationsUnusedInstanceEnum = "NOT_IN_USE"
	HostStorageRecommendationsUnusedInstanceIsNotDetermined HostStorageRecommendationsUnusedInstanceEnum = "IS_NOT_DETERMINED"
)

var mappingHostStorageRecommendationsUnusedInstanceEnum = map[string]HostStorageRecommendationsUnusedInstanceEnum{
	"IN_USE":            HostStorageRecommendationsUnusedInstanceInUse,
	"NOT_IN_USE":        HostStorageRecommendationsUnusedInstanceNotInUse,
	"IS_NOT_DETERMINED": HostStorageRecommendationsUnusedInstanceIsNotDetermined,
}

var mappingHostStorageRecommendationsUnusedInstanceEnumLowerCase = map[string]HostStorageRecommendationsUnusedInstanceEnum{
	"in_use":            HostStorageRecommendationsUnusedInstanceInUse,
	"not_in_use":        HostStorageRecommendationsUnusedInstanceNotInUse,
	"is_not_determined": HostStorageRecommendationsUnusedInstanceIsNotDetermined,
}

// GetHostStorageRecommendationsUnusedInstanceEnumValues Enumerates the set of values for HostStorageRecommendationsUnusedInstanceEnum
func GetHostStorageRecommendationsUnusedInstanceEnumValues() []HostStorageRecommendationsUnusedInstanceEnum {
	values := make([]HostStorageRecommendationsUnusedInstanceEnum, 0)
	for _, v := range mappingHostStorageRecommendationsUnusedInstanceEnum {
		values = append(values, v)
	}
	return values
}

// GetHostStorageRecommendationsUnusedInstanceEnumStringValues Enumerates the set of values in String for HostStorageRecommendationsUnusedInstanceEnum
func GetHostStorageRecommendationsUnusedInstanceEnumStringValues() []string {
	return []string{
		"IN_USE",
		"NOT_IN_USE",
		"IS_NOT_DETERMINED",
	}
}

// GetMappingHostStorageRecommendationsUnusedInstanceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHostStorageRecommendationsUnusedInstanceEnum(val string) (HostStorageRecommendationsUnusedInstanceEnum, bool) {
	enum, ok := mappingHostStorageRecommendationsUnusedInstanceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
