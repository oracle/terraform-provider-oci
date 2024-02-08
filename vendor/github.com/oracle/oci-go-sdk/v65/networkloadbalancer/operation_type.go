// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// NetworkLoadBalancer API
//
// This describes the network load balancer API.
//

package networkloadbalancer

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateNetworkLoadBalancer OperationTypeEnum = "CREATE_NETWORK_LOAD_BALANCER"
	OperationTypeUpdateNetworkLoadBalancer OperationTypeEnum = "UPDATE_NETWORK_LOAD_BALANCER"
	OperationTypeDeleteNetworkLoadBalancer OperationTypeEnum = "DELETE_NETWORK_LOAD_BALANCER"
	OperationTypeCreateBackend             OperationTypeEnum = "CREATE_BACKEND"
	OperationTypeUpdateBackend             OperationTypeEnum = "UPDATE_BACKEND"
	OperationTypeDeleteBackend             OperationTypeEnum = "DELETE_BACKEND"
	OperationTypeCreateListener            OperationTypeEnum = "CREATE_LISTENER"
	OperationTypeUpdateListener            OperationTypeEnum = "UPDATE_LISTENER"
	OperationTypeDeleteListener            OperationTypeEnum = "DELETE_LISTENER"
	OperationTypeCreateBackendset          OperationTypeEnum = "CREATE_BACKENDSET"
	OperationTypeUpdateBackendset          OperationTypeEnum = "UPDATE_BACKENDSET"
	OperationTypeDeleteBackendset          OperationTypeEnum = "DELETE_BACKENDSET"
	OperationTypeUpdateNsgs                OperationTypeEnum = "UPDATE_NSGS"
	OperationTypeUpdateHealthChecker       OperationTypeEnum = "UPDATE_HEALTH_CHECKER"
	OperationTypeChangeCompartment         OperationTypeEnum = "CHANGE_COMPARTMENT"
	OperationTypeAttachNlbToPod            OperationTypeEnum = "ATTACH_NLB_TO_POD"
	OperationTypeDetachNlbFromPod          OperationTypeEnum = "DETACH_NLB_FROM_POD"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_NETWORK_LOAD_BALANCER": OperationTypeCreateNetworkLoadBalancer,
	"UPDATE_NETWORK_LOAD_BALANCER": OperationTypeUpdateNetworkLoadBalancer,
	"DELETE_NETWORK_LOAD_BALANCER": OperationTypeDeleteNetworkLoadBalancer,
	"CREATE_BACKEND":               OperationTypeCreateBackend,
	"UPDATE_BACKEND":               OperationTypeUpdateBackend,
	"DELETE_BACKEND":               OperationTypeDeleteBackend,
	"CREATE_LISTENER":              OperationTypeCreateListener,
	"UPDATE_LISTENER":              OperationTypeUpdateListener,
	"DELETE_LISTENER":              OperationTypeDeleteListener,
	"CREATE_BACKENDSET":            OperationTypeCreateBackendset,
	"UPDATE_BACKENDSET":            OperationTypeUpdateBackendset,
	"DELETE_BACKENDSET":            OperationTypeDeleteBackendset,
	"UPDATE_NSGS":                  OperationTypeUpdateNsgs,
	"UPDATE_HEALTH_CHECKER":        OperationTypeUpdateHealthChecker,
	"CHANGE_COMPARTMENT":           OperationTypeChangeCompartment,
	"ATTACH_NLB_TO_POD":            OperationTypeAttachNlbToPod,
	"DETACH_NLB_FROM_POD":          OperationTypeDetachNlbFromPod,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_network_load_balancer": OperationTypeCreateNetworkLoadBalancer,
	"update_network_load_balancer": OperationTypeUpdateNetworkLoadBalancer,
	"delete_network_load_balancer": OperationTypeDeleteNetworkLoadBalancer,
	"create_backend":               OperationTypeCreateBackend,
	"update_backend":               OperationTypeUpdateBackend,
	"delete_backend":               OperationTypeDeleteBackend,
	"create_listener":              OperationTypeCreateListener,
	"update_listener":              OperationTypeUpdateListener,
	"delete_listener":              OperationTypeDeleteListener,
	"create_backendset":            OperationTypeCreateBackendset,
	"update_backendset":            OperationTypeUpdateBackendset,
	"delete_backendset":            OperationTypeDeleteBackendset,
	"update_nsgs":                  OperationTypeUpdateNsgs,
	"update_health_checker":        OperationTypeUpdateHealthChecker,
	"change_compartment":           OperationTypeChangeCompartment,
	"attach_nlb_to_pod":            OperationTypeAttachNlbToPod,
	"detach_nlb_from_pod":          OperationTypeDetachNlbFromPod,
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
		"CREATE_NETWORK_LOAD_BALANCER",
		"UPDATE_NETWORK_LOAD_BALANCER",
		"DELETE_NETWORK_LOAD_BALANCER",
		"CREATE_BACKEND",
		"UPDATE_BACKEND",
		"DELETE_BACKEND",
		"CREATE_LISTENER",
		"UPDATE_LISTENER",
		"DELETE_LISTENER",
		"CREATE_BACKENDSET",
		"UPDATE_BACKENDSET",
		"DELETE_BACKENDSET",
		"UPDATE_NSGS",
		"UPDATE_HEALTH_CHECKER",
		"CHANGE_COMPARTMENT",
		"ATTACH_NLB_TO_POD",
		"DETACH_NLB_FROM_POD",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
