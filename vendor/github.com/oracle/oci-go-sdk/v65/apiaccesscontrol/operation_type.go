// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle API Access Control
//
// This service is used to restrict the control plane service apis; so that everybody won't be
// able to access those apis.
// There are two main resouces defined as a part of this service
// 1. PrivilegedApiControl: This is created by the customer which defines which service apis are
//    controlled and who can access it.
// 2. PrivilegedApiRequest: This is a request object again created by the customer operators who           seek access to those privileged apis. After a request is obtained based on the                       PrivilegedAccessControl for which the api belongs to, either it can be approved so that the          requested person can execute the service apis or it will wait for the customer to approve it.
//

package apiaccesscontrol

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreatePrivilegedApiControl                OperationTypeEnum = "CREATE_PRIVILEGED_API_CONTROL"
	OperationTypeGetPrivilegedApiControl                   OperationTypeEnum = "GET_PRIVILEGED_API_CONTROL"
	OperationTypeListPrivilegedApiControl                  OperationTypeEnum = "LIST_PRIVILEGED_API_CONTROL"
	OperationTypeUpdatePrivilegedApiControl                OperationTypeEnum = "UPDATE_PRIVILEGED_API_CONTROL"
	OperationTypeDeletePrivilegedApiControl                OperationTypeEnum = "DELETE_PRIVILEGED_API_CONTROL"
	OperationTypeMovePrivilegedApiControl                  OperationTypeEnum = "MOVE_PRIVILEGED_API_CONTROL"
	OperationTypeCreatePrivilegedApiRequest                OperationTypeEnum = "CREATE_PRIVILEGED_API_REQUEST"
	OperationTypeGetPrivilegedApiRequest                   OperationTypeEnum = "GET_PRIVILEGED_API_REQUEST"
	OperationTypeListPrivilegedApiRequest                  OperationTypeEnum = "LIST_PRIVILEGED_API_REQUEST"
	OperationTypeApprovePrivilegedApiRequest               OperationTypeEnum = "APPROVE_PRIVILEGED_API_REQUEST"
	OperationTypeRejectPrivilegedApiRequest                OperationTypeEnum = "REJECT_PRIVILEGED_API_REQUEST"
	OperationTypeRevokePrivilegedApiRequest                OperationTypeEnum = "REVOKE_PRIVILEGED_API_REQUEST"
	OperationTypeClosePrivilegedApiRequest                 OperationTypeEnum = "CLOSE_PRIVILEGED_API_REQUEST"
	OperationTypeTimeoutPrivilegedApiRequest               OperationTypeEnum = "TIMEOUT_PRIVILEGED_API_REQUEST"
	OperationTypeCustomerApprovalCheckPrivilegedApiRequest OperationTypeEnum = "CUSTOMER_APPROVAL_CHECK_PRIVILEGED_API_REQUEST"
	OperationTypeCheckPrivilegedApiRequestStatus           OperationTypeEnum = "CHECK_PRIVILEGED_API_REQUEST_STATUS"
	OperationTypeGetPrivilegedApiWorkRequest               OperationTypeEnum = "GET_PRIVILEGED_API_WORK_REQUEST"
	OperationTypeListPrivilegedApiWorkRequest              OperationTypeEnum = "LIST_PRIVILEGED_API_WORK_REQUEST"
	OperationTypeListPrivilegedApiWorkRequestErrors        OperationTypeEnum = "LIST_PRIVILEGED_API_WORK_REQUEST_ERRORS"
	OperationTypeListPrivilegedApiWorkRequestLogs          OperationTypeEnum = "LIST_PRIVILEGED_API_WORK_REQUEST_LOGS"
	OperationTypeCancelPrivilegedApiWorkRequest            OperationTypeEnum = "CANCEL_PRIVILEGED_API_WORK_REQUEST"
	OperationTypeListApiMetadata                           OperationTypeEnum = "LIST_API_METADATA"
	OperationTypeListApiMetadataByEntityType               OperationTypeEnum = "LIST_API_METADATA_BY_ENTITY_TYPE"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_PRIVILEGED_API_CONTROL":                  OperationTypeCreatePrivilegedApiControl,
	"GET_PRIVILEGED_API_CONTROL":                     OperationTypeGetPrivilegedApiControl,
	"LIST_PRIVILEGED_API_CONTROL":                    OperationTypeListPrivilegedApiControl,
	"UPDATE_PRIVILEGED_API_CONTROL":                  OperationTypeUpdatePrivilegedApiControl,
	"DELETE_PRIVILEGED_API_CONTROL":                  OperationTypeDeletePrivilegedApiControl,
	"MOVE_PRIVILEGED_API_CONTROL":                    OperationTypeMovePrivilegedApiControl,
	"CREATE_PRIVILEGED_API_REQUEST":                  OperationTypeCreatePrivilegedApiRequest,
	"GET_PRIVILEGED_API_REQUEST":                     OperationTypeGetPrivilegedApiRequest,
	"LIST_PRIVILEGED_API_REQUEST":                    OperationTypeListPrivilegedApiRequest,
	"APPROVE_PRIVILEGED_API_REQUEST":                 OperationTypeApprovePrivilegedApiRequest,
	"REJECT_PRIVILEGED_API_REQUEST":                  OperationTypeRejectPrivilegedApiRequest,
	"REVOKE_PRIVILEGED_API_REQUEST":                  OperationTypeRevokePrivilegedApiRequest,
	"CLOSE_PRIVILEGED_API_REQUEST":                   OperationTypeClosePrivilegedApiRequest,
	"TIMEOUT_PRIVILEGED_API_REQUEST":                 OperationTypeTimeoutPrivilegedApiRequest,
	"CUSTOMER_APPROVAL_CHECK_PRIVILEGED_API_REQUEST": OperationTypeCustomerApprovalCheckPrivilegedApiRequest,
	"CHECK_PRIVILEGED_API_REQUEST_STATUS":            OperationTypeCheckPrivilegedApiRequestStatus,
	"GET_PRIVILEGED_API_WORK_REQUEST":                OperationTypeGetPrivilegedApiWorkRequest,
	"LIST_PRIVILEGED_API_WORK_REQUEST":               OperationTypeListPrivilegedApiWorkRequest,
	"LIST_PRIVILEGED_API_WORK_REQUEST_ERRORS":        OperationTypeListPrivilegedApiWorkRequestErrors,
	"LIST_PRIVILEGED_API_WORK_REQUEST_LOGS":          OperationTypeListPrivilegedApiWorkRequestLogs,
	"CANCEL_PRIVILEGED_API_WORK_REQUEST":             OperationTypeCancelPrivilegedApiWorkRequest,
	"LIST_API_METADATA":                              OperationTypeListApiMetadata,
	"LIST_API_METADATA_BY_ENTITY_TYPE":               OperationTypeListApiMetadataByEntityType,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_privileged_api_control":                  OperationTypeCreatePrivilegedApiControl,
	"get_privileged_api_control":                     OperationTypeGetPrivilegedApiControl,
	"list_privileged_api_control":                    OperationTypeListPrivilegedApiControl,
	"update_privileged_api_control":                  OperationTypeUpdatePrivilegedApiControl,
	"delete_privileged_api_control":                  OperationTypeDeletePrivilegedApiControl,
	"move_privileged_api_control":                    OperationTypeMovePrivilegedApiControl,
	"create_privileged_api_request":                  OperationTypeCreatePrivilegedApiRequest,
	"get_privileged_api_request":                     OperationTypeGetPrivilegedApiRequest,
	"list_privileged_api_request":                    OperationTypeListPrivilegedApiRequest,
	"approve_privileged_api_request":                 OperationTypeApprovePrivilegedApiRequest,
	"reject_privileged_api_request":                  OperationTypeRejectPrivilegedApiRequest,
	"revoke_privileged_api_request":                  OperationTypeRevokePrivilegedApiRequest,
	"close_privileged_api_request":                   OperationTypeClosePrivilegedApiRequest,
	"timeout_privileged_api_request":                 OperationTypeTimeoutPrivilegedApiRequest,
	"customer_approval_check_privileged_api_request": OperationTypeCustomerApprovalCheckPrivilegedApiRequest,
	"check_privileged_api_request_status":            OperationTypeCheckPrivilegedApiRequestStatus,
	"get_privileged_api_work_request":                OperationTypeGetPrivilegedApiWorkRequest,
	"list_privileged_api_work_request":               OperationTypeListPrivilegedApiWorkRequest,
	"list_privileged_api_work_request_errors":        OperationTypeListPrivilegedApiWorkRequestErrors,
	"list_privileged_api_work_request_logs":          OperationTypeListPrivilegedApiWorkRequestLogs,
	"cancel_privileged_api_work_request":             OperationTypeCancelPrivilegedApiWorkRequest,
	"list_api_metadata":                              OperationTypeListApiMetadata,
	"list_api_metadata_by_entity_type":               OperationTypeListApiMetadataByEntityType,
}

