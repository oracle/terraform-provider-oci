---
subcategory: "Globally Distributed Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_globally_distributed_database_private_endpoints"
sidebar_current: "docs-oci-datasource-globally_distributed_database-private_endpoints"
description: |-
  Provides the list of Private Endpoints in Oracle Cloud Infrastructure Globally Distributed Database service
---

# Data Source: oci_globally_distributed_database_private_endpoints
This data source provides the list of Private Endpoints in Oracle Cloud Infrastructure Globally Distributed Database service.

List of PrivateEndpoints.


## Example Usage

```hcl
data "oci_globally_distributed_database_private_endpoints" "test_private_endpoints" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.private_endpoint_display_name
	state = var.private_endpoint_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only private endpoint that match the entire name given. The match is not case sensitive.
* `state` - (Optional) A filter to return only resources their lifecycleState matches the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `private_endpoint_collection` - The list of private_endpoint_collection.

### PrivateEndpoint Reference

The following attributes are exported:

* `compartment_id` - Identifier of the compartment in which private endpoint exists.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - PrivateEndpoint description.
* `display_name` - PrivateEndpoint display name.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The identifier of the Private Endpoint.
* `lifecycle_state_details` - Detailed message for the lifecycle state.
* `nsg_ids` - The OCIDs of the network security groups that the private endpoint belongs to. 
* `private_ip` - IP address of the Private Endpoint.
* `proxy_compute_instance_id` - The identifier of the proxy compute instance.
* `sharded_databases` - The OCIDs of sharded databases that consumes the given private endpoint.
* `state` - Lifecycle states for private endpoint.
* `subnet_id` - Identifier of the subnet in which private endpoint exists.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the PrivateEndpoint was first created. An RFC3339 formatted datetime string
* `time_updated` - The time the Private Endpoint was last updated. An RFC3339 formatted datetime string
* `vcn_id` - Identifier of the VCN in which subnet exists.

