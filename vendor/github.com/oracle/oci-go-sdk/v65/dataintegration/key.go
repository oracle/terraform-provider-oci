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

// Key The key object.
type Key interface {
}

type key struct {
	JsonData  []byte
	ModelType string `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *key) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerkey key
	s := struct {
		Model Unmarshalerkey
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *key) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "FOREIGN_KEY":
		mm := ForeignKey{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for Key: %s.", m.ModelType)
		return *m, nil
	}
}

func (m key) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m key) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// KeyModelTypeEnum Enum with underlying type: string
type KeyModelTypeEnum string

// Set of constants representing the allowable values for KeyModelTypeEnum
const (
	KeyModelTypeForeignKey KeyModelTypeEnum = "FOREIGN_KEY"
)

var mappingKeyModelTypeEnum = map[string]KeyModelTypeEnum{
	"FOREIGN_KEY": KeyModelTypeForeignKey,
}

var mappingKeyModelTypeEnumLowerCase = map[string]KeyModelTypeEnum{
	"foreign_key": KeyModelTypeForeignKey,
}

// GetKeyModelTypeEnumValues Enumerates the set of values for KeyModelTypeEnum
func GetKeyModelTypeEnumValues() []KeyModelTypeEnum {
	values := make([]KeyModelTypeEnum, 0)
	for _, v := range mappingKeyModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetKeyModelTypeEnumStringValues Enumerates the set of values in String for KeyModelTypeEnum
func GetKeyModelTypeEnumStringValues() []string {
	return []string{
		"FOREIGN_KEY",
	}
}

// GetMappingKeyModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingKeyModelTypeEnum(val string) (KeyModelTypeEnum, bool) {
	enum, ok := mappingKeyModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
