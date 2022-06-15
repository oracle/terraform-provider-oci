---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_route_table"
sidebar_current: "docs-oci-resource-core-route_table"
description: |-
  Provides the Route Table resource in Oracle Cloud Infrastructure Core service
---

# oci_core_route_table
This resource provides the Route Table resource in Oracle Cloud Infrastructure Core service.

Creates a new route table for the specified VCN. In the request you must also include at least one route
rule for the new route table. For information on the number of rules you can have in a route table, see
[Service Limits](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/servicelimits.htm). For general information about route
tables in your VCN and the types of targets you can use in route rules,
see [Route Tables](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingroutetables.htm).

For the purposes of access control, you must provide the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want the route
table to reside. Notice that the route table doesn't have to be in the same compartment as the VCN, subnets,
or other Networking Service components. If you're not sure which compartment to use, put the route
table in the same compartment as the VCN. For more information about compartments and access control, see
[Overview of the IAM Service](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm). For information about OCIDs, see
[Resource Identifiers](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).

You may optionally specify a *display name* for the route table, otherwise a default is provided.
It does not have to be unique, and you can change it. Avoid entering confidential information.

For more information on configuring a VCN's default route table, see [Managing Default VCN Resources](/docs/providers/oci/guides/managing_default_resources.html)

## Example Usage

```hcl
resource "oci_core_route_table" "test_route_table" {
	#Required
	compartment_id = var.compartment_id
	vcn_id = oci_core_vcn.test_vcn.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.route_table_display_name
	freeform_tags = {"Department"= "Finance"}
	route_rules {
		#Required
		network_entity_id = oci_core_internet_gateway.test_internet_gateway.id

		#Optional
		cidr_block = var.route_table_route_rules_cidr_block
		description = var.route_table_route_rules_description
		destination = var.route_table_route_rules_destination
		destination_type = var.route_table_route_rules_destination_type
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to contain the route table.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `route_rules` - (Optional) (Updatable) The collection of rules used for routing destination IPs to network devices. 
	* `cidr_block` - (Optional) (Updatable) Deprecated. Instead use `destination` and `destinationType`. Requests that include both `cidrBlock` and `destination` will be rejected.

		A destination IP address range in CIDR notation. Matching packets will be routed to the indicated network entity (the target).

		Cannot be an IPv6 CIDR.

		Example: `0.0.0.0/0` 
	* `description` - (Optional) (Updatable) An optional description of your choice for the rule. 
	* `destination` - (Optional) (Updatable) Conceptually, this is the range of IP addresses used for matching when routing traffic. Required if you provide a `destinationType`.

		Allowed values:
		* IP address range in CIDR notation. Can be an IPv4 or IPv6 CIDR. For example: `192.168.1.0/24` or `2001:0db8:0123:45::/56`. If you set this to an IPv6 CIDR, the route rule's target can only be a DRG or internet gateway. IPv6 addressing is supported for all commercial and government regions. See [IPv6 Addresses](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/ipv6.htm).
		* The `cidrBlock` value for a [Service](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Service/), if you're setting up a route rule for traffic destined for a particular `Service` through a service gateway. For example: `oci-phx-objectstorage`. 
	* `destination_type` - (Optional) (Updatable) Type of destination for the rule. Required if you provide a `destination`.
		* `CIDR_BLOCK`: If the rule's `destination` is an IP address range in CIDR notation.
		* `SERVICE_CIDR_BLOCK`: If the rule's `destination` is the `cidrBlock` value for a [Service](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Service/) (the rule is for traffic destined for a particular `Service` through a service gateway). 
	* `network_entity_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the route rule's target. For information about the type of targets you can specify, see [Route Tables](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingroutetables.htm). 
* `vcn_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN the route table belongs to.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Route Table
	* `update` - (Defaults to 20 minutes), when updating the Route Table
	* `delete` - (Defaults to 20 minutes), when destroying the Route Table


## Import

RouteTables can be imported using the `id`, e.g.

```
$ terraform import oci_core_route_table.test_route_table "id"
```

