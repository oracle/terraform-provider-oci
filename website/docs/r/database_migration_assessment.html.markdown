---
subcategory: "Database Migration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_migration_assessment"
sidebar_current: "docs-oci-resource-database_migration-assessment"
description: |-
  Provides the Assessment resource in Oracle Cloud Infrastructure Database Migration service
---

# oci_database_migration_assessment
This resource provides the Assessment resource in Oracle Cloud Infrastructure Database Migration service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/database-migration/latest/Assessment

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/databasemigration

Create an Assessment resource that contains all the details to perform the
database assessment operation, such as source and destination database
details, network throughput, accepted downtime etc.


## Example Usage

```hcl
resource "oci_database_migration_assessment" "test_assessment" {
	#Required
	acceptable_downtime = var.assessment_acceptable_downtime
	compartment_id = var.compartment_id
	database_combination = var.assessment_database_combination
	database_data_size = var.assessment_database_data_size
	ddl_expectation = var.assessment_ddl_expectation
	network_speed_megabit_per_second = var.assessment_network_speed_megabit_per_second
	source_database_connection {
		#Required
		id = var.assessment_source_database_connection_id
	}
	target_database_connection {

		#Optional
		connection_type = var.assessment_target_database_connection_connection_type
		database_version = var.assessment_target_database_connection_database_version
		id = var.assessment_target_database_connection_id
		technology_sub_type = var.assessment_target_database_connection_technology_sub_type
		technology_type = var.assessment_target_database_connection_technology_type
	}

	#Optional
	bulk_include_exclude_data = var.assessment_bulk_include_exclude_data
	creation_type = var.assessment_creation_type
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.assessment_description
	display_name = var.assessment_display_name
	exclude_objects {
		#Required
		object = var.assessment_exclude_objects_object

		#Optional
		is_omit_excluded_table_from_replication = var.assessment_exclude_objects_is_omit_excluded_table_from_replication
		owner = var.assessment_exclude_objects_owner
		schema = var.assessment_exclude_objects_schema
		type = var.assessment_exclude_objects_type
	}
	freeform_tags = var.assessment_freeform_tags
	include_objects {
		#Required
		object = var.assessment_include_objects_object

		#Optional
		is_omit_excluded_table_from_replication = var.assessment_include_objects_is_omit_excluded_table_from_replication
		owner = var.assessment_include_objects_owner
		schema = var.assessment_include_objects_schema
		type = var.assessment_include_objects_type
	}
}
```

## Argument Reference

The following arguments are supported:

* `acceptable_downtime` - (Required) (Updatable) Time allowed for the application downtime.
* `bulk_include_exclude_data` - (Optional) Specifies the database objects to be excluded from the migration in bulk. The definition accepts input in a CSV format, newline separated for each entry. More details can be found in the documentation. 
* `compartment_id` - (Required) (Updatable) The OCID of the resource being referenced.
* `creation_type` - (Optional) (Updatable) The type of assessment creation.
* `database_combination` - (Required) (Updatable) The combination of source and target databases participating in a migration. Example: ORACLE means the migration is meant for migrating Oracle source and target databases. 
* `database_data_size` - (Required) (Updatable) The size of a source database.
* `ddl_expectation` - (Required) (Updatable) DDL expectation values.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) A user-friendly description. Does not have to be unique, and it's changeable.  Avoid entering confidential information. 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable.  Avoid entering confidential information. 
* `exclude_objects` - (Optional) Database objects to exclude from migration, cannot be specified alongside 'includeObjects' 
	* `is_omit_excluded_table_from_replication` - (Applicable when database_combination=ORACLE) Whether an excluded table should be omitted from replication. Only valid for database objects  that have are of type TABLE and object status EXCLUDE. 
	* `object` - (Required) Name of the object (regular expression is allowed)
	* `owner` - (Required when database_combination=ORACLE) Owner of the object (regular expression is allowed)
	* `schema` - (Required when database_combination=MYSQL) Schema of the object (regular expression is allowed)
	* `type` - (Optional) Type of object to exclude. If not specified, matching owners and object names of type TABLE would be excluded. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.  For more information, see Resource Tags. Example: {"Department": "Finance"} 
* `include_objects` - (Optional) Database objects to include from migration, cannot be specified alongside 'excludeObjects'
	* `is_omit_excluded_table_from_replication` - (Applicable when database_combination=ORACLE) Whether an excluded table should be omitted from replication. Only valid for database objects  that have are of type TABLE and object status EXCLUDE. 
	* `object` - (Required) Name of the object (regular expression is allowed)
	* `owner` - (Required when database_combination=ORACLE) Owner of the object (regular expression is allowed)
	* `schema` - (Required when database_combination=MYSQL) Schema of the object (regular expression is allowed)
	* `type` - (Optional) Type of object to exclude. If not specified, matching owners and object names of type TABLE would be excluded. 
* `network_speed_megabit_per_second` - (Required) (Updatable) A network speed in Megabits per second.
* `source_database_connection` - (Required) (Updatable) Source Assessment Connection object
	* `id` - (Required) (Updatable) The OCID of the resource being referenced.
* `target_database_connection` - (Required) (Updatable) Target Assessment Connection object
	* `connection_type` - (Optional) (Updatable) Defines the type of connection. For example, ORACLE.
	* `database_version` - (Optional) (Updatable) The database version
	* `id` - (Optional) (Updatable) The OCID of the resource being referenced.
	* `technology_sub_type` - (Optional) (Updatable) Technology sub-type e.g. ADW_SHARED, ADW_DEDICATED, ATP_SHARED, ATP_DEDICATED
	* `technology_type` - (Optional) (Updatable) The technology type.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Assessment
	* `update` - (Defaults to 20 minutes), when updating the Assessment
	* `delete` - (Defaults to 20 minutes), when destroying the Assessment


## Import

Assessments can be imported using the `id`, e.g.

```
$ terraform import oci_database_migration_assessment.test_assessment "id"
```

