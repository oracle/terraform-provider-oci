---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_sql_collections"
sidebar_current: "docs-oci-datasource-data_safe-sql_collections"
description: |-
  Provides the list of Sql Collections in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_sql_collections
This data source provides the list of Sql Collections in Oracle Cloud Infrastructure Data Safe service.

Retrieves a list of all SQL collections in Data Safe.

The ListSqlCollections operation returns only the SQL collections in the specified `compartmentId`.

The parameter `accessLevel` specifies whether to return only those compartments for which the
requestor has INSPECT permissions on at least one resource directly
or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
Principal doesn't have access to even one of the child compartments. This is valid only when
`compartmentIdInSubtree` is set to `true`.

The parameter `compartmentIdInSubtree` applies when you perform ListSqlCollections on the
`compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
To get a full list of all compartments and subcompartments in the tenancy (root compartment),
set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.


## Example Usage

```hcl
data "oci_data_safe_sql_collections" "test_sql_collections" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.sql_collection_access_level
	compartment_id_in_subtree = var.sql_collection_compartment_id_in_subtree
	db_user_name = oci_identity_user.test_user.name
	display_name = var.sql_collection_display_name
	sql_collection_id = oci_data_safe_sql_collection.test_sql_collection.id
	state = var.sql_collection_state
	target_id = oci_cloud_guard_target.test_target.id
	time_created_greater_than_or_equal_to = var.sql_collection_time_created_greater_than_or_equal_to
	time_created_less_than = var.sql_collection_time_created_less_than
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `db_user_name` - (Optional) A filter to return only items that match the specified user name.
* `display_name` - (Optional) A filter to return only resources that match the specified display name. 
* `sql_collection_id` - (Optional) An optional filter to return only resources that match the specified OCID of the SQL collection resource.
* `state` - (Optional) The current state of the SQL collection.
* `target_id` - (Optional) A filter to return only items related to a specific target OCID.
* `time_created_greater_than_or_equal_to` - (Optional) A filter to return only the resources that were created after the specified date and time, as defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Using TimeCreatedGreaterThanOrEqualToQueryParam parameter retrieves all resources created after that date.

	**Example:** 2016-12-19T16:39:57.600Z 
* `time_created_less_than` - (Optional) Search for resources that were created before a specific date. Specifying this parameter corresponding `timeCreatedLessThan` parameter will retrieve all resources created before the specified created date, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by RFC 3339.

	**Example:** 2016-12-19T16:39:57.600Z 


## Attributes Reference

The following attributes are exported:

* `sql_collection_collection` - The list of sql_collection_collection.

### SqlCollection Reference

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

