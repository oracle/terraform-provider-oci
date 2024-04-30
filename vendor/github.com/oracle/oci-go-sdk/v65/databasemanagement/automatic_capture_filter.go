// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AutomaticCaptureFilter An automatic capture filter that enables you to capture only those SQL statements
// that you want, and exclude noncritical statements.
type AutomaticCaptureFilter struct {

	// The name of the automatic capture filter.
	// - AUTO_CAPTURE_SQL_TEXT: Search pattern to apply to SQL text.
	// - AUTO_CAPTURE_PARSING_SCHEMA_NAME: Parsing schema to include or exclude for SQL plan management auto capture.
	// - AUTO_CAPTURE_MODULE: Module to include or exclude for SQL plan management auto capture.
	// - AUTO_CAPTURE_ACTION: Action to include or exclude for SQL plan management automatic capture.
	Name AutomaticCaptureFilterNameEnum `mandatory:"false" json:"name,omitempty"`

	// A list of filter values to include.
	ValuesToInclude []string `mandatory:"false" json:"valuesToInclude"`

	// A list of filter values to exclude.
	ValuesToExclude []string `mandatory:"false" json:"valuesToExclude"`

	// The time the filter value was last updated.
	TimeLastModified *common.SDKTime `mandatory:"false" json:"timeLastModified"`

	// The database user who last updated the filter value.
	ModifiedBy *string `mandatory:"false" json:"modifiedBy"`
}

func (m AutomaticCaptureFilter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutomaticCaptureFilter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAutomaticCaptureFilterNameEnum(string(m.Name)); !ok && m.Name != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Name: %s. Supported values are: %s.", m.Name, strings.Join(GetAutomaticCaptureFilterNameEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AutomaticCaptureFilterNameEnum Enum with underlying type: string
type AutomaticCaptureFilterNameEnum string

// Set of constants representing the allowable values for AutomaticCaptureFilterNameEnum
const (
	AutomaticCaptureFilterNameSqlText           AutomaticCaptureFilterNameEnum = "AUTO_CAPTURE_SQL_TEXT"
	AutomaticCaptureFilterNameParsingSchemaName AutomaticCaptureFilterNameEnum = "AUTO_CAPTURE_PARSING_SCHEMA_NAME"
	AutomaticCaptureFilterNameModule            AutomaticCaptureFilterNameEnum = "AUTO_CAPTURE_MODULE"
	AutomaticCaptureFilterNameAction            AutomaticCaptureFilterNameEnum = "AUTO_CAPTURE_ACTION"
)

var mappingAutomaticCaptureFilterNameEnum = map[string]AutomaticCaptureFilterNameEnum{
	"AUTO_CAPTURE_SQL_TEXT":            AutomaticCaptureFilterNameSqlText,
	"AUTO_CAPTURE_PARSING_SCHEMA_NAME": AutomaticCaptureFilterNameParsingSchemaName,
	"AUTO_CAPTURE_MODULE":              AutomaticCaptureFilterNameModule,
	"AUTO_CAPTURE_ACTION":              AutomaticCaptureFilterNameAction,
}

var mappingAutomaticCaptureFilterNameEnumLowerCase = map[string]AutomaticCaptureFilterNameEnum{
	"auto_capture_sql_text":            AutomaticCaptureFilterNameSqlText,
	"auto_capture_parsing_schema_name": AutomaticCaptureFilterNameParsingSchemaName,
	"auto_capture_module":              AutomaticCaptureFilterNameModule,
	"auto_capture_action":              AutomaticCaptureFilterNameAction,
}

// GetAutomaticCaptureFilterNameEnumValues Enumerates the set of values for AutomaticCaptureFilterNameEnum
func GetAutomaticCaptureFilterNameEnumValues() []AutomaticCaptureFilterNameEnum {
	values := make([]AutomaticCaptureFilterNameEnum, 0)
	for _, v := range mappingAutomaticCaptureFilterNameEnum {
		values = append(values, v)
	}
	return values
}

// GetAutomaticCaptureFilterNameEnumStringValues Enumerates the set of values in String for AutomaticCaptureFilterNameEnum
func GetAutomaticCaptureFilterNameEnumStringValues() []string {
	return []string{
		"AUTO_CAPTURE_SQL_TEXT",
		"AUTO_CAPTURE_PARSING_SCHEMA_NAME",
		"AUTO_CAPTURE_MODULE",
		"AUTO_CAPTURE_ACTION",
	}
}

// GetMappingAutomaticCaptureFilterNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutomaticCaptureFilterNameEnum(val string) (AutomaticCaptureFilterNameEnum, bool) {
	enum, ok := mappingAutomaticCaptureFilterNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
