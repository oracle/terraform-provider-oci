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

// ExadataIormConfigUpdateDetails IORM Setting details for this Exadata System to be updated
type ExadataIormConfigUpdateDetails struct {

	// Value for the IORM objective
	// Default is "Auto"
	Objective ExadataIormConfigUpdateDetailsObjectiveEnum `mandatory:"false" json:"objective,omitempty"`

	// Array of IORM Setting for all the database in
	// this Exadata DB System
	DbPlans []DbIormConfigUpdateDetail `mandatory:"false" json:"dbPlans"`
}

func (m ExadataIormConfigUpdateDetails) String() string {
	return common.PointerString(m)
}

// ExadataIormConfigUpdateDetailsObjectiveEnum Enum with underlying type: string
type ExadataIormConfigUpdateDetailsObjectiveEnum string

// Set of constants representing the allowable values for ExadataIormConfigUpdateDetailsObjectiveEnum
const (
	ExadataIormConfigUpdateDetailsObjectiveLowLatency     ExadataIormConfigUpdateDetailsObjectiveEnum = "LOW_LATENCY"
	ExadataIormConfigUpdateDetailsObjectiveHighThroughput ExadataIormConfigUpdateDetailsObjectiveEnum = "HIGH_THROUGHPUT"
	ExadataIormConfigUpdateDetailsObjectiveBalanced       ExadataIormConfigUpdateDetailsObjectiveEnum = "BALANCED"
	ExadataIormConfigUpdateDetailsObjectiveAuto           ExadataIormConfigUpdateDetailsObjectiveEnum = "AUTO"
	ExadataIormConfigUpdateDetailsObjectiveBasic          ExadataIormConfigUpdateDetailsObjectiveEnum = "BASIC"
)

var mappingExadataIormConfigUpdateDetailsObjective = map[string]ExadataIormConfigUpdateDetailsObjectiveEnum{
	"LOW_LATENCY":     ExadataIormConfigUpdateDetailsObjectiveLowLatency,
	"HIGH_THROUGHPUT": ExadataIormConfigUpdateDetailsObjectiveHighThroughput,
	"BALANCED":        ExadataIormConfigUpdateDetailsObjectiveBalanced,
	"AUTO":            ExadataIormConfigUpdateDetailsObjectiveAuto,
	"BASIC":           ExadataIormConfigUpdateDetailsObjectiveBasic,
}

// GetExadataIormConfigUpdateDetailsObjectiveEnumValues Enumerates the set of values for ExadataIormConfigUpdateDetailsObjectiveEnum
func GetExadataIormConfigUpdateDetailsObjectiveEnumValues() []ExadataIormConfigUpdateDetailsObjectiveEnum {
	values := make([]ExadataIormConfigUpdateDetailsObjectiveEnum, 0)
	for _, v := range mappingExadataIormConfigUpdateDetailsObjective {
		values = append(values, v)
	}
	return values
}
