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

// DataGuardGroupMember The member of a Data Guard group. Represents either a PRIMARY or a STANDBY Data Guard instance.
type DataGuardGroupMember struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB system, Cloud VM cluster or VM cluster.
	DbSystemId *string `mandatory:"true" json:"dbSystemId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database.
	DatabaseId *string `mandatory:"true" json:"databaseId"`

	// The role of the reporting database in this Data Guard association.
	Role DataGuardGroupMemberRoleEnum `mandatory:"true" json:"role"`

	// The lag time between updates to the primary database and application of the redo data on the standby database,
	// as computed by the reporting database.
	// Example: `1 second`
	ApplyLag *string `mandatory:"false" json:"applyLag"`

	// The rate at which redo logs are synced between the associated databases.
	// Example: `102.96 MByte/s`
	ApplyRate *string `mandatory:"false" json:"applyRate"`

	// The rate at which redo logs are transported between the associated databases.
	// Example: `1 second`
	TransportLag *string `mandatory:"false" json:"transportLag"`

	// The date and time when last redo transport has been done.
	TransportLagRefresh *string `mandatory:"false" json:"transportLagRefresh"`

	// The redo transport type to use for this Data Guard association.  Valid values depend on the specified `protectionMode`:
	// * MAXIMUM_AVAILABILITY - SYNC or FASTSYNC
	// * MAXIMUM_PERFORMANCE - ASYNC
	// * MAXIMUM_PROTECTION - SYNC
	// For more information, see
	// Redo Transport Services (http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-redo-transport-services.htm#SBYDB00400)
	// in the Oracle Data Guard documentation.
	// **IMPORTANT** - The only transport type currently supported by the Database service is ASYNC.
	TransportType DataGuardGroupMemberTransportTypeEnum `mandatory:"false" json:"transportType,omitempty"`

	// True if active Data Guard is enabled.
	IsActiveDataGuardEnabled *bool `mandatory:"false" json:"isActiveDataGuardEnabled"`

	// The switchover readiness status of the Data Guard member.
	SwitchoverReadiness DataGuardGroupMemberSwitchoverReadinessEnum `mandatory:"false" json:"switchoverReadiness,omitempty"`

	// The message explaining switchover readiness status.
	// Example: `Address failed checks to avoid extended downtime.`
	SwitchoverReadinessMessage *string `mandatory:"false" json:"switchoverReadinessMessage"`

	// The failover readiness status of the Data Guard member.
	FailoverReadiness DataGuardGroupMemberFailoverReadinessEnum `mandatory:"false" json:"failoverReadiness,omitempty"`

	// The message explaining failover readiness status.
	// Example: `This standby database is not failover ready.`
	FailoverReadinessMessage *string `mandatory:"false" json:"failoverReadinessMessage"`

	// The Data loss exposure is the redo transport lag between the primary and standby databases.
	// Example: `2 seconds`
	DataLossExposure *string `mandatory:"false" json:"dataLossExposure"`

	// The date and time when the last successful Data Guard refresh occurred.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m DataGuardGroupMember) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataGuardGroupMember) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDataGuardGroupMemberRoleEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetDataGuardGroupMemberRoleEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDataGuardGroupMemberTransportTypeEnum(string(m.TransportType)); !ok && m.TransportType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TransportType: %s. Supported values are: %s.", m.TransportType, strings.Join(GetDataGuardGroupMemberTransportTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDataGuardGroupMemberSwitchoverReadinessEnum(string(m.SwitchoverReadiness)); !ok && m.SwitchoverReadiness != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SwitchoverReadiness: %s. Supported values are: %s.", m.SwitchoverReadiness, strings.Join(GetDataGuardGroupMemberSwitchoverReadinessEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDataGuardGroupMemberFailoverReadinessEnum(string(m.FailoverReadiness)); !ok && m.FailoverReadiness != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FailoverReadiness: %s. Supported values are: %s.", m.FailoverReadiness, strings.Join(GetDataGuardGroupMemberFailoverReadinessEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DataGuardGroupMemberRoleEnum Enum with underlying type: string
type DataGuardGroupMemberRoleEnum string

// Set of constants representing the allowable values for DataGuardGroupMemberRoleEnum
const (
	DataGuardGroupMemberRolePrimary         DataGuardGroupMemberRoleEnum = "PRIMARY"
	DataGuardGroupMemberRoleStandby         DataGuardGroupMemberRoleEnum = "STANDBY"
	DataGuardGroupMemberRoleDisabledStandby DataGuardGroupMemberRoleEnum = "DISABLED_STANDBY"
	DataGuardGroupMemberRoleSnapshotStandby DataGuardGroupMemberRoleEnum = "SNAPSHOT_STANDBY"
)

var mappingDataGuardGroupMemberRoleEnum = map[string]DataGuardGroupMemberRoleEnum{
	"PRIMARY":          DataGuardGroupMemberRolePrimary,
	"STANDBY":          DataGuardGroupMemberRoleStandby,
	"DISABLED_STANDBY": DataGuardGroupMemberRoleDisabledStandby,
	"SNAPSHOT_STANDBY": DataGuardGroupMemberRoleSnapshotStandby,
}

var mappingDataGuardGroupMemberRoleEnumLowerCase = map[string]DataGuardGroupMemberRoleEnum{
	"primary":          DataGuardGroupMemberRolePrimary,
	"standby":          DataGuardGroupMemberRoleStandby,
	"disabled_standby": DataGuardGroupMemberRoleDisabledStandby,
	"snapshot_standby": DataGuardGroupMemberRoleSnapshotStandby,
}

// GetDataGuardGroupMemberRoleEnumValues Enumerates the set of values for DataGuardGroupMemberRoleEnum
func GetDataGuardGroupMemberRoleEnumValues() []DataGuardGroupMemberRoleEnum {
	values := make([]DataGuardGroupMemberRoleEnum, 0)
	for _, v := range mappingDataGuardGroupMemberRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetDataGuardGroupMemberRoleEnumStringValues Enumerates the set of values in String for DataGuardGroupMemberRoleEnum
func GetDataGuardGroupMemberRoleEnumStringValues() []string {
	return []string{
		"PRIMARY",
		"STANDBY",
		"DISABLED_STANDBY",
		"SNAPSHOT_STANDBY",
	}
}

// GetMappingDataGuardGroupMemberRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataGuardGroupMemberRoleEnum(val string) (DataGuardGroupMemberRoleEnum, bool) {
	enum, ok := mappingDataGuardGroupMemberRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DataGuardGroupMemberTransportTypeEnum Enum with underlying type: string
type DataGuardGroupMemberTransportTypeEnum string

// Set of constants representing the allowable values for DataGuardGroupMemberTransportTypeEnum
const (
	DataGuardGroupMemberTransportTypeSync     DataGuardGroupMemberTransportTypeEnum = "SYNC"
	DataGuardGroupMemberTransportTypeAsync    DataGuardGroupMemberTransportTypeEnum = "ASYNC"
	DataGuardGroupMemberTransportTypeFastsync DataGuardGroupMemberTransportTypeEnum = "FASTSYNC"
)

var mappingDataGuardGroupMemberTransportTypeEnum = map[string]DataGuardGroupMemberTransportTypeEnum{
	"SYNC":     DataGuardGroupMemberTransportTypeSync,
	"ASYNC":    DataGuardGroupMemberTransportTypeAsync,
	"FASTSYNC": DataGuardGroupMemberTransportTypeFastsync,
}

var mappingDataGuardGroupMemberTransportTypeEnumLowerCase = map[string]DataGuardGroupMemberTransportTypeEnum{
	"sync":     DataGuardGroupMemberTransportTypeSync,
	"async":    DataGuardGroupMemberTransportTypeAsync,
	"fastsync": DataGuardGroupMemberTransportTypeFastsync,
}

// GetDataGuardGroupMemberTransportTypeEnumValues Enumerates the set of values for DataGuardGroupMemberTransportTypeEnum
func GetDataGuardGroupMemberTransportTypeEnumValues() []DataGuardGroupMemberTransportTypeEnum {
	values := make([]DataGuardGroupMemberTransportTypeEnum, 0)
	for _, v := range mappingDataGuardGroupMemberTransportTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDataGuardGroupMemberTransportTypeEnumStringValues Enumerates the set of values in String for DataGuardGroupMemberTransportTypeEnum
func GetDataGuardGroupMemberTransportTypeEnumStringValues() []string {
	return []string{
		"SYNC",
		"ASYNC",
		"FASTSYNC",
	}
}

// GetMappingDataGuardGroupMemberTransportTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataGuardGroupMemberTransportTypeEnum(val string) (DataGuardGroupMemberTransportTypeEnum, bool) {
	enum, ok := mappingDataGuardGroupMemberTransportTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DataGuardGroupMemberSwitchoverReadinessEnum Enum with underlying type: string
type DataGuardGroupMemberSwitchoverReadinessEnum string

// Set of constants representing the allowable values for DataGuardGroupMemberSwitchoverReadinessEnum
const (
	DataGuardGroupMemberSwitchoverReadinessHealthy  DataGuardGroupMemberSwitchoverReadinessEnum = "HEALTHY"
	DataGuardGroupMemberSwitchoverReadinessWarning  DataGuardGroupMemberSwitchoverReadinessEnum = "WARNING"
	DataGuardGroupMemberSwitchoverReadinessCritical DataGuardGroupMemberSwitchoverReadinessEnum = "CRITICAL"
)

var mappingDataGuardGroupMemberSwitchoverReadinessEnum = map[string]DataGuardGroupMemberSwitchoverReadinessEnum{
	"HEALTHY":  DataGuardGroupMemberSwitchoverReadinessHealthy,
	"WARNING":  DataGuardGroupMemberSwitchoverReadinessWarning,
	"CRITICAL": DataGuardGroupMemberSwitchoverReadinessCritical,
}

var mappingDataGuardGroupMemberSwitchoverReadinessEnumLowerCase = map[string]DataGuardGroupMemberSwitchoverReadinessEnum{
	"healthy":  DataGuardGroupMemberSwitchoverReadinessHealthy,
	"warning":  DataGuardGroupMemberSwitchoverReadinessWarning,
	"critical": DataGuardGroupMemberSwitchoverReadinessCritical,
}

// GetDataGuardGroupMemberSwitchoverReadinessEnumValues Enumerates the set of values for DataGuardGroupMemberSwitchoverReadinessEnum
func GetDataGuardGroupMemberSwitchoverReadinessEnumValues() []DataGuardGroupMemberSwitchoverReadinessEnum {
	values := make([]DataGuardGroupMemberSwitchoverReadinessEnum, 0)
	for _, v := range mappingDataGuardGroupMemberSwitchoverReadinessEnum {
		values = append(values, v)
	}
	return values
}

// GetDataGuardGroupMemberSwitchoverReadinessEnumStringValues Enumerates the set of values in String for DataGuardGroupMemberSwitchoverReadinessEnum
func GetDataGuardGroupMemberSwitchoverReadinessEnumStringValues() []string {
	return []string{
		"HEALTHY",
		"WARNING",
		"CRITICAL",
	}
}

// GetMappingDataGuardGroupMemberSwitchoverReadinessEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataGuardGroupMemberSwitchoverReadinessEnum(val string) (DataGuardGroupMemberSwitchoverReadinessEnum, bool) {
	enum, ok := mappingDataGuardGroupMemberSwitchoverReadinessEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DataGuardGroupMemberFailoverReadinessEnum Enum with underlying type: string
type DataGuardGroupMemberFailoverReadinessEnum string

// Set of constants representing the allowable values for DataGuardGroupMemberFailoverReadinessEnum
const (
	DataGuardGroupMemberFailoverReadinessHealthy  DataGuardGroupMemberFailoverReadinessEnum = "HEALTHY"
	DataGuardGroupMemberFailoverReadinessWarning  DataGuardGroupMemberFailoverReadinessEnum = "WARNING"
	DataGuardGroupMemberFailoverReadinessCritical DataGuardGroupMemberFailoverReadinessEnum = "CRITICAL"
)

var mappingDataGuardGroupMemberFailoverReadinessEnum = map[string]DataGuardGroupMemberFailoverReadinessEnum{
	"HEALTHY":  DataGuardGroupMemberFailoverReadinessHealthy,
	"WARNING":  DataGuardGroupMemberFailoverReadinessWarning,
	"CRITICAL": DataGuardGroupMemberFailoverReadinessCritical,
}

var mappingDataGuardGroupMemberFailoverReadinessEnumLowerCase = map[string]DataGuardGroupMemberFailoverReadinessEnum{
	"healthy":  DataGuardGroupMemberFailoverReadinessHealthy,
	"warning":  DataGuardGroupMemberFailoverReadinessWarning,
	"critical": DataGuardGroupMemberFailoverReadinessCritical,
}

// GetDataGuardGroupMemberFailoverReadinessEnumValues Enumerates the set of values for DataGuardGroupMemberFailoverReadinessEnum
func GetDataGuardGroupMemberFailoverReadinessEnumValues() []DataGuardGroupMemberFailoverReadinessEnum {
	values := make([]DataGuardGroupMemberFailoverReadinessEnum, 0)
	for _, v := range mappingDataGuardGroupMemberFailoverReadinessEnum {
		values = append(values, v)
	}
	return values
}

// GetDataGuardGroupMemberFailoverReadinessEnumStringValues Enumerates the set of values in String for DataGuardGroupMemberFailoverReadinessEnum
func GetDataGuardGroupMemberFailoverReadinessEnumStringValues() []string {
	return []string{
		"HEALTHY",
		"WARNING",
		"CRITICAL",
	}
}

// GetMappingDataGuardGroupMemberFailoverReadinessEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataGuardGroupMemberFailoverReadinessEnum(val string) (DataGuardGroupMemberFailoverReadinessEnum, bool) {
	enum, ok := mappingDataGuardGroupMemberFailoverReadinessEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
