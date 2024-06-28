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

// DataObjectPowerColumnUnit Unit details of a data object column of POWER unit category.
type DataObjectPowerColumnUnit struct {

	// Display name of the column's unit.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Power unit.
	Unit DataObjectPowerColumnUnitUnitEnum `mandatory:"false" json:"unit,omitempty"`
}

// GetDisplayName returns DisplayName
func (m DataObjectPowerColumnUnit) GetDisplayName() *string {
	return m.DisplayName
}

func (m DataObjectPowerColumnUnit) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataObjectPowerColumnUnit) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDataObjectPowerColumnUnitUnitEnum(string(m.Unit)); !ok && m.Unit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Unit: %s. Supported values are: %s.", m.Unit, strings.Join(GetDataObjectPowerColumnUnitUnitEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DataObjectPowerColumnUnit) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDataObjectPowerColumnUnit DataObjectPowerColumnUnit
	s := struct {
		DiscriminatorParam string `json:"unitCategory"`
		MarshalTypeDataObjectPowerColumnUnit
	}{
		"POWER",
		(MarshalTypeDataObjectPowerColumnUnit)(m),
	}

	return json.Marshal(&s)
}

// DataObjectPowerColumnUnitUnitEnum Enum with underlying type: string
type DataObjectPowerColumnUnitUnitEnum string

// Set of constants representing the allowable values for DataObjectPowerColumnUnitUnitEnum
const (
	DataObjectPowerColumnUnitUnitAmp      DataObjectPowerColumnUnitUnitEnum = "AMP"
	DataObjectPowerColumnUnitUnitWatt     DataObjectPowerColumnUnitUnitEnum = "WATT"
	DataObjectPowerColumnUnitUnitKiloWatt DataObjectPowerColumnUnitUnitEnum = "KILO_WATT"
	DataObjectPowerColumnUnitUnitMegaWatt DataObjectPowerColumnUnitUnitEnum = "MEGA_WATT"
	DataObjectPowerColumnUnitUnitGigaWatt DataObjectPowerColumnUnitUnitEnum = "GIGA_WATT"
)

var mappingDataObjectPowerColumnUnitUnitEnum = map[string]DataObjectPowerColumnUnitUnitEnum{
	"AMP":       DataObjectPowerColumnUnitUnitAmp,
	"WATT":      DataObjectPowerColumnUnitUnitWatt,
	"KILO_WATT": DataObjectPowerColumnUnitUnitKiloWatt,
	"MEGA_WATT": DataObjectPowerColumnUnitUnitMegaWatt,
	"GIGA_WATT": DataObjectPowerColumnUnitUnitGigaWatt,
}

var mappingDataObjectPowerColumnUnitUnitEnumLowerCase = map[string]DataObjectPowerColumnUnitUnitEnum{
	"amp":       DataObjectPowerColumnUnitUnitAmp,
	"watt":      DataObjectPowerColumnUnitUnitWatt,
	"kilo_watt": DataObjectPowerColumnUnitUnitKiloWatt,
	"mega_watt": DataObjectPowerColumnUnitUnitMegaWatt,
	"giga_watt": DataObjectPowerColumnUnitUnitGigaWatt,
}

// GetDataObjectPowerColumnUnitUnitEnumValues Enumerates the set of values for DataObjectPowerColumnUnitUnitEnum
func GetDataObjectPowerColumnUnitUnitEnumValues() []DataObjectPowerColumnUnitUnitEnum {
	values := make([]DataObjectPowerColumnUnitUnitEnum, 0)
	for _, v := range mappingDataObjectPowerColumnUnitUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetDataObjectPowerColumnUnitUnitEnumStringValues Enumerates the set of values in String for DataObjectPowerColumnUnitUnitEnum
func GetDataObjectPowerColumnUnitUnitEnumStringValues() []string {
	return []string{
		"AMP",
		"WATT",
		"KILO_WATT",
		"MEGA_WATT",
		"GIGA_WATT",
	}
}

// GetMappingDataObjectPowerColumnUnitUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataObjectPowerColumnUnitUnitEnum(val string) (DataObjectPowerColumnUnitUnitEnum, bool) {
	enum, ok := mappingDataObjectPowerColumnUnitUnitEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
