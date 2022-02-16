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
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AwrDbSqlReport) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAwrDbSqlReportFormatEnum(string(m.Format)); !ok && m.Format != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Format: %s. Supported values are: %s.", m.Format, strings.Join(GetAwrDbSqlReportFormatEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingAwrDbSqlReportFormatEnum = map[string]AwrDbSqlReportFormatEnum{
	"HTML": AwrDbSqlReportFormatHtml,
	"TEXT": AwrDbSqlReportFormatText,
}

// GetAwrDbSqlReportFormatEnumValues Enumerates the set of values for AwrDbSqlReportFormatEnum
func GetAwrDbSqlReportFormatEnumValues() []AwrDbSqlReportFormatEnum {
	values := make([]AwrDbSqlReportFormatEnum, 0)
	for _, v := range mappingAwrDbSqlReportFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetAwrDbSqlReportFormatEnumStringValues Enumerates the set of values in String for AwrDbSqlReportFormatEnum
func GetAwrDbSqlReportFormatEnumStringValues() []string {
	return []string{
		"HTML",
		"TEXT",
	}
}

// GetMappingAwrDbSqlReportFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAwrDbSqlReportFormatEnum(val string) (AwrDbSqlReportFormatEnum, bool) {
	mappingAwrDbSqlReportFormatEnumIgnoreCase := make(map[string]AwrDbSqlReportFormatEnum)
	for k, v := range mappingAwrDbSqlReportFormatEnum {
		mappingAwrDbSqlReportFormatEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingAwrDbSqlReportFormatEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
