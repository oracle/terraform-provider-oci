---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_database_table_statistics"
sidebar_current: "docs-oci-datasource-database_management-managed_database_table_statistics"
description: |-
  Provides the list of Managed Database Table Statistics in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_database_table_statistics
This data source provides the list of Managed Database Table Statistics in Oracle Cloud Infrastructure Database Management service.

Gets the number of database table objects grouped by different statuses such as
Not Stale Stats, Stale Stats, and No Stats. This also includes the percentage of each status.


## Example Usage

```hcl
data "oci_database_management_managed_database_table_statistics" "test_managed_database_table_statistics" {
	#Required
	managed_database_id = oci_database_management_managed_database.test_managed_database.id
}
```

## Argument Reference

The following arguments are supported:

* `managed_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.


## Attributes Reference

The following attributes are exported:

* `table_statistics_collection` - The list of table_statistics_collection.

### ManagedDatabaseTableStatistic Reference

The following attributes are exported:

* `items` - The list of table statistics statuses.
	* `count` - The number of objects aggregated by status category.
	* `percentage` - The percentage of objects with a particular status.
	* `type` - The valid status categories of table statistics.

