// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Bridge API
//
// API for Oracle Cloud Bridge service.
//

package cloudbridge

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ComputeProperties Compute related properties.
type ComputeProperties struct {

	// Primary IP address of the compute instance.
	PrimaryIp *string `mandatory:"false" json:"primaryIp"`

	// Fully Qualified DNS Name.
	DnsName *string `mandatory:"false" json:"dnsName"`

	// Information about the asset.
	Description *string `mandatory:"false" json:"description"`

	// Number of CPUs.
	CoresCount *int `mandatory:"false" json:"coresCount"`

	// CPU model name.
	CpuModel *string `mandatory:"false" json:"cpuModel"`

	// Number of GPU devices.
	GpuDevicesCount *int `mandatory:"false" json:"gpuDevicesCount"`

	// List of GPU devices attached to a virtual machine.
	GpuDevices []GpuDevice `mandatory:"false" json:"gpuDevices"`

	// Number of threads per core.
	ThreadsPerCoreCount *int `mandatory:"false" json:"threadsPerCoreCount"`

	// Memory size in MBs.
	MemoryInMBs *int64 `mandatory:"false" json:"memoryInMBs"`

	// Whether Pmem is enabled. Decides if NVDIMMs are used as a permanent memory.
	IsPmemEnabled *bool `mandatory:"false" json:"isPmemEnabled"`

	// Pmem size in MBs.
	PmemInMBs *int64 `mandatory:"false" json:"pmemInMBs"`

	// Operating system.
	OperatingSystem *string `mandatory:"false" json:"operatingSystem"`

	// Operating system version.
	OperatingSystemVersion *string `mandatory:"false" json:"operatingSystemVersion"`

	// Host name of the VM.
	HostName *string `mandatory:"false" json:"hostName"`

	// The current power state of the virtual machine.
	PowerState *string `mandatory:"false" json:"powerState"`

	// Guest state.
	GuestState *string `mandatory:"false" json:"guestState"`

	// Whether Trusted Platform Module (TPM) is enabled.
	IsTpmEnabled *bool `mandatory:"false" json:"isTpmEnabled"`

	// Number of connected networks.
	ConnectedNetworks *int `mandatory:"false" json:"connectedNetworks"`

	// Number of network ethernet cards.
	NicsCount *int `mandatory:"false" json:"nicsCount"`

	// List of network ethernet cards attached to a virtual machine.
	Nics []Nic `mandatory:"false" json:"nics"`

	// Provision storage size in MBs.
	StorageProvisionedInMBs *int64 `mandatory:"false" json:"storageProvisionedInMBs"`

	// Number of disks.
	DisksCount *int `mandatory:"false" json:"disksCount"`

	// Lists the set of disks belonging to the virtual machine. This list is unordered.
	Disks []Disk `mandatory:"false" json:"disks"`

	// Information about firmware type for this virtual machine.
	Firmware *string `mandatory:"false" json:"firmware"`

	// Latency sensitivity.
	LatencySensitivity *string `mandatory:"false" json:"latencySensitivity"`

	// The properties of the NVDIMMs attached to a virtual machine.
	Nvdimms []Nvdimm `mandatory:"false" json:"nvdimms"`

	NvdimmController *NvdimmController `mandatory:"false" json:"nvdimmController"`

	ScsiController *ScsiController `mandatory:"false" json:"scsiController"`

	// Hardware version.
	HardwareVersion *string `mandatory:"false" json:"hardwareVersion"`
}

func (m ComputeProperties) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ComputeProperties) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
