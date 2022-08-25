---
subcategory: "Fusion Apps"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fusion_apps_fusion_environment_data_masking_activities"
sidebar_current: "docs-oci-datasource-fusion_apps-fusion_environment_data_masking_activities"
description: |-
  Provides the list of Fusion Environment Data Masking Activities in Oracle Cloud Infrastructure Fusion Apps service
---

# Data Source: oci_fusion_apps_fusion_environment_data_masking_activities
This data source provides the list of Fusion Environment Data Masking Activities in Oracle Cloud Infrastructure Fusion Apps service.

Returns a list of DataMaskingActivities.


## Example Usage

```hcl
data "oci_fusion_apps_fusion_environment_data_masking_activities" "test_fusion_environment_data_masking_activities" {
	#Required
	fusion_environment_id = oci_fusion_apps_fusion_environment.test_fusion_environment.id

	#Optional
	state = var.fusion_environment_data_masking_activity_state
}
```

## Argument Reference

The following arguments are supported:

* `fusion_environment_id` - (Required) unique FusionEnvironment identifier
* `state` - (Optional) A filter that returns all resources that match the specified status


## Attributes Reference

The following attributes are exported:

* `data_masking_activity_collection` - The list of data_masking_activity_collection.

### FusionEnvironmentDataMaskingActivity Reference

The following attributes are exported:

* `fusion_environment_id` - Fusion Environment Identifier.
* `id` - Unique identifier that is immutable on creation.
* `state` - The current state of the DataMaskingActivity.
* `time_masking_finish` - The time the data masking activity ended. An RFC3339 formatted datetime string.
* `time_masking_start` - The time the data masking activity started. An RFC3339 formatted datetime string.

