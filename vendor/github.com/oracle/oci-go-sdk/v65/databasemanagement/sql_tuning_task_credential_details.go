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

// SqlTuningTaskCredentialDetails The credential used to connect to the database.
type SqlTuningTaskCredentialDetails interface {

	// The user name used to connect to the database.
	GetUsername() *string

	// The role of the database user.
	GetRole() SqlTuningTaskCredentialDetailsRoleEnum
}

type sqltuningtaskcredentialdetails struct {
	JsonData                    []byte
	Username                    *string                                `mandatory:"true" json:"username"`
	Role                        SqlTuningTaskCredentialDetailsRoleEnum `mandatory:"true" json:"role"`
	SqlTuningTaskCredentialType string                                 `json:"sqlTuningTaskCredentialType"`
}

// UnmarshalJSON unmarshals json
func (m *sqltuningtaskcredentialdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersqltuningtaskcredentialdetails sqltuningtaskcredentialdetails
	s := struct {
		Model Unmarshalersqltuningtaskcredentialdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Username = s.Model.Username
	m.Role = s.Model.Role
	m.SqlTuningTaskCredentialType = s.Model.SqlTuningTaskCredentialType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *sqltuningtaskcredentialdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SqlTuningTaskCredentialType {
	case "SECRET":
		mm := SqlTuningTaskSecretCredentialDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PASSWORD":
		mm := SqlTuningTaskPasswordCredentialDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for SqlTuningTaskCredentialDetails: %s.", m.SqlTuningTaskCredentialType)
		return *m, nil
	}
}

// GetUsername returns Username
func (m sqltuningtaskcredentialdetails) GetUsername() *string {
	return m.Username
}

// GetRole returns Role
func (m sqltuningtaskcredentialdetails) GetRole() SqlTuningTaskCredentialDetailsRoleEnum {
	return m.Role
}

func (m sqltuningtaskcredentialdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m sqltuningtaskcredentialdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSqlTuningTaskCredentialDetailsRoleEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetSqlTuningTaskCredentialDetailsRoleEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SqlTuningTaskCredentialDetailsRoleEnum Enum with underlying type: string
type SqlTuningTaskCredentialDetailsRoleEnum string

// Set of constants representing the allowable values for SqlTuningTaskCredentialDetailsRoleEnum
const (
	SqlTuningTaskCredentialDetailsRoleNormal SqlTuningTaskCredentialDetailsRoleEnum = "NORMAL"
	SqlTuningTaskCredentialDetailsRoleSysdba SqlTuningTaskCredentialDetailsRoleEnum = "SYSDBA"
)

var mappingSqlTuningTaskCredentialDetailsRoleEnum = map[string]SqlTuningTaskCredentialDetailsRoleEnum{
	"NORMAL": SqlTuningTaskCredentialDetailsRoleNormal,
	"SYSDBA": SqlTuningTaskCredentialDetailsRoleSysdba,
}

var mappingSqlTuningTaskCredentialDetailsRoleEnumLowerCase = map[string]SqlTuningTaskCredentialDetailsRoleEnum{
	"normal": SqlTuningTaskCredentialDetailsRoleNormal,
	"sysdba": SqlTuningTaskCredentialDetailsRoleSysdba,
}

// GetSqlTuningTaskCredentialDetailsRoleEnumValues Enumerates the set of values for SqlTuningTaskCredentialDetailsRoleEnum
func GetSqlTuningTaskCredentialDetailsRoleEnumValues() []SqlTuningTaskCredentialDetailsRoleEnum {
	values := make([]SqlTuningTaskCredentialDetailsRoleEnum, 0)
	for _, v := range mappingSqlTuningTaskCredentialDetailsRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlTuningTaskCredentialDetailsRoleEnumStringValues Enumerates the set of values in String for SqlTuningTaskCredentialDetailsRoleEnum
func GetSqlTuningTaskCredentialDetailsRoleEnumStringValues() []string {
	return []string{
		"NORMAL",
		"SYSDBA",
	}
}

// GetMappingSqlTuningTaskCredentialDetailsRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlTuningTaskCredentialDetailsRoleEnum(val string) (SqlTuningTaskCredentialDetailsRoleEnum, bool) {
	enum, ok := mappingSqlTuningTaskCredentialDetailsRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
