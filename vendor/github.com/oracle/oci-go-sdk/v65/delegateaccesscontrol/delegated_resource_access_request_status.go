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

// DelegatedResourceAccessRequestStatusEnum Enum with underlying type: string
type DelegatedResourceAccessRequestStatusEnum string

// Set of constants representing the allowable values for DelegatedResourceAccessRequestStatusEnum
const (
	DelegatedResourceAccessRequestStatusCreated                   DelegatedResourceAccessRequestStatusEnum = "CREATED"
	DelegatedResourceAccessRequestStatusApprovalWaiting           DelegatedResourceAccessRequestStatusEnum = "APPROVAL_WAITING"
	DelegatedResourceAccessRequestStatusOperatorAssignmentWaiting DelegatedResourceAccessRequestStatusEnum = "OPERATOR_ASSIGNMENT_WAITING"
	DelegatedResourceAccessRequestStatusPreapproved               DelegatedResourceAccessRequestStatusEnum = "PREAPPROVED"
	DelegatedResourceAccessRequestStatusApproved                  DelegatedResourceAccessRequestStatusEnum = "APPROVED"
	DelegatedResourceAccessRequestStatusApprovedForFuture         DelegatedResourceAccessRequestStatusEnum = "APPROVED_FOR_FUTURE"
	DelegatedResourceAccessRequestStatusRejected                  DelegatedResourceAccessRequestStatusEnum = "REJECTED"
	DelegatedResourceAccessRequestStatusDeployed                  DelegatedResourceAccessRequestStatusEnum = "DEPLOYED"
	DelegatedResourceAccessRequestStatusDeployFailed              DelegatedResourceAccessRequestStatusEnum = "DEPLOY_FAILED"
	DelegatedResourceAccessRequestStatusUndeployed                DelegatedResourceAccessRequestStatusEnum = "UNDEPLOYED"
	DelegatedResourceAccessRequestStatusUndeployFailed            DelegatedResourceAccessRequestStatusEnum = "UNDEPLOY_FAILED"
	DelegatedResourceAccessRequestStatusCloseFailed               DelegatedResourceAccessRequestStatusEnum = "CLOSE_FAILED"
	DelegatedResourceAccessRequestStatusRevokeFailed              DelegatedResourceAccessRequestStatusEnum = "REVOKE_FAILED"
	DelegatedResourceAccessRequestStatusExpiryFailed              DelegatedResourceAccessRequestStatusEnum = "EXPIRY_FAILED"
	DelegatedResourceAccessRequestStatusRevoking                  DelegatedResourceAccessRequestStatusEnum = "REVOKING"
	DelegatedResourceAccessRequestStatusRevoked                   DelegatedResourceAccessRequestStatusEnum = "REVOKED"
	DelegatedResourceAccessRequestStatusExtending                 DelegatedResourceAccessRequestStatusEnum = "EXTENDING"
	DelegatedResourceAccessRequestStatusExtended                  DelegatedResourceAccessRequestStatusEnum = "EXTENDED"
	DelegatedResourceAccessRequestStatusExtensionRejected         DelegatedResourceAccessRequestStatusEnum = "EXTENSION_REJECTED"
	DelegatedResourceAccessRequestStatusExtensionFailed           DelegatedResourceAccessRequestStatusEnum = "EXTENSION_FAILED"
	DelegatedResourceAccessRequestStatusCompleting                DelegatedResourceAccessRequestStatusEnum = "COMPLETING"
	DelegatedResourceAccessRequestStatusCompleted                 DelegatedResourceAccessRequestStatusEnum = "COMPLETED"
	DelegatedResourceAccessRequestStatusExpired                   DelegatedResourceAccessRequestStatusEnum = "EXPIRED"
)

var mappingDelegatedResourceAccessRequestStatusEnum = map[string]DelegatedResourceAccessRequestStatusEnum{
	"CREATED":                     DelegatedResourceAccessRequestStatusCreated,
	"APPROVAL_WAITING":            DelegatedResourceAccessRequestStatusApprovalWaiting,
	"OPERATOR_ASSIGNMENT_WAITING": DelegatedResourceAccessRequestStatusOperatorAssignmentWaiting,
	"PREAPPROVED":                 DelegatedResourceAccessRequestStatusPreapproved,
	"APPROVED":                    DelegatedResourceAccessRequestStatusApproved,
	"APPROVED_FOR_FUTURE":         DelegatedResourceAccessRequestStatusApprovedForFuture,
	"REJECTED":                    DelegatedResourceAccessRequestStatusRejected,
	"DEPLOYED":                    DelegatedResourceAccessRequestStatusDeployed,
	"DEPLOY_FAILED":               DelegatedResourceAccessRequestStatusDeployFailed,
	"UNDEPLOYED":                  DelegatedResourceAccessRequestStatusUndeployed,
	"UNDEPLOY_FAILED":             DelegatedResourceAccessRequestStatusUndeployFailed,
	"CLOSE_FAILED":                DelegatedResourceAccessRequestStatusCloseFailed,
	"REVOKE_FAILED":               DelegatedResourceAccessRequestStatusRevokeFailed,
	"EXPIRY_FAILED":               DelegatedResourceAccessRequestStatusExpiryFailed,
	"REVOKING":                    DelegatedResourceAccessRequestStatusRevoking,
	"REVOKED":                     DelegatedResourceAccessRequestStatusRevoked,
	"EXTENDING":                   DelegatedResourceAccessRequestStatusExtending,
	"EXTENDED":                    DelegatedResourceAccessRequestStatusExtended,
	"EXTENSION_REJECTED":          DelegatedResourceAccessRequestStatusExtensionRejected,
	"EXTENSION_FAILED":            DelegatedResourceAccessRequestStatusExtensionFailed,
	"COMPLETING":                  DelegatedResourceAccessRequestStatusCompleting,
	"COMPLETED":                   DelegatedResourceAccessRequestStatusCompleted,
	"EXPIRED":                     DelegatedResourceAccessRequestStatusExpired,
}

