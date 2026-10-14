---
subcategory: "Database Migration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_migration_migration_data_verification_table_row_counts"
sidebar_current: "docs-oci-datasource-database_migration-migration_data_verification_table_row_counts"
description: |-
  Provides the list of Migration Data Verification Table Row Counts in Oracle Cloud Infrastructure Database Migration service
---

# Data Source: oci_database_migration_migration_data_verification_table_row_counts
This data source provides the list of Migration Data Verification Table Row Counts in Oracle Cloud Infrastructure Database Migration service.

Lists table row count comparison results for a migration.

The migration must already have Data Verification results.


## Example Usage

```hcl
data "oci_database_migration_migration_data_verification_table_row_counts" "test_migration_data_verification_table_row_counts" {
	#Required
	compartment_id = var.compartment_id
	migration_id = var.migration_id

	#Optional
	migration_data_verification_table_row_count_filter = null
	min_abs_delta_percent = null
	owner = null
	table_name = null
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources. 
* `migration_data_verification_table_row_count_filter` - (Optional) Free-text filter applied by the service to relevant fields for the report. 
* `migration_id` - (Required) The OCID of the migration 
* `min_abs_delta_percent` - (Optional) Minimum absolute deltaPercent threshold (magnitude) to return. The service filters results where `abs(deltaPercent) >= minAbsDeltaPercent`. Must be non-negative.

	Note: This filter applies to the object type counts report, which uses `deltaPercent`. Table row count reports use `variancePercent` instead. 
* `owner` - (Optional) A filter to return only results for a specific owner. 
* `table_name` - (Optional) A filter to return only results for a specific table. 


## Attributes Reference

The following attributes are exported:

* `data_verification_table_row_count_collection` - The list of data_verification_table_row_count_collection.

### MigrationDataVerificationTableRowCount Reference

The following attributes are exported:

* `items` - Items in collection.
	* `database_combination` - The combination of source and target databases participating in a migration. Example: ORACLE means the migration is meant for migrating Oracle source and target databases. 
	* `owner` - Schema/owner name.

		This field is omitted (null/empty depending on backend serialization) for database-wide object types that are not schema-scoped.

		Oracle non-schema-scoped object types: USER, ROLE, PROFILE, TABLESPACE, DATABASE_LINK, CONTROLFILE, DATAFILE, REDO_LOG, DIRECTORY, LIBRARY, CONTEXT.

		MySQL non-schema-scoped object types: USER, ROLE, SERVER, TABLESPACE, LOGFILE_GROUP. 
	* `source_row_count` - Row count in source.
	* `table_name` - Table name.
	* `target_row_count` - Row count in target.
	* `time_last_source_statistics_collection` - Time of last statistics collection in the source used for the dictionary-based row count estimate. This value is reported per row because statistics can be collected at different times per object. This timestamp is stored in UTC. 
	* `time_last_target_statistics_collection` - Time of last statistics collection in the target used for the dictionary-based row count estimate. This value is reported per row because statistics can be collected at different times per object. This timestamp is stored in UTC. 
	* `variance_percent` - Percentage variance between source and target row counts. This value is expected to be non-negative.

		A negative value indicates that the target has rows not present in the source, which should be treated as an inconsistency (i.e., a failed comparison/report result). 
* `total_items` - Total number of items matching the request.
