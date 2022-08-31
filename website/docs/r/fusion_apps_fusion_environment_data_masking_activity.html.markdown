---
subcategory: "Fusion Apps"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fusion_apps_fusion_environment_data_masking_activity"
sidebar_current: "docs-oci-resource-fusion_apps-fusion_environment_data_masking_activity"
description: |-
  Provides the Fusion Environment Data Masking Activity resource in Oracle Cloud Infrastructure Fusion Apps service
---

# oci_fusion_apps_fusion_environment_data_masking_activity
This resource provides the Fusion Environment Data Masking Activity resource in Oracle Cloud Infrastructure Fusion Apps service.

Creates a new DataMaskingActivity.


## Example Usage

```hcl
resource "oci_fusion_apps_fusion_environment_data_masking_activity" "test_fusion_environment_data_masking_activity" {
	#Required
	fusion_environment_id = oci_fusion_apps_fusion_environment.test_fusion_environment.id

	#Optional
	is_resume_data_masking = var.fusion_environment_data_masking_activity_is_resume_data_masking
}
```

## Argument Reference

The following arguments are supported:

* `fusion_environment_id` - (Required) unique FusionEnvironment identifier
* `is_resume_data_masking` - (Optional) This allows the Data Safe service to resume the previously failed data masking activity.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `fusion_environment_id` - Fusion Environment Identifier.
* `id` - Unique identifier that is immutable on creation.
* `state` - The current state of the DataMaskingActivity.
* `time_masking_finish` - The time the data masking activity ended. An RFC3339 formatted datetime string.
* `time_masking_start` - The time the data masking activity started. An RFC3339 formatted datetime string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Fusion Environment Data Masking Activity
	* `update` - (Defaults to 20 minutes), when updating the Fusion Environment Data Masking Activity
	* `delete` - (Defaults to 20 minutes), when destroying the Fusion Environment Data Masking Activity


## Import

FusionEnvironmentDataMaskingActivities can be imported using the `id`, e.g.

```
$ terraform import oci_fusion_apps_fusion_environment_data_masking_activity.test_fusion_environment_data_masking_activity "fusionEnvironments/{fusionEnvironmentId}/dataMaskingActivities/{dataMaskingActivityId}" 
```

