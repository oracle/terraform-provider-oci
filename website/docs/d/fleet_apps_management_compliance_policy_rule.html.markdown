---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_compliance_policy_rule"
sidebar_current: "docs-oci-datasource-fleet_apps_management-compliance_policy_rule"
description: |-
  Provides details about a specific Compliance Policy Rule in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_compliance_policy_rule
This data source provides details about a specific Compliance Policy Rule resource in Oracle Cloud Infrastructure Fleet Apps Management service.

Gets information about a CompliancePolicyRule.

## Example Usage

```hcl
data "oci_fleet_apps_management_compliance_policy_rule" "test_compliance_policy_rule" {
	#Required
	compliance_policy_rule_id = oci_fleet_apps_management_compliance_policy_rule.test_compliance_policy_rule.id
}
```

## Argument Reference

The following arguments are supported:

* `compliance_policy_rule_id` - (Required) unique CompliancePolicyRule identifier.


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

