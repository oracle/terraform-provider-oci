// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Integration API
//
// Oracle Integration API.
//

package integration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ConvertInstanceDetails Details for converting integration instance to Disaster Recovery Enabled instance type
type ConvertInstanceDetails struct {

	// Convert given instance to specified DR instance
	ConversionType ConvertInstanceDetailsConversionTypeEnum `mandatory:"true" json:"conversionType"`

	// Conversion phase for convert instance operation.
	ConversionPhase ConvertInstanceDetailsConversionPhaseEnum `mandatory:"false" json:"conversionPhase,omitempty"`
}

func (m ConvertInstanceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ConvertInstanceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingConvertInstanceDetailsConversionTypeEnum(string(m.ConversionType)); !ok && m.ConversionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConversionType: %s. Supported values are: %s.", m.ConversionType, strings.Join(GetConvertInstanceDetailsConversionTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingConvertInstanceDetailsConversionPhaseEnum(string(m.ConversionPhase)); !ok && m.ConversionPhase != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConversionPhase: %s. Supported values are: %s.", m.ConversionPhase, strings.Join(GetConvertInstanceDetailsConversionPhaseEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ConvertInstanceDetailsConversionTypeEnum Enum with underlying type: string
type ConvertInstanceDetailsConversionTypeEnum string

// Set of constants representing the allowable values for ConvertInstanceDetailsConversionTypeEnum
const (
	ConvertInstanceDetailsConversionTypeDisasterRecovery ConvertInstanceDetailsConversionTypeEnum = "DISASTER_RECOVERY"
	ConvertInstanceDetailsConversionTypeDevelopmentShape ConvertInstanceDetailsConversionTypeEnum = "DEVELOPMENT_SHAPE"
	ConvertInstanceDetailsConversionTypeProductionShape  ConvertInstanceDetailsConversionTypeEnum = "PRODUCTION_SHAPE"
)

var mappingConvertInstanceDetailsConversionTypeEnum = map[string]ConvertInstanceDetailsConversionTypeEnum{
	"DISASTER_RECOVERY": ConvertInstanceDetailsConversionTypeDisasterRecovery,
	"DEVELOPMENT_SHAPE": ConvertInstanceDetailsConversionTypeDevelopmentShape,
	"PRODUCTION_SHAPE":  ConvertInstanceDetailsConversionTypeProductionShape,
}

var mappingConvertInstanceDetailsConversionTypeEnumLowerCase = map[string]ConvertInstanceDetailsConversionTypeEnum{
	"disaster_recovery": ConvertInstanceDetailsConversionTypeDisasterRecovery,
	"development_shape": ConvertInstanceDetailsConversionTypeDevelopmentShape,
	"production_shape":  ConvertInstanceDetailsConversionTypeProductionShape,
}

// GetConvertInstanceDetailsConversionTypeEnumValues Enumerates the set of values for ConvertInstanceDetailsConversionTypeEnum
func GetConvertInstanceDetailsConversionTypeEnumValues() []ConvertInstanceDetailsConversionTypeEnum {
	values := make([]ConvertInstanceDetailsConversionTypeEnum, 0)
	for _, v := range mappingConvertInstanceDetailsConversionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetConvertInstanceDetailsConversionTypeEnumStringValues Enumerates the set of values in String for ConvertInstanceDetailsConversionTypeEnum
func GetConvertInstanceDetailsConversionTypeEnumStringValues() []string {
	return []string{
		"DISASTER_RECOVERY",
		"DEVELOPMENT_SHAPE",
		"PRODUCTION_SHAPE",
	}
}

// GetMappingConvertInstanceDetailsConversionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConvertInstanceDetailsConversionTypeEnum(val string) (ConvertInstanceDetailsConversionTypeEnum, bool) {
	enum, ok := mappingConvertInstanceDetailsConversionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ConvertInstanceDetailsConversionPhaseEnum Enum with underlying type: string
type ConvertInstanceDetailsConversionPhaseEnum string

// Set of constants representing the allowable values for ConvertInstanceDetailsConversionPhaseEnum
const (
	ConvertInstanceDetailsConversionPhaseBeginMigration    ConvertInstanceDetailsConversionPhaseEnum = "BEGIN_MIGRATION"
	ConvertInstanceDetailsConversionPhaseCompleteMigration ConvertInstanceDetailsConversionPhaseEnum = "COMPLETE_MIGRATION"
	ConvertInstanceDetailsConversionPhaseRollbackMigration ConvertInstanceDetailsConversionPhaseEnum = "ROLLBACK_MIGRATION"
)

var mappingConvertInstanceDetailsConversionPhaseEnum = map[string]ConvertInstanceDetailsConversionPhaseEnum{
	"BEGIN_MIGRATION":    ConvertInstanceDetailsConversionPhaseBeginMigration,
	"COMPLETE_MIGRATION": ConvertInstanceDetailsConversionPhaseCompleteMigration,
	"ROLLBACK_MIGRATION": ConvertInstanceDetailsConversionPhaseRollbackMigration,
}

var mappingConvertInstanceDetailsConversionPhaseEnumLowerCase = map[string]ConvertInstanceDetailsConversionPhaseEnum{
	"begin_migration":    ConvertInstanceDetailsConversionPhaseBeginMigration,
	"complete_migration": ConvertInstanceDetailsConversionPhaseCompleteMigration,
	"rollback_migration": ConvertInstanceDetailsConversionPhaseRollbackMigration,
}

// GetConvertInstanceDetailsConversionPhaseEnumValues Enumerates the set of values for ConvertInstanceDetailsConversionPhaseEnum
func GetConvertInstanceDetailsConversionPhaseEnumValues() []ConvertInstanceDetailsConversionPhaseEnum {
	values := make([]ConvertInstanceDetailsConversionPhaseEnum, 0)
	for _, v := range mappingConvertInstanceDetailsConversionPhaseEnum {
		values = append(values, v)
	}
	return values
}

// GetConvertInstanceDetailsConversionPhaseEnumStringValues Enumerates the set of values in String for ConvertInstanceDetailsConversionPhaseEnum
func GetConvertInstanceDetailsConversionPhaseEnumStringValues() []string {
	return []string{
		"BEGIN_MIGRATION",
		"COMPLETE_MIGRATION",
		"ROLLBACK_MIGRATION",
	}
}

// GetMappingConvertInstanceDetailsConversionPhaseEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConvertInstanceDetailsConversionPhaseEnum(val string) (ConvertInstanceDetailsConversionPhaseEnum, bool) {
	enum, ok := mappingConvertInstanceDetailsConversionPhaseEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
