---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_target_alert_policy_associations"
sidebar_current: "docs-oci-datasource-data_safe-target_alert_policy_associations"
description: |-
  Provides the list of Target Alert Policy Associations in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_target_alert_policy_associations
This data source provides the list of Target Alert Policy Associations in Oracle Cloud Infrastructure Data Safe service.

Gets a list of all target-alert policy associations.


## Example Usage

```hcl
data "oci_data_safe_target_alert_policy_associations" "test_target_alert_policy_associations" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.target_alert_policy_association_access_level
	alert_policy_id = oci_data_safe_alert_policy.test_alert_policy.id
	compartment_id_in_subtree = var.target_alert_policy_association_compartment_id_in_subtree
	state = var.target_alert_policy_association_state
	target_alert_policy_association_id = oci_data_safe_target_alert_policy_association.test_target_alert_policy_association.id
	target_id = oci_cloud_guard_target.test_target.id
	time_created_greater_than_or_equal_to = var.target_alert_policy_association_time_created_greater_than_or_equal_to
	time_created_less_than = var.target_alert_policy_association_time_created_less_than
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
* `alert_policy_id` - (Optional) A filter to return policy by it's OCID.
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
* `state` - (Optional) An optional filter to return only alert policies that have the given life-cycle state.
* `target_alert_policy_association_id` - (Optional) A filter to return only items related to a specific target-alert policy association ID.
* `target_id` - (Optional) A filter to return only items related to a specific target OCID.
* `time_created_greater_than_or_equal_to` - (Optional) A filter to return only the resources that were created after the specified date and time, as defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Using TimeCreatedGreaterThanOrEqualToQueryParam parameter retrieves all resources created after that date.

  **Example:** 2016-12-19T16:39:57.600Z
* `time_created_less_than` - (Optional) Search for resources that were created before a specific date. Specifying this parameter corresponding `timeCreatedLessThan` parameter will retrieve all resources created before the specified created date, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by RFC 3339.

  **Example:** 2016-12-19T16:39:57.600Z


## Attributes Reference

The following attributes are exported:

* `target_alert_policy_association_collection` - The list of target_alert_policy_association_collection.

### TargetAlertPolicyAssociation Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the policy.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}`
* `description` - Describes the target-alert policy association.
* `display_name` - The display name of the target-alert policy association.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}`
* `id` - The OCID of the target-alert policy association.
* `is_enabled` - Indicates if the target-alert policy association is enabled or disabled by user.
* `lifecycle_details` - Details about the current state of the target-alert policy association.
* `policy_id` - The OCID of the alert policy.
* `state` - The current state of the target-alert policy association.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `target_id` - The OCID of the target on which alert policy is to be applied.
* `time_created` - Creation date and time of the alert policy, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_updated` - Last date and time the alert policy was updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
