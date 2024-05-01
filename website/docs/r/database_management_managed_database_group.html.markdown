---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_database_group"
sidebar_current: "docs-oci-resource-database_management-managed_database_group"
description: |-
  Provides the Managed Database Group resource in Oracle Cloud Infrastructure Database Management service
---

# oci_database_management_managed_database_group
This resource provides the Managed Database Group resource in Oracle Cloud Infrastructure Database Management service.

Creates a Managed Database Group. The group does not contain any
Managed Databases when it is created, and they must be added later.


## Example Usage

```hcl
resource "oci_database_management_managed_database_group" "test_managed_database_group" {
	#Required
	compartment_id = var.compartment_id
	name = var.managed_database_group_name

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.managed_database_group_description
	freeform_tags = {"Department"= "Finance"}
        managed_databases {
          id = var.managed_database_id
        }
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the Managed Database Group resides. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) The information specified by the user about the Managed Database Group.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `name` - (Required) The name of the Managed Database Group. Valid characters are uppercase or lowercase letters, numbers, and "_". The name of the Managed Database Group cannot be modified. It must be unique in the compartment and must begin with an alphabetic character. 
* `managed_databases` - (Optional) (Updatable) Set of Managed Databases that the user wants to add to the Managed Database Group. Specifying a block will add the Managed Database to Managed Database Group and removing the block will remove Managed Database from the Managed Database Group.
    * `id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed database that needs to be added to the Managed Database Group. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - The information specified by the user about the Managed Database Group.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database Group.
* `managed_databases` - A list of Managed Databases in the Managed Database Group.
	* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the Managed Database resides.
	* `database_sub_type` - The subtype of the Oracle Database. Indicates whether the database is a Container Database, Pluggable Database, Non-container Database, Autonomous Database, or Autonomous Container Database. 
	* `database_type` - The type of Oracle Database installation.
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
	* `deployment_type` - The infrastructure used to deploy the Oracle Database.
	* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
	* `name` - The name of the Managed Database.
	* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
	* `time_added` - The date and time the Managed Database was added to the group.
	* `workload_type` - The workload type of the Autonomous Database.
* `name` - The name of the Managed Database Group.
* `state` - The current lifecycle state of the Managed Database Group.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the Managed Database Group was created.
* `time_updated` - The date and time the Managed Database Group was last updated.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Managed Database Group
	* `update` - (Defaults to 20 minutes), when updating the Managed Database Group
	* `delete` - (Defaults to 20 minutes), when destroying the Managed Database Group


## Import

ManagedDatabaseGroups can be imported using the `id`, e.g.

```
$ terraform import oci_database_management_managed_database_group.test_managed_database_group "id"
```

