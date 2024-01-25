---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_database_sql_plan_baseline"
sidebar_current: "docs-oci-datasource-database_management-managed_database_sql_plan_baseline"
description: |-
  Provides details about a specific Managed Database Sql Plan Baseline in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_database_sql_plan_baseline
This data source provides details about a specific Managed Database Sql Plan Baseline resource in Oracle Cloud Infrastructure Database Management service.

Gets the SQL plan baseline details for the specified planName.


## Example Usage

```hcl
data "oci_database_management_managed_database_sql_plan_baseline" "test_managed_database_sql_plan_baseline" {
	#Required
	managed_database_id = oci_database_management_managed_database.test_managed_database.id
	plan_name = var.managed_database_sql_plan_baseline_plan_name

	#Optional
	opc_named_credential_id = var.managed_database_sql_plan_baseline_opc_named_credential_id
}
```

## Argument Reference

The following arguments are supported:

* `managed_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
* `opc_named_credential_id` - (Optional) The OCID of the Named Credential.
* `plan_name` - (Required) The plan name of the SQL plan baseline.


## Attributes Reference

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

