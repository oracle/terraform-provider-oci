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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// DatabaseParameterUpdateStatus The result of database parameter update.
type DatabaseParameterUpdateStatus struct {

	// The status of the parameter update.
	Status DatabaseParameterUpdateStatusStatusEnum `mandatory:"false" json:"status,omitempty"`

	// An error code that defines the failure or `null` if the parameter
	// was updated successfully.
	ErrorCode *string `mandatory:"false" json:"errorCode"`

	// The error message indicating the reason for failure or `null` if
	// the parameter was updated successfully.
	ErrorMessage *string `mandatory:"false" json:"errorMessage"`
}

func (m DatabaseParameterUpdateStatus) String() string {
	return common.PointerString(m)
}

// DatabaseParameterUpdateStatusStatusEnum Enum with underlying type: string
type DatabaseParameterUpdateStatusStatusEnum string

// Set of constants representing the allowable values for DatabaseParameterUpdateStatusStatusEnum
const (
	DatabaseParameterUpdateStatusStatusSucceeded DatabaseParameterUpdateStatusStatusEnum = "SUCCEEDED"
	DatabaseParameterUpdateStatusStatusFailed    DatabaseParameterUpdateStatusStatusEnum = "FAILED"
)

var mappingDatabaseParameterUpdateStatusStatus = map[string]DatabaseParameterUpdateStatusStatusEnum{
	"SUCCEEDED": DatabaseParameterUpdateStatusStatusSucceeded,
	"FAILED":    DatabaseParameterUpdateStatusStatusFailed,
}

// GetDatabaseParameterUpdateStatusStatusEnumValues Enumerates the set of values for DatabaseParameterUpdateStatusStatusEnum
func GetDatabaseParameterUpdateStatusStatusEnumValues() []DatabaseParameterUpdateStatusStatusEnum {
	values := make([]DatabaseParameterUpdateStatusStatusEnum, 0)
	for _, v := range mappingDatabaseParameterUpdateStatusStatus {
		values = append(values, v)
	}
	return values
}
