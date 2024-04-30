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

// SqlTuningSetAdminCredentialDetails The credential to connect to the database to perform Sql tuning set administration tasks.
type SqlTuningSetAdminCredentialDetails interface {

	// The user to connect to the database.
	GetUsername() *string

	// The role of the database user.
	GetRole() SqlTuningSetAdminCredentialDetailsRoleEnum
}

type sqltuningsetadmincredentialdetails struct {
	JsonData                        []byte
	Username                        *string                                    `mandatory:"true" json:"username"`
	Role                            SqlTuningSetAdminCredentialDetailsRoleEnum `mandatory:"true" json:"role"`
	SqlTuningSetAdminCredentialType string                                     `json:"sqlTuningSetAdminCredentialType"`
}

// UnmarshalJSON unmarshals json
func (m *sqltuningsetadmincredentialdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersqltuningsetadmincredentialdetails sqltuningsetadmincredentialdetails
	s := struct {
		Model Unmarshalersqltuningsetadmincredentialdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Username = s.Model.Username
	m.Role = s.Model.Role
	m.SqlTuningSetAdminCredentialType = s.Model.SqlTuningSetAdminCredentialType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *sqltuningsetadmincredentialdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SqlTuningSetAdminCredentialType {
	case "PASSWORD":
		mm := SqlTuningSetAdminPasswordCredentialDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SECRET":
		mm := SqlTuningSetAdminSecretCredentialDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for SqlTuningSetAdminCredentialDetails: %s.", m.SqlTuningSetAdminCredentialType)
		return *m, nil
	}
}

// GetUsername returns Username
func (m sqltuningsetadmincredentialdetails) GetUsername() *string {
	return m.Username
}

// GetRole returns Role
func (m sqltuningsetadmincredentialdetails) GetRole() SqlTuningSetAdminCredentialDetailsRoleEnum {
	return m.Role
}

func (m sqltuningsetadmincredentialdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m sqltuningsetadmincredentialdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSqlTuningSetAdminCredentialDetailsRoleEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetSqlTuningSetAdminCredentialDetailsRoleEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SqlTuningSetAdminCredentialDetailsRoleEnum Enum with underlying type: string
type SqlTuningSetAdminCredentialDetailsRoleEnum string

// Set of constants representing the allowable values for SqlTuningSetAdminCredentialDetailsRoleEnum
const (
	SqlTuningSetAdminCredentialDetailsRoleNormal SqlTuningSetAdminCredentialDetailsRoleEnum = "NORMAL"
	SqlTuningSetAdminCredentialDetailsRoleSysdba SqlTuningSetAdminCredentialDetailsRoleEnum = "SYSDBA"
)

var mappingSqlTuningSetAdminCredentialDetailsRoleEnum = map[string]SqlTuningSetAdminCredentialDetailsRoleEnum{
	"NORMAL": SqlTuningSetAdminCredentialDetailsRoleNormal,
	"SYSDBA": SqlTuningSetAdminCredentialDetailsRoleSysdba,
}

var mappingSqlTuningSetAdminCredentialDetailsRoleEnumLowerCase = map[string]SqlTuningSetAdminCredentialDetailsRoleEnum{
	"normal": SqlTuningSetAdminCredentialDetailsRoleNormal,
	"sysdba": SqlTuningSetAdminCredentialDetailsRoleSysdba,
}

// GetSqlTuningSetAdminCredentialDetailsRoleEnumValues Enumerates the set of values for SqlTuningSetAdminCredentialDetailsRoleEnum
func GetSqlTuningSetAdminCredentialDetailsRoleEnumValues() []SqlTuningSetAdminCredentialDetailsRoleEnum {
	values := make([]SqlTuningSetAdminCredentialDetailsRoleEnum, 0)
	for _, v := range mappingSqlTuningSetAdminCredentialDetailsRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlTuningSetAdminCredentialDetailsRoleEnumStringValues Enumerates the set of values in String for SqlTuningSetAdminCredentialDetailsRoleEnum
func GetSqlTuningSetAdminCredentialDetailsRoleEnumStringValues() []string {
	return []string{
		"NORMAL",
		"SYSDBA",
	}
}

// GetMappingSqlTuningSetAdminCredentialDetailsRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlTuningSetAdminCredentialDetailsRoleEnum(val string) (SqlTuningSetAdminCredentialDetailsRoleEnum, bool) {
	enum, ok := mappingSqlTuningSetAdminCredentialDetailsRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SqlTuningSetAdminCredentialDetailsSqlTuningSetAdminCredentialTypeEnum Enum with underlying type: string
type SqlTuningSetAdminCredentialDetailsSqlTuningSetAdminCredentialTypeEnum string

// Set of constants representing the allowable values for SqlTuningSetAdminCredentialDetailsSqlTuningSetAdminCredentialTypeEnum
const (
	SqlTuningSetAdminCredentialDetailsSqlTuningSetAdminCredentialTypeSecret   SqlTuningSetAdminCredentialDetailsSqlTuningSetAdminCredentialTypeEnum = "SECRET"
	SqlTuningSetAdminCredentialDetailsSqlTuningSetAdminCredentialTypePassword SqlTuningSetAdminCredentialDetailsSqlTuningSetAdminCredentialTypeEnum = "PASSWORD"
)

var mappingSqlTuningSetAdminCredentialDetailsSqlTuningSetAdminCredentialTypeEnum = map[string]SqlTuningSetAdminCredentialDetailsSqlTuningSetAdminCredentialTypeEnum{
	"SECRET":   SqlTuningSetAdminCredentialDetailsSqlTuningSetAdminCredentialTypeSecret,
	"PASSWORD": SqlTuningSetAdminCredentialDetailsSqlTuningSetAdminCredentialTypePassword,
}

var mappingSqlTuningSetAdminCredentialDetailsSqlTuningSetAdminCredentialTypeEnumLowerCase = map[string]SqlTuningSetAdminCredentialDetailsSqlTuningSetAdminCredentialTypeEnum{
	"secret":   SqlTuningSetAdminCredentialDetailsSqlTuningSetAdminCredentialTypeSecret,
	"password": SqlTuningSetAdminCredentialDetailsSqlTuningSetAdminCredentialTypePassword,
}

// GetSqlTuningSetAdminCredentialDetailsSqlTuningSetAdminCredentialTypeEnumValues Enumerates the set of values for SqlTuningSetAdminCredentialDetailsSqlTuningSetAdminCredentialTypeEnum
func GetSqlTuningSetAdminCredentialDetailsSqlTuningSetAdminCredentialTypeEnumValues() []SqlTuningSetAdminCredentialDetailsSqlTuningSetAdminCredentialTypeEnum {
	values := make([]SqlTuningSetAdminCredentialDetailsSqlTuningSetAdminCredentialTypeEnum, 0)
	for _, v := range mappingSqlTuningSetAdminCredentialDetailsSqlTuningSetAdminCredentialTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlTuningSetAdminCredentialDetailsSqlTuningSetAdminCredentialTypeEnumStringValues Enumerates the set of values in String for SqlTuningSetAdminCredentialDetailsSqlTuningSetAdminCredentialTypeEnum
func GetSqlTuningSetAdminCredentialDetailsSqlTuningSetAdminCredentialTypeEnumStringValues() []string {
	return []string{
		"SECRET",
		"PASSWORD",
	}
}

// GetMappingSqlTuningSetAdminCredentialDetailsSqlTuningSetAdminCredentialTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlTuningSetAdminCredentialDetailsSqlTuningSetAdminCredentialTypeEnum(val string) (SqlTuningSetAdminCredentialDetailsSqlTuningSetAdminCredentialTypeEnum, bool) {
	enum, ok := mappingSqlTuningSetAdminCredentialDetailsSqlTuningSetAdminCredentialTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
