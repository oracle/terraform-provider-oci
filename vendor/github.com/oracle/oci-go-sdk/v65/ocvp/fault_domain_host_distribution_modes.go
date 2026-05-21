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

// FaultDomainHostDistributionModesEnum Enum with underlying type: string
type FaultDomainHostDistributionModesEnum string

// Set of constants representing the allowable values for FaultDomainHostDistributionModesEnum
const (
	FaultDomainHostDistributionModesEvenlyDistributed   FaultDomainHostDistributionModesEnum = "EVENLY_DISTRIBUTED"
	FaultDomainHostDistributionModesUnevenlyDistributed FaultDomainHostDistributionModesEnum = "UNEVENLY_DISTRIBUTED"
)

var mappingFaultDomainHostDistributionModesEnum = map[string]FaultDomainHostDistributionModesEnum{
	"EVENLY_DISTRIBUTED":   FaultDomainHostDistributionModesEvenlyDistributed,
	"UNEVENLY_DISTRIBUTED": FaultDomainHostDistributionModesUnevenlyDistributed,
}

var mappingFaultDomainHostDistributionModesEnumLowerCase = map[string]FaultDomainHostDistributionModesEnum{
	"evenly_distributed":   FaultDomainHostDistributionModesEvenlyDistributed,
	"unevenly_distributed": FaultDomainHostDistributionModesUnevenlyDistributed,
}

// GetFaultDomainHostDistributionModesEnumValues Enumerates the set of values for FaultDomainHostDistributionModesEnum
func GetFaultDomainHostDistributionModesEnumValues() []FaultDomainHostDistributionModesEnum {
	values := make([]FaultDomainHostDistributionModesEnum, 0)
	for _, v := range mappingFaultDomainHostDistributionModesEnum {
		values = append(values, v)
	}
	return values
}

// GetFaultDomainHostDistributionModesEnumStringValues Enumerates the set of values in String for FaultDomainHostDistributionModesEnum
func GetFaultDomainHostDistributionModesEnumStringValues() []string {
	return []string{
		"EVENLY_DISTRIBUTED",
		"UNEVENLY_DISTRIBUTED",
	}
}

// GetMappingFaultDomainHostDistributionModesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFaultDomainHostDistributionModesEnum(val string) (FaultDomainHostDistributionModesEnum, bool) {
	enum, ok := mappingFaultDomainHostDistributionModesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
