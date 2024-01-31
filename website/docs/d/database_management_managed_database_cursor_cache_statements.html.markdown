---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_database_cursor_cache_statements"
sidebar_current: "docs-oci-datasource-database_management-managed_database_cursor_cache_statements"
description: |-
  Provides the list of Managed Database Cursor Cache Statements in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_database_cursor_cache_statements
This data source provides the list of Managed Database Cursor Cache Statements in Oracle Cloud Infrastructure Database Management service.

Lists the SQL statements from shared SQL area, also called the cursor cache.


## Example Usage

```hcl
data "oci_database_management_managed_database_cursor_cache_statements" "test_managed_database_cursor_cache_statements" {
	#Required
	managed_database_id = oci_database_management_managed_database.test_managed_database.id

	#Optional
	opc_named_credential_id = var.managed_database_cursor_cache_statement_opc_named_credential_id
	sql_text = var.managed_database_cursor_cache_statement_sql_text
}
```

## Argument Reference

The following arguments are supported:

* `managed_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
* `opc_named_credential_id` - (Optional) The OCID of the Named Credential.
* `sql_text` - (Optional) A filter to return all the SQL plan baselines that match the SQL text. By default, the search is case insensitive. To run an exact or case-sensitive search, double-quote the search string. You may also use the '%' symbol as a wildcard. 


## Attributes Reference

The following attributes are exported:

* `cursor_cache_statement_collection` - The list of cursor_cache_statement_collection.

### ManagedDatabaseCursorCacheStatement Reference

The following attributes are exported:

* `items` - A list of SQL statements in the cursor cache.
	* `schema` - The name of the parsing schema.
	* `sql_id` - The SQL statement identifier. Identifies a SQL statement in the cursor cache.
	* `sql_text` - The first thousand characters of the SQL text.