// GetOperationTypeEnumValues Enumerates the set of values for OperationTypeEnum
func GetOperationTypeEnumValues() []OperationTypeEnum {
	values := make([]OperationTypeEnum, 0)
	for _, v := range mappingOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationTypeEnumStringValues Enumerates the set of values in String for OperationTypeEnum
func GetOperationTypeEnumStringValues() []string {
	return []string{
		"CREATE_PRIVILEGED_API_CONTROL",
		"GET_PRIVILEGED_API_CONTROL",
		"LIST_PRIVILEGED_API_CONTROL",
		"UPDATE_PRIVILEGED_API_CONTROL",
		"DELETE_PRIVILEGED_API_CONTROL",
		"MOVE_PRIVILEGED_API_CONTROL",
		"CREATE_PRIVILEGED_API_REQUEST",
		"GET_PRIVILEGED_API_REQUEST",
		"LIST_PRIVILEGED_API_REQUEST",
		"APPROVE_PRIVILEGED_API_REQUEST",
		"REJECT_PRIVILEGED_API_REQUEST",
		"REVOKE_PRIVILEGED_API_REQUEST",
		"CLOSE_PRIVILEGED_API_REQUEST",
		"TIMEOUT_PRIVILEGED_API_REQUEST",
		"CUSTOMER_APPROVAL_CHECK_PRIVILEGED_API_REQUEST",
		"CHECK_PRIVILEGED_API_REQUEST_STATUS",
		"GET_PRIVILEGED_API_WORK_REQUEST",
		"LIST_PRIVILEGED_API_WORK_REQUEST",
		"LIST_PRIVILEGED_API_WORK_REQUEST_ERRORS",
		"LIST_PRIVILEGED_API_WORK_REQUEST_LOGS",
		"CANCEL_PRIVILEGED_API_WORK_REQUEST",
		"LIST_API_METADATA",
		"LIST_API_METADATA_BY_ENTITY_TYPE",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
