---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_sql_collection"
sidebar_current: "docs-oci-resource-data_safe-sql_collection"
description: |-
  Provides the Sql Collection resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_sql_collection
This resource provides the Sql Collection resource in Oracle Cloud Infrastructure Data Safe service.

Creates a new SQL collection resource.


## Example Usage

```hcl
resource "oci_data_safe_sql_collection" "test_sql_collection" {
	#Required
	compartment_id = var.compartment_id
	db_user_name = oci_identity_user.test_user.name
	target_id = oci_cloud_guard_target.test_target.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.sql_collection_description
	display_name = var.sql_collection_display_name
	freeform_tags = {"Department"= "Finance"}
	sql_level = var.sql_collection_sql_level
	status = var.sql_collection_status
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment containing the SQL collection.
* `db_user_name` - (Required) The database user name.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) The description of the SQL collection.
* `display_name` - (Optional) (Updatable) The display name of the SQL collection. The name does not have to be unique, and it is changeable.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `sql_level` - (Optional) Specifies the level of SQL that will be collected. USER_ISSUED_SQL - User issued SQL statements only. ALL_SQL - Includes all SQL statements including SQL statement issued inside PL/SQL units. 
* `status` - (Optional) Specifies if the SqlCollection has to be started after creation. Enabled indicates that the SqlCollection will be started after creation.
* `target_id` - (Required) The OCID of the target corresponding to the security policy deployment.
* `generate_sql_firewall_policy_trigger` - (Optional) (Updatable) An optional property when incremented triggers Generate Sql Firewall Policy. Could be set to any integer value.
* `purge_logs_trigger` - (Optional) (Updatable) An optional property when incremented triggers Purge Logs. Could be set to any integer value.
* `refresh_log_insights_trigger` - (Optional) (Updatable) An optional property when incremented triggers Refresh Log Insights. Could be set to any integer value.
* `start_trigger` - (Optional) (Updatable) An optional property when incremented triggers Start. Could be set to any integer value.
* `stop_trigger` - (Optional) (Updatable) An optional property when incremented triggers Stop. Could be set to any integer value.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the SQL collection.
* `db_user_name` - The database user name.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the SQL collection.
* `display_name` - The display name of the SQL collection.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the SQL collection.
* `lifecycle_details` - Details about the current state of the SQL collection in Data Safe.
* `sql_level` - Specifies the level of SQL that will be collected. USER_ISSUED_SQL - User issued SQL statements only. ALL_SQL - Includes all SQL statements including SQL statement issued inside PL/SQL units. 
* `state` - The current state of the SQL collection.
* `status` - Specifies if the status of the SqlCollection. Enabled indicates that the collecting is in progress.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `target_id` - The OCID of the target corresponding to the security policy deployment.
* `time_created` - The time that the SQL collection was created, in the format defined by RFC3339.
* `time_last_started` - The timestamp of the most recent SqlCollection start operation, in the format defined by RFC3339.
* `time_last_stopped` - The timestamp of the most recent SqlCollection stop operation, in the format defined by RFC3339.
* `time_updated` - The last date and time the SQL collection was updated, in the format defined by RFC3339.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Sql Collection
	* `update` - (Defaults to 20 minutes), when updating the Sql Collection
	* `delete` - (Defaults to 20 minutes), when destroying the Sql Collection


## Import

SqlCollections can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_sql_collection.test_sql_collection "id"
```

