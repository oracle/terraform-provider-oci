// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// MySqlConfigurationDataSummary Configuration Variables for MySQL.
type MySqlConfigurationDataSummary struct {

	// The name of variable
	Name *string `mandatory:"true" json:"name"`

	// The value of variable
	Value *string `mandatory:"true" json:"value"`

	// The source from which the variable was most recently set
	Source MySqlConfigurationDataSummarySourceEnum `mandatory:"true" json:"source"`

	// Minimum value of variable
	MinValue *float32 `mandatory:"true" json:"minValue"`

	// Maximum value of variable
	MaxValue *float32 `mandatory:"true" json:"maxValue"`

	// type of variable
	Type *string `mandatory:"true" json:"type"`

	// default value of variable
	DefaultValue *string `mandatory:"true" json:"defaultValue"`

	// Time when value was set
	TimeSet *common.SDKTime `mandatory:"true" json:"timeSet"`

	// Host from where this value was set. Empty for MySql Database System
	HostSet *string `mandatory:"true" json:"hostSet"`

	// User who set this value. Empty for MySql Database System
	UserSet *string `mandatory:"true" json:"userSet"`

	// Whether variable can be set dynamically or not
	IsDynamic *bool `mandatory:"true" json:"isDynamic"`

	// whether variable is set at server startup
	IsInit *bool `mandatory:"true" json:"isInit"`

	// Whether this variable is configurable
	IsConfigurable *bool `mandatory:"true" json:"isConfigurable"`

	// If the variable was set from an option file, VARIABLE_PATH is the path name of that file. Otherwise, the value is the empty string.
	Path *string `mandatory:"true" json:"path"`

	// Description of the variable
	Description *string `mandatory:"true" json:"description"`

	// Comma separated list of possible values for the variable in value:valueDescription format
	PossibleValues *string `mandatory:"true" json:"possibleValues"`

	// Comma separated list of MySql versions where this variable is supported
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
