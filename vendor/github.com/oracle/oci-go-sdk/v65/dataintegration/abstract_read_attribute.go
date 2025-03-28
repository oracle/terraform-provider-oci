// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// AbstractReadAttribute The abstract read attribute.
type AbstractReadAttribute interface {
}

type abstractreadattribute struct {
	JsonData  []byte
	ModelType string `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *abstractreadattribute) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerabstractreadattribute abstractreadattribute
	s := struct {
		Model Unmarshalerabstractreadattribute
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *abstractreadattribute) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "ORACLE_READ_ATTRIBUTE":
		mm := OracleReadAttributes{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BICC_READ_ATTRIBUTE":
		mm := BiccReadAttributes{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BIP_READ_ATTRIBUTE":
		mm := BipReadAttributes{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLEREADATTRIBUTE":
		mm := OracleReadAttribute{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for AbstractReadAttribute: %s.", m.ModelType)
		return *m, nil
	}
}

func (m abstractreadattribute) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m abstractreadattribute) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AbstractReadAttributeModelTypeEnum Enum with underlying type: string
type AbstractReadAttributeModelTypeEnum string

// Set of constants representing the allowable values for AbstractReadAttributeModelTypeEnum
const (
	AbstractReadAttributeModelTypeOraclereadattribute AbstractReadAttributeModelTypeEnum = "ORACLEREADATTRIBUTE"
	AbstractReadAttributeModelTypeOracleReadAttribute AbstractReadAttributeModelTypeEnum = "ORACLE_READ_ATTRIBUTE"
	AbstractReadAttributeModelTypeBiccReadAttribute   AbstractReadAttributeModelTypeEnum = "BICC_READ_ATTRIBUTE"
	AbstractReadAttributeModelTypeBipReadAttribute    AbstractReadAttributeModelTypeEnum = "BIP_READ_ATTRIBUTE"
)

var mappingAbstractReadAttributeModelTypeEnum = map[string]AbstractReadAttributeModelTypeEnum{
	"ORACLEREADATTRIBUTE":   AbstractReadAttributeModelTypeOraclereadattribute,
	"ORACLE_READ_ATTRIBUTE": AbstractReadAttributeModelTypeOracleReadAttribute,
	"BICC_READ_ATTRIBUTE":   AbstractReadAttributeModelTypeBiccReadAttribute,
	"BIP_READ_ATTRIBUTE":    AbstractReadAttributeModelTypeBipReadAttribute,
}

var mappingAbstractReadAttributeModelTypeEnumLowerCase = map[string]AbstractReadAttributeModelTypeEnum{
	"oraclereadattribute":   AbstractReadAttributeModelTypeOraclereadattribute,
	"oracle_read_attribute": AbstractReadAttributeModelTypeOracleReadAttribute,
	"bicc_read_attribute":   AbstractReadAttributeModelTypeBiccReadAttribute,
	"bip_read_attribute":    AbstractReadAttributeModelTypeBipReadAttribute,
}

// GetAbstractReadAttributeModelTypeEnumValues Enumerates the set of values for AbstractReadAttributeModelTypeEnum
func GetAbstractReadAttributeModelTypeEnumValues() []AbstractReadAttributeModelTypeEnum {
	values := make([]AbstractReadAttributeModelTypeEnum, 0)
	for _, v := range mappingAbstractReadAttributeModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAbstractReadAttributeModelTypeEnumStringValues Enumerates the set of values in String for AbstractReadAttributeModelTypeEnum
func GetAbstractReadAttributeModelTypeEnumStringValues() []string {
	return []string{
		"ORACLEREADATTRIBUTE",
		"ORACLE_READ_ATTRIBUTE",
		"BICC_READ_ATTRIBUTE",
		"BIP_READ_ATTRIBUTE",
	}
}

// GetMappingAbstractReadAttributeModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAbstractReadAttributeModelTypeEnum(val string) (AbstractReadAttributeModelTypeEnum, bool) {
	enum, ok := mappingAbstractReadAttributeModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
