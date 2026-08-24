---
subcategory: "Database Migration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_migration_migration_data_verification_object_statuses"
sidebar_current: "docs-oci-datasource-database_migration-migration_data_verification_object_statuses"
description: |-
  Provides the list of Migration Data Verification Object Statuses in Oracle Cloud Infrastructure Database Migration service
---

# Data Source: oci_database_migration_migration_data_verification_object_statuses
This data source provides the list of Migration Data Verification Object Statuses in Oracle Cloud Infrastructure Database Migration service.

Lists per-object status comparison results for a migration.

The migration must already have Data Verification results.

The `objectType` filter is a database-specific string and is validated by the backend.


## Example Usage

```hcl
data "oci_database_migration_migration_data_verification_object_statuses" "test_migration_data_verification_object_statuses" {
	#Required
	compartment_id = var.compartment_id
	migration_id = var.migration_id

	#Optional
	is_match = null
	migration_data_verification_object_status_filter = null
	name = null
	object_type = null
	owner = null
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources. 
* `is_match` - (Optional) Match filter for object status results.

	When `true`, returns only rows where source and target match. When `false`, returns only rows with mismatches. When omitted, returns all rows, ordered by mismatch first. 
* `migration_data_verification_object_status_filter` - (Optional) Free-text filter applied by the service to relevant fields for the report. 
* `migration_id` - (Required) The OCID of the migration 
* `name` - (Optional) A filter to return only resources that match the entire name given. 
* `object_type` - (Optional) A filter to return only results for a specific object type.

	The allowed values depend on the migration's `databaseCombination`:
	* Oracle migrations: `OracleDatabaseObjectTypes`
	* MySQL migrations: `MySqlDatabaseObjectTypes` 
* `owner` - (Optional) A filter to return only results for a specific owner. 


## Attributes Reference

The following attributes are exported:

* `data_verification_object_status_collection` - The list of data_verification_object_status_collection.

### MigrationDataVerificationObjectStatus Reference

The following attributes are exported:

* `items` - Items in collection.
	* `database_combination` - The combination of source and target databases participating in a migration. Example: ORACLE means the migration is meant for migrating Oracle source and target databases. 
	* `object` - Database object name.
	* `object_type` - Oracle database object type.
	* `owner` - Schema/owner name.

		This field is omitted (null/empty depending on backend serialization) for database-wide object types that are not schema-scoped.

		Oracle non-schema-scoped object types: USER, ROLE, PROFILE, TABLESPACE, DATABASE_LINK, CONTROLFILE, DATAFILE, REDO_LOG, DIRECTORY, LIBRARY, CONTEXT.

		MySQL non-schema-scoped object types: USER, ROLE, SERVER, TABLESPACE, LOGFILE_GROUP. 
	* `status_in_source` - Object status in source.
	* `status_in_target` - Object status in target.
* `total_items` - Total number of items matching the request.
