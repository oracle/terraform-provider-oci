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

// DataObjectTemperatureColumnUnit Unit details of a data object column of TEMPERATURE unit category.
type DataObjectTemperatureColumnUnit struct {

	// Display name of the column's unit.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Temparature unit.
	Unit DataObjectTemperatureColumnUnitUnitEnum `mandatory:"false" json:"unit,omitempty"`
}

// GetDisplayName returns DisplayName
func (m DataObjectTemperatureColumnUnit) GetDisplayName() *string {
	return m.DisplayName
}

func (m DataObjectTemperatureColumnUnit) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataObjectTemperatureColumnUnit) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDataObjectTemperatureColumnUnitUnitEnum(string(m.Unit)); !ok && m.Unit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Unit: %s. Supported values are: %s.", m.Unit, strings.Join(GetDataObjectTemperatureColumnUnitUnitEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DataObjectTemperatureColumnUnit) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDataObjectTemperatureColumnUnit DataObjectTemperatureColumnUnit
	s := struct {
		DiscriminatorParam string `json:"unitCategory"`
		MarshalTypeDataObjectTemperatureColumnUnit
	}{
		"TEMPERATURE",
		(MarshalTypeDataObjectTemperatureColumnUnit)(m),
	}

	return json.Marshal(&s)
}

// DataObjectTemperatureColumnUnitUnitEnum Enum with underlying type: string
type DataObjectTemperatureColumnUnitUnitEnum string

// Set of constants representing the allowable values for DataObjectTemperatureColumnUnitUnitEnum
const (
	DataObjectTemperatureColumnUnitUnitCelsius    DataObjectTemperatureColumnUnitUnitEnum = "CELSIUS"
	DataObjectTemperatureColumnUnitUnitFahrenheit DataObjectTemperatureColumnUnitUnitEnum = "FAHRENHEIT"
)

var mappingDataObjectTemperatureColumnUnitUnitEnum = map[string]DataObjectTemperatureColumnUnitUnitEnum{
	"CELSIUS":    DataObjectTemperatureColumnUnitUnitCelsius,
	"FAHRENHEIT": DataObjectTemperatureColumnUnitUnitFahrenheit,
}

var mappingDataObjectTemperatureColumnUnitUnitEnumLowerCase = map[string]DataObjectTemperatureColumnUnitUnitEnum{
	"celsius":    DataObjectTemperatureColumnUnitUnitCelsius,
	"fahrenheit": DataObjectTemperatureColumnUnitUnitFahrenheit,
}

// GetDataObjectTemperatureColumnUnitUnitEnumValues Enumerates the set of values for DataObjectTemperatureColumnUnitUnitEnum
func GetDataObjectTemperatureColumnUnitUnitEnumValues() []DataObjectTemperatureColumnUnitUnitEnum {
	values := make([]DataObjectTemperatureColumnUnitUnitEnum, 0)
	for _, v := range mappingDataObjectTemperatureColumnUnitUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetDataObjectTemperatureColumnUnitUnitEnumStringValues Enumerates the set of values in String for DataObjectTemperatureColumnUnitUnitEnum
func GetDataObjectTemperatureColumnUnitUnitEnumStringValues() []string {
	return []string{
		"CELSIUS",
		"FAHRENHEIT",
	}
}

// GetMappingDataObjectTemperatureColumnUnitUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataObjectTemperatureColumnUnitUnitEnum(val string) (DataObjectTemperatureColumnUnitUnitEnum, bool) {
	enum, ok := mappingDataObjectTemperatureColumnUnitUnitEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
