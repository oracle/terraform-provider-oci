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

// OlvmVirtualMachine Represents a virtual machine.
type OlvmVirtualMachine struct {

	// Specifies if and how the auto CPU and NUMA configuration is applied.
	AutoPinningPolicy OlvmVirtualMachineAutoPinningPolicyEnum `mandatory:"false" json:"autoPinningPolicy,omitempty"`

	Bios *OlvmBios `mandatory:"false" json:"bios"`

	// Free text containing comments about this object.
	Comment *string `mandatory:"false" json:"comment"`

	Console *OlvmConsole `mandatory:"false" json:"console"`

	Cpu *OlvmCpu `mandatory:"false" json:"cpu"`

	// Type representing the CPU and NUMA pinning policy.
	CpuPinningPolicy OlvmVirtualMachineCpuPinningPolicyEnum `mandatory:"false" json:"cpuPinningPolicy,omitempty"`

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

	// Fully qualified domain name of the virtual machine.
	Fqdn *string `mandatory:"false" json:"fqdn"`

	GuestOperatingSystem *OlvmGuestOperatingSystem `mandatory:"false" json:"guestOperatingSystem"`

	// What time zone is used by the virtual machine (as returned by guest agent).
	GuestTimeZone *string `mandatory:"false" json:"guestTimeZone"`

	// Indicates whether the virtual machine has snapshots with disks in ILLEGAL state.
	IsContainIllegalImages *bool `mandatory:"false" json:"isContainIllegalImages"`

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

	// Name of the property.
	VirtualMachineName *string `mandatory:"false" json:"virtualMachineName"`

	// Indicates if virtual machine configuration has been changed and requires restart of the virtual machine.
	IsNextRunConfigurationExists *bool `mandatory:"false" json:"isNextRunConfigurationExists"`

	// How the NUMA topology is applied.
	NumaTuneMode OlvmVirtualMachineNumaTuneModeEnum `mandatory:"false" json:"numaTuneMode,omitempty"`

	// The origin of this virtual machine.
	Origin *string `mandatory:"false" json:"origin"`

	Os *OlvmOperatingSystem `mandatory:"false" json:"os"`

	PlacementPolicy *OlvmVmPlacementPolicy `mandatory:"false" json:"placementPolicy"`

	// If true, the virtual machine has been started using the run once command, meaning it???s configuration might differ from the stored one for the purpose of this single run.
	IsRunOnce *bool `mandatory:"false" json:"isRunOnce"`

	SerialNumber *OlvmSerialNumber `mandatory:"false" json:"serialNumber"`

	SmallIcon *OlvmIcon `mandatory:"false" json:"smallIcon"`

	// If true, the sound card is added to the virtual machine.
	IsSoundcardEnabled *bool `mandatory:"false" json:"isSoundcardEnabled"`

	// If true, the virtual machine will be initially in 'paused' state after start.
	IsStartPaused *bool `mandatory:"false" json:"isStartPaused"`

	// The date in which the virtual machine was started.
	TimeStart *common.SDKTime `mandatory:"false" json:"timeStart"`

	// If true, the virtual machine is stateless - it???s state (disks) are rolled-back after shutdown.
	IsStateless *bool `mandatory:"false" json:"isStateless"`

	Status *OlvmVmStatus `mandatory:"false" json:"status"`

	// Human readable detail of current status.
	StatusDetail *string `mandatory:"false" json:"statusDetail"`

	// The reason the virtual machine was stopped.
	StopReason *string `mandatory:"false" json:"stopReason"`

	// The date in which the virtual machine was stopped.
	TimeStop *common.SDKTime `mandatory:"false" json:"timeStop"`

	// Determines how the virtual machine will be resumed after storage error.
	StorageErrorResumeBehavior OlvmVirtualMachineStorageErrorResumeBehaviorEnum `mandatory:"false" json:"storageErrorResumeBehavior,omitempty"`

	// Name of time zone.
	Timezone *string `mandatory:"false" json:"timezone"`

	// If true, a TPM device is added to the virtual machine.
	IsTpmEnabled *bool `mandatory:"false" json:"isTpmEnabled"`

	// If true, the network data transfer will be encrypted during virtual machine live migration.
	IsTunnelMigration *bool `mandatory:"false" json:"isTunnelMigration"`

	Type *OlvmVmType `mandatory:"false" json:"type"`

	Usb *OlvmUsb `mandatory:"false" json:"usb"`

	// If true, the virtual machine is reconfigured to the latest version of it???s template when it is started.
	IsUseLatestTemplateVersion *bool `mandatory:"false" json:"isUseLatestTemplateVersion"`

	// Indicates if Virtio SCSI Support is enabled.
	IsVirtioScsiEnabled *bool `mandatory:"false" json:"isVirtioScsiEnabled"`

	// Number of queues for a Virtio-SCSI contoller this field requires virtioScsiMultiQueuesEnabled to be true see virtioScsiMultiQueuesEnabled for more info
	VirtioScsiMultiQueues *int `mandatory:"false" json:"virtioScsiMultiQueues"`

	// If true, the Virtio-SCSI devices will obtain a number of multiple queues depending on the available virtual Cpus and disks, or according to the specified virtioScsiMultiQueues
	IsVirtioScsiMultiQueuesEnabled *bool `mandatory:"false" json:"isVirtioScsiMultiQueuesEnabled"`
}

