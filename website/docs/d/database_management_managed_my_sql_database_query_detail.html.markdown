---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_my_sql_database_query_detail"
sidebar_current: "docs-oci-datasource-database_management-managed_my_sql_database_query_detail"
description: |-
  Provides details about a specific Managed My Sql Database Query Detail in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_my_sql_database_query_detail
This data source provides details about a specific Managed My Sql Database Query Detail resource in Oracle Cloud Infrastructure Database Management service.

Retrieves query sample details, explain plan and potential warnings for a given digest.


## Example Usage

```hcl
data "oci_database_management_managed_my_sql_database_query_detail" "test_managed_my_sql_database_query_detail" {
	#Required
	digest = var.managed_my_sql_database_query_detail_digest
	managed_my_sql_database_id = oci_database_management_managed_my_sql_database.test_managed_my_sql_database.id
}
```

## Argument Reference

The following arguments are supported:

* `digest` - (Required) The digest of a MySQL normalized query.
* `managed_my_sql_database_id` - (Required) The OCID of the Managed MySQL Database.


## Attributes Reference

The following attributes are exported:

* `query_explain_plan` - The explain plan for a given MySQL query.
	* `json_explain` - The json format of the explain plan.
	* `json_explain_version` - The version of the Json format of MySQL Explain.
* `query_messages` - The errors, warnings and notes that could be raised by the execution of the query.
	* `code` - The MySQL code of the raised error, warning or note.
	* `level` - The level of severity of the MySQL message.
	* `message_text` - The MySQL message text of the raised error, warning or note.
* `query_sample_details` - The details of a query sample including the query text, execution time and other details.
	* `execution_time` - The total amount of time that has been spent executing the query sample.
	* `host` - The host from which the query sample was run.
	* `mysql_instance` - The MySQL instance against which the query sample was run.
	* `query_sample_text` - The query sample mapped by MySQL to a given normalized query.
	* `thread_id` - The thread id of the connection.
	* `time_query_sample_seen` - The date and time the query sample was last seen.
	* `user` - The user who ran the query sample.

