---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_subsetting_reports"
sidebar_current: "docs-oci-datasource-data_safe-subsetting_reports"
description: |-
  Provides the list of Subsetting Reports in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_subsetting_reports
This data source provides the list of Subsetting Reports in Oracle Cloud Infrastructure Data Safe service.

Gets a list of subsetting reports based on the specified query parameters.

## Example Usage

```hcl
data "oci_data_safe_subsetting_reports" "test_subsetting_reports" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.subsetting_report_access_level
	compartment_id_in_subtree = var.subsetting_report_compartment_id_in_subtree
	subsetting_policy_id = oci_data_safe_subsetting_policy.test_subsetting_policy.id
	target_database_group_id = oci_data_safe_target_database_group.test_target_database_group.id
	target_id = oci_cloud_guard_target.test_target.id
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `subsetting_policy_id` - (Optional) A filter to return only the resources that match the specified subsetting policy OCID.
* `target_database_group_id` - (Optional) A filter to return the target database group that matches the specified OCID.
* `target_id` - (Optional) A filter to return only items related to a specific target OCID.


## Attributes Reference

The following attributes are exported:

* `subsetting_report_collection` - The list of subsetting_report_collection.

### SubsettingReport Reference

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

