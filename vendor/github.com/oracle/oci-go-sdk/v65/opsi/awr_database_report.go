// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AwrDatabaseReport The result of the AWR report.
type AwrDatabaseReport struct {

	// The name of the query result.
	Name *string `mandatory:"true" json:"name"`

	// The version of the query result.
	Version *string `mandatory:"false" json:"version"`

	// The time taken to query the database tier (in seconds).
	DbQueryTimeInSecs *float64 `mandatory:"false" json:"dbQueryTimeInSecs"`

	// The content of the report.
	Content *string `mandatory:"false" json:"content"`

	// The format of the report.
	Format AwrDatabaseReportFormatEnum `mandatory:"false" json:"format,omitempty"`
}

// GetName returns Name
func (m AwrDatabaseReport) GetName() *string {
	return m.Name
}

// GetVersion returns Version
func (m AwrDatabaseReport) GetVersion() *string {
	return m.Version
}

// GetDbQueryTimeInSecs returns DbQueryTimeInSecs
func (m AwrDatabaseReport) GetDbQueryTimeInSecs() *float64 {
	return m.DbQueryTimeInSecs
}

func (m AwrDatabaseReport) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AwrDatabaseReport) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAwrDatabaseReportFormatEnum(string(m.Format)); !ok && m.Format != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Format: %s. Supported values are: %s.", m.Format, strings.Join(GetAwrDatabaseReportFormatEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AwrDatabaseReport) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAwrDatabaseReport AwrDatabaseReport
	s := struct {
		DiscriminatorParam string `json:"awrResultType"`
		MarshalTypeAwrDatabaseReport
	}{
		"AWRDB_DB_REPORT",
		(MarshalTypeAwrDatabaseReport)(m),
	}

	return json.Marshal(&s)
}

// AwrDatabaseReportFormatEnum Enum with underlying type: string
type AwrDatabaseReportFormatEnum string

// Set of constants representing the allowable values for AwrDatabaseReportFormatEnum
const (
	AwrDatabaseReportFormatHtml AwrDatabaseReportFormatEnum = "HTML"
	AwrDatabaseReportFormatText AwrDatabaseReportFormatEnum = "TEXT"
	AwrDatabaseReportFormatXml  AwrDatabaseReportFormatEnum = "XML"
)

var mappingAwrDatabaseReportFormatEnum = map[string]AwrDatabaseReportFormatEnum{
	"HTML": AwrDatabaseReportFormatHtml,
	"TEXT": AwrDatabaseReportFormatText,
	"XML":  AwrDatabaseReportFormatXml,
}

var mappingAwrDatabaseReportFormatEnumLowerCase = map[string]AwrDatabaseReportFormatEnum{
	"html": AwrDatabaseReportFormatHtml,
	"text": AwrDatabaseReportFormatText,
	"xml":  AwrDatabaseReportFormatXml,
}

// GetAwrDatabaseReportFormatEnumValues Enumerates the set of values for AwrDatabaseReportFormatEnum
func GetAwrDatabaseReportFormatEnumValues() []AwrDatabaseReportFormatEnum {
	values := make([]AwrDatabaseReportFormatEnum, 0)
	for _, v := range mappingAwrDatabaseReportFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetAwrDatabaseReportFormatEnumStringValues Enumerates the set of values in String for AwrDatabaseReportFormatEnum
func GetAwrDatabaseReportFormatEnumStringValues() []string {
	return []string{
		"HTML",
		"TEXT",
		"XML",
	}
}

// GetMappingAwrDatabaseReportFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAwrDatabaseReportFormatEnum(val string) (AwrDatabaseReportFormatEnum, bool) {
	enum, ok := mappingAwrDatabaseReportFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
