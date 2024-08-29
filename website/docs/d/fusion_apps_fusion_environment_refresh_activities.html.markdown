---
subcategory: "Fusion Apps"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fusion_apps_fusion_environment_refresh_activities"
sidebar_current: "docs-oci-datasource-fusion_apps-fusion_environment_refresh_activities"
description: |-
  Provides the list of Fusion Environment Refresh Activities in Oracle Cloud Infrastructure Fusion Apps service
---

# Data Source: oci_fusion_apps_fusion_environment_refresh_activities
This data source provides the list of Fusion Environment Refresh Activities in Oracle Cloud Infrastructure Fusion Apps service.

Returns a list of RefreshActivities.


## Example Usage

```hcl
data "oci_fusion_apps_fusion_environment_refresh_activities" "test_fusion_environment_refresh_activities" {
	#Required
	fusion_environment_id = oci_fusion_apps_fusion_environment.test_fusion_environment.id

	#Optional
	display_name = var.fusion_environment_refresh_activity_display_name
	state = var.fusion_environment_refresh_activity_state
	time_expected_finish_less_than_or_equal_to = var.fusion_environment_refresh_activity_time_expected_finish_less_than_or_equal_to
	time_scheduled_start_greater_than_or_equal_to = var.fusion_environment_refresh_activity_time_scheduled_start_greater_than_or_equal_to
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `fusion_environment_id` - (Required) unique FusionEnvironment identifier
* `state` - (Optional) A filter that returns all resources that match the specified status
* `time_expected_finish_less_than_or_equal_to` - (Optional) A filter that returns all resources that end before this date
* `time_scheduled_start_greater_than_or_equal_to` - (Optional) A filter that returns all resources that are scheduled after this date


## Attributes Reference

The following attributes are exported:

* `refresh_activity_collection` - The list of refresh_activity_collection.

### FusionEnvironmentRefreshActivity Reference

The following attributes are exported:

* `display_name` - A friendly name for the refresh activity. Can be changed later.
* `id` - The unique identifier (OCID) of the refresh activity. Can't be changed after creation.
* `is_data_masking_opted` - Represents if the customer opted for Data Masking or not during refreshActivity.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `refresh_issue_details_list` - Details of refresh investigation information, each item represents a different issue.
	* `refresh_issues` - Detail reasons of refresh failure or validation failure that needs to be shown to customer.
* `service_availability` - Service availability / impact during refresh activity execution up down
* `source_fusion_environment_id` - The OCID of the Fusion environment that is the source environment for the refresh.
* `state` - The current state of the refreshActivity.
* `time_accepted` - The time the refresh activity record was created. An RFC3339 formatted datetime string.
* `time_expected_finish` - The time the refresh activity is scheduled to end. An RFC3339 formatted datetime string.
* `time_finished` - The time the refresh activity actually completed / cancelled / failed. An RFC3339 formatted datetime string.
* `time_of_restoration_point` - The date and time of the most recent source environment backup used for the environment refresh.
* `time_scheduled_start` - The time the refresh activity is scheduled to start. An RFC3339 formatted datetime string.
* `time_updated` - The time the refresh activity record was updated. An RFC3339 formatted datetime string.

