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

// WorkRequestOperationTypeEnum Enum with underlying type: string
type WorkRequestOperationTypeEnum string

// Set of constants representing the allowable values for WorkRequestOperationTypeEnum
const (
	WorkRequestOperationTypeCreateDbManagementPrivateEndpoint WorkRequestOperationTypeEnum = "CREATE_DB_MANAGEMENT_PRIVATE_ENDPOINT"
	WorkRequestOperationTypeDeleteDbManagementPrivateEndpoint WorkRequestOperationTypeEnum = "DELETE_DB_MANAGEMENT_PRIVATE_ENDPOINT"
)

var mappingWorkRequestOperationType = map[string]WorkRequestOperationTypeEnum{
	"CREATE_DB_MANAGEMENT_PRIVATE_ENDPOINT": WorkRequestOperationTypeCreateDbManagementPrivateEndpoint,
	"DELETE_DB_MANAGEMENT_PRIVATE_ENDPOINT": WorkRequestOperationTypeDeleteDbManagementPrivateEndpoint,
}

// GetWorkRequestOperationTypeEnumValues Enumerates the set of values for WorkRequestOperationTypeEnum
func GetWorkRequestOperationTypeEnumValues() []WorkRequestOperationTypeEnum {
	values := make([]WorkRequestOperationTypeEnum, 0)
	for _, v := range mappingWorkRequestOperationType {
		values = append(values, v)
	}
	return values
}
