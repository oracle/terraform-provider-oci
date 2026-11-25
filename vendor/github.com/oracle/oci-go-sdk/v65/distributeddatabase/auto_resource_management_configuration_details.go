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

// AutoResourceManagementConfigurationDetails Automatic resource management configuration details for the Globally distributed database.
type AutoResourceManagementConfigurationDetails struct {

	// Flag indicating if autoResourceManagement is enabled or disabled.
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// Time period to wait for the database to get stable after autoResourceManagement event.
	CoolOffPeriodInMinutes *int `mandatory:"false" json:"coolOffPeriodInMinutes"`

	// Maximum number of move replication unit attempts allowed within the configured period per database.
	MaxMoveRuAttempts *int `mandatory:"false" json:"maxMoveRuAttempts"`

	// Time unit applicable to maxMoveRuAttempts.
	MaxMoveRuUnit AutoResourceManagementConfigurationDetailsMaxMoveRuUnitEnum `mandatory:"false" json:"maxMoveRuUnit,omitempty"`

	// Mode of autoResourceManagement execution.
	Mode AutoResourceManagementConfigurationDetailsModeEnum `mandatory:"false" json:"mode,omitempty"`

	// The action that will be taken when autoResourceManagement is triggered.
	ActionType AutoResourceManagementConfigurationDetailsActionTypeEnum `mandatory:"false" json:"actionType,omitempty"`

	// The list of notification topic ids OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) to receive autoResourceManagement events.
	NotificationTopicIds []string `mandatory:"false" json:"notificationTopicIds"`

	// The list of stream ids OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) that receive shard-related telemetry.
	StreamIds []string `mandatory:"false" json:"streamIds"`
}

