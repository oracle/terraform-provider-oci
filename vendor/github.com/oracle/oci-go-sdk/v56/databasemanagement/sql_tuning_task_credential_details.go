// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// SqlTuningTaskCredentialDetails The credential to be used to connect to the database.
type SqlTuningTaskCredentialDetails interface {

	// The user to connect to the database.
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
		return *m, nil
	}
}

//GetUsername returns Username
func (m sqltuningtaskcredentialdetails) GetUsername() *string {
	return m.Username
}

//GetRole returns Role
func (m sqltuningtaskcredentialdetails) GetRole() SqlTuningTaskCredentialDetailsRoleEnum {
	return m.Role
}

func (m sqltuningtaskcredentialdetails) String() string {
	return common.PointerString(m)
}

// SqlTuningTaskCredentialDetailsRoleEnum Enum with underlying type: string
type SqlTuningTaskCredentialDetailsRoleEnum string

// Set of constants representing the allowable values for SqlTuningTaskCredentialDetailsRoleEnum
const (
	SqlTuningTaskCredentialDetailsRoleNormal SqlTuningTaskCredentialDetailsRoleEnum = "NORMAL"
	SqlTuningTaskCredentialDetailsRoleSysdba SqlTuningTaskCredentialDetailsRoleEnum = "SYSDBA"
)

var mappingSqlTuningTaskCredentialDetailsRole = map[string]SqlTuningTaskCredentialDetailsRoleEnum{
	"NORMAL": SqlTuningTaskCredentialDetailsRoleNormal,
	"SYSDBA": SqlTuningTaskCredentialDetailsRoleSysdba,
}

// GetSqlTuningTaskCredentialDetailsRoleEnumValues Enumerates the set of values for SqlTuningTaskCredentialDetailsRoleEnum
func GetSqlTuningTaskCredentialDetailsRoleEnumValues() []SqlTuningTaskCredentialDetailsRoleEnum {
	values := make([]SqlTuningTaskCredentialDetailsRoleEnum, 0)
	for _, v := range mappingSqlTuningTaskCredentialDetailsRole {
		values = append(values, v)
	}
	return values
}
