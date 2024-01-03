// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OciFunction The information about the OCI Function.
type OciFunction struct {

	// Ocid of the OCI Function.
	FunctionId *string `mandatory:"false" json:"functionId"`

	// Region where the OCI Function is deployed.
	RegionId *string `mandatory:"false" json:"regionId"`

	FnConfigDefinition *ConfigDefinition `mandatory:"false" json:"fnConfigDefinition"`

	InputShape *Shape `mandatory:"false" json:"inputShape"`

	OutputShape *Shape `mandatory:"false" json:"outputShape"`

	// The type of the OCI Function object.
	ModelType OciFunctionModelTypeEnum `mandatory:"false" json:"modelType,omitempty"`

	// The key identifying the OCI Function operator object, use this to identiy this instance within the dataflow.
	Key *string `mandatory:"false" json:"key"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	// The OCI Function payload format.
	PayloadFormat OciFunctionPayloadFormatEnum `mandatory:"false" json:"payloadFormat,omitempty"`

	FnConfigDef *FunctionConfigurationDefinition `mandatory:"false" json:"fnConfigDef"`
}

func (m OciFunction) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OciFunction) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOciFunctionModelTypeEnum(string(m.ModelType)); !ok && m.ModelType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ModelType: %s. Supported values are: %s.", m.ModelType, strings.Join(GetOciFunctionModelTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOciFunctionPayloadFormatEnum(string(m.PayloadFormat)); !ok && m.PayloadFormat != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PayloadFormat: %s. Supported values are: %s.", m.PayloadFormat, strings.Join(GetOciFunctionPayloadFormatEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OciFunctionModelTypeEnum Enum with underlying type: string
type OciFunctionModelTypeEnum string

// Set of constants representing the allowable values for OciFunctionModelTypeEnum
const (
	OciFunctionModelTypeOciFunction OciFunctionModelTypeEnum = "OCI_FUNCTION"
)

var mappingOciFunctionModelTypeEnum = map[string]OciFunctionModelTypeEnum{
	"OCI_FUNCTION": OciFunctionModelTypeOciFunction,
}

var mappingOciFunctionModelTypeEnumLowerCase = map[string]OciFunctionModelTypeEnum{
	"oci_function": OciFunctionModelTypeOciFunction,
}

// GetOciFunctionModelTypeEnumValues Enumerates the set of values for OciFunctionModelTypeEnum
func GetOciFunctionModelTypeEnumValues() []OciFunctionModelTypeEnum {
	values := make([]OciFunctionModelTypeEnum, 0)
	for _, v := range mappingOciFunctionModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOciFunctionModelTypeEnumStringValues Enumerates the set of values in String for OciFunctionModelTypeEnum
func GetOciFunctionModelTypeEnumStringValues() []string {
	return []string{
		"OCI_FUNCTION",
	}
}

// GetMappingOciFunctionModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOciFunctionModelTypeEnum(val string) (OciFunctionModelTypeEnum, bool) {
	enum, ok := mappingOciFunctionModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OciFunctionPayloadFormatEnum Enum with underlying type: string
type OciFunctionPayloadFormatEnum string

// Set of constants representing the allowable values for OciFunctionPayloadFormatEnum
const (
	OciFunctionPayloadFormatJson      OciFunctionPayloadFormatEnum = "JSON"
	OciFunctionPayloadFormatAvro      OciFunctionPayloadFormatEnum = "AVRO"
	OciFunctionPayloadFormatJsonbytes OciFunctionPayloadFormatEnum = "JSONBYTES"
)

var mappingOciFunctionPayloadFormatEnum = map[string]OciFunctionPayloadFormatEnum{
	"JSON":      OciFunctionPayloadFormatJson,
	"AVRO":      OciFunctionPayloadFormatAvro,
	"JSONBYTES": OciFunctionPayloadFormatJsonbytes,
}

var mappingOciFunctionPayloadFormatEnumLowerCase = map[string]OciFunctionPayloadFormatEnum{
	"json":      OciFunctionPayloadFormatJson,
	"avro":      OciFunctionPayloadFormatAvro,
	"jsonbytes": OciFunctionPayloadFormatJsonbytes,
}

// GetOciFunctionPayloadFormatEnumValues Enumerates the set of values for OciFunctionPayloadFormatEnum
func GetOciFunctionPayloadFormatEnumValues() []OciFunctionPayloadFormatEnum {
	values := make([]OciFunctionPayloadFormatEnum, 0)
	for _, v := range mappingOciFunctionPayloadFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetOciFunctionPayloadFormatEnumStringValues Enumerates the set of values in String for OciFunctionPayloadFormatEnum
func GetOciFunctionPayloadFormatEnumStringValues() []string {
	return []string{
		"JSON",
		"AVRO",
		"JSONBYTES",
	}
}

// GetMappingOciFunctionPayloadFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOciFunctionPayloadFormatEnum(val string) (OciFunctionPayloadFormatEnum, bool) {
	enum, ok := mappingOciFunctionPayloadFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
