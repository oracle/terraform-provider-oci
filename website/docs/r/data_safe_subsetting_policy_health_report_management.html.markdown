---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_subsetting_policy_health_report_management"
sidebar_current: "docs-oci-resource-data_safe-subsetting_policy_health_report_management"
description: |-
  Provides the Subsetting Policy Health Report Management resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_subsetting_policy_health_report_management
This resource manages the health report for a subsetting policy and target database.

During creation, if a matching report does not already exist, the resource generates the report and waits for the asynchronous operation to complete. The report details and logs can be read with the related Data Safe data sources.

## Example Usage

```hcl
resource "oci_data_safe_subsetting_policy_health_report_management" "test_subsetting_policy_health_report_management" {
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

* `check_type` - (Optional) The health check to run. Allowed values are `ALL` and `TABLESPACE_CHECK`.
* `compartment_id` - (Optional) The OCID of the compartment containing the health report. If omitted, it is resolved from the subsetting policy.
* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
* `freeform_tags` - (Optional) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
* `subsetting_policy_id` - (Required) The OCID of the subsetting policy.
* `tablespace` - (Optional) The tablespace to check.
* `target_credentials` - (Required) Credentials for connecting to the target database.
	* `password` - (Required) The password for the target database user.
	* `user_name` - (Required) The user name for the target database.
* `target_id` - (Optional) The OCID of the target database. If omitted, the target is resolved from the subsetting policy.

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the health report.
* `defined_tags` - Defined tags for this resource.
* `display_name` - The display name of the health report.
* `error_count` - The count of errors in the subsetting health report.
* `freeform_tags` - Free-form tags for this resource.
* `id` - The OCID of the health report.
* `state` - The current state of the health report.
* `subsetting_policy_id` - The OCID of the subsetting policy.
* `target_id` - The OCID of the target database for which this report was created.
* `time_created` - The date and time the report was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_updated` - The date and time the report was last updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `warning_count` - The count of warnings in the subsetting health report.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for the asynchronous operation:
	* `create` - (Defaults to 20 minutes), when creating the Subsetting Policy Health Report Management
	* `delete` - (Defaults to 20 minutes), when destroying the Subsetting Policy Health Report Management

## Import

Import is not supported for this resource.
