---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_database_alert_log_count"
sidebar_current: "docs-oci-datasource-database_management-managed_database_alert_log_count"
description: |-
  Provides details about a specific Managed Database Alert Log Count in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_database_alert_log_count
This data source provides details about a specific Managed Database Alert Log Count resource in Oracle Cloud Infrastructure Database Management service.

Get the counts of alert logs for the specified Managed Database.


## Example Usage

```hcl
data "oci_database_management_managed_database_alert_log_count" "test_managed_database_alert_log_count" {
	#Required
	managed_database_id = oci_database_management_managed_database.test_managed_database.id

	#Optional
	group_by = var.managed_database_alert_log_count_group_by
	is_regular_expression = var.managed_database_alert_log_count_is_regular_expression
	level_filter = var.managed_database_alert_log_count_level_filter
	log_search_text = var.managed_database_alert_log_count_log_search_text
	time_greater_than_or_equal_to = var.managed_database_alert_log_count_time_greater_than_or_equal_to
	time_less_than_or_equal_to = var.managed_database_alert_log_count_time_less_than_or_equal_to
	type_filter = var.managed_database_alert_log_count_type_filter
}
```

## Argument Reference

The following arguments are supported:

* `group_by` - (Optional) The optional parameter used to group different alert logs.
* `is_regular_expression` - (Optional) The flag to indicate whether the search text is regular expression or not.
* `level_filter` - (Optional) The optional parameter to filter the alert logs by log level.
* `log_search_text` - (Optional) The optional query parameter to filter the attention or alert logs by search text.
* `managed_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
* `time_greater_than_or_equal_to` - (Optional) The optional greater than or equal to timestamp to filter the logs.
* `time_less_than_or_equal_to` - (Optional) The optional less than or equal to timestamp to filter the logs.
* `type_filter` - (Optional) The optional parameter to filter the attention or alert logs by type.


## Attributes Reference

The following attributes are exported:

* `items` - An array of the counts of different urgency or type of alert logs.
	* `category` - The category of different alert logs.
	* `count` - The count of alert logs with specific category.
* `managed_database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.

