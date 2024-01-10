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

// DataObjectOtherStandardColumnUnit Unit details of a data object column of OTHER_STANDARD unit category.
type DataObjectOtherStandardColumnUnit struct {

	// Display name of the column's unit.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Other standard column unit.
	Unit DataObjectOtherStandardColumnUnitUnitEnum `mandatory:"false" json:"unit,omitempty"`
}

// GetDisplayName returns DisplayName
func (m DataObjectOtherStandardColumnUnit) GetDisplayName() *string {
	return m.DisplayName
}

func (m DataObjectOtherStandardColumnUnit) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataObjectOtherStandardColumnUnit) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDataObjectOtherStandardColumnUnitUnitEnum(string(m.Unit)); !ok && m.Unit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Unit: %s. Supported values are: %s.", m.Unit, strings.Join(GetDataObjectOtherStandardColumnUnitUnitEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DataObjectOtherStandardColumnUnit) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDataObjectOtherStandardColumnUnit DataObjectOtherStandardColumnUnit
	s := struct {
		DiscriminatorParam string `json:"unitCategory"`
		MarshalTypeDataObjectOtherStandardColumnUnit
	}{
		"OTHER_STANDARD",
		(MarshalTypeDataObjectOtherStandardColumnUnit)(m),
	}

	return json.Marshal(&s)
}

// DataObjectOtherStandardColumnUnitUnitEnum Enum with underlying type: string
type DataObjectOtherStandardColumnUnitUnitEnum string

// Set of constants representing the allowable values for DataObjectOtherStandardColumnUnitUnitEnum
const (
	DataObjectOtherStandardColumnUnitUnitPercentage  DataObjectOtherStandardColumnUnitUnitEnum = "PERCENTAGE"
	DataObjectOtherStandardColumnUnitUnitCount       DataObjectOtherStandardColumnUnitUnitEnum = "COUNT"
	DataObjectOtherStandardColumnUnitUnitIo          DataObjectOtherStandardColumnUnitUnitEnum = "IO"
	DataObjectOtherStandardColumnUnitUnitBoolean     DataObjectOtherStandardColumnUnitUnitEnum = "BOOLEAN"
	DataObjectOtherStandardColumnUnitUnitOperation   DataObjectOtherStandardColumnUnitUnitEnum = "OPERATION"
	DataObjectOtherStandardColumnUnitUnitTransaction DataObjectOtherStandardColumnUnitUnitEnum = "TRANSACTION"
	DataObjectOtherStandardColumnUnitUnitConnection  DataObjectOtherStandardColumnUnitUnitEnum = "CONNECTION"
	DataObjectOtherStandardColumnUnitUnitAccess      DataObjectOtherStandardColumnUnitUnitEnum = "ACCESS"
	DataObjectOtherStandardColumnUnitUnitRequest     DataObjectOtherStandardColumnUnitUnitEnum = "REQUEST"
	DataObjectOtherStandardColumnUnitUnitMessage     DataObjectOtherStandardColumnUnitUnitEnum = "MESSAGE"
	DataObjectOtherStandardColumnUnitUnitExecution   DataObjectOtherStandardColumnUnitUnitEnum = "EXECUTION"
	DataObjectOtherStandardColumnUnitUnitLogons      DataObjectOtherStandardColumnUnitUnitEnum = "LOGONS"
	DataObjectOtherStandardColumnUnitUnitThread      DataObjectOtherStandardColumnUnitUnitEnum = "THREAD"
	DataObjectOtherStandardColumnUnitUnitError       DataObjectOtherStandardColumnUnitUnitEnum = "ERROR"
)

var mappingDataObjectOtherStandardColumnUnitUnitEnum = map[string]DataObjectOtherStandardColumnUnitUnitEnum{
	"PERCENTAGE":  DataObjectOtherStandardColumnUnitUnitPercentage,
	"COUNT":       DataObjectOtherStandardColumnUnitUnitCount,
	"IO":          DataObjectOtherStandardColumnUnitUnitIo,
	"BOOLEAN":     DataObjectOtherStandardColumnUnitUnitBoolean,
	"OPERATION":   DataObjectOtherStandardColumnUnitUnitOperation,
	"TRANSACTION": DataObjectOtherStandardColumnUnitUnitTransaction,
	"CONNECTION":  DataObjectOtherStandardColumnUnitUnitConnection,
	"ACCESS":      DataObjectOtherStandardColumnUnitUnitAccess,
	"REQUEST":     DataObjectOtherStandardColumnUnitUnitRequest,
	"MESSAGE":     DataObjectOtherStandardColumnUnitUnitMessage,
	"EXECUTION":   DataObjectOtherStandardColumnUnitUnitExecution,
	"LOGONS":      DataObjectOtherStandardColumnUnitUnitLogons,
	"THREAD":      DataObjectOtherStandardColumnUnitUnitThread,
	"ERROR":       DataObjectOtherStandardColumnUnitUnitError,
}

var mappingDataObjectOtherStandardColumnUnitUnitEnumLowerCase = map[string]DataObjectOtherStandardColumnUnitUnitEnum{
	"percentage":  DataObjectOtherStandardColumnUnitUnitPercentage,
	"count":       DataObjectOtherStandardColumnUnitUnitCount,
	"io":          DataObjectOtherStandardColumnUnitUnitIo,
	"boolean":     DataObjectOtherStandardColumnUnitUnitBoolean,
	"operation":   DataObjectOtherStandardColumnUnitUnitOperation,
	"transaction": DataObjectOtherStandardColumnUnitUnitTransaction,
	"connection":  DataObjectOtherStandardColumnUnitUnitConnection,
	"access":      DataObjectOtherStandardColumnUnitUnitAccess,
	"request":     DataObjectOtherStandardColumnUnitUnitRequest,
	"message":     DataObjectOtherStandardColumnUnitUnitMessage,
	"execution":   DataObjectOtherStandardColumnUnitUnitExecution,
	"logons":      DataObjectOtherStandardColumnUnitUnitLogons,
	"thread":      DataObjectOtherStandardColumnUnitUnitThread,
	"error":       DataObjectOtherStandardColumnUnitUnitError,
}

// GetDataObjectOtherStandardColumnUnitUnitEnumValues Enumerates the set of values for DataObjectOtherStandardColumnUnitUnitEnum
func GetDataObjectOtherStandardColumnUnitUnitEnumValues() []DataObjectOtherStandardColumnUnitUnitEnum {
	values := make([]DataObjectOtherStandardColumnUnitUnitEnum, 0)
	for _, v := range mappingDataObjectOtherStandardColumnUnitUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetDataObjectOtherStandardColumnUnitUnitEnumStringValues Enumerates the set of values in String for DataObjectOtherStandardColumnUnitUnitEnum
func GetDataObjectOtherStandardColumnUnitUnitEnumStringValues() []string {
	return []string{
		"PERCENTAGE",
		"COUNT",
		"IO",
		"BOOLEAN",
		"OPERATION",
		"TRANSACTION",
		"CONNECTION",
		"ACCESS",
		"REQUEST",
		"MESSAGE",
		"EXECUTION",
		"LOGONS",
		"THREAD",
		"ERROR",
	}
}

// GetMappingDataObjectOtherStandardColumnUnitUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataObjectOtherStandardColumnUnitUnitEnum(val string) (DataObjectOtherStandardColumnUnitUnitEnum, bool) {
	enum, ok := mappingDataObjectOtherStandardColumnUnitUnitEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
