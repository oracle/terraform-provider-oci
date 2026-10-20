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

// DataVerificationTableRowCountSummary Table row count comparison result line.
type DataVerificationTableRowCountSummary interface {

	// Table name.
	GetTableName() *string

	// Row count in source.
	GetSourceRowCount() *int

	// Row count in target.
	GetTargetRowCount() *int

	// Percentage variance between source and target row counts.
	// This value is expected to be non-negative.
	// A negative value indicates that the target has rows not present in the source,
	// which should be treated as an inconsistency (i.e., a failed comparison/report result).
	GetVariancePercent() *float64

	// Schema/owner name.
	// This field is omitted (null/empty depending on backend serialization) for database-wide object types
	// that are not schema-scoped.
	// Oracle non-schema-scoped object types: USER, ROLE, PROFILE, TABLESPACE, DATABASE_LINK, CONTROLFILE,
	// DATAFILE, REDO_LOG, DIRECTORY, LIBRARY, CONTEXT.
	// MySQL non-schema-scoped object types: USER, ROLE, SERVER, TABLESPACE, LOGFILE_GROUP.
	GetOwner() *string
}

type dataverificationtablerowcountsummary struct {
	JsonData            []byte
	Owner               *string  `mandatory:"false" json:"owner"`
	TableName           *string  `mandatory:"true" json:"tableName"`
	SourceRowCount      *int     `mandatory:"true" json:"sourceRowCount"`
	TargetRowCount      *int     `mandatory:"true" json:"targetRowCount"`
	VariancePercent     *float64 `mandatory:"true" json:"variancePercent"`
	DatabaseCombination string   `json:"databaseCombination"`
}

// UnmarshalJSON unmarshals json
func (m *dataverificationtablerowcountsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdataverificationtablerowcountsummary dataverificationtablerowcountsummary
	s := struct {
		Model Unmarshalerdataverificationtablerowcountsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.TableName = s.Model.TableName
	m.SourceRowCount = s.Model.SourceRowCount
	m.TargetRowCount = s.Model.TargetRowCount
	m.VariancePercent = s.Model.VariancePercent
	m.Owner = s.Model.Owner
	m.DatabaseCombination = s.Model.DatabaseCombination

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *dataverificationtablerowcountsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.DatabaseCombination {
	case "MYSQL":
		mm := MySqlDataVerificationTableRowCountSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE":
		mm := OracleDataVerificationTableRowCountSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for DataVerificationTableRowCountSummary: %s.", m.DatabaseCombination)
		return *m, nil
	}
}

// GetOwner returns Owner
func (m dataverificationtablerowcountsummary) GetOwner() *string {
	return m.Owner
}

// GetTableName returns TableName
func (m dataverificationtablerowcountsummary) GetTableName() *string {
	return m.TableName
}

// GetSourceRowCount returns SourceRowCount
func (m dataverificationtablerowcountsummary) GetSourceRowCount() *int {
	return m.SourceRowCount
}

// GetTargetRowCount returns TargetRowCount
func (m dataverificationtablerowcountsummary) GetTargetRowCount() *int {
	return m.TargetRowCount
}

// GetVariancePercent returns VariancePercent
func (m dataverificationtablerowcountsummary) GetVariancePercent() *float64 {
	return m.VariancePercent
}

func (m dataverificationtablerowcountsummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m dataverificationtablerowcountsummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