var mappingDelegatedResourceAccessRequestStatusEnumLowerCase = map[string]DelegatedResourceAccessRequestStatusEnum{
	"created":                     DelegatedResourceAccessRequestStatusCreated,
	"approval_waiting":            DelegatedResourceAccessRequestStatusApprovalWaiting,
	"operator_assignment_waiting": DelegatedResourceAccessRequestStatusOperatorAssignmentWaiting,
	"preapproved":                 DelegatedResourceAccessRequestStatusPreapproved,
	"approved":                    DelegatedResourceAccessRequestStatusApproved,
	"approved_for_future":         DelegatedResourceAccessRequestStatusApprovedForFuture,
	"rejected":                    DelegatedResourceAccessRequestStatusRejected,
	"deployed":                    DelegatedResourceAccessRequestStatusDeployed,
	"deploy_failed":               DelegatedResourceAccessRequestStatusDeployFailed,
	"undeployed":                  DelegatedResourceAccessRequestStatusUndeployed,
	"undeploy_failed":             DelegatedResourceAccessRequestStatusUndeployFailed,
	"close_failed":                DelegatedResourceAccessRequestStatusCloseFailed,
	"revoke_failed":               DelegatedResourceAccessRequestStatusRevokeFailed,
	"expiry_failed":               DelegatedResourceAccessRequestStatusExpiryFailed,
	"revoking":                    DelegatedResourceAccessRequestStatusRevoking,
	"revoked":                     DelegatedResourceAccessRequestStatusRevoked,
	"extending":                   DelegatedResourceAccessRequestStatusExtending,
	"extended":                    DelegatedResourceAccessRequestStatusExtended,
	"extension_rejected":          DelegatedResourceAccessRequestStatusExtensionRejected,
	"extension_failed":            DelegatedResourceAccessRequestStatusExtensionFailed,
	"completing":                  DelegatedResourceAccessRequestStatusCompleting,
	"completed":                   DelegatedResourceAccessRequestStatusCompleted,
	"expired":                     DelegatedResourceAccessRequestStatusExpired,
}

// GetDelegatedResourceAccessRequestStatusEnumValues Enumerates the set of values for DelegatedResourceAccessRequestStatusEnum
func GetDelegatedResourceAccessRequestStatusEnumValues() []DelegatedResourceAccessRequestStatusEnum {
	values := make([]DelegatedResourceAccessRequestStatusEnum, 0)
	for _, v := range mappingDelegatedResourceAccessRequestStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDelegatedResourceAccessRequestStatusEnumStringValues Enumerates the set of values in String for DelegatedResourceAccessRequestStatusEnum
func GetDelegatedResourceAccessRequestStatusEnumStringValues() []string {
	return []string{
		"CREATED",
		"APPROVAL_WAITING",
		"OPERATOR_ASSIGNMENT_WAITING",
		"PREAPPROVED",
		"APPROVED",
		"APPROVED_FOR_FUTURE",
		"REJECTED",
		"DEPLOYED",
		"DEPLOY_FAILED",
		"UNDEPLOYED",
		"UNDEPLOY_FAILED",
		"CLOSE_FAILED",
		"REVOKE_FAILED",
		"EXPIRY_FAILED",
		"REVOKING",
		"REVOKED",
		"EXTENDING",
		"EXTENDED",
		"EXTENSION_REJECTED",
		"EXTENSION_FAILED",
		"COMPLETING",
		"COMPLETED",
		"EXPIRED",
	}
}

// GetMappingDelegatedResourceAccessRequestStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDelegatedResourceAccessRequestStatusEnum(val string) (DelegatedResourceAccessRequestStatusEnum, bool) {
	enum, ok := mappingDelegatedResourceAccessRequestStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
