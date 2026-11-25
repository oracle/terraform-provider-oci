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

// ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetails Automatic resource management configuration details for the Globally distributed autonomous database.
type ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetails struct {

	// Flag indicating if autoResourceManagement is enabled or disabled.
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// Time period to wait for the database to get stable after autoResourceManagement event.
	CoolOffPeriodInMinutes *int `mandatory:"false" json:"coolOffPeriodInMinutes"`

	// Maximum number of move replication unit attempts allowed within the configured period per database.
	MaxMoveRuAttempts *int `mandatory:"false" json:"maxMoveRuAttempts"`

	// Time unit applicable to maxMoveRuAttempts.
	MaxMoveRuUnit ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum `mandatory:"false" json:"maxMoveRuUnit,omitempty"`

	// Mode of autoResourceManagement execution.
	Mode ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnum `mandatory:"false" json:"mode,omitempty"`

	// The action that will be taken when autoResourceManagement is triggered.
	ActionType ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnum `mandatory:"false" json:"actionType,omitempty"`

	// The list of notification topic ids OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) to receive autoResourceManagement events.
	NotificationTopicIds []string `mandatory:"false" json:"notificationTopicIds"`

	// The list of stream ids OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) that receive shard-related telemetry.
	StreamIds []string `mandatory:"false" json:"streamIds"`
}

func (m ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum(string(m.MaxMoveRuUnit)); !ok && m.MaxMoveRuUnit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MaxMoveRuUnit: %s. Supported values are: %s.", m.MaxMoveRuUnit, strings.Join(GetConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnumStringValues(), ",")))
	}
	if _, ok := GetMappingConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnum(string(m.Mode)); !ok && m.Mode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Mode: %s. Supported values are: %s.", m.Mode, strings.Join(GetConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnum(string(m.ActionType)); !ok && m.ActionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ActionType: %s. Supported values are: %s.", m.ActionType, strings.Join(GetConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum Enum with underlying type: string
type ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum string

// Set of constants representing the allowable values for ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum
const (
	ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitHour ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum = "HOUR"
	ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitDay  ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum = "DAY"
)

var mappingConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum = map[string]ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum{
	"HOUR": ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitHour,
	"DAY":  ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitDay,
}

var mappingConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnumLowerCase = map[string]ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum{
	"hour": ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitHour,
	"day":  ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitDay,
}

// GetConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnumValues Enumerates the set of values for ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum
func GetConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnumValues() []ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum {
	values := make([]ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum, 0)
	for _, v := range mappingConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnumStringValues Enumerates the set of values in String for ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum
func GetConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnumStringValues() []string {
	return []string{
		"HOUR",
		"DAY",
	}
}

// GetMappingConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum(val string) (ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum, bool) {
	enum, ok := mappingConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnum Enum with underlying type: string
type ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnum string

// Set of constants representing the allowable values for ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnum
const (
	ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsModeAutomatic      ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnum = "AUTOMATIC"
	ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsModeRecommendation ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnum = "RECOMMENDATION"
)

var mappingConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnum = map[string]ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnum{
	"AUTOMATIC":      ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsModeAutomatic,
	"RECOMMENDATION": ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsModeRecommendation,
}

var mappingConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnumLowerCase = map[string]ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnum{
	"automatic":      ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsModeAutomatic,
	"recommendation": ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsModeRecommendation,
}

// GetConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnumValues Enumerates the set of values for ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnum
func GetConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnumValues() []ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnum {
	values := make([]ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnum, 0)
	for _, v := range mappingConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnum {
		values = append(values, v)
	}
	return values
}

// GetConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnumStringValues Enumerates the set of values in String for ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnum
func GetConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnumStringValues() []string {
	return []string{
		"AUTOMATIC",
		"RECOMMENDATION",
	}
}

// GetMappingConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnum(val string) (ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnum, bool) {
	enum, ok := mappingConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnum Enum with underlying type: string
type ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnum string

// Set of constants representing the allowable values for ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnum
const (
	ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeMove ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnum = "MOVE"
)

var mappingConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnum = map[string]ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnum{
	"MOVE": ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeMove,
}

var mappingConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnumLowerCase = map[string]ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnum{
	"move": ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeMove,
}

// GetConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnumValues Enumerates the set of values for ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnum
func GetConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnumValues() []ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnum {
	values := make([]ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnum, 0)
	for _, v := range mappingConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnumStringValues Enumerates the set of values in String for ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnum
func GetConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnumStringValues() []string {
	return []string{
		"MOVE",
	}
}

// GetMappingConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnum(val string) (ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnum, bool) {
	enum, ok := mappingConfigureDistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
