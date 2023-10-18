---
subcategory: "Network Firewall"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_firewall_network_firewall_policy_application_group"
sidebar_current: "docs-oci-resource-network_firewall-network_firewall_policy_application_group"
description: |-
  Provides the Network Firewall Policy Application Group resource in Oracle Cloud Infrastructure Network Firewall service
---

# oci_network_firewall_network_firewall_policy_application_group
This resource provides the Network Firewall Policy Application Group resource in Oracle Cloud Infrastructure Network Firewall service.

Creates a new ApplicationGroup inside the Network Firewall Policy.


## Example Usage

```hcl
resource "oci_network_firewall_network_firewall_policy_application_group" "test_network_firewall_policy_application_group" {
	#Required
	apps = var.network_firewall_policy_application_group_apps
	name = var.network_firewall_policy_application_group_name
	network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
}
```

## Argument Reference

The following arguments are supported:

* `apps` - (Required) (Updatable) Collection of application names. The apps referenced in the application group must already be present in the policy before being used in the application group.
* `name` - (Required) Name of the application group.
* `network_firewall_policy_id` - (Required) Unique Network Firewall Policy identifier


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `apps` - List of apps in the group.
* `name` - Name of the application group.
* `parent_resource_id` - OCID of the Network Firewall Policy this application group belongs to.
* `total_apps` - Count of total applications in the given application group.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Network Firewall Policy Application Group
	* `update` - (Defaults to 20 minutes), when updating the Network Firewall Policy Application Group
	* `delete` - (Defaults to 20 minutes), when destroying the Network Firewall Policy Application Group


## Import

NetworkFirewallPolicyApplicationGroups can be imported using the `name`, e.g.

```
$ terraform import oci_network_firewall_network_firewall_policy_application_group.test_network_firewall_policy_application_group "networkFirewallPolicies/{networkFirewallPolicyId}/applicationGroups/{applicationGroupName}" 
```

