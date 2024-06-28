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

// DataObjectTimeColumnUnit Unit details of a data object column of TIME unit category.
type DataObjectTimeColumnUnit struct {

	// Display name of the column's unit.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Time unit.
	Unit DataObjectTimeColumnUnitUnitEnum `mandatory:"false" json:"unit,omitempty"`
}

// GetDisplayName returns DisplayName
func (m DataObjectTimeColumnUnit) GetDisplayName() *string {
	return m.DisplayName
}

func (m DataObjectTimeColumnUnit) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataObjectTimeColumnUnit) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDataObjectTimeColumnUnitUnitEnum(string(m.Unit)); !ok && m.Unit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Unit: %s. Supported values are: %s.", m.Unit, strings.Join(GetDataObjectTimeColumnUnitUnitEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DataObjectTimeColumnUnit) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDataObjectTimeColumnUnit DataObjectTimeColumnUnit
	s := struct {
		DiscriminatorParam string `json:"unitCategory"`
		MarshalTypeDataObjectTimeColumnUnit
	}{
		"TIME",
		(MarshalTypeDataObjectTimeColumnUnit)(m),
	}

	return json.Marshal(&s)
}

// DataObjectTimeColumnUnitUnitEnum Enum with underlying type: string
type DataObjectTimeColumnUnitUnitEnum string

// Set of constants representing the allowable values for DataObjectTimeColumnUnitUnitEnum
const (
	DataObjectTimeColumnUnitUnitNanoSecond  DataObjectTimeColumnUnitUnitEnum = "NANO_SECOND"
	DataObjectTimeColumnUnitUnitMicroSecond DataObjectTimeColumnUnitUnitEnum = "MICRO_SECOND"
	DataObjectTimeColumnUnitUnitMilliSecond DataObjectTimeColumnUnitUnitEnum = "MILLI_SECOND"
	DataObjectTimeColumnUnitUnitCentiSecond DataObjectTimeColumnUnitUnitEnum = "CENTI_SECOND"
	DataObjectTimeColumnUnitUnitSecond      DataObjectTimeColumnUnitUnitEnum = "SECOND"
	DataObjectTimeColumnUnitUnitHour        DataObjectTimeColumnUnitUnitEnum = "HOUR"
	DataObjectTimeColumnUnitUnitDay         DataObjectTimeColumnUnitUnitEnum = "DAY"
	DataObjectTimeColumnUnitUnitWeek        DataObjectTimeColumnUnitUnitEnum = "WEEK"
	DataObjectTimeColumnUnitUnitMonth       DataObjectTimeColumnUnitUnitEnum = "MONTH"
	DataObjectTimeColumnUnitUnitYear        DataObjectTimeColumnUnitUnitEnum = "YEAR"
	DataObjectTimeColumnUnitUnitMinute      DataObjectTimeColumnUnitUnitEnum = "MINUTE"
)

var mappingDataObjectTimeColumnUnitUnitEnum = map[string]DataObjectTimeColumnUnitUnitEnum{
	"NANO_SECOND":  DataObjectTimeColumnUnitUnitNanoSecond,
	"MICRO_SECOND": DataObjectTimeColumnUnitUnitMicroSecond,
	"MILLI_SECOND": DataObjectTimeColumnUnitUnitMilliSecond,
	"CENTI_SECOND": DataObjectTimeColumnUnitUnitCentiSecond,
	"SECOND":       DataObjectTimeColumnUnitUnitSecond,
	"HOUR":         DataObjectTimeColumnUnitUnitHour,
	"DAY":          DataObjectTimeColumnUnitUnitDay,
	"WEEK":         DataObjectTimeColumnUnitUnitWeek,
	"MONTH":        DataObjectTimeColumnUnitUnitMonth,
	"YEAR":         DataObjectTimeColumnUnitUnitYear,
	"MINUTE":       DataObjectTimeColumnUnitUnitMinute,
}

var mappingDataObjectTimeColumnUnitUnitEnumLowerCase = map[string]DataObjectTimeColumnUnitUnitEnum{
	"nano_second":  DataObjectTimeColumnUnitUnitNanoSecond,
	"micro_second": DataObjectTimeColumnUnitUnitMicroSecond,
	"milli_second": DataObjectTimeColumnUnitUnitMilliSecond,
	"centi_second": DataObjectTimeColumnUnitUnitCentiSecond,
	"second":       DataObjectTimeColumnUnitUnitSecond,
	"hour":         DataObjectTimeColumnUnitUnitHour,
	"day":          DataObjectTimeColumnUnitUnitDay,
	"week":         DataObjectTimeColumnUnitUnitWeek,
	"month":        DataObjectTimeColumnUnitUnitMonth,
	"year":         DataObjectTimeColumnUnitUnitYear,
	"minute":       DataObjectTimeColumnUnitUnitMinute,
}

// GetDataObjectTimeColumnUnitUnitEnumValues Enumerates the set of values for DataObjectTimeColumnUnitUnitEnum
func GetDataObjectTimeColumnUnitUnitEnumValues() []DataObjectTimeColumnUnitUnitEnum {
	values := make([]DataObjectTimeColumnUnitUnitEnum, 0)
	for _, v := range mappingDataObjectTimeColumnUnitUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetDataObjectTimeColumnUnitUnitEnumStringValues Enumerates the set of values in String for DataObjectTimeColumnUnitUnitEnum
func GetDataObjectTimeColumnUnitUnitEnumStringValues() []string {
	return []string{
		"NANO_SECOND",
		"MICRO_SECOND",
		"MILLI_SECOND",
		"CENTI_SECOND",
		"SECOND",
		"HOUR",
		"DAY",
		"WEEK",
		"MONTH",
		"YEAR",
		"MINUTE",
	}
}

// GetMappingDataObjectTimeColumnUnitUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataObjectTimeColumnUnitUnitEnum(val string) (DataObjectTimeColumnUnitUnitEnum, bool) {
	enum, ok := mappingDataObjectTimeColumnUnitUnitEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
