// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package service_mesh

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_service_mesh_access_policies", ServiceMeshAccessPoliciesDataSource())
	tfresource.RegisterDatasource("oci_service_mesh_access_policy", ServiceMeshAccessPolicyDataSource())
	tfresource.RegisterDatasource("oci_service_mesh_ingress_gateway", ServiceMeshIngressGatewayDataSource())
	tfresource.RegisterDatasource("oci_service_mesh_ingress_gateway_route_table", ServiceMeshIngressGatewayRouteTableDataSource())
	tfresource.RegisterDatasource("oci_service_mesh_ingress_gateway_route_tables", ServiceMeshIngressGatewayRouteTablesDataSource())
	tfresource.RegisterDatasource("oci_service_mesh_ingress_gateways", ServiceMeshIngressGatewaysDataSource())
	tfresource.RegisterDatasource("oci_service_mesh_mesh", ServiceMeshMeshDataSource())
	tfresource.RegisterDatasource("oci_service_mesh_meshes", ServiceMeshMeshesDataSource())
	tfresource.RegisterDatasource("oci_service_mesh_proxy_detail", ServiceMeshProxyDetailDataSource())
	tfresource.RegisterDatasource("oci_service_mesh_virtual_deployment", ServiceMeshVirtualDeploymentDataSource())
	tfresource.RegisterDatasource("oci_service_mesh_virtual_deployments", ServiceMeshVirtualDeploymentsDataSource())
	tfresource.RegisterDatasource("oci_service_mesh_virtual_service", ServiceMeshVirtualServiceDataSource())
	tfresource.RegisterDatasource("oci_service_mesh_virtual_service_route_table", ServiceMeshVirtualServiceRouteTableDataSource())
	tfresource.RegisterDatasource("oci_service_mesh_virtual_service_route_tables", ServiceMeshVirtualServiceRouteTablesDataSource())
	tfresource.RegisterDatasource("oci_service_mesh_virtual_services", ServiceMeshVirtualServicesDataSource())
}
