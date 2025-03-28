// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

// GetName returns Name
func (m AwrDbReport) GetName() *string {
	return m.Name
}

// GetVersion returns Version
func (m AwrDbReport) GetVersion() *string {
	return m.Version
}

// GetQueryKey returns QueryKey
func (m AwrDbReport) GetQueryKey() *string {
	return m.QueryKey
}

// GetDbQueryTimeInSecs returns DbQueryTimeInSecs
func (m AwrDbReport) GetDbQueryTimeInSecs() *float64 {
	return m.DbQueryTimeInSecs
}

func (m AwrDbReport) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AwrDbReport) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAwrDbReportFormatEnum(string(m.Format)); !ok && m.Format != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Format: %s. Supported values are: %s.", m.Format, strings.Join(GetAwrDbReportFormatEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingAwrDbReportFormatEnum = map[string]AwrDbReportFormatEnum{
	"HTML": AwrDbReportFormatHtml,
	"TEXT": AwrDbReportFormatText,
	"XML":  AwrDbReportFormatXml,
}

var mappingAwrDbReportFormatEnumLowerCase = map[string]AwrDbReportFormatEnum{
	"html": AwrDbReportFormatHtml,
	"text": AwrDbReportFormatText,
	"xml":  AwrDbReportFormatXml,
}

// GetAwrDbReportFormatEnumValues Enumerates the set of values for AwrDbReportFormatEnum
func GetAwrDbReportFormatEnumValues() []AwrDbReportFormatEnum {
	values := make([]AwrDbReportFormatEnum, 0)
	for _, v := range mappingAwrDbReportFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetAwrDbReportFormatEnumStringValues Enumerates the set of values in String for AwrDbReportFormatEnum
func GetAwrDbReportFormatEnumStringValues() []string {
	return []string{
		"HTML",
		"TEXT",
		"XML",
	}
}

// GetMappingAwrDbReportFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAwrDbReportFormatEnum(val string) (AwrDbReportFormatEnum, bool) {
	enum, ok := mappingAwrDbReportFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
