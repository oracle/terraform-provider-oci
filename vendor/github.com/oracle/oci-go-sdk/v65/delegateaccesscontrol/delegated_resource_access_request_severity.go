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

// DelegatedResourceAccessRequestSeverityEnum Enum with underlying type: string
type DelegatedResourceAccessRequestSeverityEnum string

// Set of constants representing the allowable values for DelegatedResourceAccessRequestSeverityEnum
const (
	DelegatedResourceAccessRequestSeverityS1 DelegatedResourceAccessRequestSeverityEnum = "S1"
	DelegatedResourceAccessRequestSeverityS2 DelegatedResourceAccessRequestSeverityEnum = "S2"
	DelegatedResourceAccessRequestSeverityS3 DelegatedResourceAccessRequestSeverityEnum = "S3"
	DelegatedResourceAccessRequestSeverityS4 DelegatedResourceAccessRequestSeverityEnum = "S4"
)

var mappingDelegatedResourceAccessRequestSeverityEnum = map[string]DelegatedResourceAccessRequestSeverityEnum{
	"S1": DelegatedResourceAccessRequestSeverityS1,
	"S2": DelegatedResourceAccessRequestSeverityS2,
	"S3": DelegatedResourceAccessRequestSeverityS3,
	"S4": DelegatedResourceAccessRequestSeverityS4,
}

var mappingDelegatedResourceAccessRequestSeverityEnumLowerCase = map[string]DelegatedResourceAccessRequestSeverityEnum{
	"s1": DelegatedResourceAccessRequestSeverityS1,
	"s2": DelegatedResourceAccessRequestSeverityS2,
	"s3": DelegatedResourceAccessRequestSeverityS3,
	"s4": DelegatedResourceAccessRequestSeverityS4,
}

// GetDelegatedResourceAccessRequestSeverityEnumValues Enumerates the set of values for DelegatedResourceAccessRequestSeverityEnum
func GetDelegatedResourceAccessRequestSeverityEnumValues() []DelegatedResourceAccessRequestSeverityEnum {
	values := make([]DelegatedResourceAccessRequestSeverityEnum, 0)
	for _, v := range mappingDelegatedResourceAccessRequestSeverityEnum {
		values = append(values, v)
	}
	return values
}

// GetDelegatedResourceAccessRequestSeverityEnumStringValues Enumerates the set of values in String for DelegatedResourceAccessRequestSeverityEnum
func GetDelegatedResourceAccessRequestSeverityEnumStringValues() []string {
	return []string{
		"S1",
		"S2",
		"S3",
		"S4",
	}
}

// GetMappingDelegatedResourceAccessRequestSeverityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDelegatedResourceAccessRequestSeverityEnum(val string) (DelegatedResourceAccessRequestSeverityEnum, bool) {
	enum, ok := mappingDelegatedResourceAccessRequestSeverityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
