// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AutonomousDbVersionSummary The supported Autonomous AI Database version.
type AutonomousDbVersionSummary struct {

	// A valid Oracle AI Database version for Autonomous AI Database.
	Version *string `mandatory:"true" json:"version"`

	// The Autonomous AI Database workload type. The following values are valid:
	// - OLTP - indicates an Autonomous AI Transaction Processing database
	// - DW - indicates an Autonomous AI Lakehouse database
	// - AJD - indicates an Autonomous AI JSON Database
	// - APEX - indicates an Autonomous AI Database with the Oracle APEX AI Application Development workload type.
	// - LH - indicates an Oracle Autonomous AI Lakehouse database
	//
	// This cannot be updated in parallel with any of the following: licenseModel, dbEdition, cpuCoreCount, computeCount, computeModel, adminPassword, whitelistedIps, isMTLSConnectionRequired, privateEndpointLabel, nsgIds, dbVersion, isRefreshable, dbName, scheduledOperations, dbToolsDetails, isLocalDataGuardEnabled, or isFreeTier.
	DbWorkload AutonomousDbVersionSummaryDbWorkloadEnum `mandatory:"false" json:"dbWorkload,omitempty"`

	// True if the database uses dedicated Exadata infrastructure (https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html).
	IsDedicated *bool `mandatory:"false" json:"isDedicated"`

	// A URL that points to a detailed description of the Autonomous AI Database version.
	Details *string `mandatory:"false" json:"details"`

	// True if this version of the Oracle AI Database software can be used for Always-Free Autonomous AI Databases.
	IsFreeTierEnabled *bool `mandatory:"false" json:"isFreeTierEnabled"`

	// True if this Oracle AI Database software version can be used for Autonomous AI Databases for Developers.
	IsDevTierEnabled *bool `mandatory:"false" json:"isDevTierEnabled"`

	// True if this version of the Oracle AI Database software has payments enabled.
	IsPaidEnabled *bool `mandatory:"false" json:"isPaidEnabled"`

	// True if this version of the Oracle AI Database software's default is free.
	IsDefaultForFree *bool `mandatory:"false" json:"isDefaultForFree"`

	// True if this version of the Oracle AI Database software's default is paid.
	IsDefaultForPaid *bool `mandatory:"false" json:"isDefaultForPaid"`
}

func (m AutonomousDbVersionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutonomousDbVersionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAutonomousDbVersionSummaryDbWorkloadEnum(string(m.DbWorkload)); !ok && m.DbWorkload != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DbWorkload: %s. Supported values are: %s.", m.DbWorkload, strings.Join(GetAutonomousDbVersionSummaryDbWorkloadEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AutonomousDbVersionSummaryDbWorkloadEnum Enum with underlying type: string
type AutonomousDbVersionSummaryDbWorkloadEnum string

// Set of constants representing the allowable values for AutonomousDbVersionSummaryDbWorkloadEnum
const (
	AutonomousDbVersionSummaryDbWorkloadOltp AutonomousDbVersionSummaryDbWorkloadEnum = "OLTP"
	AutonomousDbVersionSummaryDbWorkloadDw   AutonomousDbVersionSummaryDbWorkloadEnum = "DW"
	AutonomousDbVersionSummaryDbWorkloadAjd  AutonomousDbVersionSummaryDbWorkloadEnum = "AJD"
	AutonomousDbVersionSummaryDbWorkloadApex AutonomousDbVersionSummaryDbWorkloadEnum = "APEX"
	AutonomousDbVersionSummaryDbWorkloadLh   AutonomousDbVersionSummaryDbWorkloadEnum = "LH"
)

var mappingAutonomousDbVersionSummaryDbWorkloadEnum = map[string]AutonomousDbVersionSummaryDbWorkloadEnum{
	"OLTP": AutonomousDbVersionSummaryDbWorkloadOltp,
	"DW":   AutonomousDbVersionSummaryDbWorkloadDw,
	"AJD":  AutonomousDbVersionSummaryDbWorkloadAjd,
	"APEX": AutonomousDbVersionSummaryDbWorkloadApex,
	"LH":   AutonomousDbVersionSummaryDbWorkloadLh,
}

var mappingAutonomousDbVersionSummaryDbWorkloadEnumLowerCase = map[string]AutonomousDbVersionSummaryDbWorkloadEnum{
	"oltp": AutonomousDbVersionSummaryDbWorkloadOltp,
	"dw":   AutonomousDbVersionSummaryDbWorkloadDw,
	"ajd":  AutonomousDbVersionSummaryDbWorkloadAjd,
	"apex": AutonomousDbVersionSummaryDbWorkloadApex,
	"lh":   AutonomousDbVersionSummaryDbWorkloadLh,
}

// GetAutonomousDbVersionSummaryDbWorkloadEnumValues Enumerates the set of values for AutonomousDbVersionSummaryDbWorkloadEnum
func GetAutonomousDbVersionSummaryDbWorkloadEnumValues() []AutonomousDbVersionSummaryDbWorkloadEnum {
	values := make([]AutonomousDbVersionSummaryDbWorkloadEnum, 0)
	for _, v := range mappingAutonomousDbVersionSummaryDbWorkloadEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDbVersionSummaryDbWorkloadEnumStringValues Enumerates the set of values in String for AutonomousDbVersionSummaryDbWorkloadEnum
func GetAutonomousDbVersionSummaryDbWorkloadEnumStringValues() []string {
	return []string{
		"OLTP",
		"DW",
		"AJD",
		"APEX",
		"LH",
	}
}

// GetMappingAutonomousDbVersionSummaryDbWorkloadEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDbVersionSummaryDbWorkloadEnum(val string) (AutonomousDbVersionSummaryDbWorkloadEnum, bool) {
	enum, ok := mappingAutonomousDbVersionSummaryDbWorkloadEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
