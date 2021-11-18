---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_virtual_circuit"
sidebar_current: "docs-oci-resource-core-virtual_circuit"
description: |-
  Provides the Virtual Circuit resource in Oracle Cloud Infrastructure Core service
---

# oci_core_virtual_circuit
This resource provides the Virtual Circuit resource in Oracle Cloud Infrastructure Core service.

Creates a new virtual circuit to use with Oracle Cloud
Infrastructure FastConnect. For more information, see
[FastConnect Overview](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/fastconnect.htm).

For the purposes of access control, you must provide the OCID of the
compartment where you want the virtual circuit to reside. If you're
not sure which compartment to use, put the virtual circuit in the
same compartment with the DRG it's using. For more information about
compartments and access control, see
[Overview of the IAM Service](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm).
For information about OCIDs, see
[Resource Identifiers](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).

You may optionally specify a *display name* for the virtual circuit.
It does not have to be unique, and you can change it. Avoid entering confidential information.

**Important:** When creating a virtual circuit, you specify a DRG for
the traffic to flow through. Make sure you attach the DRG to your
VCN and confirm the VCN's routing sends traffic to the DRG. Otherwise
traffic will not flow. For more information, see
[Route Tables](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingroutetables.htm).


## Example Usage

```hcl
resource "oci_core_virtual_circuit" "test_virtual_circuit" {
	#Required
	compartment_id = var.compartment_id
	type = var.virtual_circuit_type

	#Optional
	bandwidth_shape_name = var.virtual_circuit_bandwidth_shape_name
	cross_connect_mappings {

		#Optional
		bgp_md5auth_key = var.virtual_circuit_cross_connect_mappings_bgp_md5auth_key
		cross_connect_or_cross_connect_group_id = oci_core_cross_connect_or_cross_connect_group.test_cross_connect_or_cross_connect_group.id
		customer_bgp_peering_ip = var.virtual_circuit_cross_connect_mappings_customer_bgp_peering_ip
		customer_bgp_peering_ipv6 = var.virtual_circuit_cross_connect_mappings_customer_bgp_peering_ipv6
		oracle_bgp_peering_ip = var.virtual_circuit_cross_connect_mappings_oracle_bgp_peering_ip
		oracle_bgp_peering_ipv6 = var.virtual_circuit_cross_connect_mappings_oracle_bgp_peering_ipv6
		vlan = var.virtual_circuit_cross_connect_mappings_vlan
	}
	customer_asn = var.virtual_circuit_customer_asn
	customer_bgp_asn = var.virtual_circuit_customer_bgp_asn
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.virtual_circuit_display_name
	freeform_tags = {"Department"= "Finance"}
	ip_mtu = var.virtual_circuit_ip_mtu
	gateway_id = oci_core_gateway.test_gateway.id
	provider_service_id = data.oci_core_fast_connect_provider_services.test_fast_connect_provider_services.fast_connect_provider_services.0.id
	provider_service_key_name = var.virtual_circuit_provider_service_key_name
	public_prefixes {
		#Required
		cidr_block = var.virtual_circuit_public_prefixes_cidr_block
	}
	region = var.virtual_circuit_region
	routing_policy = var.virtual_circuit_routing_policy
}
```

## Argument Reference

The following arguments are supported:

