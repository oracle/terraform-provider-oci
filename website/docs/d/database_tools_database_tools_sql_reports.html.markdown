---
subcategory: "Database Tools"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_tools_database_tools_sql_reports"
sidebar_current: "docs-oci-datasource-database_tools-database_tools_sql_reports"
description: |-
  Provides the list of Database Tools Sql Reports in Oracle Cloud Infrastructure Database Tools service
---

# Data Source: oci_database_tools_database_tools_sql_reports
This data source provides the list of Database Tools Sql Reports in Oracle Cloud Infrastructure Database Tools service.

Returns a list of Database Tools SQL reports.

## Example Usage

```hcl
data "oci_database_tools_database_tools_sql_reports" "test_database_tools_sql_reports" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.database_tools_sql_report_display_name
	state = var.database_tools_sql_report_state
	type = var.database_tools_sql_report_type
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire specified display name.
* `state` - (Optional) A filter to return resources only when their `databaseToolsSqlReportLifecycleState` matches the specified `databaseToolsSqlReportLifecycleState`.
* `type` - (Optional) A filter to return only resources with one of the specified type values.


## Attributes Reference

The following attributes are exported:

* `database_tools_sql_report_collection` - The list of database_tools_sql_report_collection.

### DatabaseToolsSqlReport Reference

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

