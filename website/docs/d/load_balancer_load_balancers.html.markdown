---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_load_balancer_load_balancers"
sidebar_current: "docs-oci-datasource-load_balancer-load_balancers"
description: |-
  Provides the list of Load Balancers in Oracle Cloud Infrastructure Load Balancer service
---

# Data Source: oci_load_balancer_load_balancers
Other supported legacy names/aliases:
  * `oci_load_balancer`

This data source provides the list of Load Balancers in Oracle Cloud Infrastructure Load Balancer service.

Lists all load balancers in the specified compartment.

## Example Usage

```hcl
data "oci_load_balancer_load_balancers" "test_load_balancers" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	detail = "${var.load_balancer_detail}"
	display_name = "${var.load_balancer_display_name}"
	state = "${var.load_balancer_state}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the load balancers to list.
* `detail` - (Optional) The level of detail to return for each result. Can be `full` or `simple`.  Example: `full` 
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly.  Example: `example_load_balancer` 
* `state` - (Optional) A filter to return only resources that match the given lifecycle state.  Example: `SUCCEEDED` 


## Attributes Reference

The following attributes are exported:

* `load_balancers` - The list of load_balancers.

### LoadBalancer Reference

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
* `ip_addresses` - An array of IP addresses. Deprecated: use ip_address_details instead.
* `is_private` - Whether the load balancer has a VCN-local (private) IP address.

	If "true", the service assigns a private IP address to the load balancer. The load balancer requires only one subnet to host both the primary and secondary load balancers. The private IP address is local to the subnet. The load balancer is accessible only from within the VCN that contains the associated subnet, or as further restricted by your security list rules. The load balancer can route traffic to any backend server that is reachable from the VCN.

	For a private load balancer, both the primary and secondary load balancer hosts are within the same Availability Domain.

	If "false", the service assigns a public IP address to the load balancer. A load balancer with a public IP address requires two subnets, each in a different Availability Domain. One subnet hosts the primary load balancer and the other hosts the secondary (standby) load balancer. A public load balancer is accessible from the internet, depending on your VCN's [security list rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securitylists.htm).

	Example: `true` 
* `shape` - A template that determines the total pre-provisioned bandwidth (ingress plus egress). To get a list of available shapes, use the [ListShapes](https://docs.cloud.oracle.com/iaas/api/#/en/loadbalancer/20170115/LoadBalancerShape/ListShapes) operation.  Example: `100Mbps` 
* `state` - The current state of the load balancer. 
* `subnet_ids` - An array of subnet [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `time_created` - The date and time the load balancer was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 

