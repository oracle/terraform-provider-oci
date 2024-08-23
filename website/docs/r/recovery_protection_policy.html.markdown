---
subcategory: "Recovery"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_recovery_protection_policy"
sidebar_current: "docs-oci-resource-recovery-protection_policy"
description: |-
  Provides the Protection Policy resource in Oracle Cloud Infrastructure Recovery service
---

# oci_recovery_protection_policy
This resource provides the Protection Policy resource in Oracle Cloud Infrastructure Recovery service.

Creates a new Protection Policy.


## Example Usage

```hcl
resource "oci_recovery_protection_policy" "test_protection_policy" {
	#Required
	backup_retention_period_in_days = var.protection_policy_backup_retention_period_in_days
	compartment_id = var.compartment_id
	display_name = var.protection_policy_display_name

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	must_enforce_cloud_locality = var.protection_policy_must_enforce_cloud_locality
	policy_locked_date_time = var.protection_policy_policy_locked_date_time
}
```

## Argument Reference

The following arguments are supported:

* `backup_retention_period_in_days` - (Required) (Updatable) The maximum number of days to retain backups for a protected database.
* `compartment_id` - (Required) (Updatable) Compartment Identifier
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`. For more information, see [Resource Tags](https://docs.oracle.com/en-us/iaas/Content/General/Concepts/resourcetags.htm) 
* `display_name` - (Required) (Updatable) A user provided name for the protection policy. The 'displayName' does not have to be unique, and it can be modified. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `must_enforce_cloud_locality` - (Optional) Indicates whether the protection policy enforces Recovery Service to retain backups in the same cloud service environment where your Oracle Database is provisioned. This parameter is applicable if your Oracle Database runs in a different cloud service environment, such as Microsoft Azure. If you set the mustEnforceCloudLocality parameter to TRUE, then Recovery Service stores the database backups locally in the same cloud service environment where the database resides. For example, if your Oracle Database is provisioned on Microsoft Azure, then Recovery Service stores the database backups in Azure. Note: You cannot change the mustEnforceCloudLocality setting for a protection policy after you create it. 
* `policy_locked_date_time` - (Optional) (Updatable) An RFC3339 formatted datetime string that specifies the exact date and time for the retention lock to take effect and permanently lock the retention period defined in the policy.
	* The retention lock feature controls whether Recovery Service strictly preserves backups for the duration defined in a policy. Retention lock is useful to enforce recovery window compliance and to prevent unintentional modifications to protected database backups.
	* Recovery Service enforces a 14-day delay before the retention lock set for a policy can take effect. Therefore, you must set policyLockedDateTime  to a date that occurs 14 days after the current date.
	* For example, assuming that the current date is Aug 1, 2023 9 pm, you can set policyLockedDateTime  to '2023-08-15T21:00:00.600Z' (Aug 15, 2023, 9:00 pm), or greater.
	* During the 14-day delay period, you can either increase or decrease the retention period in the policy.
	* However, you are only allowed to increase the retention period on or after the retention lock date.
	* You cannot change the value of policyLockedDateTime if the retention lock is already in effect. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `backup_retention_period_in_days` - The maximum number of days to retain backups for a protected database. Specify a period ranging from a minimum 14 days to a maximum 95 days. For example, specify the value 55 if you want to retain backups for 55 days.
* `compartment_id` - The OCID of the compartment that contains the protection policy.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`. For more information, see [Resource Tags](https://docs.oracle.com/en-us/iaas/Content/General/Concepts/resourcetags.htm) 
* `display_name` - A user provided name for the protection policy.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The protection policy OCID.
* `is_predefined_policy` - Set to TRUE if the policy is Oracle-defined, and FALSE for a user-defined custom policy. You can modify only the custom policies.
* `lifecycle_details` - Detailed description about the current lifecycle state of the protection policy. For example, it can be used to provide actionable information for a resource in a Failed state.
* `must_enforce_cloud_locality` - Indicates whether the protection policy enforces Recovery Service to retain backups in the same cloud service environment where your Oracle Database is provisioned.
* `policy_locked_date_time` - An RFC3339 formatted datetime string that specifies the exact date and time for the retention lock to take effect and permanently lock the retention period defined in the policy. 
* `state` - The current state of the protection policy. 
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}`. For more information, see [Resource Tags](https://docs.oracle.com/en-us/iaas/Content/General/Concepts/resourcetags.htm) 
* `time_created` - An RFC3339 formatted datetime string that indicates the created time for the protection policy. For example: '2020-05-22T21:10:29.600Z'. 
* `time_updated` - An RFC3339 formatted datetime string that indicates the updated time for the protection policy. For example: '2020-05-22T21:10:29.600Z'. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Protection Policy
	* `update` - (Defaults to 20 minutes), when updating the Protection Policy
	* `delete` - (Defaults to 20 minutes), when destroying the Protection Policy


## Import

ProtectionPolicies can be imported using the `id`, e.g.

```
$ terraform import oci_recovery_protection_policy.test_protection_policy "id"
```

