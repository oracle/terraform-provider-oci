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

// DbServerPatchingDetails The scheduling details for the quarterly maintenance window. Patching and system updates take place during the maintenance window.
type DbServerPatchingDetails struct {

	// Estimated time, in minutes, to patch one database server.
	EstimatedPatchDuration *int `mandatory:"false" json:"estimatedPatchDuration"`

	// The status of the patching operation.
	PatchingStatus DbServerPatchingDetailsPatchingStatusEnum `mandatory:"false" json:"patchingStatus,omitempty"`

	// The time when the patching operation started.
	TimePatchingStarted *common.SDKTime `mandatory:"false" json:"timePatchingStarted"`

	// The time when the patching operation ended.
	TimePatchingEnded *common.SDKTime `mandatory:"false" json:"timePatchingEnded"`
}

func (m DbServerPatchingDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbServerPatchingDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDbServerPatchingDetailsPatchingStatusEnum(string(m.PatchingStatus)); !ok && m.PatchingStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PatchingStatus: %s. Supported values are: %s.", m.PatchingStatus, strings.Join(GetDbServerPatchingDetailsPatchingStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DbServerPatchingDetailsPatchingStatusEnum Enum with underlying type: string
type DbServerPatchingDetailsPatchingStatusEnum string

// Set of constants representing the allowable values for DbServerPatchingDetailsPatchingStatusEnum
const (
	DbServerPatchingDetailsPatchingStatusScheduled             DbServerPatchingDetailsPatchingStatusEnum = "SCHEDULED"
	DbServerPatchingDetailsPatchingStatusMaintenanceInProgress DbServerPatchingDetailsPatchingStatusEnum = "MAINTENANCE_IN_PROGRESS"
	DbServerPatchingDetailsPatchingStatusFailed                DbServerPatchingDetailsPatchingStatusEnum = "FAILED"
	DbServerPatchingDetailsPatchingStatusComplete              DbServerPatchingDetailsPatchingStatusEnum = "COMPLETE"
)

var mappingDbServerPatchingDetailsPatchingStatusEnum = map[string]DbServerPatchingDetailsPatchingStatusEnum{
	"SCHEDULED":               DbServerPatchingDetailsPatchingStatusScheduled,
	"MAINTENANCE_IN_PROGRESS": DbServerPatchingDetailsPatchingStatusMaintenanceInProgress,
	"FAILED":                  DbServerPatchingDetailsPatchingStatusFailed,
	"COMPLETE":                DbServerPatchingDetailsPatchingStatusComplete,
}

var mappingDbServerPatchingDetailsPatchingStatusEnumLowerCase = map[string]DbServerPatchingDetailsPatchingStatusEnum{
	"scheduled":               DbServerPatchingDetailsPatchingStatusScheduled,
	"maintenance_in_progress": DbServerPatchingDetailsPatchingStatusMaintenanceInProgress,
	"failed":                  DbServerPatchingDetailsPatchingStatusFailed,
	"complete":                DbServerPatchingDetailsPatchingStatusComplete,
}

// GetDbServerPatchingDetailsPatchingStatusEnumValues Enumerates the set of values for DbServerPatchingDetailsPatchingStatusEnum
func GetDbServerPatchingDetailsPatchingStatusEnumValues() []DbServerPatchingDetailsPatchingStatusEnum {
	values := make([]DbServerPatchingDetailsPatchingStatusEnum, 0)
	for _, v := range mappingDbServerPatchingDetailsPatchingStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDbServerPatchingDetailsPatchingStatusEnumStringValues Enumerates the set of values in String for DbServerPatchingDetailsPatchingStatusEnum
func GetDbServerPatchingDetailsPatchingStatusEnumStringValues() []string {
	return []string{
		"SCHEDULED",
		"MAINTENANCE_IN_PROGRESS",
		"FAILED",
		"COMPLETE",
	}
}

// GetMappingDbServerPatchingDetailsPatchingStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbServerPatchingDetailsPatchingStatusEnum(val string) (DbServerPatchingDetailsPatchingStatusEnum, bool) {
	enum, ok := mappingDbServerPatchingDetailsPatchingStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
