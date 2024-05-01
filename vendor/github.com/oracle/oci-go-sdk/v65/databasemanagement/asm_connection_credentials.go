// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AsmConnectionCredentials The credentials used to connect to the ASM instance. Currently only the `DETAILS` type
// is supported for creating MACS connector credentials.
type AsmConnectionCredentials interface {
}

type asmconnectioncredentials struct {
	JsonData       []byte
	CredentialType string `json:"credentialType"`
}

// UnmarshalJSON unmarshals json
func (m *asmconnectioncredentials) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerasmconnectioncredentials asmconnectioncredentials
	s := struct {
		Model Unmarshalerasmconnectioncredentials
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CredentialType = s.Model.CredentialType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *asmconnectioncredentials) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.CredentialType {
	case "NAME_REFERENCE":
		mm := AsmConnectionCredentailsByName{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DETAILS":
		mm := AsmConnectionCredentialsByDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for AsmConnectionCredentials: %s.", m.CredentialType)
		return *m, nil
	}
}

func (m asmconnectioncredentials) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m asmconnectioncredentials) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AsmConnectionCredentialsCredentialTypeEnum Enum with underlying type: string
type AsmConnectionCredentialsCredentialTypeEnum string

// Set of constants representing the allowable values for AsmConnectionCredentialsCredentialTypeEnum
const (
	AsmConnectionCredentialsCredentialTypeNameReference AsmConnectionCredentialsCredentialTypeEnum = "NAME_REFERENCE"
	AsmConnectionCredentialsCredentialTypeDetails       AsmConnectionCredentialsCredentialTypeEnum = "DETAILS"
)

var mappingAsmConnectionCredentialsCredentialTypeEnum = map[string]AsmConnectionCredentialsCredentialTypeEnum{
	"NAME_REFERENCE": AsmConnectionCredentialsCredentialTypeNameReference,
	"DETAILS":        AsmConnectionCredentialsCredentialTypeDetails,
}

var mappingAsmConnectionCredentialsCredentialTypeEnumLowerCase = map[string]AsmConnectionCredentialsCredentialTypeEnum{
	"name_reference": AsmConnectionCredentialsCredentialTypeNameReference,
	"details":        AsmConnectionCredentialsCredentialTypeDetails,
}

// GetAsmConnectionCredentialsCredentialTypeEnumValues Enumerates the set of values for AsmConnectionCredentialsCredentialTypeEnum
func GetAsmConnectionCredentialsCredentialTypeEnumValues() []AsmConnectionCredentialsCredentialTypeEnum {
	values := make([]AsmConnectionCredentialsCredentialTypeEnum, 0)
	for _, v := range mappingAsmConnectionCredentialsCredentialTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAsmConnectionCredentialsCredentialTypeEnumStringValues Enumerates the set of values in String for AsmConnectionCredentialsCredentialTypeEnum
func GetAsmConnectionCredentialsCredentialTypeEnumStringValues() []string {
	return []string{
		"NAME_REFERENCE",
		"DETAILS",
	}
}

// GetMappingAsmConnectionCredentialsCredentialTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAsmConnectionCredentialsCredentialTypeEnum(val string) (AsmConnectionCredentialsCredentialTypeEnum, bool) {
	enum, ok := mappingAsmConnectionCredentialsCredentialTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
