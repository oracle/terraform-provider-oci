---
subcategory: "Database Migration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_migration_connection_databaseconnectiontypes"
sidebar_current: "docs-oci-datasource-database_migration-connection_databaseconnectiontypes"
description: |-
  Provides the list of Connection Databaseconnectiontypes in Oracle Cloud Infrastructure Database Migration service
---

# Data Source: oci_database_migration_connection_databaseconnectiontypes
This data source provides the list of Connection Databaseconnectiontypes in Oracle Cloud Infrastructure Database Migration service.

List supported Database Types, Sub-types and Versions.

## Example Usage

```hcl
data "oci_database_migration_connection_databaseconnectiontypes" "test_connection_databaseconnectiontypes" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	connection_type = var.connection_databaseconnectiontype_connection_type
	source_connection_id = oci_database_migration_connection.test_connection.id
	technology_type = var.connection_databaseconnectiontype_technology_type
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources. 
* `connection_type` - (Optional) The array of connection types.
* `source_connection_id` - (Optional) The OCID of the source connection.
* `technology_type` - (Optional) The array of technology types.


## Attributes Reference

The following attributes are exported:

* `database_connection_type_collection` - The list of database_connection_type_collection.

### ConnectionDatabaseconnectiontype Reference

The following attributes are exported:

* `items` - Items in collection.
	* `connection_type` - Defines the type of connection. For example, ORACLE.
	* `technology_types` - Array of technology type objects
		* `database_versions` - Array of database versions
		* `technology_sub_types` - Array of technology sub-types e.g. ADW_SHARED.
			* `database_versions` - Array of database versions
			* `technology_sub_type` - Technology sub-type e.g. ADW_SHARED.
			* `technology_sub_type_display_name` - A user-friendly name. Does not have to be unique, and it's changeable.  Avoid entering confidential information. 
		* `technology_type` - The technology type.

