// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"strings"
)

// SecurityPolicyEntryStateDeploymentStatusEnum Enum with underlying type: string
type SecurityPolicyEntryStateDeploymentStatusEnum string

// Set of constants representing the allowable values for SecurityPolicyEntryStateDeploymentStatusEnum
const (
	SecurityPolicyEntryStateDeploymentStatusCreated           SecurityPolicyEntryStateDeploymentStatusEnum = "CREATED"
	SecurityPolicyEntryStateDeploymentStatusModified          SecurityPolicyEntryStateDeploymentStatusEnum = "MODIFIED"
	SecurityPolicyEntryStateDeploymentStatusConflict          SecurityPolicyEntryStateDeploymentStatusEnum = "CONFLICT"
	SecurityPolicyEntryStateDeploymentStatusConnectivityIssue SecurityPolicyEntryStateDeploymentStatusEnum = "CONNECTIVITY_ISSUE"
	SecurityPolicyEntryStateDeploymentStatusUnsupportedSyntax SecurityPolicyEntryStateDeploymentStatusEnum = "UNSUPPORTED_SYNTAX"
	SecurityPolicyEntryStateDeploymentStatusUnknownError      SecurityPolicyEntryStateDeploymentStatusEnum = "UNKNOWN_ERROR"
	SecurityPolicyEntryStateDeploymentStatusUnauthorized      SecurityPolicyEntryStateDeploymentStatusEnum = "UNAUTHORIZED"
	SecurityPolicyEntryStateDeploymentStatusDeleted           SecurityPolicyEntryStateDeploymentStatusEnum = "DELETED"
)

var mappingSecurityPolicyEntryStateDeploymentStatusEnum = map[string]SecurityPolicyEntryStateDeploymentStatusEnum{
	"CREATED":            SecurityPolicyEntryStateDeploymentStatusCreated,
	"MODIFIED":           SecurityPolicyEntryStateDeploymentStatusModified,
	"CONFLICT":           SecurityPolicyEntryStateDeploymentStatusConflict,
	"CONNECTIVITY_ISSUE": SecurityPolicyEntryStateDeploymentStatusConnectivityIssue,
	"UNSUPPORTED_SYNTAX": SecurityPolicyEntryStateDeploymentStatusUnsupportedSyntax,
	"UNKNOWN_ERROR":      SecurityPolicyEntryStateDeploymentStatusUnknownError,
	"UNAUTHORIZED":       SecurityPolicyEntryStateDeploymentStatusUnauthorized,
	"DELETED":            SecurityPolicyEntryStateDeploymentStatusDeleted,
}

var mappingSecurityPolicyEntryStateDeploymentStatusEnumLowerCase = map[string]SecurityPolicyEntryStateDeploymentStatusEnum{
	"created":            SecurityPolicyEntryStateDeploymentStatusCreated,
	"modified":           SecurityPolicyEntryStateDeploymentStatusModified,
	"conflict":           SecurityPolicyEntryStateDeploymentStatusConflict,
	"connectivity_issue": SecurityPolicyEntryStateDeploymentStatusConnectivityIssue,
	"unsupported_syntax": SecurityPolicyEntryStateDeploymentStatusUnsupportedSyntax,
	"unknown_error":      SecurityPolicyEntryStateDeploymentStatusUnknownError,
	"unauthorized":       SecurityPolicyEntryStateDeploymentStatusUnauthorized,
	"deleted":            SecurityPolicyEntryStateDeploymentStatusDeleted,
}

// GetSecurityPolicyEntryStateDeploymentStatusEnumValues Enumerates the set of values for SecurityPolicyEntryStateDeploymentStatusEnum
func GetSecurityPolicyEntryStateDeploymentStatusEnumValues() []SecurityPolicyEntryStateDeploymentStatusEnum {
	values := make([]SecurityPolicyEntryStateDeploymentStatusEnum, 0)
	for _, v := range mappingSecurityPolicyEntryStateDeploymentStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetSecurityPolicyEntryStateDeploymentStatusEnumStringValues Enumerates the set of values in String for SecurityPolicyEntryStateDeploymentStatusEnum
func GetSecurityPolicyEntryStateDeploymentStatusEnumStringValues() []string {
	return []string{
		"CREATED",
		"MODIFIED",
		"CONFLICT",
		"CONNECTIVITY_ISSUE",
		"UNSUPPORTED_SYNTAX",
		"UNKNOWN_ERROR",
		"UNAUTHORIZED",
		"DELETED",
	}
}

// GetMappingSecurityPolicyEntryStateDeploymentStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecurityPolicyEntryStateDeploymentStatusEnum(val string) (SecurityPolicyEntryStateDeploymentStatusEnum, bool) {
	enum, ok := mappingSecurityPolicyEntryStateDeploymentStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
