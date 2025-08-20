// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CloudAsmConnectionCredentials The credentials used to connect to the Cloud ASM instance. Currently only the `DETAILS` type
// is supported for creating MACS connector credentials.
type CloudAsmConnectionCredentials interface {
}

type cloudasmconnectioncredentials struct {
	JsonData       []byte
	CredentialType string `json:"credentialType"`
}

// UnmarshalJSON unmarshals json
func (m *cloudasmconnectioncredentials) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercloudasmconnectioncredentials cloudasmconnectioncredentials
	s := struct {
		Model Unmarshalercloudasmconnectioncredentials
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CredentialType = s.Model.CredentialType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *cloudasmconnectioncredentials) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.CredentialType {
	case "DETAILS":
		mm := CloudAsmConnectionCredentialsByDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NAME_REFERENCE":
		mm := CloudAsmConnectionCredentialsByName{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CloudAsmConnectionCredentials: %s.", m.CredentialType)
		return *m, nil
	}
}

func (m cloudasmconnectioncredentials) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m cloudasmconnectioncredentials) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CloudAsmConnectionCredentialsCredentialTypeEnum Enum with underlying type: string
type CloudAsmConnectionCredentialsCredentialTypeEnum string

// Set of constants representing the allowable values for CloudAsmConnectionCredentialsCredentialTypeEnum
const (
	CloudAsmConnectionCredentialsCredentialTypeNameReference CloudAsmConnectionCredentialsCredentialTypeEnum = "NAME_REFERENCE"
	CloudAsmConnectionCredentialsCredentialTypeDetails       CloudAsmConnectionCredentialsCredentialTypeEnum = "DETAILS"
)

var mappingCloudAsmConnectionCredentialsCredentialTypeEnum = map[string]CloudAsmConnectionCredentialsCredentialTypeEnum{
	"NAME_REFERENCE": CloudAsmConnectionCredentialsCredentialTypeNameReference,
	"DETAILS":        CloudAsmConnectionCredentialsCredentialTypeDetails,
}

var mappingCloudAsmConnectionCredentialsCredentialTypeEnumLowerCase = map[string]CloudAsmConnectionCredentialsCredentialTypeEnum{
	"name_reference": CloudAsmConnectionCredentialsCredentialTypeNameReference,
	"details":        CloudAsmConnectionCredentialsCredentialTypeDetails,
}

// GetCloudAsmConnectionCredentialsCredentialTypeEnumValues Enumerates the set of values for CloudAsmConnectionCredentialsCredentialTypeEnum
func GetCloudAsmConnectionCredentialsCredentialTypeEnumValues() []CloudAsmConnectionCredentialsCredentialTypeEnum {
	values := make([]CloudAsmConnectionCredentialsCredentialTypeEnum, 0)
	for _, v := range mappingCloudAsmConnectionCredentialsCredentialTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCloudAsmConnectionCredentialsCredentialTypeEnumStringValues Enumerates the set of values in String for CloudAsmConnectionCredentialsCredentialTypeEnum
func GetCloudAsmConnectionCredentialsCredentialTypeEnumStringValues() []string {
	return []string{
		"NAME_REFERENCE",
		"DETAILS",
	}
}

// GetMappingCloudAsmConnectionCredentialsCredentialTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCloudAsmConnectionCredentialsCredentialTypeEnum(val string) (CloudAsmConnectionCredentialsCredentialTypeEnum, bool) {
	enum, ok := mappingCloudAsmConnectionCredentialsCredentialTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
