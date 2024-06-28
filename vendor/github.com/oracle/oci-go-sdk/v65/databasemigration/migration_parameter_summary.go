// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MigrationParameterSummary Migration parameter response object.
type MigrationParameterSummary struct {

	// Parameter name.
	Name *string `mandatory:"true" json:"name"`

	// Parameter data type.
	DataType AdvancedParameterDataTypesEnum `mandatory:"true" json:"dataType"`

	// The combination of source and target databases participating in a migration.
	// Example: ORACLE means the migration is meant for migrating Oracle source and target databases.
	DatabaseCombination DatabaseCombinationEnum `mandatory:"true" json:"databaseCombination"`

	// Parameter display name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Parameter name description.
	Description *string `mandatory:"true" json:"description"`

	// Parameter category name.
	CategoryName *string `mandatory:"true" json:"categoryName"`

	// Parameter category display name.
	CategoryDisplayName *string `mandatory:"true" json:"categoryDisplayName"`

	// Migration Stage.
	MigrationType MigrationTypesEnum `mandatory:"true" json:"migrationType"`

	// Parameter documentation URL link.
	DocUrlLink *string `mandatory:"false" json:"docUrlLink"`

	// Default value for a parameter.
	DefaultValue *string `mandatory:"false" json:"defaultValue"`

	// Parameter minimum value.
	MinValue *float32 `mandatory:"false" json:"minValue"`

	// Parameter maximum value.
	MaxValue *float32 `mandatory:"false" json:"maxValue"`

	// Hint text for parameter value.
	HintText *string `mandatory:"false" json:"hintText"`
}

func (m MigrationParameterSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MigrationParameterSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAdvancedParameterDataTypesEnum(string(m.DataType)); !ok && m.DataType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataType: %s. Supported values are: %s.", m.DataType, strings.Join(GetAdvancedParameterDataTypesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabaseCombinationEnum(string(m.DatabaseCombination)); !ok && m.DatabaseCombination != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseCombination: %s. Supported values are: %s.", m.DatabaseCombination, strings.Join(GetDatabaseCombinationEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMigrationTypesEnum(string(m.MigrationType)); !ok && m.MigrationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MigrationType: %s. Supported values are: %s.", m.MigrationType, strings.Join(GetMigrationTypesEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
