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

// MySqlConfigurationDataSummary The configuration variables for a MySQL Database.
type MySqlConfigurationDataSummary struct {

	// The name of the configuration variable
	Name *string `mandatory:"true" json:"name"`

	// The value of the variable.
	Value *string `mandatory:"true" json:"value"`

	// The source from which the variable was most recently set.
	Source MySqlConfigurationDataSummarySourceEnum `mandatory:"true" json:"source"`

	// The minimum value of the variable.
	MinValue *float32 `mandatory:"true" json:"minValue"`

	// The maximum value of the variable.
	MaxValue *float32 `mandatory:"true" json:"maxValue"`

	// The type of variable.
	Type *string `mandatory:"true" json:"type"`

	// The default value of the variable.
	DefaultValue *string `mandatory:"true" json:"defaultValue"`

	// The time when the value of the variable was set.
	TimeSet *common.SDKTime `mandatory:"true" json:"timeSet"`

	// The host from where the value of the variable was set. This is empty for a MySQL Database System.
	HostSet *string `mandatory:"true" json:"hostSet"`

	// The user who sets the value of the variable. This is empty for a MySQL Database System.
	UserSet *string `mandatory:"true" json:"userSet"`

	// Indicates whether the variable can be set dynamically or not.
	IsDynamic *bool `mandatory:"true" json:"isDynamic"`

	// Indicates whether the variable is set at server startup.
	IsInit *bool `mandatory:"true" json:"isInit"`

	// Indicates whether the variable is configurable.
	IsConfigurable *bool `mandatory:"true" json:"isConfigurable"`

	// The path name of the option file (VARIABLE_PATH), if the variable was set in an option file. If the variable was not set in an
	Path *string `mandatory:"true" json:"path"`

	// The description of the variable.
	Description *string `mandatory:"true" json:"description"`

	// The comma-separated list of possible values for the variable in value:valueDescription format.
	PossibleValues *string `mandatory:"true" json:"possibleValues"`

	// The comma-separated list of MySQL versions that support the variable.
	SupportedVersions *string `mandatory:"true" json:"supportedVersions"`
}

func (m MySqlConfigurationDataSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MySqlConfigurationDataSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMySqlConfigurationDataSummarySourceEnum(string(m.Source)); !ok && m.Source != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Source: %s. Supported values are: %s.", m.Source, strings.Join(GetMySqlConfigurationDataSummarySourceEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MySqlConfigurationDataSummarySourceEnum Enum with underlying type: string
type MySqlConfigurationDataSummarySourceEnum string

// Set of constants representing the allowable values for MySqlConfigurationDataSummarySourceEnum
const (
	MySqlConfigurationDataSummarySourceCompiled    MySqlConfigurationDataSummarySourceEnum = "COMPILED"
	MySqlConfigurationDataSummarySourceGlobal      MySqlConfigurationDataSummarySourceEnum = "GLOBAL"
	MySqlConfigurationDataSummarySourceServer      MySqlConfigurationDataSummarySourceEnum = "SERVER"
	MySqlConfigurationDataSummarySourceExplicit    MySqlConfigurationDataSummarySourceEnum = "EXPLICIT"
	MySqlConfigurationDataSummarySourceExtra       MySqlConfigurationDataSummarySourceEnum = "EXTRA"
	MySqlConfigurationDataSummarySourceUser        MySqlConfigurationDataSummarySourceEnum = "USER"
	MySqlConfigurationDataSummarySourceLogin       MySqlConfigurationDataSummarySourceEnum = "LOGIN"
	MySqlConfigurationDataSummarySourceCommandLine MySqlConfigurationDataSummarySourceEnum = "COMMAND_LINE"
	MySqlConfigurationDataSummarySourcePersisted   MySqlConfigurationDataSummarySourceEnum = "PERSISTED"
	MySqlConfigurationDataSummarySourceDynamic     MySqlConfigurationDataSummarySourceEnum = "DYNAMIC"
)

var mappingMySqlConfigurationDataSummarySourceEnum = map[string]MySqlConfigurationDataSummarySourceEnum{
	"COMPILED":     MySqlConfigurationDataSummarySourceCompiled,
	"GLOBAL":       MySqlConfigurationDataSummarySourceGlobal,
	"SERVER":       MySqlConfigurationDataSummarySourceServer,
	"EXPLICIT":     MySqlConfigurationDataSummarySourceExplicit,
	"EXTRA":        MySqlConfigurationDataSummarySourceExtra,
	"USER":         MySqlConfigurationDataSummarySourceUser,
	"LOGIN":        MySqlConfigurationDataSummarySourceLogin,
	"COMMAND_LINE": MySqlConfigurationDataSummarySourceCommandLine,
	"PERSISTED":    MySqlConfigurationDataSummarySourcePersisted,
	"DYNAMIC":      MySqlConfigurationDataSummarySourceDynamic,
}

var mappingMySqlConfigurationDataSummarySourceEnumLowerCase = map[string]MySqlConfigurationDataSummarySourceEnum{
	"compiled":     MySqlConfigurationDataSummarySourceCompiled,
	"global":       MySqlConfigurationDataSummarySourceGlobal,
	"server":       MySqlConfigurationDataSummarySourceServer,
	"explicit":     MySqlConfigurationDataSummarySourceExplicit,
	"extra":        MySqlConfigurationDataSummarySourceExtra,
	"user":         MySqlConfigurationDataSummarySourceUser,
	"login":        MySqlConfigurationDataSummarySourceLogin,
	"command_line": MySqlConfigurationDataSummarySourceCommandLine,
	"persisted":    MySqlConfigurationDataSummarySourcePersisted,
	"dynamic":      MySqlConfigurationDataSummarySourceDynamic,
}

// GetMySqlConfigurationDataSummarySourceEnumValues Enumerates the set of values for MySqlConfigurationDataSummarySourceEnum
func GetMySqlConfigurationDataSummarySourceEnumValues() []MySqlConfigurationDataSummarySourceEnum {
	values := make([]MySqlConfigurationDataSummarySourceEnum, 0)
	for _, v := range mappingMySqlConfigurationDataSummarySourceEnum {
		values = append(values, v)
	}
	return values
}

// GetMySqlConfigurationDataSummarySourceEnumStringValues Enumerates the set of values in String for MySqlConfigurationDataSummarySourceEnum
func GetMySqlConfigurationDataSummarySourceEnumStringValues() []string {
	return []string{
		"COMPILED",
		"GLOBAL",
		"SERVER",
		"EXPLICIT",
		"EXTRA",
		"USER",
		"LOGIN",
		"COMMAND_LINE",
		"PERSISTED",
		"DYNAMIC",
	}
}

// GetMappingMySqlConfigurationDataSummarySourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMySqlConfigurationDataSummarySourceEnum(val string) (MySqlConfigurationDataSummarySourceEnum, bool) {
	enum, ok := mappingMySqlConfigurationDataSummarySourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
