---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_virtual_circuit"
sidebar_current: "docs-oci-datasource-core-virtual_circuit"
description: |-
  Provides details about a specific Virtual Circuit in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_virtual_circuit
This data source provides details about a specific Virtual Circuit resource in Oracle Cloud Infrastructure Core service.

Gets the specified virtual circuit's information.

## Example Usage

```hcl
data "oci_core_virtual_circuit" "test_virtual_circuit" {
	#Required
	virtual_circuit_id = oci_core_virtual_circuit.test_virtual_circuit.id
}
```

## Argument Reference

The following arguments are supported:

* `virtual_circuit_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the virtual circuit.


## Attributes Reference

The following attributes are exported:

* `bandwidth_shape_name` - The provisioned data rate of the connection. To get a list of the available bandwidth levels (that is, shapes), see [ListFastConnectProviderServiceVirtualCircuitBandwidthShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/FastConnectProviderService/ListFastConnectProviderVirtualCircuitBandwidthShapes).  Example: `10 Gbps` 
* `bgp_admin_state` - Set to `ENABLED` (the default) to activate the BGP session of the virtual circuit, set to `DISABLED` to deactivate the virtual circuit. 
* `bgp_ipv6session_state` - The state of the Ipv6 BGP session associated with the virtual circuit.
* `bgp_management` - Deprecated. Instead use the information in [FastConnectProviderService](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/FastConnectProviderService/). 
* `bgp_session_state` - The state of the Ipv4 BGP session associated with the virtual circuit.
* `compartment_id` - The OCID of the compartment containing the virtual circuit.
* `cross_connect_mappings` - An array of mappings, each containing properties for a cross-connect or cross-connect group that is associated with this virtual circuit.
	* `bgp_md5auth_key` - The key for BGP MD5 authentication. Only applicable if your system requires MD5 authentication. If empty or not set (null), that means you don't use BGP MD5 authentication.
	* `cross_connect_or_cross_connect_group_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cross-connect or cross-connect group for this mapping. Specified by the owner of the cross-connect or cross-connect group (the customer if the customer is colocated with Oracle, or the provider if the customer is connecting via provider). 
	* `customer_bgp_peering_ip` - The BGP IPv4 address for the router on the other end of the BGP session from Oracle. Specified by the owner of that router. If the session goes from Oracle to a customer, this is the BGP IPv4 address of the customer's edge router. If the session goes from Oracle to a provider, this is the BGP IPv4 address of the provider's edge router. Must use a subnet mask from /28 to /31.

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
* `gateway_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the customer's [dynamic routing gateway (DRG)](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Drg) that this virtual circuit uses. Applicable only to private virtual circuits. 
* `id` - The virtual circuit's Oracle ID ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)).
* `ip_mtu` - The layer 3 IP MTU to use on this virtual circuit.
* `is_bfd_enabled` - Set to `true` to enable BFD for IPv4 BGP peering, or set to `false` to disable BFD. If this is not set, the default is `false`. 
* `is_transport_mode` - Set to `true` for the virtual circuit to carry only encrypted traffic, or set to `false` for the virtual circuit to carry unencrypted traffic. If this is not set, the default is `false`. 
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
* `virtual_circuit_redundancy_metadata` - Redundancy level details of the virtual circuit 
	* `configured_redundancy_level` - The configured redundancy level of the virtual circuit
	* `ipv4bgp_session_redundancy_status` - IPV4 BGP redundancy status indicates if the configured redundancy level is met
	* `ipv6bgp_session_redundancy_status` - IPV6 BGP redundancy status indicates if the configured redundancy level is met

