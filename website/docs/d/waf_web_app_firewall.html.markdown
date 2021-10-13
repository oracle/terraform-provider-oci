---
subcategory: "Waf"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_waf_web_app_firewall"
sidebar_current: "docs-oci-datasource-waf-web_app_firewall"
description: |-
  Provides details about a specific Web App Firewall in Oracle Cloud Infrastructure Waf service
---

# Data Source: oci_waf_web_app_firewall
This data source provides details about a specific Web App Firewall resource in Oracle Cloud Infrastructure Waf service.

Gets a WebAppFirewall by OCID.

## Example Usage

```hcl
data "oci_waf_web_app_firewall" "test_web_app_firewall" {
	#Required
	web_app_firewall_id = oci_waf_web_app_firewall.test_web_app_firewall.id
}
```

## Argument Reference

The following arguments are supported:

* `web_app_firewall_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the WebAppFirewall.


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

