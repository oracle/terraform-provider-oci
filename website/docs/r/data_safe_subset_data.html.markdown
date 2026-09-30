---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_subset_data"
sidebar_current: "docs-oci-resource-data_safe-subset_data"
description: |-
  Provides the Subset Data resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_subset_data
This resource starts an asynchronous data subsetting operation in Oracle Cloud Infrastructure Data Safe.

The resource is an action resource: it starts the operation during creation and is intentionally non-deletable. Configure the subsetting policy and rules before creating it.

## Example Usage

```hcl
resource "oci_data_safe_subset_data" "test_subset_data" {
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

The resource does not export service attributes. Its `id` is set to the subsetting policy OCID after the operation completes.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for the asynchronous operation:
	* `create` - (Defaults to 20 minutes), when starting the subsetting operation

## Import

Import is not supported for this resource.