* `bandwidth_shape_name` - (Optional) (Updatable) The provisioned data rate of the connection. To get a list of the available bandwidth levels (that is, shapes), see [ListFastConnectProviderServiceVirtualCircuitBandwidthShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/FastConnectProviderService/ListFastConnectProviderVirtualCircuitBandwidthShapes).  Example: `10 Gbps` 
* `compartment_id` - (Required) (Updatable) The OCID of the compartment to contain the virtual circuit. 
* `cross_connect_mappings` - (Optional) (Updatable) Create a `CrossConnectMapping` for each cross-connect or cross-connect group this virtual circuit will run on. 
	* `bgp_md5auth_key` - (Optional) (Updatable) The key for BGP MD5 authentication. Only applicable if your system requires MD5 authentication. If empty or not set (null), that means you don't use BGP MD5 authentication. 
	* `cross_connect_or_cross_connect_group_id` - (Optional) (Updatable) The OCID of the cross-connect or cross-connect group for this mapping. Specified by the owner of the cross-connect or cross-connect group (the customer if the customer is colocated with Oracle, or the provider if the customer is connecting via provider). 
	* `customer_bgp_peering_ip` - (Optional) (Updatable) The BGP IPv4 address for the router on the other end of the BGP session from Oracle. Specified by the owner of that router. If the session goes from Oracle to a customer, this is the BGP IPv4 address of the customer's edge router. If the session goes from Oracle to a provider, this is the BGP IPv4 address of the provider's edge router. Must use a /30 or /31 subnet mask.

		There's one exception: for a public virtual circuit, Oracle specifies the BGP IPv4 addresses.

		Example: `10.0.0.18/31` 
	* `customer_bgp_peering_ipv6` - (Optional) (Updatable) IPv6 is currently supported only in the Government Cloud. The BGP IPv6 address for the router on the other end of the BGP session from Oracle. Specified by the owner of that router. If the session goes from Oracle to a customer, this is the BGP IPv6 address of the customer's edge router. If the session goes from Oracle to a provider, this is the BGP IPv6 address of the provider's edge router. Only subnet masks from /64 up to /127 are allowed.

		There's one exception: for a public virtual circuit, Oracle specifies the BGP IPv6 addresses.

		IPv6 addressing is supported for all commercial and government regions. See [IPv6 Addresses](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/ipv6.htm).

		Example: `2001:db8::1/64` 
	* `oracle_bgp_peering_ip` - (Optional) (Updatable) The IPv4 address for Oracle's end of the BGP session. Must use a /30 or /31 subnet mask. If the session goes from Oracle to a customer's edge router, the customer specifies this information. If the session goes from Oracle to a provider's edge router, the provider specifies this.

		There's one exception: for a public virtual circuit, Oracle specifies the BGP IPv4 addresses.

		Example: `10.0.0.19/31` 
	* `oracle_bgp_peering_ipv6` - (Optional) (Updatable) IPv6 is currently supported only in the Government Cloud. The IPv6 address for Oracle's end of the BGP session.  Only subnet masks from /64 up to /127 are allowed. If the session goes from Oracle to a customer's edge router, the customer specifies this information. If the session goes from Oracle to a provider's edge router, the provider specifies this.

		There's one exception: for a public virtual circuit, Oracle specifies the BGP IPv6 addresses.

		Note that IPv6 addressing is currently supported only in certain regions. See [IPv6 Addresses](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/ipv6.htm).

		Example: `2001:db8::2/64` 
	* `vlan` - (Optional) (Updatable) The number of the specific VLAN (on the cross-connect or cross-connect group) that is assigned to this virtual circuit. Specified by the owner of the cross-connect or cross-connect group (the customer if the customer is colocated with Oracle, or the provider if the customer is connecting via provider).  Example: `200` 
