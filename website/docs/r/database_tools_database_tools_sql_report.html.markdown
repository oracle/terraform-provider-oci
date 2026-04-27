---
subcategory: "Database Tools"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_tools_database_tools_sql_report"
sidebar_current: "docs-oci-resource-database_tools-database_tools_sql_report"
description: |-
  Provides the Database Tools Sql Report resource in Oracle Cloud Infrastructure Database Tools service
---

# oci_database_tools_database_tools_sql_report
This resource provides the Database Tools Sql Report resource in Oracle Cloud Infrastructure Database Tools service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/database-tools/latest/DatabaseToolsSqlReport

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/databaseTools

Creates a new Database Tools  Sql Report.


## Example Usage

```hcl
resource "oci_database_tools_database_tools_sql_report" "test_database_tools_sql_report" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.database_tools_sql_report_display_name
	source = var.database_tools_sql_report_source
	type = var.database_tools_sql_report_type

	#Optional
	columns {
		#Required
		description = var.database_tools_sql_report_columns_description
		name = var.database_tools_sql_report_columns_name
		type = var.database_tools_sql_report_columns_type
	}
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.database_tools_sql_report_description
	freeform_tags = {"bar-key"= "value"}
	instructions = var.database_tools_sql_report_instructions
	locks {
		#Required
		type = var.database_tools_sql_report_locks_type

		#Optional
		message = var.database_tools_sql_report_locks_message
		related_resource_id = oci_cloud_guard_resource.test_resource.id
		time_created = var.database_tools_sql_report_locks_time_created
	}
	purpose = var.database_tools_sql_report_purpose
	variables {
		#Required
		description = var.database_tools_sql_report_variables_description
		name = var.database_tools_sql_report_variables_name
		type = var.database_tools_sql_report_variables_type
	}
}
```

## Argument Reference

The following arguments are supported:

* `columns` - (Optional) (Updatable) Descriptive information on columns referenced in the Database Tools SQL Report source.
	* `description` - (Required) (Updatable) The description of the column
	* `name` - (Required) (Updatable) The name of the column
	* `type` - (Required) (Updatable) The type of the column
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the Database Tools SQL report.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) A description of the SQL report.
* `display_name` - (Required) (Updatable) A user-friendly name. Does not have to be unique and can be updated. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `instructions` - (Optional) (Updatable) Instructions on how to use the SQL report. Step-by-step guidance for an AI agent on how to execute or fill in parameters for the report.
* `locks` - (Optional) Locks associated with this resource.
	* `message` - (Optional) A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked. 
	* `related_resource_id` - (Optional) The id of the resource that is locking this resource. Indicates that deleting this resource will remove the lock. 
	* `time_created` - (Optional) When the lock was created.
	* `type` - (Required) Type of the lock.
* `purpose` - (Optional) (Updatable) Purpose of the Database Tools SQL report. Scenario or conditions describing when or why this report should be used. Provides selection criteria to AI agents to improve report selection accuracy.
* `source` - (Required) (Updatable) SQL query executed to generate the report.
* `type` - (Required) (Updatable) The Database Tools SQL report type.
* `variables` - (Optional) (Updatable) Variables referenced in the Database Tools SQL Report source.
	* `description` - (Required) (Updatable) The description of the variable
	* `name` - (Required) (Updatable) The name of the variable
	* `type` - (Required) (Updatable) The type of the variable


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `columns` - Descriptive information on columns referenced in the Database Tools SQL Report source.
	* `description` - The description of the column
	* `name` - The name of the column
	* `type` - The type of the column
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the Database Tools SQL report.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A description of the SQL report.
* `display_name` - A user-friendly name. Does not have to be unique and can be updated. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Tools SQL report.
* `instructions` - Instructions on how to use the SQL report. Step-by-step guidance for an AI agent on how to execute or fill in parameters for the report.
* `lifecycle_details` - A message describing the current state in more detail.
* `locks` - Locks associated with this resource.
	* `message` - A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked. 
	* `related_resource_id` - The id of the resource that is locking this resource. Indicates that deleting this resource will remove the lock. 
	* `time_created` - When the lock was created.
	* `type` - Type of the lock.
* `purpose` - Purpose of the Database Tools SQL report. Scenario or conditions describing when or why this report should be used. Provides selection criteria to AI agents to improve report selection accuracy.
* `source` - SQL query executed to generate the report.
* `state` - The current state of the Database Tools SQL report.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the Database Tools SQL report was created. An RFC3339 formatted datetime string.
* `time_updated` - The time the Database Tools SQL report was updated. An RFC3339 formatted datetime string.
* `type` - The Database Tools SQL report type.
* `variables` - Variables referenced in the Database Tools SQL Report source.
	* `description` - The description of the variable
	* `name` - The name of the variable
	* `type` - The type of the variable

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Database Tools Sql Report
	* `update` - (Defaults to 20 minutes), when updating the Database Tools Sql Report
	* `delete` - (Defaults to 20 minutes), when destroying the Database Tools Sql Report


## Import

DatabaseToolsSqlReports can be imported using the `id`, e.g.

```
$ terraform import oci_database_tools_database_tools_sql_report.test_database_tools_sql_report "id"
```

