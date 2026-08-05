// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DataGuardGroup Details of Data Guard setup that the given database is part of.
// Also includes information about databases part of this Data Guard group and properties for their Data Guard configuration.
type DataGuardGroup struct {

	// List of Data Guard members, representing each database that is part of Data Guard.
	Members []DataGuardGroupMember `mandatory:"false" json:"members"`

	// The protection mode of this Data Guard. For more information, see
	// Oracle Data Guard Protection Modes (http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000)
	// in the Oracle Data Guard documentation.
	ProtectionMode DataGuardGroupProtectionModeEnum `mandatory:"false" json:"protectionMode,omitempty"`

	// Specifies readiness of Managed Automatic failover.
	ManagedAutoFailOverReadiness DataGuardGroupManagedAutoFailOverReadinessEnum `mandatory:"false" json:"managedAutoFailOverReadiness,omitempty"`
}

func (m DataGuardGroup) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataGuardGroup) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDataGuardGroupProtectionModeEnum(string(m.ProtectionMode)); !ok && m.ProtectionMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ProtectionMode: %s. Supported values are: %s.", m.ProtectionMode, strings.Join(GetDataGuardGroupProtectionModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDataGuardGroupManagedAutoFailOverReadinessEnum(string(m.ManagedAutoFailOverReadiness)); !ok && m.ManagedAutoFailOverReadiness != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ManagedAutoFailOverReadiness: %s. Supported values are: %s.", m.ManagedAutoFailOverReadiness, strings.Join(GetDataGuardGroupManagedAutoFailOverReadinessEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DataGuardGroupProtectionModeEnum Enum with underlying type: string
type DataGuardGroupProtectionModeEnum string

// Set of constants representing the allowable values for DataGuardGroupProtectionModeEnum
const (
	DataGuardGroupProtectionModeAvailability DataGuardGroupProtectionModeEnum = "MAXIMUM_AVAILABILITY"
	DataGuardGroupProtectionModePerformance  DataGuardGroupProtectionModeEnum = "MAXIMUM_PERFORMANCE"
	DataGuardGroupProtectionModeProtection   DataGuardGroupProtectionModeEnum = "MAXIMUM_PROTECTION"
)

var mappingDataGuardGroupProtectionModeEnum = map[string]DataGuardGroupProtectionModeEnum{
	"MAXIMUM_AVAILABILITY": DataGuardGroupProtectionModeAvailability,
	"MAXIMUM_PERFORMANCE":  DataGuardGroupProtectionModePerformance,
	"MAXIMUM_PROTECTION":   DataGuardGroupProtectionModeProtection,
}

var mappingDataGuardGroupProtectionModeEnumLowerCase = map[string]DataGuardGroupProtectionModeEnum{
	"maximum_availability": DataGuardGroupProtectionModeAvailability,
	"maximum_performance":  DataGuardGroupProtectionModePerformance,
	"maximum_protection":   DataGuardGroupProtectionModeProtection,
}

// GetDataGuardGroupProtectionModeEnumValues Enumerates the set of values for DataGuardGroupProtectionModeEnum
func GetDataGuardGroupProtectionModeEnumValues() []DataGuardGroupProtectionModeEnum {
	values := make([]DataGuardGroupProtectionModeEnum, 0)
	for _, v := range mappingDataGuardGroupProtectionModeEnum {
		values = append(values, v)
	}
	return values
}

// GetDataGuardGroupProtectionModeEnumStringValues Enumerates the set of values in String for DataGuardGroupProtectionModeEnum
func GetDataGuardGroupProtectionModeEnumStringValues() []string {
	return []string{
		"MAXIMUM_AVAILABILITY",
		"MAXIMUM_PERFORMANCE",
		"MAXIMUM_PROTECTION",
	}
}

// GetMappingDataGuardGroupProtectionModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataGuardGroupProtectionModeEnum(val string) (DataGuardGroupProtectionModeEnum, bool) {
	enum, ok := mappingDataGuardGroupProtectionModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DataGuardGroupManagedAutoFailOverReadinessEnum Enum with underlying type: string
type DataGuardGroupManagedAutoFailOverReadinessEnum string

// Set of constants representing the allowable values for DataGuardGroupManagedAutoFailOverReadinessEnum
const (
	DataGuardGroupManagedAutoFailOverReadinessHealthy  DataGuardGroupManagedAutoFailOverReadinessEnum = "HEALTHY"
	DataGuardGroupManagedAutoFailOverReadinessCritical DataGuardGroupManagedAutoFailOverReadinessEnum = "CRITICAL"
)

var mappingDataGuardGroupManagedAutoFailOverReadinessEnum = map[string]DataGuardGroupManagedAutoFailOverReadinessEnum{
	"HEALTHY":  DataGuardGroupManagedAutoFailOverReadinessHealthy,
	"CRITICAL": DataGuardGroupManagedAutoFailOverReadinessCritical,
}

var mappingDataGuardGroupManagedAutoFailOverReadinessEnumLowerCase = map[string]DataGuardGroupManagedAutoFailOverReadinessEnum{
	"healthy":  DataGuardGroupManagedAutoFailOverReadinessHealthy,
	"critical": DataGuardGroupManagedAutoFailOverReadinessCritical,
}

// GetDataGuardGroupManagedAutoFailOverReadinessEnumValues Enumerates the set of values for DataGuardGroupManagedAutoFailOverReadinessEnum
func GetDataGuardGroupManagedAutoFailOverReadinessEnumValues() []DataGuardGroupManagedAutoFailOverReadinessEnum {
	values := make([]DataGuardGroupManagedAutoFailOverReadinessEnum, 0)
	for _, v := range mappingDataGuardGroupManagedAutoFailOverReadinessEnum {
		values = append(values, v)
	}
	return values
}

// GetDataGuardGroupManagedAutoFailOverReadinessEnumStringValues Enumerates the set of values in String for DataGuardGroupManagedAutoFailOverReadinessEnum
func GetDataGuardGroupManagedAutoFailOverReadinessEnumStringValues() []string {
	return []string{
		"HEALTHY",
		"CRITICAL",
	}
}

// GetMappingDataGuardGroupManagedAutoFailOverReadinessEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataGuardGroupManagedAutoFailOverReadinessEnum(val string) (DataGuardGroupManagedAutoFailOverReadinessEnum, bool) {
	enum, ok := mappingDataGuardGroupManagedAutoFailOverReadinessEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
