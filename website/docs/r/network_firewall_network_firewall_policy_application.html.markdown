---
subcategory: "Network Firewall"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_firewall_network_firewall_policy_application"
sidebar_current: "docs-oci-resource-network_firewall-network_firewall_policy_application"
description: |-
  Provides the Network Firewall Policy Application resource in Oracle Cloud Infrastructure Network Firewall service
---

# oci_network_firewall_network_firewall_policy_application
This resource provides the Network Firewall Policy Application resource in Oracle Cloud Infrastructure Network Firewall service.

Creates a new Application inside the Network Firewall Policy.


## Example Usage

```hcl
resource "oci_network_firewall_network_firewall_policy_application" "test_network_firewall_policy_application" {
	#Required
	icmp_type = var.network_firewall_policy_application_icmp_type
	name = var.network_firewall_policy_application_name
	network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
	type = var.network_firewall_policy_application_type

	#Optional
	icmp_code = var.network_firewall_policy_application_icmp_code
}
```

## Argument Reference

The following arguments are supported:

* `icmp_code` - (Optional) (Updatable) The value of the ICMP/ICMP_V6 message Code (subtype) field as defined by [RFC 4443](https://www.rfc-editor.org/rfc/rfc4443.html#section-2.1).
* `icmp_type` - (Required) (Updatable) The value of the ICMP/IMCP_V6 message Type field as defined by [RFC 4443](https://www.rfc-editor.org/rfc/rfc4443.html#section-2.1).
* `name` - (Required) Name of the application
* `network_firewall_policy_id` - (Required) Unique Network Firewall Policy identifier
* `type` - (Required) Describes the type of application. The accepted values are - * ICMP * ICMP_V6


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `icmp_code` - The value of the ICMP/ICMP_V6 message Code (subtype) field as defined by [RFC 4443](https://www.rfc-editor.org/rfc/rfc4443.html#section-2.1).
* `icmp_type` - The value of the ICMP/ICMP_V6 message Type field as defined by [RFC 4443](https://www.rfc-editor.org/rfc/rfc4443.html#section-2.1).
* `name` - Name of the application.
* `parent_resource_id` - OCID of the Network Firewall Policy this application belongs to.
* `type` - Describes the type of application.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Network Firewall Policy Application
	* `update` - (Defaults to 20 minutes), when updating the Network Firewall Policy Application
	* `delete` - (Defaults to 20 minutes), when destroying the Network Firewall Policy Application


## Import

NetworkFirewallPolicyApplications can be imported using the `name`, e.g.

```
$ terraform import oci_network_firewall_network_firewall_policy_application.test_network_firewall_policy_application "networkFirewallPolicies/{networkFirewallPolicyId}/applications/{applicationName}" 
```

