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

// OlvmOperatingSystem Information describing the operating system. This is used for both virtual machines and hosts.
type OlvmOperatingSystem struct {

	// Configuration of the boot sequence of a virtual machine.
	Boot []OlvmBootDeviceEnum `mandatory:"false" json:"boot"`

	// Custom kernel parameters for starting the virtual machine if Linux operating system is used.
	CmdLine *string `mandatory:"false" json:"cmdLine"`

	// A custom part of the host kernel command line.
	CustomKernelCmdLine *string `mandatory:"false" json:"customKernelCmdLine"`

	// Path to custom initial ramdisk on ISO storage domain if Linux operating system is used.
	InitRd *string `mandatory:"false" json:"initRd"`

	// Path to custom kernel on ISO storage domain if Linux operating system is used.
	Kernel *string `mandatory:"false" json:"kernel"`

	// The host kernel command line as reported by a running host.
	ReportedKernelCmdLine *string `mandatory:"false" json:"reportedKernelCmdLine"`

	// Operating system name in human readable form
	Type *string `mandatory:"false" json:"type"`

	Version *OlvmVersion `mandatory:"false" json:"version"`
}

func (m OlvmOperatingSystem) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OlvmOperatingSystem) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
