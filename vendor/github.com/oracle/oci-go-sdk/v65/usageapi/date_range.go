// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Usage API
//
// Use the Usage API to view your Oracle Cloud usage and costs. The API allows you to request data that meets the specified filter criteria, and to group that data by the dimension of your choosing. The Usage API is used by the Cost Analysis tool in the Console. Also see Using the Usage API (https://docs.cloud.oracle.com/Content/Billing/Concepts/costanalysisoverview.htm#cost_analysis_using_the_api) for more information.
//

package usageapi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DateRange Static or dynamic date range `dateRangeType`,
// which corresponds with type-specific characteristics.
type DateRange interface {
}

type daterange struct {
	JsonData      []byte
	DateRangeType string `json:"dateRangeType"`
}

// UnmarshalJSON unmarshals json
func (m *daterange) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdaterange daterange
	s := struct {
		Model Unmarshalerdaterange
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DateRangeType = s.Model.DateRangeType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *daterange) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.DateRangeType {
	case "STATIC":
		mm := StaticDateRange{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DYNAMIC":
		mm := DynamicDateRange{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DateRange: %s.", m.DateRangeType)
		return *m, nil
	}
}

func (m daterange) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m daterange) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DateRangeDateRangeTypeEnum Enum with underlying type: string
type DateRangeDateRangeTypeEnum string

// Set of constants representing the allowable values for DateRangeDateRangeTypeEnum
const (
	DateRangeDateRangeTypeStatic  DateRangeDateRangeTypeEnum = "STATIC"
	DateRangeDateRangeTypeDynamic DateRangeDateRangeTypeEnum = "DYNAMIC"
)

var mappingDateRangeDateRangeTypeEnum = map[string]DateRangeDateRangeTypeEnum{
	"STATIC":  DateRangeDateRangeTypeStatic,
	"DYNAMIC": DateRangeDateRangeTypeDynamic,
}

var mappingDateRangeDateRangeTypeEnumLowerCase = map[string]DateRangeDateRangeTypeEnum{
	"static":  DateRangeDateRangeTypeStatic,
	"dynamic": DateRangeDateRangeTypeDynamic,
}

// GetDateRangeDateRangeTypeEnumValues Enumerates the set of values for DateRangeDateRangeTypeEnum
func GetDateRangeDateRangeTypeEnumValues() []DateRangeDateRangeTypeEnum {
	values := make([]DateRangeDateRangeTypeEnum, 0)
	for _, v := range mappingDateRangeDateRangeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDateRangeDateRangeTypeEnumStringValues Enumerates the set of values in String for DateRangeDateRangeTypeEnum
func GetDateRangeDateRangeTypeEnumStringValues() []string {
	return []string{
		"STATIC",
		"DYNAMIC",
	}
}

// GetMappingDateRangeDateRangeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDateRangeDateRangeTypeEnum(val string) (DateRangeDateRangeTypeEnum, bool) {
	enum, ok := mappingDateRangeDateRangeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
