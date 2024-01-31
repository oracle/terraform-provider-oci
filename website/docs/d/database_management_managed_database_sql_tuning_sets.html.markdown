---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_database_sql_tuning_sets"
sidebar_current: "docs-oci-datasource-database_management-managed_database_sql_tuning_sets"
description: |-
  Provides the list of Managed Database Sql Tuning Sets in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_database_sql_tuning_sets
This data source provides the list of Managed Database Sql Tuning Sets in Oracle Cloud Infrastructure Database Management service.

Lists the SQL tuning sets for the specified Managed Database.


## Example Usage

```hcl
data "oci_database_management_managed_database_sql_tuning_sets" "test_managed_database_sql_tuning_sets" {
	#Required
	managed_database_id = oci_database_management_managed_database.test_managed_database.id

	#Optional
	name_contains = var.managed_database_sql_tuning_set_name_contains
	opc_named_credential_id = var.managed_database_sql_tuning_set_opc_named_credential_id
	owner = var.managed_database_sql_tuning_set_owner
}
```

## Argument Reference

The following arguments are supported:

* `managed_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
* `name_contains` - (Optional) Allow searching the name of the SQL tuning set by partial matching. The search is case insensitive.
* `opc_named_credential_id` - (Optional) The OCID of the Named Credential.
* `owner` - (Optional) The owner of the SQL tuning set.


## Attributes Reference

The following attributes are exported:

* `sql_tuning_set_collection` - The list of sql_tuning_set_collection.

### ManagedDatabaseSqlTuningSet Reference

The following attributes are exported:

* `items` - The details in the SQL tuning set summary.
	* `description` - The description of the SQL tuning set.
	* `error_message` - Latest execution error of the plsql that was submitted as a scheduler job.
	* `id` - The unique Sql tuning set identifier. This is not OCID.
	* `name` - The name of the SQL tuning set.
	* `owner` - The owner of the SQL tuning set.
	* `scheduled_job_name` - Name of the Sql tuning set scheduler job.
	* `statement_counts` - The number of SQL statements in the SQL tuning set.
	* `status` - Current status of the Sql tuning set.
	* `time_created` - The created time of the Sql tuning set.
	* `time_last_modified` - Last modified time of the Sql tuning set.
* `managed_database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.

