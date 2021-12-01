---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_route_tables"
sidebar_current: "docs-oci-datasource-core-route_tables"
description: |-
  Provides the list of Route Tables in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_route_tables
This data source provides the list of Route Tables in Oracle Cloud Infrastructure Core service.

Lists the route tables in the specified VCN and specified compartment.
If the VCN ID is not provided, then the list includes the route tables from all VCNs in the specified compartment.
The response includes the default route table that automatically comes with
each VCN in the specified compartment, plus any route tables you've created.


## Example Usage

```hcl
data "oci_core_route_tables" "test_route_tables" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.route_table_display_name
	state = var.route_table_state
	vcn_id = oci_core_vcn.test_vcn.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `state` - (Optional) A filter to only return resources that match the given lifecycle state. The state value is case-insensitive. 
* `vcn_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN.


## Attributes Reference

The following attributes are exported:

* `route_tables` - The list of route_tables.

### RouteTable Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the route table.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The route table's Oracle ID ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)).
* `route_rules` - The collection of rules for routing destination IPs to network devices. 
	* `cidr_block` - Deprecated. Instead use `destination` and `destinationType`. Requests that include both `cidrBlock` and `destination` will be rejected.

		A destination IP address range in CIDR notation. Matching packets will be routed to the indicated network entity (the target).

		Cannot be an IPv6 CIDR.

		Example: `0.0.0.0/0` 
	* `description` - An optional description of your choice for the rule. 
	* `destination` - Conceptually, this is the range of IP addresses used for matching when routing traffic. Required if you provide a `destinationType`.

		Allowed values:
		* IP address range in CIDR notation. Can be an IPv4 or IPv6 CIDR. For example: `192.168.1.0/24` or `2001:0db8:0123:45::/56`. If you set this to an IPv6 CIDR, the route rule's target can only be a DRG or internet gateway. IPv6 addressing is supported for all commercial and government regions. See [IPv6 Addresses](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/ipv6.htm).
		* The `cidrBlock` value for a [Service](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Service/), if you're setting up a route rule for traffic destined for a particular `Service` through a service gateway. For example: `oci-phx-objectstorage`. 
	* `destination_type` - Type of destination for the rule. Required if you provide a `destination`.
		* `CIDR_BLOCK`: If the rule's `destination` is an IP address range in CIDR notation.
		* `SERVICE_CIDR_BLOCK`: If the rule's `destination` is the `cidrBlock` value for a [Service](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Service/) (the rule is for traffic destined for a particular `Service` through a service gateway). 
	* `network_entity_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the route rule's target. For information about the type of targets you can specify, see [Route Tables](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingroutetables.htm). 
* `state` - The route table's current state.
* `time_created` - The date and time the route table was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `vcn_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN the route table list belongs to.

