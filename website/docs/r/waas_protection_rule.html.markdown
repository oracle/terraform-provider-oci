---
subcategory: "Web Application Acceleration and Security"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_waas_protection_rule"
sidebar_current: "docs-oci-resource-waas-protection_rule"
description: |-
  Provides the Protection Rule resource in Oracle Cloud Infrastructure Web Application Acceleration and Security service
---

# oci_waas_protection_rule
This resource provides the Protection Rule resource in Oracle Cloud Infrastructure Web Application Acceleration and Security service.

Updates the action for each specified protection rule. Requests can either be allowed, blocked, or trigger an alert if they meet the parameters of an applied rule. For more information on protection rules, see [WAF Protection Rules](https://docs.cloud.oracle.com/iaas/Content/WAF/Tasks/wafprotectionrules.htm).
This operation can update or disable protection rules depending on the structure of the request body.
Protection rules can be updated by changing the properties of the protection rule object with the rule's key specified in the key field.

## Example Usage

```hcl
resource "oci_waas_protection_rule" "test_protection_rule" {
	#Required
	waas_policy_id = oci_waas_waas_policy.test_waas_policy.id
    key            = var.key
    
    #Optional
    action         = "DETECT"
      exclusions = {
        exclusions = ["example.com"]
        target     = "REQUEST_COOKIES"
      }
}
```

## Argument Reference

The following arguments are supported:

* `waas_policy_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the WAAS policy.
* `key` - (Required) (Updatable) The unique key of the protection rule.
* `action` - (Optional) (Updatable) The action to take when the traffic is detected as malicious. If unspecified, defaults to `OFF`.
* `exclusions` - (Optional) (Updatable)
	* `exclusions` - An array of The target property of a request that would allow it to bypass the protection rule. For example, when `target` is `REQUEST_COOKIE_NAMES`, the list may include names of cookies to exclude from the protection rule. When the target is `ARGS`, the list may include strings of URL query parameters and values from form-urlencoded XML, JSON, AMP, or POST payloads to exclude from the protection rule. `Exclusions` properties must not contain whitespace, comma or |. **Note:** If protection rules have been enabled that utilize the `maxArgumentCount` or `maxTotalNameLengthOfArguments` properties, and the `target` property has been set to `ARGS`, it is important that the `exclusions` properties be defined to honor those protection rule settings in a consistent manner.
	* `target` - The target of the exclusion.

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Protection Rule
	* `update` - (Defaults to 20 minutes), when updating the Protection Rule
	* `delete` - (Defaults to 20 minutes), when destroying the Protection Rule


## Import

ProtectionRules can be imported using the `id`, e.g.
```
$ terraform import oci_waas_protection_rule.test_protection_rule "waasPolicyId/{waasPolicyId}/key/{key}" 
```

