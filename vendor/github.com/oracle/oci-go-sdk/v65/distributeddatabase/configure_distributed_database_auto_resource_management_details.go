// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage the Globally distributed databases.
//

package distributeddatabase

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ConfigureDistributedDatabaseAutoResourceManagementDetails AutoResourceManagement configuration details for the Globally distributed database.
type ConfigureDistributedDatabaseAutoResourceManagementDetails struct {

	// Flag indicating if autoResourceManagement is enabled or disabled.
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// Time period to wait for the database to get stable after autoResourceManagement event.
	CoolOffPeriodInMinutes *int `mandatory:"false" json:"coolOffPeriodInMinutes"`

	// Maximum number of move replication unit attempts allowed within the configured period per database.
	MaxMoveRuAttempts *int `mandatory:"false" json:"maxMoveRuAttempts"`

	// Time unit applicable to maxMoveRuAttempts.
	MaxMoveRuUnit ConfigureDistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum `mandatory:"false" json:"maxMoveRuUnit,omitempty"`

	// Mode of autoResourceManagement execution.
	Mode ConfigureDistributedDatabaseAutoResourceManagementDetailsModeEnum `mandatory:"false" json:"mode,omitempty"`

	// The action that will be taken when autoResourceManagement is triggered.
	ActionType ConfigureDistributedDatabaseAutoResourceManagementDetailsActionTypeEnum `mandatory:"false" json:"actionType,omitempty"`

	// The list of notification topic ids OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) to receive autoResourceManagement events.
	NotificationTopicIds []string `mandatory:"false" json:"notificationTopicIds"`

	// The list of stream ids OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) that receive shard-related telemetry.
	StreamIds []string `mandatory:"false" json:"streamIds"`
}

