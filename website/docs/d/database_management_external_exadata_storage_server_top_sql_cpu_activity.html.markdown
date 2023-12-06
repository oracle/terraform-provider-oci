---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_external_exadata_storage_server_top_sql_cpu_activity"
sidebar_current: "docs-oci-datasource-database_management-external_exadata_storage_server_top_sql_cpu_activity"
description: |-
  Provides details about a specific External Exadata Storage Server Top Sql Cpu Activity in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_external_exadata_storage_server_top_sql_cpu_activity
This data source provides details about a specific External Exadata Storage Server Top Sql Cpu Activity resource in Oracle Cloud Infrastructure Database Management service.

Gets the SQL IDs with the top CPU activity from the Exadata storage server.


## Example Usage

```hcl
data "oci_database_management_external_exadata_storage_server_top_sql_cpu_activity" "test_external_exadata_storage_server_top_sql_cpu_activity" {
	#Required
	external_exadata_storage_server_id = oci_database_management_external_exadata_storage_server.test_external_exadata_storage_server.id
}
```

## Argument Reference

The following arguments are supported:

* `external_exadata_storage_server_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata storage server.


## Attributes Reference

The following attributes are exported:

* `activity` - A list of sql CPU activity.
	* `cpu_activity` - The CPU activity percentage.
	* `database_name` - The database name.
	* `sql_id` - The SQL ID.

