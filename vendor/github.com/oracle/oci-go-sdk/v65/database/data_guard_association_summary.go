// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DataGuardAssociationSummary The properties that define a Data Guard association.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an
// administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
// For information about endpoints and signing API requests, see
// About the API (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm). For information about available SDKs and tools, see
// SDKS and Other Tools (https://docs.cloud.oracle.com/Content/API/Concepts/sdks.htm).
type DataGuardAssociationSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Data Guard association.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the reporting database.
	DatabaseId *string `mandatory:"true" json:"databaseId"`

	// The role of the reporting database in this Data Guard association.
	Role DataGuardAssociationSummaryRoleEnum `mandatory:"true" json:"role"`

	// The current state of the Data Guard association.
	LifecycleState DataGuardAssociationSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the DB system containing the associated
	// peer database.
	PeerDbSystemId *string `mandatory:"true" json:"peerDbSystemId"`

	// The role of the peer database in this Data Guard association.
	PeerRole DataGuardAssociationSummaryPeerRoleEnum `mandatory:"true" json:"peerRole"`

	// The protection mode of this Data Guard association. For more information, see
	// Oracle Data Guard Protection Modes (http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000)
	// in the Oracle Data Guard documentation.
	ProtectionMode DataGuardAssociationSummaryProtectionModeEnum `mandatory:"true" json:"protectionMode"`

	// Additional information about the current lifecycleState, if available.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Database Home containing the associated peer database.
	PeerDbHomeId *string `mandatory:"false" json:"peerDbHomeId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the associated peer database.
	PeerDatabaseId *string `mandatory:"false" json:"peerDatabaseId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the peer database's Data Guard association.
	PeerDataGuardAssociationId *string `mandatory:"false" json:"peerDataGuardAssociationId"`

	// The lag time between updates to the primary database and application of the redo data on the standby database,
	// as computed by the reporting database.
	// Example: `9 seconds`
	ApplyLag *string `mandatory:"false" json:"applyLag"`

	// The rate at which redo logs are synced between the associated databases.
	// Example: `180 Mb per second`
	ApplyRate *string `mandatory:"false" json:"applyRate"`

	// The redo transport type used by this Data Guard association.  For more information, see
	// Redo Transport Services (http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-redo-transport-services.htm#SBYDB00400)
	// in the Oracle Data Guard documentation.
	TransportType DataGuardAssociationSummaryTransportTypeEnum `mandatory:"false" json:"transportType,omitempty"`

	// The date and time the Data Guard association was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// True if active Data Guard is enabled.
	IsActiveDataGuardEnabled *bool `mandatory:"false" json:"isActiveDataGuardEnabled"`
}

