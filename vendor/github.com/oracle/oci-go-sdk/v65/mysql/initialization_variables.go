// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InitializationVariables User-defined service variables set only at DB system initialization. These variables cannot be changed later at runtime.
type InitializationVariables struct {

	//
	// Represents the MySQL server system variable lower_case_table_names (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_lower_case_table_names).
	// lowerCaseTableNames controls case-sensitivity of tables and schema names and how they are stored in the DB System.
	// Valid values are:
	//   - CASE_SENSITIVE - (default) Table and schema name comparisons are case-sensitive and stored as specified. (lower_case_table_names=0)
	//   - CASE_INSENSITIVE_LOWERCASE - Table and schema name comparisons are not case-sensitive and stored in lowercase. (lower_case_table_names=1)
	LowerCaseTableNames InitializationVariablesLowerCaseTableNamesEnum `mandatory:"false" json:"lowerCaseTableNames,omitempty"`
}

func (m InitializationVariables) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InitializationVariables) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingInitializationVariablesLowerCaseTableNamesEnum(string(m.LowerCaseTableNames)); !ok && m.LowerCaseTableNames != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LowerCaseTableNames: %s. Supported values are: %s.", m.LowerCaseTableNames, strings.Join(GetInitializationVariablesLowerCaseTableNamesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InitializationVariablesLowerCaseTableNamesEnum Enum with underlying type: string
type InitializationVariablesLowerCaseTableNamesEnum string

// Set of constants representing the allowable values for InitializationVariablesLowerCaseTableNamesEnum
const (
	InitializationVariablesLowerCaseTableNamesSensitive            InitializationVariablesLowerCaseTableNamesEnum = "CASE_SENSITIVE"
	InitializationVariablesLowerCaseTableNamesInsensitiveLowercase InitializationVariablesLowerCaseTableNamesEnum = "CASE_INSENSITIVE_LOWERCASE"
)

var mappingInitializationVariablesLowerCaseTableNamesEnum = map[string]InitializationVariablesLowerCaseTableNamesEnum{
	"CASE_SENSITIVE":             InitializationVariablesLowerCaseTableNamesSensitive,
	"CASE_INSENSITIVE_LOWERCASE": InitializationVariablesLowerCaseTableNamesInsensitiveLowercase,
}

var mappingInitializationVariablesLowerCaseTableNamesEnumLowerCase = map[string]InitializationVariablesLowerCaseTableNamesEnum{
	"case_sensitive":             InitializationVariablesLowerCaseTableNamesSensitive,
	"case_insensitive_lowercase": InitializationVariablesLowerCaseTableNamesInsensitiveLowercase,
}

// GetInitializationVariablesLowerCaseTableNamesEnumValues Enumerates the set of values for InitializationVariablesLowerCaseTableNamesEnum
func GetInitializationVariablesLowerCaseTableNamesEnumValues() []InitializationVariablesLowerCaseTableNamesEnum {
	values := make([]InitializationVariablesLowerCaseTableNamesEnum, 0)
	for _, v := range mappingInitializationVariablesLowerCaseTableNamesEnum {
		values = append(values, v)
	}
	return values
}

// GetInitializationVariablesLowerCaseTableNamesEnumStringValues Enumerates the set of values in String for InitializationVariablesLowerCaseTableNamesEnum
func GetInitializationVariablesLowerCaseTableNamesEnumStringValues() []string {
	return []string{
		"CASE_SENSITIVE",
		"CASE_INSENSITIVE_LOWERCASE",
	}
}

// GetMappingInitializationVariablesLowerCaseTableNamesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInitializationVariablesLowerCaseTableNamesEnum(val string) (InitializationVariablesLowerCaseTableNamesEnum, bool) {
	enum, ok := mappingInitializationVariablesLowerCaseTableNamesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
