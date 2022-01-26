// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// Datafile The details of a data file.
type Datafile struct {

	// The filename (including the path) of the data file or temp file.
	Name *string `mandatory:"true" json:"name"`

	// The status of the file. INVALID status is used when the file number is not in use, for example, a file in a tablespace that was removed.
	Status DatafileStatusEnum `mandatory:"false" json:"status,omitempty"`

	// The online status of the file.
	OnlineStatus DatafileOnlineStatusEnum `mandatory:"false" json:"onlineStatus,omitempty"`

	// Indicates whether the data file is auto-extensible.
	IsAutoExtensible *bool `mandatory:"false" json:"isAutoExtensible"`

	// The lost write protection status of the file.
	LostWriteProtect DatafileLostWriteProtectEnum `mandatory:"false" json:"lostWriteProtect,omitempty"`

	// Type of tablespace this file belongs to. If it's for a shared tablespace, for a local temporary tablespace for RIM (read-only) instances, or for local temporary tablespace for all instance types.
	Shared DatafileSharedEnum `mandatory:"false" json:"shared,omitempty"`

	// Instance ID of the instance to which the temp file belongs. This column has a NULL value for temp files that belong to shared tablespaces.
	InstanceId *float32 `mandatory:"false" json:"instanceId"`

	// The maximum file size in KB.
	MaxSizeKB *float32 `mandatory:"false" json:"maxSizeKB"`

	// The allocated file size in KB.
	AllocatedSizeKB *float32 `mandatory:"false" json:"allocatedSizeKB"`

	// The size of the file available for user data in KB. The actual size of the file minus the USER_BYTES value is used to store file-related metadata.
	UserSizeKB *float32 `mandatory:"false" json:"userSizeKB"`

	// The number of blocks used as auto-extension increment.
	IncrementBy *float32 `mandatory:"false" json:"incrementBy"`

	// The free space available in the data file in KB.
	FreeSpaceKB *float32 `mandatory:"false" json:"freeSpaceKB"`

	// The total space used in the data file in KB.
	UsedSpaceKB *float32 `mandatory:"false" json:"usedSpaceKB"`

	// The percentage of used space out of the maximum available space in the file.
	UsedPercentAvailable *float64 `mandatory:"false" json:"usedPercentAvailable"`

	// The percentage of used space out of the total allocated space in the file.
	UsedPercentAllocated *float64 `mandatory:"false" json:"usedPercentAllocated"`
}

func (m Datafile) String() string {
	return common.PointerString(m)
}

// DatafileStatusEnum Enum with underlying type: string
type DatafileStatusEnum string

// Set of constants representing the allowable values for DatafileStatusEnum
const (
	DatafileStatusAvailable DatafileStatusEnum = "AVAILABLE"
	DatafileStatusInvalid   DatafileStatusEnum = "INVALID"
	DatafileStatusOffline   DatafileStatusEnum = "OFFLINE"
	DatafileStatusOnline    DatafileStatusEnum = "ONLINE"
	DatafileStatusUnknown   DatafileStatusEnum = "UNKNOWN"
)

var mappingDatafileStatus = map[string]DatafileStatusEnum{
	"AVAILABLE": DatafileStatusAvailable,
	"INVALID":   DatafileStatusInvalid,
	"OFFLINE":   DatafileStatusOffline,
	"ONLINE":    DatafileStatusOnline,
	"UNKNOWN":   DatafileStatusUnknown,
}

// GetDatafileStatusEnumValues Enumerates the set of values for DatafileStatusEnum
func GetDatafileStatusEnumValues() []DatafileStatusEnum {
	values := make([]DatafileStatusEnum, 0)
	for _, v := range mappingDatafileStatus {
		values = append(values, v)
	}
	return values
}

// DatafileOnlineStatusEnum Enum with underlying type: string
type DatafileOnlineStatusEnum string

// Set of constants representing the allowable values for DatafileOnlineStatusEnum
const (
	DatafileOnlineStatusSysoff  DatafileOnlineStatusEnum = "SYSOFF"
	DatafileOnlineStatusSystem  DatafileOnlineStatusEnum = "SYSTEM"
	DatafileOnlineStatusOffline DatafileOnlineStatusEnum = "OFFLINE"
	DatafileOnlineStatusOnline  DatafileOnlineStatusEnum = "ONLINE"
	DatafileOnlineStatusRecover DatafileOnlineStatusEnum = "RECOVER"
)

var mappingDatafileOnlineStatus = map[string]DatafileOnlineStatusEnum{
	"SYSOFF":  DatafileOnlineStatusSysoff,
	"SYSTEM":  DatafileOnlineStatusSystem,
	"OFFLINE": DatafileOnlineStatusOffline,
	"ONLINE":  DatafileOnlineStatusOnline,
	"RECOVER": DatafileOnlineStatusRecover,
}

// GetDatafileOnlineStatusEnumValues Enumerates the set of values for DatafileOnlineStatusEnum
func GetDatafileOnlineStatusEnumValues() []DatafileOnlineStatusEnum {
	values := make([]DatafileOnlineStatusEnum, 0)
	for _, v := range mappingDatafileOnlineStatus {
		values = append(values, v)
	}
	return values
}

// DatafileLostWriteProtectEnum Enum with underlying type: string
type DatafileLostWriteProtectEnum string

// Set of constants representing the allowable values for DatafileLostWriteProtectEnum
const (
	DatafileLostWriteProtectEnabled    DatafileLostWriteProtectEnum = "ENABLED"
	DatafileLostWriteProtectProtectOff DatafileLostWriteProtectEnum = "PROTECT_OFF"
	DatafileLostWriteProtectSuspend    DatafileLostWriteProtectEnum = "SUSPEND"
)

var mappingDatafileLostWriteProtect = map[string]DatafileLostWriteProtectEnum{
	"ENABLED":     DatafileLostWriteProtectEnabled,
	"PROTECT_OFF": DatafileLostWriteProtectProtectOff,
	"SUSPEND":     DatafileLostWriteProtectSuspend,
}

// GetDatafileLostWriteProtectEnumValues Enumerates the set of values for DatafileLostWriteProtectEnum
func GetDatafileLostWriteProtectEnumValues() []DatafileLostWriteProtectEnum {
	values := make([]DatafileLostWriteProtectEnum, 0)
	for _, v := range mappingDatafileLostWriteProtect {
		values = append(values, v)
	}
	return values
}

// DatafileSharedEnum Enum with underlying type: string
type DatafileSharedEnum string

// Set of constants representing the allowable values for DatafileSharedEnum
const (
	DatafileSharedShared      DatafileSharedEnum = "SHARED"
	DatafileSharedLocalForRim DatafileSharedEnum = "LOCAL_FOR_RIM"
	DatafileSharedLocalForAll DatafileSharedEnum = "LOCAL_FOR_ALL"
)

var mappingDatafileShared = map[string]DatafileSharedEnum{
	"SHARED":        DatafileSharedShared,
	"LOCAL_FOR_RIM": DatafileSharedLocalForRim,
	"LOCAL_FOR_ALL": DatafileSharedLocalForAll,
}

// GetDatafileSharedEnumValues Enumerates the set of values for DatafileSharedEnum
func GetDatafileSharedEnumValues() []DatafileSharedEnum {
	values := make([]DatafileSharedEnum, 0)
	for _, v := range mappingDatafileShared {
		values = append(values, v)
	}
	return values
}
