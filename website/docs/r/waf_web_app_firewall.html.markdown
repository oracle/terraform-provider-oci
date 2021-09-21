---
subcategory: "Waf"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_waf_web_app_firewall"
sidebar_current: "docs-oci-resource-waf-web_app_firewall"
description: |-
  Provides the Web App Firewall resource in Oracle Cloud Infrastructure Waf service
---

# oci_waf_web_app_firewall
This resource provides the Web App Firewall resource in Oracle Cloud Infrastructure Waf service.

Creates a new WebAppFirewall.


## Example Usage

```hcl
resource "oci_waf_web_app_firewall" "test_web_app_firewall" {
	#Required
	backend_type = var.web_app_firewall_backend_type
	compartment_id = var.compartment_id
	load_balancer_id = oci_load_balancer_load_balancer.test_load_balancer.id
	web_app_firewall_policy_id = oci_waf_web_app_firewall_policy.test_web_app_firewall_policy.id

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	display_name = var.web_app_firewall_display_name
	freeform_tags = {"bar-key"= "value"}
	system_tags = var.web_app_firewall_system_tags
}
```

## Argument Reference

The following arguments are supported:

* `backend_type` - (Required) Type of the WebAppFirewall, as example LOAD_BALANCER.
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Optional) (Updatable) WebAppFirewall display name, can be renamed.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `load_balancer_id` - (Required) LoadBalancer [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) to which the WebAppFirewallPolicy is attached to.
* `system_tags` - (Optional) (Updatable) Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `web_app_firewall_policy_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of WebAppFirewallPolicy, which is attached to the resource.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `backend_type` - Type of the WebAppFirewall, as example LOAD_BALANCER.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - WebAppFirewall display name, can be renamed.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the WebAppFirewall.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in FAILED state. 
* `load_balancer_id` - LoadBalancer [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) to which the WebAppFirewallPolicy is attached to.
* `state` - The current state of the WebAppFirewall.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the WebAppFirewall was created. An RFC3339 formatted datetime string.
* `time_updated` - The time the WebAppFirewall was updated. An RFC3339 formatted datetime string.
* `web_app_firewall_policy_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of WebAppFirewallPolicy, which is attached to the resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Web App Firewall
	* `update` - (Defaults to 20 minutes), when updating the Web App Firewall
	* `delete` - (Defaults to 20 minutes), when destroying the Web App Firewall


## Import

WebAppFirewalls can be imported using the `id`, e.g.

```
$ terraform import oci_waf_web_app_firewall.test_web_app_firewall "id"
```

