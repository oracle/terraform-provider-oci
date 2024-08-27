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

// DelegationControlResourceTypeEnum Enum with underlying type: string
type DelegationControlResourceTypeEnum string

// Set of constants representing the allowable values for DelegationControlResourceTypeEnum
const (
	DelegationControlResourceTypeVmcluster      DelegationControlResourceTypeEnum = "VMCLUSTER"
	DelegationControlResourceTypeCloudvmcluster DelegationControlResourceTypeEnum = "CLOUDVMCLUSTER"
)

var mappingDelegationControlResourceTypeEnum = map[string]DelegationControlResourceTypeEnum{
	"VMCLUSTER":      DelegationControlResourceTypeVmcluster,
	"CLOUDVMCLUSTER": DelegationControlResourceTypeCloudvmcluster,
}

var mappingDelegationControlResourceTypeEnumLowerCase = map[string]DelegationControlResourceTypeEnum{
	"vmcluster":      DelegationControlResourceTypeVmcluster,
	"cloudvmcluster": DelegationControlResourceTypeCloudvmcluster,
}

// GetDelegationControlResourceTypeEnumValues Enumerates the set of values for DelegationControlResourceTypeEnum
func GetDelegationControlResourceTypeEnumValues() []DelegationControlResourceTypeEnum {
	values := make([]DelegationControlResourceTypeEnum, 0)
	for _, v := range mappingDelegationControlResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDelegationControlResourceTypeEnumStringValues Enumerates the set of values in String for DelegationControlResourceTypeEnum
func GetDelegationControlResourceTypeEnumStringValues() []string {
	return []string{
		"VMCLUSTER",
		"CLOUDVMCLUSTER",
	}
}

// GetMappingDelegationControlResourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDelegationControlResourceTypeEnum(val string) (DelegationControlResourceTypeEnum, bool) {
	enum, ok := mappingDelegationControlResourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
