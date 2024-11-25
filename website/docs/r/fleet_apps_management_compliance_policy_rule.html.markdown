---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_compliance_policy_rule"
sidebar_current: "docs-oci-resource-fleet_apps_management-compliance_policy_rule"
description: |-
  Provides the Compliance Policy Rule resource in Oracle Cloud Infrastructure Fleet Apps Management service
---

# oci_fleet_apps_management_compliance_policy_rule
This resource provides the Compliance Policy Rule resource in Oracle Cloud Infrastructure Fleet Apps Management service.

Creates a CompliancePolicyRule.


## Example Usage

```hcl
resource "oci_fleet_apps_management_compliance_policy_rule" "test_compliance_policy_rule" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.compliance_policy_rule_display_name
	patch_selection {
		#Required
		selection_type = var.compliance_policy_rule_patch_selection_selection_type

		#Optional
		days_since_release = var.compliance_policy_rule_patch_selection_days_since_release
		patch_level = var.compliance_policy_rule_patch_selection_patch_level
		patch_name = oci_fleet_apps_management_patch.test_patch.name
	}
	patch_type = var.compliance_policy_rule_patch_type
	product_version {
		#Required
		version = var.compliance_policy_rule_product_version_version

		#Optional
		is_applicable_for_all_higher_versions = var.compliance_policy_rule_product_version_is_applicable_for_all_higher_versions
	}

	#Optional
	compliance_policy_id = oci_fleet_apps_management_compliance_policy.test_compliance_policy.id
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	grace_period = var.compliance_policy_rule_grace_period
	severity = var.compliance_policy_rule_severity
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment the CompliancePolicyRule belongs to.
* `compliance_policy_id` - (Optional) Unique OCID of the CompliancePolicy this CompliancePolicyRule belongs to.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Required) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `grace_period` - (Optional) (Updatable) Grace period in days,weeks,months or years the exemption is applicable for the rule. This enables a grace period when Fleet Application Management doesn't report the product as noncompliant when patch is not applied. 
* `patch_selection` - (Required) (Updatable) Patch Selection Details
	* `days_since_release` - (Required when selection_type=PATCH_RELEASE_DATE) (Updatable) Days passed since patch release.
	* `patch_level` - (Required when selection_type=PATCH_LEVEL) (Updatable) Patch Name.
	* `patch_name` - (Required when selection_type=PATCH_NAME) (Updatable) Patch Name.
	* `selection_type` - (Required) (Updatable) Selection type for the Patch. 
* `patch_type` - (Required) (Updatable) PlatformConfiguration OCID for the patch type to which this CompliancePolicyRule applies.
* `product_version` - (Required) (Updatable) A specific product version or a specific version and succeeding. Example: 12.1 or 12.1 and above for Oracle WebLogic Application server. The policy applies to the next version only, and not to other versions such as, 12.1.x. 
	* `is_applicable_for_all_higher_versions` - (Optional) (Updatable) Is rule applicable to all higher versions also
	* `version` - (Required) (Updatable) Product version the rule is applicable.
* `severity` - (Optional) (Updatable) Severity to which this CompliancePolicyRule applies.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment the CompliancePolicyRule belongs to.
* `compliance_policy_id` - Unique OCID of the CompliancePolicy.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `grace_period` - Grace period in days,weeks,months or years the exemption is applicable for the rule. This enables a grace period when Fleet Application Management doesn't report the product as noncompliant when patch is not applied. 
* `id` - Unique OCID of the CompliancePolicyRule.
* `lifecycle_details` - A message that describes the current state of the CompliancePolicyRule in more detail. For example, can be used to provide actionable information for a resource in the Failed state. 
* `patch_selection` - Patch Selection Details
	* `days_since_release` - Days passed since patch release.
	* `patch_level` - Patch Name.
	* `patch_name` - Patch Name.
	* `selection_type` - Selection type for the Patch. 
* `patch_type` - PlatformConfiguration OCID for the patch type to which this CompliancePolicyRule applies.
* `product_version` - A specific product version or a specific version and succeeding. Example: 12.1 or 12.1 and above for Oracle WebLogic Application server. The policy applies to the next version only, and not to other versions such as, 12.1.x. 
	* `is_applicable_for_all_higher_versions` - Is rule applicable to all higher versions also
	* `version` - Product version the rule is applicable.
* `severity` - Severity to which this CompliancePolicyRule applies.
* `state` - The current state of the CompliancePolicyRule.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the CompliancePolicyRule was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the CompliancePolicyRule was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Compliance Policy Rule
	* `update` - (Defaults to 20 minutes), when updating the Compliance Policy Rule
	* `delete` - (Defaults to 20 minutes), when destroying the Compliance Policy Rule


## Import

CompliancePolicyRules can be imported using the `id`, e.g.

```
$ terraform import oci_fleet_apps_management_compliance_policy_rule.test_compliance_policy_rule "id"
```

