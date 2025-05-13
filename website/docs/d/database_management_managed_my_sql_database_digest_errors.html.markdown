---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_my_sql_database_digest_errors"
sidebar_current: "docs-oci-datasource-database_management-managed_my_sql_database_digest_errors"
description: |-
  Provides the list of Managed My Sql Database Digest Errors in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_my_sql_database_digest_errors
This data source provides the list of Managed My Sql Database Digest Errors in Oracle Cloud Infrastructure Database Management service.

Retrieves any potential errors for a given digest.


## Example Usage

```hcl
data "oci_database_management_managed_my_sql_database_digest_errors" "test_managed_my_sql_database_digest_errors" {
	#Required
	digest = var.managed_my_sql_database_digest_error_digest
	managed_my_sql_database_id = oci_database_management_managed_my_sql_database.test_managed_my_sql_database.id
}
```

## Argument Reference

The following arguments are supported:

* `digest` - (Required) The digest of a MySQL normalized query.
* `managed_my_sql_database_id` - (Required) The OCID of the Managed MySQL Database.


## Attributes Reference

The following attributes are exported:

* `my_sql_digest_errors_collection` - The list of my_sql_digest_errors_collection.

### ManagedMySqlDatabaseDigestError Reference

The following attributes are exported:

* `items` - The unique set of errors for a given digest.
	* `error` - The MySQL error, warning or note raised when a query is run, if any.
		* `code` - The MySQL code of the raised error, warning or note.
		* `level` - The level of severity of the MySQL message.
		* `message_text` - The MySQL message text of the raised error, warning or note.
	* `occurrence_count` - The number of times a MySQL error is encountered.

