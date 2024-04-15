---
subcategory: "Fusion Apps"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fusion_apps_fusion_environment"
sidebar_current: "docs-oci-resource-fusion_apps-fusion_environment"
description: |-
  Provides the Fusion Environment resource in Oracle Cloud Infrastructure Fusion Apps service
---

# oci_fusion_apps_fusion_environment
This resource provides the Fusion Environment resource in Oracle Cloud Infrastructure Fusion Apps service.

Creates a new FusionEnvironment.


## Example Usage

```hcl
resource "oci_fusion_apps_fusion_environment" "test_fusion_environment" {
	#Required
	compartment_id = var.compartment_id
	create_fusion_environment_admin_user_details {
		#Required
		email_address = var.fusion_environment_create_fusion_environment_admin_user_details_email_address
		first_name = var.fusion_environment_create_fusion_environment_admin_user_details_first_name
		last_name = var.fusion_environment_create_fusion_environment_admin_user_details_last_name
		username = var.fusion_environment_create_fusion_environment_admin_user_details_username

		#Optional
		password = var.fusion_environment_create_fusion_environment_admin_user_details_password
	}
	display_name = var.fusion_environment_display_name
	fusion_environment_family_id = oci_fusion_apps_fusion_environment_family.test_fusion_environment_family.id
	fusion_environment_type = var.fusion_environment_fusion_environment_type

	#Optional
	additional_language_packs = var.fusion_environment_additional_language_packs
	defined_tags = {"foo-namespace.bar-key"= "value"}
	dns_prefix = var.fusion_environment_dns_prefix
	freeform_tags = {"bar-key"= "value"}
	kms_key_id = oci_kms_key.test_key.id
	maintenance_policy {

		#Optional
		environment_maintenance_override = var.fusion_environment_maintenance_policy_environment_maintenance_override
		monthly_patching_override = var.fusion_environment_maintenance_policy_monthly_patching_override
	}
	rules {
		#Required
		action = var.fusion_environment_rules_action
		conditions {
			#Required
			attribute_name = var.fusion_environment_rules_conditions_attribute_name
			attribute_value = var.fusion_environment_rules_conditions_attribute_value
		}

		#Optional
		description = var.fusion_environment_rules_description
	}
}
```

## Argument Reference

The following arguments are supported:

* `additional_language_packs` - (Optional) (Updatable) Language packs.
* `compartment_id` - (Required) (Updatable) The unique identifier (OCID) of the compartment where the Fusion Environment is located.
* `create_fusion_environment_admin_user_details` - (Required) The credentials for the Fusion Applications service administrator.
	* `email_address` - (Required) The email address for the administrator.
	* `first_name` - (Required) The administrator's first name.
	* `last_name` - (Required) The administrator's last name.
	* `password` - (Optional) The password for the administrator.
	* `username` - (Required) The username for the administrator.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Required) (Updatable) FusionEnvironment Identifier can be renamed.
* `dns_prefix` - (Optional) DNS prefix.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `fusion_environment_family_id` - (Required) The unique identifier (OCID) of the Fusion Environment Family that the Fusion Environment belongs to.
* `fusion_environment_type` - (Required) The type of environment. Valid values are Production, Test, or Development.
* `kms_key_id` - (Optional) (Updatable) byok kms keyId
* `maintenance_policy` - (Optional) (Updatable) The policy that specifies the maintenance and upgrade preferences for an environment. For more information about the options, see [Understanding Environment Maintenance](https://docs.cloud.oracle.com/iaas/Content/fusion-applications/plan-environment-family.htm#about-env-maintenance).
	* `environment_maintenance_override` - (Optional) (Updatable) User choice to upgrade both test and prod pods at the same time. Overrides fusion environment families'.
	* `monthly_patching_override` - (Optional) (Updatable) When "ENABLED", the Fusion environment is patched monthly. When "DISABLED", the Fusion environment is not patched monthly. This setting overrides the environment family setting. When not set, the environment follows the environment family policy.
* `rules` - (Optional) (Updatable) Rules.
	* `action` - (Required) (Updatable) Rule type
	* `conditions` - (Required) (Updatable) 
		* `attribute_name` - (Required) (Updatable) RuleCondition type
		* `attribute_value` - (Required) (Updatable) The OCID of the originating VCN that an incoming packet must match. You can use this condition in conjunction with `SourceVcnIpAddressCondition`. **NOTE:** If you define this condition for a rule without a `SourceVcnIpAddressCondition`, this condition matches all incoming traffic in the specified VCN. 
	* `description` - (Optional) (Updatable) A brief description of the access control rule. Avoid entering confidential information. example: `192.168.0.0/16 and 2001:db8::/32 are trusted clients. Whitelist them.` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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
* `is_break_glass_enabled` - If it's true, then the Break Glass feature is enabled
* `kms_key_id` - BYOK key id
* `kms_key_info` - BYOK key info
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `lockbox_id` - The lockbox Id of this fusion environment. If there's no lockbox id, this field will be null
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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Fusion Environment
	* `update` - (Defaults to 20 minutes), when updating the Fusion Environment
	* `delete` - (Defaults to 20 minutes), when destroying the Fusion Environment


## Import

FusionEnvironments can be imported using the `id`, e.g.

```
$ terraform import oci_fusion_apps_fusion_environment.test_fusion_environment "id"
```

