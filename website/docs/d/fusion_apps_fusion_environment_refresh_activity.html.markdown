---
subcategory: "Fusion Apps"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fusion_apps_fusion_environment_refresh_activity"
sidebar_current: "docs-oci-datasource-fusion_apps-fusion_environment_refresh_activity"
description: |-
  Provides details about a specific Fusion Environment Refresh Activity in Oracle Cloud Infrastructure Fusion Apps service
---

# Data Source: oci_fusion_apps_fusion_environment_refresh_activity
This data source provides details about a specific Fusion Environment Refresh Activity resource in Oracle Cloud Infrastructure Fusion Apps service.

Gets a RefreshActivity by identifier

## Example Usage

```hcl
data "oci_fusion_apps_fusion_environment_refresh_activity" "test_fusion_environment_refresh_activity" {
	#Required
	fusion_environment_id = oci_fusion_apps_fusion_environment.test_fusion_environment.id
	refresh_activity_id = oci_fusion_apps_refresh_activity.test_refresh_activity.id
}
```

## Argument Reference

The following arguments are supported:

* `fusion_environment_id` - (Required) unique FusionEnvironment identifier
* `refresh_activity_id` - (Required) The unique identifier (OCID) of the Refresh activity.


## Attributes Reference

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

