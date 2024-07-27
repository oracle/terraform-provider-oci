---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_alert"
sidebar_current: "docs-oci-resource-data_safe-alert"
description: |-
	Provides the Alert resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_alert
This resource provides the Alert resource in Oracle Cloud Infrastructure Data Safe service.

Updates the status of the specified alert.

## Example Usage

```hcl
resource "oci_data_safe_alert" "test_alert" {
	#Required
	alert_id = oci_data_safe_alert.test_alert.id

	#Optional
	comment = var.alert_comment
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
	status = var.alert_status
}
```

## Argument Reference

The following arguments are supported:

* `alert_id` - (Required) The OCID of alert.
* `comment` - (Optional) (Updatable) A comment can be entered to track the alert changes done by the user.
* `compartment_id` - (Optional) (Updatable) The OCID of the compartment that contains the alert.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}`
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}`
* `status` - (Optional) (Updatable) The status of the alert.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `alert_policy_rule_key` - The key of the rule of alert policy that triggered alert.
* `alert_policy_rule_name` - The display name of the rule of alert policy that triggered alert.
* `alert_type` - Type of the alert. Indicates the Data Safe feature triggering the alert.
* `comment` - A comment for the alert. Entered by the user.
* `compartment_id` - The OCID of the compartment that contains the alert.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}`
* `description` - The description of the alert.
* `display_name` - The display name of the alert.
* `feature_details` - Map that contains maps of values. Example: `{"Operations": {"CostCenter": "42"}}`
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}`
* `id` - The OCID of the alert.
* `operation` - The operation (event) that triggered alert.
* `operation_status` - The result of the operation (event) that triggered alert.
* `operation_time` - Creation date and time of the operation that triggered alert, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `policy_id` - The OCID of the policy that triggered alert.
* `resource_name` - The resource endpoint that triggered the alert.
* `severity` - Severity level of the alert.
* `state` - The current state of the alert.
* `status` - The status of the alert.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `target_ids` - Array of OCIDs of the target database which are associated with the alert.
* `target_names` - Array of names of the target database.
* `time_created` - Creation date and time of the alert, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_updated` - Last date and time the alert was updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
* `create` - (Defaults to 20 minutes), when creating the Alert
* `update` - (Defaults to 20 minutes), when updating the Alert
* `delete` - (Defaults to 20 minutes), when destroying the Alert


## Import

Alerts can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_alert.test_alert "id"
```
