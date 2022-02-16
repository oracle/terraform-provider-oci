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
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// DataGuardAssociation The representation of DataGuardAssociation
type DataGuardAssociation struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Data Guard association.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the reporting database.
	DatabaseId *string `mandatory:"true" json:"databaseId"`

	// The role of the reporting database in this Data Guard association.
	Role DataGuardAssociationRoleEnum `mandatory:"true" json:"role"`

	// The current state of the Data Guard association.
	LifecycleState DataGuardAssociationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the DB system containing the associated
	// peer database.
	PeerDbSystemId *string `mandatory:"true" json:"peerDbSystemId"`

	// The role of the peer database in this Data Guard association.
	PeerRole DataGuardAssociationPeerRoleEnum `mandatory:"true" json:"peerRole"`

	// The protection mode of this Data Guard association. For more information, see
	// Oracle Data Guard Protection Modes (http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000)
	// in the Oracle Data Guard documentation.
	ProtectionMode DataGuardAssociationProtectionModeEnum `mandatory:"true" json:"protectionMode"`

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
	TransportType DataGuardAssociationTransportTypeEnum `mandatory:"false" json:"transportType,omitempty"`

	// The date and time the Data Guard association was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// True if active Data Guard is enabled.
	IsActiveDataGuardEnabled *bool `mandatory:"false" json:"isActiveDataGuardEnabled"`
}

