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

// DataObjectFrequencyColumnUnit Unit details of a data object column of FREQEUENCY unit category.
type DataObjectFrequencyColumnUnit struct {

	// Display name of the column's unit.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Frequency unit.
	Unit DataObjectFrequencyColumnUnitUnitEnum `mandatory:"false" json:"unit,omitempty"`
}

// GetDisplayName returns DisplayName
func (m DataObjectFrequencyColumnUnit) GetDisplayName() *string {
	return m.DisplayName
}

func (m DataObjectFrequencyColumnUnit) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataObjectFrequencyColumnUnit) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDataObjectFrequencyColumnUnitUnitEnum(string(m.Unit)); !ok && m.Unit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Unit: %s. Supported values are: %s.", m.Unit, strings.Join(GetDataObjectFrequencyColumnUnitUnitEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DataObjectFrequencyColumnUnit) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDataObjectFrequencyColumnUnit DataObjectFrequencyColumnUnit
	s := struct {
		DiscriminatorParam string `json:"unitCategory"`
		MarshalTypeDataObjectFrequencyColumnUnit
	}{
		"FREQUENCY",
		(MarshalTypeDataObjectFrequencyColumnUnit)(m),
	}

	return json.Marshal(&s)
}

// DataObjectFrequencyColumnUnitUnitEnum Enum with underlying type: string
type DataObjectFrequencyColumnUnitUnitEnum string

// Set of constants representing the allowable values for DataObjectFrequencyColumnUnitUnitEnum
const (
	DataObjectFrequencyColumnUnitUnitHertz     DataObjectFrequencyColumnUnitUnitEnum = "HERTZ"
	DataObjectFrequencyColumnUnitUnitKiloHertz DataObjectFrequencyColumnUnitUnitEnum = "KILO_HERTZ"
	DataObjectFrequencyColumnUnitUnitMegaHertz DataObjectFrequencyColumnUnitUnitEnum = "MEGA_HERTZ"
	DataObjectFrequencyColumnUnitUnitGigaHertz DataObjectFrequencyColumnUnitUnitEnum = "GIGA_HERTZ"
	DataObjectFrequencyColumnUnitUnitTeraHertz DataObjectFrequencyColumnUnitUnitEnum = "TERA_HERTZ"
)

var mappingDataObjectFrequencyColumnUnitUnitEnum = map[string]DataObjectFrequencyColumnUnitUnitEnum{
	"HERTZ":      DataObjectFrequencyColumnUnitUnitHertz,
	"KILO_HERTZ": DataObjectFrequencyColumnUnitUnitKiloHertz,
	"MEGA_HERTZ": DataObjectFrequencyColumnUnitUnitMegaHertz,
	"GIGA_HERTZ": DataObjectFrequencyColumnUnitUnitGigaHertz,
	"TERA_HERTZ": DataObjectFrequencyColumnUnitUnitTeraHertz,
}

var mappingDataObjectFrequencyColumnUnitUnitEnumLowerCase = map[string]DataObjectFrequencyColumnUnitUnitEnum{
	"hertz":      DataObjectFrequencyColumnUnitUnitHertz,
	"kilo_hertz": DataObjectFrequencyColumnUnitUnitKiloHertz,
	"mega_hertz": DataObjectFrequencyColumnUnitUnitMegaHertz,
	"giga_hertz": DataObjectFrequencyColumnUnitUnitGigaHertz,
	"tera_hertz": DataObjectFrequencyColumnUnitUnitTeraHertz,
}

// GetDataObjectFrequencyColumnUnitUnitEnumValues Enumerates the set of values for DataObjectFrequencyColumnUnitUnitEnum
func GetDataObjectFrequencyColumnUnitUnitEnumValues() []DataObjectFrequencyColumnUnitUnitEnum {
	values := make([]DataObjectFrequencyColumnUnitUnitEnum, 0)
	for _, v := range mappingDataObjectFrequencyColumnUnitUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetDataObjectFrequencyColumnUnitUnitEnumStringValues Enumerates the set of values in String for DataObjectFrequencyColumnUnitUnitEnum
func GetDataObjectFrequencyColumnUnitUnitEnumStringValues() []string {
	return []string{
		"HERTZ",
		"KILO_HERTZ",
		"MEGA_HERTZ",
		"GIGA_HERTZ",
		"TERA_HERTZ",
	}
}

// GetMappingDataObjectFrequencyColumnUnitUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataObjectFrequencyColumnUnitUnitEnum(val string) (DataObjectFrequencyColumnUnitUnitEnum, bool) {
	enum, ok := mappingDataObjectFrequencyColumnUnitUnitEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
