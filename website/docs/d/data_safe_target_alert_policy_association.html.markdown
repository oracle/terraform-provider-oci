---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_target_alert_policy_association"
sidebar_current: "docs-oci-datasource-data_safe-target_alert_policy_association"
description: |-
  Provides details about a specific Target Alert Policy Association in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_target_alert_policy_association
This data source provides details about a specific Target Alert Policy Association resource in Oracle Cloud Infrastructure Data Safe service.

Gets the details of target-alert policy association by its ID.

## Example Usage

```hcl
data "oci_data_safe_target_alert_policy_association" "test_target_alert_policy_association" {
	#Required
	target_alert_policy_association_id = oci_data_safe_target_alert_policy_association.test_target_alert_policy_association.id
}
```

## Argument Reference

The following arguments are supported:

* `target_alert_policy_association_id` - (Required) The OCID of the target-alert policy association.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the policy.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - Describes the target-alert policy association.
* `display_name` - The display name of the target-alert policy association.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the target-alert policy association.
* `is_enabled` - Indicates if the target-alert policy association is enabled or disabled by user.
* `policy_id` - The OCID of the alert policy.
* `state` - The current state of the target-alert policy association.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `target_id` - The OCID of the target on which alert policy is to be applied.
* `time_created` - Creation date and time of the alert policy, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_updated` - Last date and time the alert policy was updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).

