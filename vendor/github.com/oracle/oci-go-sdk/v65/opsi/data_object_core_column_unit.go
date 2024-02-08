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

// DataObjectCoreColumnUnit Unit details of a data object column of CORE unit category.
type DataObjectCoreColumnUnit struct {

	// Display name of the column's unit.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Core unit.
	Unit DataObjectCoreColumnUnitUnitEnum `mandatory:"false" json:"unit,omitempty"`
}

// GetDisplayName returns DisplayName
func (m DataObjectCoreColumnUnit) GetDisplayName() *string {
	return m.DisplayName
}

func (m DataObjectCoreColumnUnit) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataObjectCoreColumnUnit) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDataObjectCoreColumnUnitUnitEnum(string(m.Unit)); !ok && m.Unit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Unit: %s. Supported values are: %s.", m.Unit, strings.Join(GetDataObjectCoreColumnUnitUnitEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DataObjectCoreColumnUnit) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDataObjectCoreColumnUnit DataObjectCoreColumnUnit
	s := struct {
		DiscriminatorParam string `json:"unitCategory"`
		MarshalTypeDataObjectCoreColumnUnit
	}{
		"CORE",
		(MarshalTypeDataObjectCoreColumnUnit)(m),
	}

	return json.Marshal(&s)
}

// DataObjectCoreColumnUnitUnitEnum Enum with underlying type: string
type DataObjectCoreColumnUnitUnitEnum string

// Set of constants representing the allowable values for DataObjectCoreColumnUnitUnitEnum
const (
	DataObjectCoreColumnUnitUnitCore      DataObjectCoreColumnUnitUnitEnum = "CORE"
	DataObjectCoreColumnUnitUnitMilliCore DataObjectCoreColumnUnitUnitEnum = "MILLI_CORE"
)

var mappingDataObjectCoreColumnUnitUnitEnum = map[string]DataObjectCoreColumnUnitUnitEnum{
	"CORE":       DataObjectCoreColumnUnitUnitCore,
	"MILLI_CORE": DataObjectCoreColumnUnitUnitMilliCore,
}

var mappingDataObjectCoreColumnUnitUnitEnumLowerCase = map[string]DataObjectCoreColumnUnitUnitEnum{
	"core":       DataObjectCoreColumnUnitUnitCore,
	"milli_core": DataObjectCoreColumnUnitUnitMilliCore,
}

// GetDataObjectCoreColumnUnitUnitEnumValues Enumerates the set of values for DataObjectCoreColumnUnitUnitEnum
func GetDataObjectCoreColumnUnitUnitEnumValues() []DataObjectCoreColumnUnitUnitEnum {
	values := make([]DataObjectCoreColumnUnitUnitEnum, 0)
	for _, v := range mappingDataObjectCoreColumnUnitUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetDataObjectCoreColumnUnitUnitEnumStringValues Enumerates the set of values in String for DataObjectCoreColumnUnitUnitEnum
func GetDataObjectCoreColumnUnitUnitEnumStringValues() []string {
	return []string{
		"CORE",
		"MILLI_CORE",
	}
}

// GetMappingDataObjectCoreColumnUnitUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataObjectCoreColumnUnitUnitEnum(val string) (DataObjectCoreColumnUnitUnitEnum, bool) {
	enum, ok := mappingDataObjectCoreColumnUnitUnitEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
