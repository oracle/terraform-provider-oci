---
subcategory: "Fusion Apps"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fusion_apps_fusion_environment_data_masking_activity"
sidebar_current: "docs-oci-datasource-fusion_apps-fusion_environment_data_masking_activity"
description: |-
  Provides details about a specific Fusion Environment Data Masking Activity in Oracle Cloud Infrastructure Fusion Apps service
---

# Data Source: oci_fusion_apps_fusion_environment_data_masking_activity
This data source provides details about a specific Fusion Environment Data Masking Activity resource in Oracle Cloud Infrastructure Fusion Apps service.

Gets a DataMaskingActivity by identifier

## Example Usage

```hcl
data "oci_fusion_apps_fusion_environment_data_masking_activity" "test_fusion_environment_data_masking_activity" {
	#Required
	data_masking_activity_id = oci_fusion_apps_data_masking_activity.test_data_masking_activity.id
	fusion_environment_id = oci_fusion_apps_fusion_environment.test_fusion_environment.id
}
```

## Argument Reference

The following arguments are supported:

* `data_masking_activity_id` - (Required) Unique DataMasking run identifier.
* `fusion_environment_id` - (Required) unique FusionEnvironment identifier


## Attributes Reference

The following attributes are exported:

* `fusion_environment_id` - Fusion Environment Identifier.
* `id` - Unique identifier that is immutable on creation.
* `state` - The current state of the DataMaskingActivity.
* `time_masking_finish` - The time the data masking activity ended. An RFC3339 formatted datetime string.
* `time_masking_start` - The time the data masking activity started. An RFC3339 formatted datetime string.

