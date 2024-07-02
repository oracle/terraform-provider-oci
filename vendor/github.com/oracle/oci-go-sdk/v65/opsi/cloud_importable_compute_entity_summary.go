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

// CloudImportableComputeEntitySummary A compute host entity that can be imported into Operations Insights.
type CloudImportableComputeEntitySummary struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Compute Instance
	ComputeId *string `mandatory:"true" json:"computeId"`

	// The Display Name (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm#Display) of the Compute Instance
	ComputeDisplayName *string `mandatory:"true" json:"computeDisplayName"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The host name. The host name is unique amongst the hosts managed by the same management agent.
	HostName *string `mandatory:"true" json:"hostName"`

	// Platform type.
	// Supported platformType(s) for MACS-managed external host insight: [LINUX, SOLARIS, WINDOWS].
	// Supported platformType(s) for MACS-managed cloud host insight: [LINUX].
	// Supported platformType(s) for EM-managed external host insight: [LINUX, SOLARIS, SUNOS, ZLINUX, WINDOWS, AIX, HP-UX].
	PlatformType CloudImportableComputeEntitySummaryPlatformTypeEnum `mandatory:"true" json:"platformType"`
}

// GetComputeId returns ComputeId
func (m CloudImportableComputeEntitySummary) GetComputeId() *string {
	return m.ComputeId
}

// GetComputeDisplayName returns ComputeDisplayName
func (m CloudImportableComputeEntitySummary) GetComputeDisplayName() *string {
	return m.ComputeDisplayName
}

// GetCompartmentId returns CompartmentId
func (m CloudImportableComputeEntitySummary) GetCompartmentId() *string {
	return m.CompartmentId
}

func (m CloudImportableComputeEntitySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CloudImportableComputeEntitySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCloudImportableComputeEntitySummaryPlatformTypeEnum(string(m.PlatformType)); !ok && m.PlatformType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlatformType: %s. Supported values are: %s.", m.PlatformType, strings.Join(GetCloudImportableComputeEntitySummaryPlatformTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CloudImportableComputeEntitySummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCloudImportableComputeEntitySummary CloudImportableComputeEntitySummary
	s := struct {
		DiscriminatorParam string `json:"entitySource"`
		MarshalTypeCloudImportableComputeEntitySummary
	}{
		"MACS_MANAGED_CLOUD_HOST",
		(MarshalTypeCloudImportableComputeEntitySummary)(m),
	}

	return json.Marshal(&s)
}

// CloudImportableComputeEntitySummaryPlatformTypeEnum Enum with underlying type: string
type CloudImportableComputeEntitySummaryPlatformTypeEnum string

// Set of constants representing the allowable values for CloudImportableComputeEntitySummaryPlatformTypeEnum
const (
	CloudImportableComputeEntitySummaryPlatformTypeLinux   CloudImportableComputeEntitySummaryPlatformTypeEnum = "LINUX"
	CloudImportableComputeEntitySummaryPlatformTypeSolaris CloudImportableComputeEntitySummaryPlatformTypeEnum = "SOLARIS"
	CloudImportableComputeEntitySummaryPlatformTypeSunos   CloudImportableComputeEntitySummaryPlatformTypeEnum = "SUNOS"
	CloudImportableComputeEntitySummaryPlatformTypeZlinux  CloudImportableComputeEntitySummaryPlatformTypeEnum = "ZLINUX"
	CloudImportableComputeEntitySummaryPlatformTypeWindows CloudImportableComputeEntitySummaryPlatformTypeEnum = "WINDOWS"
	CloudImportableComputeEntitySummaryPlatformTypeAix     CloudImportableComputeEntitySummaryPlatformTypeEnum = "AIX"
	CloudImportableComputeEntitySummaryPlatformTypeHpUx    CloudImportableComputeEntitySummaryPlatformTypeEnum = "HP_UX"
)

var mappingCloudImportableComputeEntitySummaryPlatformTypeEnum = map[string]CloudImportableComputeEntitySummaryPlatformTypeEnum{
	"LINUX":   CloudImportableComputeEntitySummaryPlatformTypeLinux,
	"SOLARIS": CloudImportableComputeEntitySummaryPlatformTypeSolaris,
	"SUNOS":   CloudImportableComputeEntitySummaryPlatformTypeSunos,
	"ZLINUX":  CloudImportableComputeEntitySummaryPlatformTypeZlinux,
	"WINDOWS": CloudImportableComputeEntitySummaryPlatformTypeWindows,
	"AIX":     CloudImportableComputeEntitySummaryPlatformTypeAix,
	"HP_UX":   CloudImportableComputeEntitySummaryPlatformTypeHpUx,
}

var mappingCloudImportableComputeEntitySummaryPlatformTypeEnumLowerCase = map[string]CloudImportableComputeEntitySummaryPlatformTypeEnum{
	"linux":   CloudImportableComputeEntitySummaryPlatformTypeLinux,
	"solaris": CloudImportableComputeEntitySummaryPlatformTypeSolaris,
	"sunos":   CloudImportableComputeEntitySummaryPlatformTypeSunos,
	"zlinux":  CloudImportableComputeEntitySummaryPlatformTypeZlinux,
	"windows": CloudImportableComputeEntitySummaryPlatformTypeWindows,
	"aix":     CloudImportableComputeEntitySummaryPlatformTypeAix,
	"hp_ux":   CloudImportableComputeEntitySummaryPlatformTypeHpUx,
}

// GetCloudImportableComputeEntitySummaryPlatformTypeEnumValues Enumerates the set of values for CloudImportableComputeEntitySummaryPlatformTypeEnum
func GetCloudImportableComputeEntitySummaryPlatformTypeEnumValues() []CloudImportableComputeEntitySummaryPlatformTypeEnum {
	values := make([]CloudImportableComputeEntitySummaryPlatformTypeEnum, 0)
	for _, v := range mappingCloudImportableComputeEntitySummaryPlatformTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCloudImportableComputeEntitySummaryPlatformTypeEnumStringValues Enumerates the set of values in String for CloudImportableComputeEntitySummaryPlatformTypeEnum
func GetCloudImportableComputeEntitySummaryPlatformTypeEnumStringValues() []string {
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

// GetMappingCloudImportableComputeEntitySummaryPlatformTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCloudImportableComputeEntitySummaryPlatformTypeEnum(val string) (CloudImportableComputeEntitySummaryPlatformTypeEnum, bool) {
	enum, ok := mappingCloudImportableComputeEntitySummaryPlatformTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
