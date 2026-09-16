// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RunDataVerificationDetails Details for running Data Verification.
// This request is modeled with a type discriminator so future Data Verification run options can be
// added without changing the required request body contract.
// Note: The migration database combination is inferred from the migrationId in the path and is not part
// of this request body.
type RunDataVerificationDetails interface {
}

type rundataverificationdetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *rundataverificationdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerrundataverificationdetails rundataverificationdetails
	s := struct {
		Model Unmarshalerrundataverificationdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *rundataverificationdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "DEFAULT":
		mm := DefaultRunDataVerificationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for RunDataVerificationDetails: %s.", m.Type)
		return *m, nil
	}
}

func (m rundataverificationdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m rundataverificationdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RunDataVerificationDetailsTypeEnum Enum with underlying type: string
type RunDataVerificationDetailsTypeEnum string

// Set of constants representing the allowable values for RunDataVerificationDetailsTypeEnum
const (
	RunDataVerificationDetailsTypeDefault RunDataVerificationDetailsTypeEnum = "DEFAULT"
)

var mappingRunDataVerificationDetailsTypeEnum = map[string]RunDataVerificationDetailsTypeEnum{
	"DEFAULT": RunDataVerificationDetailsTypeDefault,
}

var mappingRunDataVerificationDetailsTypeEnumLowerCase = map[string]RunDataVerificationDetailsTypeEnum{
	"default": RunDataVerificationDetailsTypeDefault,
}

// GetRunDataVerificationDetailsTypeEnumValues Enumerates the set of values for RunDataVerificationDetailsTypeEnum
func GetRunDataVerificationDetailsTypeEnumValues() []RunDataVerificationDetailsTypeEnum {
	values := make([]RunDataVerificationDetailsTypeEnum, 0)
	for _, v := range mappingRunDataVerificationDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRunDataVerificationDetailsTypeEnumStringValues Enumerates the set of values in String for RunDataVerificationDetailsTypeEnum
func GetRunDataVerificationDetailsTypeEnumStringValues() []string {
	return []string{
		"DEFAULT",
	}
}

// GetMappingRunDataVerificationDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRunDataVerificationDetailsTypeEnum(val string) (RunDataVerificationDetailsTypeEnum, bool) {
	enum, ok := mappingRunDataVerificationDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
