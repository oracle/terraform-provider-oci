// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// NetworkLoadBalancer API
//
// A description of the network load balancer API
//

package networkloadbalancer

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
)

var mappingOperationType = map[string]OperationTypeEnum{
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
}

// GetOperationTypeEnumValues Enumerates the set of values for OperationTypeEnum
func GetOperationTypeEnumValues() []OperationTypeEnum {
	values := make([]OperationTypeEnum, 0)
	for _, v := range mappingOperationType {
		values = append(values, v)
	}
	return values
}
