---
subcategory: "Network Firewall"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_firewall_network_firewall_policy_service"
sidebar_current: "docs-oci-resource-network_firewall-network_firewall_policy_service"
description: |-
  Provides the Network Firewall Policy Service resource in Oracle Cloud Infrastructure Network Firewall service
---

# oci_network_firewall_network_firewall_policy_service
This resource provides the Network Firewall Policy Service resource in Oracle Cloud Infrastructure Network Firewall service.

Creates a new Service for the Network Firewall Policy.


## Example Usage

```hcl
resource "oci_network_firewall_network_firewall_policy_service" "test_network_firewall_policy_service" {
	#Required
	name = var.network_firewall_policy_service_name
	network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
	port_ranges {
		#Required
		minimum_port = var.network_firewall_policy_service_port_ranges_minimum_port

		#Optional
		maximum_port = var.network_firewall_policy_service_port_ranges_maximum_port
	}
	type = var.network_firewall_policy_service_type
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the service
* `network_firewall_policy_id` - (Required) Unique Network Firewall Policy identifier
* `port_ranges` - (Required) (Updatable) List of port-ranges to be used.
	* `maximum_port` - (Optional) (Updatable) The maximum port in the range (inclusive), which may be absent for a single-port range.
	* `minimum_port` - (Required) (Updatable) The minimum port in the range (inclusive), or the sole port of a single-port range.
* `type` - (Optional) Describes the type of Service. The accepted values are 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `name` - Name of the service.
* `parent_resource_id` - OCID of the Network Firewall Policy this service belongs to.
* `port_ranges` - List of port-ranges used.
	* `maximum_port` - The maximum port in the range (inclusive), which may be absent for a single-port range.
	* `minimum_port` - The minimum port in the range (inclusive), or the sole port of a single-port range.
* `type` - Describes the type of service.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Network Firewall Policy Service
	* `update` - (Defaults to 20 minutes), when updating the Network Firewall Policy Service
	* `delete` - (Defaults to 20 minutes), when destroying the Network Firewall Policy Service


## Import

NetworkFirewallPolicyServices can be imported using the `name`, e.g.

```
$ terraform import oci_network_firewall_network_firewall_policy_service.test_network_firewall_policy_service "networkFirewallPolicies/{networkFirewallPolicyId}/services/{serviceName}" 
```

