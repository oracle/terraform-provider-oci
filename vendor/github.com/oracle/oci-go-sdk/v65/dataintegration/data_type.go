// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DataType A `DataType` object is a simple primitive type that describes the type of a single atomic unit of data.  For example, `INT`, `VARCHAR`, `NUMBER`, and so on.
type DataType struct {

	// The key of the object.
	Key *string `mandatory:"false" json:"key"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// A user defined description for the object.
	Description *string `mandatory:"false" json:"description"`

	// The data type system name.
	TypeSystemName *string `mandatory:"false" json:"typeSystemName"`

	ConfigDefinition *ConfigDefinition `mandatory:"false" json:"configDefinition"`

	// The data type.
	DtType DataTypeDtTypeEnum `mandatory:"false" json:"dtType,omitempty"`
}

// GetKey returns Key
func (m DataType) GetKey() *string {
	return m.Key
}

// GetModelVersion returns ModelVersion
func (m DataType) GetModelVersion() *string {
	return m.ModelVersion
}

// GetParentRef returns ParentRef
func (m DataType) GetParentRef() *ParentReference {
	return m.ParentRef
}

// GetName returns Name
func (m DataType) GetName() *string {
	return m.Name
}

// GetObjectStatus returns ObjectStatus
func (m DataType) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetDescription returns Description
func (m DataType) GetDescription() *string {
	return m.Description
}

func (m DataType) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataType) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDataTypeDtTypeEnum(string(m.DtType)); !ok && m.DtType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DtType: %s. Supported values are: %s.", m.DtType, strings.Join(GetDataTypeDtTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DataType) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDataType DataType
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeDataType
	}{
		"DATA_TYPE",
		(MarshalTypeDataType)(m),
	}

	return json.Marshal(&s)
}

// DataTypeDtTypeEnum Enum with underlying type: string
type DataTypeDtTypeEnum string

// Set of constants representing the allowable values for DataTypeDtTypeEnum
const (
	DataTypeDtTypePrimitive  DataTypeDtTypeEnum = "PRIMITIVE"
	DataTypeDtTypeStructured DataTypeDtTypeEnum = "STRUCTURED"
)

var mappingDataTypeDtTypeEnum = map[string]DataTypeDtTypeEnum{
	"PRIMITIVE":  DataTypeDtTypePrimitive,
	"STRUCTURED": DataTypeDtTypeStructured,
}

var mappingDataTypeDtTypeEnumLowerCase = map[string]DataTypeDtTypeEnum{
	"primitive":  DataTypeDtTypePrimitive,
	"structured": DataTypeDtTypeStructured,
}

// GetDataTypeDtTypeEnumValues Enumerates the set of values for DataTypeDtTypeEnum
func GetDataTypeDtTypeEnumValues() []DataTypeDtTypeEnum {
	values := make([]DataTypeDtTypeEnum, 0)
	for _, v := range mappingDataTypeDtTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDataTypeDtTypeEnumStringValues Enumerates the set of values in String for DataTypeDtTypeEnum
func GetDataTypeDtTypeEnumStringValues() []string {
	return []string{
		"PRIMITIVE",
		"STRUCTURED",
	}
}

// GetMappingDataTypeDtTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataTypeDtTypeEnum(val string) (DataTypeDtTypeEnum, bool) {
	enum, ok := mappingDataTypeDtTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
