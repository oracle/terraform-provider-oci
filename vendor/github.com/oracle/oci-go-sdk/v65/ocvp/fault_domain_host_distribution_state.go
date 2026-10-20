// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud VMware Solution API
//
// Use the Oracle Cloud VMware API to create SDDCs and manage ESXi hosts and software.
// For more information, see Oracle Cloud VMware Solution (https://docs.oracle.com/iaas/Content/VMware/Concepts/ocvsoverview.htm).
//

package ocvp

import (
	"strings"
)

// FaultDomainHostDistributionStateEnum Enum with underlying type: string
type FaultDomainHostDistributionStateEnum string

// Set of constants representing the allowable values for FaultDomainHostDistributionStateEnum
const (
	FaultDomainHostDistributionStateEvenlyDistributed   FaultDomainHostDistributionStateEnum = "EVENLY_DISTRIBUTED"
	FaultDomainHostDistributionStateUnevenlyDistributed FaultDomainHostDistributionStateEnum = "UNEVENLY_DISTRIBUTED"
	FaultDomainHostDistributionStateUnsupported         FaultDomainHostDistributionStateEnum = "UNSUPPORTED"
)

var mappingFaultDomainHostDistributionStateEnum = map[string]FaultDomainHostDistributionStateEnum{
	"EVENLY_DISTRIBUTED":   FaultDomainHostDistributionStateEvenlyDistributed,
	"UNEVENLY_DISTRIBUTED": FaultDomainHostDistributionStateUnevenlyDistributed,
	"UNSUPPORTED":          FaultDomainHostDistributionStateUnsupported,
}

var mappingFaultDomainHostDistributionStateEnumLowerCase = map[string]FaultDomainHostDistributionStateEnum{
	"evenly_distributed":   FaultDomainHostDistributionStateEvenlyDistributed,
	"unevenly_distributed": FaultDomainHostDistributionStateUnevenlyDistributed,
	"unsupported":          FaultDomainHostDistributionStateUnsupported,
}

// GetFaultDomainHostDistributionStateEnumValues Enumerates the set of values for FaultDomainHostDistributionStateEnum
func GetFaultDomainHostDistributionStateEnumValues() []FaultDomainHostDistributionStateEnum {
	values := make([]FaultDomainHostDistributionStateEnum, 0)
	for _, v := range mappingFaultDomainHostDistributionStateEnum {
		values = append(values, v)
	}
	return values
}

// GetFaultDomainHostDistributionStateEnumStringValues Enumerates the set of values in String for FaultDomainHostDistributionStateEnum
func GetFaultDomainHostDistributionStateEnumStringValues() []string {
	return []string{
		"EVENLY_DISTRIBUTED",
		"UNEVENLY_DISTRIBUTED",
		"UNSUPPORTED",
	}
}

// GetMappingFaultDomainHostDistributionStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFaultDomainHostDistributionStateEnum(val string) (FaultDomainHostDistributionStateEnum, bool) {
	enum, ok := mappingFaultDomainHostDistributionStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
