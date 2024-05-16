// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.cloud.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"strings"
)

// BackboneFailoverPolicyEnum Enum with underlying type: string
type BackboneFailoverPolicyEnum string

// Set of constants representing the allowable values for BackboneFailoverPolicyEnum
const (
	BackboneFailoverPolicyFailoverToInternet      BackboneFailoverPolicyEnum = "FAILOVER_TO_INTERNET"
	BackboneFailoverPolicyDoNotFailoverToInternet BackboneFailoverPolicyEnum = "DO_NOT_FAILOVER_TO_INTERNET"
)

var mappingBackboneFailoverPolicyEnum = map[string]BackboneFailoverPolicyEnum{
	"FAILOVER_TO_INTERNET":        BackboneFailoverPolicyFailoverToInternet,
	"DO_NOT_FAILOVER_TO_INTERNET": BackboneFailoverPolicyDoNotFailoverToInternet,
}

var mappingBackboneFailoverPolicyEnumLowerCase = map[string]BackboneFailoverPolicyEnum{
	"failover_to_internet":        BackboneFailoverPolicyFailoverToInternet,
	"do_not_failover_to_internet": BackboneFailoverPolicyDoNotFailoverToInternet,
}

// GetBackboneFailoverPolicyEnumValues Enumerates the set of values for BackboneFailoverPolicyEnum
func GetBackboneFailoverPolicyEnumValues() []BackboneFailoverPolicyEnum {
	values := make([]BackboneFailoverPolicyEnum, 0)
	for _, v := range mappingBackboneFailoverPolicyEnum {
		values = append(values, v)
	}
	return values
}

// GetBackboneFailoverPolicyEnumStringValues Enumerates the set of values in String for BackboneFailoverPolicyEnum
func GetBackboneFailoverPolicyEnumStringValues() []string {
	return []string{
		"FAILOVER_TO_INTERNET",
		"DO_NOT_FAILOVER_TO_INTERNET",
	}
}

// GetMappingBackboneFailoverPolicyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBackboneFailoverPolicyEnum(val string) (BackboneFailoverPolicyEnum, bool) {
	enum, ok := mappingBackboneFailoverPolicyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
