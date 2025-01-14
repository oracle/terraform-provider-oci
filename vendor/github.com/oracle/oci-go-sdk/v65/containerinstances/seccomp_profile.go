// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Instance API
//
// A description of the Container Instance API
//

package containerinstances

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SeccompProfile Seccomp profile to be applied to the container.
type SeccompProfile interface {
}

type seccompprofile struct {
	JsonData           []byte
	SeccompProfileType string `json:"seccompProfileType"`
}

// UnmarshalJSON unmarshals json
func (m *seccompprofile) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerseccompprofile seccompprofile
	s := struct {
		Model Unmarshalerseccompprofile
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SeccompProfileType = s.Model.SeccompProfileType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *seccompprofile) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SeccompProfileType {
	case "RUNTIME_DEFAULT":
		mm := RuntimeDefaultSeccompProfile{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for SeccompProfile: %s.", m.SeccompProfileType)
		return *m, nil
	}
}

func (m seccompprofile) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m seccompprofile) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SeccompProfileSeccompProfileTypeEnum Enum with underlying type: string
type SeccompProfileSeccompProfileTypeEnum string

// Set of constants representing the allowable values for SeccompProfileSeccompProfileTypeEnum
const (
	SeccompProfileSeccompProfileTypeRuntimeDefault SeccompProfileSeccompProfileTypeEnum = "RUNTIME_DEFAULT"
)

var mappingSeccompProfileSeccompProfileTypeEnum = map[string]SeccompProfileSeccompProfileTypeEnum{
	"RUNTIME_DEFAULT": SeccompProfileSeccompProfileTypeRuntimeDefault,
}

var mappingSeccompProfileSeccompProfileTypeEnumLowerCase = map[string]SeccompProfileSeccompProfileTypeEnum{
	"runtime_default": SeccompProfileSeccompProfileTypeRuntimeDefault,
}

// GetSeccompProfileSeccompProfileTypeEnumValues Enumerates the set of values for SeccompProfileSeccompProfileTypeEnum
func GetSeccompProfileSeccompProfileTypeEnumValues() []SeccompProfileSeccompProfileTypeEnum {
	values := make([]SeccompProfileSeccompProfileTypeEnum, 0)
	for _, v := range mappingSeccompProfileSeccompProfileTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSeccompProfileSeccompProfileTypeEnumStringValues Enumerates the set of values in String for SeccompProfileSeccompProfileTypeEnum
func GetSeccompProfileSeccompProfileTypeEnumStringValues() []string {
	return []string{
		"RUNTIME_DEFAULT",
	}
}

// GetMappingSeccompProfileSeccompProfileTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSeccompProfileSeccompProfileTypeEnum(val string) (SeccompProfileSeccompProfileTypeEnum, bool) {
	enum, ok := mappingSeccompProfileSeccompProfileTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
