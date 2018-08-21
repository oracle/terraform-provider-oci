---
layout: "oci"
page_title: "OCI: oci_core_virtual_circuit"
sidebar_current: "docs-oci-resource-core-virtual_circuit"
description: |-
  Creates and manages an OCI VirtualCircuit
---

# oci_core_virtual_circuit
The `oci_core_virtual_circuit` resource creates and manages an OCI VirtualCircuit

Creates a new virtual circuit to use with Oracle Cloud
Infrastructure FastConnect. For more information, see
[FastConnect Overview](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/fastconnect.htm).

For the purposes of access control, you must provide the OCID of the
compartment where you want the virtual circuit to reside. If you're
not sure which compartment to use, put the virtual circuit in the
same compartment with the DRG it's using. For more information about
compartments and access control, see
[Overview of the IAM Service](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/overview.htm).
For information about OCIDs, see
[Resource Identifiers](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).

You may optionally specify a *display name* for the virtual circuit.
It does not have to be unique, and you can change it. Avoid entering confidential information.

**Important:** When creating a virtual circuit, you specify a DRG for
the traffic to flow through. Make sure you attach the DRG to your
VCN and confirm the VCN's routing sends traffic to the DRG. Otherwise
traffic will not flow. For more information, see
[Route Tables](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingroutetables.htm).


## Example Usage

```hcl
resource "oci_core_virtual_circuit" "test_virtual_circuit" {
	#Required
	compartment_id = "${var.compartment_id}"
	type = "${var.virtual_circuit_type}"

	#Optional
	bandwidth_shape_name = "${var.virtual_circuit_bandwidth_shape_name}"
	cross_connect_mappings {

		#Optional
		bgp_md5auth_key = "${var.virtual_circuit_cross_connect_mappings_bgp_md5auth_key}"
		cross_connect_or_cross_connect_group_id = "${oci_core_cross_connect_or_cross_connect_group.test_cross_connect_or_cross_connect_group.id}"
		customer_bgp_peering_ip = "${var.virtual_circuit_cross_connect_mappings_customer_bgp_peering_ip}"
		oracle_bgp_peering_ip = "${var.virtual_circuit_cross_connect_mappings_oracle_bgp_peering_ip}"
		vlan = "${var.virtual_circuit_cross_connect_mappings_vlan}"
	}
	customer_bgp_asn = "${var.virtual_circuit_customer_bgp_asn}"
	display_name = "${var.virtual_circuit_display_name}"
	gateway_id = "${oci_core_gateway.test_gateway.id}"
	provider_service_id = "${oci_core_provider_service.test_provider_service.id}"
	public_prefixes {
		#Required
		cidr_block = "${var.virtual_circuit_public_prefixes_cidr_block}"
	}
	region = "${var.virtual_circuit_region}"
}
```

## Argument Reference

The following arguments are supported:

* `bandwidth_shape_name` - (Optional) (Updatable) The provisioned data rate of the connection.  To get a list of the available bandwidth levels (that is, shapes), see [ListFastConnectProviderServiceVirtualCircuitBandwidthShapes](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/VirtualCircuitBandwidthShape/ListFastConnectProviderVirtualCircuitBandwidthShapes).  Example: `10 Gbps` 
* `compartment_id` - (Required) The OCID of the compartment to contain the virtual circuit. 
* `cross_connect_mappings` - (Optional) (Updatable) Create a `CrossConnectMapping` for each cross-connect or cross-connect group this virtual circuit will run on. 
	* `bgp_md5auth_key` - (Optional) (Updatable) The key for BGP MD5 authentication. Only applicable if your system requires MD5 authentication. If empty or not set (null), that means you don't use BGP MD5 authentication. 
	* `cross_connect_or_cross_connect_group_id` - (Optional) (Updatable) The OCID of the cross-connect or cross-connect group for this mapping. Specified by the owner of the cross-connect or cross-connect group (the customer if the customer is colocated with Oracle, or the provider if the customer is connecting via provider). 
	* `customer_bgp_peering_ip` - (Optional) (Updatable) The BGP IP address for the router on the other end of the BGP session from Oracle. Specified by the owner of that router. If the session goes from Oracle to a customer, this is the BGP IP address of the customer's edge router. If the session goes from Oracle to a provider, this is the BGP IP address of the provider's edge router. Must use a /30 or /31 subnet mask.

		There's one exception: for a public virtual circuit, Oracle specifies the BGP IP addresses.  Example: `10.0.0.18/31` 
	* `oracle_bgp_peering_ip` - (Optional) (Updatable) The IP address for Oracle's end of the BGP session. Must use a /30 or /31 subnet mask. If the session goes from Oracle to a customer's edge router, the customer specifies this information. If the session goes from Oracle to a provider's edge router, the provider specifies this.

		There's one exception: for a public virtual circuit, Oracle specifies the BGP IP addresses.  Example: `10.0.0.19/31` 
	* `vlan` - (Optional) (Updatable) The number of the specific VLAN (on the cross-connect or cross-connect group) that is assigned to this virtual circuit. Specified by the owner of the cross-connect or cross-connect group (the customer if the customer is colocated with Oracle, or the provider if the customer is connecting via provider).  Example: `200` 
