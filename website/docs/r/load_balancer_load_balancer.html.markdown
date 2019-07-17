---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_load_balancer_load_balancer"
sidebar_current: "docs-oci-resource-load_balancer-load_balancer"
description: |-
  Provides the Load Balancer resource in Oracle Cloud Infrastructure Load Balancer service
---

# oci_load_balancer_load_balancer
This resource provides the Load Balancer resource in Oracle Cloud Infrastructure Load Balancer service.

Creates a new load balancer in the specified compartment. For general information about load balancers,
see [Overview of the Load Balancing Service](https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).

For the purposes of access control, you must provide the OCID of the compartment where you want
the load balancer to reside. Notice that the load balancer doesn't have to be in the same compartment as the VCN
or backend set. If you're not sure which compartment to use, put the load balancer in the same compartment as the VCN.
For information about access control and compartments, see
[Overview of the IAM Service](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm).

You must specify a display name for the load balancer. It does not have to be unique, and you can change it.

For information about Availability Domains, see
[Regions and Availability Domains](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm).
To get a list of Availability Domains, use the `ListAvailabilityDomains` operation
in the Identity and Access Management Service API.

All Oracle Cloud Infrastructure resources, including load balancers, get an Oracle-assigned,
unique ID called an Oracle Cloud Identifier (OCID). When you create a resource, you can find its OCID
in the response. You can also retrieve a resource's OCID by using a List API operation on that resource type,
or by viewing the resource in the Console. Fore more information, see
[Resource Identifiers](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).

When you create a load balancer, the system assigns an IP address.
To get the IP address, use the [GetLoadBalancer](https://docs.cloud.oracle.com/iaas/api/#/en/loadbalancer/20170115/LoadBalancer/GetLoadBalancer) operation.


## Supported Aliases

* `oci_load_balancer`

## Example Usage

```hcl
resource "oci_load_balancer_load_balancer" "test_load_balancer" {
	#Required
	compartment_id = "${var.compartment_id}"
	display_name = "${var.load_balancer_display_name}"
	shape = "${var.load_balancer_shape}"
	subnet_ids = "${var.load_balancer_subnet_ids}"

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
	is_private = "${var.load_balancer_is_private}"
	network_security_group_ids = "${var.load_balancer_network_security_group_ids}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to create the load balancer.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Required) (Updatable) A user-friendly name. It does not have to be unique, and it is changeable. Avoid entering confidential information.  Example: `example_load_balancer` 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `is_private` - (Optional) Whether the load balancer has a VCN-local (private) IP address.

	If "true", the service assigns a private IP address to the load balancer.

	If "false", the service assigns a public IP address to the load balancer.

	A public load balancer is accessible from the internet, depending on your VCN's [security list rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securitylists.htm). For more information about public and private load balancers, see [How Load Balancing Works](https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm#how-load-balancing-works).

	Example: `true` 
* `network_security_group_ids` - (Optional) The array of NSG [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) to be used by this Load Balancer. 
* `shape` - (Required) A template that determines the total pre-provisioned bandwidth (ingress plus egress). To get a list of available shapes, use the [ListShapes](https://docs.cloud.oracle.com/iaas/api/#/en/loadbalancer/20170115/LoadBalancerShape/ListShapes) operation.  Example: `100Mbps` 
* `subnet_ids` - (Required) An array of subnet [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the load balancer.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. It does not have to be unique, and it is changeable.  Example: `example_load_balancer` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the load balancer.
* `ip_address_details` - An array of IP addresses. 
	* `ip_address` - An IP address.  Example: `192.168.0.3` 
	* `is_public` - Whether the IP address is public or private.

		If "true", the IP address is public and accessible from the internet.

		If "false", the IP address is private and accessible only from within the associated VCN. 
* `ip_addresses` - An array of IP addresses. Deprecated: use ip_address_details instead 
* `is_private` - Whether the load balancer has a VCN-local (private) IP address.

	If "true", the service assigns a private IP address to the load balancer.

	If "false", the service assigns a public IP address to the load balancer.

	A public load balancer is accessible from the internet, depending on your VCN's [security list rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securitylists.htm). For more information about public and private load balancers, see [How Load Balancing Works](https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm#how-load-balancing-works).

	Example: `true` 
* `network_security_group_ids` - The array of NSG [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) in use by this Load Balancer. 
* `shape` - A template that determines the total pre-provisioned bandwidth (ingress plus egress). To get a list of available shapes, use the [ListShapes](https://docs.cloud.oracle.com/iaas/api/#/en/loadbalancer/20170115/LoadBalancerShape/ListShapes) operation.  Example: `100Mbps` 
* `state` - The current state of the load balancer. 
* `subnet_ids` - An array of subnet [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `time_created` - The date and time the load balancer was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 

## Import

LoadBalancers can be imported using the `id`, e.g.

```
$ terraform import oci_load_balancer_load_balancer.test_load_balancer "id"
```