* `customer_asn` - (Optional) (Updatable) Your BGP ASN (either public or private). Provide this value only if there's a BGP session that goes from your edge router to Oracle. Otherwise, leave this empty or null. Can be a 2-byte or 4-byte ASN. Uses "asplain" format.  Example: `12345` (2-byte) or `1587232876` (4-byte) 
* `customer_bgp_asn` - (Optional) (Updatable) Deprecated. Instead use `customerAsn`. If you specify values for both, the request will be rejected. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `gateway_id` - (Optional) (Updatable) For private virtual circuits only. The OCID of the [dynamic routing gateway (DRG)](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Drg) that this virtual circuit uses. 
* `ip_mtu` - (Optional) (Updatable) The layer-3 IP MTU to be used for this VirtualCircuit
* `provider_service_id` - (Optional) The OCID of the service offered by the provider (if you're connecting via a provider). To get a list of the available service offerings, see [ListFastConnectProviderServices](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/FastConnectProviderService/ListFastConnectProviderServices). 
* `provider_service_key_name` - (Optional) (Updatable) The service key name offered by the provider (if the customer is connecting via a provider). 
* `public_prefixes` - (Optional) (Updatable) For a public virtual circuit. The public IP prefixes (CIDRs) the customer wants to advertise across the connection. 
	* `cidr_block` - (Required) (Updatable) An individual public IP prefix (CIDR) to add to the public virtual circuit. All prefix sizes are allowed. 
* `region` - (Optional) The Oracle Cloud Infrastructure region where this virtual circuit is located. Example: `phx` 
* `routing_policy` - (Optional) (Updatable) The routing policy sets how routing information about the Oracle cloud is shared over a public virtual circuit. Policies available are: `ORACLE_SERVICE_NETWORK`, `REGIONAL`, `MARKET_LEVEL`, and `GLOBAL`. See [Route Filtering](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/routingonprem.htm#route_filtering) for details. By default, routing information is shared for all routes in the same market. 
* `type` - (Required) The type of IP addresses used in this virtual circuit. PRIVATE means [RFC 1918](https://tools.ietf.org/html/rfc1918) addresses (10.0.0.0/8, 172.16/12, and 192.168/16). 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `bandwidth_shape_name` - The provisioned data rate of the connection. To get a list of the available bandwidth levels (that is, shapes), see [ListFastConnectProviderServiceVirtualCircuitBandwidthShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/FastConnectProviderService/ListFastConnectProviderVirtualCircuitBandwidthShapes).  Example: `10 Gbps` 
* `bgp_ipv6session_state` - The state of the Ipv6 BGP session associated with the virtual circuit.
* `bgp_management` - Deprecated. Instead use the information in [FastConnectProviderService](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/FastConnectProviderService/). 
* `bgp_session_state` - The state of the Ipv4 BGP session associated with the virtual circuit.
* `compartment_id` - The OCID of the compartment containing the virtual circuit.
* `cross_connect_mappings` - An array of mappings, each containing properties for a cross-connect or cross-connect group that is associated with this virtual circuit. 
	* `bgp_md5auth_key` - The key for BGP MD5 authentication. Only applicable if your system requires MD5 authentication. If empty or not set (null), that means you don't use BGP MD5 authentication. 
	* `cross_connect_or_cross_connect_group_id` - The OCID of the cross-connect or cross-connect group for this mapping. Specified by the owner of the cross-connect or cross-connect group (the customer if the customer is colocated with Oracle, or the provider if the customer is connecting via provider). 
	* `customer_bgp_peering_ip` - The BGP IPv4 address for the router on the other end of the BGP session from Oracle. Specified by the owner of that router. If the session goes from Oracle to a customer, this is the BGP IPv4 address of the customer's edge router. If the session goes from Oracle to a provider, this is the BGP IPv4 address of the provider's edge router. Must use a /30 or /31 subnet mask.

		There's one exception: for a public virtual circuit, Oracle specifies the BGP IPv4 addresses.

		Example: `10.0.0.18/31` 
	* `customer_bgp_peering_ipv6` - The BGP IPv6 address for the router on the other end of the BGP session from Oracle. Specified by the owner of that router. If the session goes from Oracle to a customer, this is the BGP IPv6 address of the customer's edge router. If the session goes from Oracle to a provider, this is the BGP IPv6 address of the provider's edge router. Only subnet masks from /64 up to /127 are allowed.

		There's one exception: for a public virtual circuit, Oracle specifies the BGP IPv6 addresses.

		IPv6 addressing is supported for all commercial and government regions. See [IPv6 Addresses](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/ipv6.htm).

		Example: `2001:db8::1/64` 
	* `oracle_bgp_peering_ip` - The IPv4 address for Oracle's end of the BGP session. Must use a /30 or /31 subnet mask. If the session goes from Oracle to a customer's edge router, the customer specifies this information. If the session goes from Oracle to a provider's edge router, the provider specifies this.

		There's one exception: for a public virtual circuit, Oracle specifies the BGP IPv4 addresses.

		Example: `10.0.0.19/31` 
	* `oracle_bgp_peering_ipv6` - The IPv6 address for Oracle's end of the BGP session. Only subnet masks from /64 up to /127 are allowed. If the session goes from Oracle to a customer's edge router, the customer specifies this information. If the session goes from Oracle to a provider's edge router, the provider specifies this.

		There's one exception: for a public virtual circuit, Oracle specifies the BGP IPv6 addresses.

		Note that IPv6 addressing is currently supported only in certain regions. See [IPv6 Addresses](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/ipv6.htm).

		Example: `2001:db8::2/64` 
	* `vlan` - The number of the specific VLAN (on the cross-connect or cross-connect group) that is assigned to this virtual circuit. Specified by the owner of the cross-connect or cross-connect group (the customer if the customer is colocated with Oracle, or the provider if the customer is connecting via provider).  Example: `200` 
* `customer_asn` - The BGP ASN of the network at the other end of the BGP session from Oracle. If the session is between the customer's edge router and Oracle, the value is the customer's ASN. If the BGP session is between the provider's edge router and Oracle, the value is the provider's ASN. Can be a 2-byte or 4-byte ASN. Uses "asplain" format. 
* `customer_bgp_asn` - Deprecated. Instead use `customerAsn`. If you specify values for both, the request will be rejected. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `gateway_id` - The OCID of the customer's [dynamic routing gateway (DRG)](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Drg) that this virtual circuit uses. Applicable only to private virtual circuits. 
* `id` - The virtual circuit's Oracle ID (OCID).
* `ip_mtu` - The layer-3 IP MTU to be used for this VirtualCircuit
* `oracle_bgp_asn` - The Oracle BGP ASN.
* `provider_service_id` - The OCID of the service offered by the provider (if the customer is connecting via a provider). 
* `provider_service_key_name` - The service key name offered by the provider (if the customer is connecting via a provider). 
* `provider_state` - The provider's state in relation to this virtual circuit (if the customer is connecting via a provider). ACTIVE means the provider has provisioned the virtual circuit from their end. INACTIVE means the provider has not yet provisioned the virtual circuit, or has de-provisioned it. 
* `public_prefixes` - For a public virtual circuit. The public IP prefixes (CIDRs) the customer wants to advertise across the connection. All prefix sizes are allowed. 
* `reference_comment` - Provider-supplied reference information about this virtual circuit (if the customer is connecting via a provider). 
* `region` - The Oracle Cloud Infrastructure region where this virtual circuit is located. 
* `routing_policy` - The routing policy sets how routing information about the Oracle cloud is shared over a public virtual circuit. Policies available are: `ORACLE_SERVICE_NETWORK`, `REGIONAL`, `MARKET_LEVEL`, and `GLOBAL`. See [Route Filtering](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/routingonprem.htm#route_filtering) for details. By default, routing information is shared for all routes in the same market. 
* `service_type` - Provider service type. 
* `state` - The virtual circuit's current state. For information about the different states, see [FastConnect Overview](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/fastconnect.htm). 
* `time_created` - The date and time the virtual circuit was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `type` - Whether the virtual circuit supports private or public peering. For more information, see [FastConnect Overview](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/fastconnect.htm). 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Virtual Circuit
	* `update` - (Defaults to 20 minutes), when updating the Virtual Circuit
	* `delete` - (Defaults to 20 minutes), when destroying the Virtual Circuit


## Import

VirtualCircuits can be imported using the `id`, e.g.

```
$ terraform import oci_core_virtual_circuit.test_virtual_circuit "id"
```

