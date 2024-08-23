---
subcategory: "Recovery"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_recovery_protection_policies"
sidebar_current: "docs-oci-datasource-recovery-protection_policies"
description: |-
  Provides the list of Protection Policies in Oracle Cloud Infrastructure Recovery service
---

# Data Source: oci_recovery_protection_policies
This data source provides the list of Protection Policies in Oracle Cloud Infrastructure Recovery service.

Gets a list of protection policies based on the specified parameters.


## Example Usage

```hcl
data "oci_recovery_protection_policies" "test_protection_policies" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.protection_policy_display_name
	owner = var.protection_policy_owner
	protection_policy_id = oci_recovery_protection_policy.test_protection_policy.id
	state = var.protection_policy_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment OCID.
* `display_name` - (Optional) A filter to return only resources that match the entire 'displayname' given.
* `owner` - (Optional) A filter to return only the policies that match the owner as 'Customer' or 'Oracle'.
* `protection_policy_id` - (Optional) The protection policy OCID.
* `state` - (Optional) A filter to return only resources their lifecycleState matches the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `protection_policy_collection` - The list of protection_policy_collection.

### ProtectionPolicy Reference

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

