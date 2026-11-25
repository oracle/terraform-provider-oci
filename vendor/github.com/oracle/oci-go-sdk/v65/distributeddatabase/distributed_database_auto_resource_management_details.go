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

// DistributedDatabaseAutoResourceManagementDetails AutoResourceManagement configuration details for the Globally distributed database.
type DistributedDatabaseAutoResourceManagementDetails struct {

	// Flag indicating if autoResourceManagement is enabled or disabled.
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// Time period to wait for the database to get stable after autoResourceManagement event.
	CoolOffPeriodInMinutes *int `mandatory:"false" json:"coolOffPeriodInMinutes"`

	// Maximum number of move replication unit attempts allowed within the configured period per database.
	MaxMoveRuAttempts *int `mandatory:"false" json:"maxMoveRuAttempts"`

	// Time unit applicable to maxMoveRuAttempts.
	MaxMoveRuUnit DistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum `mandatory:"false" json:"maxMoveRuUnit,omitempty"`

	// Mode of autoResourceManagement execution.
	Mode DistributedDatabaseAutoResourceManagementDetailsModeEnum `mandatory:"false" json:"mode,omitempty"`

	// The action that will be taken when autoResourceManagement is triggered.
	ActionType DistributedDatabaseAutoResourceManagementDetailsActionTypeEnum `mandatory:"false" json:"actionType,omitempty"`

	// The list of notification topic ids OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) to receive autoResourceManagement events.
	NotificationTopicIds []string `mandatory:"false" json:"notificationTopicIds"`

	// The list of stream ids OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) that receive shard-related telemetry.
	StreamIds []string `mandatory:"false" json:"streamIds"`
}

