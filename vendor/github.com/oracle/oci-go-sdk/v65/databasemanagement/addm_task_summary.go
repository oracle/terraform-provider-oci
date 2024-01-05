// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AddmTaskSummary The object containing the ADDM task metadata.
type AddmTaskSummary struct {

	// The ID number of the ADDM task.
	TaskId *int64 `mandatory:"true" json:"taskId"`

	// The creation date of the ADDM task.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The name of the ADDM task.
	TaskName *string `mandatory:"false" json:"taskName"`

	// The description of the ADDM task.
	Description *string `mandatory:"false" json:"description"`

	// The database user who owns the ADDM task.
	DbUser *string `mandatory:"false" json:"dbUser"`

	// The status of the ADDM task.
	Status AddmTaskSummaryStatusEnum `mandatory:"false" json:"status,omitempty"`

	// A description of how the task was created.
	HowCreated AddmTaskSummaryHowCreatedEnum `mandatory:"false" json:"howCreated,omitempty"`

	// The timestamp of the beginning AWR snapshot used in the ADDM task as defined by date-time RFC3339 format.
	StartSnapshotTime *common.SDKTime `mandatory:"false" json:"startSnapshotTime"`

	// The timestamp of the ending AWR snapshot used in the ADDM task as defined by date-time RFC3339 format.
	EndSnapshotTime *common.SDKTime `mandatory:"false" json:"endSnapshotTime"`

	// The ID number of the beginning AWR snapshot.
	BeginSnapshotId *int64 `mandatory:"false" json:"beginSnapshotId"`

	// The ID number of the ending AWR snapshot.
	EndSnapshotId *int64 `mandatory:"false" json:"endSnapshotId"`

	// The number of ADDM findings.
	Findings *int64 `mandatory:"false" json:"findings"`
}

func (m AddmTaskSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AddmTaskSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAddmTaskSummaryStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetAddmTaskSummaryStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAddmTaskSummaryHowCreatedEnum(string(m.HowCreated)); !ok && m.HowCreated != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for HowCreated: %s. Supported values are: %s.", m.HowCreated, strings.Join(GetAddmTaskSummaryHowCreatedEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AddmTaskSummaryStatusEnum Enum with underlying type: string
type AddmTaskSummaryStatusEnum string

// Set of constants representing the allowable values for AddmTaskSummaryStatusEnum
const (
	AddmTaskSummaryStatusInitial     AddmTaskSummaryStatusEnum = "INITIAL"
	AddmTaskSummaryStatusExecuting   AddmTaskSummaryStatusEnum = "EXECUTING"
	AddmTaskSummaryStatusInterrupted AddmTaskSummaryStatusEnum = "INTERRUPTED"
	AddmTaskSummaryStatusCompleted   AddmTaskSummaryStatusEnum = "COMPLETED"
	AddmTaskSummaryStatusError       AddmTaskSummaryStatusEnum = "ERROR"
)

var mappingAddmTaskSummaryStatusEnum = map[string]AddmTaskSummaryStatusEnum{
	"INITIAL":     AddmTaskSummaryStatusInitial,
	"EXECUTING":   AddmTaskSummaryStatusExecuting,
	"INTERRUPTED": AddmTaskSummaryStatusInterrupted,
	"COMPLETED":   AddmTaskSummaryStatusCompleted,
	"ERROR":       AddmTaskSummaryStatusError,
}

var mappingAddmTaskSummaryStatusEnumLowerCase = map[string]AddmTaskSummaryStatusEnum{
	"initial":     AddmTaskSummaryStatusInitial,
	"executing":   AddmTaskSummaryStatusExecuting,
	"interrupted": AddmTaskSummaryStatusInterrupted,
	"completed":   AddmTaskSummaryStatusCompleted,
	"error":       AddmTaskSummaryStatusError,
}

// GetAddmTaskSummaryStatusEnumValues Enumerates the set of values for AddmTaskSummaryStatusEnum
func GetAddmTaskSummaryStatusEnumValues() []AddmTaskSummaryStatusEnum {
	values := make([]AddmTaskSummaryStatusEnum, 0)
	for _, v := range mappingAddmTaskSummaryStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetAddmTaskSummaryStatusEnumStringValues Enumerates the set of values in String for AddmTaskSummaryStatusEnum
func GetAddmTaskSummaryStatusEnumStringValues() []string {
	return []string{
		"INITIAL",
		"EXECUTING",
		"INTERRUPTED",
		"COMPLETED",
		"ERROR",
	}
}

// GetMappingAddmTaskSummaryStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAddmTaskSummaryStatusEnum(val string) (AddmTaskSummaryStatusEnum, bool) {
	enum, ok := mappingAddmTaskSummaryStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AddmTaskSummaryHowCreatedEnum Enum with underlying type: string
type AddmTaskSummaryHowCreatedEnum string

// Set of constants representing the allowable values for AddmTaskSummaryHowCreatedEnum
const (
	AddmTaskSummaryHowCreatedAuto   AddmTaskSummaryHowCreatedEnum = "AUTO"
	AddmTaskSummaryHowCreatedManual AddmTaskSummaryHowCreatedEnum = "MANUAL"
)

var mappingAddmTaskSummaryHowCreatedEnum = map[string]AddmTaskSummaryHowCreatedEnum{
	"AUTO":   AddmTaskSummaryHowCreatedAuto,
	"MANUAL": AddmTaskSummaryHowCreatedManual,
}

var mappingAddmTaskSummaryHowCreatedEnumLowerCase = map[string]AddmTaskSummaryHowCreatedEnum{
	"auto":   AddmTaskSummaryHowCreatedAuto,
	"manual": AddmTaskSummaryHowCreatedManual,
}

// GetAddmTaskSummaryHowCreatedEnumValues Enumerates the set of values for AddmTaskSummaryHowCreatedEnum
func GetAddmTaskSummaryHowCreatedEnumValues() []AddmTaskSummaryHowCreatedEnum {
	values := make([]AddmTaskSummaryHowCreatedEnum, 0)
	for _, v := range mappingAddmTaskSummaryHowCreatedEnum {
		values = append(values, v)
	}
	return values
}

// GetAddmTaskSummaryHowCreatedEnumStringValues Enumerates the set of values in String for AddmTaskSummaryHowCreatedEnum
func GetAddmTaskSummaryHowCreatedEnumStringValues() []string {
	return []string{
		"AUTO",
		"MANUAL",
	}
}

// GetMappingAddmTaskSummaryHowCreatedEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAddmTaskSummaryHowCreatedEnum(val string) (AddmTaskSummaryHowCreatedEnum, bool) {
	enum, ok := mappingAddmTaskSummaryHowCreatedEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
