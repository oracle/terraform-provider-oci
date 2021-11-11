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

* `ip_sec_connection_tunnels` - The list of ip_sec_connection_tunnels.

### IpSecConnectionTunnel Reference

The following attributes are exported:

* `bgp_session_info` - Information for establishing a BGP session for the IPSec tunnel.
	* `bgp_ipv6state` - The state of the BGP IPv6 session. 
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
* `cpe_ip` - The IP address of the CPE's VPN headend.  Example: `203.0.113.22` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `dpd_mode` - dpd mode
* `dpd_timeout_in_sec` - Dead peer detection (DPD) timeout in seconds.
* `encryption_domain_config` - Configuration information used by the encryption domain policy.
	* `cpe_traffic_selector` - Lists IPv4 or IPv6-enabled subnets in your on-premises network.
	* `oracle_traffic_selector` - Lists IPv4 or IPv6-enabled subnets in your Oracle tenancy.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the tunnel.
* `ike_version` - Internet Key Exchange protocol version. 
* `nat_translation_enabled` - Whether NAT-T Enabled on the tunnel
* `oracle_can_initiate` - Indicates whether Oracle can either initiate the tunnel or respond, or respond only.
* `phase_one_details` - Tunnel detail information specific to IPSec phase 1.
	* `custom_authentication_algorithm` - Custom authentication algorithm 
	* `custom_dh_group` - Custom Diffie-Hellman group. 
	* `custom_encryption_algorithm` - Custom encryption algorithm. 
	* `is_custom_phase_one_config` - Indicates whether custom phase one configuration is enabled.
	* `is_ike_established` - Indicates whether IKE Phase 1 is established.
	* `lifetime` - The total configured lifetime of an IKE security association.
	* `negotiated_authentication_algorithm` - The negotiated authentication algorithm.
	* `negotiated_dh_group` - The negotiated Diffie-Hellman group.
	* `negotiated_encryption_algorithm` - The negotiated encryption algorithm.
	* `remaining_lifetime` - The lifetime remaining before the key is refreshed.
	* `remaining_lifetime_last_retrieved` - The date and time we retrieved the remaining lifetime, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `phase_two_details` - Tunnel detail information specific to IPSec phase 2.
	* `custom_authentication_algorithm` - Phase Two authentication algorithm supported during tunnel negotiation. 
	* `custom_encryption_algorithm` - Custom Encryption Algorithm 
	* `dh_group` - Proposed Diffie-Hellman group. 
	* `is_custom_phase_two_config` - Indicates whether custom phase two configuration is enabled.
	* `is_esp_established` - ESP Phase 2 established
	* `is_pfs_enabled` - Is PFS (perfect forward secrecy) enabled
	* `lifetime` - The total configured lifetime of an IKE security association.
	* `negotiated_authentication_algorithm` - The negotiated authentication algorithm.
	* `negotiated_dh_group` - The negotiated Diffie-Hellman group.
	* `negotiated_encryption_algorithm` - The negotiated encryption algorithm.
	* `remaining_lifetime` - The lifetime remaining before the key is refreshed.
	* `remaining_lifetime_last_retrieved` - The date and time we retrieved the remaining lifetime, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `routing` - The type of routing used for this tunnel (either BGP dynamic routing or static routing). 
* `state` - The tunnel's lifecycle state.
* `status` - The status of the tunnel based on IPSec protocol characteristics.
* `time_created` - The date and time the IPSec connection tunnel was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_status_updated` - When the status of the tunnel last changed, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `vpn_ip` - The IP address of Oracle's VPN headend.  Example: `203.0.113.21` 

