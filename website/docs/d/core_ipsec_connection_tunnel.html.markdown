---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_ipsec_connection_tunnel"
sidebar_current: "docs-oci-datasource-core-ip_sec_connection_tunnel"
description: |-
  Provides details about a specific Ip Sec Connection Tunnel in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_ipsec_connection_tunnel
This data source provides details about a specific Ip Sec Connection Tunnel resource in Oracle Cloud Infrastructure Core service.

Gets the specified IPSec connection's specified tunnel basic information.


## Example Usage

```hcl
data "oci_core_ipsec_connection_tunnel" "test_ip_sec_connection_tunnel" {
	#Required
	ipsec_id = oci_core_ipsec.test_ipsec.id
	tunnel_id = data.oci_core_ipsec_connection_tunnels.test_ip_sec_connection_tunnels.ip_sec_connection_tunnels[0].id
}
```

## Argument Reference

The following arguments are supported:

* `ipsec_id` - (Required) The OCID of the IPSec connection.
* `tunnel_id` - (Required) The OCID of the IPSec connection's tunnel.


## Attributes Reference

The following attributes are exported:

* `bgp_session_info` - Information needed to establish a BGP Session on an interface. 
	* `bgp_state` - the state of the BGP. 
	* `bgp_ipv6_state` - The state of the BGP IPv6 session.
	* `customer_bgp_asn` - This is the value of the remote Bgp ASN in asplain format, as a string. Example: 1587232876 (4 byte ASN) or 12345 (2 byte ASN) 
	* `customer_interface_ip` - This is the IPv4 Address used in the BGP peering session for the non-Oracle router. Example: 10.0.0.2/31 
	* `customer_interface_ipv6` - The IPv6 address for the CPE end of the inside tunnel interface.
	* `oracle_bgp_asn` - This is the value of the Oracle Bgp ASN in asplain format, as a string. Example: 1587232876 (4 byte ASN) or 12345 (2 byte ASN) 
	* `oracle_interface_ip` - This is the IPv4 Address used in the BGP peering session for the Oracle router. Example: 10.0.0.1/31 
    * `oracle_interface_ipv6` - The IPv6 address for the Oracle end of the inside tunnel interface.
* `compartment_id` - The OCID of the compartment containing the tunnel.
* `cpe_ip` - The IP address of Cpe headend.  Example: `129.146.17.50` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `dpd_mode` - Dead peer detection (DPD) mode set on the Oracle side of the connection.
* `dpd_timeout_in_sec` - DPD timeout in seconds.
* `encryption_domain_config` - Configuration information used by the encryption domain policy.
	* `cpe_traffic_selector` - Lists IPv4 or IPv6-enabled subnets in your on-premises network.
	* `oracle_traffic_selector` - Lists IPv4 or IPv6-enabled subnets in your Oracle tenancy.
* `nat_translation_enabled` - By default (the `AUTO` setting), IKE sends packets with a source and destination port set to 500, and when it detects that the port used to forward packets has changed (most likely because a NAT device is between the CPE device and the Oracle VPN headend) it will try to negotiate the use of NAT-T.

	The `ENABLED` option sets the IKE protocol to use port 4500 instead of 500 and forces encapsulating traffic with the ESP protocol inside UDP packets.

	The `DISABLED` option directs IKE to completely refuse to negotiate NAT-T even if it senses there may be a NAT device in use.
* `oracle_can_initiate` - Indicates whether Oracle can only respond to a request to start an IPSec tunnel from the CPE device, or both respond to and initiate requests.
* `id` - The tunnel's Oracle ID (OCID).
* `routing` - the routing strategy used for this tunnel, either static route or BGP dynamic routing
* `ike_version` - Internet Key Exchange protocol version.
* `state` - The IPSec connection's tunnel's lifecycle state.
* `status` - The tunnel's current state.
* `time_created` - The date and time the IPSec connection tunnel was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `time_status_updated` - When the status of the tunnel last changed, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `vpn_ip` - The IP address of Oracle's VPN headend.  Example: `129.146.17.50` 
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

