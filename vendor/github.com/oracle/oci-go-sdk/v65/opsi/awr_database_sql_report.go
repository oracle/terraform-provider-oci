// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AwrDatabaseSqlReport The result of the AWR SQL report.
type AwrDatabaseSqlReport struct {

	// The name of the query result.
	Name *string `mandatory:"true" json:"name"`

	// The version of the query result.
	Version *string `mandatory:"false" json:"version"`

	// The time taken to query the database tier (in seconds).
	DbQueryTimeInSecs *float64 `mandatory:"false" json:"dbQueryTimeInSecs"`

	// The content of the report.
	Content *string `mandatory:"false" json:"content"`

	// The format of the report.
	Format AwrDatabaseSqlReportFormatEnum `mandatory:"false" json:"format,omitempty"`
}

// GetName returns Name
func (m AwrDatabaseSqlReport) GetName() *string {
	return m.Name
}

// GetVersion returns Version
func (m AwrDatabaseSqlReport) GetVersion() *string {
	return m.Version
}

// GetDbQueryTimeInSecs returns DbQueryTimeInSecs
func (m AwrDatabaseSqlReport) GetDbQueryTimeInSecs() *float64 {
	return m.DbQueryTimeInSecs
}

func (m AwrDatabaseSqlReport) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AwrDatabaseSqlReport) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAwrDatabaseSqlReportFormatEnum(string(m.Format)); !ok && m.Format != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Format: %s. Supported values are: %s.", m.Format, strings.Join(GetAwrDatabaseSqlReportFormatEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AwrDatabaseSqlReport) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAwrDatabaseSqlReport AwrDatabaseSqlReport
	s := struct {
		DiscriminatorParam string `json:"awrResultType"`
		MarshalTypeAwrDatabaseSqlReport
	}{
		"AWRDB_SQL_REPORT",
		(MarshalTypeAwrDatabaseSqlReport)(m),
	}

	return json.Marshal(&s)
}

// AwrDatabaseSqlReportFormatEnum Enum with underlying type: string
type AwrDatabaseSqlReportFormatEnum string

// Set of constants representing the allowable values for AwrDatabaseSqlReportFormatEnum
const (
	AwrDatabaseSqlReportFormatHtml AwrDatabaseSqlReportFormatEnum = "HTML"
	AwrDatabaseSqlReportFormatText AwrDatabaseSqlReportFormatEnum = "TEXT"
)

var mappingAwrDatabaseSqlReportFormatEnum = map[string]AwrDatabaseSqlReportFormatEnum{
	"HTML": AwrDatabaseSqlReportFormatHtml,
	"TEXT": AwrDatabaseSqlReportFormatText,
}

var mappingAwrDatabaseSqlReportFormatEnumLowerCase = map[string]AwrDatabaseSqlReportFormatEnum{
	"html": AwrDatabaseSqlReportFormatHtml,
	"text": AwrDatabaseSqlReportFormatText,
}

// GetAwrDatabaseSqlReportFormatEnumValues Enumerates the set of values for AwrDatabaseSqlReportFormatEnum
func GetAwrDatabaseSqlReportFormatEnumValues() []AwrDatabaseSqlReportFormatEnum {
	values := make([]AwrDatabaseSqlReportFormatEnum, 0)
	for _, v := range mappingAwrDatabaseSqlReportFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetAwrDatabaseSqlReportFormatEnumStringValues Enumerates the set of values in String for AwrDatabaseSqlReportFormatEnum
func GetAwrDatabaseSqlReportFormatEnumStringValues() []string {
	return []string{
		"HTML",
		"TEXT",
	}
}

// GetMappingAwrDatabaseSqlReportFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAwrDatabaseSqlReportFormatEnum(val string) (AwrDatabaseSqlReportFormatEnum, bool) {
	enum, ok := mappingAwrDatabaseSqlReportFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
