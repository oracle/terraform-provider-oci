// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MySqlDataVerificationObjectStatusSummary Per-object status comparison line for MySQL migrations.
type MySqlDataVerificationObjectStatusSummary struct {

	// Database object name.
	ObjectName *string `mandatory:"true" json:"objectName"`

	// Object status in source.
	StatusInSource *string `mandatory:"true" json:"statusInSource"`

	// Object status in target.
	StatusInTarget *string `mandatory:"true" json:"statusInTarget"`

	// Schema/owner name.
	// This field is omitted (null/empty depending on backend serialization) for database-wide object types
	// that are not schema-scoped.
	// Oracle non-schema-scoped object types: USER, ROLE, PROFILE, TABLESPACE, DATABASE_LINK, CONTROLFILE,
	// DATAFILE, REDO_LOG, DIRECTORY, LIBRARY, CONTEXT.
	// MySQL non-schema-scoped object types: USER, ROLE, SERVER, TABLESPACE, LOGFILE_GROUP.
	Owner *string `mandatory:"false" json:"owner"`

	// MySQL database object type.
	ObjectType MySqlDatabaseObjectTypesEnum `mandatory:"true" json:"objectType"`
}

// GetOwner returns Owner
func (m MySqlDataVerificationObjectStatusSummary) GetOwner() *string {
	return m.Owner
}

// GetObjectName returns ObjectName
func (m MySqlDataVerificationObjectStatusSummary) GetObjectName() *string {
	return m.ObjectName
}

// GetStatusInSource returns StatusInSource
func (m MySqlDataVerificationObjectStatusSummary) GetStatusInSource() *string {
	return m.StatusInSource
}

// GetStatusInTarget returns StatusInTarget
func (m MySqlDataVerificationObjectStatusSummary) GetStatusInTarget() *string {
	return m.StatusInTarget
}

func (m MySqlDataVerificationObjectStatusSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MySqlDataVerificationObjectStatusSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingMySqlDatabaseObjectTypesEnum(string(m.ObjectType)); !ok && m.ObjectType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ObjectType: %s. Supported values are: %s.", m.ObjectType, strings.Join(GetMySqlDatabaseObjectTypesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m MySqlDataVerificationObjectStatusSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeMySqlDataVerificationObjectStatusSummary MySqlDataVerificationObjectStatusSummary
	s := struct {
		DiscriminatorParam string `json:"databaseCombination"`
		MarshalTypeMySqlDataVerificationObjectStatusSummary
	}{
		"MYSQL",
		(MarshalTypeMySqlDataVerificationObjectStatusSummary)(m),
	}

	return json.Marshal(&s)
}
