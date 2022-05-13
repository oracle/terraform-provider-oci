// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Mesh API
//
// Use the Service Mesh API to manage mesh, virtual service, access policy and other mesh related items.
//

package servicemesh

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateMesh                     OperationTypeEnum = "CREATE_MESH"
	OperationTypeUpdateMesh                     OperationTypeEnum = "UPDATE_MESH"
	OperationTypeDeleteMesh                     OperationTypeEnum = "DELETE_MESH"
	OperationTypeMoveMesh                       OperationTypeEnum = "MOVE_MESH"
	OperationTypeCreateAccessPolicy             OperationTypeEnum = "CREATE_ACCESS_POLICY"
	OperationTypeUpdateAccessPolicy             OperationTypeEnum = "UPDATE_ACCESS_POLICY"
	OperationTypeDeleteAccessPolicy             OperationTypeEnum = "DELETE_ACCESS_POLICY"
	OperationTypeMoveAccessPolicy               OperationTypeEnum = "MOVE_ACCESS_POLICY"
	OperationTypeCreateVirtualService           OperationTypeEnum = "CREATE_VIRTUAL_SERVICE"
	OperationTypeUpdateVirtualService           OperationTypeEnum = "UPDATE_VIRTUAL_SERVICE"
	OperationTypeDeleteVirtualService           OperationTypeEnum = "DELETE_VIRTUAL_SERVICE"
	OperationTypeMoveVirtualService             OperationTypeEnum = "MOVE_VIRTUAL_SERVICE"
	OperationTypeCreateVirtualServiceRouteTable OperationTypeEnum = "CREATE_VIRTUAL_SERVICE_ROUTE_TABLE"
	OperationTypeUpdateVirtualServiceRouteTable OperationTypeEnum = "UPDATE_VIRTUAL_SERVICE_ROUTE_TABLE"
	OperationTypeDeleteVirtualServiceRouteTable OperationTypeEnum = "DELETE_VIRTUAL_SERVICE_ROUTE_TABLE"
	OperationTypeMoveVirtualServiceRouteTable   OperationTypeEnum = "MOVE_VIRTUAL_SERVICE_ROUTE_TABLE"
	OperationTypeCreateVirtualDeployment        OperationTypeEnum = "CREATE_VIRTUAL_DEPLOYMENT"
	OperationTypeUpdateVirtualDeployment        OperationTypeEnum = "UPDATE_VIRTUAL_DEPLOYMENT"
	OperationTypeDeleteVirtualDeployment        OperationTypeEnum = "DELETE_VIRTUAL_DEPLOYMENT"
	OperationTypeMoveVirtualDeployment          OperationTypeEnum = "MOVE_VIRTUAL_DEPLOYMENT"
	OperationTypeCreateIngressGateway           OperationTypeEnum = "CREATE_INGRESS_GATEWAY"
	OperationTypeUpdateIngressGateway           OperationTypeEnum = "UPDATE_INGRESS_GATEWAY"
	OperationTypeDeleteIngressGateway           OperationTypeEnum = "DELETE_INGRESS_GATEWAY"
	OperationTypeMoveIngressGateway             OperationTypeEnum = "MOVE_INGRESS_GATEWAY"
	OperationTypeCreateIngressGatewayRouteTable OperationTypeEnum = "CREATE_INGRESS_GATEWAY_ROUTE_TABLE"
	OperationTypeUpdateIngressGatewayRouteTable OperationTypeEnum = "UPDATE_INGRESS_GATEWAY_ROUTE_TABLE"
	OperationTypeDeleteIngressGatewayRouteTable OperationTypeEnum = "DELETE_INGRESS_GATEWAY_ROUTE_TABLE"
	OperationTypeMoveIngressGatewayRouteTable   OperationTypeEnum = "MOVE_INGRESS_GATEWAY_ROUTE_TABLE"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_MESH":                        OperationTypeCreateMesh,
	"UPDATE_MESH":                        OperationTypeUpdateMesh,
	"DELETE_MESH":                        OperationTypeDeleteMesh,
	"MOVE_MESH":                          OperationTypeMoveMesh,
	"CREATE_ACCESS_POLICY":               OperationTypeCreateAccessPolicy,
	"UPDATE_ACCESS_POLICY":               OperationTypeUpdateAccessPolicy,
	"DELETE_ACCESS_POLICY":               OperationTypeDeleteAccessPolicy,
	"MOVE_ACCESS_POLICY":                 OperationTypeMoveAccessPolicy,
	"CREATE_VIRTUAL_SERVICE":             OperationTypeCreateVirtualService,
	"UPDATE_VIRTUAL_SERVICE":             OperationTypeUpdateVirtualService,
	"DELETE_VIRTUAL_SERVICE":             OperationTypeDeleteVirtualService,
	"MOVE_VIRTUAL_SERVICE":               OperationTypeMoveVirtualService,
	"CREATE_VIRTUAL_SERVICE_ROUTE_TABLE": OperationTypeCreateVirtualServiceRouteTable,
	"UPDATE_VIRTUAL_SERVICE_ROUTE_TABLE": OperationTypeUpdateVirtualServiceRouteTable,
	"DELETE_VIRTUAL_SERVICE_ROUTE_TABLE": OperationTypeDeleteVirtualServiceRouteTable,
	"MOVE_VIRTUAL_SERVICE_ROUTE_TABLE":   OperationTypeMoveVirtualServiceRouteTable,
	"CREATE_VIRTUAL_DEPLOYMENT":          OperationTypeCreateVirtualDeployment,
	"UPDATE_VIRTUAL_DEPLOYMENT":          OperationTypeUpdateVirtualDeployment,
	"DELETE_VIRTUAL_DEPLOYMENT":          OperationTypeDeleteVirtualDeployment,
	"MOVE_VIRTUAL_DEPLOYMENT":            OperationTypeMoveVirtualDeployment,
	"CREATE_INGRESS_GATEWAY":             OperationTypeCreateIngressGateway,
	"UPDATE_INGRESS_GATEWAY":             OperationTypeUpdateIngressGateway,
	"DELETE_INGRESS_GATEWAY":             OperationTypeDeleteIngressGateway,
	"MOVE_INGRESS_GATEWAY":               OperationTypeMoveIngressGateway,
	"CREATE_INGRESS_GATEWAY_ROUTE_TABLE": OperationTypeCreateIngressGatewayRouteTable,
	"UPDATE_INGRESS_GATEWAY_ROUTE_TABLE": OperationTypeUpdateIngressGatewayRouteTable,
	"DELETE_INGRESS_GATEWAY_ROUTE_TABLE": OperationTypeDeleteIngressGatewayRouteTable,
	"MOVE_INGRESS_GATEWAY_ROUTE_TABLE":   OperationTypeMoveIngressGatewayRouteTable,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_mesh":                        OperationTypeCreateMesh,
	"update_mesh":                        OperationTypeUpdateMesh,
	"delete_mesh":                        OperationTypeDeleteMesh,
	"move_mesh":                          OperationTypeMoveMesh,
	"create_access_policy":               OperationTypeCreateAccessPolicy,
	"update_access_policy":               OperationTypeUpdateAccessPolicy,
	"delete_access_policy":               OperationTypeDeleteAccessPolicy,
	"move_access_policy":                 OperationTypeMoveAccessPolicy,
	"create_virtual_service":             OperationTypeCreateVirtualService,
	"update_virtual_service":             OperationTypeUpdateVirtualService,
	"delete_virtual_service":             OperationTypeDeleteVirtualService,
	"move_virtual_service":               OperationTypeMoveVirtualService,
	"create_virtual_service_route_table": OperationTypeCreateVirtualServiceRouteTable,
	"update_virtual_service_route_table": OperationTypeUpdateVirtualServiceRouteTable,
	"delete_virtual_service_route_table": OperationTypeDeleteVirtualServiceRouteTable,
	"move_virtual_service_route_table":   OperationTypeMoveVirtualServiceRouteTable,
	"create_virtual_deployment":          OperationTypeCreateVirtualDeployment,
	"update_virtual_deployment":          OperationTypeUpdateVirtualDeployment,
	"delete_virtual_deployment":          OperationTypeDeleteVirtualDeployment,
	"move_virtual_deployment":            OperationTypeMoveVirtualDeployment,
	"create_ingress_gateway":             OperationTypeCreateIngressGateway,
	"update_ingress_gateway":             OperationTypeUpdateIngressGateway,
	"delete_ingress_gateway":             OperationTypeDeleteIngressGateway,
	"move_ingress_gateway":               OperationTypeMoveIngressGateway,
	"create_ingress_gateway_route_table": OperationTypeCreateIngressGatewayRouteTable,
	"update_ingress_gateway_route_table": OperationTypeUpdateIngressGatewayRouteTable,
	"delete_ingress_gateway_route_table": OperationTypeDeleteIngressGatewayRouteTable,
	"move_ingress_gateway_route_table":   OperationTypeMoveIngressGatewayRouteTable,
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
		"CREATE_MESH",
		"UPDATE_MESH",
		"DELETE_MESH",
		"MOVE_MESH",
		"CREATE_ACCESS_POLICY",
		"UPDATE_ACCESS_POLICY",
		"DELETE_ACCESS_POLICY",
		"MOVE_ACCESS_POLICY",
		"CREATE_VIRTUAL_SERVICE",
		"UPDATE_VIRTUAL_SERVICE",
		"DELETE_VIRTUAL_SERVICE",
		"MOVE_VIRTUAL_SERVICE",
		"CREATE_VIRTUAL_SERVICE_ROUTE_TABLE",
		"UPDATE_VIRTUAL_SERVICE_ROUTE_TABLE",
		"DELETE_VIRTUAL_SERVICE_ROUTE_TABLE",
		"MOVE_VIRTUAL_SERVICE_ROUTE_TABLE",
		"CREATE_VIRTUAL_DEPLOYMENT",
		"UPDATE_VIRTUAL_DEPLOYMENT",
		"DELETE_VIRTUAL_DEPLOYMENT",
		"MOVE_VIRTUAL_DEPLOYMENT",
		"CREATE_INGRESS_GATEWAY",
		"UPDATE_INGRESS_GATEWAY",
		"DELETE_INGRESS_GATEWAY",
		"MOVE_INGRESS_GATEWAY",
		"CREATE_INGRESS_GATEWAY_ROUTE_TABLE",
		"UPDATE_INGRESS_GATEWAY_ROUTE_TABLE",
		"DELETE_INGRESS_GATEWAY_ROUTE_TABLE",
		"MOVE_INGRESS_GATEWAY_ROUTE_TABLE",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
