// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools Runtime API
//
// Use the Database Tools Runtime API to connect to databases through Database Tools Connections.
//

package databasetoolsruntime

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummary Description of a database API gateway config setting to be provided as an advanced property.
type DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummary struct {

	// A string that uniquely identifies a Database Tools database API gateway config global settings resource.
	Key *string `mandatory:"true" json:"key"`

	// A user-friendly name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The data type of a database API gateway config setting.
	DataType DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDataTypeEnum `mandatory:"true" json:"dataType"`

	// The config types that support this advanced property. The supported types are GLOBAL and POOL.
	ConfigTypes []DatabaseApiGatewayConfigAdvancedPropertyConfigTypeEnum `mandatory:"false" json:"configTypes,omitempty"`

	// The category of the Database Tools database API gateway config global setting.
	CategoryKey *string `mandatory:"false" json:"categoryKey"`

	// A user-friendly name of a category.
	CategoryDisplayName *string `mandatory:"false" json:"categoryDisplayName"`

	// The type of database (as determined by a type of Database Tools connection) to which this setting applies.
	// The advancedProperty applies to all types of Database Tools connection when null. This is only applicable
	// when configTypes includes POOL.
	DatabaseToolsConnectionTypes []DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDatabaseToolsConnectionTypesEnum `mandatory:"false" json:"databaseToolsConnectionTypes,omitempty"`

	// Uniform resource locator (URL) of documentation related to this setting.
	DocumentationUrl *string `mandatory:"false" json:"documentationUrl"`

	// A user-friendly description of a database API gateway config setting.
	Description *string `mandatory:"false" json:"description"`

	// The default value (if applicable) of a database API gateway config setting.
	DefaultValue *string `mandatory:"false" json:"defaultValue"`

	// A minimum numeric value (if applicable) of a database API gateway config setting.
	MinValue *int `mandatory:"false" json:"minValue"`

	// A maximum numeric value (if applicable) of a database API gateway config setting.
	MaxValue *int `mandatory:"false" json:"maxValue"`

	// A list of string values (if applicable) supported by this database API gateway config setting.
	ListOfValues []string `mandatory:"false" json:"listOfValues"`

	// Hint text for a database API gateway config setting.
	HintText *string `mandatory:"false" json:"hintText"`
}

func (m DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDataTypeEnum(string(m.DataType)); !ok && m.DataType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataType: %s. Supported values are: %s.", m.DataType, strings.Join(GetDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDataTypeEnumStringValues(), ",")))
	}

	for _, val := range m.ConfigTypes {
		if _, ok := GetMappingDatabaseApiGatewayConfigAdvancedPropertyConfigTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConfigTypes: %s. Supported values are: %s.", val, strings.Join(GetDatabaseApiGatewayConfigAdvancedPropertyConfigTypeEnumStringValues(), ",")))
		}
	}

	for _, val := range m.DatabaseToolsConnectionTypes {
		if _, ok := GetMappingDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDatabaseToolsConnectionTypesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseToolsConnectionTypes: %s. Supported values are: %s.", val, strings.Join(GetDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDatabaseToolsConnectionTypesEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDataTypeEnum Enum with underlying type: string
type DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDataTypeEnum string

// Set of constants representing the allowable values for DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDataTypeEnum
const (
	DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDataTypeString   DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDataTypeEnum = "STRING"
	DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDataTypeNumber   DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDataTypeEnum = "NUMBER"
	DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDataTypeDuration DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDataTypeEnum = "DURATION"
	DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDataTypeBoolean  DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDataTypeEnum = "BOOLEAN"
)

var mappingDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDataTypeEnum = map[string]DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDataTypeEnum{
	"STRING":   DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDataTypeString,
	"NUMBER":   DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDataTypeNumber,
	"DURATION": DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDataTypeDuration,
	"BOOLEAN":  DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDataTypeBoolean,
}

var mappingDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDataTypeEnumLowerCase = map[string]DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDataTypeEnum{
	"string":   DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDataTypeString,
	"number":   DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDataTypeNumber,
	"duration": DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDataTypeDuration,
	"boolean":  DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDataTypeBoolean,
}

// GetDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDataTypeEnumValues Enumerates the set of values for DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDataTypeEnum
func GetDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDataTypeEnumValues() []DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDataTypeEnum {
	values := make([]DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDataTypeEnum, 0)
	for _, v := range mappingDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDataTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDataTypeEnumStringValues Enumerates the set of values in String for DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDataTypeEnum
func GetDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDataTypeEnumStringValues() []string {
	return []string{
		"STRING",
		"NUMBER",
		"DURATION",
		"BOOLEAN",
	}
}

// GetMappingDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDataTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDataTypeEnum(val string) (DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDataTypeEnum, bool) {
	enum, ok := mappingDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDataTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDatabaseToolsConnectionTypesEnum Enum with underlying type: string
type DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDatabaseToolsConnectionTypesEnum string

// Set of constants representing the allowable values for DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDatabaseToolsConnectionTypesEnum
const (
	DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDatabaseToolsConnectionTypesOracleDatabase DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDatabaseToolsConnectionTypesEnum = "ORACLE_DATABASE"
	DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDatabaseToolsConnectionTypesMysql          DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDatabaseToolsConnectionTypesEnum = "MYSQL"
)

var mappingDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDatabaseToolsConnectionTypesEnum = map[string]DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDatabaseToolsConnectionTypesEnum{
	"ORACLE_DATABASE": DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDatabaseToolsConnectionTypesOracleDatabase,
	"MYSQL":           DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDatabaseToolsConnectionTypesMysql,
}

var mappingDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDatabaseToolsConnectionTypesEnumLowerCase = map[string]DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDatabaseToolsConnectionTypesEnum{
	"oracle_database": DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDatabaseToolsConnectionTypesOracleDatabase,
	"mysql":           DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDatabaseToolsConnectionTypesMysql,
}

// GetDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDatabaseToolsConnectionTypesEnumValues Enumerates the set of values for DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDatabaseToolsConnectionTypesEnum
func GetDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDatabaseToolsConnectionTypesEnumValues() []DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDatabaseToolsConnectionTypesEnum {
	values := make([]DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDatabaseToolsConnectionTypesEnum, 0)
	for _, v := range mappingDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDatabaseToolsConnectionTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDatabaseToolsConnectionTypesEnumStringValues Enumerates the set of values in String for DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDatabaseToolsConnectionTypesEnum
func GetDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDatabaseToolsConnectionTypesEnumStringValues() []string {
	return []string{
		"ORACLE_DATABASE",
		"MYSQL",
	}
}

// GetMappingDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDatabaseToolsConnectionTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDatabaseToolsConnectionTypesEnum(val string) (DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDatabaseToolsConnectionTypesEnum, bool) {
	enum, ok := mappingDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryDatabaseToolsConnectionTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
