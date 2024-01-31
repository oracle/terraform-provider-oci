---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_database_sql_plan_baselines"
sidebar_current: "docs-oci-datasource-database_management-managed_database_sql_plan_baselines"
description: |-
  Provides the list of Managed Database Sql Plan Baselines in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_database_sql_plan_baselines
This data source provides the list of Managed Database Sql Plan Baselines in Oracle Cloud Infrastructure Database Management service.

Lists the SQL plan baselines for the specified Managed Database.


## Example Usage

```hcl
data "oci_database_management_managed_database_sql_plan_baselines" "test_managed_database_sql_plan_baselines" {
	#Required
	managed_database_id = oci_database_management_managed_database.test_managed_database.id

	#Optional
	is_accepted = var.managed_database_sql_plan_baseline_is_accepted
	is_adaptive = var.managed_database_sql_plan_baseline_is_adaptive
	is_auto_purged = var.managed_database_sql_plan_baseline_is_auto_purged
	is_enabled = var.managed_database_sql_plan_baseline_is_enabled
	is_fixed = var.managed_database_sql_plan_baseline_is_fixed
	is_never_executed = var.managed_database_sql_plan_baseline_is_never_executed
	is_reproduced = var.managed_database_sql_plan_baseline_is_reproduced
	opc_named_credential_id = var.managed_database_sql_plan_baseline_opc_named_credential_id
	origin = var.managed_database_sql_plan_baseline_origin
	plan_name = var.managed_database_sql_plan_baseline_plan_name
	sql_handle = var.managed_database_sql_plan_baseline_sql_handle
	sql_text = var.managed_database_sql_plan_baseline_sql_text
	time_last_executed_greater_than = var.managed_database_sql_plan_baseline_time_last_executed_greater_than
	time_last_executed_less_than = var.managed_database_sql_plan_baseline_time_last_executed_less_than
}
```

## Argument Reference

The following arguments are supported:

* `is_accepted` - (Optional) A filter to return only SQL plan baselines that are either accepted or not accepted. By default, all SQL plan baselines are returned. 
* `is_adaptive` - (Optional) A filter to return only SQL plan baselines that are either adaptive or not adaptive. By default, all SQL plan baselines are returned. 
* `is_auto_purged` - (Optional) A filter to return only SQL plan baselines that are either auto-purged or not auto-purged. By default, all SQL plan baselines are returned. 
* `is_enabled` - (Optional) A filter to return only SQL plan baselines that are either enabled or not enabled. By default, all SQL plan baselines are returned. 
* `is_fixed` - (Optional) A filter to return only SQL plan baselines that are either fixed or not fixed. By default, all SQL plan baselines are returned. 
* `is_never_executed` - (Optional) A filter to return only SQL plan baselines that are not executed till now. By default, all SQL plan baselines are returned. 
* `is_reproduced` - (Optional) A filter to return only SQL plan baselines that were either reproduced or not reproduced by the optimizer. By default, all SQL plan baselines are returned. 
* `managed_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
* `opc_named_credential_id` - (Optional) The OCID of the Named Credential.
* `origin` - (Optional) A filter to return all the SQL plan baselines that match the origin.
* `plan_name` - (Optional) A filter to return only SQL plan baselines that match the plan name.
* `sql_handle` - (Optional) A filter to return all the SQL plan baselines for the specified SQL handle.
* `sql_text` - (Optional) A filter to return all the SQL plan baselines that match the SQL text. By default, the search is case insensitive. To run an exact or case-sensitive search, double-quote the search string. You may also use the '%' symbol as a wildcard. 
* `time_last_executed_greater_than` - (Optional) A filter to return only SQL plan baselines whose last execution time is after the specified value. By default, all SQL plan baselines are returned. 
* `time_last_executed_less_than` - (Optional) A filter to return only SQL plan baselines whose last execution time is before the specified value. By default, all SQL plan baselines are returned. 


## Attributes Reference

The following attributes are exported:

* `sql_plan_baseline_collection` - The list of sql_plan_baseline_collection.

### ManagedDatabaseSqlPlanBaseline Reference

The following attributes are exported:

* `accepted` - Indicates whether the plan baseline is accepted (`YES`) or not (`NO`).
* `action` - The application action.
* `adaptive` - Indicates whether a plan that is automatically captured by SQL plan management is marked adaptive or not.

	When a new adaptive plan is found for a SQL statement that has an existing SQL plan baseline, that new plan will be added to the SQL plan baseline as an unaccepted plan, and the `ADAPTIVE` property will be marked `YES`. When this new plan is verified (either manually or via the auto evolve task), the plan will be test executed and the final plan determined at execution will become an accepted plan if its performance is better than the existing plan baseline. At this point, the value of the `ADAPTIVE` property is set to `NO` since the plan is no longer adaptive, but resolved. 
* `auto_purge` - Indicates whether the plan baseline is auto-purged (`YES`) or not (`NO`).
* `enabled` - Indicates whether the plan baseline is enabled (`YES`) or disabled (`NO`).
* `execution_plan` - The execution plan for the SQL statement.
* `fixed` - Indicates whether the plan baseline is fixed (`YES`) or not (`NO`).
* `module` - The application module name.
* `origin` - The origin of the SQL plan baseline.
* `plan_name` - The unique plan identifier.
* `reproduced` - Indicates whether the optimizer was able to reproduce the plan (`YES`) or not (`NO`). The value is set to `YES` when a plan is initially added to the plan baseline. 
* `sql_handle` - The unique SQL identifier.
* `sql_text` - The SQL text.
* `time_created` - The date and time when the plan baseline was created.
* `time_last_executed` - The date and time when the plan baseline was last executed.

	**Note:** For performance reasons, database does not update this value immediately after each execution of the plan baseline. Therefore, the plan baseline may have been executed more recently than this value indicates. 
* `time_last_modified` - The date and time when the plan baseline was last modified.

