// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// SchemaDriftConfig The configuration for handling schema drift in a Source or Target operator.
type SchemaDriftConfig struct {

	// The setting for how to handle extra columns/fields.  NULL_FILLUP means that nulls will be loaded into the target for extra columns.
	ExtraColumnHandling SchemaDriftConfigExtraColumnHandlingEnum `mandatory:"false" json:"extraColumnHandling,omitempty"`

	// The setting for how to handle missing columns/fields.  NULL_SELECT means that null values will be selected from the source for missing columns.
	MissingColumnHandling SchemaDriftConfigMissingColumnHandlingEnum `mandatory:"false" json:"missingColumnHandling,omitempty"`

	// The setting for how to handle columns/fields with changed data types.
	DataTypeChangeHandling SchemaDriftConfigDataTypeChangeHandlingEnum `mandatory:"false" json:"dataTypeChangeHandling,omitempty"`

	// If true, display a validation warning for schema changes, even if they are allowed.
	IsValidationWarningIfAllowed *bool `mandatory:"false" json:"isValidationWarningIfAllowed"`
}

func (m SchemaDriftConfig) String() string {
	return common.PointerString(m)
}

// SchemaDriftConfigExtraColumnHandlingEnum Enum with underlying type: string
type SchemaDriftConfigExtraColumnHandlingEnum string

// Set of constants representing the allowable values for SchemaDriftConfigExtraColumnHandlingEnum
const (
	SchemaDriftConfigExtraColumnHandlingAllow      SchemaDriftConfigExtraColumnHandlingEnum = "ALLOW"
	SchemaDriftConfigExtraColumnHandlingNullFillup SchemaDriftConfigExtraColumnHandlingEnum = "NULL_FILLUP"
	SchemaDriftConfigExtraColumnHandlingDoNotAllow SchemaDriftConfigExtraColumnHandlingEnum = "DO_NOT_ALLOW"
)

var mappingSchemaDriftConfigExtraColumnHandling = map[string]SchemaDriftConfigExtraColumnHandlingEnum{
	"ALLOW":        SchemaDriftConfigExtraColumnHandlingAllow,
	"NULL_FILLUP":  SchemaDriftConfigExtraColumnHandlingNullFillup,
	"DO_NOT_ALLOW": SchemaDriftConfigExtraColumnHandlingDoNotAllow,
}

// GetSchemaDriftConfigExtraColumnHandlingEnumValues Enumerates the set of values for SchemaDriftConfigExtraColumnHandlingEnum
func GetSchemaDriftConfigExtraColumnHandlingEnumValues() []SchemaDriftConfigExtraColumnHandlingEnum {
	values := make([]SchemaDriftConfigExtraColumnHandlingEnum, 0)
	for _, v := range mappingSchemaDriftConfigExtraColumnHandling {
		values = append(values, v)
	}
	return values
}

// SchemaDriftConfigMissingColumnHandlingEnum Enum with underlying type: string
type SchemaDriftConfigMissingColumnHandlingEnum string

// Set of constants representing the allowable values for SchemaDriftConfigMissingColumnHandlingEnum
const (
	SchemaDriftConfigMissingColumnHandlingAllow      SchemaDriftConfigMissingColumnHandlingEnum = "ALLOW"
	SchemaDriftConfigMissingColumnHandlingNullSelect SchemaDriftConfigMissingColumnHandlingEnum = "NULL_SELECT"
	SchemaDriftConfigMissingColumnHandlingDoNotAllow SchemaDriftConfigMissingColumnHandlingEnum = "DO_NOT_ALLOW"
)

var mappingSchemaDriftConfigMissingColumnHandling = map[string]SchemaDriftConfigMissingColumnHandlingEnum{
	"ALLOW":        SchemaDriftConfigMissingColumnHandlingAllow,
	"NULL_SELECT":  SchemaDriftConfigMissingColumnHandlingNullSelect,
	"DO_NOT_ALLOW": SchemaDriftConfigMissingColumnHandlingDoNotAllow,
}

// GetSchemaDriftConfigMissingColumnHandlingEnumValues Enumerates the set of values for SchemaDriftConfigMissingColumnHandlingEnum
func GetSchemaDriftConfigMissingColumnHandlingEnumValues() []SchemaDriftConfigMissingColumnHandlingEnum {
	values := make([]SchemaDriftConfigMissingColumnHandlingEnum, 0)
	for _, v := range mappingSchemaDriftConfigMissingColumnHandling {
		values = append(values, v)
	}
	return values
}

// SchemaDriftConfigDataTypeChangeHandlingEnum Enum with underlying type: string
type SchemaDriftConfigDataTypeChangeHandlingEnum string

// Set of constants representing the allowable values for SchemaDriftConfigDataTypeChangeHandlingEnum
const (
	SchemaDriftConfigDataTypeChangeHandlingAllow            SchemaDriftConfigDataTypeChangeHandlingEnum = "ALLOW"
	SchemaDriftConfigDataTypeChangeHandlingDoCastIfPossible SchemaDriftConfigDataTypeChangeHandlingEnum = "DO_CAST_IF_POSSIBLE"
	SchemaDriftConfigDataTypeChangeHandlingDoNotAllow       SchemaDriftConfigDataTypeChangeHandlingEnum = "DO_NOT_ALLOW"
)

var mappingSchemaDriftConfigDataTypeChangeHandling = map[string]SchemaDriftConfigDataTypeChangeHandlingEnum{
	"ALLOW":               SchemaDriftConfigDataTypeChangeHandlingAllow,
	"DO_CAST_IF_POSSIBLE": SchemaDriftConfigDataTypeChangeHandlingDoCastIfPossible,
	"DO_NOT_ALLOW":        SchemaDriftConfigDataTypeChangeHandlingDoNotAllow,
}

// GetSchemaDriftConfigDataTypeChangeHandlingEnumValues Enumerates the set of values for SchemaDriftConfigDataTypeChangeHandlingEnum
func GetSchemaDriftConfigDataTypeChangeHandlingEnumValues() []SchemaDriftConfigDataTypeChangeHandlingEnum {
	values := make([]SchemaDriftConfigDataTypeChangeHandlingEnum, 0)
	for _, v := range mappingSchemaDriftConfigDataTypeChangeHandling {
		values = append(values, v)
	}
	return values
}
