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

// AutomaticCaptureFilterDetails The details of a capture filter used to include or exclude SQL statements
// in the initial automatic plan capture.
type AutomaticCaptureFilterDetails struct {

	// The name of the automatic capture filter.
	// - AUTO_CAPTURE_SQL_TEXT: Search pattern to apply to SQL text.
	// - AUTO_CAPTURE_PARSING_SCHEMA_NAME: Parsing schema to include or exclude for SQL plan management auto capture.
	// - AUTO_CAPTURE_MODULE: Module to include or exclude for SQL plan management auto capture.
	// - AUTO_CAPTURE_ACTION: Action to include or exclude for SQL plan management automatic capture.
	Name AutomaticCaptureFilterDetailsNameEnum `mandatory:"true" json:"name"`

	// A list of filter values to include.
	ValuesToInclude []string `mandatory:"false" json:"valuesToInclude"`

	// A list of filter values to exclude.
	ValuesToExclude []string `mandatory:"false" json:"valuesToExclude"`
}

func (m AutomaticCaptureFilterDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutomaticCaptureFilterDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAutomaticCaptureFilterDetailsNameEnum(string(m.Name)); !ok && m.Name != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Name: %s. Supported values are: %s.", m.Name, strings.Join(GetAutomaticCaptureFilterDetailsNameEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AutomaticCaptureFilterDetailsNameEnum Enum with underlying type: string
type AutomaticCaptureFilterDetailsNameEnum string

// Set of constants representing the allowable values for AutomaticCaptureFilterDetailsNameEnum
const (
	AutomaticCaptureFilterDetailsNameSqlText           AutomaticCaptureFilterDetailsNameEnum = "AUTO_CAPTURE_SQL_TEXT"
	AutomaticCaptureFilterDetailsNameParsingSchemaName AutomaticCaptureFilterDetailsNameEnum = "AUTO_CAPTURE_PARSING_SCHEMA_NAME"
	AutomaticCaptureFilterDetailsNameModule            AutomaticCaptureFilterDetailsNameEnum = "AUTO_CAPTURE_MODULE"
	AutomaticCaptureFilterDetailsNameAction            AutomaticCaptureFilterDetailsNameEnum = "AUTO_CAPTURE_ACTION"
)

var mappingAutomaticCaptureFilterDetailsNameEnum = map[string]AutomaticCaptureFilterDetailsNameEnum{
	"AUTO_CAPTURE_SQL_TEXT":            AutomaticCaptureFilterDetailsNameSqlText,
	"AUTO_CAPTURE_PARSING_SCHEMA_NAME": AutomaticCaptureFilterDetailsNameParsingSchemaName,
	"AUTO_CAPTURE_MODULE":              AutomaticCaptureFilterDetailsNameModule,
	"AUTO_CAPTURE_ACTION":              AutomaticCaptureFilterDetailsNameAction,
}

var mappingAutomaticCaptureFilterDetailsNameEnumLowerCase = map[string]AutomaticCaptureFilterDetailsNameEnum{
	"auto_capture_sql_text":            AutomaticCaptureFilterDetailsNameSqlText,
	"auto_capture_parsing_schema_name": AutomaticCaptureFilterDetailsNameParsingSchemaName,
	"auto_capture_module":              AutomaticCaptureFilterDetailsNameModule,
	"auto_capture_action":              AutomaticCaptureFilterDetailsNameAction,
}

// GetAutomaticCaptureFilterDetailsNameEnumValues Enumerates the set of values for AutomaticCaptureFilterDetailsNameEnum
func GetAutomaticCaptureFilterDetailsNameEnumValues() []AutomaticCaptureFilterDetailsNameEnum {
	values := make([]AutomaticCaptureFilterDetailsNameEnum, 0)
	for _, v := range mappingAutomaticCaptureFilterDetailsNameEnum {
		values = append(values, v)
	}
	return values
}

// GetAutomaticCaptureFilterDetailsNameEnumStringValues Enumerates the set of values in String for AutomaticCaptureFilterDetailsNameEnum
func GetAutomaticCaptureFilterDetailsNameEnumStringValues() []string {
	return []string{
		"AUTO_CAPTURE_SQL_TEXT",
		"AUTO_CAPTURE_PARSING_SCHEMA_NAME",
		"AUTO_CAPTURE_MODULE",
		"AUTO_CAPTURE_ACTION",
	}
}

// GetMappingAutomaticCaptureFilterDetailsNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutomaticCaptureFilterDetailsNameEnum(val string) (AutomaticCaptureFilterDetailsNameEnum, bool) {
	enum, ok := mappingAutomaticCaptureFilterDetailsNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
