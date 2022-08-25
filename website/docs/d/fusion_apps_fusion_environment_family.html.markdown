---
subcategory: "Fusion Apps"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fusion_apps_fusion_environment_family"
sidebar_current: "docs-oci-datasource-fusion_apps-fusion_environment_family"
description: |-
  Provides details about a specific Fusion Environment Family in Oracle Cloud Infrastructure Fusion Apps service
---

# Data Source: oci_fusion_apps_fusion_environment_family
This data source provides details about a specific Fusion Environment Family resource in Oracle Cloud Infrastructure Fusion Apps service.

Retrieves a fusion environment family identified by its OCID.

## Example Usage

```hcl
data "oci_fusion_apps_fusion_environment_family" "test_fusion_environment_family" {
	#Required
	fusion_environment_family_id = oci_fusion_apps_fusion_environment_family.test_fusion_environment_family.id
}
```

## Argument Reference

The following arguments are supported:

* `fusion_environment_family_id` - (Required) The unique identifier (OCID) of the FusionEnvironmentFamily.


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

