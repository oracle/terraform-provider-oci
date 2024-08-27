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

// ServiceProviderServiceTypeEnum Enum with underlying type: string
type ServiceProviderServiceTypeEnum string

// Set of constants representing the allowable values for ServiceProviderServiceTypeEnum
const (
	ServiceProviderServiceTypeTroubleshooting  ServiceProviderServiceTypeEnum = "TROUBLESHOOTING"
	ServiceProviderServiceTypeAssistedPatching ServiceProviderServiceTypeEnum = "ASSISTED_PATCHING"
)

var mappingServiceProviderServiceTypeEnum = map[string]ServiceProviderServiceTypeEnum{
	"TROUBLESHOOTING":   ServiceProviderServiceTypeTroubleshooting,
	"ASSISTED_PATCHING": ServiceProviderServiceTypeAssistedPatching,
}

var mappingServiceProviderServiceTypeEnumLowerCase = map[string]ServiceProviderServiceTypeEnum{
	"troubleshooting":   ServiceProviderServiceTypeTroubleshooting,
	"assisted_patching": ServiceProviderServiceTypeAssistedPatching,
}

// GetServiceProviderServiceTypeEnumValues Enumerates the set of values for ServiceProviderServiceTypeEnum
func GetServiceProviderServiceTypeEnumValues() []ServiceProviderServiceTypeEnum {
	values := make([]ServiceProviderServiceTypeEnum, 0)
	for _, v := range mappingServiceProviderServiceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetServiceProviderServiceTypeEnumStringValues Enumerates the set of values in String for ServiceProviderServiceTypeEnum
func GetServiceProviderServiceTypeEnumStringValues() []string {
	return []string{
		"TROUBLESHOOTING",
		"ASSISTED_PATCHING",
	}
}

// GetMappingServiceProviderServiceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingServiceProviderServiceTypeEnum(val string) (ServiceProviderServiceTypeEnum, bool) {
	enum, ok := mappingServiceProviderServiceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