func (m DistributedDatabaseAutoResourceManagementDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DistributedDatabaseAutoResourceManagementDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum(string(m.MaxMoveRuUnit)); !ok && m.MaxMoveRuUnit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MaxMoveRuUnit: %s. Supported values are: %s.", m.MaxMoveRuUnit, strings.Join(GetDistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDistributedDatabaseAutoResourceManagementDetailsModeEnum(string(m.Mode)); !ok && m.Mode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Mode: %s. Supported values are: %s.", m.Mode, strings.Join(GetDistributedDatabaseAutoResourceManagementDetailsModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDistributedDatabaseAutoResourceManagementDetailsActionTypeEnum(string(m.ActionType)); !ok && m.ActionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ActionType: %s. Supported values are: %s.", m.ActionType, strings.Join(GetDistributedDatabaseAutoResourceManagementDetailsActionTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum Enum with underlying type: string
type DistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum string

// Set of constants representing the allowable values for DistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum
const (
	DistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitHour DistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum = "HOUR"
	DistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitDay  DistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum = "DAY"
)

var mappingDistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum = map[string]DistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum{
	"HOUR": DistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitHour,
	"DAY":  DistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitDay,
}

var mappingDistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnumLowerCase = map[string]DistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum{
	"hour": DistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitHour,
	"day":  DistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitDay,
}

// GetDistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnumValues Enumerates the set of values for DistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum
func GetDistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnumValues() []DistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum {
	values := make([]DistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum, 0)
	for _, v := range mappingDistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnumStringValues Enumerates the set of values in String for DistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum
func GetDistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnumStringValues() []string {
	return []string{
		"HOUR",
		"DAY",
	}
}

// GetMappingDistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum(val string) (DistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnum, bool) {
	enum, ok := mappingDistributedDatabaseAutoResourceManagementDetailsMaxMoveRuUnitEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DistributedDatabaseAutoResourceManagementDetailsModeEnum Enum with underlying type: string
type DistributedDatabaseAutoResourceManagementDetailsModeEnum string

// Set of constants representing the allowable values for DistributedDatabaseAutoResourceManagementDetailsModeEnum
const (
	DistributedDatabaseAutoResourceManagementDetailsModeAutomatic      DistributedDatabaseAutoResourceManagementDetailsModeEnum = "AUTOMATIC"
	DistributedDatabaseAutoResourceManagementDetailsModeRecommendation DistributedDatabaseAutoResourceManagementDetailsModeEnum = "RECOMMENDATION"
)

var mappingDistributedDatabaseAutoResourceManagementDetailsModeEnum = map[string]DistributedDatabaseAutoResourceManagementDetailsModeEnum{
	"AUTOMATIC":      DistributedDatabaseAutoResourceManagementDetailsModeAutomatic,
	"RECOMMENDATION": DistributedDatabaseAutoResourceManagementDetailsModeRecommendation,
}

var mappingDistributedDatabaseAutoResourceManagementDetailsModeEnumLowerCase = map[string]DistributedDatabaseAutoResourceManagementDetailsModeEnum{
	"automatic":      DistributedDatabaseAutoResourceManagementDetailsModeAutomatic,
	"recommendation": DistributedDatabaseAutoResourceManagementDetailsModeRecommendation,
}

// GetDistributedDatabaseAutoResourceManagementDetailsModeEnumValues Enumerates the set of values for DistributedDatabaseAutoResourceManagementDetailsModeEnum
func GetDistributedDatabaseAutoResourceManagementDetailsModeEnumValues() []DistributedDatabaseAutoResourceManagementDetailsModeEnum {
	values := make([]DistributedDatabaseAutoResourceManagementDetailsModeEnum, 0)
	for _, v := range mappingDistributedDatabaseAutoResourceManagementDetailsModeEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedDatabaseAutoResourceManagementDetailsModeEnumStringValues Enumerates the set of values in String for DistributedDatabaseAutoResourceManagementDetailsModeEnum
func GetDistributedDatabaseAutoResourceManagementDetailsModeEnumStringValues() []string {
	return []string{
		"AUTOMATIC",
		"RECOMMENDATION",
	}
}

// GetMappingDistributedDatabaseAutoResourceManagementDetailsModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedDatabaseAutoResourceManagementDetailsModeEnum(val string) (DistributedDatabaseAutoResourceManagementDetailsModeEnum, bool) {
	enum, ok := mappingDistributedDatabaseAutoResourceManagementDetailsModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DistributedDatabaseAutoResourceManagementDetailsActionTypeEnum Enum with underlying type: string
type DistributedDatabaseAutoResourceManagementDetailsActionTypeEnum string

// Set of constants representing the allowable values for DistributedDatabaseAutoResourceManagementDetailsActionTypeEnum
const (
	DistributedDatabaseAutoResourceManagementDetailsActionTypeMove DistributedDatabaseAutoResourceManagementDetailsActionTypeEnum = "MOVE"
)

var mappingDistributedDatabaseAutoResourceManagementDetailsActionTypeEnum = map[string]DistributedDatabaseAutoResourceManagementDetailsActionTypeEnum{
	"MOVE": DistributedDatabaseAutoResourceManagementDetailsActionTypeMove,
}

var mappingDistributedDatabaseAutoResourceManagementDetailsActionTypeEnumLowerCase = map[string]DistributedDatabaseAutoResourceManagementDetailsActionTypeEnum{
	"move": DistributedDatabaseAutoResourceManagementDetailsActionTypeMove,
}

// GetDistributedDatabaseAutoResourceManagementDetailsActionTypeEnumValues Enumerates the set of values for DistributedDatabaseAutoResourceManagementDetailsActionTypeEnum
func GetDistributedDatabaseAutoResourceManagementDetailsActionTypeEnumValues() []DistributedDatabaseAutoResourceManagementDetailsActionTypeEnum {
	values := make([]DistributedDatabaseAutoResourceManagementDetailsActionTypeEnum, 0)
	for _, v := range mappingDistributedDatabaseAutoResourceManagementDetailsActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedDatabaseAutoResourceManagementDetailsActionTypeEnumStringValues Enumerates the set of values in String for DistributedDatabaseAutoResourceManagementDetailsActionTypeEnum
func GetDistributedDatabaseAutoResourceManagementDetailsActionTypeEnumStringValues() []string {
	return []string{
		"MOVE",
	}
}

// GetMappingDistributedDatabaseAutoResourceManagementDetailsActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedDatabaseAutoResourceManagementDetailsActionTypeEnum(val string) (DistributedDatabaseAutoResourceManagementDetailsActionTypeEnum, bool) {
	enum, ok := mappingDistributedDatabaseAutoResourceManagementDetailsActionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
