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

// DelegatedResourceAccessRequestAuditTypeEnum Enum with underlying type: string
type DelegatedResourceAccessRequestAuditTypeEnum string

// Set of constants representing the allowable values for DelegatedResourceAccessRequestAuditTypeEnum
const (
	DelegatedResourceAccessRequestAuditTypeCommandAudit          DelegatedResourceAccessRequestAuditTypeEnum = "COMMAND_AUDIT"
	DelegatedResourceAccessRequestAuditTypeCommandKeystrokeAudit DelegatedResourceAccessRequestAuditTypeEnum = "COMMAND_KEYSTROKE_AUDIT"
)

var mappingDelegatedResourceAccessRequestAuditTypeEnum = map[string]DelegatedResourceAccessRequestAuditTypeEnum{
	"COMMAND_AUDIT":           DelegatedResourceAccessRequestAuditTypeCommandAudit,
	"COMMAND_KEYSTROKE_AUDIT": DelegatedResourceAccessRequestAuditTypeCommandKeystrokeAudit,
}

var mappingDelegatedResourceAccessRequestAuditTypeEnumLowerCase = map[string]DelegatedResourceAccessRequestAuditTypeEnum{
	"command_audit":           DelegatedResourceAccessRequestAuditTypeCommandAudit,
	"command_keystroke_audit": DelegatedResourceAccessRequestAuditTypeCommandKeystrokeAudit,
}

// GetDelegatedResourceAccessRequestAuditTypeEnumValues Enumerates the set of values for DelegatedResourceAccessRequestAuditTypeEnum
func GetDelegatedResourceAccessRequestAuditTypeEnumValues() []DelegatedResourceAccessRequestAuditTypeEnum {
	values := make([]DelegatedResourceAccessRequestAuditTypeEnum, 0)
	for _, v := range mappingDelegatedResourceAccessRequestAuditTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDelegatedResourceAccessRequestAuditTypeEnumStringValues Enumerates the set of values in String for DelegatedResourceAccessRequestAuditTypeEnum
func GetDelegatedResourceAccessRequestAuditTypeEnumStringValues() []string {
	return []string{
		"COMMAND_AUDIT",
		"COMMAND_KEYSTROKE_AUDIT",
	}
}

// GetMappingDelegatedResourceAccessRequestAuditTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDelegatedResourceAccessRequestAuditTypeEnum(val string) (DelegatedResourceAccessRequestAuditTypeEnum, bool) {
	enum, ok := mappingDelegatedResourceAccessRequestAuditTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
