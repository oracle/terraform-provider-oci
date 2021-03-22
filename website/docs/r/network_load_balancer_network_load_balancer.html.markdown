---
subcategory: "Network Load Balancer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_load_balancer_network_load_balancer"
sidebar_current: "docs-oci-resource-network_load_balancer-network_load_balancer"
description: |-
  Provides the Network Load Balancer resource in Oracle Cloud Infrastructure Network Load Balancer service
---

# oci_network_load_balancer_network_load_balancer
This resource provides the Network Load Balancer resource in Oracle Cloud Infrastructure Network Load Balancer service.

Creates a network load balancer.


## Example Usage

```hcl
resource "oci_network_load_balancer_network_load_balancer" "test_network_load_balancer" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.network_load_balancer_display_name
	subnet_id = oci_core_subnet.test_subnet.id
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
	is_preserve_source_destination = var.network_load_balancer_is_preserve_source_destination
	is_private = var.network_load_balancer_is_private
	network_security_group_ids = var.network_load_balancer_network_security_group_ids
	reserved_ips {

		#Optional
		id = var.network_load_balancer_reserved_ips_id
	}
}
```

## Argument Reference

The following arguments are supported:
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the network load balancer.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Required) (Updatable) Network load balancer identifier, which can be renamed.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `is_preserve_source_destination` - (Optional) (Updatable) This parameter can be enabled only if backends are compute OCIDs. When enabled, the skipSourceDestinationCheck parameter is automatically enabled on the load balancer VNIC, and packets are sent to the backend with the entire IP header intact. 
* `is_private` - (Optional) Whether the network load balancer has a virtual cloud network-local (private) IP address.

	If "true", then the service assigns a private IP address to the network load balancer.

	If "false", then the service assigns a public IP address to the network load balancer.

	A public network load balancer is accessible from the internet, depending on the [security list rules](https://docs.cloud.oracle.com/iaas/Content/network/Concepts/securitylists.htm) for your virtual cloud network. For more information about public and private network load balancers, see [How Network Load Balancing Works](https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm#how-network-load-balancing-works). This value is true by default.

	Example: `true` 
* `network_security_group_ids` - (Optional) (Updatable) An array of network security groups [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) associated with the network load balancer.

	During the creation of the network load balancer, the service adds the new load balancer to the specified network security groups.

	The benefits of associating the network load balancer with network security groups include:
	*  Network security groups define network security rules to govern ingress and egress traffic for the network load balancer.
	*  The network security rules of other resources can reference the network security groups associated with the network load balancer to ensure access.

	Example: ["ocid1.nsg.oc1.phx.unique_ID"] 
* `reserved_ips` - (Optional) An array of reserved Ips. 
	* `id` - (Optional) OCID of the reserved public IP address created with the virtual cloud network.

		Reserved public IP addresses are IP addresses that are registered using the virtual cloud network API.

		Create a reserved public IP address. When you create the network load balancer, enter the OCID of the reserved public IP address in the reservedIp field to attach the IP address to the network load balancer. This task configures the network load balancer to listen to traffic on this IP address.

		Reserved public IP addresses are not deleted when the network load balancer is deleted. The IP addresses become unattached from the network load balancer.

		Example: "ocid1.publicip.oc1.phx.unique_ID" 
* `subnet_id` - (Required) The subnet in which the network load balancer is spawned [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the network load balancer.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name, which does not have to be unique, and can be changed.  Example: `example_load_balancer` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network load balancer.
* `ip_addresses` - An array of IP addresses. 
	* `ip_address` - An IP address.  Example: `192.168.0.3` 
	* `is_public` - Whether the IP address is public or private.

		If "true", then the IP address is public and accessible from the internet.

		If "false", then the IP address is private and accessible only from within the associated virtual cloud network. 
	* `reserved_ip` - An object representing a reserved IP address to be attached or that is already attached to a network load balancer. 
		* `id` - OCID of the reserved public IP address created with the virtual cloud network.

			Reserved public IP addresses are IP addresses that are registered using the virtual cloud network API.

			Create a reserved public IP address. When you create the network load balancer, enter the OCID of the reserved public IP address in the reservedIp field to attach the IP address to the network load balancer. This task configures the network load balancer to listen to traffic on this IP address.

			Reserved public IP addresses are not deleted when the network load balancer is deleted. The IP addresses become unattached from the network load balancer.

			Example: "ocid1.publicip.oc1.phx.unique_ID" 
* `is_preserve_source_destination` - When enabled, the skipSourceDestinationCheck parameter is automatically enabled on the load balancer VNIC. Packets are sent to the backend set without any changes to the source and destination IP. 
* `is_private` - Whether the network load balancer has a virtual cloud network-local (private) IP address.

	If "true", then the service assigns a private IP address to the network load balancer.

	If "false", then the service assigns a public IP address to the network load balancer.

	A public network load balancer is accessible from the internet, depending the [security list rules](https://docs.cloud.oracle.com/iaas/Content/network/Concepts/securitylists.htm) for your virtual cloudn network. For more information about public and private network load balancers, see [How Network Load Balancing Works](https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm#how-network-load-balancing-works). This value is true by default.

	Example: `true` 
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state. 
* `network_security_group_ids` - An array of network security groups [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) associated with the network load balancer.

	During the creation of the network load balancer, the service adds the new load balancer to the specified network security groups.

	The benefits of associating the network load balancer with network security groups include:
	*  Network security groups define network security rules to govern ingress and egress traffic for the network load balancer.
	*  The network security rules of other resources can reference the network security groups associated with the network load balancer to ensure access.

	Example: ["ocid1.nsg.oc1.phx.unique_ID"] 
* `state` - The current state of the network load balancer.
* `subnet_id` - The subnet in which the network load balancer is spawned [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)."
* `system_tags` - Key-value pair representing system tags' keys and values scoped to a namespace. Example: `{"bar-key": "value"}` 
* `time_created` - The date and time the network load balancer was created, in the format defined by RFC3339.  Example: `2020-05-01T21:10:29.600Z` 
* `time_updated` - The time the network load balancer was updated. An RFC3339 formatted date-time string.  Example: `2020-05-01T22:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Network Load Balancer
	* `update` - (Defaults to 20 minutes), when updating the Network Load Balancer
	* `delete` - (Defaults to 20 minutes), when destroying the Network Load Balancer


## Import

NetworkLoadBalancers can be imported using the `id`, e.g.

```
$ terraform import oci_network_load_balancer_network_load_balancer.test_network_load_balancer "id"
```

