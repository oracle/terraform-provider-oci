---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_sql_collection"
sidebar_current: "docs-oci-datasource-data_safe-sql_collection"
description: |-
  Provides details about a specific Sql Collection in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_sql_collection
This data source provides details about a specific Sql Collection resource in Oracle Cloud Infrastructure Data Safe service.

Gets a SQL collection by identifier.

## Example Usage

```hcl
data "oci_data_safe_sql_collection" "test_sql_collection" {
	#Required
	sql_collection_id = oci_data_safe_sql_collection.test_sql_collection.id
}
```

## Argument Reference

The following arguments are supported:

* `sql_collection_id` - (Required) The OCID of the SQL collection resource.


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

