// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OlvmHost Type representing a host.
type OlvmHost struct {

	// The host address (FQDN/IP).
	Address *string `mandatory:"false" json:"address"`

	// The host auto non uniform memory access (NUMA) status.
	AutoNumaStatus OlvmHostAutoNumaStatusEnum `mandatory:"false" json:"autoNumaStatus,omitempty"`

	Certificate *OlvmCertificate `mandatory:"false" json:"certificate"`

	// Free text containing comments about this object.
	Comment *string `mandatory:"false" json:"comment"`

	Cpu *OlvmCpu `mandatory:"false" json:"cpu"`

	// A human-readable description in plain text.
	Description *string `mandatory:"false" json:"description"`

	// Specifies whether host device passthrough is enabled on this host.
	IsHostDevicePassthrough *bool `mandatory:"false" json:"isHostDevicePassthrough"`

	Display *OlvmDisplay `mandatory:"false" json:"display"`

	// Status of storage domain.
	ExternalStatus OlvmHostExternalStatusEnum `mandatory:"false" json:"externalStatus,omitempty"`

	HardwareInformation *OlvmHardwareInformation `mandatory:"false" json:"hardwareInformation"`

	HostedEngine *OlvmHostedEngine `mandatory:"false" json:"hostedEngine"`

	// A unique identifier.
	Id *string `mandatory:"false" json:"id"`

	Iscsi *OlvmIscsiDetails `mandatory:"false" json:"iscsi"`

	// The host KDUMP status.
	KdumpStatus OlvmHostKdumpStatusEnum `mandatory:"false" json:"kdumpStatus,omitempty"`

	Ksm *Ksm `mandatory:"false" json:"ksm"`

	LibvirtVersion *OlvmVersion `mandatory:"false" json:"libvirtVersion"`

	// The max scheduling memory on this host in bytes.
	MaxSchedulingMemoryInBytes *int64 `mandatory:"false" json:"maxSchedulingMemoryInBytes"`

	// The amount of physical memory on this host in bytes.
	MemoryInBytes *int64 `mandatory:"false" json:"memoryInBytes"`

	// Name of the property.
	Name *string `mandatory:"false" json:"name"`

	// Specifies whether a network-related operation, such as 'setup networks', 'sync networks', or 'refresh capabilities', is currently being executed on this host.
	IsNetworkOperationInProgress *bool `mandatory:"false" json:"isNetworkOperationInProgress"`

	// Specifies whether non uniform memory access (NUMA) is supported on this host.
	IsNumaSupported *bool `mandatory:"false" json:"isNumaSupported"`

	Os *OlvmOperatingSystem `mandatory:"false" json:"os"`

	// Specifies whether we should override firewall definitions
	IsOverrideIpTables *bool `mandatory:"false" json:"isOverrideIpTables"`

	// Indicates if the host has correctly configured OVN.
	IsOvnConfigured *bool `mandatory:"false" json:"isOvnConfigured"`

	// The host port.
	Port *int `mandatory:"false" json:"port"`

	PowerManagement *OlvmPowerManagement `mandatory:"false" json:"powerManagement"`

	// Specifies whether the host should be reinstalled.
	IsReinstallationRequired *bool `mandatory:"false" json:"isReinstallationRequired"`

	SeLinux *OlvmSelinux `mandatory:"false" json:"seLinux"`

	Spm *OlvmSpm `mandatory:"false" json:"spm"`

	HostStatus *OlvmHostStatus `mandatory:"false" json:"hostStatus"`

	// The host status details.
	StatusDetail *string `mandatory:"false" json:"statusDetail"`

	Summary *OlvmVmSummary `mandatory:"false" json:"summary"`

	// Indicates if transparent huge pages (THP) support is enabled.
	IsTransparentHugePagesEnabled *bool `mandatory:"false" json:"isTransparentHugePagesEnabled"`

	Type *OlvmHostType `mandatory:"false" json:"type"`

	// Specified whether there is an oVirt-related update on this host.
	IsUpdateAvailable *bool `mandatory:"false" json:"isUpdateAvailable"`

	Version *OlvmVersion `mandatory:"false" json:"version"`

	// Specifies the vGPU placement strategy.
	VgpuPlacement OlvmHostVgpuPlacementEnum `mandatory:"false" json:"vgpuPlacement,omitempty"`
}

