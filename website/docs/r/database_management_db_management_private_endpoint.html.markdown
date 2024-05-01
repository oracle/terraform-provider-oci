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
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.db_management_private_endpoint_description
	freeform_tags = {"Department"= "Finance"}
	is_cluster = var.db_management_private_endpoint_is_cluster
	nsg_ids = var.db_management_private_endpoint_nsg_ids
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) The description of the private endpoint.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `is_cluster` - (Optional) Specifies whether the Database Management private endpoint will be used for Oracle Databases in a cluster.
* `name` - (Required) (Updatable) The display name of the Database Management private endpoint.
* `nsg_ids` - (Optional) (Updatable) The OCIDs of the Network Security Groups to which the Database Management private endpoint belongs. 
* `subnet_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the Database Management private endpoint.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Management private endpoint.
* `is_cluster` - Specifies whether the Database Management private endpoint can be used for Oracle Databases in a cluster.
* `name` - The display name of the Database Management private endpoint.
* `nsg_ids` - The OCIDs of the Network Security Groups to which the Database Management private endpoint belongs. 
* `private_ip` - The IP addresses assigned to the Database Management private endpoint. 
* `state` - The current lifecycle state of the Database Management private endpoint.
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the Database Managament private endpoint was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `vcn_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Db Management Private Endpoint
	* `update` - (Defaults to 20 minutes), when updating the Db Management Private Endpoint
	* `delete` - (Defaults to 20 minutes), when destroying the Db Management Private Endpoint


## Import

DbManagementPrivateEndpoints can be imported using the `id`, e.g.

```
$ terraform import oci_database_management_db_management_private_endpoint.test_db_management_private_endpoint "id"
```