func (m DataGuardAssociationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataGuardAssociationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDataGuardAssociationSummaryRoleEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetDataGuardAssociationSummaryRoleEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDataGuardAssociationSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDataGuardAssociationSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDataGuardAssociationSummaryPeerRoleEnum(string(m.PeerRole)); !ok && m.PeerRole != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PeerRole: %s. Supported values are: %s.", m.PeerRole, strings.Join(GetDataGuardAssociationSummaryPeerRoleEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDataGuardAssociationSummaryProtectionModeEnum(string(m.ProtectionMode)); !ok && m.ProtectionMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ProtectionMode: %s. Supported values are: %s.", m.ProtectionMode, strings.Join(GetDataGuardAssociationSummaryProtectionModeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDataGuardAssociationSummaryTransportTypeEnum(string(m.TransportType)); !ok && m.TransportType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TransportType: %s. Supported values are: %s.", m.TransportType, strings.Join(GetDataGuardAssociationSummaryTransportTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DataGuardAssociationSummaryRoleEnum Enum with underlying type: string
type DataGuardAssociationSummaryRoleEnum string

// Set of constants representing the allowable values for DataGuardAssociationSummaryRoleEnum
const (
	DataGuardAssociationSummaryRolePrimary         DataGuardAssociationSummaryRoleEnum = "PRIMARY"
	DataGuardAssociationSummaryRoleStandby         DataGuardAssociationSummaryRoleEnum = "STANDBY"
	DataGuardAssociationSummaryRoleDisabledStandby DataGuardAssociationSummaryRoleEnum = "DISABLED_STANDBY"
)

var mappingDataGuardAssociationSummaryRoleEnum = map[string]DataGuardAssociationSummaryRoleEnum{
	"PRIMARY":          DataGuardAssociationSummaryRolePrimary,
	"STANDBY":          DataGuardAssociationSummaryRoleStandby,
	"DISABLED_STANDBY": DataGuardAssociationSummaryRoleDisabledStandby,
}

var mappingDataGuardAssociationSummaryRoleEnumLowerCase = map[string]DataGuardAssociationSummaryRoleEnum{
	"primary":          DataGuardAssociationSummaryRolePrimary,
	"standby":          DataGuardAssociationSummaryRoleStandby,
	"disabled_standby": DataGuardAssociationSummaryRoleDisabledStandby,
}

// GetDataGuardAssociationSummaryRoleEnumValues Enumerates the set of values for DataGuardAssociationSummaryRoleEnum
func GetDataGuardAssociationSummaryRoleEnumValues() []DataGuardAssociationSummaryRoleEnum {
	values := make([]DataGuardAssociationSummaryRoleEnum, 0)
	for _, v := range mappingDataGuardAssociationSummaryRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetDataGuardAssociationSummaryRoleEnumStringValues Enumerates the set of values in String for DataGuardAssociationSummaryRoleEnum
func GetDataGuardAssociationSummaryRoleEnumStringValues() []string {
	return []string{
		"PRIMARY",
		"STANDBY",
		"DISABLED_STANDBY",
	}
}

// GetMappingDataGuardAssociationSummaryRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataGuardAssociationSummaryRoleEnum(val string) (DataGuardAssociationSummaryRoleEnum, bool) {
	enum, ok := mappingDataGuardAssociationSummaryRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DataGuardAssociationSummaryLifecycleStateEnum Enum with underlying type: string
type DataGuardAssociationSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for DataGuardAssociationSummaryLifecycleStateEnum
const (
	DataGuardAssociationSummaryLifecycleStateProvisioning DataGuardAssociationSummaryLifecycleStateEnum = "PROVISIONING"
	DataGuardAssociationSummaryLifecycleStateAvailable    DataGuardAssociationSummaryLifecycleStateEnum = "AVAILABLE"
	DataGuardAssociationSummaryLifecycleStateUpdating     DataGuardAssociationSummaryLifecycleStateEnum = "UPDATING"
	DataGuardAssociationSummaryLifecycleStateTerminating  DataGuardAssociationSummaryLifecycleStateEnum = "TERMINATING"
	DataGuardAssociationSummaryLifecycleStateTerminated   DataGuardAssociationSummaryLifecycleStateEnum = "TERMINATED"
	DataGuardAssociationSummaryLifecycleStateFailed       DataGuardAssociationSummaryLifecycleStateEnum = "FAILED"
	DataGuardAssociationSummaryLifecycleStateUpgrading    DataGuardAssociationSummaryLifecycleStateEnum = "UPGRADING"
)

var mappingDataGuardAssociationSummaryLifecycleStateEnum = map[string]DataGuardAssociationSummaryLifecycleStateEnum{
	"PROVISIONING": DataGuardAssociationSummaryLifecycleStateProvisioning,
	"AVAILABLE":    DataGuardAssociationSummaryLifecycleStateAvailable,
	"UPDATING":     DataGuardAssociationSummaryLifecycleStateUpdating,
	"TERMINATING":  DataGuardAssociationSummaryLifecycleStateTerminating,
	"TERMINATED":   DataGuardAssociationSummaryLifecycleStateTerminated,
	"FAILED":       DataGuardAssociationSummaryLifecycleStateFailed,
	"UPGRADING":    DataGuardAssociationSummaryLifecycleStateUpgrading,
}

var mappingDataGuardAssociationSummaryLifecycleStateEnumLowerCase = map[string]DataGuardAssociationSummaryLifecycleStateEnum{
	"provisioning": DataGuardAssociationSummaryLifecycleStateProvisioning,
	"available":    DataGuardAssociationSummaryLifecycleStateAvailable,
	"updating":     DataGuardAssociationSummaryLifecycleStateUpdating,
	"terminating":  DataGuardAssociationSummaryLifecycleStateTerminating,
	"terminated":   DataGuardAssociationSummaryLifecycleStateTerminated,
	"failed":       DataGuardAssociationSummaryLifecycleStateFailed,
	"upgrading":    DataGuardAssociationSummaryLifecycleStateUpgrading,
}

// GetDataGuardAssociationSummaryLifecycleStateEnumValues Enumerates the set of values for DataGuardAssociationSummaryLifecycleStateEnum
func GetDataGuardAssociationSummaryLifecycleStateEnumValues() []DataGuardAssociationSummaryLifecycleStateEnum {
	values := make([]DataGuardAssociationSummaryLifecycleStateEnum, 0)
	for _, v := range mappingDataGuardAssociationSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDataGuardAssociationSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for DataGuardAssociationSummaryLifecycleStateEnum
func GetDataGuardAssociationSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"UPDATING",
		"TERMINATING",
		"TERMINATED",
		"FAILED",
		"UPGRADING",
	}
}

// GetMappingDataGuardAssociationSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataGuardAssociationSummaryLifecycleStateEnum(val string) (DataGuardAssociationSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingDataGuardAssociationSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DataGuardAssociationSummaryPeerRoleEnum Enum with underlying type: string
type DataGuardAssociationSummaryPeerRoleEnum string

// Set of constants representing the allowable values for DataGuardAssociationSummaryPeerRoleEnum
const (
	DataGuardAssociationSummaryPeerRolePrimary         DataGuardAssociationSummaryPeerRoleEnum = "PRIMARY"
	DataGuardAssociationSummaryPeerRoleStandby         DataGuardAssociationSummaryPeerRoleEnum = "STANDBY"
	DataGuardAssociationSummaryPeerRoleDisabledStandby DataGuardAssociationSummaryPeerRoleEnum = "DISABLED_STANDBY"
)

var mappingDataGuardAssociationSummaryPeerRoleEnum = map[string]DataGuardAssociationSummaryPeerRoleEnum{
	"PRIMARY":          DataGuardAssociationSummaryPeerRolePrimary,
	"STANDBY":          DataGuardAssociationSummaryPeerRoleStandby,
	"DISABLED_STANDBY": DataGuardAssociationSummaryPeerRoleDisabledStandby,
}

var mappingDataGuardAssociationSummaryPeerRoleEnumLowerCase = map[string]DataGuardAssociationSummaryPeerRoleEnum{
	"primary":          DataGuardAssociationSummaryPeerRolePrimary,
	"standby":          DataGuardAssociationSummaryPeerRoleStandby,
	"disabled_standby": DataGuardAssociationSummaryPeerRoleDisabledStandby,
}

// GetDataGuardAssociationSummaryPeerRoleEnumValues Enumerates the set of values for DataGuardAssociationSummaryPeerRoleEnum
func GetDataGuardAssociationSummaryPeerRoleEnumValues() []DataGuardAssociationSummaryPeerRoleEnum {
	values := make([]DataGuardAssociationSummaryPeerRoleEnum, 0)
	for _, v := range mappingDataGuardAssociationSummaryPeerRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetDataGuardAssociationSummaryPeerRoleEnumStringValues Enumerates the set of values in String for DataGuardAssociationSummaryPeerRoleEnum
func GetDataGuardAssociationSummaryPeerRoleEnumStringValues() []string {
	return []string{
		"PRIMARY",
		"STANDBY",
		"DISABLED_STANDBY",
	}
}

// GetMappingDataGuardAssociationSummaryPeerRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataGuardAssociationSummaryPeerRoleEnum(val string) (DataGuardAssociationSummaryPeerRoleEnum, bool) {
	enum, ok := mappingDataGuardAssociationSummaryPeerRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DataGuardAssociationSummaryProtectionModeEnum Enum with underlying type: string
type DataGuardAssociationSummaryProtectionModeEnum string

// Set of constants representing the allowable values for DataGuardAssociationSummaryProtectionModeEnum
const (
	DataGuardAssociationSummaryProtectionModeAvailability DataGuardAssociationSummaryProtectionModeEnum = "MAXIMUM_AVAILABILITY"
	DataGuardAssociationSummaryProtectionModePerformance  DataGuardAssociationSummaryProtectionModeEnum = "MAXIMUM_PERFORMANCE"
	DataGuardAssociationSummaryProtectionModeProtection   DataGuardAssociationSummaryProtectionModeEnum = "MAXIMUM_PROTECTION"
)

var mappingDataGuardAssociationSummaryProtectionModeEnum = map[string]DataGuardAssociationSummaryProtectionModeEnum{
	"MAXIMUM_AVAILABILITY": DataGuardAssociationSummaryProtectionModeAvailability,
	"MAXIMUM_PERFORMANCE":  DataGuardAssociationSummaryProtectionModePerformance,
	"MAXIMUM_PROTECTION":   DataGuardAssociationSummaryProtectionModeProtection,
}

var mappingDataGuardAssociationSummaryProtectionModeEnumLowerCase = map[string]DataGuardAssociationSummaryProtectionModeEnum{
	"maximum_availability": DataGuardAssociationSummaryProtectionModeAvailability,
	"maximum_performance":  DataGuardAssociationSummaryProtectionModePerformance,
	"maximum_protection":   DataGuardAssociationSummaryProtectionModeProtection,
}

// GetDataGuardAssociationSummaryProtectionModeEnumValues Enumerates the set of values for DataGuardAssociationSummaryProtectionModeEnum
func GetDataGuardAssociationSummaryProtectionModeEnumValues() []DataGuardAssociationSummaryProtectionModeEnum {
	values := make([]DataGuardAssociationSummaryProtectionModeEnum, 0)
	for _, v := range mappingDataGuardAssociationSummaryProtectionModeEnum {
		values = append(values, v)
	}
	return values
}

// GetDataGuardAssociationSummaryProtectionModeEnumStringValues Enumerates the set of values in String for DataGuardAssociationSummaryProtectionModeEnum
func GetDataGuardAssociationSummaryProtectionModeEnumStringValues() []string {
	return []string{
		"MAXIMUM_AVAILABILITY",
		"MAXIMUM_PERFORMANCE",
		"MAXIMUM_PROTECTION",
	}
}

// GetMappingDataGuardAssociationSummaryProtectionModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataGuardAssociationSummaryProtectionModeEnum(val string) (DataGuardAssociationSummaryProtectionModeEnum, bool) {
	enum, ok := mappingDataGuardAssociationSummaryProtectionModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DataGuardAssociationSummaryTransportTypeEnum Enum with underlying type: string
type DataGuardAssociationSummaryTransportTypeEnum string

// Set of constants representing the allowable values for DataGuardAssociationSummaryTransportTypeEnum
const (
	DataGuardAssociationSummaryTransportTypeSync     DataGuardAssociationSummaryTransportTypeEnum = "SYNC"
	DataGuardAssociationSummaryTransportTypeAsync    DataGuardAssociationSummaryTransportTypeEnum = "ASYNC"
	DataGuardAssociationSummaryTransportTypeFastsync DataGuardAssociationSummaryTransportTypeEnum = "FASTSYNC"
)

var mappingDataGuardAssociationSummaryTransportTypeEnum = map[string]DataGuardAssociationSummaryTransportTypeEnum{
	"SYNC":     DataGuardAssociationSummaryTransportTypeSync,
	"ASYNC":    DataGuardAssociationSummaryTransportTypeAsync,
	"FASTSYNC": DataGuardAssociationSummaryTransportTypeFastsync,
}

var mappingDataGuardAssociationSummaryTransportTypeEnumLowerCase = map[string]DataGuardAssociationSummaryTransportTypeEnum{
	"sync":     DataGuardAssociationSummaryTransportTypeSync,
	"async":    DataGuardAssociationSummaryTransportTypeAsync,
	"fastsync": DataGuardAssociationSummaryTransportTypeFastsync,
}

// GetDataGuardAssociationSummaryTransportTypeEnumValues Enumerates the set of values for DataGuardAssociationSummaryTransportTypeEnum
func GetDataGuardAssociationSummaryTransportTypeEnumValues() []DataGuardAssociationSummaryTransportTypeEnum {
	values := make([]DataGuardAssociationSummaryTransportTypeEnum, 0)
	for _, v := range mappingDataGuardAssociationSummaryTransportTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDataGuardAssociationSummaryTransportTypeEnumStringValues Enumerates the set of values in String for DataGuardAssociationSummaryTransportTypeEnum
func GetDataGuardAssociationSummaryTransportTypeEnumStringValues() []string {
	return []string{
		"SYNC",
		"ASYNC",
		"FASTSYNC",
	}
}

// GetMappingDataGuardAssociationSummaryTransportTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataGuardAssociationSummaryTransportTypeEnum(val string) (DataGuardAssociationSummaryTransportTypeEnum, bool) {
	enum, ok := mappingDataGuardAssociationSummaryTransportTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
