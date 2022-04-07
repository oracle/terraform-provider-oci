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
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AbstractCallAttribute The call attributes
type AbstractCallAttribute interface {
}

type abstractcallattribute struct {
	JsonData  []byte
	ModelType string `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *abstractcallattribute) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerabstractcallattribute abstractcallattribute
	s := struct {
		Model Unmarshalerabstractcallattribute
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *abstractcallattribute) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "BIPCALLATTRIBUTE":
		mm := BipCallAttribute{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m abstractcallattribute) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m abstractcallattribute) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AbstractCallAttributeModelTypeEnum Enum with underlying type: string
type AbstractCallAttributeModelTypeEnum string

// Set of constants representing the allowable values for AbstractCallAttributeModelTypeEnum
const (
	AbstractCallAttributeModelTypeBipcallattribute AbstractCallAttributeModelTypeEnum = "BIPCALLATTRIBUTE"
)

var mappingAbstractCallAttributeModelTypeEnum = map[string]AbstractCallAttributeModelTypeEnum{
	"BIPCALLATTRIBUTE": AbstractCallAttributeModelTypeBipcallattribute,
}

var mappingAbstractCallAttributeModelTypeEnumLowerCase = map[string]AbstractCallAttributeModelTypeEnum{
	"bipcallattribute": AbstractCallAttributeModelTypeBipcallattribute,
}

// GetAbstractCallAttributeModelTypeEnumValues Enumerates the set of values for AbstractCallAttributeModelTypeEnum
func GetAbstractCallAttributeModelTypeEnumValues() []AbstractCallAttributeModelTypeEnum {
	values := make([]AbstractCallAttributeModelTypeEnum, 0)
	for _, v := range mappingAbstractCallAttributeModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAbstractCallAttributeModelTypeEnumStringValues Enumerates the set of values in String for AbstractCallAttributeModelTypeEnum
func GetAbstractCallAttributeModelTypeEnumStringValues() []string {
	return []string{
		"BIPCALLATTRIBUTE",
	}
}

// GetMappingAbstractCallAttributeModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAbstractCallAttributeModelTypeEnum(val string) (AbstractCallAttributeModelTypeEnum, bool) {
	enum, ok := mappingAbstractCallAttributeModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
