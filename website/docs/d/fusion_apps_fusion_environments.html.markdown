---
subcategory: "Fusion Apps"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fusion_apps_fusion_environments"
sidebar_current: "docs-oci-datasource-fusion_apps-fusion_environments"
description: |-
  Provides the list of Fusion Environments in Oracle Cloud Infrastructure Fusion Apps service
---

# Data Source: oci_fusion_apps_fusion_environments
This data source provides the list of Fusion Environments in Oracle Cloud Infrastructure Fusion Apps service.

Returns a list of FusionEnvironments.


## Example Usage

```hcl
data "oci_fusion_apps_fusion_environments" "test_fusion_environments" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.fusion_environment_display_name
	fusion_environment_family_id = oci_fusion_apps_fusion_environment_family.test_fusion_environment_family.id
	state = var.fusion_environment_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `fusion_environment_family_id` - (Optional) The ID of the fusion environment family in which to list resources.
* `state` - (Optional) A filter that returns all resources that match the specified lifecycle state.


## Attributes Reference

The following attributes are exported:

* `fusion_environment_collection` - The list of fusion_environment_collection.

### FusionEnvironment Reference

The following attributes are exported:

* `additional_language_packs` - Language packs
* `applied_patch_bundles` - Patch bundle names
* `compartment_id` - Compartment Identifier
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - FusionEnvironment Identifier, can be renamed
* `dns_prefix` - DNS prefix
* `domain_id` - The IDCS domain created for the fusion instance
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `fusion_environment_family_id` - FusionEnvironmentFamily Identifier
* `fusion_environment_type` - Type of the FusionEnvironment.
* `id` - Unique identifier that is immutable on creation
* `idcs_domain_url` - The IDCS Domain URL
* `kms_key_id` - BYOK key id
* `kms_key_info` - BYOK key info
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `maintenance_policy` - The policy that specifies the maintenance and upgrade preferences for an environment. For more information about the options, see [Understanding Environment Maintenance](https://docs.cloud.oracle.com/iaas/Content/fusion-applications/plan-environment-family.htm#about-env-maintenance).
	* `environment_maintenance_override` - User choice to upgrade both production and non-production environments at the same time. Overrides the Fusion environment family setting.
	* `monthly_patching_override` - Whether the Fusion environment will be updated monthly or updated on the quarterly cycle. This setting overrides the monthly patching setting of its Fusion environment family.
	* `quarterly_upgrade_begin_times` - Determines the quarterly upgrade begin times (monthly maintenance group schedule ) of the Fusion environment.
		* `begin_times_value` - The frequency and month when maintenance occurs for the Fusion environment.
		* `override_type` - Determines if the maintenance schedule of the Fusion environment is inherited from the Fusion environment family.
* `public_url` - Public URL
* `refresh` - Describes a refresh of a fusion environment
	* `source_fusion_environment_id` - The source environment id for the last refresh
	* `time_finished` - The time of when the last refresh finish
	* `time_of_restoration_point` - The point of time of the latest DB backup for the last refresh
* `rules` - Network Access Control Rules
	* `action` - Rule type
	* `conditions` - 
		* `attribute_name` - RuleCondition type
		* `attribute_value` - The OCID of the originating VCN that an incoming packet must match. You can use this condition in conjunction with `SourceVcnIpAddressCondition`. **NOTE:** If you define this condition for a rule without a `SourceVcnIpAddressCondition`, this condition matches all incoming traffic in the specified VCN. 
	* `description` - A brief description of the access control rule. Avoid entering confidential information. example: `192.168.0.0/16 and 2001:db8::/32 are trusted clients. Whitelist them.` 
* `state` - The current state of the ServiceInstance.
* `subscription_ids` - List of subscription IDs.
* `system_name` - Environment Specific Guid/ System Name
* `time_created` - The time the the FusionEnvironment was created. An RFC3339 formatted datetime string
* `time_upcoming_maintenance` - The next maintenance for this environment
* `time_updated` - The time the FusionEnvironment was updated. An RFC3339 formatted datetime string
* `version` - Version of Fusion Apps used by this environment

