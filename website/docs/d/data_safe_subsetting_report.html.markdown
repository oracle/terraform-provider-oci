---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_subsetting_report"
sidebar_current: "docs-oci-datasource-data_safe-subsetting_report"
description: |-
  Provides details about a specific Subsetting Report in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_subsetting_report
This data source provides details about a specific Subsetting Report resource in Oracle Cloud Infrastructure Data Safe service.

Gets the details of the specified subsetting report.

## Example Usage

```hcl
data "oci_data_safe_subsetting_report" "test_subsetting_report" {
	#Required
	subsetting_report_id = oci_data_safe_subsetting_report.test_subsetting_report.id
}
```

## Argument Reference

The following arguments are supported:

* `subsetting_report_id` - (Required) The OCID of the subsetting report.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the subsetting report
* `database_size_after_subsetting_in_kbs` - The size of the target database after subsetting in KBs
* `database_size_before_subsetting_in_kbs` - The size of the target database before subsetting in KBs
* `id` - The OCID of the subsetting report
* `is_redo_logging_enabled` - Indicates if redo logging was enabled during the subsetting operation 
* `is_refresh_stats_enabled` - Indicates if statistics gathering was enabled during the subsetting operation 
* `masking_policy_id` - The OCID of the masking policy associated with this subsetting report
* `masking_report_id` - The OCID of the masking report associated with this subsetting report
* `masking_work_request_id` - The OCID of the masking work request triggered after this subsetting job
* `parallel_degree` - Indicates if parallel execution was enabled during the subsetting operation 
* `recompile` - Indicates how invalid objects were recompiled post the subsetting operation 
* `state` - The current state of the subsetting report
* `subsetting_policy_id` - The OCID of the subsetting policy used
* `subsetting_status` - The status of the subsetting job
* `subsetting_work_request_id` - The OCID of the subsetting work request that resulted in this subsetting report
* `target_id` - The OCID of the target database subsetted
* `time_created` - The date and time the subsetting report was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339) 
* `time_subsetting_finished` - The date and time data subsetting finished, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339)
* `time_subsetting_started` - The date and time data subsetting started, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339)
* `total_post_subsetting_script_errors` - The total number of errors in post-subsetting script
* `total_pre_subsetting_script_errors` - The total number of errors in pre-subsetting script
* `total_subsetted_objects` - The total number of subsetted objects
* `total_subsetted_rows` - The count of rows reduced in the subsetting job
* `total_subsetted_schemas` - The total number of subsetted schemas

