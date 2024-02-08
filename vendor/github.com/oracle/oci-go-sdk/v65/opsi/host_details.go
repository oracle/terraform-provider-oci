// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// HostDetails Partial information about a host which includes id, name, type.
type HostDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the host.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The host name. The host name is unique amongst the hosts managed by the same management agent.
	HostName *string `mandatory:"true" json:"hostName"`

	// Platform type.
	// Supported platformType(s) for MACS-managed external host insight: [LINUX, SOLARIS, WINDOWS].
	// Supported platformType(s) for MACS-managed cloud host insight: [LINUX].
	// Supported platformType(s) for EM-managed external host insight: [LINUX, SOLARIS, SUNOS, ZLINUX, WINDOWS, AIX, HP-UX].
	PlatformType HostDetailsPlatformTypeEnum `mandatory:"true" json:"platformType"`

	// The identifier of the agent.
	AgentIdentifier *string `mandatory:"true" json:"agentIdentifier"`

	// The user-friendly name for the host. The name does not have to be unique.
	HostDisplayName *string `mandatory:"false" json:"hostDisplayName"`
}

func (m HostDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HostDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingHostDetailsPlatformTypeEnum(string(m.PlatformType)); !ok && m.PlatformType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlatformType: %s. Supported values are: %s.", m.PlatformType, strings.Join(GetHostDetailsPlatformTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// HostDetailsPlatformTypeEnum Enum with underlying type: string
type HostDetailsPlatformTypeEnum string

// Set of constants representing the allowable values for HostDetailsPlatformTypeEnum
const (
	HostDetailsPlatformTypeLinux   HostDetailsPlatformTypeEnum = "LINUX"
	HostDetailsPlatformTypeSolaris HostDetailsPlatformTypeEnum = "SOLARIS"
	HostDetailsPlatformTypeSunos   HostDetailsPlatformTypeEnum = "SUNOS"
	HostDetailsPlatformTypeZlinux  HostDetailsPlatformTypeEnum = "ZLINUX"
	HostDetailsPlatformTypeWindows HostDetailsPlatformTypeEnum = "WINDOWS"
	HostDetailsPlatformTypeAix     HostDetailsPlatformTypeEnum = "AIX"
	HostDetailsPlatformTypeHpUx    HostDetailsPlatformTypeEnum = "HP_UX"
)

var mappingHostDetailsPlatformTypeEnum = map[string]HostDetailsPlatformTypeEnum{
	"LINUX":   HostDetailsPlatformTypeLinux,
	"SOLARIS": HostDetailsPlatformTypeSolaris,
	"SUNOS":   HostDetailsPlatformTypeSunos,
	"ZLINUX":  HostDetailsPlatformTypeZlinux,
	"WINDOWS": HostDetailsPlatformTypeWindows,
	"AIX":     HostDetailsPlatformTypeAix,
	"HP_UX":   HostDetailsPlatformTypeHpUx,
}

var mappingHostDetailsPlatformTypeEnumLowerCase = map[string]HostDetailsPlatformTypeEnum{
	"linux":   HostDetailsPlatformTypeLinux,
	"solaris": HostDetailsPlatformTypeSolaris,
	"sunos":   HostDetailsPlatformTypeSunos,
	"zlinux":  HostDetailsPlatformTypeZlinux,
	"windows": HostDetailsPlatformTypeWindows,
	"aix":     HostDetailsPlatformTypeAix,
	"hp_ux":   HostDetailsPlatformTypeHpUx,
}

// GetHostDetailsPlatformTypeEnumValues Enumerates the set of values for HostDetailsPlatformTypeEnum
func GetHostDetailsPlatformTypeEnumValues() []HostDetailsPlatformTypeEnum {
	values := make([]HostDetailsPlatformTypeEnum, 0)
	for _, v := range mappingHostDetailsPlatformTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetHostDetailsPlatformTypeEnumStringValues Enumerates the set of values in String for HostDetailsPlatformTypeEnum
func GetHostDetailsPlatformTypeEnumStringValues() []string {
	return []string{
		"LINUX",
		"SOLARIS",
		"SUNOS",
		"ZLINUX",
		"WINDOWS",
		"AIX",
		"HP_UX",
	}
}

// GetMappingHostDetailsPlatformTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHostDetailsPlatformTypeEnum(val string) (HostDetailsPlatformTypeEnum, bool) {
	enum, ok := mappingHostDetailsPlatformTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
