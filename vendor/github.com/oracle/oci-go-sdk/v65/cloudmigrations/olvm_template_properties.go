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

// OlvmTemplateProperties OLVM Template properties.
type OlvmTemplateProperties struct {

	// Name of the property.
	TemplateName *string `mandatory:"true" json:"templateName"`

	// Specifies if and how the auto CPU and NUMA configuration is applied.
	AutoPinningPolicy OlvmTemplatePropertiesAutoPinningPolicyEnum `mandatory:"false" json:"autoPinningPolicy,omitempty"`

	Bios *OlvmBios `mandatory:"false" json:"bios"`

	// Free text containing comments about this object.
	Comment *string `mandatory:"false" json:"comment"`

	Console *OlvmConsole `mandatory:"false" json:"console"`

	Cpu *OlvmCpu `mandatory:"false" json:"cpu"`

	// Type representing the CPU and NUMA pinning policy.
	CpuPinningPolicy OlvmTemplatePropertiesCpuPinningPolicyEnum `mandatory:"false" json:"cpuPinningPolicy,omitempty"`

	// Number of CPU Shares
	CpuShares *int `mandatory:"false" json:"cpuShares"`

	// Creation time.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	CustomCompatibilityVersion *OlvmVersion `mandatory:"false" json:"customCompatibilityVersion"`

	// Custom CPU model
	CustomCpuModel *string `mandatory:"false" json:"customCpuModel"`

	// Custom Emulated machine
	CustomEmulatedMachine *string `mandatory:"false" json:"customEmulatedMachine"`

	// Custom properties applied to the vNIC profile.
	CustomProperties []OlvmCustomProperty `mandatory:"false" json:"customProperties"`

	// If true, the virtual machine cannot be deleted.
	IsDeleteProtected *bool `mandatory:"false" json:"isDeleteProtected"`

	// A human-readable description in plain text.
	Description *string `mandatory:"false" json:"description"`

	Display *OlvmDisplay `mandatory:"false" json:"display"`

	Domain *OlvmDomain `mandatory:"false" json:"domain"`

	HighAvailability *OlvmHighAvailability `mandatory:"false" json:"highAvailability"`

	Initialization *OlvmInitialization `mandatory:"false" json:"initialization"`

	// For performance tuning of IO threading.
	IoThreads *int `mandatory:"false" json:"ioThreads"`

	LargeIcon *OlvmIcon `mandatory:"false" json:"largeIcon"`

	StorageDomainLease *OlvmStorageDomainProperties `mandatory:"false" json:"storageDomainLease"`

	// The virtual machine???s memory, in bytes.
	MemoryInBytes *int64 `mandatory:"false" json:"memoryInBytes"`

	MemoryPolicy *OlvmMemoryPolicy `mandatory:"false" json:"memoryPolicy"`

	MigrationPolicy *OlvmMigrationOptions `mandatory:"false" json:"migrationPolicy"`

	// Maximum time the virtual machine can be non responsive during its live migration to another host in ms.
	MigrationDowntimeInMs *int `mandatory:"false" json:"migrationDowntimeInMs"`

	// If true, each virtual interface will get the optimal number of queues, depending on the available virtual Cpus.
	IsMultiQueuesEnabled *bool `mandatory:"false" json:"isMultiQueuesEnabled"`

	// The origin of this virtual machine.
	Origin *string `mandatory:"false" json:"origin"`

	Os *OlvmOperatingSystem `mandatory:"false" json:"os"`

	PlacementPolicy *OlvmVmPlacementPolicy `mandatory:"false" json:"placementPolicy"`

	SerialNumber *OlvmSerialNumber `mandatory:"false" json:"serialNumber"`

	SmallIcon *OlvmIcon `mandatory:"false" json:"smallIcon"`

	// If true, the sound card is added to the virtual machine.
	IsSoundcardEnabled *bool `mandatory:"false" json:"isSoundcardEnabled"`

	// If true, the virtual machine will be initially in 'paused' state after start.
	IsStartPaused *bool `mandatory:"false" json:"isStartPaused"`

	// If true, the virtual machine is stateless - it???s state (disks) are rolled-back after shutdown.
	IsStateless *bool `mandatory:"false" json:"isStateless"`

	// The status of the template.
	Status OlvmTemplatePropertiesStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Determines how the virtual machine will be resumed after storage error.
	StorageErrorResumeBehavior OlvmTemplatePropertiesStorageErrorResumeBehaviorEnum `mandatory:"false" json:"storageErrorResumeBehavior,omitempty"`

	// Name of time zone.
	Timezone *string `mandatory:"false" json:"timezone"`

	// If true, a TPM device is added to the virtual machine.
	IsTpmEnabled *bool `mandatory:"false" json:"isTpmEnabled"`

	// If true, the network data transfer will be encrypted during virtual machine live migration.
	IsTunnelMigration *bool `mandatory:"false" json:"isTunnelMigration"`

	Type *OlvmVmType `mandatory:"false" json:"type"`

	Usb *OlvmUsb `mandatory:"false" json:"usb"`

	Version *OlvmTemplateVersion `mandatory:"false" json:"version"`

	// Indicates if Virtio SCSI Support is enabled.
	IsVirtioScsiEnabled *bool `mandatory:"false" json:"isVirtioScsiEnabled"`

	// Number of queues for a Virtio-SCSI contoller this field requires virtioScsiMultiQueuesEnabled to be true see virtioScsiMultiQueuesEnabled for more info
	VirtioScsiMultiQueues *int `mandatory:"false" json:"virtioScsiMultiQueues"`

	// If true, the Virtio-SCSI devices will obtain a number of multiple queues depending on the available virtual Cpus and disks, or according to the specified virtioScsiMultiQueues
	IsVirtioScsiMultiQueuesEnabled *bool `mandatory:"false" json:"isVirtioScsiMultiQueuesEnabled"`

	Vm *OlvmVirtualMachine `mandatory:"false" json:"vm"`
}

