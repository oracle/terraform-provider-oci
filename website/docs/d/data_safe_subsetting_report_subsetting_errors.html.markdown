---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_subsetting_report_subsetting_errors"
sidebar_current: "docs-oci-datasource-data_safe-subsetting_report_subsetting_errors"
description: |-
  Provides the list of Subsetting Report Subsetting Errors in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_subsetting_report_subsetting_errors
This data source provides the list of Subsetting Report Subsetting Errors in Oracle Cloud Infrastructure Data Safe service.

Gets a list of subsetting errors in a subsetting run based on the specified query parameters.


## Example Usage

```hcl
data "oci_data_safe_subsetting_report_subsetting_errors" "test_subsetting_report_subsetting_errors" {
	#Required
	subsetting_report_id = oci_data_safe_subsetting_report.test_subsetting_report.id

	#Optional
	step_name = var.subsetting_report_subsetting_error_step_name
}
```

## Argument Reference

The following arguments are supported:

* `step_name` - (Optional) A filter to return only subsetting errors that match the specified step name.
* `subsetting_report_id` - (Required) The OCID of the subsetting report.


## Attributes Reference

The following attributes are exported:

* `subsetting_error_collection` - The list of subsetting_error_collection.

### SubsettingReportSubsettingError Reference

The following attributes are exported:

* `items` - An array of subsetting error objects.
	* `error` - The text of the subsetting error.
	* `failed_statement` - The statement resulting into the error.
	* `step_name` - The stepName of the subsetting error.
	* `time_created` - The date and time the error entry was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 

