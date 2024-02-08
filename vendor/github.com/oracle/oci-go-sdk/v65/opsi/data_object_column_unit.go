// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DataObjectColumnUnit Unit details of a data object column.
type DataObjectColumnUnit interface {

	// Display name of the column's unit.
	GetDisplayName() *string
}

type dataobjectcolumnunit struct {
	JsonData     []byte
	DisplayName  *string `mandatory:"false" json:"displayName"`
	UnitCategory string  `json:"unitCategory"`
}

// UnmarshalJSON unmarshals json
func (m *dataobjectcolumnunit) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdataobjectcolumnunit dataobjectcolumnunit
	s := struct {
		Model Unmarshalerdataobjectcolumnunit
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayName = s.Model.DisplayName
	m.UnitCategory = s.Model.UnitCategory

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *dataobjectcolumnunit) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.UnitCategory {
	case "CORE":
		mm := DataObjectCoreColumnUnit{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TIME":
		mm := DataObjectTimeColumnUnit{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OTHER_STANDARD":
		mm := DataObjectOtherStandardColumnUnit{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CUSTOM":
		mm := DataObjectCustomColumnUnit{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TEMPERATURE":
		mm := DataObjectTemperatureColumnUnit{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "POWER":
		mm := DataObjectPowerColumnUnit{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RATE":
		mm := DataObjectRateColumnUnit{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FREQUENCY":
		mm := DataObjectFrequencyColumnUnit{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DATA_SIZE":
		mm := DataObjectDataSizeColumnUnit{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DataObjectColumnUnit: %s.", m.UnitCategory)
		return *m, nil
	}
}

// GetDisplayName returns DisplayName
func (m dataobjectcolumnunit) GetDisplayName() *string {
	return m.DisplayName
}

func (m dataobjectcolumnunit) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m dataobjectcolumnunit) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DataObjectColumnUnitUnitCategoryEnum Enum with underlying type: string
type DataObjectColumnUnitUnitCategoryEnum string

// Set of constants representing the allowable values for DataObjectColumnUnitUnitCategoryEnum
const (
	DataObjectColumnUnitUnitCategoryDataSize      DataObjectColumnUnitUnitCategoryEnum = "DATA_SIZE"
	DataObjectColumnUnitUnitCategoryTime          DataObjectColumnUnitUnitCategoryEnum = "TIME"
	DataObjectColumnUnitUnitCategoryPower         DataObjectColumnUnitUnitCategoryEnum = "POWER"
	DataObjectColumnUnitUnitCategoryTemperature   DataObjectColumnUnitUnitCategoryEnum = "TEMPERATURE"
	DataObjectColumnUnitUnitCategoryCore          DataObjectColumnUnitUnitCategoryEnum = "CORE"
	DataObjectColumnUnitUnitCategoryRate          DataObjectColumnUnitUnitCategoryEnum = "RATE"
	DataObjectColumnUnitUnitCategoryFrequency     DataObjectColumnUnitUnitCategoryEnum = "FREQUENCY"
	DataObjectColumnUnitUnitCategoryOtherStandard DataObjectColumnUnitUnitCategoryEnum = "OTHER_STANDARD"
	DataObjectColumnUnitUnitCategoryCustom        DataObjectColumnUnitUnitCategoryEnum = "CUSTOM"
)

var mappingDataObjectColumnUnitUnitCategoryEnum = map[string]DataObjectColumnUnitUnitCategoryEnum{
	"DATA_SIZE":      DataObjectColumnUnitUnitCategoryDataSize,
	"TIME":           DataObjectColumnUnitUnitCategoryTime,
	"POWER":          DataObjectColumnUnitUnitCategoryPower,
	"TEMPERATURE":    DataObjectColumnUnitUnitCategoryTemperature,
	"CORE":           DataObjectColumnUnitUnitCategoryCore,
	"RATE":           DataObjectColumnUnitUnitCategoryRate,
	"FREQUENCY":      DataObjectColumnUnitUnitCategoryFrequency,
	"OTHER_STANDARD": DataObjectColumnUnitUnitCategoryOtherStandard,
	"CUSTOM":         DataObjectColumnUnitUnitCategoryCustom,
}

var mappingDataObjectColumnUnitUnitCategoryEnumLowerCase = map[string]DataObjectColumnUnitUnitCategoryEnum{
	"data_size":      DataObjectColumnUnitUnitCategoryDataSize,
	"time":           DataObjectColumnUnitUnitCategoryTime,
	"power":          DataObjectColumnUnitUnitCategoryPower,
	"temperature":    DataObjectColumnUnitUnitCategoryTemperature,
	"core":           DataObjectColumnUnitUnitCategoryCore,
	"rate":           DataObjectColumnUnitUnitCategoryRate,
	"frequency":      DataObjectColumnUnitUnitCategoryFrequency,
	"other_standard": DataObjectColumnUnitUnitCategoryOtherStandard,
	"custom":         DataObjectColumnUnitUnitCategoryCustom,
}

// GetDataObjectColumnUnitUnitCategoryEnumValues Enumerates the set of values for DataObjectColumnUnitUnitCategoryEnum
func GetDataObjectColumnUnitUnitCategoryEnumValues() []DataObjectColumnUnitUnitCategoryEnum {
	values := make([]DataObjectColumnUnitUnitCategoryEnum, 0)
	for _, v := range mappingDataObjectColumnUnitUnitCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetDataObjectColumnUnitUnitCategoryEnumStringValues Enumerates the set of values in String for DataObjectColumnUnitUnitCategoryEnum
func GetDataObjectColumnUnitUnitCategoryEnumStringValues() []string {
	return []string{
		"DATA_SIZE",
		"TIME",
		"POWER",
		"TEMPERATURE",
		"CORE",
		"RATE",
		"FREQUENCY",
		"OTHER_STANDARD",
		"CUSTOM",
	}
}

// GetMappingDataObjectColumnUnitUnitCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataObjectColumnUnitUnitCategoryEnum(val string) (DataObjectColumnUnitUnitCategoryEnum, bool) {
	enum, ok := mappingDataObjectColumnUnitUnitCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
