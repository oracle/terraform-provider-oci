---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_ipsec_connection_tunnels"
sidebar_current: "docs-oci-datasource-core-ipsec_connection_tunnels"
description: |-
  Provides the list of Ip Sec Connection Tunnels in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_ipsec_connection_tunnels
This data source provides the list of Ip Sec Connection Tunnels in Oracle Cloud Infrastructure Core service.

Lists the tunnel information for the specified IPSec connection.


## Example Usage

```hcl
data "oci_core_ipsec_connection_tunnels" "test_ip_sec_connection_tunnels" {
	#Required
	ipsec_id = oci_core_ipsec.test_ipsec.id
}
```

## Argument Reference

The following arguments are supported:

* `ipsec_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IPSec connection.


## Attributes Reference

The following attributes are exported:

* `ip_sec_connection_tunnels` - The list of two ip_sec_connection_tunnels.

### IpSecConnectionTunnel Reference

The following attributes are exported:

* `associated_virtual_circuits` - The list of virtual circuit [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)s over which your network can reach this tunnel. 
* `bgp_session_info` - Information for establishing a BGP session for the IPSec tunnel.
	* `bgp_ipv6_state` - The state of the BGP IPv6 session. 
	* `bgp_state` - The state of the BGP session. 
	* `customer_bgp_asn` - If the tunnel's `routing` attribute is set to `BGP` (see [IPSecConnectionTunnel](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/IPSecConnectionTunnel/)), this ASN is required and used for the tunnel's BGP session. This is the ASN of the network on the CPE end of the BGP session. Can be a 2-byte or 4-byte ASN. Uses "asplain" format.

		If the tunnel uses static routing, the `customerBgpAsn` must be null.

		Example: `12345` (2-byte) or `1587232876` (4-byte) 
	* `customer_interface_ip` - The IP address for the CPE end of the inside tunnel interface.

		If the tunnel's `routing` attribute is set to `BGP` (see [IPSecConnectionTunnel](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/IPSecConnectionTunnel/)), this IP address is required and used for the tunnel's BGP session.

		If `routing` is instead set to `STATIC`, this IP address is optional. You can set this IP address so you can troubleshoot or monitor the tunnel.

		The value must be a /30 or /31.

		Example: `10.0.0.5/31` 
	* `customer_interface_ipv6` - The IPv6 address for the CPE end of the inside tunnel interface. This IP address is optional.

		If the tunnel's `routing` attribute is set to `BGP` (see [IPSecConnectionTunnel](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/IPSecConnectionTunnel/)), this IP address is used for the tunnel's BGP session.

		If `routing` is instead set to `STATIC`, you can set this IP address to troubleshoot or monitor the tunnel.

		Only subnet masks from /64 up to /127 are allowed.

		Example: `2001:db8::1/64` 
	* `oracle_bgp_asn` - The Oracle BGP ASN. 
	* `oracle_interface_ip` - The IP address for the Oracle end of the inside tunnel interface.

		If the tunnel's `routing` attribute is set to `BGP` (see [IPSecConnectionTunnel](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/IPSecConnectionTunnel/)), this IP address is required and used for the tunnel's BGP session.

		If `routing` is instead set to `STATIC`, this IP address is optional. You can set this IP address so you can troubleshoot or monitor the tunnel.

		The value must be a /30 or /31.

		Example: `10.0.0.4/31` 
	* `oracle_interface_ipv6` - The IPv6 address for the Oracle end of the inside tunnel interface. This IP address is optional.

		If the tunnel's `routing` attribute is set to `BGP` (see [IPSecConnectionTunnel](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/IPSecConnectionTunnel/)), this IP address is used for the tunnel's BGP session.

		If `routing` is instead set to `STATIC`, you can set this IP address to troubleshoot or monitor the tunnel.

		Only subnet masks from /64 up to /127 are allowed.

		Example: `2001:db8::1/64` 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the tunnel. 
* `cpe_ip` - The IP address of the CPE device's VPN headend.  Example: `203.0.113.22` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `dpd_mode` - Dead peer detection (DPD) mode set on the Oracle side of the connection. This mode sets whether Oracle can only respond to a request from the CPE device to start DPD, or both respond to and initiate requests. 
* `dpd_timeout_in_sec` - DPD timeout in seconds.
* `encryption_domain_config` - Configuration information used by the encryption domain policy.
	* `cpe_traffic_selector` - Lists IPv4 or IPv6-enabled subnets in your on-premises network.
	* `oracle_traffic_selector` - Lists IPv4 or IPv6-enabled subnets in your Oracle tenancy.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the tunnel.
* `ike_version` - Internet Key Exchange protocol version. 
* `nat_translation_enabled` - By default (the `AUTO` setting), IKE sends packets with a source and destination port set to 500, and when it detects that the port used to forward packets has changed (most likely because a NAT device is between the CPE device and the Oracle VPN headend) it will try to negotiate the use of NAT-T.

	The `ENABLED` option sets the IKE protocol to use port 4500 instead of 500 and forces encapsulating traffic with the ESP protocol inside UDP packets.

	The `DISABLED` option directs IKE to completely refuse to negotiate NAT-T even if it senses there may be a NAT device in use.

* `oracle_can_initiate` - Indicates whether Oracle can only respond to a request to start an IPSec tunnel from the CPE device, or both respond to and initiate requests. 
* `phase_one_details` - IPSec tunnel details specific to ISAKMP phase one.
	* `custom_authentication_algorithm` - The proposed custom authentication algorithm.
	* `custom_dh_group` - The proposed custom Diffie-Hellman group.
	* `custom_encryption_algorithm` - The proposed custom encryption algorithm.
	* `is_custom_phase_one_config` - Indicates whether custom phase one configuration is enabled. If this option is not enabled, default settings are proposed. 
	* `is_ike_established` - Indicates whether IKE phase one is established.
	* `lifetime` - The total configured lifetime of the IKE security association.
	* `negotiated_authentication_algorithm` - The negotiated authentication algorithm.
	* `negotiated_dh_group` - The negotiated Diffie-Hellman group.
	* `negotiated_encryption_algorithm` - The negotiated encryption algorithm.
	* `remaining_lifetime_int` - The remaining lifetime before the key is refreshed.
	* `remaining_lifetime_last_retrieved` - The date and time we retrieved the remaining lifetime, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `phase_two_details` - IPsec tunnel detail information specific to phase two.
	* `custom_authentication_algorithm` - Phase two authentication algorithm proposed during tunnel negotiation. 
	* `custom_encryption_algorithm` - The proposed custom phase two encryption algorithm. 
	* `dh_group` - The proposed Diffie-Hellman group. 
	* `is_custom_phase_two_config` - Indicates whether custom phase two configuration is enabled. If this option is not enabled, default settings are proposed. 
	* `is_esp_established` - Indicates that ESP phase two is established.
	* `is_pfs_enabled` - Indicates that PFS (perfect forward secrecy) is enabled.
	* `lifetime` - The total configured lifetime of the IKE security association.
	* `negotiated_authentication_algorithm` - The negotiated phase two authentication algorithm.
	* `negotiated_dh_group` - The negotiated Diffie-Hellman group.
	* `negotiated_encryption_algorithm` - The negotiated encryption algorithm.
	* `remaining_lifetime_int` - The remaining lifetime before the key is refreshed.
	* `remaining_lifetime_last_retrieved` - The date and time the remaining lifetime was last retrieved, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `routing` - The type of routing used for this tunnel (BGP dynamic routing, static routing, or policy-based routing). 
* `state` - The tunnel's lifecycle state.
* `status` - The status of the tunnel based on IPSec protocol characteristics.
* `time_created` - The date and time the IPSec tunnel was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_status_updated` - When the status of the IPSec tunnel last changed, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `vpn_ip` - The IP address of the Oracle VPN headend for the connection.  Example: `203.0.113.21` 

