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

// ResourceAssessmentStrategy Migration strategy for the resource to be migrated.
type ResourceAssessmentStrategy interface {

	// The type of resource.
	GetResourceType() ResourceAssessmentStrategyResourceTypeEnum
}

type resourceassessmentstrategy struct {
	JsonData     []byte
	ResourceType ResourceAssessmentStrategyResourceTypeEnum `mandatory:"true" json:"resourceType"`
	StrategyType string                                     `json:"strategyType"`
}

// UnmarshalJSON unmarshals json
func (m *resourceassessmentstrategy) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerresourceassessmentstrategy resourceassessmentstrategy
	s := struct {
		Model Unmarshalerresourceassessmentstrategy
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ResourceType = s.Model.ResourceType
	m.StrategyType = s.Model.StrategyType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *resourceassessmentstrategy) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.StrategyType {
	case "PEAK":
		mm := PeakResourceAssessmentStrategy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PERCENTILE":
		mm := PercentileResourceAssessmentStrategy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AVERAGE":
		mm := AverageResourceAssessmentStrategy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AS_IS":
		mm := AsIsResourceAssessmentStrategy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ResourceAssessmentStrategy: %s.", m.StrategyType)
		return *m, nil
	}
}

// GetResourceType returns ResourceType
func (m resourceassessmentstrategy) GetResourceType() ResourceAssessmentStrategyResourceTypeEnum {
	return m.ResourceType
}

func (m resourceassessmentstrategy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m resourceassessmentstrategy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingResourceAssessmentStrategyResourceTypeEnum(string(m.ResourceType)); !ok && m.ResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceType: %s. Supported values are: %s.", m.ResourceType, strings.Join(GetResourceAssessmentStrategyResourceTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ResourceAssessmentStrategyResourceTypeEnum Enum with underlying type: string
type ResourceAssessmentStrategyResourceTypeEnum string

// Set of constants representing the allowable values for ResourceAssessmentStrategyResourceTypeEnum
const (
	ResourceAssessmentStrategyResourceTypeCpu    ResourceAssessmentStrategyResourceTypeEnum = "CPU"
	ResourceAssessmentStrategyResourceTypeMemory ResourceAssessmentStrategyResourceTypeEnum = "MEMORY"
	ResourceAssessmentStrategyResourceTypeAll    ResourceAssessmentStrategyResourceTypeEnum = "ALL"
)

var mappingResourceAssessmentStrategyResourceTypeEnum = map[string]ResourceAssessmentStrategyResourceTypeEnum{
	"CPU":    ResourceAssessmentStrategyResourceTypeCpu,
	"MEMORY": ResourceAssessmentStrategyResourceTypeMemory,
	"ALL":    ResourceAssessmentStrategyResourceTypeAll,
}

var mappingResourceAssessmentStrategyResourceTypeEnumLowerCase = map[string]ResourceAssessmentStrategyResourceTypeEnum{
	"cpu":    ResourceAssessmentStrategyResourceTypeCpu,
	"memory": ResourceAssessmentStrategyResourceTypeMemory,
	"all":    ResourceAssessmentStrategyResourceTypeAll,
}

// GetResourceAssessmentStrategyResourceTypeEnumValues Enumerates the set of values for ResourceAssessmentStrategyResourceTypeEnum
func GetResourceAssessmentStrategyResourceTypeEnumValues() []ResourceAssessmentStrategyResourceTypeEnum {
	values := make([]ResourceAssessmentStrategyResourceTypeEnum, 0)
	for _, v := range mappingResourceAssessmentStrategyResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetResourceAssessmentStrategyResourceTypeEnumStringValues Enumerates the set of values in String for ResourceAssessmentStrategyResourceTypeEnum
func GetResourceAssessmentStrategyResourceTypeEnumStringValues() []string {
	return []string{
		"CPU",
		"MEMORY",
		"ALL",
	}
}

// GetMappingResourceAssessmentStrategyResourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResourceAssessmentStrategyResourceTypeEnum(val string) (ResourceAssessmentStrategyResourceTypeEnum, bool) {
	enum, ok := mappingResourceAssessmentStrategyResourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ResourceAssessmentStrategyStrategyTypeEnum Enum with underlying type: string
type ResourceAssessmentStrategyStrategyTypeEnum string

// Set of constants representing the allowable values for ResourceAssessmentStrategyStrategyTypeEnum
const (
	ResourceAssessmentStrategyStrategyTypeAsIs       ResourceAssessmentStrategyStrategyTypeEnum = "AS_IS"
	ResourceAssessmentStrategyStrategyTypeAverage    ResourceAssessmentStrategyStrategyTypeEnum = "AVERAGE"
	ResourceAssessmentStrategyStrategyTypePeak       ResourceAssessmentStrategyStrategyTypeEnum = "PEAK"
	ResourceAssessmentStrategyStrategyTypePercentile ResourceAssessmentStrategyStrategyTypeEnum = "PERCENTILE"
)

var mappingResourceAssessmentStrategyStrategyTypeEnum = map[string]ResourceAssessmentStrategyStrategyTypeEnum{
	"AS_IS":      ResourceAssessmentStrategyStrategyTypeAsIs,
	"AVERAGE":    ResourceAssessmentStrategyStrategyTypeAverage,
	"PEAK":       ResourceAssessmentStrategyStrategyTypePeak,
	"PERCENTILE": ResourceAssessmentStrategyStrategyTypePercentile,
}

var mappingResourceAssessmentStrategyStrategyTypeEnumLowerCase = map[string]ResourceAssessmentStrategyStrategyTypeEnum{
	"as_is":      ResourceAssessmentStrategyStrategyTypeAsIs,
	"average":    ResourceAssessmentStrategyStrategyTypeAverage,
	"peak":       ResourceAssessmentStrategyStrategyTypePeak,
	"percentile": ResourceAssessmentStrategyStrategyTypePercentile,
}

// GetResourceAssessmentStrategyStrategyTypeEnumValues Enumerates the set of values for ResourceAssessmentStrategyStrategyTypeEnum
func GetResourceAssessmentStrategyStrategyTypeEnumValues() []ResourceAssessmentStrategyStrategyTypeEnum {
	values := make([]ResourceAssessmentStrategyStrategyTypeEnum, 0)
	for _, v := range mappingResourceAssessmentStrategyStrategyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetResourceAssessmentStrategyStrategyTypeEnumStringValues Enumerates the set of values in String for ResourceAssessmentStrategyStrategyTypeEnum
func GetResourceAssessmentStrategyStrategyTypeEnumStringValues() []string {
	return []string{
		"AS_IS",
		"AVERAGE",
		"PEAK",
		"PERCENTILE",
	}
}

// GetMappingResourceAssessmentStrategyStrategyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResourceAssessmentStrategyStrategyTypeEnum(val string) (ResourceAssessmentStrategyStrategyTypeEnum, bool) {
	enum, ok := mappingResourceAssessmentStrategyStrategyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
