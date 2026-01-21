---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_db_connection_bundles"
sidebar_current: "docs-oci-datasource-database-db_connection_bundles"
description: |-
  Provides the list of Db Connection Bundles in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_db_connection_bundles
This data source provides the list of Db Connection Bundles in Oracle Cloud Infrastructure Database service.

Lists all database connection bundles that match the query parameters.


## Example Usage

```hcl
data "oci_database_db_connection_bundles" "test_db_connection_bundles" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	associated_resource_id = oci_cloud_guard_resource.test_resource.id
	db_connection_bundle_type = var.db_connection_bundle_db_connection_bundle_type
	display_name = var.db_connection_bundle_display_name
	state = var.db_connection_bundle_state
}
```

## Argument Reference

The following arguments are supported:

* `associated_resource_id` - (Optional) The OCID of the VM cluster associated with the connection bundle. If the parameter is set to null, all bundles are returned.
* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `db_connection_bundle_type` - (Optional) A filter that returns only resources that match the specified database connection bundle type.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. The match is not case sensitive.
* `state` - (Optional) A filter that returns only resources that match the given lifecycle state. The state value is case-insensitive.


## Attributes Reference

The following attributes are exported:

* `db_connection_bundles` - The list of db_connection_bundles.

### DbConnectionBundle Reference

The following attributes are exported:

* `associated_resource_details` - Details about the resources associated with the connection bundle.
	* `resource_ids` - The OCIDs of the associated resources.
* `compartment_id` - The OCID of the compartment containing the database connection bundle.
* `db_connection_bundle_type` - The type of the database connection bundle.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
* `display_name` - Display name for the connection bundle.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
* `id` - The OCID of the database connection bundle.
* `is_protected` - True for the default, service-created Database Connection Bundle.
* `state` - The current lifecycle state of the database connection bundle.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
* `time_created` - The time the database connection bundle was created. An RFC3339 formatted datetime string.
* `time_last_refreshed` - The time the database connection bundle was last refreshed. An RFC3339 formatted datetime string.
* `time_updated` - The time the database connection bundle was updated. An RFC3339 formatted datetime string.

