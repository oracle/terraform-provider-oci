package service_mesh

import (
	oci_service_mesh "github.com/oracle/oci-go-sdk/v65/servicemesh"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("service_mesh", serviceMeshResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportServiceMeshVirtualServiceHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_service_mesh_virtual_service",
	DatasourceClass:        "oci_service_mesh_virtual_services",
	DatasourceItemsAttr:    "virtual_service_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "virtual_service",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_service_mesh.VirtualServiceLifecycleStateActive),
	},
}

var exportServiceMeshAccessPolicyHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_service_mesh_access_policy",
	DatasourceClass:        "oci_service_mesh_access_policies",
	DatasourceItemsAttr:    "access_policy_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "access_policy",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_service_mesh.AccessPolicyLifecycleStateActive),
	},
}

var exportServiceMeshMeshHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_service_mesh_mesh",
	DatasourceClass:        "oci_service_mesh_meshes",
	DatasourceItemsAttr:    "mesh_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "mesh",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_service_mesh.MeshLifecycleStateActive),
	},
}

var exportServiceMeshIngressGatewayRouteTableHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_service_mesh_ingress_gateway_route_table",
	DatasourceClass:        "oci_service_mesh_ingress_gateway_route_tables",
	DatasourceItemsAttr:    "ingress_gateway_route_table_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "ingress_gateway_route_table",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_service_mesh.IngressGatewayRouteTableLifecycleStateActive),
	},
}

var exportServiceMeshVirtualServiceRouteTableHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_service_mesh_virtual_service_route_table",
	DatasourceClass:        "oci_service_mesh_virtual_service_route_tables",
	DatasourceItemsAttr:    "virtual_service_route_table_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "virtual_service_route_table",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_service_mesh.VirtualServiceRouteTableLifecycleStateActive),
	},
}

var exportServiceMeshVirtualDeploymentHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_service_mesh_virtual_deployment",
	DatasourceClass:        "oci_service_mesh_virtual_deployments",
	DatasourceItemsAttr:    "virtual_deployment_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "virtual_deployment",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_service_mesh.VirtualDeploymentLifecycleStateActive),
	},
}

var exportServiceMeshIngressGatewayHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_service_mesh_ingress_gateway",
	DatasourceClass:        "oci_service_mesh_ingress_gateways",
	DatasourceItemsAttr:    "ingress_gateway_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "ingress_gateway",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_service_mesh.IngressGatewayLifecycleStateActive),
	},
}

var serviceMeshResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportServiceMeshVirtualServiceHints},
		{TerraformResourceHints: exportServiceMeshAccessPolicyHints},
		{TerraformResourceHints: exportServiceMeshMeshHints},
		{TerraformResourceHints: exportServiceMeshIngressGatewayRouteTableHints},
		{TerraformResourceHints: exportServiceMeshVirtualServiceRouteTableHints},
		{TerraformResourceHints: exportServiceMeshVirtualDeploymentHints},
		{TerraformResourceHints: exportServiceMeshIngressGatewayHints},
	},
}
