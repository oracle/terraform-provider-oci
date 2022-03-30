// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Connectivity Management API
//
// Use the DCMS APIs to perform Metadata/Data operations.
//

package dataconnectivity

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v64/common"
	"strings"
)

// AbstractWriteAttribute The abstract write attribute.
type AbstractWriteAttribute interface {
}

type abstractwriteattribute struct {
	JsonData  []byte
	ModelType string `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *abstractwriteattribute) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerabstractwriteattribute abstractwriteattribute
	s := struct {
		Model Unmarshalerabstractwriteattribute
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *abstractwriteattribute) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "ORACLE_ATP_WRITE_ATTRIBUTE":
		mm := OracleAtpWriteAttributes{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HDFS_WRITE_ATTRIBUTE":
		mm := HdfsWriteAttributes{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_WRITE_ATTRIBUTE":
		mm := OracleWriteAttributes{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_ADWC_WRITE_ATTRIBUTE":
		mm := OracleAdwcWriteAttributes{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OBJECT_STORAGE_WRITE_ATTRIBUTE":
		mm := ObjectStorageWriteAttributes{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m abstractwriteattribute) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m abstractwriteattribute) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AbstractWriteAttributeModelTypeEnum Enum with underlying type: string
type AbstractWriteAttributeModelTypeEnum string

// Set of constants representing the allowable values for AbstractWriteAttributeModelTypeEnum
const (
	AbstractWriteAttributeModelTypeOracleWriteAttribute        AbstractWriteAttributeModelTypeEnum = "ORACLE_WRITE_ATTRIBUTE"
	AbstractWriteAttributeModelTypeOracleAtpWriteAttribute     AbstractWriteAttributeModelTypeEnum = "ORACLE_ATP_WRITE_ATTRIBUTE"
	AbstractWriteAttributeModelTypeOracleAdwcWriteAttribute    AbstractWriteAttributeModelTypeEnum = "ORACLE_ADWC_WRITE_ATTRIBUTE"
	AbstractWriteAttributeModelTypeObjectStorageWriteAttribute AbstractWriteAttributeModelTypeEnum = "OBJECT_STORAGE_WRITE_ATTRIBUTE"
	AbstractWriteAttributeModelTypeHdfsWriteAttribute          AbstractWriteAttributeModelTypeEnum = "HDFS_WRITE_ATTRIBUTE"
)

var mappingAbstractWriteAttributeModelTypeEnum = map[string]AbstractWriteAttributeModelTypeEnum{
	"ORACLE_WRITE_ATTRIBUTE":         AbstractWriteAttributeModelTypeOracleWriteAttribute,
	"ORACLE_ATP_WRITE_ATTRIBUTE":     AbstractWriteAttributeModelTypeOracleAtpWriteAttribute,
	"ORACLE_ADWC_WRITE_ATTRIBUTE":    AbstractWriteAttributeModelTypeOracleAdwcWriteAttribute,
	"OBJECT_STORAGE_WRITE_ATTRIBUTE": AbstractWriteAttributeModelTypeObjectStorageWriteAttribute,
	"HDFS_WRITE_ATTRIBUTE":           AbstractWriteAttributeModelTypeHdfsWriteAttribute,
}

var mappingAbstractWriteAttributeModelTypeEnumLowerCase = map[string]AbstractWriteAttributeModelTypeEnum{
	"oracle_write_attribute":         AbstractWriteAttributeModelTypeOracleWriteAttribute,
	"oracle_atp_write_attribute":     AbstractWriteAttributeModelTypeOracleAtpWriteAttribute,
	"oracle_adwc_write_attribute":    AbstractWriteAttributeModelTypeOracleAdwcWriteAttribute,
	"object_storage_write_attribute": AbstractWriteAttributeModelTypeObjectStorageWriteAttribute,
	"hdfs_write_attribute":           AbstractWriteAttributeModelTypeHdfsWriteAttribute,
}

// GetAbstractWriteAttributeModelTypeEnumValues Enumerates the set of values for AbstractWriteAttributeModelTypeEnum
func GetAbstractWriteAttributeModelTypeEnumValues() []AbstractWriteAttributeModelTypeEnum {
	values := make([]AbstractWriteAttributeModelTypeEnum, 0)
	for _, v := range mappingAbstractWriteAttributeModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAbstractWriteAttributeModelTypeEnumStringValues Enumerates the set of values in String for AbstractWriteAttributeModelTypeEnum
func GetAbstractWriteAttributeModelTypeEnumStringValues() []string {
	return []string{
		"ORACLE_WRITE_ATTRIBUTE",
		"ORACLE_ATP_WRITE_ATTRIBUTE",
		"ORACLE_ADWC_WRITE_ATTRIBUTE",
		"OBJECT_STORAGE_WRITE_ATTRIBUTE",
		"HDFS_WRITE_ATTRIBUTE",
	}
}

// GetMappingAbstractWriteAttributeModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAbstractWriteAttributeModelTypeEnum(val string) (AbstractWriteAttributeModelTypeEnum, bool) {
	enum, ok := mappingAbstractWriteAttributeModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
