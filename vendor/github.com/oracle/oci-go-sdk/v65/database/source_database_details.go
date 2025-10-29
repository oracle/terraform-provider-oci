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

// SourceDatabaseDetails Source Autonomous AI Database details.
type SourceDatabaseDetails struct {

	// Autonomous VM cluster's user-friendly name.
	AutonomousVmClusterDisplayName *string `mandatory:"false" json:"autonomousVmClusterDisplayName"`

	// Autonomous Container Database name.
	AutonomousContainerDatabaseName *string `mandatory:"false" json:"autonomousContainerDatabaseName"`

	// The user-provided name for the Autonomous Container Database.
	AutonomousContainerDatabaseDisplayName *string `mandatory:"false" json:"autonomousContainerDatabaseDisplayName"`

	// Customer Contacts for the Autonomous Container Database. Setting this to an empty list removes all customer contacts.
	AutonomousContainerDatabaseCustomerContacts []CustomerContact `mandatory:"false" json:"autonomousContainerDatabaseCustomerContacts"`

	// DST Time-Zone File version of the Autonomous Container Database.
	AutonomousContainerDatabaseDstFileVersion *string `mandatory:"false" json:"autonomousContainerDatabaseDstFileVersion"`

	// Autonomous AI Database's name.
	AutonomousDatabaseName *string `mandatory:"false" json:"autonomousDatabaseName"`

	// Customer Contacts for the Autonomous AI Database.
	AutonomousDatabaseCustomerContacts []CustomerContact `mandatory:"false" json:"autonomousDatabaseCustomerContacts"`

	// The Autonomous AI Database workload type. The following values are valid:
	// - OLTP - indicates an Autonomous AI Transaction Processing database
	// - DW - indicates an Autonomous AI Lakehouse database
	// - AJD - indicates an Autonomous AI JSON Database
	// - APEX - indicates an Autonomous AI Database with the Oracle APEX AI Application Development workload type.
	// - LH - indicates an Oracle Autonomous AI Lakehouse database
	//
	// This cannot be updated in parallel with any of the following: licenseModel, dbEdition, cpuCoreCount, computeCount, computeModel, adminPassword, whitelistedIps, isMTLSConnectionRequired, privateEndpointLabel, nsgIds, dbVersion, isRefreshable, dbName, scheduledOperations, dbToolsDetails, isLocalDataGuardEnabled, or isFreeTier.
	DbWorkload SourceDatabaseDetailsDbWorkloadEnum `mandatory:"false" json:"dbWorkload,omitempty"`
}

func (m SourceDatabaseDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SourceDatabaseDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSourceDatabaseDetailsDbWorkloadEnum(string(m.DbWorkload)); !ok && m.DbWorkload != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DbWorkload: %s. Supported values are: %s.", m.DbWorkload, strings.Join(GetSourceDatabaseDetailsDbWorkloadEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SourceDatabaseDetailsDbWorkloadEnum Enum with underlying type: string
type SourceDatabaseDetailsDbWorkloadEnum string

// Set of constants representing the allowable values for SourceDatabaseDetailsDbWorkloadEnum
const (
	SourceDatabaseDetailsDbWorkloadOltp SourceDatabaseDetailsDbWorkloadEnum = "OLTP"
	SourceDatabaseDetailsDbWorkloadDw   SourceDatabaseDetailsDbWorkloadEnum = "DW"
	SourceDatabaseDetailsDbWorkloadAjd  SourceDatabaseDetailsDbWorkloadEnum = "AJD"
	SourceDatabaseDetailsDbWorkloadApex SourceDatabaseDetailsDbWorkloadEnum = "APEX"
	SourceDatabaseDetailsDbWorkloadLh   SourceDatabaseDetailsDbWorkloadEnum = "LH"
)

var mappingSourceDatabaseDetailsDbWorkloadEnum = map[string]SourceDatabaseDetailsDbWorkloadEnum{
	"OLTP": SourceDatabaseDetailsDbWorkloadOltp,
	"DW":   SourceDatabaseDetailsDbWorkloadDw,
	"AJD":  SourceDatabaseDetailsDbWorkloadAjd,
	"APEX": SourceDatabaseDetailsDbWorkloadApex,
	"LH":   SourceDatabaseDetailsDbWorkloadLh,
}

var mappingSourceDatabaseDetailsDbWorkloadEnumLowerCase = map[string]SourceDatabaseDetailsDbWorkloadEnum{
	"oltp": SourceDatabaseDetailsDbWorkloadOltp,
	"dw":   SourceDatabaseDetailsDbWorkloadDw,
	"ajd":  SourceDatabaseDetailsDbWorkloadAjd,
	"apex": SourceDatabaseDetailsDbWorkloadApex,
	"lh":   SourceDatabaseDetailsDbWorkloadLh,
}

// GetSourceDatabaseDetailsDbWorkloadEnumValues Enumerates the set of values for SourceDatabaseDetailsDbWorkloadEnum
func GetSourceDatabaseDetailsDbWorkloadEnumValues() []SourceDatabaseDetailsDbWorkloadEnum {
	values := make([]SourceDatabaseDetailsDbWorkloadEnum, 0)
	for _, v := range mappingSourceDatabaseDetailsDbWorkloadEnum {
		values = append(values, v)
	}
	return values
}

// GetSourceDatabaseDetailsDbWorkloadEnumStringValues Enumerates the set of values in String for SourceDatabaseDetailsDbWorkloadEnum
func GetSourceDatabaseDetailsDbWorkloadEnumStringValues() []string {
	return []string{
		"OLTP",
		"DW",
		"AJD",
		"APEX",
		"LH",
	}
}

// GetMappingSourceDatabaseDetailsDbWorkloadEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSourceDatabaseDetailsDbWorkloadEnum(val string) (SourceDatabaseDetailsDbWorkloadEnum, bool) {
	enum, ok := mappingSourceDatabaseDetailsDbWorkloadEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
