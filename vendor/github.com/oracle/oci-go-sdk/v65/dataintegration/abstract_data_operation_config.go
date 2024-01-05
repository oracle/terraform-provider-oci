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

// AbstractDataOperationConfig The information about the data operation.
type AbstractDataOperationConfig interface {

	// This map is used for passing extra metatdata configuration that is required by read / write operation.
	GetMetadataConfigProperties() map[string]string

	// this map is used for passing BIP report parameter values.
	GetDerivedAttributes() map[string]string

	GetCallAttribute() *BipCallAttribute
}

type abstractdataoperationconfig struct {
	JsonData                 []byte
	MetadataConfigProperties map[string]string `mandatory:"false" json:"metadataConfigProperties"`
	DerivedAttributes        map[string]string `mandatory:"false" json:"derivedAttributes"`
	CallAttribute            *BipCallAttribute `mandatory:"false" json:"callAttribute"`
	ModelType                string            `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *abstractdataoperationconfig) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerabstractdataoperationconfig abstractdataoperationconfig
	s := struct {
		Model Unmarshalerabstractdataoperationconfig
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.MetadataConfigProperties = s.Model.MetadataConfigProperties
	m.DerivedAttributes = s.Model.DerivedAttributes
	m.CallAttribute = s.Model.CallAttribute
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *abstractdataoperationconfig) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "WRITE_OPERATION_CONFIG":
		mm := WriteOperationConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "READ_OPERATION_CONFIG":
		mm := ReadOperationConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for AbstractDataOperationConfig: %s.", m.ModelType)
		return *m, nil
	}
}

// GetMetadataConfigProperties returns MetadataConfigProperties
func (m abstractdataoperationconfig) GetMetadataConfigProperties() map[string]string {
	return m.MetadataConfigProperties
}

// GetDerivedAttributes returns DerivedAttributes
func (m abstractdataoperationconfig) GetDerivedAttributes() map[string]string {
	return m.DerivedAttributes
}

// GetCallAttribute returns CallAttribute
func (m abstractdataoperationconfig) GetCallAttribute() *BipCallAttribute {
	return m.CallAttribute
}

func (m abstractdataoperationconfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m abstractdataoperationconfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AbstractDataOperationConfigModelTypeEnum Enum with underlying type: string
type AbstractDataOperationConfigModelTypeEnum string

// Set of constants representing the allowable values for AbstractDataOperationConfigModelTypeEnum
const (
	AbstractDataOperationConfigModelTypeReadOperationConfig  AbstractDataOperationConfigModelTypeEnum = "READ_OPERATION_CONFIG"
	AbstractDataOperationConfigModelTypeWriteOperationConfig AbstractDataOperationConfigModelTypeEnum = "WRITE_OPERATION_CONFIG"
)

var mappingAbstractDataOperationConfigModelTypeEnum = map[string]AbstractDataOperationConfigModelTypeEnum{
	"READ_OPERATION_CONFIG":  AbstractDataOperationConfigModelTypeReadOperationConfig,
	"WRITE_OPERATION_CONFIG": AbstractDataOperationConfigModelTypeWriteOperationConfig,
}

var mappingAbstractDataOperationConfigModelTypeEnumLowerCase = map[string]AbstractDataOperationConfigModelTypeEnum{
	"read_operation_config":  AbstractDataOperationConfigModelTypeReadOperationConfig,
	"write_operation_config": AbstractDataOperationConfigModelTypeWriteOperationConfig,
}

// GetAbstractDataOperationConfigModelTypeEnumValues Enumerates the set of values for AbstractDataOperationConfigModelTypeEnum
func GetAbstractDataOperationConfigModelTypeEnumValues() []AbstractDataOperationConfigModelTypeEnum {
	values := make([]AbstractDataOperationConfigModelTypeEnum, 0)
	for _, v := range mappingAbstractDataOperationConfigModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAbstractDataOperationConfigModelTypeEnumStringValues Enumerates the set of values in String for AbstractDataOperationConfigModelTypeEnum
func GetAbstractDataOperationConfigModelTypeEnumStringValues() []string {
	return []string{
		"READ_OPERATION_CONFIG",
		"WRITE_OPERATION_CONFIG",
	}
}

// GetMappingAbstractDataOperationConfigModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAbstractDataOperationConfigModelTypeEnum(val string) (AbstractDataOperationConfigModelTypeEnum, bool) {
	enum, ok := mappingAbstractDataOperationConfigModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
