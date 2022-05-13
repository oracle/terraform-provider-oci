---
subcategory: "Service Mesh"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_service_mesh_ingress_gateway_route_tables"
sidebar_current: "docs-oci-datasource-service_mesh-ingress_gateway_route_tables"
description: |-
  Provides the list of Ingress Gateway Route Tables in Oracle Cloud Infrastructure Service Mesh service
---

# Data Source: oci_service_mesh_ingress_gateway_route_tables
This data source provides the list of Ingress Gateway Route Tables in Oracle Cloud Infrastructure Service Mesh service.

Returns a list of IngressGatewayRouteTable objects.


## Example Usage

```hcl
data "oci_service_mesh_ingress_gateway_route_tables" "test_ingress_gateway_route_tables" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	id = var.ingress_gateway_route_table_id
	ingress_gateway_id = oci_service_mesh_ingress_gateway.test_ingress_gateway.id
	name = var.ingress_gateway_route_table_name
	state = var.ingress_gateway_route_table_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `id` - (Optional) Unique IngressGatewayRouteTable identifier.
* `ingress_gateway_id` - (Optional) Unique IngressGateway identifier.
* `name` - (Optional) A filter to return only resources that match the entire name given.
* `state` - (Optional) A filter to return only resources that match the life cycle state given.


## Attributes Reference

The following attributes are exported:

* `ingress_gateway_route_table_collection` - The list of ingress_gateway_route_table_collection.

### IngressGatewayRouteTable Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Description of the resource. It can be changed after creation. Avoid entering confidential information.  Example: `This is my new resource` 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier that is immutable on creation.
* `ingress_gateway_id` - The OCID of the ingress gateway.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in a Failed state.
* `name` - A user-friendly name. The name must be unique within the same ingress gateway and cannot be changed after creation. Avoid entering confidential information.  Example: `My unique resource name` 
* `priority` - The priority of the route table. A lower value means a higher priority. The routes are declared based on the priority.
* `route_rules` - The route rules for the ingress gateway.
	* `destinations` - The destination of the request.
		* `port` - The port on the virtual service to target. Mandatory if the virtual deployments are listening on multiple ports. 
		* `type` - Type of the traffic target.
		* `virtual_service_id` - The OCID of the virtual service where the request will be routed.
		* `weight` - Weight of traffic target.
	* `ingress_gateway_host` - The ingress gateway host to which the route rule attaches. If not specified, the route rule gets attached to all hosts on the ingress gateway. 
		* `name` - Name of the ingress gateway host that this route should apply to. 
		* `port` - The port of the ingress gateway host listener. Leave empty to match all ports for the host. 
	* `is_grpc` - If true, the rule will check that the content-type header has a application/grpc or one of the various application/grpc+ values. 
	* `is_host_rewrite_enabled` - If true, the hostname will be rewritten to the target virtual deployment's DNS hostname. 
	* `is_path_rewrite_enabled` - If true, the matched path prefix will be rewritten to '/' before being directed to the target virtual deployment. 
	* `path` - Route to match
	* `path_type` - Match type for the route
	* `type` - Type of protocol.
* `state` - The current state of the Resource.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time when this resource was created in an RFC3339 formatted datetime string.
* `time_updated` - The time when this resource was updated in an RFC3339 formatted datetime string.

