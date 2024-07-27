---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_alert_policy"
sidebar_current: "docs-oci-datasource-data_safe-alert_policy"
description: |-
  Provides details about a specific Alert Policy in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_alert_policy
This data source provides details about a specific Alert Policy resource in Oracle Cloud Infrastructure Data Safe service.

Gets the details of alert policy by its ID.

## Example Usage

```hcl
data "oci_data_safe_alert_policy" "test_alert_policy" {
  #Required
  alert_policy_id = oci_data_safe_alert_policy.test_alert_policy.id
}
```

## Argument Reference

The following arguments are supported:

* `alert_policy_id` - (Required) The OCID of the alert policy.


## Attributes Reference

The following attributes are exported:

* `alert_policy_type` - Indicates the Data Safe feature to which the alert policy belongs.
* `compartment_id` - The OCID of the compartment that contains the alert policy.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}`
* `description` - The description of the alert policy.
* `display_name` - The display name of the alert policy.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}`
* `id` - The OCID of the alert policy.
* `is_user_defined` - Indicates if the alert policy is user-defined (true) or pre-defined (false).
* `lifecycle_details` - Details about the current state of the alert policy.
* `severity` - Severity level of the alert raised by this policy.
* `state` - The current state of the alert.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `time_created` - Creation date and time of the alert policy, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_updated` - Last date and time the alert policy was updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
