---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_database_sql_plan_baseline_configuration"
sidebar_current: "docs-oci-datasource-database_management-managed_database_sql_plan_baseline_configuration"
description: |-
  Provides details about a specific Managed Database Sql Plan Baseline Configuration in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_database_sql_plan_baseline_configuration
This data source provides details about a specific Managed Database Sql Plan Baseline Configuration resource in Oracle Cloud Infrastructure Database Management service.

Gets the configuration details of SQL plan baselines for the specified
Managed Database. The details include the settings for the capture and use of
SQL plan baselines, SPM Evolve Advisor task, and SQL Management Base.


## Example Usage

```hcl
data "oci_database_management_managed_database_sql_plan_baseline_configuration" "test_managed_database_sql_plan_baseline_configuration" {
	#Required
	managed_database_id = oci_database_management_managed_database.test_managed_database.id

	#Optional
	opc_named_credential_id = var.managed_database_sql_plan_baseline_configuration_opc_named_credential_id
}
```

## Argument Reference

The following arguments are supported:

* `managed_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
* `opc_named_credential_id` - (Optional) The OCID of the Named Credential.


## Attributes Reference

The following attributes are exported:

* `auto_capture_filters` - The capture filters used in automatic initial plan capture.
	* `modified_by` - The database user who last updated the filter value.
	* `name` - The name of the automatic capture filter.
		* AUTO_CAPTURE_SQL_TEXT: Search pattern to apply to SQL text.
		* AUTO_CAPTURE_PARSING_SCHEMA_NAME: Parsing schema to include or exclude for SQL plan management auto capture.
		* AUTO_CAPTURE_MODULE: Module to include or exclude for SQL plan management auto capture.
		* AUTO_CAPTURE_ACTION: Action to include or exclude for SQL plan management automatic capture. 
	* `time_last_modified` - The time the filter value was last updated.
	* `values_to_exclude` - A list of filter values to exclude.
	* `values_to_include` - A list of filter values to include.
* `auto_spm_evolve_task_parameters` - The set of parameters used in an SPM evolve task.
	* `allowed_time_limit` - The global time limit in seconds. This is the total time allowed for the task.
	* `alternate_plan_baselines` - Determines which alternative plans should be loaded.
	* `alternate_plan_limit` - Specifies the maximum number of plans to load in total (that is, not the limit for each SQL statement). A value of zero indicates `UNLIMITED` number of plans. 
	* `alternate_plan_sources` - Determines which sources to search for additional plans.
	* `are_plans_auto_accepted` - Specifies whether to accept recommended plans automatically.
* `is_auto_spm_evolve_task_enabled` - Indicates whether the Automatic SPM Evolve Advisor task is enabled (`true`) or not (`false`).
* `is_automatic_initial_plan_capture_enabled` - Indicates whether the automatic capture of SQL plan baselines is enabled (`true`) or not (`false`).
* `is_high_frequency_auto_spm_evolve_task_enabled` - Indicates whether the high frequency Automatic SPM Evolve Advisor task is enabled (`true`) or not (`false`).
* `is_sql_plan_baselines_usage_enabled` - Indicates whether the database uses SQL plan baselines (`true`) or not (`false`).
* `plan_retention_weeks` - The number of weeks to retain unused plans before they are purged.
* `space_budget_mb` - The maximum `SYSAUX` space that can be used for SQL Management Base in MB.
* `space_budget_percent` - The maximum percent of `SYSAUX` space that can be used for SQL Management Base.
* `space_used_mb` - The space used by SQL Management Base in MB.

