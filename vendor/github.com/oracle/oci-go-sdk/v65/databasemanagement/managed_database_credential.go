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

// ManagedDatabaseCredential The credential used to connect to the Managed Database and obtain the details of the optimizer statistics tasks.
type ManagedDatabaseCredential interface {

	// The user name used to connect to the database.
	GetUsername() *string

	// The role of the database user.
	GetRole() ManagedDatabaseCredentialRoleEnum
}

type manageddatabasecredential struct {
	JsonData       []byte
	Username       *string                           `mandatory:"true" json:"username"`
	Role           ManagedDatabaseCredentialRoleEnum `mandatory:"true" json:"role"`
	CredentialType string                            `json:"credentialType"`
}

// UnmarshalJSON unmarshals json
func (m *manageddatabasecredential) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalermanageddatabasecredential manageddatabasecredential
	s := struct {
		Model Unmarshalermanageddatabasecredential
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Username = s.Model.Username
	m.Role = s.Model.Role
	m.CredentialType = s.Model.CredentialType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *manageddatabasecredential) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.CredentialType {
	case "PASSWORD":
		mm := ManagedDatabasePasswordCredential{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SECRET":
		mm := ManagedDatabaseSecretCredential{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ManagedDatabaseCredential: %s.", m.CredentialType)
		return *m, nil
	}
}

// GetUsername returns Username
func (m manageddatabasecredential) GetUsername() *string {
	return m.Username
}

// GetRole returns Role
func (m manageddatabasecredential) GetRole() ManagedDatabaseCredentialRoleEnum {
	return m.Role
}

func (m manageddatabasecredential) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m manageddatabasecredential) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingManagedDatabaseCredentialRoleEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetManagedDatabaseCredentialRoleEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ManagedDatabaseCredentialRoleEnum Enum with underlying type: string
type ManagedDatabaseCredentialRoleEnum string

// Set of constants representing the allowable values for ManagedDatabaseCredentialRoleEnum
const (
	ManagedDatabaseCredentialRoleNormal ManagedDatabaseCredentialRoleEnum = "NORMAL"
	ManagedDatabaseCredentialRoleSysdba ManagedDatabaseCredentialRoleEnum = "SYSDBA"
)

var mappingManagedDatabaseCredentialRoleEnum = map[string]ManagedDatabaseCredentialRoleEnum{
	"NORMAL": ManagedDatabaseCredentialRoleNormal,
	"SYSDBA": ManagedDatabaseCredentialRoleSysdba,
}

var mappingManagedDatabaseCredentialRoleEnumLowerCase = map[string]ManagedDatabaseCredentialRoleEnum{
	"normal": ManagedDatabaseCredentialRoleNormal,
	"sysdba": ManagedDatabaseCredentialRoleSysdba,
}

// GetManagedDatabaseCredentialRoleEnumValues Enumerates the set of values for ManagedDatabaseCredentialRoleEnum
func GetManagedDatabaseCredentialRoleEnumValues() []ManagedDatabaseCredentialRoleEnum {
	values := make([]ManagedDatabaseCredentialRoleEnum, 0)
	for _, v := range mappingManagedDatabaseCredentialRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetManagedDatabaseCredentialRoleEnumStringValues Enumerates the set of values in String for ManagedDatabaseCredentialRoleEnum
func GetManagedDatabaseCredentialRoleEnumStringValues() []string {
	return []string{
		"NORMAL",
		"SYSDBA",
	}
}

// GetMappingManagedDatabaseCredentialRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingManagedDatabaseCredentialRoleEnum(val string) (ManagedDatabaseCredentialRoleEnum, bool) {
	enum, ok := mappingManagedDatabaseCredentialRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ManagedDatabaseCredentialCredentialTypeEnum Enum with underlying type: string
type ManagedDatabaseCredentialCredentialTypeEnum string

// Set of constants representing the allowable values for ManagedDatabaseCredentialCredentialTypeEnum
const (
	ManagedDatabaseCredentialCredentialTypeSecret   ManagedDatabaseCredentialCredentialTypeEnum = "SECRET"
	ManagedDatabaseCredentialCredentialTypePassword ManagedDatabaseCredentialCredentialTypeEnum = "PASSWORD"
)

var mappingManagedDatabaseCredentialCredentialTypeEnum = map[string]ManagedDatabaseCredentialCredentialTypeEnum{
	"SECRET":   ManagedDatabaseCredentialCredentialTypeSecret,
	"PASSWORD": ManagedDatabaseCredentialCredentialTypePassword,
}

var mappingManagedDatabaseCredentialCredentialTypeEnumLowerCase = map[string]ManagedDatabaseCredentialCredentialTypeEnum{
	"secret":   ManagedDatabaseCredentialCredentialTypeSecret,
	"password": ManagedDatabaseCredentialCredentialTypePassword,
}

// GetManagedDatabaseCredentialCredentialTypeEnumValues Enumerates the set of values for ManagedDatabaseCredentialCredentialTypeEnum
func GetManagedDatabaseCredentialCredentialTypeEnumValues() []ManagedDatabaseCredentialCredentialTypeEnum {
	values := make([]ManagedDatabaseCredentialCredentialTypeEnum, 0)
	for _, v := range mappingManagedDatabaseCredentialCredentialTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetManagedDatabaseCredentialCredentialTypeEnumStringValues Enumerates the set of values in String for ManagedDatabaseCredentialCredentialTypeEnum
func GetManagedDatabaseCredentialCredentialTypeEnumStringValues() []string {
	return []string{
		"SECRET",
		"PASSWORD",
	}
}

// GetMappingManagedDatabaseCredentialCredentialTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingManagedDatabaseCredentialCredentialTypeEnum(val string) (ManagedDatabaseCredentialCredentialTypeEnum, bool) {
	enum, ok := mappingManagedDatabaseCredentialCredentialTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
