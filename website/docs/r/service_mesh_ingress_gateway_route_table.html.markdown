---
subcategory: "Service Mesh"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_service_mesh_ingress_gateway_route_table"
sidebar_current: "docs-oci-resource-service_mesh-ingress_gateway_route_table"
description: |-
  Provides the Ingress Gateway Route Table resource in Oracle Cloud Infrastructure Service Mesh service
---

# oci_service_mesh_ingress_gateway_route_table
This resource provides the Ingress Gateway Route Table resource in Oracle Cloud Infrastructure Service Mesh service.

Creates a new IngressGatewayRouteTable.


## Example Usage

```hcl
resource "oci_service_mesh_ingress_gateway_route_table" "test_ingress_gateway_route_table" {
	#Required
	compartment_id = var.compartment_id
	ingress_gateway_id = oci_service_mesh_ingress_gateway.test_ingress_gateway.id
	name = var.ingress_gateway_route_table_name
	route_rules {
		#Required
		destinations {
			#Required
			virtual_service_id = oci_service_mesh_virtual_service.test_virtual_service.id

			#Optional
			port = var.ingress_gateway_route_table_route_rules_destinations_port
			weight = var.ingress_gateway_route_table_route_rules_destinations_weight
		}
		type = var.ingress_gateway_route_table_route_rules_type

		#Optional
		ingress_gateway_host {
			#Required
			name = var.ingress_gateway_route_table_route_rules_ingress_gateway_host_name

			#Optional
			port = var.ingress_gateway_route_table_route_rules_ingress_gateway_host_port
		}
		is_grpc = var.ingress_gateway_route_table_route_rules_is_grpc
		is_host_rewrite_enabled = var.ingress_gateway_route_table_route_rules_is_host_rewrite_enabled
		is_path_rewrite_enabled = var.ingress_gateway_route_table_route_rules_is_path_rewrite_enabled
		path = var.ingress_gateway_route_table_route_rules_path
		path_type = var.ingress_gateway_route_table_route_rules_path_type
		request_timeout_in_ms = var.ingress_gateway_route_table_route_rules_request_timeout_in_ms
	}

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.ingress_gateway_route_table_description
	freeform_tags = {"bar-key"= "value"}
	priority = var.ingress_gateway_route_table_priority
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) Description of the resource. It can be changed after creation. Avoid entering confidential information.  Example: `This is my new resource` 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `ingress_gateway_id` - (Required) The OCID of the service mesh in which this access policy is created.
* `name` - (Required) A user-friendly name. The name must be unique within the same ingress gateway and cannot be changed after creation. Avoid entering confidential information.  Example: `My unique resource name` 
* `priority` - (Optional) (Updatable) The priority of the route table. Lower value means higher priority. The routes are declared based on the priority.
* `route_rules` - (Required) (Updatable) The route rules for the ingress gateway.
	* `destinations` - (Required) (Updatable) The destination of the request.
		* `port` - (Optional) (Updatable) The port on the virtual service to target. Mandatory if the virtual deployments are listening on multiple ports. 
		* `type` - (Required) (Updatable) Type of the traffic target.
		* `virtual_service_id` - (Required) (Updatable) The OCID of the virtual service where the request will be routed.
		* `weight` - (Optional) (Updatable) Weight of traffic target.
	* `ingress_gateway_host` - (Optional) (Updatable) The ingress gateway host to which the route rule attaches. If not specified, the route rule gets attached to all hosts on the ingress gateway. 
		* `name` - (Required) (Updatable) Name of the ingress gateway host that this route should apply to. 
		* `port` - (Optional) (Updatable) The port of the ingress gateway host listener. Leave empty to match all ports for the host. 
	* `is_grpc` - (Applicable when type=HTTP) (Updatable) If true, the rule will check that the content-type header has a application/grpc or one of the various application/grpc+ values. 
	* `is_host_rewrite_enabled` - (Applicable when type=HTTP) (Updatable) If true, the hostname will be rewritten to the target virtual deployment's DNS hostname. 
	* `is_path_rewrite_enabled` - (Applicable when type=HTTP) (Updatable) If true, the matched path prefix will be rewritten to '/' before being directed to the target virtual deployment. 
	* `path` - (Applicable when type=HTTP) (Updatable) Route to match
	* `path_type` - (Applicable when type=HTTP) (Updatable) Match type for the route
	* `request_timeout_in_ms` - (Applicable when type=HTTP) (Updatable) The maximum duration in milliseconds for the upstream service to respond to a request.  If provided, the timeout value overrides the default timeout of 15 seconds for the HTTP based route rules, and disabled (no timeout) when 'isGrpc' is true.  The value 0 (zero) indicates that the timeout is disabled.  For streaming responses from the upstream service, consider either keeping the timeout disabled or set a sufficiently high value. 
	* `type` - (Required) (Updatable) Type of protocol.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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
	* `request_timeout_in_ms` - The maximum duration in milliseconds for the upstream service to respond to a request.  If provided, the timeout value overrides the default timeout of 15 seconds for the HTTP based route rules, and disabled (no timeout) when 'isGrpc' is true.  The value 0 (zero) indicates that the timeout is disabled.  For streaming responses from the upstream service, consider either keeping the timeout disabled or set a sufficiently high value. 
	* `type` - Type of protocol.
* `state` - The current state of the Resource.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time when this resource was created in an RFC3339 formatted datetime string.
* `time_updated` - The time when this resource was updated in an RFC3339 formatted datetime string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Ingress Gateway Route Table
	* `update` - (Defaults to 20 minutes), when updating the Ingress Gateway Route Table
	* `delete` - (Defaults to 20 minutes), when destroying the Ingress Gateway Route Table


## Import

IngressGatewayRouteTables can be imported using the `id`, e.g.

```
$ terraform import oci_service_mesh_ingress_gateway_route_table.test_ingress_gateway_route_table "id"
```

