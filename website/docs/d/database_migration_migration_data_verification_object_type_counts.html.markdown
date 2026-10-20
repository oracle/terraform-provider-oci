---
subcategory: "Database Migration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_migration_migration_data_verification_object_type_counts"
sidebar_current: "docs-oci-datasource-database_migration-migration_data_verification_object_type_counts"
description: |-
  Provides the list of Migration Data Verification Object Type Counts in Oracle Cloud Infrastructure Database Migration service
---

# Data Source: oci_database_migration_migration_data_verification_object_type_counts
This data source provides the list of Migration Data Verification Object Type Counts in Oracle Cloud Infrastructure Database Migration service.

Lists object type count comparison results for a migration.

The migration must already have Data Verification results.

The `objectType` filter is a database-specific string and is validated by the backend.


## Example Usage

```hcl
data "oci_database_migration_migration_data_verification_object_type_counts" "test_migration_data_verification_object_type_counts" {
	#Required
	compartment_id = var.compartment_id
	migration_id = var.migration_id

	#Optional
	migration_data_verification_object_type_count_filter = null
	min_abs_delta_percent = null
	object_type = null
	schema_name = null
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources. 
* `migration_data_verification_object_type_count_filter` - (Optional) Free-text filter applied by the service to relevant fields for the report. 
* `migration_id` - (Required) The OCID of the migration 
* `min_abs_delta_percent` - (Optional) Minimum absolute deltaPercent threshold (magnitude) to return. The service filters results where `abs(deltaPercent) >= minAbsDeltaPercent`. Must be non-negative.

	Note: This filter applies to the object type counts report, which uses `deltaPercent`. Table row count reports use `variancePercent` instead. 
* `object_type` - (Optional) A filter to return only results for a specific object type.

	The allowed values depend on the migration's `databaseCombination`:
	* Oracle migrations: `OracleDatabaseObjectTypes`
	* MySQL migrations: `MySqlDatabaseObjectTypes` 
* `schema_name` - (Optional) A filter to return only results for a specific schema/owner. 


## Attributes Reference

The following attributes are exported:

* `data_verification_object_type_count_collection` - The list of data_verification_object_type_count_collection.

### MigrationDataVerificationObjectTypeCount Reference

The following attributes are exported:

* `items` - Items in collection.
	* `database_combination` - The combination of source and target databases participating in a migration. Example: ORACLE means the migration is meant for migrating Oracle source and target databases. 
	* `delta_percent` - Percentage delta between source and target counts. This value is expected to be non-negative.

		If a negative value is encountered, it indicates that the target has objects not present in the source, which should be treated as an inconsistency (i.e., a failed comparison/report result). 
	* `object_type` - Oracle database object type.
	* `schema_name` - Schema/owner name.

		This field is omitted (null/empty depending on backend serialization) for database-wide object types that are not schema-scoped.

		Oracle non-schema-scoped object types: USER, ROLE, PROFILE, TABLESPACE, DATABASE_LINK, CONTROLFILE, DATAFILE, REDO_LOG, DIRECTORY, LIBRARY, CONTEXT.

		MySQL non-schema-scoped object types: USER, ROLE, SERVER, TABLESPACE, LOGFILE_GROUP. 
	* `source_invalid_count` - Number of invalid objects in source.
	* `source_object_count` - Number of objects in source.
	* `target_invalid_count` - Number of invalid objects in target.
	* `target_object_count` - Number of objects in target.
* `total_items` - Total number of items matching the request.
