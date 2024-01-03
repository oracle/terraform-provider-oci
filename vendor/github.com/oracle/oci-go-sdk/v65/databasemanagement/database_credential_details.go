// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatabaseCredentialDetails The credential to connect to the database to perform tablespace administration tasks.
type DatabaseCredentialDetails interface {
}

type databasecredentialdetails struct {
	JsonData       []byte
	CredentialType string `json:"credentialType"`
}

// UnmarshalJSON unmarshals json
func (m *databasecredentialdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabasecredentialdetails databasecredentialdetails
	s := struct {
		Model Unmarshalerdatabasecredentialdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CredentialType = s.Model.CredentialType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databasecredentialdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.CredentialType {
	case "SECRET":
		mm := DatabaseSecretCredentialDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NAMED_CREDENTIAL":
		mm := DatabaseNamedCredentialDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PASSWORD":
		mm := DatabasePasswordCredentialDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DatabaseCredentialDetails: %s.", m.CredentialType)
		return *m, nil
	}
}

func (m databasecredentialdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databasecredentialdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseCredentialDetailsCredentialTypeEnum Enum with underlying type: string
type DatabaseCredentialDetailsCredentialTypeEnum string

// Set of constants representing the allowable values for DatabaseCredentialDetailsCredentialTypeEnum
const (
	DatabaseCredentialDetailsCredentialTypeSecret          DatabaseCredentialDetailsCredentialTypeEnum = "SECRET"
	DatabaseCredentialDetailsCredentialTypePassword        DatabaseCredentialDetailsCredentialTypeEnum = "PASSWORD"
	DatabaseCredentialDetailsCredentialTypeNamedCredential DatabaseCredentialDetailsCredentialTypeEnum = "NAMED_CREDENTIAL"
)

var mappingDatabaseCredentialDetailsCredentialTypeEnum = map[string]DatabaseCredentialDetailsCredentialTypeEnum{
	"SECRET":           DatabaseCredentialDetailsCredentialTypeSecret,
	"PASSWORD":         DatabaseCredentialDetailsCredentialTypePassword,
	"NAMED_CREDENTIAL": DatabaseCredentialDetailsCredentialTypeNamedCredential,
}

var mappingDatabaseCredentialDetailsCredentialTypeEnumLowerCase = map[string]DatabaseCredentialDetailsCredentialTypeEnum{
	"secret":           DatabaseCredentialDetailsCredentialTypeSecret,
	"password":         DatabaseCredentialDetailsCredentialTypePassword,
	"named_credential": DatabaseCredentialDetailsCredentialTypeNamedCredential,
}

// GetDatabaseCredentialDetailsCredentialTypeEnumValues Enumerates the set of values for DatabaseCredentialDetailsCredentialTypeEnum
func GetDatabaseCredentialDetailsCredentialTypeEnumValues() []DatabaseCredentialDetailsCredentialTypeEnum {
	values := make([]DatabaseCredentialDetailsCredentialTypeEnum, 0)
	for _, v := range mappingDatabaseCredentialDetailsCredentialTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseCredentialDetailsCredentialTypeEnumStringValues Enumerates the set of values in String for DatabaseCredentialDetailsCredentialTypeEnum
func GetDatabaseCredentialDetailsCredentialTypeEnumStringValues() []string {
	return []string{
		"SECRET",
		"PASSWORD",
		"NAMED_CREDENTIAL",
	}
}

// GetMappingDatabaseCredentialDetailsCredentialTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseCredentialDetailsCredentialTypeEnum(val string) (DatabaseCredentialDetailsCredentialTypeEnum, bool) {
	enum, ok := mappingDatabaseCredentialDetailsCredentialTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
