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

// DataObjectColumnMetadata Metadata of a column in a data object resultset.
type DataObjectColumnMetadata struct {

	// Name of the column.
	Name *string `mandatory:"true" json:"name"`

	// Category of the column.
	Category DataObjectColumnMetadataCategoryEnum `mandatory:"false" json:"category,omitempty"`

	// Type of a data object column.
	DataType *string `mandatory:"false" json:"dataType"`

	// Type name of a data object column.
	DataTypeName DataObjectColumnMetadataDataTypeNameEnum `mandatory:"false" json:"dataTypeName,omitempty"`

	// Display name of the column.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of the column.
	Description *string `mandatory:"false" json:"description"`

	// Group name of the column.
	GroupName *string `mandatory:"false" json:"groupName"`

	UnitDetails DataObjectColumnUnit `mandatory:"false" json:"unitDetails"`
}

func (m DataObjectColumnMetadata) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataObjectColumnMetadata) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDataObjectColumnMetadataCategoryEnum(string(m.Category)); !ok && m.Category != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Category: %s. Supported values are: %s.", m.Category, strings.Join(GetDataObjectColumnMetadataCategoryEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDataObjectColumnMetadataDataTypeNameEnum(string(m.DataTypeName)); !ok && m.DataTypeName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataTypeName: %s. Supported values are: %s.", m.DataTypeName, strings.Join(GetDataObjectColumnMetadataDataTypeNameEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *DataObjectColumnMetadata) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Category     DataObjectColumnMetadataCategoryEnum     `json:"category"`
		DataType     *string                                  `json:"dataType"`
		DataTypeName DataObjectColumnMetadataDataTypeNameEnum `json:"dataTypeName"`
		DisplayName  *string                                  `json:"displayName"`
		Description  *string                                  `json:"description"`
		GroupName    *string                                  `json:"groupName"`
		UnitDetails  dataobjectcolumnunit                     `json:"unitDetails"`
		Name         *string                                  `json:"name"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Category = model.Category

	m.DataType = model.DataType

	m.DataTypeName = model.DataTypeName

	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.GroupName = model.GroupName

	nn, e = model.UnitDetails.UnmarshalPolymorphicJSON(model.UnitDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.UnitDetails = nn.(DataObjectColumnUnit)
	} else {
		m.UnitDetails = nil
	}

	m.Name = model.Name

	return
}

// DataObjectColumnMetadataCategoryEnum Enum with underlying type: string
type DataObjectColumnMetadataCategoryEnum string

// Set of constants representing the allowable values for DataObjectColumnMetadataCategoryEnum
const (
	DataObjectColumnMetadataCategoryDimension     DataObjectColumnMetadataCategoryEnum = "DIMENSION"
	DataObjectColumnMetadataCategoryMetric        DataObjectColumnMetadataCategoryEnum = "METRIC"
	DataObjectColumnMetadataCategoryTimeDimension DataObjectColumnMetadataCategoryEnum = "TIME_DIMENSION"
	DataObjectColumnMetadataCategoryUnknown       DataObjectColumnMetadataCategoryEnum = "UNKNOWN"
)

var mappingDataObjectColumnMetadataCategoryEnum = map[string]DataObjectColumnMetadataCategoryEnum{
	"DIMENSION":      DataObjectColumnMetadataCategoryDimension,
	"METRIC":         DataObjectColumnMetadataCategoryMetric,
	"TIME_DIMENSION": DataObjectColumnMetadataCategoryTimeDimension,
	"UNKNOWN":        DataObjectColumnMetadataCategoryUnknown,
}

var mappingDataObjectColumnMetadataCategoryEnumLowerCase = map[string]DataObjectColumnMetadataCategoryEnum{
	"dimension":      DataObjectColumnMetadataCategoryDimension,
	"metric":         DataObjectColumnMetadataCategoryMetric,
	"time_dimension": DataObjectColumnMetadataCategoryTimeDimension,
	"unknown":        DataObjectColumnMetadataCategoryUnknown,
}

// GetDataObjectColumnMetadataCategoryEnumValues Enumerates the set of values for DataObjectColumnMetadataCategoryEnum
func GetDataObjectColumnMetadataCategoryEnumValues() []DataObjectColumnMetadataCategoryEnum {
	values := make([]DataObjectColumnMetadataCategoryEnum, 0)
	for _, v := range mappingDataObjectColumnMetadataCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetDataObjectColumnMetadataCategoryEnumStringValues Enumerates the set of values in String for DataObjectColumnMetadataCategoryEnum
func GetDataObjectColumnMetadataCategoryEnumStringValues() []string {
	return []string{
		"DIMENSION",
		"METRIC",
		"TIME_DIMENSION",
		"UNKNOWN",
	}
}

// GetMappingDataObjectColumnMetadataCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataObjectColumnMetadataCategoryEnum(val string) (DataObjectColumnMetadataCategoryEnum, bool) {
	enum, ok := mappingDataObjectColumnMetadataCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DataObjectColumnMetadataDataTypeNameEnum Enum with underlying type: string
type DataObjectColumnMetadataDataTypeNameEnum string

// Set of constants representing the allowable values for DataObjectColumnMetadataDataTypeNameEnum
const (
	DataObjectColumnMetadataDataTypeNameNumber    DataObjectColumnMetadataDataTypeNameEnum = "NUMBER"
	DataObjectColumnMetadataDataTypeNameTimestamp DataObjectColumnMetadataDataTypeNameEnum = "TIMESTAMP"
	DataObjectColumnMetadataDataTypeNameVarchar2  DataObjectColumnMetadataDataTypeNameEnum = "VARCHAR2"
	DataObjectColumnMetadataDataTypeNameOther     DataObjectColumnMetadataDataTypeNameEnum = "OTHER"
)

var mappingDataObjectColumnMetadataDataTypeNameEnum = map[string]DataObjectColumnMetadataDataTypeNameEnum{
	"NUMBER":    DataObjectColumnMetadataDataTypeNameNumber,
	"TIMESTAMP": DataObjectColumnMetadataDataTypeNameTimestamp,
	"VARCHAR2":  DataObjectColumnMetadataDataTypeNameVarchar2,
	"OTHER":     DataObjectColumnMetadataDataTypeNameOther,
}

var mappingDataObjectColumnMetadataDataTypeNameEnumLowerCase = map[string]DataObjectColumnMetadataDataTypeNameEnum{
	"number":    DataObjectColumnMetadataDataTypeNameNumber,
	"timestamp": DataObjectColumnMetadataDataTypeNameTimestamp,
	"varchar2":  DataObjectColumnMetadataDataTypeNameVarchar2,
	"other":     DataObjectColumnMetadataDataTypeNameOther,
}

// GetDataObjectColumnMetadataDataTypeNameEnumValues Enumerates the set of values for DataObjectColumnMetadataDataTypeNameEnum
func GetDataObjectColumnMetadataDataTypeNameEnumValues() []DataObjectColumnMetadataDataTypeNameEnum {
	values := make([]DataObjectColumnMetadataDataTypeNameEnum, 0)
	for _, v := range mappingDataObjectColumnMetadataDataTypeNameEnum {
		values = append(values, v)
	}
	return values
}

// GetDataObjectColumnMetadataDataTypeNameEnumStringValues Enumerates the set of values in String for DataObjectColumnMetadataDataTypeNameEnum
func GetDataObjectColumnMetadataDataTypeNameEnumStringValues() []string {
	return []string{
		"NUMBER",
		"TIMESTAMP",
		"VARCHAR2",
		"OTHER",
	}
}

// GetMappingDataObjectColumnMetadataDataTypeNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataObjectColumnMetadataDataTypeNameEnum(val string) (DataObjectColumnMetadataDataTypeNameEnum, bool) {
	enum, ok := mappingDataObjectColumnMetadataDataTypeNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
