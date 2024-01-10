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

// VmwareVmProperties VMware virtual machine related properties.
type VmwareVmProperties struct {

	// Cluster name.
	Cluster *string `mandatory:"false" json:"cluster"`

	// Customer fields.
	CustomerFields []string `mandatory:"false" json:"customerFields"`

	// Customer defined tags.
	CustomerTags []CustomerTag `mandatory:"false" json:"customerTags"`

	// vCenter-specific identifier of the virtual machine.
	InstanceUuid *string `mandatory:"false" json:"instanceUuid"`

	// Path directory of the asset.
	Path *string `mandatory:"false" json:"path"`

	// VMware tools status.
	VmwareToolsStatus *string `mandatory:"false" json:"vmwareToolsStatus"`

	// Whether changed block tracking for this VM's disk is active.
	IsDisksUuidEnabled *bool `mandatory:"false" json:"isDisksUuidEnabled"`

	// Indicates that change tracking is supported for virtual disks of this virtual machine.
	// However, even if change tracking is supported, it might not be available for all disks of the virtual machine.
	IsDisksCbtEnabled *bool `mandatory:"false" json:"isDisksCbtEnabled"`

	// Fault tolerance state.
	FaultToleranceState *string `mandatory:"false" json:"faultToleranceState"`

	// Fault tolerance bandwidth.
	FaultToleranceBandwidth *int `mandatory:"false" json:"faultToleranceBandwidth"`

	// Fault tolerance to secondary latency.
	FaultToleranceSecondaryLatency *int `mandatory:"false" json:"faultToleranceSecondaryLatency"`
}

func (m VmwareVmProperties) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VmwareVmProperties) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
