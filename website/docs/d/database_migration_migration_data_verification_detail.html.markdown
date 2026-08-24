---
subcategory: "Database Migration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_migration_migration_data_verification_detail"
sidebar_current: "docs-oci-datasource-database_migration-migration_data_verification_detail"
description: |-
  Provides details about a specific Migration Data Verification Detail in Oracle Cloud Infrastructure Database Migration service
---

# Data Source: oci_database_migration_migration_data_verification_detail
This data source provides details about a specific Migration Data Verification Detail resource in Oracle Cloud Infrastructure Database Migration service.

Returns availability and report status information used by the Data Verification UI.

The migration must already have Data Verification results.


## Example Usage

```hcl
data "oci_database_migration_migration_data_verification_detail" "test_migration_data_verification_detail" {
	#Required
	migration_id = var.migration_id
}
```

## Argument Reference

The following arguments are supported:

* `migration_id` - (Required) The OCID of the migration 


## Attributes Reference

The following attributes are exported:

* `lifecycle_details` - A human-readable string that provides more details about the current lifecycle state.

	This field is intended to follow Oracle Cloud Infrastructure patterns and may include an internal error code and message when `lifecycleState` is `FAILED` (for example, `DMS-xxxx: <message>`). When `lifecycleState` is `SUCCEEDED`, this field may contain a non-failing disclaimer, such as unavailable table row-count comparisons caused by missing optimizer statistics. 
* `migration_id` - The OCID of the resource being referenced.
* `state` - Overall lifecycle state of Data Verification generation for the migration. This is a single, feature-level state (all Data Verification reports are generated as one batch). 
* `time_completed` - An RFC3339 formatted datetime string such as `2016-08-25T21:10:29.600Z`. 
* `time_created` - An RFC3339 formatted datetime string such as `2016-08-25T21:10:29.600Z`. 
* `time_started` - An RFC3339 formatted datetime string such as `2016-08-25T21:10:29.600Z`. 
* `time_updated` - An RFC3339 formatted datetime string such as `2016-08-25T21:10:29.600Z`. 
