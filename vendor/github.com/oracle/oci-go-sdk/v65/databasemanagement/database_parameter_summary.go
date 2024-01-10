// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatabaseParameterSummary A summary of the database parameter.
type DatabaseParameterSummary struct {

	// The parameter name.
	Name *string `mandatory:"true" json:"name"`

	// The parameter type.
	Type DatabaseParameterSummaryTypeEnum `mandatory:"true" json:"type"`

	// The parameter value.
	Value *string `mandatory:"true" json:"value"`

	// The parameter value in a user-friendly format. For example, if the `value` property shows the value 262144 for a big integer parameter, then the `displayValue` property will show the value 256K.
	DisplayValue *string `mandatory:"true" json:"displayValue"`

	// The parameter number.
	Number *float32 `mandatory:"false" json:"number"`

	// Indicates whether the parameter is set to the default value (`TRUE`) or the parameter value was specified in the parameter file (`FALSE`).
	IsDefault *bool `mandatory:"false" json:"isDefault"`

	// Indicates whether the parameter can be changed with `ALTER SESSION` (`TRUE`) or not (`FALSE`)
	IsSessionModifiable *bool `mandatory:"false" json:"isSessionModifiable"`

	// Indicates whether the parameter can be changed with `ALTER SYSTEM` and when the change takes effect:
	// - IMMEDIATE: Parameter can be changed with `ALTER SYSTEM` regardless of the type of parameter file used to start the instance. The change takes effect immediately.
	// - DEFERRED: Parameter can be changed with `ALTER SYSTEM` regardless of the type of parameter file used to start the instance. The change takes effect in subsequent sessions.
	// - FALSE: Parameter cannot be changed with `ALTER SYSTEM` unless a server parameter file was used to start the instance. The change takes effect in subsequent instances.
	IsSystemModifiable DatabaseParameterSummaryIsSystemModifiableEnum `mandatory:"false" json:"isSystemModifiable,omitempty"`

	// Indicates whether the parameter can be modified on a per-PDB basis (`TRUE`) or not (`FALSE`). In a non-CDB, the value of this property is `null`.
	IsPdbModifiable *bool `mandatory:"false" json:"isPdbModifiable"`

	// For parameters that can be changed with `ALTER SYSTEM`, indicates whether the value of the parameter can be different for every instance (`TRUE`) or whether the parameter must have the same value for all Real Application Clusters instances (`FALSE`). For other parameters, this is always `FALSE`.
	IsInstanceModifiable *bool `mandatory:"false" json:"isInstanceModifiable"`

	// Indicates how the parameter was modified. If an `ALTER SYSTEM` was performed, the value will be `MODIFIED`.
	IsModified DatabaseParameterSummaryIsModifiedEnum `mandatory:"false" json:"isModified,omitempty"`

	// Indicates whether Oracle adjusted the input value to a more suitable value.
	IsAdjusted *bool `mandatory:"false" json:"isAdjusted"`

	// Indicates whether the parameter has been deprecated (`TRUE`) or not (`FALSE`).
	IsDeprecated *bool `mandatory:"false" json:"isDeprecated"`

	// Indicates whether the parameter is a basic parameter (`TRUE`) or not (`FALSE`).
	IsBasic *bool `mandatory:"false" json:"isBasic"`

	// The description of the parameter.
	Description *string `mandatory:"false" json:"description"`

	// The position (ordinal number) of the parameter value. Useful only for parameters whose values are lists of strings.
	Ordinal *float32 `mandatory:"false" json:"ordinal"`

	// The comments associated with the most recent update.
	UpdateComment *string `mandatory:"false" json:"updateComment"`

	// The ID of the database container to which the data pertains.
	// Possible values include:
	// - `0`: This value is used for data that pertain to the entire CDB. This value is also used for data in non-CDBs.
	// - `1`: This value is used for data that pertain to only the root container.
	// - `n`: Where n is the applicable container ID for the data.
	ContainerId *float32 `mandatory:"false" json:"containerId"`

	// The parameter category.
	Category *string `mandatory:"false" json:"category"`

	// Applicable in case of Oracle Real Application Clusters (Oracle RAC) databases.
	// A `UNIQUE` parameter is one which is unique to each Oracle Real Application
	// Clusters (Oracle RAC) instance. For example, the parameter `INSTANCE_NUMBER`
	// must have different values in each instance. An `IDENTICAL` parameter must
	// have the same value for every instance. For example, the parameter
	// `DB_BLOCK_SIZE` must have the same value in all instances.
	Constraint DatabaseParameterSummaryConstraintEnum `mandatory:"false" json:"constraint,omitempty"`

	// The database instance SID for which the parameter is defined.
	Sid *string `mandatory:"false" json:"sid"`

	// Indicates whether the parameter was specified in the server parameter file (`TRUE`) or not (`FALSE`). Applicable only when the parameter source is `SPFILE`.
	IsSpecified *bool `mandatory:"false" json:"isSpecified"`

	// A list of allowed values for this parameter.
	AllowedValues []AllowedParameterValue `mandatory:"false" json:"allowedValues"`
}

