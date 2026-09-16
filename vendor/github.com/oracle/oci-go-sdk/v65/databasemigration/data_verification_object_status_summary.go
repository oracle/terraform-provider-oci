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

// DataVerificationObjectStatusSummary Per-object status comparison line.
type DataVerificationObjectStatusSummary interface {

	// Database object name.
	GetObjectName() *string

	// Object status in source.
	GetStatusInSource() *string

	// Object status in target.
	GetStatusInTarget() *string

	// Schema/owner name.
	// This field is omitted (null/empty depending on backend serialization) for database-wide object types
	// that are not schema-scoped.
	// Oracle non-schema-scoped object types: USER, ROLE, PROFILE, TABLESPACE, DATABASE_LINK, CONTROLFILE,
	// DATAFILE, REDO_LOG, DIRECTORY, LIBRARY, CONTEXT.
	// MySQL non-schema-scoped object types: USER, ROLE, SERVER, TABLESPACE, LOGFILE_GROUP.
	GetOwner() *string
}

type dataverificationobjectstatussummary struct {
	JsonData            []byte
	Owner               *string `mandatory:"false" json:"owner"`
	ObjectName          *string `mandatory:"true" json:"objectName"`
	StatusInSource      *string `mandatory:"true" json:"statusInSource"`
	StatusInTarget      *string `mandatory:"true" json:"statusInTarget"`
	DatabaseCombination string  `json:"databaseCombination"`
}

// UnmarshalJSON unmarshals json
func (m *dataverificationobjectstatussummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdataverificationobjectstatussummary dataverificationobjectstatussummary
	s := struct {
		Model Unmarshalerdataverificationobjectstatussummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ObjectName = s.Model.ObjectName
	m.StatusInSource = s.Model.StatusInSource
	m.StatusInTarget = s.Model.StatusInTarget
	m.Owner = s.Model.Owner
	m.DatabaseCombination = s.Model.DatabaseCombination

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *dataverificationobjectstatussummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.DatabaseCombination {
	case "MYSQL":
		mm := MySqlDataVerificationObjectStatusSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE":
		mm := OracleDataVerificationObjectStatusSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for DataVerificationObjectStatusSummary: %s.", m.DatabaseCombination)
		return *m, nil
	}
}

// GetOwner returns Owner
func (m dataverificationobjectstatussummary) GetOwner() *string {
	return m.Owner
}

// GetObjectName returns ObjectName
func (m dataverificationobjectstatussummary) GetObjectName() *string {
	return m.ObjectName
}

// GetStatusInSource returns StatusInSource
func (m dataverificationobjectstatussummary) GetStatusInSource() *string {
	return m.StatusInSource
}

// GetStatusInTarget returns StatusInTarget
func (m dataverificationobjectstatussummary) GetStatusInTarget() *string {
	return m.StatusInTarget
}

func (m dataverificationobjectstatussummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m dataverificationobjectstatussummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
