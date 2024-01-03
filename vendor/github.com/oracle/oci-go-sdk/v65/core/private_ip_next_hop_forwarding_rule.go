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

// PrivateIpNextHopForwardingRuleEnum Enum with underlying type: string
type PrivateIpNextHopForwardingRuleEnum string

// Set of constants representing the allowable values for PrivateIpNextHopForwardingRuleEnum
const (
	PrivateIpNextHopForwardingRuleNotSpecified     PrivateIpNextHopForwardingRuleEnum = "NOT_SPECIFIED"
	PrivateIpNextHopForwardingRuleSkipPortSharding PrivateIpNextHopForwardingRuleEnum = "SKIP_PORT_SHARDING"
	PrivateIpNextHopForwardingRuleWildcardListener PrivateIpNextHopForwardingRuleEnum = "WILDCARD_LISTENER"
)

var mappingPrivateIpNextHopForwardingRuleEnum = map[string]PrivateIpNextHopForwardingRuleEnum{
	"NOT_SPECIFIED":      PrivateIpNextHopForwardingRuleNotSpecified,
	"SKIP_PORT_SHARDING": PrivateIpNextHopForwardingRuleSkipPortSharding,
	"WILDCARD_LISTENER":  PrivateIpNextHopForwardingRuleWildcardListener,
}

var mappingPrivateIpNextHopForwardingRuleEnumLowerCase = map[string]PrivateIpNextHopForwardingRuleEnum{
	"not_specified":      PrivateIpNextHopForwardingRuleNotSpecified,
	"skip_port_sharding": PrivateIpNextHopForwardingRuleSkipPortSharding,
	"wildcard_listener":  PrivateIpNextHopForwardingRuleWildcardListener,
}

// GetPrivateIpNextHopForwardingRuleEnumValues Enumerates the set of values for PrivateIpNextHopForwardingRuleEnum
func GetPrivateIpNextHopForwardingRuleEnumValues() []PrivateIpNextHopForwardingRuleEnum {
	values := make([]PrivateIpNextHopForwardingRuleEnum, 0)
	for _, v := range mappingPrivateIpNextHopForwardingRuleEnum {
		values = append(values, v)
	}
	return values
}

// GetPrivateIpNextHopForwardingRuleEnumStringValues Enumerates the set of values in String for PrivateIpNextHopForwardingRuleEnum
func GetPrivateIpNextHopForwardingRuleEnumStringValues() []string {
	return []string{
		"NOT_SPECIFIED",
		"SKIP_PORT_SHARDING",
		"WILDCARD_LISTENER",
	}
}

// GetMappingPrivateIpNextHopForwardingRuleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPrivateIpNextHopForwardingRuleEnum(val string) (PrivateIpNextHopForwardingRuleEnum, bool) {
	enum, ok := mappingPrivateIpNextHopForwardingRuleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
