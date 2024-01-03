// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Firewall API
//
// Use the Network Firewall API to create network firewalls and configure policies that regulates network traffic in and across VCNs.
//

package networkfirewall

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateNetworkFirewall       OperationTypeEnum = "CREATE_NETWORK_FIREWALL"
	OperationTypeUpdateNetworkFirewall       OperationTypeEnum = "UPDATE_NETWORK_FIREWALL"
	OperationTypeDeleteNetworkFirewall       OperationTypeEnum = "DELETE_NETWORK_FIREWALL"
	OperationTypeMoveNetworkFirewall         OperationTypeEnum = "MOVE_NETWORK_FIREWALL"
	OperationTypeCreateNetworkFirewallPolicy OperationTypeEnum = "CREATE_NETWORK_FIREWALL_POLICY"
	OperationTypeUpdateNetworkFirewallPolicy OperationTypeEnum = "UPDATE_NETWORK_FIREWALL_POLICY"
	OperationTypeDeleteNetworkFirewallPolicy OperationTypeEnum = "DELETE_NETWORK_FIREWALL_POLICY"
	OperationTypeMoveNetworkFirewallPolicy   OperationTypeEnum = "MOVE_NETWORK_FIREWALL_POLICY"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_NETWORK_FIREWALL":        OperationTypeCreateNetworkFirewall,
	"UPDATE_NETWORK_FIREWALL":        OperationTypeUpdateNetworkFirewall,
	"DELETE_NETWORK_FIREWALL":        OperationTypeDeleteNetworkFirewall,
	"MOVE_NETWORK_FIREWALL":          OperationTypeMoveNetworkFirewall,
	"CREATE_NETWORK_FIREWALL_POLICY": OperationTypeCreateNetworkFirewallPolicy,
	"UPDATE_NETWORK_FIREWALL_POLICY": OperationTypeUpdateNetworkFirewallPolicy,
	"DELETE_NETWORK_FIREWALL_POLICY": OperationTypeDeleteNetworkFirewallPolicy,
	"MOVE_NETWORK_FIREWALL_POLICY":   OperationTypeMoveNetworkFirewallPolicy,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_network_firewall":        OperationTypeCreateNetworkFirewall,
	"update_network_firewall":        OperationTypeUpdateNetworkFirewall,
	"delete_network_firewall":        OperationTypeDeleteNetworkFirewall,
	"move_network_firewall":          OperationTypeMoveNetworkFirewall,
	"create_network_firewall_policy": OperationTypeCreateNetworkFirewallPolicy,
	"update_network_firewall_policy": OperationTypeUpdateNetworkFirewallPolicy,
	"delete_network_firewall_policy": OperationTypeDeleteNetworkFirewallPolicy,
	"move_network_firewall_policy":   OperationTypeMoveNetworkFirewallPolicy,
}

// GetOperationTypeEnumValues Enumerates the set of values for OperationTypeEnum
func GetOperationTypeEnumValues() []OperationTypeEnum {
	values := make([]OperationTypeEnum, 0)
	for _, v := range mappingOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationTypeEnumStringValues Enumerates the set of values in String for OperationTypeEnum
func GetOperationTypeEnumStringValues() []string {
	return []string{
		"CREATE_NETWORK_FIREWALL",
		"UPDATE_NETWORK_FIREWALL",
		"DELETE_NETWORK_FIREWALL",
		"MOVE_NETWORK_FIREWALL",
		"CREATE_NETWORK_FIREWALL_POLICY",
		"UPDATE_NETWORK_FIREWALL_POLICY",
		"DELETE_NETWORK_FIREWALL_POLICY",
		"MOVE_NETWORK_FIREWALL_POLICY",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
