// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management Service API. Use this API to for all FAMS related activities.
// To manage fleets,view complaince report for the Fleet,scedule patches and other lifecycle activities
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CredentialEntitySpecificDetails Credential Details
type CredentialEntitySpecificDetails interface {
}

type credentialentityspecificdetails struct {
	JsonData        []byte
	CredentialLevel string `json:"credentialLevel"`
}

// UnmarshalJSON unmarshals json
func (m *credentialentityspecificdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercredentialentityspecificdetails credentialentityspecificdetails
	s := struct {
		Model Unmarshalercredentialentityspecificdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CredentialLevel = s.Model.CredentialLevel

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *credentialentityspecificdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.CredentialLevel {
	case "TARGET":
		mm := TargetCredentialEntitySpecificDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CredentialEntitySpecificDetails: %s.", m.CredentialLevel)
		return *m, nil
	}
}

func (m credentialentityspecificdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m credentialentityspecificdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CredentialEntitySpecificDetailsCredentialLevelEnum Enum with underlying type: string
type CredentialEntitySpecificDetailsCredentialLevelEnum string

// Set of constants representing the allowable values for CredentialEntitySpecificDetailsCredentialLevelEnum
const (
	CredentialEntitySpecificDetailsCredentialLevelTarget CredentialEntitySpecificDetailsCredentialLevelEnum = "TARGET"
)

var mappingCredentialEntitySpecificDetailsCredentialLevelEnum = map[string]CredentialEntitySpecificDetailsCredentialLevelEnum{
	"TARGET": CredentialEntitySpecificDetailsCredentialLevelTarget,
}

var mappingCredentialEntitySpecificDetailsCredentialLevelEnumLowerCase = map[string]CredentialEntitySpecificDetailsCredentialLevelEnum{
	"target": CredentialEntitySpecificDetailsCredentialLevelTarget,
}

// GetCredentialEntitySpecificDetailsCredentialLevelEnumValues Enumerates the set of values for CredentialEntitySpecificDetailsCredentialLevelEnum
func GetCredentialEntitySpecificDetailsCredentialLevelEnumValues() []CredentialEntitySpecificDetailsCredentialLevelEnum {
	values := make([]CredentialEntitySpecificDetailsCredentialLevelEnum, 0)
	for _, v := range mappingCredentialEntitySpecificDetailsCredentialLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetCredentialEntitySpecificDetailsCredentialLevelEnumStringValues Enumerates the set of values in String for CredentialEntitySpecificDetailsCredentialLevelEnum
func GetCredentialEntitySpecificDetailsCredentialLevelEnumStringValues() []string {
	return []string{
		"TARGET",
	}
}

// GetMappingCredentialEntitySpecificDetailsCredentialLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCredentialEntitySpecificDetailsCredentialLevelEnum(val string) (CredentialEntitySpecificDetailsCredentialLevelEnum, bool) {
	enum, ok := mappingCredentialEntitySpecificDetailsCredentialLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
