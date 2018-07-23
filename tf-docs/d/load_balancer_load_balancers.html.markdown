---
layout: "oci"
page_title: "OCI: oci_load_balancer_load_balancers"
sidebar_current: "docs-oci-datasource-load_balancer-load_balancers"
description: |-
Provides a list of LoadBalancers
---
# Data Source: oci_load_balancer_load_balancers
`oci_load_balancers` is the old name for `oci_load_balancer_load_balancers`. Both names are supported but `oci_load_balancer_load_balancers` is used in the docs.
The LoadBalancers data source allows access to the list of OCI load_balancers

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

* `compartment_id` - (Required) The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the compartment containing the load balancers to list.
* `detail` - (Optional) The level of detail to return for each result. Can be `full` or `simple`.  Example: `full` 
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly.  Example: `example_load_balancer` 
* `state` - (Optional) A filter to return only resources that match the given lifecycle state.  Example: `SUCCEEDED` 


## Attributes Reference

The following attributes are exported:

* `load_balancers` - The list of load_balancers.

### LoadBalancer Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the compartment containing the load balancer.
* `display_name` - A user-friendly name. It does not have to be unique, and it is changeable.  Example: `example_load_balancer` 
* `id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the load balancer.
* `ip_addresses` - An array of IP addresses. 
	* `ip_address` - An IP address.  Example: `192.168.0.3` 
	* `is_public` - Whether the IP address is public or private.  If "true", the IP address is public and accessible from the internet.  If "false", the IP address is private and accessible only from within the associated VCN. 
* `is_private` - Whether the load balancer has a VCN-local (private) IP address.  If "true", the service assigns a private IP address to the load balancer. The load balancer requires only one subnet to host both the primary and secondary load balancers. The private IP address is local to the subnet. The load balancer is accessible only from within the VCN that contains the associated subnet, or as further restricted by your security list rules. The load balancer can route traffic to any backend server that is reachable from the VCN.  For a private load balancer, both the primary and secondary load balancer hosts are within the same Availability Domain.  If "false", the service assigns a public IP address to the load balancer. A load balancer with a public IP address requires two subnets, each in a different Availability Domain. One subnet hosts the primary load balancer and the other hosts the secondary (standby) load balancer. A public load balancer is accessible from the internet, depending on your VCN's [security list rules](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/securitylists.htm).  Example: `true` 
* `shape` - A template that determines the total pre-provisioned bandwidth (ingress plus egress). To get a list of available shapes, use the [ListShapes](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/LoadBalancerShape/ListShapes) operation.  Example: `100Mbps` 
* `state` - The current state of the load balancer. 
* `subnet_ids` - An array of subnet [OCIDs](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).
* `time_created` - The date and time the load balancer was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 

