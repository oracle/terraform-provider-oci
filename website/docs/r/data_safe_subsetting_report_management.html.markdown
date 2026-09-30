---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_subsetting_report_management"
sidebar_current: "docs-oci-resource-data_safe-subsetting_report_management"
description: |-
  Provides the Subsetting Report Management resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_subsetting_report_management
This resource manages the subsetting report for a subsetting policy and target database.

During creation, if a matching report does not already exist, the resource starts the subsetting operation and waits for the report. The related report details can be read with `oci_data_safe_subsetting_report`.

## Example Usage

```hcl
resource "oci_data_safe_subsetting_report_management" "test_subsetting_report_management" {
  subsetting_policy_id = oci_data_safe_subsetting_policy.test_subsetting_policy.id
  target_id            = var.target_id

  target_credentials {
    user_name = var.target_database_user_name
    password  = var.target_database_password
  }

  depends_on = [oci_data_safe_subsetting_policy_subsetting_rule.test_subsetting_policy_subsetting_rule]
}
```

## Argument Reference

The following arguments are supported:

* `is_redo_logging_enabled` - (Optional) Whether redo logging is enabled during the subsetting operation.
* `is_refresh_stats_enabled` - (Optional) Whether statistics gathering is enabled after the subsetting operation.
* `is_rerun` - (Optional) Whether to rerun a previously started subsetting operation.
* `masking` - (Optional) Whether masking is enabled during subsetting. Allowed values are `ENABLED` and `DISABLED`.
* `parallel_degree` - (Optional) The degree of parallel execution used during subsetting. Use `NONE`, `DEFAULT`, or an integer value.
* `re_run_from_step` - (Optional) The step from which to rerun the operation. Allowed values are `PRE_SUBSETTING_SCRIPT` and `POST_SUBSETTING_SCRIPT`.
* `recompile` - (Optional) How invalid objects are recompiled after subsetting. Allowed values are `SERIAL`, `PARALLEL`, and `NONE`.
* `subsetting_policy_id` - (Required) The OCID of the subsetting policy.
* `tablespace` - (Optional) The tablespace used for the subsetting operation.
* `target_credentials` - (Required) Credentials for connecting to the target database.
	* `password` - (Required) The password for the target database user.
	* `user_name` - (Required) The user name for the target database.
* `target_id` - (Optional) The OCID of the target database. If omitted, the target is resolved from the subsetting policy.

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the subsetting report.
* `database_size_after_subsetting_in_kbs` - The size of the target database after subsetting in KBs.
* `database_size_before_subsetting_in_kbs` - The size of the target database before subsetting in KBs.
* `id` - The OCID of the subsetting report.
* `is_redo_logging_enabled` - Whether redo logging was enabled during the subsetting operation.
* `is_refresh_stats_enabled` - Whether statistics gathering was enabled during the subsetting operation.
* `masking_policy_id` - The OCID of the masking policy associated with the subsetting report.
* `masking_report_id` - The OCID of the masking report associated with the subsetting report.
* `masking_work_request_id` - The OCID of the masking work request triggered after the subsetting job.
* `parallel_degree` - The degree of parallel execution used during subsetting.
* `recompile` - How invalid objects were recompiled after subsetting.
* `state` - The current state of the subsetting report.
* `subsetting_policy_id` - The OCID of the subsetting policy used.
* `subsetting_status` - The status of the subsetting job.
* `subsetting_work_request_id` - The OCID of the work request that resulted in this subsetting report.
* `target_id` - The OCID of the target database subsetted.
* `time_created` - The date and time the subsetting report was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_subsetting_finished` - The date and time data subsetting finished, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_subsetting_started` - The date and time data subsetting started, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `total_post_subsetting_script_errors` - The total number of errors in the post-subsetting script.
* `total_pre_subsetting_script_errors` - The total number of errors in the pre-subsetting script.
* `total_subsetted_objects` - The total number of subsetted objects.
* `total_subsetted_rows` - The count of rows reduced in the subsetting job.
* `total_subsetted_schemas` - The total number of subsetted schemas.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for the asynchronous operation:
	* `create` - (Defaults to 20 minutes), when creating the Subsetting Report Management
	* `delete` - (Defaults to 20 minutes), when destroying the Subsetting Report Management

## Import

Import is not supported for this resource.
