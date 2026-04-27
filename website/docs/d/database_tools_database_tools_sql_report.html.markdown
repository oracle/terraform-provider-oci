---
subcategory: "Database Tools"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_tools_database_tools_sql_report"
sidebar_current: "docs-oci-datasource-database_tools-database_tools_sql_report"
description: |-
  Provides details about a specific Database Tools Sql Report in Oracle Cloud Infrastructure Database Tools service
---

# Data Source: oci_database_tools_database_tools_sql_report
This data source provides details about a specific Database Tools Sql Report resource in Oracle Cloud Infrastructure Database Tools service.

Gets details of the specified Database Tools SQL report.

## Example Usage

```hcl
data "oci_database_tools_database_tools_sql_report" "test_database_tools_sql_report" {
	#Required
	database_tools_sql_report_id = oci_database_tools_database_tools_sql_report.test_database_tools_sql_report.id
}
```

## Argument Reference

The following arguments are supported:

* `database_tools_sql_report_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a Database Tools SQL Report.


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

