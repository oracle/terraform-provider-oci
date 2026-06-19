---
subcategory: "Psql"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_psql_db_system_pitr_detail"
sidebar_current: "docs-oci-datasource-psql-db_system_pitr_detail"
description: |-
  Provides details about a specific Db System Pitr Detail in Oracle Cloud Infrastructure Psql service
---

# Data Source: oci_psql_db_system_pitr_detail
This data source provides details about a specific Db System Pitr Detail resource in Oracle Cloud Infrastructure Psql service.

Gets the database system PITR details.

## Example Usage

```hcl
data "oci_psql_db_system_pitr_detail" "test_db_system_pitr_detail" {
	#Required
	db_system_id = oci_psql_db_system.test_db_system.id
}
```

## Argument Reference

The following arguments are supported:

* `db_system_id` - (Required) A unique identifier for the database system.


## Attributes Reference

The following attributes are exported:

* `pitr_state` - The current state of the point-in-time recovery of the db system.
* `recovery_time_windows` - List of point-in-time recovery windows.
	* `time_recovery_window_end` - Latest timestamp in the PITR window to which the database can be restored. Timestamps later than this are not recoverable. The value must be an [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp.  Example: `2016-08-25T21:10:29Z`
	* `time_recovery_window_start` - Earliest timestamp in the PITR window to which the database can be restored. Timestamps earlier than this are not recoverable. The value must be an [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp.  Example: `2016-08-25T21:10:29Z`
