---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_masking_report"
sidebar_current: "docs-oci-datasource-data_safe-masking_report"
description: |-
  Provides details about a specific Masking Report in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_masking_report
This data source provides details about a specific Masking Report resource in Oracle Cloud Infrastructure Data Safe service.

Gets the details of the specified masking report.

## Example Usage

```hcl
data "oci_data_safe_masking_report" "test_masking_report" {
	#Required
	masking_report_id = oci_data_safe_masking_report.test_masking_report.id
}
```

## Argument Reference

The following arguments are supported:

* `masking_report_id` - (Required) The OCID of the masking report.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the masking report.
* `id` - The OCID of the masking report.
* `masking_policy_id` - The OCID of the masking policy used.
* `masking_work_request_id` - The OCID of the masking work request that resulted in this masking report.
* `target_id` - The OCID of the target database masked.
* `time_masking_finished` - The date and time data masking finished, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339)
* `time_masking_started` - The date and time data masking started, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339)
* `total_masked_columns` - The total number of masked columns.
* `total_masked_objects` - The total number of unique objects (tables and editioning views) that contain the masked columns.
* `total_masked_schemas` - The total number of unique schemas that contain the masked columns.
* `total_masked_sensitive_types` - The total number of unique sensitive types associated with the masked columns.
* `total_masked_values` - The total number of masked values.

