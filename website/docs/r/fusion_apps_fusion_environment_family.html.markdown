---
subcategory: "Fusion Apps"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fusion_apps_fusion_environment_family"
sidebar_current: "docs-oci-resource-fusion_apps-fusion_environment_family"
description: |-
  Provides the Fusion Environment Family resource in Oracle Cloud Infrastructure Fusion Apps service
---

# oci_fusion_apps_fusion_environment_family
This resource provides the Fusion Environment Family resource in Oracle Cloud Infrastructure Fusion Apps service.

Creates a new FusionEnvironmentFamily.


## Example Usage

```hcl
resource "oci_fusion_apps_fusion_environment_family" "test_fusion_environment_family" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.fusion_environment_family_display_name
	subscription_ids = var.fusion_environment_family_subscription_ids

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	family_maintenance_policy {

		#Optional
		concurrent_maintenance = var.fusion_environment_family_family_maintenance_policy_concurrent_maintenance
		is_monthly_patching_enabled = var.fusion_environment_family_family_maintenance_policy_is_monthly_patching_enabled
		quarterly_upgrade_begin_times = var.fusion_environment_family_family_maintenance_policy_quarterly_upgrade_begin_times
	}
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment where the environment family is located.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Required) (Updatable) A friendly name for the environment family. The name must contain only letters, numbers, dashes, and underscores. Can be changed later.
* `family_maintenance_policy` - (Optional) (Updatable) The policy that specifies the maintenance and upgrade preferences for an environment. For more information about the options, see [Understanding Environment Maintenance](https://docs.cloud.oracle.com/iaas/Content/fusion-applications/plan-environment-family.htm#about-env-maintenance).
	* `concurrent_maintenance` - (Optional) (Updatable) Option to upgrade both production and non-production environments at the same time. When set to PROD both types of environnments are upgraded on the production schedule. When set to NON_PROD both types of environments are upgraded on the non-production schedule.
	* `is_monthly_patching_enabled` - (Optional) (Updatable) When True, monthly patching is enabled for the environment family.
	* `quarterly_upgrade_begin_times` - (Optional) The quarterly maintenance month group schedule of the Fusion environment family.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `subscription_ids` - (Required) (Updatable) The list of the IDs of the applications subscriptions that are associated with the environment family.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment where the environment family is located.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - A friendly name for the environment family. The name must contain only letters, numbers, dashes, and underscores. Can be changed later.
* `family_maintenance_policy` - The policy that specifies the maintenance and upgrade preferences for an environment. For more information about the options, see [Understanding Environment Maintenance](https://docs.cloud.oracle.com/iaas/Content/fusion-applications/plan-environment-family.htm#about-env-maintenance).
	* `concurrent_maintenance` - Option to upgrade both production and non-production environments at the same time. When set to PROD both types of environnments are upgraded on the production schedule. When set to NON_PROD both types of environments are upgraded on the non-production schedule.
	* `is_monthly_patching_enabled` - When True, monthly patching is enabled for the environment family.
	* `quarterly_upgrade_begin_times` - The quarterly maintenance month group schedule of the Fusion environment family.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The unique identifier (OCID) of the environment family. Can't be changed after creation.
* `is_subscription_update_needed` - When set to True, a subscription update is required for the environment family.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `state` - The current state of the FusionEnvironmentFamily.
* `subscription_ids` - The list of the IDs of the applications subscriptions that are associated with the environment family.
* `system_name` - Environment Specific Guid/ System Name
* `time_created` - The time the the FusionEnvironmentFamily was created. An RFC3339 formatted datetime string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Fusion Environment Family
	* `update` - (Defaults to 20 minutes), when updating the Fusion Environment Family
	* `delete` - (Defaults to 20 minutes), when destroying the Fusion Environment Family


## Import

FusionEnvironmentFamilies can be imported using the `id`, e.g.

```
$ terraform import oci_fusion_apps_fusion_environment_family.test_fusion_environment_family "id"
```

