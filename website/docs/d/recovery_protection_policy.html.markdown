---
subcategory: "Recovery"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_recovery_protection_policy"
sidebar_current: "docs-oci-datasource-recovery-protection_policy"
description: |-
  Provides details about a specific Protection Policy in Oracle Cloud Infrastructure Recovery service
---

# Data Source: oci_recovery_protection_policy
This data source provides details about a specific Protection Policy resource in Oracle Cloud Infrastructure Recovery service.

Gets information about a specified protection policy.

## Example Usage

```hcl
data "oci_recovery_protection_policy" "test_protection_policy" {
	#Required
	protection_policy_id = oci_recovery_protection_policy.test_protection_policy.id
}
```

## Argument Reference

The following arguments are supported:

* `protection_policy_id` - (Required) The protection policy OCID.


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
* `policy_locked_date_time` - An RFC3339 formatted datetime string that specifies the exact date and time for the retention lock to take effect and permanently lock the retention period defined in the policy. 
* `state` - The current state of the protection policy. 
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}`. For more information, see [Resource Tags](https://docs.oracle.com/en-us/iaas/Content/General/Concepts/resourcetags.htm) 
* `time_created` - An RFC3339 formatted datetime string that indicates the created time for the protection policy. For example: '2020-05-22T21:10:29.600Z'. 
* `time_updated` - An RFC3339 formatted datetime string that indicates the updated time for the protection policy. For example: '2020-05-22T21:10:29.600Z'. 

