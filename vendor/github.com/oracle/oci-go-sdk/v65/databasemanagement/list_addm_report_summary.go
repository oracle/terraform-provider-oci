// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// ListAddmReportSummary The object containing the ADDM report metadata.
type ListAddmReportSummary struct {

	// The ID number of the ADDM task.
	TaskId *int64 `mandatory:"true" json:"taskId"`

	// The creation date of the ADDM report.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The name of the ADDM report.
	TaskName *string `mandatory:"false" json:"taskName"`

	// The description of the ADDM report.
	Description *string `mandatory:"false" json:"description"`

	// The database user who owns the ADDM report.
	DbUser *string `mandatory:"false" json:"dbUser"`

	// The status of the ADDM report.
	Status ListAddmReportSummaryStatusEnum `mandatory:"false" json:"status,omitempty"`

	// A description of how the report was created.
	HowCreated ListAddmReportSummaryHowCreatedEnum `mandatory:"false" json:"howCreated,omitempty"`

	// The timestamp of the beginning AWR snapshot used in the ADDM report as defined by date-time RFC3339 format.
	StartSnapShotTime *common.SDKTime `mandatory:"false" json:"startSnapShotTime"`

	// The timestamp of the ending AWR snapshot used in the ADDM report as defined by date-time RFC3339 format.
	EndSnapshotTime *common.SDKTime `mandatory:"false" json:"endSnapshotTime"`

	// The ID number of the beginning AWR snapshot.
	BeginSnaphotId *int64 `mandatory:"false" json:"beginSnaphotId"`

	// The ID number of the ending AWR snapshot.
	EndSnapshotId *int64 `mandatory:"false" json:"endSnapshotId"`

	// The number of ADDM findings.
	Findings *int64 `mandatory:"false" json:"findings"`
}

func (m ListAddmReportSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ListAddmReportSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingListAddmReportSummaryStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetListAddmReportSummaryStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAddmReportSummaryHowCreatedEnum(string(m.HowCreated)); !ok && m.HowCreated != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for HowCreated: %s. Supported values are: %s.", m.HowCreated, strings.Join(GetListAddmReportSummaryHowCreatedEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAddmReportSummaryStatusEnum Enum with underlying type: string
type ListAddmReportSummaryStatusEnum string

// Set of constants representing the allowable values for ListAddmReportSummaryStatusEnum
const (
	ListAddmReportSummaryStatusInitial     ListAddmReportSummaryStatusEnum = "INITIAL"
	ListAddmReportSummaryStatusExecuting   ListAddmReportSummaryStatusEnum = "EXECUTING"
	ListAddmReportSummaryStatusInterrupted ListAddmReportSummaryStatusEnum = "INTERRUPTED"
	ListAddmReportSummaryStatusCompleted   ListAddmReportSummaryStatusEnum = "COMPLETED"
	ListAddmReportSummaryStatusError       ListAddmReportSummaryStatusEnum = "ERROR"
)

var mappingListAddmReportSummaryStatusEnum = map[string]ListAddmReportSummaryStatusEnum{
	"INITIAL":     ListAddmReportSummaryStatusInitial,
	"EXECUTING":   ListAddmReportSummaryStatusExecuting,
	"INTERRUPTED": ListAddmReportSummaryStatusInterrupted,
	"COMPLETED":   ListAddmReportSummaryStatusCompleted,
	"ERROR":       ListAddmReportSummaryStatusError,
}

var mappingListAddmReportSummaryStatusEnumLowerCase = map[string]ListAddmReportSummaryStatusEnum{
	"initial":     ListAddmReportSummaryStatusInitial,
	"executing":   ListAddmReportSummaryStatusExecuting,
	"interrupted": ListAddmReportSummaryStatusInterrupted,
	"completed":   ListAddmReportSummaryStatusCompleted,
	"error":       ListAddmReportSummaryStatusError,
}

// GetListAddmReportSummaryStatusEnumValues Enumerates the set of values for ListAddmReportSummaryStatusEnum
func GetListAddmReportSummaryStatusEnumValues() []ListAddmReportSummaryStatusEnum {
	values := make([]ListAddmReportSummaryStatusEnum, 0)
	for _, v := range mappingListAddmReportSummaryStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListAddmReportSummaryStatusEnumStringValues Enumerates the set of values in String for ListAddmReportSummaryStatusEnum
func GetListAddmReportSummaryStatusEnumStringValues() []string {
	return []string{
		"INITIAL",
		"EXECUTING",
		"INTERRUPTED",
		"COMPLETED",
		"ERROR",
	}
}

// GetMappingListAddmReportSummaryStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAddmReportSummaryStatusEnum(val string) (ListAddmReportSummaryStatusEnum, bool) {
	enum, ok := mappingListAddmReportSummaryStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAddmReportSummaryHowCreatedEnum Enum with underlying type: string
type ListAddmReportSummaryHowCreatedEnum string

// Set of constants representing the allowable values for ListAddmReportSummaryHowCreatedEnum
const (
	ListAddmReportSummaryHowCreatedAuto   ListAddmReportSummaryHowCreatedEnum = "AUTO"
	ListAddmReportSummaryHowCreatedManual ListAddmReportSummaryHowCreatedEnum = "MANUAL"
)

var mappingListAddmReportSummaryHowCreatedEnum = map[string]ListAddmReportSummaryHowCreatedEnum{
	"AUTO":   ListAddmReportSummaryHowCreatedAuto,
	"MANUAL": ListAddmReportSummaryHowCreatedManual,
}

var mappingListAddmReportSummaryHowCreatedEnumLowerCase = map[string]ListAddmReportSummaryHowCreatedEnum{
	"auto":   ListAddmReportSummaryHowCreatedAuto,
	"manual": ListAddmReportSummaryHowCreatedManual,
}

// GetListAddmReportSummaryHowCreatedEnumValues Enumerates the set of values for ListAddmReportSummaryHowCreatedEnum
func GetListAddmReportSummaryHowCreatedEnumValues() []ListAddmReportSummaryHowCreatedEnum {
	values := make([]ListAddmReportSummaryHowCreatedEnum, 0)
	for _, v := range mappingListAddmReportSummaryHowCreatedEnum {
		values = append(values, v)
	}
	return values
}

// GetListAddmReportSummaryHowCreatedEnumStringValues Enumerates the set of values in String for ListAddmReportSummaryHowCreatedEnum
func GetListAddmReportSummaryHowCreatedEnumStringValues() []string {
	return []string{
		"AUTO",
		"MANUAL",
	}
}

// GetMappingListAddmReportSummaryHowCreatedEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAddmReportSummaryHowCreatedEnum(val string) (ListAddmReportSummaryHowCreatedEnum, bool) {
	enum, ok := mappingListAddmReportSummaryHowCreatedEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
