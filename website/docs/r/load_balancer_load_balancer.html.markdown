---
layout: "oci"
page_title: "OCI: oci_load_balancer_load_balancer"
sidebar_current: "docs-oci-resource-load_balancer-load_balancer"
description: |-
  Creates and manages an OCI LoadBalancer
---

# oci_load_balancer_load_balancer
`oci_load_balancer` is the old name for `oci_load_balancer_load_balancer`. Both names are supported but `oci_load_balancer_load_balancer` is used in the docs.
The `oci_load_balancer_load_balancer` resource creates and manages an OCI LoadBalancer

Creates a new load balancer in the specified compartment. For general information about load balancers,
see [Overview of the Load Balancing Service](https://docs.us-phoenix-1.oraclecloud.com/Content/Balance/Concepts/balanceoverview.htm).

For the purposes of access control, you must provide the OCID of the compartment where you want
the load balancer to reside. Notice that the load balancer doesn't have to be in the same compartment as the VCN
or backend set. If you're not sure which compartment to use, put the load balancer in the same compartment as the VCN.
For information about access control and compartments, see
[Overview of the IAM Service](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/overview.htm).

You must specify a display name for the load balancer. It does not have to be unique, and you can change it.

For information about Availability Domains, see
[Regions and Availability Domains](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/regions.htm).
To get a list of Availability Domains, use the `ListAvailabilityDomains` operation
in the Identity and Access Management Service API.

All Oracle Cloud Infrastructure resources, including load balancers, get an Oracle-assigned,
unique ID called an Oracle Cloud Identifier (OCID). When you create a resource, you can find its OCID
in the response. You can also retrieve a resource's OCID by using a List API operation on that resource type,
or by viewing the resource in the Console. Fore more information, see
[Resource Identifiers](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).

After you send your request, the new object's state will temporarily be PROVISIONING. Before using the
object, first make sure its state has changed to RUNNING.

When you create a load balancer, the system assigns an IP address.
To get the IP address, use the [GetLoadBalancer](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/LoadBalancer/GetLoadBalancer) operation.


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
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the compartment in which to create the load balancer.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Required) (Updatable) A user-friendly name. It does not have to be unique, and it is changeable. Avoid entering confidential information.  Example: `example_load_balancer` 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `is_private` - (Optional) Whether the load balancer has a VCN-local (private) IP address.  If "true", the service assigns a private IP address to the load balancer. The load balancer requires only one subnet to host both the primary and secondary load balancers. The private IP address is local to the subnet. The load balancer is accessible only from within the VCN that contains the associated subnet, or as further restricted by your security list rules. The load balancer can route traffic to any backend server that is reachable from the VCN.  For a private load balancer, both the primary and secondary load balancer hosts are within the same Availability Domain.  If "false", the service assigns a public IP address to the load balancer. A load balancer with a public IP address requires two subnets, each in a different Availability Domain. One subnet hosts the primary load balancer and the other hosts the secondary (standby) load balancer. A public load balancer is accessible from the internet, depending on your VCN's [security list rules](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/securitylists.htm).  Example: `true` 
* `shape` - (Required) A template that determines the total pre-provisioned bandwidth (ingress plus egress). To get a list of available shapes, use the [ListShapes](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/LoadBalancerShape/ListShapes) operation.  Example: `100Mbps` 
* `subnet_ids` - (Required) An array of subnet [OCIDs](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the compartment containing the load balancer.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. It does not have to be unique, and it is changeable.  Example: `example_load_balancer` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the load balancer.
* `ip_addresses` - An array of IP addresses.
* `is_private` - Whether the load balancer has a VCN-local (private) IP address.  If "true", the service assigns a private IP address to the load balancer. The load balancer requires only one subnet to host both the primary and secondary load balancers. The private IP address is local to the subnet. The load balancer is accessible only from within the VCN that contains the associated subnet, or as further restricted by your security list rules. The load balancer can route traffic to any backend server that is reachable from the VCN.  For a private load balancer, both the primary and secondary load balancer hosts are within the same Availability Domain.  If "false", the service assigns a public IP address to the load balancer. A load balancer with a public IP address requires two subnets, each in a different Availability Domain. One subnet hosts the primary load balancer and the other hosts the secondary (standby) load balancer. A public load balancer is accessible from the internet, depending on your VCN's [security list rules](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/securitylists.htm).  Example: `true` 
* `shape` - A template that determines the total pre-provisioned bandwidth (ingress plus egress). To get a list of available shapes, use the [ListShapes](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/LoadBalancerShape/ListShapes) operation.  Example: `100Mbps` 
* `state` - The current state of the load balancer. 
* `subnet_ids` - An array of subnet [OCIDs](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).
* `time_created` - The date and time the load balancer was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 

## Import

LoadBalancers can be imported using the `id`, e.g.

```
$ terraform import oci_load_balancer_load_balancer.test_load_balancer "id"
```
