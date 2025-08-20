// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// UpdateDataGuardDetails The properties for updating a standby database.
type UpdateDataGuardDetails struct {

	// The administrator password of the primary database in this Data Guard association.
	// **The password MUST be the same as the primary admin password.**
	DatabaseAdminPassword *string `mandatory:"false" json:"databaseAdminPassword"`

	// The protection mode of this Data Guard. For more information, see
	// Oracle Data Guard Protection Modes (http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000)
	// in the Oracle Data Guard documentation.
	ProtectionMode UpdateDataGuardDetailsProtectionModeEnum `mandatory:"false" json:"protectionMode,omitempty"`

	// The redo transport type to use for this Data Guard association.  Valid values depend on the specified `protectionMode`:
	// * MAXIMUM_AVAILABILITY - SYNC or FASTSYNC
	// * MAXIMUM_PERFORMANCE - ASYNC
	// * MAXIMUM_PROTECTION - SYNC
	// For more information, see
	// Redo Transport Services (http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-redo-transport-services.htm#SBYDB00400)
	// in the Oracle Data Guard documentation.
	// **IMPORTANT** - The only transport type currently supported by the Database service is ASYNC.
	TransportType UpdateDataGuardDetailsTransportTypeEnum `mandatory:"false" json:"transportType,omitempty"`

	// True if active Data Guard is enabled.
	IsActiveDataGuardEnabled *bool `mandatory:"false" json:"isActiveDataGuardEnabled"`
}

func (m UpdateDataGuardDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateDataGuardDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateDataGuardDetailsProtectionModeEnum(string(m.ProtectionMode)); !ok && m.ProtectionMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ProtectionMode: %s. Supported values are: %s.", m.ProtectionMode, strings.Join(GetUpdateDataGuardDetailsProtectionModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUpdateDataGuardDetailsTransportTypeEnum(string(m.TransportType)); !ok && m.TransportType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TransportType: %s. Supported values are: %s.", m.TransportType, strings.Join(GetUpdateDataGuardDetailsTransportTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateDataGuardDetailsProtectionModeEnum Enum with underlying type: string
type UpdateDataGuardDetailsProtectionModeEnum string

// Set of constants representing the allowable values for UpdateDataGuardDetailsProtectionModeEnum
const (
	UpdateDataGuardDetailsProtectionModeAvailability UpdateDataGuardDetailsProtectionModeEnum = "MAXIMUM_AVAILABILITY"
	UpdateDataGuardDetailsProtectionModePerformance  UpdateDataGuardDetailsProtectionModeEnum = "MAXIMUM_PERFORMANCE"
	UpdateDataGuardDetailsProtectionModeProtection   UpdateDataGuardDetailsProtectionModeEnum = "MAXIMUM_PROTECTION"
)

var mappingUpdateDataGuardDetailsProtectionModeEnum = map[string]UpdateDataGuardDetailsProtectionModeEnum{
	"MAXIMUM_AVAILABILITY": UpdateDataGuardDetailsProtectionModeAvailability,
	"MAXIMUM_PERFORMANCE":  UpdateDataGuardDetailsProtectionModePerformance,
	"MAXIMUM_PROTECTION":   UpdateDataGuardDetailsProtectionModeProtection,
}

var mappingUpdateDataGuardDetailsProtectionModeEnumLowerCase = map[string]UpdateDataGuardDetailsProtectionModeEnum{
	"maximum_availability": UpdateDataGuardDetailsProtectionModeAvailability,
	"maximum_performance":  UpdateDataGuardDetailsProtectionModePerformance,
	"maximum_protection":   UpdateDataGuardDetailsProtectionModeProtection,
}

// GetUpdateDataGuardDetailsProtectionModeEnumValues Enumerates the set of values for UpdateDataGuardDetailsProtectionModeEnum
func GetUpdateDataGuardDetailsProtectionModeEnumValues() []UpdateDataGuardDetailsProtectionModeEnum {
	values := make([]UpdateDataGuardDetailsProtectionModeEnum, 0)
	for _, v := range mappingUpdateDataGuardDetailsProtectionModeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateDataGuardDetailsProtectionModeEnumStringValues Enumerates the set of values in String for UpdateDataGuardDetailsProtectionModeEnum
func GetUpdateDataGuardDetailsProtectionModeEnumStringValues() []string {
	return []string{
		"MAXIMUM_AVAILABILITY",
		"MAXIMUM_PERFORMANCE",
		"MAXIMUM_PROTECTION",
	}
}

// GetMappingUpdateDataGuardDetailsProtectionModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateDataGuardDetailsProtectionModeEnum(val string) (UpdateDataGuardDetailsProtectionModeEnum, bool) {
	enum, ok := mappingUpdateDataGuardDetailsProtectionModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateDataGuardDetailsTransportTypeEnum Enum with underlying type: string
type UpdateDataGuardDetailsTransportTypeEnum string

// Set of constants representing the allowable values for UpdateDataGuardDetailsTransportTypeEnum
const (
	UpdateDataGuardDetailsTransportTypeSync     UpdateDataGuardDetailsTransportTypeEnum = "SYNC"
	UpdateDataGuardDetailsTransportTypeAsync    UpdateDataGuardDetailsTransportTypeEnum = "ASYNC"
	UpdateDataGuardDetailsTransportTypeFastsync UpdateDataGuardDetailsTransportTypeEnum = "FASTSYNC"
)

var mappingUpdateDataGuardDetailsTransportTypeEnum = map[string]UpdateDataGuardDetailsTransportTypeEnum{
	"SYNC":     UpdateDataGuardDetailsTransportTypeSync,
	"ASYNC":    UpdateDataGuardDetailsTransportTypeAsync,
	"FASTSYNC": UpdateDataGuardDetailsTransportTypeFastsync,
}

var mappingUpdateDataGuardDetailsTransportTypeEnumLowerCase = map[string]UpdateDataGuardDetailsTransportTypeEnum{
	"sync":     UpdateDataGuardDetailsTransportTypeSync,
	"async":    UpdateDataGuardDetailsTransportTypeAsync,
	"fastsync": UpdateDataGuardDetailsTransportTypeFastsync,
}

// GetUpdateDataGuardDetailsTransportTypeEnumValues Enumerates the set of values for UpdateDataGuardDetailsTransportTypeEnum
func GetUpdateDataGuardDetailsTransportTypeEnumValues() []UpdateDataGuardDetailsTransportTypeEnum {
	values := make([]UpdateDataGuardDetailsTransportTypeEnum, 0)
	for _, v := range mappingUpdateDataGuardDetailsTransportTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateDataGuardDetailsTransportTypeEnumStringValues Enumerates the set of values in String for UpdateDataGuardDetailsTransportTypeEnum
func GetUpdateDataGuardDetailsTransportTypeEnumStringValues() []string {
	return []string{
		"SYNC",
		"ASYNC",
		"FASTSYNC",
	}
}

// GetMappingUpdateDataGuardDetailsTransportTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateDataGuardDetailsTransportTypeEnum(val string) (UpdateDataGuardDetailsTransportTypeEnum, bool) {
	enum, ok := mappingUpdateDataGuardDetailsTransportTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
