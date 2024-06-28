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

// DataObjectDataSizeColumnUnit Unit details of a data object column of DATA_SIZE unit category.
type DataObjectDataSizeColumnUnit struct {

	// Display name of the column's unit.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Data size unit.
	Unit DataObjectDataSizeColumnUnitUnitEnum `mandatory:"false" json:"unit,omitempty"`
}

// GetDisplayName returns DisplayName
func (m DataObjectDataSizeColumnUnit) GetDisplayName() *string {
	return m.DisplayName
}

func (m DataObjectDataSizeColumnUnit) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataObjectDataSizeColumnUnit) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDataObjectDataSizeColumnUnitUnitEnum(string(m.Unit)); !ok && m.Unit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Unit: %s. Supported values are: %s.", m.Unit, strings.Join(GetDataObjectDataSizeColumnUnitUnitEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DataObjectDataSizeColumnUnit) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDataObjectDataSizeColumnUnit DataObjectDataSizeColumnUnit
	s := struct {
		DiscriminatorParam string `json:"unitCategory"`
		MarshalTypeDataObjectDataSizeColumnUnit
	}{
		"DATA_SIZE",
		(MarshalTypeDataObjectDataSizeColumnUnit)(m),
	}

	return json.Marshal(&s)
}

// DataObjectDataSizeColumnUnitUnitEnum Enum with underlying type: string
type DataObjectDataSizeColumnUnitUnitEnum string

// Set of constants representing the allowable values for DataObjectDataSizeColumnUnitUnitEnum
const (
	DataObjectDataSizeColumnUnitUnitCharacter DataObjectDataSizeColumnUnitUnitEnum = "CHARACTER"
	DataObjectDataSizeColumnUnitUnitBlock     DataObjectDataSizeColumnUnitUnitEnum = "BLOCK"
	DataObjectDataSizeColumnUnitUnitBit       DataObjectDataSizeColumnUnitUnitEnum = "BIT"
	DataObjectDataSizeColumnUnitUnitByte      DataObjectDataSizeColumnUnitUnitEnum = "BYTE"
	DataObjectDataSizeColumnUnitUnitKiloByte  DataObjectDataSizeColumnUnitUnitEnum = "KILO_BYTE"
	DataObjectDataSizeColumnUnitUnitMegaByte  DataObjectDataSizeColumnUnitUnitEnum = "MEGA_BYTE"
	DataObjectDataSizeColumnUnitUnitGigaByte  DataObjectDataSizeColumnUnitUnitEnum = "GIGA_BYTE"
	DataObjectDataSizeColumnUnitUnitTeraByte  DataObjectDataSizeColumnUnitUnitEnum = "TERA_BYTE"
	DataObjectDataSizeColumnUnitUnitPetaByte  DataObjectDataSizeColumnUnitUnitEnum = "PETA_BYTE"
	DataObjectDataSizeColumnUnitUnitExaByte   DataObjectDataSizeColumnUnitUnitEnum = "EXA_BYTE"
	DataObjectDataSizeColumnUnitUnitZettaByte DataObjectDataSizeColumnUnitUnitEnum = "ZETTA_BYTE"
	DataObjectDataSizeColumnUnitUnitYottaByte DataObjectDataSizeColumnUnitUnitEnum = "YOTTA_BYTE"
)

var mappingDataObjectDataSizeColumnUnitUnitEnum = map[string]DataObjectDataSizeColumnUnitUnitEnum{
	"CHARACTER":  DataObjectDataSizeColumnUnitUnitCharacter,
	"BLOCK":      DataObjectDataSizeColumnUnitUnitBlock,
	"BIT":        DataObjectDataSizeColumnUnitUnitBit,
	"BYTE":       DataObjectDataSizeColumnUnitUnitByte,
	"KILO_BYTE":  DataObjectDataSizeColumnUnitUnitKiloByte,
	"MEGA_BYTE":  DataObjectDataSizeColumnUnitUnitMegaByte,
	"GIGA_BYTE":  DataObjectDataSizeColumnUnitUnitGigaByte,
	"TERA_BYTE":  DataObjectDataSizeColumnUnitUnitTeraByte,
	"PETA_BYTE":  DataObjectDataSizeColumnUnitUnitPetaByte,
	"EXA_BYTE":   DataObjectDataSizeColumnUnitUnitExaByte,
	"ZETTA_BYTE": DataObjectDataSizeColumnUnitUnitZettaByte,
	"YOTTA_BYTE": DataObjectDataSizeColumnUnitUnitYottaByte,
}

var mappingDataObjectDataSizeColumnUnitUnitEnumLowerCase = map[string]DataObjectDataSizeColumnUnitUnitEnum{
	"character":  DataObjectDataSizeColumnUnitUnitCharacter,
	"block":      DataObjectDataSizeColumnUnitUnitBlock,
	"bit":        DataObjectDataSizeColumnUnitUnitBit,
	"byte":       DataObjectDataSizeColumnUnitUnitByte,
	"kilo_byte":  DataObjectDataSizeColumnUnitUnitKiloByte,
	"mega_byte":  DataObjectDataSizeColumnUnitUnitMegaByte,
	"giga_byte":  DataObjectDataSizeColumnUnitUnitGigaByte,
	"tera_byte":  DataObjectDataSizeColumnUnitUnitTeraByte,
	"peta_byte":  DataObjectDataSizeColumnUnitUnitPetaByte,
	"exa_byte":   DataObjectDataSizeColumnUnitUnitExaByte,
	"zetta_byte": DataObjectDataSizeColumnUnitUnitZettaByte,
	"yotta_byte": DataObjectDataSizeColumnUnitUnitYottaByte,
}

// GetDataObjectDataSizeColumnUnitUnitEnumValues Enumerates the set of values for DataObjectDataSizeColumnUnitUnitEnum
func GetDataObjectDataSizeColumnUnitUnitEnumValues() []DataObjectDataSizeColumnUnitUnitEnum {
	values := make([]DataObjectDataSizeColumnUnitUnitEnum, 0)
	for _, v := range mappingDataObjectDataSizeColumnUnitUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetDataObjectDataSizeColumnUnitUnitEnumStringValues Enumerates the set of values in String for DataObjectDataSizeColumnUnitUnitEnum
func GetDataObjectDataSizeColumnUnitUnitEnumStringValues() []string {
	return []string{
		"CHARACTER",
		"BLOCK",
		"BIT",
		"BYTE",
		"KILO_BYTE",
		"MEGA_BYTE",
		"GIGA_BYTE",
		"TERA_BYTE",
		"PETA_BYTE",
		"EXA_BYTE",
		"ZETTA_BYTE",
		"YOTTA_BYTE",
	}
}

// GetMappingDataObjectDataSizeColumnUnitUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataObjectDataSizeColumnUnitUnitEnum(val string) (DataObjectDataSizeColumnUnitUnitEnum, bool) {
	enum, ok := mappingDataObjectDataSizeColumnUnitUnitEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