func (m OlvmHost) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OlvmHost) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOlvmHostAutoNumaStatusEnum(string(m.AutoNumaStatus)); !ok && m.AutoNumaStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AutoNumaStatus: %s. Supported values are: %s.", m.AutoNumaStatus, strings.Join(GetOlvmHostAutoNumaStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOlvmHostExternalStatusEnum(string(m.ExternalStatus)); !ok && m.ExternalStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExternalStatus: %s. Supported values are: %s.", m.ExternalStatus, strings.Join(GetOlvmHostExternalStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOlvmHostKdumpStatusEnum(string(m.KdumpStatus)); !ok && m.KdumpStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for KdumpStatus: %s. Supported values are: %s.", m.KdumpStatus, strings.Join(GetOlvmHostKdumpStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOlvmHostVgpuPlacementEnum(string(m.VgpuPlacement)); !ok && m.VgpuPlacement != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for VgpuPlacement: %s. Supported values are: %s.", m.VgpuPlacement, strings.Join(GetOlvmHostVgpuPlacementEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OlvmHostAutoNumaStatusEnum Enum with underlying type: string
type OlvmHostAutoNumaStatusEnum string

// Set of constants representing the allowable values for OlvmHostAutoNumaStatusEnum
const (
	OlvmHostAutoNumaStatusDisable OlvmHostAutoNumaStatusEnum = "DISABLE"
	OlvmHostAutoNumaStatusEnable  OlvmHostAutoNumaStatusEnum = "ENABLE"
	OlvmHostAutoNumaStatusUnknown OlvmHostAutoNumaStatusEnum = "UNKNOWN"
)

var mappingOlvmHostAutoNumaStatusEnum = map[string]OlvmHostAutoNumaStatusEnum{
	"DISABLE": OlvmHostAutoNumaStatusDisable,
	"ENABLE":  OlvmHostAutoNumaStatusEnable,
	"UNKNOWN": OlvmHostAutoNumaStatusUnknown,
}

var mappingOlvmHostAutoNumaStatusEnumLowerCase = map[string]OlvmHostAutoNumaStatusEnum{
	"disable": OlvmHostAutoNumaStatusDisable,
	"enable":  OlvmHostAutoNumaStatusEnable,
	"unknown": OlvmHostAutoNumaStatusUnknown,
}

// GetOlvmHostAutoNumaStatusEnumValues Enumerates the set of values for OlvmHostAutoNumaStatusEnum
func GetOlvmHostAutoNumaStatusEnumValues() []OlvmHostAutoNumaStatusEnum {
	values := make([]OlvmHostAutoNumaStatusEnum, 0)
	for _, v := range mappingOlvmHostAutoNumaStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmHostAutoNumaStatusEnumStringValues Enumerates the set of values in String for OlvmHostAutoNumaStatusEnum
func GetOlvmHostAutoNumaStatusEnumStringValues() []string {
	return []string{
		"DISABLE",
		"ENABLE",
		"UNKNOWN",
	}
}

// GetMappingOlvmHostAutoNumaStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmHostAutoNumaStatusEnum(val string) (OlvmHostAutoNumaStatusEnum, bool) {
	enum, ok := mappingOlvmHostAutoNumaStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OlvmHostExternalStatusEnum Enum with underlying type: string
type OlvmHostExternalStatusEnum string

// Set of constants representing the allowable values for OlvmHostExternalStatusEnum
const (
	OlvmHostExternalStatusError   OlvmHostExternalStatusEnum = "ERROR"
	OlvmHostExternalStatusFailure OlvmHostExternalStatusEnum = "FAILURE"
	OlvmHostExternalStatusInfo    OlvmHostExternalStatusEnum = "INFO"
	OlvmHostExternalStatusOk      OlvmHostExternalStatusEnum = "OK"
	OlvmHostExternalStatusWarning OlvmHostExternalStatusEnum = "WARNING"
)

var mappingOlvmHostExternalStatusEnum = map[string]OlvmHostExternalStatusEnum{
	"ERROR":   OlvmHostExternalStatusError,
	"FAILURE": OlvmHostExternalStatusFailure,
	"INFO":    OlvmHostExternalStatusInfo,
	"OK":      OlvmHostExternalStatusOk,
	"WARNING": OlvmHostExternalStatusWarning,
}

var mappingOlvmHostExternalStatusEnumLowerCase = map[string]OlvmHostExternalStatusEnum{
	"error":   OlvmHostExternalStatusError,
	"failure": OlvmHostExternalStatusFailure,
	"info":    OlvmHostExternalStatusInfo,
	"ok":      OlvmHostExternalStatusOk,
	"warning": OlvmHostExternalStatusWarning,
}

// GetOlvmHostExternalStatusEnumValues Enumerates the set of values for OlvmHostExternalStatusEnum
func GetOlvmHostExternalStatusEnumValues() []OlvmHostExternalStatusEnum {
	values := make([]OlvmHostExternalStatusEnum, 0)
	for _, v := range mappingOlvmHostExternalStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmHostExternalStatusEnumStringValues Enumerates the set of values in String for OlvmHostExternalStatusEnum
func GetOlvmHostExternalStatusEnumStringValues() []string {
	return []string{
		"ERROR",
		"FAILURE",
		"INFO",
		"OK",
		"WARNING",
	}
}

// GetMappingOlvmHostExternalStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmHostExternalStatusEnum(val string) (OlvmHostExternalStatusEnum, bool) {
	enum, ok := mappingOlvmHostExternalStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OlvmHostKdumpStatusEnum Enum with underlying type: string
type OlvmHostKdumpStatusEnum string

// Set of constants representing the allowable values for OlvmHostKdumpStatusEnum
const (
	OlvmHostKdumpStatusDisabled OlvmHostKdumpStatusEnum = "DISABLED"
	OlvmHostKdumpStatusEnabled  OlvmHostKdumpStatusEnum = "ENABLED"
	OlvmHostKdumpStatusUnknown  OlvmHostKdumpStatusEnum = "UNKNOWN"
)

var mappingOlvmHostKdumpStatusEnum = map[string]OlvmHostKdumpStatusEnum{
	"DISABLED": OlvmHostKdumpStatusDisabled,
	"ENABLED":  OlvmHostKdumpStatusEnabled,
	"UNKNOWN":  OlvmHostKdumpStatusUnknown,
}

var mappingOlvmHostKdumpStatusEnumLowerCase = map[string]OlvmHostKdumpStatusEnum{
	"disabled": OlvmHostKdumpStatusDisabled,
	"enabled":  OlvmHostKdumpStatusEnabled,
	"unknown":  OlvmHostKdumpStatusUnknown,
}

// GetOlvmHostKdumpStatusEnumValues Enumerates the set of values for OlvmHostKdumpStatusEnum
func GetOlvmHostKdumpStatusEnumValues() []OlvmHostKdumpStatusEnum {
	values := make([]OlvmHostKdumpStatusEnum, 0)
	for _, v := range mappingOlvmHostKdumpStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmHostKdumpStatusEnumStringValues Enumerates the set of values in String for OlvmHostKdumpStatusEnum
func GetOlvmHostKdumpStatusEnumStringValues() []string {
	return []string{
		"DISABLED",
		"ENABLED",
		"UNKNOWN",
	}
}

// GetMappingOlvmHostKdumpStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmHostKdumpStatusEnum(val string) (OlvmHostKdumpStatusEnum, bool) {
	enum, ok := mappingOlvmHostKdumpStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OlvmHostVgpuPlacementEnum Enum with underlying type: string
type OlvmHostVgpuPlacementEnum string

// Set of constants representing the allowable values for OlvmHostVgpuPlacementEnum
const (
	OlvmHostVgpuPlacementConsolidated OlvmHostVgpuPlacementEnum = "CONSOLIDATED"
	OlvmHostVgpuPlacementSeparated    OlvmHostVgpuPlacementEnum = "SEPARATED"
)

var mappingOlvmHostVgpuPlacementEnum = map[string]OlvmHostVgpuPlacementEnum{
	"CONSOLIDATED": OlvmHostVgpuPlacementConsolidated,
	"SEPARATED":    OlvmHostVgpuPlacementSeparated,
}

var mappingOlvmHostVgpuPlacementEnumLowerCase = map[string]OlvmHostVgpuPlacementEnum{
	"consolidated": OlvmHostVgpuPlacementConsolidated,
	"separated":    OlvmHostVgpuPlacementSeparated,
}

// GetOlvmHostVgpuPlacementEnumValues Enumerates the set of values for OlvmHostVgpuPlacementEnum
func GetOlvmHostVgpuPlacementEnumValues() []OlvmHostVgpuPlacementEnum {
	values := make([]OlvmHostVgpuPlacementEnum, 0)
	for _, v := range mappingOlvmHostVgpuPlacementEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmHostVgpuPlacementEnumStringValues Enumerates the set of values in String for OlvmHostVgpuPlacementEnum
func GetOlvmHostVgpuPlacementEnumStringValues() []string {
	return []string{
		"CONSOLIDATED",
		"SEPARATED",
	}
}

// GetMappingOlvmHostVgpuPlacementEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmHostVgpuPlacementEnum(val string) (OlvmHostVgpuPlacementEnum, bool) {
	enum, ok := mappingOlvmHostVgpuPlacementEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
