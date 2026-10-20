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

// DataVerificationObjectTypeCountSummary Object type count comparison summary line.
type DataVerificationObjectTypeCountSummary interface {

	// Number of objects in source.
	GetSourceObjectCount() *int

	// Number of objects in target.
	GetTargetObjectCount() *int

	// Percentage delta between source and target counts.
	// This value is expected to be non-negative.
	// If a negative value is encountered, it indicates that the target has objects not present in the source,
	// which should be treated as an inconsistency (i.e., a failed comparison/report result).
	GetDeltaPercent() *float64

	// Schema/owner name.
	// This field is omitted (null/empty depending on backend serialization) for database-wide object types
	// that are not schema-scoped.
	// Oracle non-schema-scoped object types: USER, ROLE, PROFILE, TABLESPACE, DATABASE_LINK, CONTROLFILE,
	// DATAFILE, REDO_LOG, DIRECTORY, LIBRARY, CONTEXT.
	// MySQL non-schema-scoped object types: USER, ROLE, SERVER, TABLESPACE, LOGFILE_GROUP.
	GetSchemaName() *string
}

type dataverificationobjecttypecountsummary struct {
	JsonData            []byte
	SchemaName          *string  `mandatory:"false" json:"schemaName"`
	SourceObjectCount   *int     `mandatory:"true" json:"sourceObjectCount"`
	TargetObjectCount   *int     `mandatory:"true" json:"targetObjectCount"`
	DeltaPercent        *float64 `mandatory:"true" json:"deltaPercent"`
	DatabaseCombination string   `json:"databaseCombination"`
}

// UnmarshalJSON unmarshals json
func (m *dataverificationobjecttypecountsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdataverificationobjecttypecountsummary dataverificationobjecttypecountsummary
	s := struct {
		Model Unmarshalerdataverificationobjecttypecountsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SourceObjectCount = s.Model.SourceObjectCount
	m.TargetObjectCount = s.Model.TargetObjectCount
	m.DeltaPercent = s.Model.DeltaPercent
	m.SchemaName = s.Model.SchemaName
	m.DatabaseCombination = s.Model.DatabaseCombination

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *dataverificationobjecttypecountsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.DatabaseCombination {
	case "MYSQL":
		mm := MySqlDataVerificationObjectTypeCountSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE":
		mm := OracleDataVerificationObjectTypeCountSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for DataVerificationObjectTypeCountSummary: %s.", m.DatabaseCombination)
		return *m, nil
	}
}

// GetSchemaName returns SchemaName
func (m dataverificationobjecttypecountsummary) GetSchemaName() *string {
	return m.SchemaName
}

// GetSourceObjectCount returns SourceObjectCount
func (m dataverificationobjecttypecountsummary) GetSourceObjectCount() *int {
	return m.SourceObjectCount
}

// GetTargetObjectCount returns TargetObjectCount
func (m dataverificationobjecttypecountsummary) GetTargetObjectCount() *int {
	return m.TargetObjectCount
}

// GetDeltaPercent returns DeltaPercent
func (m dataverificationobjecttypecountsummary) GetDeltaPercent() *float64 {
	return m.DeltaPercent
}

func (m dataverificationobjecttypecountsummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m dataverificationobjecttypecountsummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
