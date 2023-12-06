---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_database_sql_tuning_set"
sidebar_current: "docs-oci-datasource-database_management-managed_database_sql_tuning_set"
description: |-
  Provides details about a specific Managed Database Sql Tuning Set in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_database_sql_tuning_set
This data source provides details about a specific Managed Database Sql Tuning Set resource in Oracle Cloud Infrastructure Database Management service.

Lists the SQL tuning sets for the specified Managed Database.


## Example Usage

```hcl
data "oci_database_management_managed_database_sql_tuning_set" "test_managed_database_sql_tuning_set" {
	#Required
	managed_database_id = oci_database_management_managed_database.test_managed_database.id

	#Optional
	name_contains = var.managed_database_sql_tuning_set_name_contains
	owner = var.managed_database_sql_tuning_set_owner
}
```

## Argument Reference

The following arguments are supported:

* `managed_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
* `name_contains` - (Optional) Allow searching the name of the SQL tuning set by partial matching. The search is case insensitive.
* `owner` - (Optional) The owner of the SQL tuning set.


## Attributes Reference

The following attributes are exported:

* `items` - The details in the SQL tuning set summary.
	* `description` - The description of the SQL tuning set.
	* `name` - The name of the SQL tuning set.
	* `owner` - The owner of the SQL tuning set.
	* `statement_counts` - The number of SQL statements in the SQL tuning set.
* `managed_database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.