func (m DataGuardAssociation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataGuardAssociation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDataGuardAssociationRoleEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetDataGuardAssociationRoleEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDataGuardAssociationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDataGuardAssociationLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDataGuardAssociationPeerRoleEnum(string(m.PeerRole)); !ok && m.PeerRole != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PeerRole: %s. Supported values are: %s.", m.PeerRole, strings.Join(GetDataGuardAssociationPeerRoleEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDataGuardAssociationProtectionModeEnum(string(m.ProtectionMode)); !ok && m.ProtectionMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ProtectionMode: %s. Supported values are: %s.", m.ProtectionMode, strings.Join(GetDataGuardAssociationProtectionModeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDataGuardAssociationTransportTypeEnum(string(m.TransportType)); !ok && m.TransportType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TransportType: %s. Supported values are: %s.", m.TransportType, strings.Join(GetDataGuardAssociationTransportTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DataGuardAssociationRoleEnum Enum with underlying type: string
type DataGuardAssociationRoleEnum string

// Set of constants representing the allowable values for DataGuardAssociationRoleEnum
const (
	DataGuardAssociationRolePrimary         DataGuardAssociationRoleEnum = "PRIMARY"
	DataGuardAssociationRoleStandby         DataGuardAssociationRoleEnum = "STANDBY"
	DataGuardAssociationRoleDisabledStandby DataGuardAssociationRoleEnum = "DISABLED_STANDBY"
)

var mappingDataGuardAssociationRoleEnum = map[string]DataGuardAssociationRoleEnum{
	"PRIMARY":          DataGuardAssociationRolePrimary,
	"STANDBY":          DataGuardAssociationRoleStandby,
	"DISABLED_STANDBY": DataGuardAssociationRoleDisabledStandby,
}

// GetDataGuardAssociationRoleEnumValues Enumerates the set of values for DataGuardAssociationRoleEnum
func GetDataGuardAssociationRoleEnumValues() []DataGuardAssociationRoleEnum {
	values := make([]DataGuardAssociationRoleEnum, 0)
	for _, v := range mappingDataGuardAssociationRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetDataGuardAssociationRoleEnumStringValues Enumerates the set of values in String for DataGuardAssociationRoleEnum
func GetDataGuardAssociationRoleEnumStringValues() []string {
	return []string{
		"PRIMARY",
		"STANDBY",
		"DISABLED_STANDBY",
	}
}

// GetMappingDataGuardAssociationRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataGuardAssociationRoleEnum(val string) (DataGuardAssociationRoleEnum, bool) {
	mappingDataGuardAssociationRoleEnumIgnoreCase := make(map[string]DataGuardAssociationRoleEnum)
	for k, v := range mappingDataGuardAssociationRoleEnum {
		mappingDataGuardAssociationRoleEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingDataGuardAssociationRoleEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// DataGuardAssociationLifecycleStateEnum Enum with underlying type: string
type DataGuardAssociationLifecycleStateEnum string

// Set of constants representing the allowable values for DataGuardAssociationLifecycleStateEnum
const (
	DataGuardAssociationLifecycleStateProvisioning DataGuardAssociationLifecycleStateEnum = "PROVISIONING"
	DataGuardAssociationLifecycleStateAvailable    DataGuardAssociationLifecycleStateEnum = "AVAILABLE"
	DataGuardAssociationLifecycleStateUpdating     DataGuardAssociationLifecycleStateEnum = "UPDATING"
	DataGuardAssociationLifecycleStateTerminating  DataGuardAssociationLifecycleStateEnum = "TERMINATING"
	DataGuardAssociationLifecycleStateTerminated   DataGuardAssociationLifecycleStateEnum = "TERMINATED"
	DataGuardAssociationLifecycleStateFailed       DataGuardAssociationLifecycleStateEnum = "FAILED"
)

var mappingDataGuardAssociationLifecycleStateEnum = map[string]DataGuardAssociationLifecycleStateEnum{
	"PROVISIONING": DataGuardAssociationLifecycleStateProvisioning,
	"AVAILABLE":    DataGuardAssociationLifecycleStateAvailable,
	"UPDATING":     DataGuardAssociationLifecycleStateUpdating,
	"TERMINATING":  DataGuardAssociationLifecycleStateTerminating,
	"TERMINATED":   DataGuardAssociationLifecycleStateTerminated,
	"FAILED":       DataGuardAssociationLifecycleStateFailed,
}

// GetDataGuardAssociationLifecycleStateEnumValues Enumerates the set of values for DataGuardAssociationLifecycleStateEnum
func GetDataGuardAssociationLifecycleStateEnumValues() []DataGuardAssociationLifecycleStateEnum {
	values := make([]DataGuardAssociationLifecycleStateEnum, 0)
	for _, v := range mappingDataGuardAssociationLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDataGuardAssociationLifecycleStateEnumStringValues Enumerates the set of values in String for DataGuardAssociationLifecycleStateEnum
func GetDataGuardAssociationLifecycleStateEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"UPDATING",
		"TERMINATING",
		"TERMINATED",
		"FAILED",
	}
}

// GetMappingDataGuardAssociationLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataGuardAssociationLifecycleStateEnum(val string) (DataGuardAssociationLifecycleStateEnum, bool) {
	mappingDataGuardAssociationLifecycleStateEnumIgnoreCase := make(map[string]DataGuardAssociationLifecycleStateEnum)
	for k, v := range mappingDataGuardAssociationLifecycleStateEnum {
		mappingDataGuardAssociationLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingDataGuardAssociationLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// DataGuardAssociationPeerRoleEnum Enum with underlying type: string
type DataGuardAssociationPeerRoleEnum string

// Set of constants representing the allowable values for DataGuardAssociationPeerRoleEnum
const (
	DataGuardAssociationPeerRolePrimary         DataGuardAssociationPeerRoleEnum = "PRIMARY"
	DataGuardAssociationPeerRoleStandby         DataGuardAssociationPeerRoleEnum = "STANDBY"
	DataGuardAssociationPeerRoleDisabledStandby DataGuardAssociationPeerRoleEnum = "DISABLED_STANDBY"
)

var mappingDataGuardAssociationPeerRoleEnum = map[string]DataGuardAssociationPeerRoleEnum{
	"PRIMARY":          DataGuardAssociationPeerRolePrimary,
	"STANDBY":          DataGuardAssociationPeerRoleStandby,
	"DISABLED_STANDBY": DataGuardAssociationPeerRoleDisabledStandby,
}

// GetDataGuardAssociationPeerRoleEnumValues Enumerates the set of values for DataGuardAssociationPeerRoleEnum
func GetDataGuardAssociationPeerRoleEnumValues() []DataGuardAssociationPeerRoleEnum {
	values := make([]DataGuardAssociationPeerRoleEnum, 0)
	for _, v := range mappingDataGuardAssociationPeerRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetDataGuardAssociationPeerRoleEnumStringValues Enumerates the set of values in String for DataGuardAssociationPeerRoleEnum
func GetDataGuardAssociationPeerRoleEnumStringValues() []string {
	return []string{
		"PRIMARY",
		"STANDBY",
		"DISABLED_STANDBY",
	}
}

// GetMappingDataGuardAssociationPeerRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataGuardAssociationPeerRoleEnum(val string) (DataGuardAssociationPeerRoleEnum, bool) {
	mappingDataGuardAssociationPeerRoleEnumIgnoreCase := make(map[string]DataGuardAssociationPeerRoleEnum)
	for k, v := range mappingDataGuardAssociationPeerRoleEnum {
		mappingDataGuardAssociationPeerRoleEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingDataGuardAssociationPeerRoleEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// DataGuardAssociationProtectionModeEnum Enum with underlying type: string
type DataGuardAssociationProtectionModeEnum string

// Set of constants representing the allowable values for DataGuardAssociationProtectionModeEnum
const (
	DataGuardAssociationProtectionModeAvailability DataGuardAssociationProtectionModeEnum = "MAXIMUM_AVAILABILITY"
	DataGuardAssociationProtectionModePerformance  DataGuardAssociationProtectionModeEnum = "MAXIMUM_PERFORMANCE"
	DataGuardAssociationProtectionModeProtection   DataGuardAssociationProtectionModeEnum = "MAXIMUM_PROTECTION"
)

var mappingDataGuardAssociationProtectionModeEnum = map[string]DataGuardAssociationProtectionModeEnum{
	"MAXIMUM_AVAILABILITY": DataGuardAssociationProtectionModeAvailability,
	"MAXIMUM_PERFORMANCE":  DataGuardAssociationProtectionModePerformance,
	"MAXIMUM_PROTECTION":   DataGuardAssociationProtectionModeProtection,
}

// GetDataGuardAssociationProtectionModeEnumValues Enumerates the set of values for DataGuardAssociationProtectionModeEnum
func GetDataGuardAssociationProtectionModeEnumValues() []DataGuardAssociationProtectionModeEnum {
	values := make([]DataGuardAssociationProtectionModeEnum, 0)
	for _, v := range mappingDataGuardAssociationProtectionModeEnum {
		values = append(values, v)
	}
	return values
}

// GetDataGuardAssociationProtectionModeEnumStringValues Enumerates the set of values in String for DataGuardAssociationProtectionModeEnum
func GetDataGuardAssociationProtectionModeEnumStringValues() []string {
	return []string{
		"MAXIMUM_AVAILABILITY",
		"MAXIMUM_PERFORMANCE",
		"MAXIMUM_PROTECTION",
	}
}

// GetMappingDataGuardAssociationProtectionModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataGuardAssociationProtectionModeEnum(val string) (DataGuardAssociationProtectionModeEnum, bool) {
	mappingDataGuardAssociationProtectionModeEnumIgnoreCase := make(map[string]DataGuardAssociationProtectionModeEnum)
	for k, v := range mappingDataGuardAssociationProtectionModeEnum {
		mappingDataGuardAssociationProtectionModeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingDataGuardAssociationProtectionModeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// DataGuardAssociationTransportTypeEnum Enum with underlying type: string
type DataGuardAssociationTransportTypeEnum string

// Set of constants representing the allowable values for DataGuardAssociationTransportTypeEnum
const (
	DataGuardAssociationTransportTypeSync     DataGuardAssociationTransportTypeEnum = "SYNC"
	DataGuardAssociationTransportTypeAsync    DataGuardAssociationTransportTypeEnum = "ASYNC"
	DataGuardAssociationTransportTypeFastsync DataGuardAssociationTransportTypeEnum = "FASTSYNC"
)

var mappingDataGuardAssociationTransportTypeEnum = map[string]DataGuardAssociationTransportTypeEnum{
	"SYNC":     DataGuardAssociationTransportTypeSync,
	"ASYNC":    DataGuardAssociationTransportTypeAsync,
	"FASTSYNC": DataGuardAssociationTransportTypeFastsync,
}

// GetDataGuardAssociationTransportTypeEnumValues Enumerates the set of values for DataGuardAssociationTransportTypeEnum
func GetDataGuardAssociationTransportTypeEnumValues() []DataGuardAssociationTransportTypeEnum {
	values := make([]DataGuardAssociationTransportTypeEnum, 0)
	for _, v := range mappingDataGuardAssociationTransportTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDataGuardAssociationTransportTypeEnumStringValues Enumerates the set of values in String for DataGuardAssociationTransportTypeEnum
func GetDataGuardAssociationTransportTypeEnumStringValues() []string {
	return []string{
		"SYNC",
		"ASYNC",
		"FASTSYNC",
	}
}

// GetMappingDataGuardAssociationTransportTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataGuardAssociationTransportTypeEnum(val string) (DataGuardAssociationTransportTypeEnum, bool) {
	mappingDataGuardAssociationTransportTypeEnumIgnoreCase := make(map[string]DataGuardAssociationTransportTypeEnum)
	for k, v := range mappingDataGuardAssociationTransportTypeEnum {
		mappingDataGuardAssociationTransportTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingDataGuardAssociationTransportTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
