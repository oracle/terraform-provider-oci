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

// MySqlDataVerificationTableRowCountSummary Table row count comparison result line for MySQL migrations.
type MySqlDataVerificationTableRowCountSummary struct {

	// Table name.
	TableName *string `mandatory:"true" json:"tableName"`

	// Row count in source.
	SourceRowCount *int `mandatory:"true" json:"sourceRowCount"`

	// Row count in target.
	TargetRowCount *int `mandatory:"true" json:"targetRowCount"`

	// Percentage variance between source and target row counts.
	// This value is expected to be non-negative.
	// A negative value indicates that the target has rows not present in the source,
	// which should be treated as an inconsistency (i.e., a failed comparison/report result).
	VariancePercent *float64 `mandatory:"true" json:"variancePercent"`

	// Time of last statistics collection in the source used for the dictionary-based row count estimate.
	// This value is reported per row because statistics can be collected at different times per object.
	// This timestamp is stored in UTC.
	TimeLastSourceStatisticsCollection *common.SDKTime `mandatory:"true" json:"timeLastSourceStatisticsCollection"`

	// Time of last statistics collection in the target used for the dictionary-based row count estimate.
	// This value is reported per row because statistics can be collected at different times per object.
	// This timestamp is stored in UTC.
	TimeLastTargetStatisticsCollection *common.SDKTime `mandatory:"true" json:"timeLastTargetStatisticsCollection"`

	// Schema/owner name.
	// This field is omitted (null/empty depending on backend serialization) for database-wide object types
	// that are not schema-scoped.
	// Oracle non-schema-scoped object types: USER, ROLE, PROFILE, TABLESPACE, DATABASE_LINK, CONTROLFILE,
	// DATAFILE, REDO_LOG, DIRECTORY, LIBRARY, CONTEXT.
	// MySQL non-schema-scoped object types: USER, ROLE, SERVER, TABLESPACE, LOGFILE_GROUP.
	Owner *string `mandatory:"false" json:"owner"`
}

// GetOwner returns Owner
func (m MySqlDataVerificationTableRowCountSummary) GetOwner() *string {
	return m.Owner
}

// GetTableName returns TableName
func (m MySqlDataVerificationTableRowCountSummary) GetTableName() *string {
	return m.TableName
}

// GetSourceRowCount returns SourceRowCount
func (m MySqlDataVerificationTableRowCountSummary) GetSourceRowCount() *int {
	return m.SourceRowCount
}

// GetTargetRowCount returns TargetRowCount
func (m MySqlDataVerificationTableRowCountSummary) GetTargetRowCount() *int {
	return m.TargetRowCount
}

// GetVariancePercent returns VariancePercent
func (m MySqlDataVerificationTableRowCountSummary) GetVariancePercent() *float64 {
	return m.VariancePercent
}

func (m MySqlDataVerificationTableRowCountSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MySqlDataVerificationTableRowCountSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m MySqlDataVerificationTableRowCountSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeMySqlDataVerificationTableRowCountSummary MySqlDataVerificationTableRowCountSummary
	s := struct {
		DiscriminatorParam string `json:"databaseCombination"`
		MarshalTypeMySqlDataVerificationTableRowCountSummary
	}{
		"MYSQL",
		(MarshalTypeMySqlDataVerificationTableRowCountSummary)(m),
	}

	return json.Marshal(&s)
}