func (m AutoResourceManagementConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutoResourceManagementConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAutoResourceManagementConfigurationDetailsMaxMoveRuUnitEnum(string(m.MaxMoveRuUnit)); !ok && m.MaxMoveRuUnit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MaxMoveRuUnit: %s. Supported values are: %s.", m.MaxMoveRuUnit, strings.Join(GetAutoResourceManagementConfigurationDetailsMaxMoveRuUnitEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutoResourceManagementConfigurationDetailsModeEnum(string(m.Mode)); !ok && m.Mode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Mode: %s. Supported values are: %s.", m.Mode, strings.Join(GetAutoResourceManagementConfigurationDetailsModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutoResourceManagementConfigurationDetailsActionTypeEnum(string(m.ActionType)); !ok && m.ActionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ActionType: %s. Supported values are: %s.", m.ActionType, strings.Join(GetAutoResourceManagementConfigurationDetailsActionTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AutoResourceManagementConfigurationDetailsMaxMoveRuUnitEnum Enum with underlying type: string
type AutoResourceManagementConfigurationDetailsMaxMoveRuUnitEnum string

// Set of constants representing the allowable values for AutoResourceManagementConfigurationDetailsMaxMoveRuUnitEnum
const (
	AutoResourceManagementConfigurationDetailsMaxMoveRuUnitHour AutoResourceManagementConfigurationDetailsMaxMoveRuUnitEnum = "HOUR"
	AutoResourceManagementConfigurationDetailsMaxMoveRuUnitDay  AutoResourceManagementConfigurationDetailsMaxMoveRuUnitEnum = "DAY"
)

var mappingAutoResourceManagementConfigurationDetailsMaxMoveRuUnitEnum = map[string]AutoResourceManagementConfigurationDetailsMaxMoveRuUnitEnum{
	"HOUR": AutoResourceManagementConfigurationDetailsMaxMoveRuUnitHour,
	"DAY":  AutoResourceManagementConfigurationDetailsMaxMoveRuUnitDay,
}

var mappingAutoResourceManagementConfigurationDetailsMaxMoveRuUnitEnumLowerCase = map[string]AutoResourceManagementConfigurationDetailsMaxMoveRuUnitEnum{
	"hour": AutoResourceManagementConfigurationDetailsMaxMoveRuUnitHour,
	"day":  AutoResourceManagementConfigurationDetailsMaxMoveRuUnitDay,
}

// GetAutoResourceManagementConfigurationDetailsMaxMoveRuUnitEnumValues Enumerates the set of values for AutoResourceManagementConfigurationDetailsMaxMoveRuUnitEnum
func GetAutoResourceManagementConfigurationDetailsMaxMoveRuUnitEnumValues() []AutoResourceManagementConfigurationDetailsMaxMoveRuUnitEnum {
	values := make([]AutoResourceManagementConfigurationDetailsMaxMoveRuUnitEnum, 0)
	for _, v := range mappingAutoResourceManagementConfigurationDetailsMaxMoveRuUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetAutoResourceManagementConfigurationDetailsMaxMoveRuUnitEnumStringValues Enumerates the set of values in String for AutoResourceManagementConfigurationDetailsMaxMoveRuUnitEnum
func GetAutoResourceManagementConfigurationDetailsMaxMoveRuUnitEnumStringValues() []string {
	return []string{
		"HOUR",
		"DAY",
	}
}

// GetMappingAutoResourceManagementConfigurationDetailsMaxMoveRuUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutoResourceManagementConfigurationDetailsMaxMoveRuUnitEnum(val string) (AutoResourceManagementConfigurationDetailsMaxMoveRuUnitEnum, bool) {
	enum, ok := mappingAutoResourceManagementConfigurationDetailsMaxMoveRuUnitEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutoResourceManagementConfigurationDetailsModeEnum Enum with underlying type: string
type AutoResourceManagementConfigurationDetailsModeEnum string

// Set of constants representing the allowable values for AutoResourceManagementConfigurationDetailsModeEnum
const (
	AutoResourceManagementConfigurationDetailsModeAutomatic      AutoResourceManagementConfigurationDetailsModeEnum = "AUTOMATIC"
	AutoResourceManagementConfigurationDetailsModeRecommendation AutoResourceManagementConfigurationDetailsModeEnum = "RECOMMENDATION"
)

var mappingAutoResourceManagementConfigurationDetailsModeEnum = map[string]AutoResourceManagementConfigurationDetailsModeEnum{
	"AUTOMATIC":      AutoResourceManagementConfigurationDetailsModeAutomatic,
	"RECOMMENDATION": AutoResourceManagementConfigurationDetailsModeRecommendation,
}

var mappingAutoResourceManagementConfigurationDetailsModeEnumLowerCase = map[string]AutoResourceManagementConfigurationDetailsModeEnum{
	"automatic":      AutoResourceManagementConfigurationDetailsModeAutomatic,
	"recommendation": AutoResourceManagementConfigurationDetailsModeRecommendation,
}

// GetAutoResourceManagementConfigurationDetailsModeEnumValues Enumerates the set of values for AutoResourceManagementConfigurationDetailsModeEnum
func GetAutoResourceManagementConfigurationDetailsModeEnumValues() []AutoResourceManagementConfigurationDetailsModeEnum {
	values := make([]AutoResourceManagementConfigurationDetailsModeEnum, 0)
	for _, v := range mappingAutoResourceManagementConfigurationDetailsModeEnum {
		values = append(values, v)
	}
	return values
}

// GetAutoResourceManagementConfigurationDetailsModeEnumStringValues Enumerates the set of values in String for AutoResourceManagementConfigurationDetailsModeEnum
func GetAutoResourceManagementConfigurationDetailsModeEnumStringValues() []string {
	return []string{
		"AUTOMATIC",
		"RECOMMENDATION",
	}
}

// GetMappingAutoResourceManagementConfigurationDetailsModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutoResourceManagementConfigurationDetailsModeEnum(val string) (AutoResourceManagementConfigurationDetailsModeEnum, bool) {
	enum, ok := mappingAutoResourceManagementConfigurationDetailsModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutoResourceManagementConfigurationDetailsActionTypeEnum Enum with underlying type: string
type AutoResourceManagementConfigurationDetailsActionTypeEnum string

// Set of constants representing the allowable values for AutoResourceManagementConfigurationDetailsActionTypeEnum
const (
	AutoResourceManagementConfigurationDetailsActionTypeMove AutoResourceManagementConfigurationDetailsActionTypeEnum = "MOVE"
)

var mappingAutoResourceManagementConfigurationDetailsActionTypeEnum = map[string]AutoResourceManagementConfigurationDetailsActionTypeEnum{
	"MOVE": AutoResourceManagementConfigurationDetailsActionTypeMove,
}

var mappingAutoResourceManagementConfigurationDetailsActionTypeEnumLowerCase = map[string]AutoResourceManagementConfigurationDetailsActionTypeEnum{
	"move": AutoResourceManagementConfigurationDetailsActionTypeMove,
}

// GetAutoResourceManagementConfigurationDetailsActionTypeEnumValues Enumerates the set of values for AutoResourceManagementConfigurationDetailsActionTypeEnum
func GetAutoResourceManagementConfigurationDetailsActionTypeEnumValues() []AutoResourceManagementConfigurationDetailsActionTypeEnum {
	values := make([]AutoResourceManagementConfigurationDetailsActionTypeEnum, 0)
	for _, v := range mappingAutoResourceManagementConfigurationDetailsActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAutoResourceManagementConfigurationDetailsActionTypeEnumStringValues Enumerates the set of values in String for AutoResourceManagementConfigurationDetailsActionTypeEnum
func GetAutoResourceManagementConfigurationDetailsActionTypeEnumStringValues() []string {
	return []string{
		"MOVE",
	}
}

// GetMappingAutoResourceManagementConfigurationDetailsActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutoResourceManagementConfigurationDetailsActionTypeEnum(val string) (AutoResourceManagementConfigurationDetailsActionTypeEnum, bool) {
	enum, ok := mappingAutoResourceManagementConfigurationDetailsActionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
