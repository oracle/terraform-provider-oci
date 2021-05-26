// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v41/common"
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

// DatabaseConnectionCredentialsCredentialTypeEnum Enum with underlying type: string
type DatabaseConnectionCredentialsCredentialTypeEnum string

// Set of constants representing the allowable values for DatabaseConnectionCredentialsCredentialTypeEnum
const (
	DatabaseConnectionCredentialsCredentialTypeNameReference DatabaseConnectionCredentialsCredentialTypeEnum = "NAME_REFERENCE"
	DatabaseConnectionCredentialsCredentialTypeDetails       DatabaseConnectionCredentialsCredentialTypeEnum = "DETAILS"
)

var mappingDatabaseConnectionCredentialsCredentialType = map[string]DatabaseConnectionCredentialsCredentialTypeEnum{
	"NAME_REFERENCE": DatabaseConnectionCredentialsCredentialTypeNameReference,
	"DETAILS":        DatabaseConnectionCredentialsCredentialTypeDetails,
}

// GetDatabaseConnectionCredentialsCredentialTypeEnumValues Enumerates the set of values for DatabaseConnectionCredentialsCredentialTypeEnum
func GetDatabaseConnectionCredentialsCredentialTypeEnumValues() []DatabaseConnectionCredentialsCredentialTypeEnum {
	values := make([]DatabaseConnectionCredentialsCredentialTypeEnum, 0)
	for _, v := range mappingDatabaseConnectionCredentialsCredentialType {
		values = append(values, v)
	}
	return values
}
