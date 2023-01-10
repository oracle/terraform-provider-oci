// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExadataIormConfigUpdateDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingExadataIormConfigUpdateDetailsObjectiveEnum(string(m.Objective)); !ok && m.Objective != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Objective: %s. Supported values are: %s.", m.Objective, strings.Join(GetExadataIormConfigUpdateDetailsObjectiveEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingExadataIormConfigUpdateDetailsObjectiveEnum = map[string]ExadataIormConfigUpdateDetailsObjectiveEnum{
	"LOW_LATENCY":     ExadataIormConfigUpdateDetailsObjectiveLowLatency,
	"HIGH_THROUGHPUT": ExadataIormConfigUpdateDetailsObjectiveHighThroughput,
	"BALANCED":        ExadataIormConfigUpdateDetailsObjectiveBalanced,
	"AUTO":            ExadataIormConfigUpdateDetailsObjectiveAuto,
	"BASIC":           ExadataIormConfigUpdateDetailsObjectiveBasic,
}

var mappingExadataIormConfigUpdateDetailsObjectiveEnumLowerCase = map[string]ExadataIormConfigUpdateDetailsObjectiveEnum{
	"low_latency":     ExadataIormConfigUpdateDetailsObjectiveLowLatency,
	"high_throughput": ExadataIormConfigUpdateDetailsObjectiveHighThroughput,
	"balanced":        ExadataIormConfigUpdateDetailsObjectiveBalanced,
	"auto":            ExadataIormConfigUpdateDetailsObjectiveAuto,
	"basic":           ExadataIormConfigUpdateDetailsObjectiveBasic,
}

// GetExadataIormConfigUpdateDetailsObjectiveEnumValues Enumerates the set of values for ExadataIormConfigUpdateDetailsObjectiveEnum
func GetExadataIormConfigUpdateDetailsObjectiveEnumValues() []ExadataIormConfigUpdateDetailsObjectiveEnum {
	values := make([]ExadataIormConfigUpdateDetailsObjectiveEnum, 0)
	for _, v := range mappingExadataIormConfigUpdateDetailsObjectiveEnum {
		values = append(values, v)
	}
	return values
}

// GetExadataIormConfigUpdateDetailsObjectiveEnumStringValues Enumerates the set of values in String for ExadataIormConfigUpdateDetailsObjectiveEnum
func GetExadataIormConfigUpdateDetailsObjectiveEnumStringValues() []string {
	return []string{
		"LOW_LATENCY",
		"HIGH_THROUGHPUT",
		"BALANCED",
		"AUTO",
		"BASIC",
	}
}

// GetMappingExadataIormConfigUpdateDetailsObjectiveEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExadataIormConfigUpdateDetailsObjectiveEnum(val string) (ExadataIormConfigUpdateDetailsObjectiveEnum, bool) {
	enum, ok := mappingExadataIormConfigUpdateDetailsObjectiveEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
