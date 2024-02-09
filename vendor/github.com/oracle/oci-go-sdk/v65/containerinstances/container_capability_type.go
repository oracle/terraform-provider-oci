// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Instance API
//
// A description of the Container Instance API
//

package containerinstances

import (
	"strings"
)

// ContainerCapabilityTypeEnum Enum with underlying type: string
type ContainerCapabilityTypeEnum string

// Set of constants representing the allowable values for ContainerCapabilityTypeEnum
const (
	ContainerCapabilityTypeCapChown          ContainerCapabilityTypeEnum = "CAP_CHOWN"
	ContainerCapabilityTypeCapDacOverride    ContainerCapabilityTypeEnum = "CAP_DAC_OVERRIDE"
	ContainerCapabilityTypeCapFsetid         ContainerCapabilityTypeEnum = "CAP_FSETID"
	ContainerCapabilityTypeCapFowner         ContainerCapabilityTypeEnum = "CAP_FOWNER"
	ContainerCapabilityTypeCapMknod          ContainerCapabilityTypeEnum = "CAP_MKNOD"
	ContainerCapabilityTypeCapNetRaw         ContainerCapabilityTypeEnum = "CAP_NET_RAW"
	ContainerCapabilityTypeCapSetgid         ContainerCapabilityTypeEnum = "CAP_SETGID"
	ContainerCapabilityTypeCapSetuid         ContainerCapabilityTypeEnum = "CAP_SETUID"
	ContainerCapabilityTypeCapSetfcap        ContainerCapabilityTypeEnum = "CAP_SETFCAP"
	ContainerCapabilityTypeCapSetpcap        ContainerCapabilityTypeEnum = "CAP_SETPCAP"
	ContainerCapabilityTypeCapNetBindService ContainerCapabilityTypeEnum = "CAP_NET_BIND_SERVICE"
	ContainerCapabilityTypeCapSysChroot      ContainerCapabilityTypeEnum = "CAP_SYS_CHROOT"
	ContainerCapabilityTypeCapKill           ContainerCapabilityTypeEnum = "CAP_KILL"
	ContainerCapabilityTypeCapAuditWrite     ContainerCapabilityTypeEnum = "CAP_AUDIT_WRITE"
	ContainerCapabilityTypeAll               ContainerCapabilityTypeEnum = "ALL"
)

var mappingContainerCapabilityTypeEnum = map[string]ContainerCapabilityTypeEnum{
	"CAP_CHOWN":            ContainerCapabilityTypeCapChown,
	"CAP_DAC_OVERRIDE":     ContainerCapabilityTypeCapDacOverride,
	"CAP_FSETID":           ContainerCapabilityTypeCapFsetid,
	"CAP_FOWNER":           ContainerCapabilityTypeCapFowner,
	"CAP_MKNOD":            ContainerCapabilityTypeCapMknod,
	"CAP_NET_RAW":          ContainerCapabilityTypeCapNetRaw,
	"CAP_SETGID":           ContainerCapabilityTypeCapSetgid,
	"CAP_SETUID":           ContainerCapabilityTypeCapSetuid,
	"CAP_SETFCAP":          ContainerCapabilityTypeCapSetfcap,
	"CAP_SETPCAP":          ContainerCapabilityTypeCapSetpcap,
	"CAP_NET_BIND_SERVICE": ContainerCapabilityTypeCapNetBindService,
	"CAP_SYS_CHROOT":       ContainerCapabilityTypeCapSysChroot,
	"CAP_KILL":             ContainerCapabilityTypeCapKill,
	"CAP_AUDIT_WRITE":      ContainerCapabilityTypeCapAuditWrite,
	"ALL":                  ContainerCapabilityTypeAll,
}

var mappingContainerCapabilityTypeEnumLowerCase = map[string]ContainerCapabilityTypeEnum{
	"cap_chown":            ContainerCapabilityTypeCapChown,
	"cap_dac_override":     ContainerCapabilityTypeCapDacOverride,
	"cap_fsetid":           ContainerCapabilityTypeCapFsetid,
	"cap_fowner":           ContainerCapabilityTypeCapFowner,
	"cap_mknod":            ContainerCapabilityTypeCapMknod,
	"cap_net_raw":          ContainerCapabilityTypeCapNetRaw,
	"cap_setgid":           ContainerCapabilityTypeCapSetgid,
	"cap_setuid":           ContainerCapabilityTypeCapSetuid,
	"cap_setfcap":          ContainerCapabilityTypeCapSetfcap,
	"cap_setpcap":          ContainerCapabilityTypeCapSetpcap,
	"cap_net_bind_service": ContainerCapabilityTypeCapNetBindService,
	"cap_sys_chroot":       ContainerCapabilityTypeCapSysChroot,
	"cap_kill":             ContainerCapabilityTypeCapKill,
	"cap_audit_write":      ContainerCapabilityTypeCapAuditWrite,
	"all":                  ContainerCapabilityTypeAll,
}

// GetContainerCapabilityTypeEnumValues Enumerates the set of values for ContainerCapabilityTypeEnum
func GetContainerCapabilityTypeEnumValues() []ContainerCapabilityTypeEnum {
	values := make([]ContainerCapabilityTypeEnum, 0)
	for _, v := range mappingContainerCapabilityTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetContainerCapabilityTypeEnumStringValues Enumerates the set of values in String for ContainerCapabilityTypeEnum
func GetContainerCapabilityTypeEnumStringValues() []string {
	return []string{
		"CAP_CHOWN",
		"CAP_DAC_OVERRIDE",
		"CAP_FSETID",
		"CAP_FOWNER",
		"CAP_MKNOD",
		"CAP_NET_RAW",
		"CAP_SETGID",
		"CAP_SETUID",
		"CAP_SETFCAP",
		"CAP_SETPCAP",
		"CAP_NET_BIND_SERVICE",
		"CAP_SYS_CHROOT",
		"CAP_KILL",
		"CAP_AUDIT_WRITE",
		"ALL",
	}
}

// GetMappingContainerCapabilityTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingContainerCapabilityTypeEnum(val string) (ContainerCapabilityTypeEnum, bool) {
	enum, ok := mappingContainerCapabilityTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
