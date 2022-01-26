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

// AwrDbReport The result of the AWR report.
type AwrDbReport struct {

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
	Format AwrDbReportFormatEnum `mandatory:"false" json:"format,omitempty"`
}

//GetName returns Name
func (m AwrDbReport) GetName() *string {
	return m.Name
}

//GetVersion returns Version
func (m AwrDbReport) GetVersion() *string {
	return m.Version
}

//GetQueryKey returns QueryKey
func (m AwrDbReport) GetQueryKey() *string {
	return m.QueryKey
}

//GetDbQueryTimeInSecs returns DbQueryTimeInSecs
func (m AwrDbReport) GetDbQueryTimeInSecs() *float64 {
	return m.DbQueryTimeInSecs
}

func (m AwrDbReport) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m AwrDbReport) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAwrDbReport AwrDbReport
	s := struct {
		DiscriminatorParam string `json:"awrResultType"`
		MarshalTypeAwrDbReport
	}{
		"AWRDB_DB_REPORT",
		(MarshalTypeAwrDbReport)(m),
	}

	return json.Marshal(&s)
}

// AwrDbReportFormatEnum Enum with underlying type: string
type AwrDbReportFormatEnum string

// Set of constants representing the allowable values for AwrDbReportFormatEnum
const (
	AwrDbReportFormatHtml AwrDbReportFormatEnum = "HTML"
	AwrDbReportFormatText AwrDbReportFormatEnum = "TEXT"
	AwrDbReportFormatXml  AwrDbReportFormatEnum = "XML"
)

var mappingAwrDbReportFormat = map[string]AwrDbReportFormatEnum{
	"HTML": AwrDbReportFormatHtml,
	"TEXT": AwrDbReportFormatText,
	"XML":  AwrDbReportFormatXml,
}

// GetAwrDbReportFormatEnumValues Enumerates the set of values for AwrDbReportFormatEnum
func GetAwrDbReportFormatEnumValues() []AwrDbReportFormatEnum {
	values := make([]AwrDbReportFormatEnum, 0)
	for _, v := range mappingAwrDbReportFormat {
		values = append(values, v)
	}
	return values
}
