---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_alert_policies"
sidebar_current: "docs-oci-datasource-data_safe-alert_policies"
description: |-
	Provides the list of Alert Policies in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_alert_policies
This data source provides the list of Alert Policies in Oracle Cloud Infrastructure Data Safe service.

Gets a list of all alert policies.


## Example Usage

```hcl
data "oci_data_safe_alert_policies" "test_alert_policies" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.alert_policy_access_level
	alert_policy_id = oci_data_safe_alert_policy.test_alert_policy.id
	compartment_id_in_subtree = var.alert_policy_compartment_id_in_subtree
	display_name = var.alert_policy_display_name
	is_user_defined = var.alert_policy_is_user_defined
	state = var.alert_policy_state
	time_created_greater_than_or_equal_to = var.alert_policy_time_created_greater_than_or_equal_to
	time_created_less_than = var.alert_policy_time_created_less_than
	type = var.alert_policy_type
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
* `alert_policy_id` - (Optional) A filter to return policy by it's OCID.
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
* `display_name` - (Optional) A filter to return only resources that match the specified display name.
* `is_user_defined` - (Optional) An optional filter to return only alert policies that are user-defined or not.
* `state` - (Optional) An optional filter to return only alert policies that have the given life-cycle state.
* `time_created_greater_than_or_equal_to` - (Optional) A filter to return only the resources that were created after the specified date and time, as defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Using TimeCreatedGreaterThanOrEqualToQueryParam parameter retrieves all resources created after that date.

  **Example:** 2016-12-19T16:39:57.600Z
* `time_created_less_than` - (Optional) Search for resources that were created before a specific date. Specifying this parameter corresponding `timeCreatedLessThan` parameter will retrieve all resources created before the specified created date, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by RFC 3339.

  **Example:** 2016-12-19T16:39:57.600Z
* `type` - (Optional) An optional filter to return only alert policies of a certain type.


## Attributes Reference

The following attributes are exported:

* `alert_policy_collection` - The list of alert_policy_collection.

### AlertPolicy Reference

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
