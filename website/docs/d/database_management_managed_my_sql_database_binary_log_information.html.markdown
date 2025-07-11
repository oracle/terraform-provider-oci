---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_my_sql_database_binary_log_information"
sidebar_current: "docs-oci-datasource-database_management-managed_my_sql_database_binary_log_information"
description: |-
  Provides details about a specific Managed My Sql Database Binary Log Information in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_my_sql_database_binary_log_information
This data source provides details about a specific Managed My Sql Database Binary Log Information resource in Oracle Cloud Infrastructure Database Management service.

Retrieves information pertaining to the binary log of a specific MySQL server.


## Example Usage

```hcl
data "oci_database_management_managed_my_sql_database_binary_log_information" "test_managed_my_sql_database_binary_log_information" {
	#Required
	managed_my_sql_database_id = oci_database_management_managed_my_sql_database.test_managed_my_sql_database.id
}
```

## Argument Reference

The following arguments are supported:

* `managed_my_sql_database_id` - (Required) The OCID of the Managed MySQL Database.


## Attributes Reference

The following attributes are exported:

* `binary_log_compression` - Indicates whether compression is enabled for transactions written to binary log files on the MySQL server.
* `binary_log_compression_percent` - The compression ratio for the binary log, expressed as a percentage.
* `binary_log_format` - The binary logging format used by the MySQL server.
* `binary_log_name` - The name of the binary log file.
* `binary_log_position` - The position within the binary log file.
* `binary_logging` - The status of binary logging on the MySQL server.

