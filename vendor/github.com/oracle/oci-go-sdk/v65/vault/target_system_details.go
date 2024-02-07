// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vault Secret Management API
//
// Use the Secret Management API to manage secrets and secret versions. For more information, see Managing Secrets (https://docs.cloud.oracle.com/Content/KeyManagement/Tasks/managingsecrets.htm).
//

package vault

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TargetSystemDetails The TargetSystemDetails provides the targetSystem type and type-specific connection metadata
type TargetSystemDetails interface {
}

type targetsystemdetails struct {
	JsonData         []byte
	TargetSystemType string `json:"targetSystemType"`
}

// UnmarshalJSON unmarshals json
func (m *targetsystemdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertargetsystemdetails targetsystemdetails
	s := struct {
		Model Unmarshalertargetsystemdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.TargetSystemType = s.Model.TargetSystemType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *targetsystemdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.TargetSystemType {
	case "ADB":
		mm := AdbTargetSystemDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FUNCTION":
		mm := FunctionTargetSystemDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for TargetSystemDetails: %s.", m.TargetSystemType)
		return *m, nil
	}
}

func (m targetsystemdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m targetsystemdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TargetSystemDetailsTargetSystemTypeEnum Enum with underlying type: string
type TargetSystemDetailsTargetSystemTypeEnum string

// Set of constants representing the allowable values for TargetSystemDetailsTargetSystemTypeEnum
const (
	TargetSystemDetailsTargetSystemTypeAdb      TargetSystemDetailsTargetSystemTypeEnum = "ADB"
	TargetSystemDetailsTargetSystemTypeFunction TargetSystemDetailsTargetSystemTypeEnum = "FUNCTION"
)

var mappingTargetSystemDetailsTargetSystemTypeEnum = map[string]TargetSystemDetailsTargetSystemTypeEnum{
	"ADB":      TargetSystemDetailsTargetSystemTypeAdb,
	"FUNCTION": TargetSystemDetailsTargetSystemTypeFunction,
}

var mappingTargetSystemDetailsTargetSystemTypeEnumLowerCase = map[string]TargetSystemDetailsTargetSystemTypeEnum{
	"adb":      TargetSystemDetailsTargetSystemTypeAdb,
	"function": TargetSystemDetailsTargetSystemTypeFunction,
}

// GetTargetSystemDetailsTargetSystemTypeEnumValues Enumerates the set of values for TargetSystemDetailsTargetSystemTypeEnum
func GetTargetSystemDetailsTargetSystemTypeEnumValues() []TargetSystemDetailsTargetSystemTypeEnum {
	values := make([]TargetSystemDetailsTargetSystemTypeEnum, 0)
	for _, v := range mappingTargetSystemDetailsTargetSystemTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTargetSystemDetailsTargetSystemTypeEnumStringValues Enumerates the set of values in String for TargetSystemDetailsTargetSystemTypeEnum
func GetTargetSystemDetailsTargetSystemTypeEnumStringValues() []string {
	return []string{
		"ADB",
		"FUNCTION",
	}
}

// GetMappingTargetSystemDetailsTargetSystemTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTargetSystemDetailsTargetSystemTypeEnum(val string) (TargetSystemDetailsTargetSystemTypeEnum, bool) {
	enum, ok := mappingTargetSystemDetailsTargetSystemTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
