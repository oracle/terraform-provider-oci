---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_masking_report_masking_errors"
sidebar_current: "docs-oci-datasource-data_safe-masking_report_masking_errors"
description: |-
  Provides the list of Masking Report Masking Errors in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_masking_report_masking_errors
This data source provides the list of Masking Report Masking Errors in Oracle Cloud Infrastructure Data Safe service.

Gets a list of masking errors in a masking run based on the specified query parameters.


## Example Usage

```hcl
data "oci_data_safe_masking_report_masking_errors" "test_masking_report_masking_errors" {
	#Required
	masking_report_id = oci_data_safe_masking_report.test_masking_report.id

	#Optional
	step_name = var.masking_report_masking_error_step_name
}
```

## Argument Reference

The following arguments are supported:

* `masking_report_id` - (Required) The OCID of the masking report.
* `step_name` - (Optional) A filter to return only masking errors that match the specified step name.


## Attributes Reference

The following attributes are exported:

* `masking_error_collection` - The list of masking_error_collection.

### MaskingReportMaskingError Reference

The following attributes are exported:

* `items` - An array of masking error objects.
	* `error` - The text of the masking error.
	* `failed_statement` - The statement resulting into the error.
	* `step_name` - The stepName of the masking error.
	* `time_created` - The date and time the error entry was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 

