---
subcategory: "Web Application Acceleration and Security"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_waas_protection_rule"
sidebar_current: "docs-oci-datasource-waas-protection_rule"
description: |-
  Provides details about a specific Protection Rule in Oracle Cloud Infrastructure Web Application Acceleration and Security service
---

# Data Source: oci_waas_protection_rule
This data source provides details about a specific Protection Rule resource in Oracle Cloud Infrastructure Web Application Acceleration and Security service.

Gets the details of a protection rule in the Web Application Firewall configuration for a WAAS policy.

## Example Usage

```hcl
data "oci_waas_protection_rule" "test_protection_rule" {
	#Required
	protection_rule_key = "${var.protection_rule_protection_rule_key}"
	waas_policy_id = "${oci_waas_waas_policy.test_waas_policy.id}"
}
```

## Argument Reference

The following arguments are supported:

* `protection_rule_key` - (Required) The protection rule key.
* `waas_policy_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the WAAS policy.


## Attributes Reference

The following attributes are exported:

* `action` - The action to take when the traffic is detected as malicious. If unspecified, defaults to `OFF`.
* `description` - The description of the protection rule.
* `exclusions` - 
	* `exclusions` - An array of The target property of a request that would allow it to bypass the protection rule. For example, when `target` is `REQUEST_COOKIE_NAMES`, the list may include names of cookies to exclude from the protection rule. When the target is `ARGS`, the list may include strings of URL query parameters and values from form-urlencoded XML, JSON, AMP, or POST payloads to exclude from the protection rule. `Exclusions` properties must not contain whitespace, comma or |. **Note:** If protection rules have been enabled that utilize the `maxArgumentCount` or `maxTotalNameLengthOfArguments` properties, and the `target` property has been set to `ARGS`, it is important that the `exclusions` properties be defined to honor those protection rule settings in a consistent manner.
	* `target` - The target of the exclusion.
* `key` - The unique key of the protection rule.
* `labels` - The list of labels for the protection rule.

	**Note:** Protection rules with a `ResponseBody` label will have no effect unless `isResponseInspected` is true.
* `mod_security_rule_ids` - The list of the ModSecurity rule IDs that apply to this protection rule. For more information about ModSecurity's open source WAF rules, see [Mod Security's documentation](https://www.modsecurity.org/CRS/Documentation/index.html).
* `name` - The name of the protection rule.