func (m ConfigureDistributedDatabaseAutoResourceManagementDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ConfigureDistributedDatabaseAutoResourceManagementDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingConfigureDistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum(string(m.MaxMoveRuUnit)); !ok && m.MaxMoveRuUnit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MaxMoveRuUnit: %s. Supported values are: %s.", m.MaxMoveRuUnit, strings.Join(GetConfigureDistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnumStringValues(), ",")))
	}
	if _, ok := GetMappingConfigureDistributedDatabaseAutoResourceManagementDetailsModeEnum(string(m.Mode)); !ok && m.Mode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Mode: %s. Supported values are: %s.", m.Mode, strings.Join(GetConfigureDistributedDatabaseAutoResourceManagementDetailsModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingConfigureDistributedDatabaseAutoResourceManagementDetailsActionTypeEnum(string(m.ActionType)); !ok && m.ActionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ActionType: %s. Supported values are: %s.", m.ActionType, strings.Join(GetConfigureDistributedDatabaseAutoResourceManagementDetailsActionTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ConfigureDistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum Enum with underlying type: string
type ConfigureDistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum string

// Set of constants representing the allowable values for ConfigureDistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum
const (
	ConfigureDistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitHour ConfigureDistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum = "HOUR"
	ConfigureDistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitDay  ConfigureDistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum = "DAY"
)

var mappingConfigureDistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum = map[string]ConfigureDistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum{
	"HOUR": ConfigureDistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitHour,
	"DAY":  ConfigureDistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitDay,
}

var mappingConfigureDistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnumLowerCase = map[string]ConfigureDistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum{
	"hour": ConfigureDistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitHour,
	"day":  ConfigureDistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitDay,
}

// GetConfigureDistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnumValues Enumerates the set of values for ConfigureDistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum
func GetConfigureDistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnumValues() []ConfigureDistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum {
	values := make([]ConfigureDistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum, 0)
	for _, v := range mappingConfigureDistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetConfigureDistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnumStringValues Enumerates the set of values in String for ConfigureDistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum
func GetConfigureDistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnumStringValues() []string {
	return []string{
		"HOUR",
		"DAY",
	}
}

// GetMappingConfigureDistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConfigureDistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum(val string) (ConfigureDistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum, bool) {
	enum, ok := mappingConfigureDistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ConfigureDistributedDatabaseAutoResourceManagementDetailsModeEnum Enum with underlying type: string
type ConfigureDistributedDatabaseAutoResourceManagementDetailsModeEnum string

// Set of constants representing the allowable values for ConfigureDistributedDatabaseAutoResourceManagementDetailsModeEnum
const (
	ConfigureDistributedDatabaseAutoResourceManagementDetailsModeAutomatic      ConfigureDistributedDatabaseAutoResourceManagementDetailsModeEnum = "AUTOMATIC"
	ConfigureDistributedDatabaseAutoResourceManagementDetailsModeRecommendation ConfigureDistributedDatabaseAutoResourceManagementDetailsModeEnum = "RECOMMENDATION"
)

var mappingConfigureDistributedDatabaseAutoResourceManagementDetailsModeEnum = map[string]ConfigureDistributedDatabaseAutoResourceManagementDetailsModeEnum{
	"AUTOMATIC":      ConfigureDistributedDatabaseAutoResourceManagementDetailsModeAutomatic,
	"RECOMMENDATION": ConfigureDistributedDatabaseAutoResourceManagementDetailsModeRecommendation,
}

var mappingConfigureDistributedDatabaseAutoResourceManagementDetailsModeEnumLowerCase = map[string]ConfigureDistributedDatabaseAutoResourceManagementDetailsModeEnum{
	"automatic":      ConfigureDistributedDatabaseAutoResourceManagementDetailsModeAutomatic,
	"recommendation": ConfigureDistributedDatabaseAutoResourceManagementDetailsModeRecommendation,
}

// GetConfigureDistributedDatabaseAutoResourceManagementDetailsModeEnumValues Enumerates the set of values for ConfigureDistributedDatabaseAutoResourceManagementDetailsModeEnum
func GetConfigureDistributedDatabaseAutoResourceManagementDetailsModeEnumValues() []ConfigureDistributedDatabaseAutoResourceManagementDetailsModeEnum {
	values := make([]ConfigureDistributedDatabaseAutoResourceManagementDetailsModeEnum, 0)
	for _, v := range mappingConfigureDistributedDatabaseAutoResourceManagementDetailsModeEnum {
		values = append(values, v)
	}
	return values
}

// GetConfigureDistributedDatabaseAutoResourceManagementDetailsModeEnumStringValues Enumerates the set of values in String for ConfigureDistributedDatabaseAutoResourceManagementDetailsModeEnum
func GetConfigureDistributedDatabaseAutoResourceManagementDetailsModeEnumStringValues() []string {
	return []string{
		"AUTOMATIC",
		"RECOMMENDATION",
	}
}

// GetMappingConfigureDistributedDatabaseAutoResourceManagementDetailsModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConfigureDistributedDatabaseAutoResourceManagementDetailsModeEnum(val string) (ConfigureDistributedDatabaseAutoResourceManagementDetailsModeEnum, bool) {
	enum, ok := mappingConfigureDistributedDatabaseAutoResourceManagementDetailsModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ConfigureDistributedDatabaseAutoResourceManagementDetailsActionTypeEnum Enum with underlying type: string
type ConfigureDistributedDatabaseAutoResourceManagementDetailsActionTypeEnum string

// Set of constants representing the allowable values for ConfigureDistributedDatabaseAutoResourceManagementDetailsActionTypeEnum
const (
	ConfigureDistributedDatabaseAutoResourceManagementDetailsActionTypeMove ConfigureDistributedDatabaseAutoResourceManagementDetailsActionTypeEnum = "MOVE"
)

var mappingConfigureDistributedDatabaseAutoResourceManagementDetailsActionTypeEnum = map[string]ConfigureDistributedDatabaseAutoResourceManagementDetailsActionTypeEnum{
	"MOVE": ConfigureDistributedDatabaseAutoResourceManagementDetailsActionTypeMove,
}

var mappingConfigureDistributedDatabaseAutoResourceManagementDetailsActionTypeEnumLowerCase = map[string]ConfigureDistributedDatabaseAutoResourceManagementDetailsActionTypeEnum{
	"move": ConfigureDistributedDatabaseAutoResourceManagementDetailsActionTypeMove,
}

// GetConfigureDistributedDatabaseAutoResourceManagementDetailsActionTypeEnumValues Enumerates the set of values for ConfigureDistributedDatabaseAutoResourceManagementDetailsActionTypeEnum
func GetConfigureDistributedDatabaseAutoResourceManagementDetailsActionTypeEnumValues() []ConfigureDistributedDatabaseAutoResourceManagementDetailsActionTypeEnum {
	values := make([]ConfigureDistributedDatabaseAutoResourceManagementDetailsActionTypeEnum, 0)
	for _, v := range mappingConfigureDistributedDatabaseAutoResourceManagementDetailsActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetConfigureDistributedDatabaseAutoResourceManagementDetailsActionTypeEnumStringValues Enumerates the set of values in String for ConfigureDistributedDatabaseAutoResourceManagementDetailsActionTypeEnum
func GetConfigureDistributedDatabaseAutoResourceManagementDetailsActionTypeEnumStringValues() []string {
	return []string{
		"MOVE",
	}
}

// GetMappingConfigureDistributedDatabaseAutoResourceManagementDetailsActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConfigureDistributedDatabaseAutoResourceManagementDetailsActionTypeEnum(val string) (ConfigureDistributedDatabaseAutoResourceManagementDetailsActionTypeEnum, bool) {
	enum, ok := mappingConfigureDistributedDatabaseAutoResourceManagementDetailsActionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
