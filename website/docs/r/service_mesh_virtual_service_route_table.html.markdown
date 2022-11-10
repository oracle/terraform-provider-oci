---
subcategory: "Service Mesh"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_service_mesh_virtual_service_route_table"
sidebar_current: "docs-oci-resource-service_mesh-virtual_service_route_table"
description: |-
  Provides the Virtual Service Route Table resource in Oracle Cloud Infrastructure Service Mesh service
---

# oci_service_mesh_virtual_service_route_table
This resource provides the Virtual Service Route Table resource in Oracle Cloud Infrastructure Service Mesh service.

Creates a new VirtualServiceRouteTable.


## Example Usage

```hcl
resource "oci_service_mesh_virtual_service_route_table" "test_virtual_service_route_table" {
	#Required
	compartment_id = var.compartment_id
	name = var.virtual_service_route_table_name
	route_rules {
		#Required
		destinations {
			#Required
			virtual_deployment_id = oci_service_mesh_virtual_deployment.test_virtual_deployment.id
			weight = var.virtual_service_route_table_route_rules_destinations_weight

			#Optional
			port = var.virtual_service_route_table_route_rules_destinations_port
		}
		type = var.virtual_service_route_table_route_rules_type

		#Optional
		is_grpc = var.virtual_service_route_table_route_rules_is_grpc
		path = var.virtual_service_route_table_route_rules_path
		path_type = var.virtual_service_route_table_route_rules_path_type
	}
	virtual_service_id = oci_service_mesh_virtual_service.test_virtual_service.id

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.virtual_service_route_table_description
	freeform_tags = {"bar-key"= "value"}
	priority = var.virtual_service_route_table_priority
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) Description of the resource. It can be changed after creation. Avoid entering confidential information.  Example: `This is my new resource` 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `name` - (Required) A user-friendly name. The name must be unique within the same virtual service and cannot be changed after creation. Avoid entering confidential information.  Example: `My unique resource name` 
* `priority` - (Optional) (Updatable) The priority of the route table. Lower value means higher priority. The routes are declared based on the priority.
* `route_rules` - (Required) (Updatable) The route rules for the virtual service.
	* `destinations` - (Required) (Updatable) The destination of the request.
		* `port` - (Optional) (Updatable) Port on virtual deployment to target. If port is missing, the rule will target all ports on the virtual deployment. 
		* `type` - (Required) (Updatable) Type of the traffic target.
		* `virtual_deployment_id` - (Required) (Updatable) The OCID of the virtual deployment where the request will be routed.
		* `weight` - (Required) (Updatable) Weight of traffic target.
	* `is_grpc` - (Applicable when type=HTTP) (Updatable) If true, the rule will check that the content-type header has a application/grpc or one of the various application/grpc+ values. 
	* `path` - (Applicable when type=HTTP) (Updatable) Route to match
	* `path_type` - (Applicable when type=HTTP) (Updatable) Match type for the route
	* `type` - (Required) (Updatable) Type of protocol.
* `virtual_service_id` - (Required) The OCID of the service mesh in which this access policy is created.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Description of the resource. It can be changed after creation. Avoid entering confidential information.  Example: `This is my new resource` 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier that is immutable on creation.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in a Failed state.
* `name` - A user-friendly name. The name must be unique within the same virtual service and cannot be changed after creation. Avoid entering confidential information.  Example: `My unique resource name` 
* `priority` - The priority of the route table. Lower value means higher priority. The routes are declared based on the priority.
* `route_rules` - The route rules for the virtual service.
	* `destinations` - The destination of the request.
		* `port` - Port on virtual deployment to target. If port is missing, the rule will target all ports on the virtual deployment. 
		* `type` - Type of the traffic target.
		* `virtual_deployment_id` - The OCID of the virtual deployment where the request will be routed.
		* `weight` - Weight of traffic target.
	* `is_grpc` - If true, the rule will check that the content-type header has a application/grpc or one of the various application/grpc+ values. 
	* `path` - Route to match
	* `path_type` - Match type for the route
	* `type` - Type of protocol.
* `state` - The current state of the Resource.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time when this resource was created in an RFC3339 formatted datetime string.
* `time_updated` - The time when this resource was updated in an RFC3339 formatted datetime string.
* `virtual_service_id` - The OCID of the virtual service in which this virtual service route table is created.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Virtual Service Route Table
	* `update` - (Defaults to 20 minutes), when updating the Virtual Service Route Table
	* `delete` - (Defaults to 20 minutes), when destroying the Virtual Service Route Table


## Import

VirtualServiceRouteTables can be imported using the `id`, e.g.

```
$ terraform import oci_service_mesh_virtual_service_route_table.test_virtual_service_route_table "id"
```

