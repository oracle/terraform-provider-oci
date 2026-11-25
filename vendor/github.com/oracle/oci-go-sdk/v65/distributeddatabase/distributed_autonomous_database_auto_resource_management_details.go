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

// DistributedAutonomousDatabaseAutoResourceManagementDetails Automatic resource management configuration details for the Globally distributed autonomous database.
type DistributedAutonomousDatabaseAutoResourceManagementDetails struct {

	// Flag indicating if autoResourceManagement is enabled or disabled.
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// Time period to wait for the database to get stable after autoResourceManagement event.
	CoolOffPeriodInMinutes *int `mandatory:"false" json:"coolOffPeriodInMinutes"`

	// Maximum number of move replication unit attempts allowed within the configured period per database.
	MaxMoveRuAttempts *int `mandatory:"false" json:"maxMoveRuAttempts"`

	// Time unit applicable to maxMoveRuAttempts.
	MaxMoveRuUnit DistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum `mandatory:"false" json:"maxMoveRuUnit,omitempty"`

	// Mode of autoResourceManagement execution.
	Mode DistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnum `mandatory:"false" json:"mode,omitempty"`

	// The action that will be taken when autoResourceManagement is triggered.
	ActionType DistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnum `mandatory:"false" json:"actionType,omitempty"`

	// The list of notification topic ids OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) to receive autoResourceManagement events.
	NotificationTopicIds []string `mandatory:"false" json:"notificationTopicIds"`

	// The list of stream ids OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) that receive shard-related telemetry.
	StreamIds []string `mandatory:"false" json:"streamIds"`
}

func (m DistributedAutonomousDatabaseAutoResourceManagementDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DistributedAutonomousDatabaseAutoResourceManagementDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum(string(m.MaxMoveRuUnit)); !ok && m.MaxMoveRuUnit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MaxMoveRuUnit: %s. Supported values are: %s.", m.MaxMoveRuUnit, strings.Join(GetDistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnum(string(m.Mode)); !ok && m.Mode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Mode: %s. Supported values are: %s.", m.Mode, strings.Join(GetDistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnum(string(m.ActionType)); !ok && m.ActionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ActionType: %s. Supported values are: %s.", m.ActionType, strings.Join(GetDistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum Enum with underlying type: string
type DistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum string

// Set of constants representing the allowable values for DistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum
const (
	DistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitHour DistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum = "HOUR"
	DistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitDay  DistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum = "DAY"
)

var mappingDistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum = map[string]DistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum{
	"HOUR": DistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitHour,
	"DAY":  DistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitDay,
}

var mappingDistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnumLowerCase = map[string]DistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum{
	"hour": DistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitHour,
	"day":  DistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitDay,
}

// GetDistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnumValues Enumerates the set of values for DistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum
func GetDistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnumValues() []DistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum {
	values := make([]DistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum, 0)
	for _, v := range mappingDistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnumStringValues Enumerates the set of values in String for DistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum
func GetDistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnumStringValues() []string {
	return []string{
		"HOUR",
		"DAY",
	}
}

// GetMappingDistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum(val string) (DistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum, bool) {
	enum, ok := mappingDistributedAutonomousDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnum Enum with underlying type: string
type DistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnum string

// Set of constants representing the allowable values for DistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnum
const (
	DistributedAutonomousDatabaseAutoResourceManagementDetailsModeAutomatic      DistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnum = "AUTOMATIC"
	DistributedAutonomousDatabaseAutoResourceManagementDetailsModeRecommendation DistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnum = "RECOMMENDATION"
)

var mappingDistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnum = map[string]DistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnum{
	"AUTOMATIC":      DistributedAutonomousDatabaseAutoResourceManagementDetailsModeAutomatic,
	"RECOMMENDATION": DistributedAutonomousDatabaseAutoResourceManagementDetailsModeRecommendation,
}

var mappingDistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnumLowerCase = map[string]DistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnum{
	"automatic":      DistributedAutonomousDatabaseAutoResourceManagementDetailsModeAutomatic,
	"recommendation": DistributedAutonomousDatabaseAutoResourceManagementDetailsModeRecommendation,
}

// GetDistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnumValues Enumerates the set of values for DistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnum
func GetDistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnumValues() []DistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnum {
	values := make([]DistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnum, 0)
	for _, v := range mappingDistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnumStringValues Enumerates the set of values in String for DistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnum
func GetDistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnumStringValues() []string {
	return []string{
		"AUTOMATIC",
		"RECOMMENDATION",
	}
}

// GetMappingDistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnum(val string) (DistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnum, bool) {
	enum, ok := mappingDistributedAutonomousDatabaseAutoResourceManagementDetailsModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnum Enum with underlying type: string
type DistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnum string

// Set of constants representing the allowable values for DistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnum
const (
	DistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeMove DistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnum = "MOVE"
)

var mappingDistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnum = map[string]DistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnum{
	"MOVE": DistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeMove,
}

var mappingDistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnumLowerCase = map[string]DistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnum{
	"move": DistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeMove,
}

// GetDistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnumValues Enumerates the set of values for DistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnum
func GetDistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnumValues() []DistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnum {
	values := make([]DistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnum, 0)
	for _, v := range mappingDistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnumStringValues Enumerates the set of values in String for DistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnum
func GetDistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnumStringValues() []string {
	return []string{
		"MOVE",
	}
}

// GetMappingDistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnum(val string) (DistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnum, bool) {
	enum, ok := mappingDistributedAutonomousDatabaseAutoResourceManagementDetailsActionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
