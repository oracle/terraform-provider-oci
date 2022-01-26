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

// AwrDbSqlReport The result of the AWR SQL report.
type AwrDbSqlReport struct {

	// The name of the query result.
	Name *string `mandatory:"true" json:"name"`

	// The version of the query result.
	Version *string `mandatory:"false" json:"version"`

	// The ID assigned to the query instance.
	QueryKey *string `mandatory:"false" json:"queryKey"`

	// The time taken to query the database tier (in seconds).
	DbQueryTimeInSecs *float64 `mandatory:"false" json:"dbQueryTimeInSecs"`

	// The content of the report.
	Content *string `mandatory:"false" json:"content"`

	// The format of the report.
	Format AwrDbSqlReportFormatEnum `mandatory:"false" json:"format,omitempty"`
}

//GetName returns Name
func (m AwrDbSqlReport) GetName() *string {
	return m.Name
}

//GetVersion returns Version
func (m AwrDbSqlReport) GetVersion() *string {
	return m.Version
}

//GetQueryKey returns QueryKey
func (m AwrDbSqlReport) GetQueryKey() *string {
	return m.QueryKey
}

//GetDbQueryTimeInSecs returns DbQueryTimeInSecs
func (m AwrDbSqlReport) GetDbQueryTimeInSecs() *float64 {
	return m.DbQueryTimeInSecs
}

func (m AwrDbSqlReport) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m AwrDbSqlReport) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAwrDbSqlReport AwrDbSqlReport
	s := struct {
		DiscriminatorParam string `json:"awrResultType"`
		MarshalTypeAwrDbSqlReport
	}{
		"AWRDB_SQL_REPORT",
		(MarshalTypeAwrDbSqlReport)(m),
	}

	return json.Marshal(&s)
}

// AwrDbSqlReportFormatEnum Enum with underlying type: string
type AwrDbSqlReportFormatEnum string

// Set of constants representing the allowable values for AwrDbSqlReportFormatEnum
const (
	AwrDbSqlReportFormatHtml AwrDbSqlReportFormatEnum = "HTML"
	AwrDbSqlReportFormatText AwrDbSqlReportFormatEnum = "TEXT"
)

var mappingAwrDbSqlReportFormat = map[string]AwrDbSqlReportFormatEnum{
	"HTML": AwrDbSqlReportFormatHtml,
	"TEXT": AwrDbSqlReportFormatText,
}

// GetAwrDbSqlReportFormatEnumValues Enumerates the set of values for AwrDbSqlReportFormatEnum
func GetAwrDbSqlReportFormatEnumValues() []AwrDbSqlReportFormatEnum {
	values := make([]AwrDbSqlReportFormatEnum, 0)
	for _, v := range mappingAwrDbSqlReportFormat {
		values = append(values, v)
	}
	return values
}
