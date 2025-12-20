---
subcategory: "Database Migration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_migration_assessments"
sidebar_current: "docs-oci-datasource-database_migration-assessments"
description: |-
  Provides the list of Assessments in Oracle Cloud Infrastructure Database Migration service
---

# Data Source: oci_database_migration_assessments
This data source provides the list of Assessments in Oracle Cloud Infrastructure Database Migration service.

List all Assessments.


## Example Usage

```hcl
data "oci_database_migration_assessments" "test_assessments" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.assessment_display_name
	lifecycle_details = var.assessment_lifecycle_details
	state = var.assessment_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources. 
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. 
* `lifecycle_details` - (Optional) The lifecycle detailed status of the Migration. 
* `state` - (Optional) The lifecycle state of the Assessment. 


## Attributes Reference

The following attributes are exported:

* `assessment_collection` - The list of assessment_collection.

### Assessment Reference

The following attributes are exported:

* `acceptable_downtime` - Time allowed for the application downtime.
* `assessment_migration_type` - The migration type of the migration to be performed.
* `compartment_id` - The OCID of the resource being referenced.
* `creation_type` - The type of assessment creation.
* `database_combination` - The combination of source and target databases participating in a migration. Example: ORACLE means the migration is meant for migrating Oracle source and target databases. 
* `database_data_size` - The size of a source database.
* `ddl_expectation` - DDL expectation values.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A user-friendly description. Does not have to be unique, and it's changeable.  Avoid entering confidential information. 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.  Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.  For more information, see Resource Tags. Example: {"Department": "Finance"} 
* `id` - The OCID of the resource being referenced.
* `is_cdb_supported` - True if CDB should be defined, false otherwise.
* `migration_id` - The OCID of the resource being referenced.
* `network_speed_megabit_per_second` - A network speed in Megabits per second.
* `source_database_connection` - Source Assessment Connection object
	* `id` - The OCID of the resource being referenced.
* `state` - The current state of the Assessment resource.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `target_database_connection` - Target Assessment Connection object
	* `connection_type` - Defines the type of connection. For example, ORACLE.
	* `database_version` - The database version
	* `id` - The OCID of the resource being referenced.
	* `technology_sub_type` - Technology sub-type e.g. ADW_SHARED, ADW_DEDICATED, ATP_SHARED, ATP_DEDICATED
	* `technology_type` - The technology type.
* `time_created` - An RFC3339 formatted datetime string such as `2016-08-25T21:10:29.600Z`. 
* `time_updated` - An RFC3339 formatted datetime string such as `2016-08-25T21:10:29.600Z`. 

