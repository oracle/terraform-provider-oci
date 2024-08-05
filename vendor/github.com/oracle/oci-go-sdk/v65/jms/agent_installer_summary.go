// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Fleets API
//
// The APIs for the Fleet Management (https://docs.oracle.com/en-us/iaas/jms/doc/fleet-management.html) feature of Java Management Service to monitor and manage the usage of Java in your enterprise. Use these APIs to manage fleets, configure managed instances to report to fleets, and gain insights into the Java workloads running on these instances by carrying out basic and advanced features.
//

package jms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AgentInstallerSummary Supported agent installer downloads.
type AgentInstallerSummary struct {

	// Unique identifier for the agent installer.
	AgentInstallerId *int64 `mandatory:"true" json:"agentInstallerId"`

	// Description of the agent installer artifact. The description typically includes the OS, architecture, and agent installer type.
	AgentInstallerDescription *string `mandatory:"true" json:"agentInstallerDescription"`

	// Approximate compressed file size in bytes.
	ApproximateFileSizeInBytes *int64 `mandatory:"true" json:"approximateFileSizeInBytes"`

	// SHA256 checksum of the agent installer.
	Sha256 *string `mandatory:"true" json:"sha256"`

	// The target operating system family for the agent installer.
	OsFamily OsFamilyEnum `mandatory:"true" json:"osFamily"`

	// The target operating system architecture for the installer.
	PlatformArchitecture PlatformArchitectureTypeEnum `mandatory:"true" json:"platformArchitecture"`

	// The package type (typically the file extension) of the agent software included in the installer.
	PackageType PackageTypeEnum `mandatory:"true" json:"packageType"`

	// Agent image version.
	AgentVersion *string `mandatory:"true" json:"agentVersion"`

	// Java version.
	JavaVersion *string `mandatory:"true" json:"javaVersion"`

	// Agent installer version.
	AgentInstallerVersion *string `mandatory:"true" json:"agentInstallerVersion"`
}

func (m AgentInstallerSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AgentInstallerSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOsFamilyEnum(string(m.OsFamily)); !ok && m.OsFamily != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OsFamily: %s. Supported values are: %s.", m.OsFamily, strings.Join(GetOsFamilyEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPlatformArchitectureTypeEnum(string(m.PlatformArchitecture)); !ok && m.PlatformArchitecture != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlatformArchitecture: %s. Supported values are: %s.", m.PlatformArchitecture, strings.Join(GetPlatformArchitectureTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPackageTypeEnum(string(m.PackageType)); !ok && m.PackageType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PackageType: %s. Supported values are: %s.", m.PackageType, strings.Join(GetPackageTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
