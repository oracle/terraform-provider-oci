// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package service_mesh

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_service_mesh_access_policy", ServiceMeshAccessPolicyResource())
	tfresource.RegisterResource("oci_service_mesh_ingress_gateway", ServiceMeshIngressGatewayResource())
	tfresource.RegisterResource("oci_service_mesh_ingress_gateway_route_table", ServiceMeshIngressGatewayRouteTableResource())
	tfresource.RegisterResource("oci_service_mesh_mesh", ServiceMeshMeshResource())
	tfresource.RegisterResource("oci_service_mesh_virtual_deployment", ServiceMeshVirtualDeploymentResource())
	tfresource.RegisterResource("oci_service_mesh_virtual_service", ServiceMeshVirtualServiceResource())
	tfresource.RegisterResource("oci_service_mesh_virtual_service_route_table", ServiceMeshVirtualServiceRouteTableResource())
}
