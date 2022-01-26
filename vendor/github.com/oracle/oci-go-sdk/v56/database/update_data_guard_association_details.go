// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// UpdateDataGuardAssociationDetails The configuration details for updating a Data Guard association for a database.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type UpdateDataGuardAssociationDetails struct {

	// A strong password for the 'SYS', 'SYSTEM', and 'PDB Admin' users to apply during standby creation.
	// The password must contain no fewer than nine characters and include:
	// * At least two uppercase characters.
	// * At least two lowercase characters.
	// * At least two numeric characters.
	// * At least two special characters. Valid special characters include "_", "#", and "-" only.
	// **The password MUST be the same as the primary admin password.**
	DatabaseAdminPassword *string `mandatory:"false" json:"databaseAdminPassword"`

	// The protection mode for the Data Guard association's primary and standby database. For more information, see
	// Oracle Data Guard Protection Modes (http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000)
	// in the Oracle Data Guard documentation.
	ProtectionMode UpdateDataGuardAssociationDetailsProtectionModeEnum `mandatory:"false" json:"protectionMode,omitempty"`

	// The redo transport type to use for this Data Guard association.  Valid values depend on the specified 'protectionMode':
	// * MAXIMUM_AVAILABILITY - Use SYNC or FASTSYNC
	// * MAXIMUM_PERFORMANCE - Use ASYNC
	// * MAXIMUM_PROTECTION - Use SYNC
	// For more information, see
	// Redo Transport Services (http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-redo-transport-services.htm#SBYDB00400)
	// in the Oracle Data Guard documentation.
	TransportType UpdateDataGuardAssociationDetailsTransportTypeEnum `mandatory:"false" json:"transportType,omitempty"`

	// True if active Data Guard is enabled. Update this parameter to change the Data Guard setting.
	IsActiveDataGuardEnabled *bool `mandatory:"false" json:"isActiveDataGuardEnabled"`
}

func (m UpdateDataGuardAssociationDetails) String() string {
	return common.PointerString(m)
}

// UpdateDataGuardAssociationDetailsProtectionModeEnum Enum with underlying type: string
type UpdateDataGuardAssociationDetailsProtectionModeEnum string

// Set of constants representing the allowable values for UpdateDataGuardAssociationDetailsProtectionModeEnum
const (
	UpdateDataGuardAssociationDetailsProtectionModeAvailability UpdateDataGuardAssociationDetailsProtectionModeEnum = "MAXIMUM_AVAILABILITY"
	UpdateDataGuardAssociationDetailsProtectionModePerformance  UpdateDataGuardAssociationDetailsProtectionModeEnum = "MAXIMUM_PERFORMANCE"
	UpdateDataGuardAssociationDetailsProtectionModeProtection   UpdateDataGuardAssociationDetailsProtectionModeEnum = "MAXIMUM_PROTECTION"
)

var mappingUpdateDataGuardAssociationDetailsProtectionMode = map[string]UpdateDataGuardAssociationDetailsProtectionModeEnum{
	"MAXIMUM_AVAILABILITY": UpdateDataGuardAssociationDetailsProtectionModeAvailability,
	"MAXIMUM_PERFORMANCE":  UpdateDataGuardAssociationDetailsProtectionModePerformance,
	"MAXIMUM_PROTECTION":   UpdateDataGuardAssociationDetailsProtectionModeProtection,
}

// GetUpdateDataGuardAssociationDetailsProtectionModeEnumValues Enumerates the set of values for UpdateDataGuardAssociationDetailsProtectionModeEnum
func GetUpdateDataGuardAssociationDetailsProtectionModeEnumValues() []UpdateDataGuardAssociationDetailsProtectionModeEnum {
	values := make([]UpdateDataGuardAssociationDetailsProtectionModeEnum, 0)
	for _, v := range mappingUpdateDataGuardAssociationDetailsProtectionMode {
		values = append(values, v)
	}
	return values
}

// UpdateDataGuardAssociationDetailsTransportTypeEnum Enum with underlying type: string
type UpdateDataGuardAssociationDetailsTransportTypeEnum string

// Set of constants representing the allowable values for UpdateDataGuardAssociationDetailsTransportTypeEnum
const (
	UpdateDataGuardAssociationDetailsTransportTypeSync     UpdateDataGuardAssociationDetailsTransportTypeEnum = "SYNC"
	UpdateDataGuardAssociationDetailsTransportTypeAsync    UpdateDataGuardAssociationDetailsTransportTypeEnum = "ASYNC"
	UpdateDataGuardAssociationDetailsTransportTypeFastsync UpdateDataGuardAssociationDetailsTransportTypeEnum = "FASTSYNC"
)

var mappingUpdateDataGuardAssociationDetailsTransportType = map[string]UpdateDataGuardAssociationDetailsTransportTypeEnum{
	"SYNC":     UpdateDataGuardAssociationDetailsTransportTypeSync,
	"ASYNC":    UpdateDataGuardAssociationDetailsTransportTypeAsync,
	"FASTSYNC": UpdateDataGuardAssociationDetailsTransportTypeFastsync,
}

// GetUpdateDataGuardAssociationDetailsTransportTypeEnumValues Enumerates the set of values for UpdateDataGuardAssociationDetailsTransportTypeEnum
func GetUpdateDataGuardAssociationDetailsTransportTypeEnumValues() []UpdateDataGuardAssociationDetailsTransportTypeEnum {
	values := make([]UpdateDataGuardAssociationDetailsTransportTypeEnum, 0)
	for _, v := range mappingUpdateDataGuardAssociationDetailsTransportType {
		values = append(values, v)
	}
	return values
}
