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

// AuditReportStatusEnum Enum with underlying type: string
type AuditReportStatusEnum string

// Set of constants representing the allowable values for AuditReportStatusEnum
const (
	AuditReportStatusNotavailable AuditReportStatusEnum = "NOTAVAILABLE"
	AuditReportStatusAvailable    AuditReportStatusEnum = "AVAILABLE"
	AuditReportStatusExpired      AuditReportStatusEnum = "EXPIRED"
	AuditReportStatusFailed       AuditReportStatusEnum = "FAILED"
)

var mappingAuditReportStatusEnum = map[string]AuditReportStatusEnum{
	"NOTAVAILABLE": AuditReportStatusNotavailable,
	"AVAILABLE":    AuditReportStatusAvailable,
	"EXPIRED":      AuditReportStatusExpired,
	"FAILED":       AuditReportStatusFailed,
}

var mappingAuditReportStatusEnumLowerCase = map[string]AuditReportStatusEnum{
	"notavailable": AuditReportStatusNotavailable,
	"available":    AuditReportStatusAvailable,
	"expired":      AuditReportStatusExpired,
	"failed":       AuditReportStatusFailed,
}

// GetAuditReportStatusEnumValues Enumerates the set of values for AuditReportStatusEnum
func GetAuditReportStatusEnumValues() []AuditReportStatusEnum {
	values := make([]AuditReportStatusEnum, 0)
	for _, v := range mappingAuditReportStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetAuditReportStatusEnumStringValues Enumerates the set of values in String for AuditReportStatusEnum
func GetAuditReportStatusEnumStringValues() []string {
	return []string{
		"NOTAVAILABLE",
		"AVAILABLE",
		"EXPIRED",
		"FAILED",
	}
}

// GetMappingAuditReportStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAuditReportStatusEnum(val string) (AuditReportStatusEnum, bool) {
	enum, ok := mappingAuditReportStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
