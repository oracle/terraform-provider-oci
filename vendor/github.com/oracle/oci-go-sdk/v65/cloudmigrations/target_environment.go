// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TargetEnvironment Description of the target environment.
type TargetEnvironment interface {

	// Target compartment identifier
	GetTargetCompartmentId() *string
}

type targetenvironment struct {
	JsonData              []byte
	TargetCompartmentId   *string `mandatory:"false" json:"targetCompartmentId"`
	TargetEnvironmentType string  `json:"targetEnvironmentType"`
}

// UnmarshalJSON unmarshals json
func (m *targetenvironment) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertargetenvironment targetenvironment
	s := struct {
		Model Unmarshalertargetenvironment
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.TargetCompartmentId = s.Model.TargetCompartmentId
	m.TargetEnvironmentType = s.Model.TargetEnvironmentType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *targetenvironment) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.TargetEnvironmentType {
	case "VM_TARGET_ENV":
		mm := VmTargetEnvironment{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for TargetEnvironment: %s.", m.TargetEnvironmentType)
		return *m, nil
	}
}

// GetTargetCompartmentId returns TargetCompartmentId
func (m targetenvironment) GetTargetCompartmentId() *string {
	return m.TargetCompartmentId
}

func (m targetenvironment) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m targetenvironment) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TargetEnvironmentTargetEnvironmentTypeEnum Enum with underlying type: string
type TargetEnvironmentTargetEnvironmentTypeEnum string

// Set of constants representing the allowable values for TargetEnvironmentTargetEnvironmentTypeEnum
const (
	TargetEnvironmentTargetEnvironmentTypeVmTargetEnv TargetEnvironmentTargetEnvironmentTypeEnum = "VM_TARGET_ENV"
)

var mappingTargetEnvironmentTargetEnvironmentTypeEnum = map[string]TargetEnvironmentTargetEnvironmentTypeEnum{
	"VM_TARGET_ENV": TargetEnvironmentTargetEnvironmentTypeVmTargetEnv,
}

var mappingTargetEnvironmentTargetEnvironmentTypeEnumLowerCase = map[string]TargetEnvironmentTargetEnvironmentTypeEnum{
	"vm_target_env": TargetEnvironmentTargetEnvironmentTypeVmTargetEnv,
}

// GetTargetEnvironmentTargetEnvironmentTypeEnumValues Enumerates the set of values for TargetEnvironmentTargetEnvironmentTypeEnum
func GetTargetEnvironmentTargetEnvironmentTypeEnumValues() []TargetEnvironmentTargetEnvironmentTypeEnum {
	values := make([]TargetEnvironmentTargetEnvironmentTypeEnum, 0)
	for _, v := range mappingTargetEnvironmentTargetEnvironmentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTargetEnvironmentTargetEnvironmentTypeEnumStringValues Enumerates the set of values in String for TargetEnvironmentTargetEnvironmentTypeEnum
func GetTargetEnvironmentTargetEnvironmentTypeEnumStringValues() []string {
	return []string{
		"VM_TARGET_ENV",
	}
}

// GetMappingTargetEnvironmentTargetEnvironmentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTargetEnvironmentTargetEnvironmentTypeEnum(val string) (TargetEnvironmentTargetEnvironmentTypeEnum, bool) {
	enum, ok := mappingTargetEnvironmentTargetEnvironmentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
