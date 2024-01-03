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

// HistoricAddmResult The details of the historic ADDM task.
type HistoricAddmResult struct {

	// The ID of the historic ADDM task.
	TaskId *int64 `mandatory:"true" json:"taskId"`

	// The creation date of the ADDM task.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Specifies whether the ADDM task returned had already existed or was newly created by the api call.
	IsNewlyCreated *bool `mandatory:"false" json:"isNewlyCreated"`

	// The name of the historic ADDM task.
	TaskName *string `mandatory:"false" json:"taskName"`

	// The description of the ADDM task.
	Description *string `mandatory:"false" json:"description"`

	// The database user who owns the historic ADDM task.
	DbUser *string `mandatory:"false" json:"dbUser"`

	// The status of the ADDM task.
	Status HistoricAddmResultStatusEnum `mandatory:"false" json:"status,omitempty"`

	// A description of how the task was created.
	HowCreated HistoricAddmResultHowCreatedEnum `mandatory:"false" json:"howCreated,omitempty"`

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

func (m HistoricAddmResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HistoricAddmResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingHistoricAddmResultStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetHistoricAddmResultStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingHistoricAddmResultHowCreatedEnum(string(m.HowCreated)); !ok && m.HowCreated != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for HowCreated: %s. Supported values are: %s.", m.HowCreated, strings.Join(GetHistoricAddmResultHowCreatedEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// HistoricAddmResultStatusEnum Enum with underlying type: string
type HistoricAddmResultStatusEnum string

// Set of constants representing the allowable values for HistoricAddmResultStatusEnum
const (
	HistoricAddmResultStatusInitial     HistoricAddmResultStatusEnum = "INITIAL"
	HistoricAddmResultStatusExecuting   HistoricAddmResultStatusEnum = "EXECUTING"
	HistoricAddmResultStatusInterrupted HistoricAddmResultStatusEnum = "INTERRUPTED"
	HistoricAddmResultStatusCompleted   HistoricAddmResultStatusEnum = "COMPLETED"
	HistoricAddmResultStatusError       HistoricAddmResultStatusEnum = "ERROR"
)

var mappingHistoricAddmResultStatusEnum = map[string]HistoricAddmResultStatusEnum{
	"INITIAL":     HistoricAddmResultStatusInitial,
	"EXECUTING":   HistoricAddmResultStatusExecuting,
	"INTERRUPTED": HistoricAddmResultStatusInterrupted,
	"COMPLETED":   HistoricAddmResultStatusCompleted,
	"ERROR":       HistoricAddmResultStatusError,
}

var mappingHistoricAddmResultStatusEnumLowerCase = map[string]HistoricAddmResultStatusEnum{
	"initial":     HistoricAddmResultStatusInitial,
	"executing":   HistoricAddmResultStatusExecuting,
	"interrupted": HistoricAddmResultStatusInterrupted,
	"completed":   HistoricAddmResultStatusCompleted,
	"error":       HistoricAddmResultStatusError,
}

// GetHistoricAddmResultStatusEnumValues Enumerates the set of values for HistoricAddmResultStatusEnum
func GetHistoricAddmResultStatusEnumValues() []HistoricAddmResultStatusEnum {
	values := make([]HistoricAddmResultStatusEnum, 0)
	for _, v := range mappingHistoricAddmResultStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetHistoricAddmResultStatusEnumStringValues Enumerates the set of values in String for HistoricAddmResultStatusEnum
func GetHistoricAddmResultStatusEnumStringValues() []string {
	return []string{
		"INITIAL",
		"EXECUTING",
		"INTERRUPTED",
		"COMPLETED",
		"ERROR",
	}
}

// GetMappingHistoricAddmResultStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHistoricAddmResultStatusEnum(val string) (HistoricAddmResultStatusEnum, bool) {
	enum, ok := mappingHistoricAddmResultStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// HistoricAddmResultHowCreatedEnum Enum with underlying type: string
type HistoricAddmResultHowCreatedEnum string

// Set of constants representing the allowable values for HistoricAddmResultHowCreatedEnum
const (
	HistoricAddmResultHowCreatedAuto   HistoricAddmResultHowCreatedEnum = "AUTO"
	HistoricAddmResultHowCreatedManual HistoricAddmResultHowCreatedEnum = "MANUAL"
)

var mappingHistoricAddmResultHowCreatedEnum = map[string]HistoricAddmResultHowCreatedEnum{
	"AUTO":   HistoricAddmResultHowCreatedAuto,
	"MANUAL": HistoricAddmResultHowCreatedManual,
}

var mappingHistoricAddmResultHowCreatedEnumLowerCase = map[string]HistoricAddmResultHowCreatedEnum{
	"auto":   HistoricAddmResultHowCreatedAuto,
	"manual": HistoricAddmResultHowCreatedManual,
}

// GetHistoricAddmResultHowCreatedEnumValues Enumerates the set of values for HistoricAddmResultHowCreatedEnum
func GetHistoricAddmResultHowCreatedEnumValues() []HistoricAddmResultHowCreatedEnum {
	values := make([]HistoricAddmResultHowCreatedEnum, 0)
	for _, v := range mappingHistoricAddmResultHowCreatedEnum {
		values = append(values, v)
	}
	return values
}

// GetHistoricAddmResultHowCreatedEnumStringValues Enumerates the set of values in String for HistoricAddmResultHowCreatedEnum
func GetHistoricAddmResultHowCreatedEnumStringValues() []string {
	return []string{
		"AUTO",
		"MANUAL",
	}
}

// GetMappingHistoricAddmResultHowCreatedEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHistoricAddmResultHowCreatedEnum(val string) (HistoricAddmResultHowCreatedEnum, bool) {
	enum, ok := mappingHistoricAddmResultHowCreatedEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
