---
subcategory: "Service Mesh"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_service_mesh_virtual_service_route_tables"
sidebar_current: "docs-oci-datasource-service_mesh-virtual_service_route_tables"
description: |-
  Provides the list of Virtual Service Route Tables in Oracle Cloud Infrastructure Service Mesh service
---

# Data Source: oci_service_mesh_virtual_service_route_tables
This data source provides the list of Virtual Service Route Tables in Oracle Cloud Infrastructure Service Mesh service.

Returns a list of VirtualServiceRouteTable objects.


## Example Usage

```hcl
data "oci_service_mesh_virtual_service_route_tables" "test_virtual_service_route_tables" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	id = var.virtual_service_route_table_id
	name = var.virtual_service_route_table_name
	state = var.virtual_service_route_table_state
	virtual_service_id = oci_service_mesh_virtual_service.test_virtual_service.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `id` - (Optional) Unique VirtualServiceRouteTable identifier.
* `name` - (Optional) A filter to return only resources that match the entire name given.
* `state` - (Optional) A filter to return only resources that match the life cycle state given.
* `virtual_service_id` - (Optional) Unique VirtualService identifier.


## Attributes Reference

The following attributes are exported:

* `virtual_service_route_table_collection` - The list of virtual_service_route_table_collection.

### VirtualServiceRouteTable Reference

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
	* `request_timeout_in_ms` - The maximum duration in milliseconds for the target service to respond to a request.  If provided, the timeout value overrides the default timeout of 15 seconds for the HTTP based route rules, and disabled (no timeout) when 'isGrpc' is true.  The value 0 (zero) indicates that the timeout is disabled.  For streaming responses from the target service, consider either keeping the timeout disabled or set a sufficiently high value. 
	* `type` - Type of protocol.
* `state` - The current state of the Resource.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time when this resource was created in an RFC3339 formatted datetime string.
* `time_updated` - The time when this resource was updated in an RFC3339 formatted datetime string.
* `virtual_service_id` - The OCID of the virtual service in which this virtual service route table is created.

