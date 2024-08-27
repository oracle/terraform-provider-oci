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

// DelegatedResourceAccessRequesterTypeEnum Enum with underlying type: string
type DelegatedResourceAccessRequesterTypeEnum string

// Set of constants representing the allowable values for DelegatedResourceAccessRequesterTypeEnum
const (
	DelegatedResourceAccessRequesterTypeOperator DelegatedResourceAccessRequesterTypeEnum = "OPERATOR"
	DelegatedResourceAccessRequesterTypeCustomer DelegatedResourceAccessRequesterTypeEnum = "CUSTOMER"
	DelegatedResourceAccessRequesterTypeSystem   DelegatedResourceAccessRequesterTypeEnum = "SYSTEM"
)

var mappingDelegatedResourceAccessRequesterTypeEnum = map[string]DelegatedResourceAccessRequesterTypeEnum{
	"OPERATOR": DelegatedResourceAccessRequesterTypeOperator,
	"CUSTOMER": DelegatedResourceAccessRequesterTypeCustomer,
	"SYSTEM":   DelegatedResourceAccessRequesterTypeSystem,
}

var mappingDelegatedResourceAccessRequesterTypeEnumLowerCase = map[string]DelegatedResourceAccessRequesterTypeEnum{
	"operator": DelegatedResourceAccessRequesterTypeOperator,
	"customer": DelegatedResourceAccessRequesterTypeCustomer,
	"system":   DelegatedResourceAccessRequesterTypeSystem,
}

// GetDelegatedResourceAccessRequesterTypeEnumValues Enumerates the set of values for DelegatedResourceAccessRequesterTypeEnum
func GetDelegatedResourceAccessRequesterTypeEnumValues() []DelegatedResourceAccessRequesterTypeEnum {
	values := make([]DelegatedResourceAccessRequesterTypeEnum, 0)
	for _, v := range mappingDelegatedResourceAccessRequesterTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDelegatedResourceAccessRequesterTypeEnumStringValues Enumerates the set of values in String for DelegatedResourceAccessRequesterTypeEnum
func GetDelegatedResourceAccessRequesterTypeEnumStringValues() []string {
	return []string{
		"OPERATOR",
		"CUSTOMER",
		"SYSTEM",
	}
}

// GetMappingDelegatedResourceAccessRequesterTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDelegatedResourceAccessRequesterTypeEnum(val string) (DelegatedResourceAccessRequesterTypeEnum, bool) {
	enum, ok := mappingDelegatedResourceAccessRequesterTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
