---
subcategory: "Fusion Apps"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fusion_apps_fusion_environment_refresh_activity"
sidebar_current: "docs-oci-resource-fusion_apps-fusion_environment_refresh_activity"
description: |-
  Provides the Fusion Environment Refresh Activity resource in Oracle Cloud Infrastructure Fusion Apps service
---

# oci_fusion_apps_fusion_environment_refresh_activity
This resource provides the Fusion Environment Refresh Activity resource in Oracle Cloud Infrastructure Fusion Apps service.

Creates a new RefreshActivity.


## Example Usage

```hcl
resource "oci_fusion_apps_fusion_environment_refresh_activity" "test_fusion_environment_refresh_activity" {
	#Required
	fusion_environment_id = oci_fusion_apps_fusion_environment.test_fusion_environment.id
	source_fusion_environment_id = oci_fusion_apps_fusion_environment.test_fusion_environment.id
	#Optional
	is_data_masking_opted = var.fusion_environment_refresh_activity_is_data_masking_opted
	time_scheduled_start = var.fusion_environment_refresh_activity_time_scheduled_start
}
```

## Argument Reference

The following arguments are supported:

* `fusion_environment_id` - (Required) unique FusionEnvironment identifier
* `is_data_masking_opted` - (Optional) Represents if the customer opted for Data Masking or not during refreshActivity.
* `source_fusion_environment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the source environment 
* `time_scheduled_start` - (Optional) (Updatable) Current time the refresh activity is scheduled to start. An RFC3339 formatted datetime string.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Fusion Environment Refresh Activity
	* `update` - (Defaults to 20 minutes), when updating the Fusion Environment Refresh Activity
	* `delete` - (Defaults to 20 minutes), when destroying the Fusion Environment Refresh Activity


## Import

FusionEnvironmentRefreshActivities can be imported using the `id`, e.g.

```
$ terraform import oci_fusion_apps_fusion_environment_refresh_activity.test_fusion_environment_refresh_activity "fusionEnvironments/{fusionEnvironmentId}/refreshActivities/{refreshActivityId}" 
```

