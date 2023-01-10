// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatabaseConnectionCredentials Credentials used to connect to the database. Currently only the `DETAILS` type is supported for creating MACS connector crendentials.
type DatabaseConnectionCredentials interface {
}

type databaseconnectioncredentials struct {
	JsonData       []byte
	CredentialType string `json:"credentialType"`
}

// UnmarshalJSON unmarshals json
func (m *databaseconnectioncredentials) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabaseconnectioncredentials databaseconnectioncredentials
	s := struct {
		Model Unmarshalerdatabaseconnectioncredentials
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CredentialType = s.Model.CredentialType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databaseconnectioncredentials) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.CredentialType {
	case "NAME_REFERENCE":
		mm := DatabaseConnectionCredentailsByName{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SSL_DETAILS":
		mm := DatabaseSslConnectionCredentials{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DETAILS":
		mm := DatabaseConnectionCredentialsByDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m databaseconnectioncredentials) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databaseconnectioncredentials) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseConnectionCredentialsCredentialTypeEnum Enum with underlying type: string
type DatabaseConnectionCredentialsCredentialTypeEnum string

// Set of constants representing the allowable values for DatabaseConnectionCredentialsCredentialTypeEnum
const (
	DatabaseConnectionCredentialsCredentialTypeNameReference DatabaseConnectionCredentialsCredentialTypeEnum = "NAME_REFERENCE"
	DatabaseConnectionCredentialsCredentialTypeDetails       DatabaseConnectionCredentialsCredentialTypeEnum = "DETAILS"
	DatabaseConnectionCredentialsCredentialTypeSslDetails    DatabaseConnectionCredentialsCredentialTypeEnum = "SSL_DETAILS"
)

var mappingDatabaseConnectionCredentialsCredentialTypeEnum = map[string]DatabaseConnectionCredentialsCredentialTypeEnum{
	"NAME_REFERENCE": DatabaseConnectionCredentialsCredentialTypeNameReference,
	"DETAILS":        DatabaseConnectionCredentialsCredentialTypeDetails,
	"SSL_DETAILS":    DatabaseConnectionCredentialsCredentialTypeSslDetails,
}

var mappingDatabaseConnectionCredentialsCredentialTypeEnumLowerCase = map[string]DatabaseConnectionCredentialsCredentialTypeEnum{
	"name_reference": DatabaseConnectionCredentialsCredentialTypeNameReference,
	"details":        DatabaseConnectionCredentialsCredentialTypeDetails,
	"ssl_details":    DatabaseConnectionCredentialsCredentialTypeSslDetails,
}

// GetDatabaseConnectionCredentialsCredentialTypeEnumValues Enumerates the set of values for DatabaseConnectionCredentialsCredentialTypeEnum
func GetDatabaseConnectionCredentialsCredentialTypeEnumValues() []DatabaseConnectionCredentialsCredentialTypeEnum {
	values := make([]DatabaseConnectionCredentialsCredentialTypeEnum, 0)
	for _, v := range mappingDatabaseConnectionCredentialsCredentialTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseConnectionCredentialsCredentialTypeEnumStringValues Enumerates the set of values in String for DatabaseConnectionCredentialsCredentialTypeEnum
func GetDatabaseConnectionCredentialsCredentialTypeEnumStringValues() []string {
	return []string{
		"NAME_REFERENCE",
		"DETAILS",
		"SSL_DETAILS",
	}
}

// GetMappingDatabaseConnectionCredentialsCredentialTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseConnectionCredentialsCredentialTypeEnum(val string) (DatabaseConnectionCredentialsCredentialTypeEnum, bool) {
	enum, ok := mappingDatabaseConnectionCredentialsCredentialTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
