---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_db_management_private_endpoint"
sidebar_current: "docs-oci-resource-database_management-db_management_private_endpoint"
description: |-
  Provides the Db Management Private Endpoint resource in Oracle Cloud Infrastructure Database Management service
---

# oci_database_management_db_management_private_endpoint
This resource provides the Db Management Private Endpoint resource in Oracle Cloud Infrastructure Database Management service.

Creates a new Database Management private endpoint.


## Example Usage

```hcl
resource "oci_database_management_db_management_private_endpoint" "test_db_management_private_endpoint" {
	#Required
	compartment_id = var.compartment_id
	name = var.db_management_private_endpoint_name
	subnet_id = oci_core_subnet.test_subnet.id

	#Optional
	description = var.db_management_private_endpoint_description
	nsg_ids = var.db_management_private_endpoint_nsg_ids
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment.
* `description` - (Optional) (Updatable) The description of the private endpoint.
* `name` - (Required) (Updatable) The display name for the private endpoint. It is changeable.
* `nsg_ids` - (Optional) (Updatable) The OCIDs of the network security groups that the private endpoint belongs to. 
* `subnet_id` - (Required) The OCID of the subnet.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment.
* `description` - The description of the private endpoint.
* `id` - The OCID of the Database Management private endpoint.
* `name` - The display name of the private endpoint.
* `nsg_ids` - The OCIDs of the network security groups that the private endpoint belongs to. 
* `private_ip` - The private IP addresses assigned to the private endpoint. 
* `state` - The current state of the private endpoint.
* `subnet_id` - The OCID of the subnet.
* `time_created` - The date and time the private endpoint was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `vcn_id` - The OCID of the VCN.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Db Management Private Endpoint
	* `update` - (Defaults to 20 minutes), when updating the Db Management Private Endpoint
	* `delete` - (Defaults to 20 minutes), when destroying the Db Management Private Endpoint


## Import

DbManagementPrivateEndpoints can be imported using the `id`, e.g.

```
$ terraform import oci_database_management_db_management_private_endpoint.test_db_management_private_endpoint "id"
```

