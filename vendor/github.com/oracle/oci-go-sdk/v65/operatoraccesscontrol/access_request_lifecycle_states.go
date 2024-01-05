// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OperatorAccessControl API
//
// Operator Access Control enables you to control the time duration and the actions an Oracle operator can perform on your Exadata Cloud@Customer infrastructure.
// Using logging service, you can view a near real-time audit report of all actions performed by an Oracle operator.
// Use the table of contents and search tool to explore the OperatorAccessControl API.
//

package operatoraccesscontrol

import (
	"strings"
)

// AccessRequestLifecycleStatesEnum Enum with underlying type: string
type AccessRequestLifecycleStatesEnum string

// Set of constants representing the allowable values for AccessRequestLifecycleStatesEnum
const (
	AccessRequestLifecycleStatesCreated           AccessRequestLifecycleStatesEnum = "CREATED"
	AccessRequestLifecycleStatesApprovalwaiting   AccessRequestLifecycleStatesEnum = "APPROVALWAITING"
	AccessRequestLifecycleStatesPreapproved       AccessRequestLifecycleStatesEnum = "PREAPPROVED"
	AccessRequestLifecycleStatesApproved          AccessRequestLifecycleStatesEnum = "APPROVED"
	AccessRequestLifecycleStatesMoreinfo          AccessRequestLifecycleStatesEnum = "MOREINFO"
	AccessRequestLifecycleStatesRejected          AccessRequestLifecycleStatesEnum = "REJECTED"
	AccessRequestLifecycleStatesDeployed          AccessRequestLifecycleStatesEnum = "DEPLOYED"
	AccessRequestLifecycleStatesDeployfailed      AccessRequestLifecycleStatesEnum = "DEPLOYFAILED"
	AccessRequestLifecycleStatesUndeployed        AccessRequestLifecycleStatesEnum = "UNDEPLOYED"
	AccessRequestLifecycleStatesUndeployfailed    AccessRequestLifecycleStatesEnum = "UNDEPLOYFAILED"
	AccessRequestLifecycleStatesClosefailed       AccessRequestLifecycleStatesEnum = "CLOSEFAILED"
	AccessRequestLifecycleStatesRevokefailed      AccessRequestLifecycleStatesEnum = "REVOKEFAILED"
	AccessRequestLifecycleStatesExpiryfailed      AccessRequestLifecycleStatesEnum = "EXPIRYFAILED"
	AccessRequestLifecycleStatesRevoking          AccessRequestLifecycleStatesEnum = "REVOKING"
	AccessRequestLifecycleStatesRevoked           AccessRequestLifecycleStatesEnum = "REVOKED"
	AccessRequestLifecycleStatesExtending         AccessRequestLifecycleStatesEnum = "EXTENDING"
	AccessRequestLifecycleStatesExtended          AccessRequestLifecycleStatesEnum = "EXTENDED"
	AccessRequestLifecycleStatesExtensionrejected AccessRequestLifecycleStatesEnum = "EXTENSIONREJECTED"
	AccessRequestLifecycleStatesCompleting        AccessRequestLifecycleStatesEnum = "COMPLETING"
	AccessRequestLifecycleStatesCompleted         AccessRequestLifecycleStatesEnum = "COMPLETED"
	AccessRequestLifecycleStatesExpired           AccessRequestLifecycleStatesEnum = "EXPIRED"
	AccessRequestLifecycleStatesApprovedforfuture AccessRequestLifecycleStatesEnum = "APPROVEDFORFUTURE"
	AccessRequestLifecycleStatesInreview          AccessRequestLifecycleStatesEnum = "INREVIEW"
)

