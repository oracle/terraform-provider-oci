// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management API
//
// API for the OS Management service. Use these API operations for working
// with Managed instances and Managed instance groups.
//

package osmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CrashEventSystemInformation Detailed information about system at the time of the crash.
type CrashEventSystemInformation struct {

	// system architecture
	Architecture ArchTypesEnum `mandatory:"false" json:"architecture,omitempty"`

	// Active ksplice kernel version (uptrack-uname -r)
	KspliceEffectiveKernelVersion *string `mandatory:"false" json:"kspliceEffectiveKernelVersion"`

	// The Operating System type of the managed instance.
	OsFamily OsFamiliesEnum `mandatory:"false" json:"osFamily,omitempty"`

	// Operating System Name (OCA value)
	OsName *string `mandatory:"false" json:"osName"`

	// Operating System Kernel Release (uname -v)
	OsKernelRelease *string `mandatory:"false" json:"osKernelRelease"`

	// Operating System Kernel Version (uname -r)
	OsKernelVersion *string `mandatory:"false" json:"osKernelVersion"`

	// Version of the OS (VERSION from /etc/os-release)
	OsSystemVersion *string `mandatory:"false" json:"osSystemVersion"`
}

func (m CrashEventSystemInformation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CrashEventSystemInformation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingArchTypesEnum(string(m.Architecture)); !ok && m.Architecture != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Architecture: %s. Supported values are: %s.", m.Architecture, strings.Join(GetArchTypesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOsFamiliesEnum(string(m.OsFamily)); !ok && m.OsFamily != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OsFamily: %s. Supported values are: %s.", m.OsFamily, strings.Join(GetOsFamiliesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
