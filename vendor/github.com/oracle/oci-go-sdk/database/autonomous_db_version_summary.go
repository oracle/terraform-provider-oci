// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// AutonomousDbVersionSummary The supported Autonomous Database version.
type AutonomousDbVersionSummary struct {

	// A valid Oracle Database version for Autonomous Database.
	Version *string `mandatory:"true" json:"version"`

	// The Autonomous Database workload type. The following values are valid:
	// - OLTP - indicates an Autonomous Transaction Processing database
	// - DW - indicates an Autonomous Data Warehouse database
	// - AJD - indicates an Autonomous JSON Database
	DbWorkload AutonomousDbVersionSummaryDbWorkloadEnum `mandatory:"false" json:"dbWorkload,omitempty"`

	// True if the database uses dedicated Exadata infrastructure (https://docs.cloud.oracle.com/Content/Database/Concepts/adbddoverview.htm).
	IsDedicated *bool `mandatory:"false" json:"isDedicated"`

	// A URL that points to a detailed description of the Autonomous Database version.
	Details *string `mandatory:"false" json:"details"`

	// True if this version of the Oracle Database software can be used for Always-Free Autonomous Databases.
	IsFreeTierEnabled *bool `mandatory:"false" json:"isFreeTierEnabled"`
}

func (m AutonomousDbVersionSummary) String() string {
	return common.PointerString(m)
}

// AutonomousDbVersionSummaryDbWorkloadEnum Enum with underlying type: string
type AutonomousDbVersionSummaryDbWorkloadEnum string

// Set of constants representing the allowable values for AutonomousDbVersionSummaryDbWorkloadEnum
const (
	AutonomousDbVersionSummaryDbWorkloadOltp AutonomousDbVersionSummaryDbWorkloadEnum = "OLTP"
	AutonomousDbVersionSummaryDbWorkloadDw   AutonomousDbVersionSummaryDbWorkloadEnum = "DW"
	AutonomousDbVersionSummaryDbWorkloadAjd  AutonomousDbVersionSummaryDbWorkloadEnum = "AJD"
)

var mappingAutonomousDbVersionSummaryDbWorkload = map[string]AutonomousDbVersionSummaryDbWorkloadEnum{
	"OLTP": AutonomousDbVersionSummaryDbWorkloadOltp,
	"DW":   AutonomousDbVersionSummaryDbWorkloadDw,
	"AJD":  AutonomousDbVersionSummaryDbWorkloadAjd,
}

// GetAutonomousDbVersionSummaryDbWorkloadEnumValues Enumerates the set of values for AutonomousDbVersionSummaryDbWorkloadEnum
func GetAutonomousDbVersionSummaryDbWorkloadEnumValues() []AutonomousDbVersionSummaryDbWorkloadEnum {
	values := make([]AutonomousDbVersionSummaryDbWorkloadEnum, 0)
	for _, v := range mappingAutonomousDbVersionSummaryDbWorkload {
		values = append(values, v)
	}
	return values
}