func (m OlvmVirtualMachine) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OlvmVirtualMachine) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOlvmVirtualMachineAutoPinningPolicyEnum(string(m.AutoPinningPolicy)); !ok && m.AutoPinningPolicy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AutoPinningPolicy: %s. Supported values are: %s.", m.AutoPinningPolicy, strings.Join(GetOlvmVirtualMachineAutoPinningPolicyEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOlvmVirtualMachineCpuPinningPolicyEnum(string(m.CpuPinningPolicy)); !ok && m.CpuPinningPolicy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CpuPinningPolicy: %s. Supported values are: %s.", m.CpuPinningPolicy, strings.Join(GetOlvmVirtualMachineCpuPinningPolicyEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOlvmVirtualMachineNumaTuneModeEnum(string(m.NumaTuneMode)); !ok && m.NumaTuneMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NumaTuneMode: %s. Supported values are: %s.", m.NumaTuneMode, strings.Join(GetOlvmVirtualMachineNumaTuneModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOlvmVirtualMachineStorageErrorResumeBehaviorEnum(string(m.StorageErrorResumeBehavior)); !ok && m.StorageErrorResumeBehavior != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for StorageErrorResumeBehavior: %s. Supported values are: %s.", m.StorageErrorResumeBehavior, strings.Join(GetOlvmVirtualMachineStorageErrorResumeBehaviorEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OlvmVirtualMachineAutoPinningPolicyEnum Enum with underlying type: string
type OlvmVirtualMachineAutoPinningPolicyEnum string

// Set of constants representing the allowable values for OlvmVirtualMachineAutoPinningPolicyEnum
const (
	OlvmVirtualMachineAutoPinningPolicyAdjust   OlvmVirtualMachineAutoPinningPolicyEnum = "ADJUST"
	OlvmVirtualMachineAutoPinningPolicyDisabled OlvmVirtualMachineAutoPinningPolicyEnum = "DISABLED"
	OlvmVirtualMachineAutoPinningPolicyExisting OlvmVirtualMachineAutoPinningPolicyEnum = "EXISTING"
)

var mappingOlvmVirtualMachineAutoPinningPolicyEnum = map[string]OlvmVirtualMachineAutoPinningPolicyEnum{
	"ADJUST":   OlvmVirtualMachineAutoPinningPolicyAdjust,
	"DISABLED": OlvmVirtualMachineAutoPinningPolicyDisabled,
	"EXISTING": OlvmVirtualMachineAutoPinningPolicyExisting,
}

var mappingOlvmVirtualMachineAutoPinningPolicyEnumLowerCase = map[string]OlvmVirtualMachineAutoPinningPolicyEnum{
	"adjust":   OlvmVirtualMachineAutoPinningPolicyAdjust,
	"disabled": OlvmVirtualMachineAutoPinningPolicyDisabled,
	"existing": OlvmVirtualMachineAutoPinningPolicyExisting,
}

// GetOlvmVirtualMachineAutoPinningPolicyEnumValues Enumerates the set of values for OlvmVirtualMachineAutoPinningPolicyEnum
func GetOlvmVirtualMachineAutoPinningPolicyEnumValues() []OlvmVirtualMachineAutoPinningPolicyEnum {
	values := make([]OlvmVirtualMachineAutoPinningPolicyEnum, 0)
	for _, v := range mappingOlvmVirtualMachineAutoPinningPolicyEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmVirtualMachineAutoPinningPolicyEnumStringValues Enumerates the set of values in String for OlvmVirtualMachineAutoPinningPolicyEnum
func GetOlvmVirtualMachineAutoPinningPolicyEnumStringValues() []string {
	return []string{
		"ADJUST",
		"DISABLED",
		"EXISTING",
	}
}

// GetMappingOlvmVirtualMachineAutoPinningPolicyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmVirtualMachineAutoPinningPolicyEnum(val string) (OlvmVirtualMachineAutoPinningPolicyEnum, bool) {
	enum, ok := mappingOlvmVirtualMachineAutoPinningPolicyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OlvmVirtualMachineCpuPinningPolicyEnum Enum with underlying type: string
type OlvmVirtualMachineCpuPinningPolicyEnum string

// Set of constants representing the allowable values for OlvmVirtualMachineCpuPinningPolicyEnum
const (
	OlvmVirtualMachineCpuPinningPolicyDedicated        OlvmVirtualMachineCpuPinningPolicyEnum = "DEDICATED"
	OlvmVirtualMachineCpuPinningPolicyIsolatedThreads  OlvmVirtualMachineCpuPinningPolicyEnum = "ISOLATED_THREADS"
	OlvmVirtualMachineCpuPinningPolicyManual           OlvmVirtualMachineCpuPinningPolicyEnum = "MANUAL"
	OlvmVirtualMachineCpuPinningPolicyNone             OlvmVirtualMachineCpuPinningPolicyEnum = "NONE"
	OlvmVirtualMachineCpuPinningPolicyResizeAndPinNuma OlvmVirtualMachineCpuPinningPolicyEnum = "RESIZE_AND_PIN_NUMA"
)

var mappingOlvmVirtualMachineCpuPinningPolicyEnum = map[string]OlvmVirtualMachineCpuPinningPolicyEnum{
	"DEDICATED":           OlvmVirtualMachineCpuPinningPolicyDedicated,
	"ISOLATED_THREADS":    OlvmVirtualMachineCpuPinningPolicyIsolatedThreads,
	"MANUAL":              OlvmVirtualMachineCpuPinningPolicyManual,
	"NONE":                OlvmVirtualMachineCpuPinningPolicyNone,
	"RESIZE_AND_PIN_NUMA": OlvmVirtualMachineCpuPinningPolicyResizeAndPinNuma,
}

var mappingOlvmVirtualMachineCpuPinningPolicyEnumLowerCase = map[string]OlvmVirtualMachineCpuPinningPolicyEnum{
	"dedicated":           OlvmVirtualMachineCpuPinningPolicyDedicated,
	"isolated_threads":    OlvmVirtualMachineCpuPinningPolicyIsolatedThreads,
	"manual":              OlvmVirtualMachineCpuPinningPolicyManual,
	"none":                OlvmVirtualMachineCpuPinningPolicyNone,
	"resize_and_pin_numa": OlvmVirtualMachineCpuPinningPolicyResizeAndPinNuma,
}

// GetOlvmVirtualMachineCpuPinningPolicyEnumValues Enumerates the set of values for OlvmVirtualMachineCpuPinningPolicyEnum
func GetOlvmVirtualMachineCpuPinningPolicyEnumValues() []OlvmVirtualMachineCpuPinningPolicyEnum {
	values := make([]OlvmVirtualMachineCpuPinningPolicyEnum, 0)
	for _, v := range mappingOlvmVirtualMachineCpuPinningPolicyEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmVirtualMachineCpuPinningPolicyEnumStringValues Enumerates the set of values in String for OlvmVirtualMachineCpuPinningPolicyEnum
func GetOlvmVirtualMachineCpuPinningPolicyEnumStringValues() []string {
	return []string{
		"DEDICATED",
		"ISOLATED_THREADS",
		"MANUAL",
		"NONE",
		"RESIZE_AND_PIN_NUMA",
	}
}

// GetMappingOlvmVirtualMachineCpuPinningPolicyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmVirtualMachineCpuPinningPolicyEnum(val string) (OlvmVirtualMachineCpuPinningPolicyEnum, bool) {
	enum, ok := mappingOlvmVirtualMachineCpuPinningPolicyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OlvmVirtualMachineNumaTuneModeEnum Enum with underlying type: string
type OlvmVirtualMachineNumaTuneModeEnum string

// Set of constants representing the allowable values for OlvmVirtualMachineNumaTuneModeEnum
const (
	OlvmVirtualMachineNumaTuneModeInterleave OlvmVirtualMachineNumaTuneModeEnum = "INTERLEAVE"
	OlvmVirtualMachineNumaTuneModePreferred  OlvmVirtualMachineNumaTuneModeEnum = "PREFERRED"
	OlvmVirtualMachineNumaTuneModeStrict     OlvmVirtualMachineNumaTuneModeEnum = "STRICT"
)

var mappingOlvmVirtualMachineNumaTuneModeEnum = map[string]OlvmVirtualMachineNumaTuneModeEnum{
	"INTERLEAVE": OlvmVirtualMachineNumaTuneModeInterleave,
	"PREFERRED":  OlvmVirtualMachineNumaTuneModePreferred,
	"STRICT":     OlvmVirtualMachineNumaTuneModeStrict,
}

var mappingOlvmVirtualMachineNumaTuneModeEnumLowerCase = map[string]OlvmVirtualMachineNumaTuneModeEnum{
	"interleave": OlvmVirtualMachineNumaTuneModeInterleave,
	"preferred":  OlvmVirtualMachineNumaTuneModePreferred,
	"strict":     OlvmVirtualMachineNumaTuneModeStrict,
}

// GetOlvmVirtualMachineNumaTuneModeEnumValues Enumerates the set of values for OlvmVirtualMachineNumaTuneModeEnum
func GetOlvmVirtualMachineNumaTuneModeEnumValues() []OlvmVirtualMachineNumaTuneModeEnum {
	values := make([]OlvmVirtualMachineNumaTuneModeEnum, 0)
	for _, v := range mappingOlvmVirtualMachineNumaTuneModeEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmVirtualMachineNumaTuneModeEnumStringValues Enumerates the set of values in String for OlvmVirtualMachineNumaTuneModeEnum
func GetOlvmVirtualMachineNumaTuneModeEnumStringValues() []string {
	return []string{
		"INTERLEAVE",
		"PREFERRED",
		"STRICT",
	}
}

// GetMappingOlvmVirtualMachineNumaTuneModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmVirtualMachineNumaTuneModeEnum(val string) (OlvmVirtualMachineNumaTuneModeEnum, bool) {
	enum, ok := mappingOlvmVirtualMachineNumaTuneModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OlvmVirtualMachineStorageErrorResumeBehaviorEnum Enum with underlying type: string
type OlvmVirtualMachineStorageErrorResumeBehaviorEnum string

// Set of constants representing the allowable values for OlvmVirtualMachineStorageErrorResumeBehaviorEnum
const (
	OlvmVirtualMachineStorageErrorResumeBehaviorAutoResume  OlvmVirtualMachineStorageErrorResumeBehaviorEnum = "AUTO_RESUME"
	OlvmVirtualMachineStorageErrorResumeBehaviorKill        OlvmVirtualMachineStorageErrorResumeBehaviorEnum = "KILL"
	OlvmVirtualMachineStorageErrorResumeBehaviorLeavePaused OlvmVirtualMachineStorageErrorResumeBehaviorEnum = "LEAVE_PAUSED"
)

var mappingOlvmVirtualMachineStorageErrorResumeBehaviorEnum = map[string]OlvmVirtualMachineStorageErrorResumeBehaviorEnum{
	"AUTO_RESUME":  OlvmVirtualMachineStorageErrorResumeBehaviorAutoResume,
	"KILL":         OlvmVirtualMachineStorageErrorResumeBehaviorKill,
	"LEAVE_PAUSED": OlvmVirtualMachineStorageErrorResumeBehaviorLeavePaused,
}

var mappingOlvmVirtualMachineStorageErrorResumeBehaviorEnumLowerCase = map[string]OlvmVirtualMachineStorageErrorResumeBehaviorEnum{
	"auto_resume":  OlvmVirtualMachineStorageErrorResumeBehaviorAutoResume,
	"kill":         OlvmVirtualMachineStorageErrorResumeBehaviorKill,
	"leave_paused": OlvmVirtualMachineStorageErrorResumeBehaviorLeavePaused,
}

// GetOlvmVirtualMachineStorageErrorResumeBehaviorEnumValues Enumerates the set of values for OlvmVirtualMachineStorageErrorResumeBehaviorEnum
func GetOlvmVirtualMachineStorageErrorResumeBehaviorEnumValues() []OlvmVirtualMachineStorageErrorResumeBehaviorEnum {
	values := make([]OlvmVirtualMachineStorageErrorResumeBehaviorEnum, 0)
	for _, v := range mappingOlvmVirtualMachineStorageErrorResumeBehaviorEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmVirtualMachineStorageErrorResumeBehaviorEnumStringValues Enumerates the set of values in String for OlvmVirtualMachineStorageErrorResumeBehaviorEnum
func GetOlvmVirtualMachineStorageErrorResumeBehaviorEnumStringValues() []string {
	return []string{
		"AUTO_RESUME",
		"KILL",
		"LEAVE_PAUSED",
	}
}

// GetMappingOlvmVirtualMachineStorageErrorResumeBehaviorEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmVirtualMachineStorageErrorResumeBehaviorEnum(val string) (OlvmVirtualMachineStorageErrorResumeBehaviorEnum, bool) {
	enum, ok := mappingOlvmVirtualMachineStorageErrorResumeBehaviorEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
