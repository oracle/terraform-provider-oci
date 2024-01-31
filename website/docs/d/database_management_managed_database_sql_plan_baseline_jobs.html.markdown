---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_database_sql_plan_baseline_jobs"
sidebar_current: "docs-oci-datasource-database_management-managed_database_sql_plan_baseline_jobs"
description: |-
  Provides the list of Managed Database Sql Plan Baseline Jobs in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_database_sql_plan_baseline_jobs
This data source provides the list of Managed Database Sql Plan Baseline Jobs in Oracle Cloud Infrastructure Database Management service.

Lists the database jobs used for loading SQL plan baselines in the specified Managed Database.


## Example Usage

```hcl
data "oci_database_management_managed_database_sql_plan_baseline_jobs" "test_managed_database_sql_plan_baseline_jobs" {
	#Required
	managed_database_id = oci_database_management_managed_database.test_managed_database.id

	#Optional
	name = var.managed_database_sql_plan_baseline_job_name
	opc_named_credential_id = var.managed_database_sql_plan_baseline_job_opc_named_credential_id
}
```

## Argument Reference

The following arguments are supported:

* `managed_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
* `name` - (Optional) A filter to return the SQL plan baseline jobs that match the name.
* `opc_named_credential_id` - (Optional) The OCID of the Named Credential.


## Attributes Reference

The following attributes are exported:

* `sql_plan_baseline_job_collection` - The list of sql_plan_baseline_job_collection.

### ManagedDatabaseSqlPlanBaselineJob Reference

The following attributes are exported:

* `items` - A list of SQL plan baseline jobs.
	* `name` - The name of the job.
	* `status` - The status of the job.
	* `time_created` - The date and time the job was created.
	* `type` - The type of the job.

