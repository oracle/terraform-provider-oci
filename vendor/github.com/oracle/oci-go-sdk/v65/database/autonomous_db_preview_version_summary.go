// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AutonomousDbPreviewVersionSummary The Autonomous Database preview version. Note that preview version software is only available for Autonomous Database Serverless instances (https://docs.oracle.com/en/cloud/paas/autonomous-database/serverless/).
type AutonomousDbPreviewVersionSummary struct {

	// A valid Autonomous Database preview version.
	Version *string `mandatory:"true" json:"version"`

	// The date and time when the preview version availability begins.
	TimePreviewBegin *common.SDKTime `mandatory:"false" json:"timePreviewBegin"`

	// The date and time when the preview version availability ends.
	TimePreviewEnd *common.SDKTime `mandatory:"false" json:"timePreviewEnd"`

	// The Autonomous Database workload type. The following values are valid:
	// - OLTP - indicates an Autonomous Transaction Processing database
	// - DW - indicates an Autonomous Data Warehouse database
	// - AJD - indicates an Autonomous JSON Database
	// - APEX - indicates an Autonomous Database with the Oracle APEX Application Development workload type.
	// This cannot be updated in parallel with any of the following: licenseModel, dbEdition, cpuCoreCount, computeCount, computeModel, adminPassword, whitelistedIps, isMTLSConnectionRequired, privateEndpointLabel, nsgIds, dbVersion, isRefreshable, dbName, scheduledOperations, dbToolsDetails, isLocalDataGuardEnabled, or isFreeTier.
	DbWorkload AutonomousDbPreviewVersionSummaryDbWorkloadEnum `mandatory:"false" json:"dbWorkload,omitempty"`

	// A URL that points to a detailed description of the preview version.
	Details *string `mandatory:"false" json:"details"`
}

func (m AutonomousDbPreviewVersionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutonomousDbPreviewVersionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAutonomousDbPreviewVersionSummaryDbWorkloadEnum(string(m.DbWorkload)); !ok && m.DbWorkload != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DbWorkload: %s. Supported values are: %s.", m.DbWorkload, strings.Join(GetAutonomousDbPreviewVersionSummaryDbWorkloadEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AutonomousDbPreviewVersionSummaryDbWorkloadEnum Enum with underlying type: string
type AutonomousDbPreviewVersionSummaryDbWorkloadEnum string

// Set of constants representing the allowable values for AutonomousDbPreviewVersionSummaryDbWorkloadEnum
const (
	AutonomousDbPreviewVersionSummaryDbWorkloadOltp AutonomousDbPreviewVersionSummaryDbWorkloadEnum = "OLTP"
	AutonomousDbPreviewVersionSummaryDbWorkloadDw   AutonomousDbPreviewVersionSummaryDbWorkloadEnum = "DW"
	AutonomousDbPreviewVersionSummaryDbWorkloadAjd  AutonomousDbPreviewVersionSummaryDbWorkloadEnum = "AJD"
	AutonomousDbPreviewVersionSummaryDbWorkloadApex AutonomousDbPreviewVersionSummaryDbWorkloadEnum = "APEX"
)

var mappingAutonomousDbPreviewVersionSummaryDbWorkloadEnum = map[string]AutonomousDbPreviewVersionSummaryDbWorkloadEnum{
	"OLTP": AutonomousDbPreviewVersionSummaryDbWorkloadOltp,
	"DW":   AutonomousDbPreviewVersionSummaryDbWorkloadDw,
	"AJD":  AutonomousDbPreviewVersionSummaryDbWorkloadAjd,
	"APEX": AutonomousDbPreviewVersionSummaryDbWorkloadApex,
}

var mappingAutonomousDbPreviewVersionSummaryDbWorkloadEnumLowerCase = map[string]AutonomousDbPreviewVersionSummaryDbWorkloadEnum{
	"oltp": AutonomousDbPreviewVersionSummaryDbWorkloadOltp,
	"dw":   AutonomousDbPreviewVersionSummaryDbWorkloadDw,
	"ajd":  AutonomousDbPreviewVersionSummaryDbWorkloadAjd,
	"apex": AutonomousDbPreviewVersionSummaryDbWorkloadApex,
}

// GetAutonomousDbPreviewVersionSummaryDbWorkloadEnumValues Enumerates the set of values for AutonomousDbPreviewVersionSummaryDbWorkloadEnum
func GetAutonomousDbPreviewVersionSummaryDbWorkloadEnumValues() []AutonomousDbPreviewVersionSummaryDbWorkloadEnum {
	values := make([]AutonomousDbPreviewVersionSummaryDbWorkloadEnum, 0)
	for _, v := range mappingAutonomousDbPreviewVersionSummaryDbWorkloadEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDbPreviewVersionSummaryDbWorkloadEnumStringValues Enumerates the set of values in String for AutonomousDbPreviewVersionSummaryDbWorkloadEnum
func GetAutonomousDbPreviewVersionSummaryDbWorkloadEnumStringValues() []string {
	return []string{
		"OLTP",
		"DW",
		"AJD",
		"APEX",
	}
}

// GetMappingAutonomousDbPreviewVersionSummaryDbWorkloadEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDbPreviewVersionSummaryDbWorkloadEnum(val string) (AutonomousDbPreviewVersionSummaryDbWorkloadEnum, bool) {
	enum, ok := mappingAutonomousDbPreviewVersionSummaryDbWorkloadEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
