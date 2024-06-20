---
subcategory: "Globally Distributed Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_globally_distributed_database_private_endpoint"
sidebar_current: "docs-oci-resource-globally_distributed_database-private_endpoint"
description: |-
  Provides the Private Endpoint resource in Oracle Cloud Infrastructure Globally Distributed Database service
---

# oci_globally_distributed_database_private_endpoint
This resource provides the Private Endpoint resource in Oracle Cloud Infrastructure Globally Distributed Database service.

Creates a PrivateEndpoint.


## Example Usage

```hcl
resource "oci_globally_distributed_database_private_endpoint" "test_private_endpoint" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.private_endpoint_display_name
	subnet_id = oci_core_subnet.test_subnet.id

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.private_endpoint_description
	freeform_tags = {"bar-key"= "value"}
	nsg_ids = var.private_endpoint_nsg_ids
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) Identifier of the compartment where private endpoint is to be created.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) PrivateEndpoint description.
* `display_name` - (Required) (Updatable) Private endpoint display name.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `nsg_ids` - (Optional) (Updatable) The OCIDs of the network security groups that the private endpoint belongs to. 
* `subnet_id` - (Required) Identifier of the customer subnet against which private endpoint is to be created.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Private Endpoint
	* `update` - (Defaults to 20 minutes), when updating the Private Endpoint
	* `delete` - (Defaults to 20 minutes), when destroying the Private Endpoint


## Import

PrivateEndpoints can be imported using the `id`, e.g.

```
$ terraform import oci_globally_distributed_database_private_endpoint.test_private_endpoint "id"
```

