// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Delegate Access Control API
//
// Oracle Delegate Access Control allows ExaCC and ExaCS customers to delegate management of their Exadata resources operators outside their tenancies.
// With Delegate Access Control, Support Providers can deliver managed services using comprehensive and robust tooling built on the OCI platform.
// Customers maintain control over who has access to the delegated resources in their tenancy and what actions can be taken.
// Enterprises managing resources across multiple tenants can use Delegate Access Control to streamline management tasks.
// Using logging service, customers can view a near real-time audit report of all actions performed by a Service Provider operator.
//

package delegateaccesscontrol

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateDelegationControl               OperationTypeEnum = "CREATE_DELEGATION_CONTROL"
	OperationTypeUpdateDelegationControl               OperationTypeEnum = "UPDATE_DELEGATION_CONTROL"
	OperationTypeDeleteDelegationControl               OperationTypeEnum = "DELETE_DELEGATION_CONTROL"
	OperationTypeMoveDelegationControl                 OperationTypeEnum = "MOVE_DELEGATION_CONTROL"
	OperationTypeCreateDelegatedResourceAccessRequest  OperationTypeEnum = "CREATE_DELEGATED_RESOURCE_ACCESS_REQUEST"
	OperationTypeApproveDelegatedResourceAccessRequest OperationTypeEnum = "APPROVE_DELEGATED_RESOURCE_ACCESS_REQUEST"
	OperationTypeRejectDelegatedResourceAccessRequest  OperationTypeEnum = "REJECT_DELEGATED_RESOURCE_ACCESS_REQUEST"
	OperationTypeRevokeDelegatedResourceAccessRequest  OperationTypeEnum = "REVOKE_DELEGATED_RESOURCE_ACCESS_REQUEST"
	OperationTypeCreateDelegationSubscription          OperationTypeEnum = "CREATE_DELEGATION_SUBSCRIPTION"
	OperationTypeUpdateDelegationSubscription          OperationTypeEnum = "UPDATE_DELEGATION_SUBSCRIPTION"
	OperationTypeDeleteDelegationSubscription          OperationTypeEnum = "DELETE_DELEGATION_SUBSCRIPTION"
	OperationTypeMoveDelegationSubscription            OperationTypeEnum = "MOVE_DELEGATION_SUBSCRIPTION"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_DELEGATION_CONTROL":                 OperationTypeCreateDelegationControl,
	"UPDATE_DELEGATION_CONTROL":                 OperationTypeUpdateDelegationControl,
	"DELETE_DELEGATION_CONTROL":                 OperationTypeDeleteDelegationControl,
	"MOVE_DELEGATION_CONTROL":                   OperationTypeMoveDelegationControl,
	"CREATE_DELEGATED_RESOURCE_ACCESS_REQUEST":  OperationTypeCreateDelegatedResourceAccessRequest,
	"APPROVE_DELEGATED_RESOURCE_ACCESS_REQUEST": OperationTypeApproveDelegatedResourceAccessRequest,
	"REJECT_DELEGATED_RESOURCE_ACCESS_REQUEST":  OperationTypeRejectDelegatedResourceAccessRequest,
	"REVOKE_DELEGATED_RESOURCE_ACCESS_REQUEST":  OperationTypeRevokeDelegatedResourceAccessRequest,
	"CREATE_DELEGATION_SUBSCRIPTION":            OperationTypeCreateDelegationSubscription,
	"UPDATE_DELEGATION_SUBSCRIPTION":            OperationTypeUpdateDelegationSubscription,
	"DELETE_DELEGATION_SUBSCRIPTION":            OperationTypeDeleteDelegationSubscription,
	"MOVE_DELEGATION_SUBSCRIPTION":              OperationTypeMoveDelegationSubscription,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_delegation_control":                 OperationTypeCreateDelegationControl,
	"update_delegation_control":                 OperationTypeUpdateDelegationControl,
	"delete_delegation_control":                 OperationTypeDeleteDelegationControl,
	"move_delegation_control":                   OperationTypeMoveDelegationControl,
	"create_delegated_resource_access_request":  OperationTypeCreateDelegatedResourceAccessRequest,
	"approve_delegated_resource_access_request": OperationTypeApproveDelegatedResourceAccessRequest,
	"reject_delegated_resource_access_request":  OperationTypeRejectDelegatedResourceAccessRequest,
	"revoke_delegated_resource_access_request":  OperationTypeRevokeDelegatedResourceAccessRequest,
	"create_delegation_subscription":            OperationTypeCreateDelegationSubscription,
	"update_delegation_subscription":            OperationTypeUpdateDelegationSubscription,
	"delete_delegation_subscription":            OperationTypeDeleteDelegationSubscription,
	"move_delegation_subscription":              OperationTypeMoveDelegationSubscription,
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
		"CREATE_DELEGATION_CONTROL",
		"UPDATE_DELEGATION_CONTROL",
		"DELETE_DELEGATION_CONTROL",
		"MOVE_DELEGATION_CONTROL",
		"CREATE_DELEGATED_RESOURCE_ACCESS_REQUEST",
		"APPROVE_DELEGATED_RESOURCE_ACCESS_REQUEST",
		"REJECT_DELEGATED_RESOURCE_ACCESS_REQUEST",
		"REVOKE_DELEGATED_RESOURCE_ACCESS_REQUEST",
		"CREATE_DELEGATION_SUBSCRIPTION",
		"UPDATE_DELEGATION_SUBSCRIPTION",
		"DELETE_DELEGATION_SUBSCRIPTION",
		"MOVE_DELEGATION_SUBSCRIPTION",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
