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

// MySqlInitialLoadSettings Optional dump settings
type MySqlInitialLoadSettings struct {

	// MySql Job mode.
	JobMode MySqlInitialLoadSettingsJobModeEnum `mandatory:"true" json:"jobMode"`

	// Enable (true) or disable (false) consistent data dumps by locking the instance for backup during the dump.
	IsConsistent *bool `mandatory:"false" json:"isConsistent"`

	// Include a statement at the start of the dump to set the time zone to UTC.
	IsTzUtc *bool `mandatory:"false" json:"isTzUtc"`

	// Apply the specified requirements for compatibility with MySQL Database Service for all tables in the dump
	// output, altering the dump files as necessary.
	Compatibility []CompatibilityOptionEnum `mandatory:"false" json:"compatibility"`

	// Import the dump even if it contains objects that already exist in the target schema in the MySQL instance.
	IsIgnoreExistingObjects *bool `mandatory:"false" json:"isIgnoreExistingObjects"`

	// The action taken in the event of errors related to GRANT or REVOKE errors.
	IsHandleGrantErrors *bool `mandatory:"false" json:"isHandleGrantErrors"`
}

func (m MySqlInitialLoadSettings) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MySqlInitialLoadSettings) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMySqlInitialLoadSettingsJobModeEnum(string(m.JobMode)); !ok && m.JobMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for JobMode: %s. Supported values are: %s.", m.JobMode, strings.Join(GetMySqlInitialLoadSettingsJobModeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MySqlInitialLoadSettingsJobModeEnum Enum with underlying type: string
type MySqlInitialLoadSettingsJobModeEnum string

// Set of constants representing the allowable values for MySqlInitialLoadSettingsJobModeEnum
const (
	MySqlInitialLoadSettingsJobModeFull   MySqlInitialLoadSettingsJobModeEnum = "FULL"
	MySqlInitialLoadSettingsJobModeSchema MySqlInitialLoadSettingsJobModeEnum = "SCHEMA"
)

var mappingMySqlInitialLoadSettingsJobModeEnum = map[string]MySqlInitialLoadSettingsJobModeEnum{
	"FULL":   MySqlInitialLoadSettingsJobModeFull,
	"SCHEMA": MySqlInitialLoadSettingsJobModeSchema,
}

var mappingMySqlInitialLoadSettingsJobModeEnumLowerCase = map[string]MySqlInitialLoadSettingsJobModeEnum{
	"full":   MySqlInitialLoadSettingsJobModeFull,
	"schema": MySqlInitialLoadSettingsJobModeSchema,
}

// GetMySqlInitialLoadSettingsJobModeEnumValues Enumerates the set of values for MySqlInitialLoadSettingsJobModeEnum
func GetMySqlInitialLoadSettingsJobModeEnumValues() []MySqlInitialLoadSettingsJobModeEnum {
	values := make([]MySqlInitialLoadSettingsJobModeEnum, 0)
	for _, v := range mappingMySqlInitialLoadSettingsJobModeEnum {
		values = append(values, v)
	}
	return values
}

// GetMySqlInitialLoadSettingsJobModeEnumStringValues Enumerates the set of values in String for MySqlInitialLoadSettingsJobModeEnum
func GetMySqlInitialLoadSettingsJobModeEnumStringValues() []string {
	return []string{
		"FULL",
		"SCHEMA",
	}
}

// GetMappingMySqlInitialLoadSettingsJobModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMySqlInitialLoadSettingsJobModeEnum(val string) (MySqlInitialLoadSettingsJobModeEnum, bool) {
	enum, ok := mappingMySqlInitialLoadSettingsJobModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
