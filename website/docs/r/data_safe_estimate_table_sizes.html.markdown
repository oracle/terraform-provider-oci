---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_estimate_table_sizes"
sidebar_current: "docs-oci-resource-data_safe-estimate_table_sizes"
description: |-
  Provides the Estimate Table Sizes resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_estimate_table_sizes
This resource starts an asynchronous table-size estimation operation for a subsetting policy in Oracle Cloud Infrastructure Data Safe.

The resource is an action resource: it starts the estimation during creation and is intentionally non-deletable. The resulting estimates can be read with `oci_data_safe_subsetting_policy_table_estimates`.

## Example Usage

```hcl
resource "oci_data_safe_estimate_table_sizes" "test_estimate_table_sizes" {
  subsetting_policy_id = oci_data_safe_subsetting_policy.test_subsetting_policy.id
  target_id            = var.target_id

  target_credentials {
    user_name = var.target_database_user_name
    password  = var.target_database_password
  }
}
```

## Argument Reference

The following arguments are supported:

* `subsetting_policy_id` - (Required) The OCID of the subsetting policy.
* `target_credentials` - (Required) Credentials for connecting to the target database.
	* `password` - (Required) The password for the target database user.
	* `user_name` - (Required) The user name for the target database.
* `target_id` - (Optional) The OCID of the target database. If omitted, the target is resolved from the subsetting policy.

## Attributes Reference

The resource does not export service attributes. Its `id` is set to the subsetting policy OCID after the operation completes.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for the asynchronous operation:
	* `create` - (Defaults to 20 minutes), when starting table-size estimation

## Import

Import is not supported for this resource.
