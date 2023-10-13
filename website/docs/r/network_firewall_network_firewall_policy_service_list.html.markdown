---
subcategory: "Network Firewall"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_firewall_network_firewall_policy_service_list"
sidebar_current: "docs-oci-resource-network_firewall-network_firewall_policy_service_list"
description: |-
  Provides the Network Firewall Policy Service List resource in Oracle Cloud Infrastructure Network Firewall service
---

# oci_network_firewall_network_firewall_policy_service_list
This resource provides the Network Firewall Policy Service List resource in Oracle Cloud Infrastructure Network Firewall service.

Creates a new ServiceList for the Network Firewall Policy.


## Example Usage

```hcl
resource "oci_network_firewall_network_firewall_policy_service_list" "test_network_firewall_policy_service_list" {
	#Required
	name = var.network_firewall_policy_service_list_name
	network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
	services = var.network_firewall_policy_service_list_services
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the service Group.
* `network_firewall_policy_id` - (Required) Unique Network Firewall Policy identifier
* `services` - (Required) (Updatable) Collection of service names. The services referenced in the service list must already be present in the policy before being used in the service list. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `name` - Name of the service Group.
* `parent_resource_id` - OCID of the Network Firewall Policy this serviceList belongs to.
* `services` - List of services in the group.
* `total_services` - Count of total services in the given service List.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Network Firewall Policy Service List
	* `update` - (Defaults to 20 minutes), when updating the Network Firewall Policy Service List
	* `delete` - (Defaults to 20 minutes), when destroying the Network Firewall Policy Service List


## Import

NetworkFirewallPolicyServiceLists can be imported using the `id`, e.g.

```
$ terraform import oci_network_firewall_network_firewall_policy_service_list.test_network_firewall_policy_service_list "networkFirewallPolicies/{networkFirewallPolicyId}/serviceLists/{serviceListName}" 
```

