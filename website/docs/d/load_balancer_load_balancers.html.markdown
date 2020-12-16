---
subcategory: "Load Balancer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_load_balancer_load_balancers"
sidebar_current: "docs-oci-datasource-load_balancer-load_balancers"
description: |-
  Provides the list of Load Balancers in Oracle Cloud Infrastructure Load Balancer service
---

# Data Source: oci_load_balancer_load_balancers
This data source provides the list of Load Balancers in Oracle Cloud Infrastructure Load Balancer service.

Lists all load balancers in the specified compartment.

## Supported Aliases

* `oci_load_balancers`

## Example Usage

```hcl
data "oci_load_balancer_load_balancers" "test_load_balancers" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	detail = var.load_balancer_detail
	display_name = var.load_balancer_display_name
	state = var.load_balancer_state
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
	* `reserved_ip` - Pre-created public IP that will be used as the IP of this load balancer. This reserved IP will not be deleted when load balancer is deleted. This ip should not be already mapped to any other resource.
		* `id` - Ocid of the pre-created public IP. That should be attahed to this load balancer.
* `ip_addresses` - An array of IP addresses. Deprecated: use ip_address_details instead.
* `is_private` - Whether the load balancer has a VCN-local (private) IP address.

	If "true", the service assigns a private IP address to the load balancer.

	If "false", the service assigns a public IP address to the load balancer.

	A public load balancer is accessible from the internet, depending on your VCN's [security list rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securitylists.htm). For more information about public and private load balancers, see [How Load Balancing Works](https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm#how-load-balancing-works).

	Example: `true` 
* `network_security_group_ids` - An array of NSG [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) associated with the load balancer.

	During the load balancer's creation, the service adds the new load balancer to the specified NSGs.

	The benefits of associating the load balancer with NSGs include:
	*  NSGs define network security rules to govern ingress and egress traffic for the load balancer.
	*  The network security rules of other resources can reference the NSGs associated with the load balancer to ensure access.

	Example: ["ocid1.nsg.oc1.phx.unique_ID"] 
* `shape` - A template that determines the total pre-provisioned bandwidth (ingress plus egress). To get a list of available shapes, use the [ListShapes](https://docs.cloud.oracle.com/iaas/api/#/en/loadbalancer/20170115/LoadBalancerShape/ListShapes) operation.  Example: `100Mbps` 
* `shape_details` - The configuration details to update load balancer to a different shape. 
	* `maximum_bandwidth_in_mbps` - Bandwidth in Mbps that determines the maximum bandwidth (ingress plus egress) that the load balancer can achieve. This bandwidth cannot always guaranteed. For a guaranteed bandwidth use the minimumBandwidthInMbps parameter.

		The values must be between minimumBandwidthInMbps and the highest limit available in multiples of 10. The highest limit available is defined in [Service Limits](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/servicelimits.htm).

		Example: `1500` 
	* `minimum_bandwidth_in_mbps` - Bandwidth in Mbps that determines the total pre-provisioned bandwidth (ingress plus egress). The values must be between 0 and the maximumBandwidthInMbps in multiples of 10. The current allowed maximum value is defined in [Service Limits](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/servicelimits.htm).  Example: `150` 
* `state` - The current state of the load balancer. 
* `subnet_ids` - An array of subnet [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the load balancer was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 

