---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_target_alert_policy_association_unassociated_target_members"
sidebar_current: "docs-oci-datasource-data_safe-target_alert_policy_association_unassociated_target_members"
description: |-
  Provides the list of Target Alert Policy Association Unassociated Target Members in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_target_alert_policy_association_unassociated_target_members
This data source provides the list of Target Alert Policy Association Unassociated Target Members in Oracle Cloud Infrastructure Data Safe service.

Gets the details of target-alert policy association and its unassociated members by its ID.

## Example Usage

```hcl
data "oci_data_safe_target_alert_policy_association_unassociated_target_members" "test_target_alert_policy_association_unassociated_target_members" {
	#Required
	target_alert_policy_association_id = oci_data_safe_target_alert_policy_association.test_target_alert_policy_association.id
}
```

## Argument Reference

The following arguments are supported:

* `target_alert_policy_association_id` - (Required) The OCID of the target-alert policy association.


## Attributes Reference

The following attributes are exported:

* `target_alert_policy_unassociated_collection` - The list of target_alert_policy_unassociated_collection.

### TargetAlertPolicyAssociationUnassociatedTargetMember Reference

The following attributes are exported:

* `items` - Array of unassociated target alert policy association summary.
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm) Example: `{"Operations.CostCenter": "42"}`
	* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}`
	* `is_enabled` - Indicates if the target-alert policy association is enabled or disabled by user.
	* `not_applied_reason` - Details on why policy is not applied on target.
	* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}`
	* `target_database_id` - The OCID of the target database that differs from the alert policy association of the target database group.