* `customer_bgp_asn` - (Optional) (Updatable) Your BGP ASN (either public or private). Provide this value only if there's a BGP session that goes from your edge router to Oracle. Otherwise, leave this empty or null. 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `gateway_id` - (Optional) (Updatable) For private virtual circuits only. The OCID of the [Dynamic Routing Gateway (DRG)](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Drg) that this virtual circuit uses. 
* `provider_service_id` - (Optional) The OCID of the service offered by the provider (if you're connecting via a provider). To get a list of the available service offerings, see [ListFastConnectProviderServices](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/FastConnectProviderService/ListFastConnectProviderServices). 
* `public_prefixes` - (Optional) For a public virtual circuit. The public IP prefixes (CIDRs) the customer wants to advertise across the connection. 
	* `cidr_block` - (Required) An individual public IP prefix (CIDR) to add to the public virtual circuit. Must be /24 or less specific. 
* `region` - (Optional) The Oracle Cloud Infrastructure region where this virtual circuit is located. Example: `phx` 
* `type` - (Required) The type of IP addresses used in this virtual circuit. PRIVATE means [RFC 1918](https://tools.ietf.org/html/rfc1918) addresses (10.0.0.0/8, 172.16/12, and 192.168/16). Only PRIVATE is supported. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `bandwidth_shape_name` - The provisioned data rate of the connection.
* `bgp_management` - BGP management option. 
* `bgp_session_state` - The state of the BGP session associated with the virtual circuit.
* `compartment_id` - The OCID of the compartment containing the virtual circuit.
* `cross_connect_mappings` - An array of mappings, each containing properties for a cross-connect or cross-connect group that is associated with this virtual circuit. 
	* `bgp_md5auth_key` - The key for BGP MD5 authentication. Only applicable if your system requires MD5 authentication. If empty or not set (null), that means you don't use BGP MD5 authentication. 
	* `cross_connect_or_cross_connect_group_id` - The OCID of the cross-connect or cross-connect group for this mapping. Specified by the owner of the cross-connect or cross-connect group (the customer if the customer is colocated with Oracle, or the provider if the customer is connecting via provider). 
	* `customer_bgp_peering_ip` - The BGP IP address for the router on the other end of the BGP session from Oracle. Specified by the owner of that router. If the session goes from Oracle to a customer, this is the BGP IP address of the customer's edge router. If the session goes from Oracle to a provider, this is the BGP IP address of the provider's edge router. Must use a /30 or /31 subnet mask.

		There's one exception: for a public virtual circuit, Oracle specifies the BGP IP addresses.  Example: `10.0.0.18/31` 
	* `oracle_bgp_peering_ip` - The IP address for Oracle's end of the BGP session. Must use a /30 or /31 subnet mask. If the session goes from Oracle to a customer's edge router, the customer specifies this information. If the session goes from Oracle to a provider's edge router, the provider specifies this.

		There's one exception: for a public virtual circuit, Oracle specifies the BGP IP addresses.  Example: `10.0.0.19/31` 
	* `vlan` - The number of the specific VLAN (on the cross-connect or cross-connect group) that is assigned to this virtual circuit. Specified by the owner of the cross-connect or cross-connect group (the customer if the customer is colocated with Oracle, or the provider if the customer is connecting via provider).  Example: `200` 
* `customer_bgp_asn` - The BGP ASN of the network at the other end of the BGP session from Oracle. If the session is between the customer's edge router and Oracle, the value is the customer's ASN. If the BGP session is between the provider's edge router and Oracle, the value is the provider's ASN. 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `gateway_id` - The OCID of the customer's [Dynamic Routing Gateway (DRG)](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Drg) that this virtual circuit uses. Applicable only to private virtual circuits. 
* `id` - The virtual circuit's Oracle ID (OCID).
* `oracle_bgp_asn` - The Oracle BGP ASN.
* `provider_service_id` - The OCID of the service offered by the provider (if the customer is connecting via a provider). 
* `provider_state` - The provider's state in relation to this virtual circuit (if the customer is connecting via a provider). ACTIVE means the provider has provisioned the virtual circuit from their end. INACTIVE means the provider has not yet provisioned the virtual circuit, or has de-provisioned it. 
* `public_prefixes` - For a public virtual circuit. The public IP prefixes (CIDRs) the customer wants to advertise across the connection. Each prefix must be /24 or less specific. 
* `reference_comment` - Provider-supplied reference information about this virtual circuit (if the customer is connecting via a provider). 
* `region` - The Oracle Cloud Infrastructure region where this virtual circuit is located. 
* `service_type` - Provider service type. 
* `state` - The virtual circuit's current state. For information about the different states, see [FastConnect Overview](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/fastconnect.htm). 
* `time_created` - The date and time the virtual circuit was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `type` - Whether the virtual circuit supports private or public peering. For more information, see [FastConnect Overview](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/fastconnect.htm). 

## Import

VirtualCircuits can be imported using the `id`, e.g.

```
$ terraform import oci_core_virtual_circuit.test_virtual_circuit "id"
```
