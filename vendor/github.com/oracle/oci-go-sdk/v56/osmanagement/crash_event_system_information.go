// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management API
//
// API for the OS Management service. Use these API operations for working
// with Managed instances and Managed instance groups.
//

package osmanagement

import (
	"github.com/oracle/oci-go-sdk/v56/common"
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