func (m DatabaseParameterSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseParameterSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDatabaseParameterSummaryTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetDatabaseParameterSummaryTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDatabaseParameterSummaryIsSystemModifiableEnum(string(m.IsSystemModifiable)); !ok && m.IsSystemModifiable != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IsSystemModifiable: %s. Supported values are: %s.", m.IsSystemModifiable, strings.Join(GetDatabaseParameterSummaryIsSystemModifiableEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabaseParameterSummaryIsModifiedEnum(string(m.IsModified)); !ok && m.IsModified != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IsModified: %s. Supported values are: %s.", m.IsModified, strings.Join(GetDatabaseParameterSummaryIsModifiedEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabaseParameterSummaryConstraintEnum(string(m.Constraint)); !ok && m.Constraint != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Constraint: %s. Supported values are: %s.", m.Constraint, strings.Join(GetDatabaseParameterSummaryConstraintEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseParameterSummaryTypeEnum Enum with underlying type: string
type DatabaseParameterSummaryTypeEnum string

// Set of constants representing the allowable values for DatabaseParameterSummaryTypeEnum
const (
	DatabaseParameterSummaryTypeBoolean    DatabaseParameterSummaryTypeEnum = "BOOLEAN"
	DatabaseParameterSummaryTypeString     DatabaseParameterSummaryTypeEnum = "STRING"
	DatabaseParameterSummaryTypeInteger    DatabaseParameterSummaryTypeEnum = "INTEGER"
	DatabaseParameterSummaryTypeFilename   DatabaseParameterSummaryTypeEnum = "FILENAME"
	DatabaseParameterSummaryTypeBigInteger DatabaseParameterSummaryTypeEnum = "BIG_INTEGER"
	DatabaseParameterSummaryTypeReserved   DatabaseParameterSummaryTypeEnum = "RESERVED"
)

var mappingDatabaseParameterSummaryTypeEnum = map[string]DatabaseParameterSummaryTypeEnum{
	"BOOLEAN":     DatabaseParameterSummaryTypeBoolean,
	"STRING":      DatabaseParameterSummaryTypeString,
	"INTEGER":     DatabaseParameterSummaryTypeInteger,
	"FILENAME":    DatabaseParameterSummaryTypeFilename,
	"BIG_INTEGER": DatabaseParameterSummaryTypeBigInteger,
	"RESERVED":    DatabaseParameterSummaryTypeReserved,
}

var mappingDatabaseParameterSummaryTypeEnumLowerCase = map[string]DatabaseParameterSummaryTypeEnum{
	"boolean":     DatabaseParameterSummaryTypeBoolean,
	"string":      DatabaseParameterSummaryTypeString,
	"integer":     DatabaseParameterSummaryTypeInteger,
	"filename":    DatabaseParameterSummaryTypeFilename,
	"big_integer": DatabaseParameterSummaryTypeBigInteger,
	"reserved":    DatabaseParameterSummaryTypeReserved,
}

// GetDatabaseParameterSummaryTypeEnumValues Enumerates the set of values for DatabaseParameterSummaryTypeEnum
func GetDatabaseParameterSummaryTypeEnumValues() []DatabaseParameterSummaryTypeEnum {
	values := make([]DatabaseParameterSummaryTypeEnum, 0)
	for _, v := range mappingDatabaseParameterSummaryTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseParameterSummaryTypeEnumStringValues Enumerates the set of values in String for DatabaseParameterSummaryTypeEnum
func GetDatabaseParameterSummaryTypeEnumStringValues() []string {
	return []string{
		"BOOLEAN",
		"STRING",
		"INTEGER",
		"FILENAME",
		"BIG_INTEGER",
		"RESERVED",
	}
}

// GetMappingDatabaseParameterSummaryTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseParameterSummaryTypeEnum(val string) (DatabaseParameterSummaryTypeEnum, bool) {
	enum, ok := mappingDatabaseParameterSummaryTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DatabaseParameterSummaryIsSystemModifiableEnum Enum with underlying type: string
type DatabaseParameterSummaryIsSystemModifiableEnum string

// Set of constants representing the allowable values for DatabaseParameterSummaryIsSystemModifiableEnum
const (
	DatabaseParameterSummaryIsSystemModifiableImmediate DatabaseParameterSummaryIsSystemModifiableEnum = "IMMEDIATE"
	DatabaseParameterSummaryIsSystemModifiableDeferred  DatabaseParameterSummaryIsSystemModifiableEnum = "DEFERRED"
	DatabaseParameterSummaryIsSystemModifiableFalse     DatabaseParameterSummaryIsSystemModifiableEnum = "FALSE"
)

var mappingDatabaseParameterSummaryIsSystemModifiableEnum = map[string]DatabaseParameterSummaryIsSystemModifiableEnum{
	"IMMEDIATE": DatabaseParameterSummaryIsSystemModifiableImmediate,
	"DEFERRED":  DatabaseParameterSummaryIsSystemModifiableDeferred,
	"FALSE":     DatabaseParameterSummaryIsSystemModifiableFalse,
}

var mappingDatabaseParameterSummaryIsSystemModifiableEnumLowerCase = map[string]DatabaseParameterSummaryIsSystemModifiableEnum{
	"immediate": DatabaseParameterSummaryIsSystemModifiableImmediate,
	"deferred":  DatabaseParameterSummaryIsSystemModifiableDeferred,
	"false":     DatabaseParameterSummaryIsSystemModifiableFalse,
}

// GetDatabaseParameterSummaryIsSystemModifiableEnumValues Enumerates the set of values for DatabaseParameterSummaryIsSystemModifiableEnum
func GetDatabaseParameterSummaryIsSystemModifiableEnumValues() []DatabaseParameterSummaryIsSystemModifiableEnum {
	values := make([]DatabaseParameterSummaryIsSystemModifiableEnum, 0)
	for _, v := range mappingDatabaseParameterSummaryIsSystemModifiableEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseParameterSummaryIsSystemModifiableEnumStringValues Enumerates the set of values in String for DatabaseParameterSummaryIsSystemModifiableEnum
func GetDatabaseParameterSummaryIsSystemModifiableEnumStringValues() []string {
	return []string{
		"IMMEDIATE",
		"DEFERRED",
		"FALSE",
	}
}

// GetMappingDatabaseParameterSummaryIsSystemModifiableEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseParameterSummaryIsSystemModifiableEnum(val string) (DatabaseParameterSummaryIsSystemModifiableEnum, bool) {
	enum, ok := mappingDatabaseParameterSummaryIsSystemModifiableEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DatabaseParameterSummaryIsModifiedEnum Enum with underlying type: string
type DatabaseParameterSummaryIsModifiedEnum string

// Set of constants representing the allowable values for DatabaseParameterSummaryIsModifiedEnum
const (
	DatabaseParameterSummaryIsModifiedModified DatabaseParameterSummaryIsModifiedEnum = "MODIFIED"
	DatabaseParameterSummaryIsModifiedFalse    DatabaseParameterSummaryIsModifiedEnum = "FALSE"
)

var mappingDatabaseParameterSummaryIsModifiedEnum = map[string]DatabaseParameterSummaryIsModifiedEnum{
	"MODIFIED": DatabaseParameterSummaryIsModifiedModified,
	"FALSE":    DatabaseParameterSummaryIsModifiedFalse,
}

var mappingDatabaseParameterSummaryIsModifiedEnumLowerCase = map[string]DatabaseParameterSummaryIsModifiedEnum{
	"modified": DatabaseParameterSummaryIsModifiedModified,
	"false":    DatabaseParameterSummaryIsModifiedFalse,
}

// GetDatabaseParameterSummaryIsModifiedEnumValues Enumerates the set of values for DatabaseParameterSummaryIsModifiedEnum
func GetDatabaseParameterSummaryIsModifiedEnumValues() []DatabaseParameterSummaryIsModifiedEnum {
	values := make([]DatabaseParameterSummaryIsModifiedEnum, 0)
	for _, v := range mappingDatabaseParameterSummaryIsModifiedEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseParameterSummaryIsModifiedEnumStringValues Enumerates the set of values in String for DatabaseParameterSummaryIsModifiedEnum
func GetDatabaseParameterSummaryIsModifiedEnumStringValues() []string {
	return []string{
		"MODIFIED",
		"FALSE",
	}
}

// GetMappingDatabaseParameterSummaryIsModifiedEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseParameterSummaryIsModifiedEnum(val string) (DatabaseParameterSummaryIsModifiedEnum, bool) {
	enum, ok := mappingDatabaseParameterSummaryIsModifiedEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DatabaseParameterSummaryConstraintEnum Enum with underlying type: string
type DatabaseParameterSummaryConstraintEnum string

// Set of constants representing the allowable values for DatabaseParameterSummaryConstraintEnum
const (
	DatabaseParameterSummaryConstraintUnique    DatabaseParameterSummaryConstraintEnum = "UNIQUE"
	DatabaseParameterSummaryConstraintIdentical DatabaseParameterSummaryConstraintEnum = "IDENTICAL"
	DatabaseParameterSummaryConstraintNone      DatabaseParameterSummaryConstraintEnum = "NONE"
)

var mappingDatabaseParameterSummaryConstraintEnum = map[string]DatabaseParameterSummaryConstraintEnum{
	"UNIQUE":    DatabaseParameterSummaryConstraintUnique,
	"IDENTICAL": DatabaseParameterSummaryConstraintIdentical,
	"NONE":      DatabaseParameterSummaryConstraintNone,
}

var mappingDatabaseParameterSummaryConstraintEnumLowerCase = map[string]DatabaseParameterSummaryConstraintEnum{
	"unique":    DatabaseParameterSummaryConstraintUnique,
	"identical": DatabaseParameterSummaryConstraintIdentical,
	"none":      DatabaseParameterSummaryConstraintNone,
}

// GetDatabaseParameterSummaryConstraintEnumValues Enumerates the set of values for DatabaseParameterSummaryConstraintEnum
func GetDatabaseParameterSummaryConstraintEnumValues() []DatabaseParameterSummaryConstraintEnum {
	values := make([]DatabaseParameterSummaryConstraintEnum, 0)
	for _, v := range mappingDatabaseParameterSummaryConstraintEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseParameterSummaryConstraintEnumStringValues Enumerates the set of values in String for DatabaseParameterSummaryConstraintEnum
func GetDatabaseParameterSummaryConstraintEnumStringValues() []string {
	return []string{
		"UNIQUE",
		"IDENTICAL",
		"NONE",
	}
}

// GetMappingDatabaseParameterSummaryConstraintEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseParameterSummaryConstraintEnum(val string) (DatabaseParameterSummaryConstraintEnum, bool) {
	enum, ok := mappingDatabaseParameterSummaryConstraintEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