func (m OlvmTemplateProperties) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OlvmTemplateProperties) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOlvmTemplatePropertiesAutoPinningPolicyEnum(string(m.AutoPinningPolicy)); !ok && m.AutoPinningPolicy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AutoPinningPolicy: %s. Supported values are: %s.", m.AutoPinningPolicy, strings.Join(GetOlvmTemplatePropertiesAutoPinningPolicyEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOlvmTemplatePropertiesCpuPinningPolicyEnum(string(m.CpuPinningPolicy)); !ok && m.CpuPinningPolicy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CpuPinningPolicy: %s. Supported values are: %s.", m.CpuPinningPolicy, strings.Join(GetOlvmTemplatePropertiesCpuPinningPolicyEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOlvmTemplatePropertiesStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetOlvmTemplatePropertiesStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOlvmTemplatePropertiesStorageErrorResumeBehaviorEnum(string(m.StorageErrorResumeBehavior)); !ok && m.StorageErrorResumeBehavior != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for StorageErrorResumeBehavior: %s. Supported values are: %s.", m.StorageErrorResumeBehavior, strings.Join(GetOlvmTemplatePropertiesStorageErrorResumeBehaviorEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OlvmTemplatePropertiesAutoPinningPolicyEnum Enum with underlying type: string
type OlvmTemplatePropertiesAutoPinningPolicyEnum string

// Set of constants representing the allowable values for OlvmTemplatePropertiesAutoPinningPolicyEnum
const (
	OlvmTemplatePropertiesAutoPinningPolicyAdjust   OlvmTemplatePropertiesAutoPinningPolicyEnum = "ADJUST"
	OlvmTemplatePropertiesAutoPinningPolicyDisabled OlvmTemplatePropertiesAutoPinningPolicyEnum = "DISABLED"
	OlvmTemplatePropertiesAutoPinningPolicyExisting OlvmTemplatePropertiesAutoPinningPolicyEnum = "EXISTING"
)

var mappingOlvmTemplatePropertiesAutoPinningPolicyEnum = map[string]OlvmTemplatePropertiesAutoPinningPolicyEnum{
	"ADJUST":   OlvmTemplatePropertiesAutoPinningPolicyAdjust,
	"DISABLED": OlvmTemplatePropertiesAutoPinningPolicyDisabled,
	"EXISTING": OlvmTemplatePropertiesAutoPinningPolicyExisting,
}

var mappingOlvmTemplatePropertiesAutoPinningPolicyEnumLowerCase = map[string]OlvmTemplatePropertiesAutoPinningPolicyEnum{
	"adjust":   OlvmTemplatePropertiesAutoPinningPolicyAdjust,
	"disabled": OlvmTemplatePropertiesAutoPinningPolicyDisabled,
	"existing": OlvmTemplatePropertiesAutoPinningPolicyExisting,
}

// GetOlvmTemplatePropertiesAutoPinningPolicyEnumValues Enumerates the set of values for OlvmTemplatePropertiesAutoPinningPolicyEnum
func GetOlvmTemplatePropertiesAutoPinningPolicyEnumValues() []OlvmTemplatePropertiesAutoPinningPolicyEnum {
	values := make([]OlvmTemplatePropertiesAutoPinningPolicyEnum, 0)
	for _, v := range mappingOlvmTemplatePropertiesAutoPinningPolicyEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmTemplatePropertiesAutoPinningPolicyEnumStringValues Enumerates the set of values in String for OlvmTemplatePropertiesAutoPinningPolicyEnum
func GetOlvmTemplatePropertiesAutoPinningPolicyEnumStringValues() []string {
	return []string{
		"ADJUST",
		"DISABLED",
		"EXISTING",
	}
}

// GetMappingOlvmTemplatePropertiesAutoPinningPolicyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmTemplatePropertiesAutoPinningPolicyEnum(val string) (OlvmTemplatePropertiesAutoPinningPolicyEnum, bool) {
	enum, ok := mappingOlvmTemplatePropertiesAutoPinningPolicyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OlvmTemplatePropertiesCpuPinningPolicyEnum Enum with underlying type: string
type OlvmTemplatePropertiesCpuPinningPolicyEnum string

// Set of constants representing the allowable values for OlvmTemplatePropertiesCpuPinningPolicyEnum
const (
	OlvmTemplatePropertiesCpuPinningPolicyDedicated        OlvmTemplatePropertiesCpuPinningPolicyEnum = "DEDICATED"
	OlvmTemplatePropertiesCpuPinningPolicyIsolatedThreads  OlvmTemplatePropertiesCpuPinningPolicyEnum = "ISOLATED_THREADS"
	OlvmTemplatePropertiesCpuPinningPolicyManual           OlvmTemplatePropertiesCpuPinningPolicyEnum = "MANUAL"
	OlvmTemplatePropertiesCpuPinningPolicyNone             OlvmTemplatePropertiesCpuPinningPolicyEnum = "NONE"
	OlvmTemplatePropertiesCpuPinningPolicyResizeAndPinNuma OlvmTemplatePropertiesCpuPinningPolicyEnum = "RESIZE_AND_PIN_NUMA"
)

var mappingOlvmTemplatePropertiesCpuPinningPolicyEnum = map[string]OlvmTemplatePropertiesCpuPinningPolicyEnum{
	"DEDICATED":           OlvmTemplatePropertiesCpuPinningPolicyDedicated,
	"ISOLATED_THREADS":    OlvmTemplatePropertiesCpuPinningPolicyIsolatedThreads,
	"MANUAL":              OlvmTemplatePropertiesCpuPinningPolicyManual,
	"NONE":                OlvmTemplatePropertiesCpuPinningPolicyNone,
	"RESIZE_AND_PIN_NUMA": OlvmTemplatePropertiesCpuPinningPolicyResizeAndPinNuma,
}

var mappingOlvmTemplatePropertiesCpuPinningPolicyEnumLowerCase = map[string]OlvmTemplatePropertiesCpuPinningPolicyEnum{
	"dedicated":           OlvmTemplatePropertiesCpuPinningPolicyDedicated,
	"isolated_threads":    OlvmTemplatePropertiesCpuPinningPolicyIsolatedThreads,
	"manual":              OlvmTemplatePropertiesCpuPinningPolicyManual,
	"none":                OlvmTemplatePropertiesCpuPinningPolicyNone,
	"resize_and_pin_numa": OlvmTemplatePropertiesCpuPinningPolicyResizeAndPinNuma,
}

// GetOlvmTemplatePropertiesCpuPinningPolicyEnumValues Enumerates the set of values for OlvmTemplatePropertiesCpuPinningPolicyEnum
func GetOlvmTemplatePropertiesCpuPinningPolicyEnumValues() []OlvmTemplatePropertiesCpuPinningPolicyEnum {
	values := make([]OlvmTemplatePropertiesCpuPinningPolicyEnum, 0)
	for _, v := range mappingOlvmTemplatePropertiesCpuPinningPolicyEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmTemplatePropertiesCpuPinningPolicyEnumStringValues Enumerates the set of values in String for OlvmTemplatePropertiesCpuPinningPolicyEnum
func GetOlvmTemplatePropertiesCpuPinningPolicyEnumStringValues() []string {
	return []string{
		"DEDICATED",
		"ISOLATED_THREADS",
		"MANUAL",
		"NONE",
		"RESIZE_AND_PIN_NUMA",
	}
}

// GetMappingOlvmTemplatePropertiesCpuPinningPolicyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmTemplatePropertiesCpuPinningPolicyEnum(val string) (OlvmTemplatePropertiesCpuPinningPolicyEnum, bool) {
	enum, ok := mappingOlvmTemplatePropertiesCpuPinningPolicyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OlvmTemplatePropertiesStatusEnum Enum with underlying type: string
type OlvmTemplatePropertiesStatusEnum string

// Set of constants representing the allowable values for OlvmTemplatePropertiesStatusEnum
const (
	OlvmTemplatePropertiesStatusIllegal OlvmTemplatePropertiesStatusEnum = "ILLEGAL"
	OlvmTemplatePropertiesStatusLocked  OlvmTemplatePropertiesStatusEnum = "LOCKED"
	OlvmTemplatePropertiesStatusOk      OlvmTemplatePropertiesStatusEnum = "OK"
)

var mappingOlvmTemplatePropertiesStatusEnum = map[string]OlvmTemplatePropertiesStatusEnum{
	"ILLEGAL": OlvmTemplatePropertiesStatusIllegal,
	"LOCKED":  OlvmTemplatePropertiesStatusLocked,
	"OK":      OlvmTemplatePropertiesStatusOk,
}

var mappingOlvmTemplatePropertiesStatusEnumLowerCase = map[string]OlvmTemplatePropertiesStatusEnum{
	"illegal": OlvmTemplatePropertiesStatusIllegal,
	"locked":  OlvmTemplatePropertiesStatusLocked,
	"ok":      OlvmTemplatePropertiesStatusOk,
}

// GetOlvmTemplatePropertiesStatusEnumValues Enumerates the set of values for OlvmTemplatePropertiesStatusEnum
func GetOlvmTemplatePropertiesStatusEnumValues() []OlvmTemplatePropertiesStatusEnum {
	values := make([]OlvmTemplatePropertiesStatusEnum, 0)
	for _, v := range mappingOlvmTemplatePropertiesStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmTemplatePropertiesStatusEnumStringValues Enumerates the set of values in String for OlvmTemplatePropertiesStatusEnum
func GetOlvmTemplatePropertiesStatusEnumStringValues() []string {
	return []string{
		"ILLEGAL",
		"LOCKED",
		"OK",
	}
}

// GetMappingOlvmTemplatePropertiesStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmTemplatePropertiesStatusEnum(val string) (OlvmTemplatePropertiesStatusEnum, bool) {
	enum, ok := mappingOlvmTemplatePropertiesStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OlvmTemplatePropertiesStorageErrorResumeBehaviorEnum Enum with underlying type: string
type OlvmTemplatePropertiesStorageErrorResumeBehaviorEnum string

// Set of constants representing the allowable values for OlvmTemplatePropertiesStorageErrorResumeBehaviorEnum
const (
	OlvmTemplatePropertiesStorageErrorResumeBehaviorAutoResume  OlvmTemplatePropertiesStorageErrorResumeBehaviorEnum = "AUTO_RESUME"
	OlvmTemplatePropertiesStorageErrorResumeBehaviorKill        OlvmTemplatePropertiesStorageErrorResumeBehaviorEnum = "KILL"
	OlvmTemplatePropertiesStorageErrorResumeBehaviorLeavePaused OlvmTemplatePropertiesStorageErrorResumeBehaviorEnum = "LEAVE_PAUSED"
)

var mappingOlvmTemplatePropertiesStorageErrorResumeBehaviorEnum = map[string]OlvmTemplatePropertiesStorageErrorResumeBehaviorEnum{
	"AUTO_RESUME":  OlvmTemplatePropertiesStorageErrorResumeBehaviorAutoResume,
	"KILL":         OlvmTemplatePropertiesStorageErrorResumeBehaviorKill,
	"LEAVE_PAUSED": OlvmTemplatePropertiesStorageErrorResumeBehaviorLeavePaused,
}

var mappingOlvmTemplatePropertiesStorageErrorResumeBehaviorEnumLowerCase = map[string]OlvmTemplatePropertiesStorageErrorResumeBehaviorEnum{
	"auto_resume":  OlvmTemplatePropertiesStorageErrorResumeBehaviorAutoResume,
	"kill":         OlvmTemplatePropertiesStorageErrorResumeBehaviorKill,
	"leave_paused": OlvmTemplatePropertiesStorageErrorResumeBehaviorLeavePaused,
}

// GetOlvmTemplatePropertiesStorageErrorResumeBehaviorEnumValues Enumerates the set of values for OlvmTemplatePropertiesStorageErrorResumeBehaviorEnum
func GetOlvmTemplatePropertiesStorageErrorResumeBehaviorEnumValues() []OlvmTemplatePropertiesStorageErrorResumeBehaviorEnum {
	values := make([]OlvmTemplatePropertiesStorageErrorResumeBehaviorEnum, 0)
	for _, v := range mappingOlvmTemplatePropertiesStorageErrorResumeBehaviorEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmTemplatePropertiesStorageErrorResumeBehaviorEnumStringValues Enumerates the set of values in String for OlvmTemplatePropertiesStorageErrorResumeBehaviorEnum
func GetOlvmTemplatePropertiesStorageErrorResumeBehaviorEnumStringValues() []string {
	return []string{
		"AUTO_RESUME",
		"KILL",
		"LEAVE_PAUSED",
	}
}

// GetMappingOlvmTemplatePropertiesStorageErrorResumeBehaviorEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmTemplatePropertiesStorageErrorResumeBehaviorEnum(val string) (OlvmTemplatePropertiesStorageErrorResumeBehaviorEnum, bool) {
	enum, ok := mappingOlvmTemplatePropertiesStorageErrorResumeBehaviorEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
