// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// HostImportableAgentEntitySummary An agent host entity that can be imported into Operations Insights.
type HostImportableAgentEntitySummary struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Management Agent
	ManagementAgentId *string `mandatory:"true" json:"managementAgentId"`

	// The Display Name (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm#Display) of the Management Agent
	ManagementAgentDisplayName *string `mandatory:"true" json:"managementAgentDisplayName"`

	// The host name. The host name is unique amongst the hosts managed by the same management agent.
	HostName *string `mandatory:"true" json:"hostName"`

	// Platform type.
	// Supported platformType(s) for MACS-managed external host insight: [LINUX, SOLARIS, WINDOWS].
	// Supported platformType(s) for MACS-managed cloud host insight: [LINUX].
	// Supported platformType(s) for EM-managed external host insight: [LINUX, SOLARIS, SUNOS, ZLINUX, WINDOWS, AIX, HP-UX].
	PlatformType HostImportableAgentEntitySummaryPlatformTypeEnum `mandatory:"true" json:"platformType"`
}

// GetManagementAgentId returns ManagementAgentId
func (m HostImportableAgentEntitySummary) GetManagementAgentId() *string {
	return m.ManagementAgentId
}

// GetManagementAgentDisplayName returns ManagementAgentDisplayName
func (m HostImportableAgentEntitySummary) GetManagementAgentDisplayName() *string {
	return m.ManagementAgentDisplayName
}

func (m HostImportableAgentEntitySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HostImportableAgentEntitySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingHostImportableAgentEntitySummaryPlatformTypeEnum(string(m.PlatformType)); !ok && m.PlatformType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlatformType: %s. Supported values are: %s.", m.PlatformType, strings.Join(GetHostImportableAgentEntitySummaryPlatformTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m HostImportableAgentEntitySummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeHostImportableAgentEntitySummary HostImportableAgentEntitySummary
	s := struct {
		DiscriminatorParam string `json:"entitySource"`
		MarshalTypeHostImportableAgentEntitySummary
	}{
		"MACS_MANAGED_EXTERNAL_HOST",
		(MarshalTypeHostImportableAgentEntitySummary)(m),
	}

	return json.Marshal(&s)
}

// HostImportableAgentEntitySummaryPlatformTypeEnum Enum with underlying type: string
type HostImportableAgentEntitySummaryPlatformTypeEnum string

// Set of constants representing the allowable values for HostImportableAgentEntitySummaryPlatformTypeEnum
const (
	HostImportableAgentEntitySummaryPlatformTypeLinux   HostImportableAgentEntitySummaryPlatformTypeEnum = "LINUX"
	HostImportableAgentEntitySummaryPlatformTypeSolaris HostImportableAgentEntitySummaryPlatformTypeEnum = "SOLARIS"
	HostImportableAgentEntitySummaryPlatformTypeSunos   HostImportableAgentEntitySummaryPlatformTypeEnum = "SUNOS"
	HostImportableAgentEntitySummaryPlatformTypeZlinux  HostImportableAgentEntitySummaryPlatformTypeEnum = "ZLINUX"
	HostImportableAgentEntitySummaryPlatformTypeWindows HostImportableAgentEntitySummaryPlatformTypeEnum = "WINDOWS"
	HostImportableAgentEntitySummaryPlatformTypeAix     HostImportableAgentEntitySummaryPlatformTypeEnum = "AIX"
	HostImportableAgentEntitySummaryPlatformTypeHpUx    HostImportableAgentEntitySummaryPlatformTypeEnum = "HP_UX"
)

var mappingHostImportableAgentEntitySummaryPlatformTypeEnum = map[string]HostImportableAgentEntitySummaryPlatformTypeEnum{
	"LINUX":   HostImportableAgentEntitySummaryPlatformTypeLinux,
	"SOLARIS": HostImportableAgentEntitySummaryPlatformTypeSolaris,
	"SUNOS":   HostImportableAgentEntitySummaryPlatformTypeSunos,
	"ZLINUX":  HostImportableAgentEntitySummaryPlatformTypeZlinux,
	"WINDOWS": HostImportableAgentEntitySummaryPlatformTypeWindows,
	"AIX":     HostImportableAgentEntitySummaryPlatformTypeAix,
	"HP_UX":   HostImportableAgentEntitySummaryPlatformTypeHpUx,
}

var mappingHostImportableAgentEntitySummaryPlatformTypeEnumLowerCase = map[string]HostImportableAgentEntitySummaryPlatformTypeEnum{
	"linux":   HostImportableAgentEntitySummaryPlatformTypeLinux,
	"solaris": HostImportableAgentEntitySummaryPlatformTypeSolaris,
	"sunos":   HostImportableAgentEntitySummaryPlatformTypeSunos,
	"zlinux":  HostImportableAgentEntitySummaryPlatformTypeZlinux,
	"windows": HostImportableAgentEntitySummaryPlatformTypeWindows,
	"aix":     HostImportableAgentEntitySummaryPlatformTypeAix,
	"hp_ux":   HostImportableAgentEntitySummaryPlatformTypeHpUx,
}

// GetHostImportableAgentEntitySummaryPlatformTypeEnumValues Enumerates the set of values for HostImportableAgentEntitySummaryPlatformTypeEnum
func GetHostImportableAgentEntitySummaryPlatformTypeEnumValues() []HostImportableAgentEntitySummaryPlatformTypeEnum {
	values := make([]HostImportableAgentEntitySummaryPlatformTypeEnum, 0)
	for _, v := range mappingHostImportableAgentEntitySummaryPlatformTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetHostImportableAgentEntitySummaryPlatformTypeEnumStringValues Enumerates the set of values in String for HostImportableAgentEntitySummaryPlatformTypeEnum
func GetHostImportableAgentEntitySummaryPlatformTypeEnumStringValues() []string {
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

// GetMappingHostImportableAgentEntitySummaryPlatformTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHostImportableAgentEntitySummaryPlatformTypeEnum(val string) (HostImportableAgentEntitySummaryPlatformTypeEnum, bool) {
	enum, ok := mappingHostImportableAgentEntitySummaryPlatformTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