var mappingAccessRequestLifecycleStatesEnum = map[string]AccessRequestLifecycleStatesEnum{
	"CREATED":           AccessRequestLifecycleStatesCreated,
	"APPROVALWAITING":   AccessRequestLifecycleStatesApprovalwaiting,
	"PREAPPROVED":       AccessRequestLifecycleStatesPreapproved,
	"APPROVED":          AccessRequestLifecycleStatesApproved,
	"MOREINFO":          AccessRequestLifecycleStatesMoreinfo,
	"REJECTED":          AccessRequestLifecycleStatesRejected,
	"DEPLOYED":          AccessRequestLifecycleStatesDeployed,
	"DEPLOYFAILED":      AccessRequestLifecycleStatesDeployfailed,
	"UNDEPLOYED":        AccessRequestLifecycleStatesUndeployed,
	"UNDEPLOYFAILED":    AccessRequestLifecycleStatesUndeployfailed,
	"CLOSEFAILED":       AccessRequestLifecycleStatesClosefailed,
	"REVOKEFAILED":      AccessRequestLifecycleStatesRevokefailed,
	"EXPIRYFAILED":      AccessRequestLifecycleStatesExpiryfailed,
	"REVOKING":          AccessRequestLifecycleStatesRevoking,
	"REVOKED":           AccessRequestLifecycleStatesRevoked,
	"EXTENDING":         AccessRequestLifecycleStatesExtending,
	"EXTENDED":          AccessRequestLifecycleStatesExtended,
	"EXTENSIONREJECTED": AccessRequestLifecycleStatesExtensionrejected,
	"COMPLETING":        AccessRequestLifecycleStatesCompleting,
	"COMPLETED":         AccessRequestLifecycleStatesCompleted,
	"EXPIRED":           AccessRequestLifecycleStatesExpired,
	"APPROVEDFORFUTURE": AccessRequestLifecycleStatesApprovedforfuture,
	"INREVIEW":          AccessRequestLifecycleStatesInreview,
}

var mappingAccessRequestLifecycleStatesEnumLowerCase = map[string]AccessRequestLifecycleStatesEnum{
	"created":           AccessRequestLifecycleStatesCreated,
	"approvalwaiting":   AccessRequestLifecycleStatesApprovalwaiting,
	"preapproved":       AccessRequestLifecycleStatesPreapproved,
	"approved":          AccessRequestLifecycleStatesApproved,
	"moreinfo":          AccessRequestLifecycleStatesMoreinfo,
	"rejected":          AccessRequestLifecycleStatesRejected,
	"deployed":          AccessRequestLifecycleStatesDeployed,
	"deployfailed":      AccessRequestLifecycleStatesDeployfailed,
	"undeployed":        AccessRequestLifecycleStatesUndeployed,
	"undeployfailed":    AccessRequestLifecycleStatesUndeployfailed,
	"closefailed":       AccessRequestLifecycleStatesClosefailed,
	"revokefailed":      AccessRequestLifecycleStatesRevokefailed,
	"expiryfailed":      AccessRequestLifecycleStatesExpiryfailed,
	"revoking":          AccessRequestLifecycleStatesRevoking,
	"revoked":           AccessRequestLifecycleStatesRevoked,
	"extending":         AccessRequestLifecycleStatesExtending,
	"extended":          AccessRequestLifecycleStatesExtended,
	"extensionrejected": AccessRequestLifecycleStatesExtensionrejected,
	"completing":        AccessRequestLifecycleStatesCompleting,
	"completed":         AccessRequestLifecycleStatesCompleted,
	"expired":           AccessRequestLifecycleStatesExpired,
	"approvedforfuture": AccessRequestLifecycleStatesApprovedforfuture,
	"inreview":          AccessRequestLifecycleStatesInreview,
}

// GetAccessRequestLifecycleStatesEnumValues Enumerates the set of values for AccessRequestLifecycleStatesEnum
func GetAccessRequestLifecycleStatesEnumValues() []AccessRequestLifecycleStatesEnum {
	values := make([]AccessRequestLifecycleStatesEnum, 0)
	for _, v := range mappingAccessRequestLifecycleStatesEnum {
		values = append(values, v)
	}
	return values
}

// GetAccessRequestLifecycleStatesEnumStringValues Enumerates the set of values in String for AccessRequestLifecycleStatesEnum
func GetAccessRequestLifecycleStatesEnumStringValues() []string {
	return []string{
		"CREATED",
		"APPROVALWAITING",
		"PREAPPROVED",
		"APPROVED",
		"MOREINFO",
		"REJECTED",
		"DEPLOYED",
		"DEPLOYFAILED",
		"UNDEPLOYED",
		"UNDEPLOYFAILED",
		"CLOSEFAILED",
		"REVOKEFAILED",
		"EXPIRYFAILED",
		"REVOKING",
		"REVOKED",
		"EXTENDING",
		"EXTENDED",
		"EXTENSIONREJECTED",
		"COMPLETING",
		"COMPLETED",
		"EXPIRED",
		"APPROVEDFORFUTURE",
		"INREVIEW",
	}
}

// GetMappingAccessRequestLifecycleStatesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAccessRequestLifecycleStatesEnum(val string) (AccessRequestLifecycleStatesEnum, bool) {
	enum, ok := mappingAccessRequestLifecycleStatesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
