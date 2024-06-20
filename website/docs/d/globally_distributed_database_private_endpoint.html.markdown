---
subcategory: "Globally Distributed Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_globally_distributed_database_private_endpoint"
sidebar_current: "docs-oci-datasource-globally_distributed_database-private_endpoint"
description: |-
  Provides details about a specific Private Endpoint in Oracle Cloud Infrastructure Globally Distributed Database service
---

# Data Source: oci_globally_distributed_database_private_endpoint
This data source provides details about a specific Private Endpoint resource in Oracle Cloud Infrastructure Globally Distributed Database service.

Get the PrivateEndpoint resource.


## Example Usage

```hcl
data "oci_globally_distributed_database_private_endpoint" "test_private_endpoint" {
	#Required
	private_endpoint_id = oci_globally_distributed_database_private_endpoint.test_private_endpoint.id
}
```

## Argument Reference

The following arguments are supported:

* `private_endpoint_id` - (Required) Oracle Sharded Database PrivateEndpoint identifier


## Attributes Reference

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
* `sharded_databases` - The OCIDs of sharded databases that consumes the given private endpoint.
* `state` - Lifecycle states for private endpoint.
* `subnet_id` - Identifier of the subnet in which private endpoint exists.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the PrivateEndpoint was first created. An RFC3339 formatted datetime string
* `time_updated` - The time the Private Endpoint was last updated. An RFC3339 formatted datetime string
* `vcn_id` - Identifier of the VCN in which subnet exists.

