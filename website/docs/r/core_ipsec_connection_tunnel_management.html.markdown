---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_ipsec_connection_tunnel_management"
sidebar_current: "docs-oci-datasource-core-ip_sec_connection_tunnel_management"
description: |-
  Provides details about a specific Ip Sec Connection Tunnel in Oracle Cloud Infrastructure Core service
---

# oci_core_ipsec_connection_tunnel_management
This resource provides the Ip Sec Connection Tunnel Management resource in Oracle Cloud Infrastructure Core service.

Updates the specified tunnel. This operation lets you change tunnel attributes such as the
routing type (BGP dynamic routing or static routing). Here are some important notes:

	* If you change the tunnel's routing type or BGP session configuration, the tunnel will go
	down while it's reprovisioned.

	* If you want to switch the tunnel's `routing` from `STATIC` to `BGP`, make sure the tunnel's
	BGP session configuration attributes have been set ([bgpSessionConfig](#/en/iaas/20160918/datatypes/BgpSessionInfo)).

	* If you want to switch the tunnel's `routing` from `BGP` to `STATIC`, make sure the
	[IPSecConnection](#/en/iaas/20160918/IPSecConnection/) already has at least one valid CIDR
	static route.

** IMPORTANT **
Destroying the `oci_core_ipsec_connection_tunnel_management` leaves the resource in its existing state. It will not destroy the tunnel and it will not return the tunnel to its default values.

## Example Usage

```hcl
resource "oci_core_ipsec_connection_tunnel_management" "test_ip_sec_connection_tunnel" {
	#Required
	ipsec_id = oci_core_ipsec.test_ipsec.id
	tunnel_id = data.oci_core_ipsec_connection_tunnels.test_ip_sec_connection_tunnels.ip_sec_connection_tunnels[0].id
	#Optional
	routing = var.ip_sec_connection_tunnel_management_routing
	bgp_session_info {
		#Optional
		customer_bgp_asn = var.ip_sec_connection_tunnel_management_bgp_session_info_customer_bgp_asn
		customer_interface_ip = var.ip_sec_connection_tunnel_management_bgp_session_info_customer_interface_ip
		oracle_interface_ip = var.ip_sec_connection_tunnel_management_bgp_session_info_oracle_interface_ip
	}
	display_name = var.ip_sec_connection_tunnel_management_display_name

    encryption_domain_config {
		#Optional
		cpe_traffic_selector = var.ip_sec_connection_tunnel_management_encryption_domain_config_cpe_traffic_selector
		oracle_traffic_selector = var.ip_sec_connection_tunnel_management_encryption_domain_config_oracle_traffic_selector
	}
	shared_secret = var.ip_sec_connection_tunnel_management_shared_secret
    ike_version = "V1"
}
```

## Argument Reference

The following arguments are supported:

* `ipsec_id` - (Required) The OCID of the IPSec connection.
* `tunnel_id` - (Required) The OCID of the IPSec connection's tunnel.
* `routing` - (Optional) The type of routing to use for this tunnel (either BGP dynamic routing, STATIC routing or POLICY routing). 
* `bgp_session_info` - (Optional) Information for establishing a BGP session for the IPSec tunnel. Required if the tunnel uses BGP dynamic routing.

	If the tunnel instead uses static routing, you may optionally provide this object and set an IP address for one or both ends of the IPSec tunnel for the purposes of troubleshooting or monitoring the tunnel. 
	* `customer_bgp_asn` - (Optional) If the tunnel's `routing` attribute is set to `BGP` (see [IPSecConnectionTunnel](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/IPSecConnectionTunnel/)), this ASN is required and used for the tunnel's BGP session. This is the ASN of the network on the CPE end of the BGP session. Can be a 2-byte or 4-byte ASN. Uses "asplain" format.

		If the tunnel's `routing` attribute is set to `STATIC`, the `customerBgpAsn` must be null.

		Example: `12345` (2-byte) or `1587232876` (4-byte) 
	* `customer_interface_ip` - (Optional) The IP address for the CPE end of the inside tunnel interface.

		If the tunnel's `routing` attribute is set to `BGP` (see [IPSecConnectionTunnel](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/IPSecConnectionTunnel/)), this IP address is required and used for the tunnel's BGP session.

		If `routing` is instead set to `STATIC`, this IP address is optional. You can set this IP address to troubleshoot or monitor the tunnel.

		The value must be a /30 or /31.

		Example: `10.0.0.5/31` 
	* `customer_interface_ipv6` - (Optional) The IPv6 address for the CPE end of the inside tunnel interface. This IP address is optional.

		If the tunnel's `routing` attribute is set to `BGP` (see [IPSecConnectionTunnel](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/IPSecConnectionTunnel/)), this IP address is used for the tunnel's BGP session.

		If `routing` is instead set to `STATIC`, you can set this IP address to troubleshoot or monitor the tunnel.

		Only subnet masks from /64 up to /127 are allowed.

		Example: `2001:db8::1/64`
	* `oracle_interface_ip` - (Optional) The IP address for the Oracle end of the inside tunnel interface.

		If the tunnel's `routing` attribute is set to `BGP` (see [IPSecConnectionTunnel](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/IPSecConnectionTunnel/)), this IP address is required and used for the tunnel's BGP session.

		If `routing` is instead set to `STATIC`, this IP address is optional. You can set this IP address to troubleshoot or monitor the tunnel.

		The value must be a /30 or /31.

		Example: `10.0.0.4/31` 
	* `oracle_interface_ipv6` - (Optional) The IPv6 address for the Oracle end of the inside tunnel interface. This IP address is optional.

		If the tunnel's `routing` attribute is set to `BGP` (see [IPSecConnectionTunnel](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/IPSecConnectionTunnel/)), this IP address is used for the tunnel's BGP session.

		If `routing` is instead set to `STATIC`, you can set this IP address to troubleshoot or monitor the tunnel.

		Only subnet masks from /64 up to /127 are allowed.

		Example: `2001:db8::1/64`
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `dpd_config` - (Optional)
	* `dpd_mode` - (Optional) This option defines whether DPD can be initiated from the Oracle side of the connection. `INITIATE_AND_RESPOND` or `RESPOND_ONLY`
	* `dpd_timeout_in_sec` - (Optional) DPD timeout in seconds. This sets the longest interval between CPE device health messages before the IPSec connection indicates it has lost contact with the CPE. The default is 20 seconds.
* `encryption_domain_config` - (Optional) Configuration information used by the encryption domain policy. Required if the tunnel uses POLICY routing.
	* `cpe_traffic_selector` - (Optional) Lists IPv4 or IPv6-enabled subnets in your on-premises network.
	* `oracle_traffic_selector` - (Optional) Lists IPv4 or IPv6-enabled subnets in your Oracle tenancy.
* `ike_version` - (Optional) Internet Key Exchange protocol version. 
* `shared_secret` - (Optional) The shared secret (pre-shared key) to use for the IPSec tunnel. If you don't provide a value, Oracle generates a value for you. You can specify your own shared secret later if you like with [UpdateIPSecConnectionTunnelSharedSecret](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/IPSecConnectionTunnelSharedSecret/UpdateIPSecConnectionTunnelSharedSecret).  Example: `EXAMPLEToUis6j1c.p8G.dVQxcmdfMO0yXMLi.lZTbYCMDGu4V8o`
* `nat_translation_enabled` - (Optional) By default (the `AUTO` setting), IKE sends packets with a source and destination port set to 500, and when it detects that the port used to forward packets has changed (most likely because a NAT device is between the CPE device and the Oracle VPN headend) it will try to negotiate the use of NAT-T.

  The `ENABLED` option sets the IKE protocol to use port 4500 instead of 500 and forces encapsulating traffic with the ESP protocol inside UDP packets.

  The `DISABLED` option directs IKE to completely refuse to negotiate NAT-T even if it senses there may be a NAT device in use.
* `oracle_can_initiate` - (Optional) Indicates whether Oracle can only respond to a request to start an IPSec tunnel from the CPE device (`RESPONDER_ONLY`), or both respond to and initiate requests (`INITIATOR_OR_RESPONDER`).
* `phase_one_details`- (Optional) Configuration details for IKE phase one (ISAKMP) configuration parameters. 

	See [PhaseOneConfigDetails](https://docs.oracle.com/en-us/iaas/api/#/en/iaas/20160918/datatypes/PhaseOneConfigDetails) for allowed values but note naming scheme follows [TunnelPhaseOneDetails](https://docs.oracle.com/en-us/iaas/api/#/en/iaas/20160918/datatypes/TunnelPhaseOneDetails).
  * `custom_authentication_algorithm` - (Optional) The custom authentication algorithm proposed during phase one tunnel negotiation.
  * `custom_dh_group` - (Optional) The custom Diffie-Hellman group proposed during phase one tunnel negotiation.
  * `custom_encryption_algorithm` - (Optional) The custom encryption algorithm proposed during phase one tunnel negotiation.
  * `is_custom_phase_one_config` - (Optional) Indicates whether custom configuration is enabled for phase one options.
  * `lifetime` - (Optional) Internet key association (IKE) session key lifetime in seconds for IPSec phase one. The default is 28800 which is equivalent to 8 hours.
* `phase_two_details` - (Optional) Configuration details for IPSec phase two configuration parameters.

	See [PhaseTwoConfigDetails](https://docs.oracle.com/en-us/iaas/api/#/en/iaas/20160918/datatypes/PhaseTwoConfigDetails) for allowed values, but note naming scheme follows [TunnelPhaseTwoDetails](https://docs.oracle.com/en-us/iaas/api/#/en/iaas/20160918/datatypes/TunnelPhaseTwoDetails). 
  * `custom_authentication_algorithm` - (Optional) The authentication algorithm proposed during phase two tunnel negotiation.
  * `custom_encryption_algorithm` - (Optional) The encryption algorithm proposed during phase two tunnel negotiation.
  * `dh_group` - (Optional) The Diffie-Hellman group used for PFS, if PFS is enabled.
  * `is_custom_phase_two_config` - (Optional) Indicates whether custom configuration is enabled for phase two options.
  * `is_pfs_enabled` - (Optional) Indicates whether perfect forward secrecy (PFS) is enabled.
  * `lifetime` - (Optional) Lifetime in seconds for the IPSec session key set in phase two. The default is 3600 which is equivalent to 1 hour.

## Attributes Reference

The following attributes are exported:

* `ipsec_id` - The OCID of the IPSec connection.
* `tunnel_id` - The OCID of the IPSec connection's tunnel.
* `routing` - The type of routing to use for this tunnel (either BGP dynamic routing, STATIC routing or POLICY routing).
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
* `encryption_domain_config` - Configuration information used by the encryption domain policy.
	* `cpe_traffic_selector` - Lists IPv4 or IPv6-enabled subnets in your on-premises network.
	* `oracle_traffic_selector` - Lists IPv4 or IPv6-enabled subnets in your Oracle tenancy.
* `id` - The tunnel's Oracle ID (OCID).
* `routing` - the routing strategy used for this tunnel, either static route or BGP dynamic routing
* `ike_version` - Internet Key Exchange protocol version.
* `state` - The IPSec connection's tunnel's lifecycle state.
* `status` - The tunnel's current state.
* `time_created` - The date and time the IPSec connection tunnel was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `time_status_updated` - When the status of the tunnel last changed, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `vpn_ip` - The IP address of Oracle's VPN headend.  Example: `129.146.17.50` 
* `dpd_mode` - Dead peer detection (DPD) mode set on the Oracle side of the connection.
* `dpd_timeout_in_sec` - DPD timeout in seconds.
* `nat_translation_enabled` - By default (the `AUTO` setting), IKE sends packets with a source and destination port set to 500, and when it detects that the port used to forward packets has changed (most likely because a NAT device is between the CPE device and the Oracle VPN headend) it will try to negotiate the use of NAT-T.

  The `ENABLED` option sets the IKE protocol to use port 4500 instead of 500 and forces encapsulating traffic with the ESP protocol inside UDP packets.

  The `DISABLED` option directs IKE to completely refuse to negotiate NAT-T even if it senses there may be a NAT device in use.
* `oracle_can_initiate` -  Indicates whether Oracle can only respond to a request to start an IPSec tunnel from the CPE device, or both respond to and initiate requests.
* `phase_one_details`- (Optional) IPSec tunnel details specific to ISAKMP phase one. See [TunnelPhaseOneDetails](https://docs.oracle.com/en-us/iaas/api/#/en/iaas/20160918/datatypes/TunnelPhaseOneDetails).
	* `custom_authentication_algorithm` - The proposed custom authentication algorithm.
	* `custom_dh_group` - The proposed custom authentication algorithm.
	* `custom_encryption_algorithm` - The proposed custom encryption algorithm.
	* `is_custom_phase_one_config` - Indicates whether custom phase one configuration is enabled. If this option is not enabled, default settings are proposed.
	* `lifetime` - The total configured lifetime of the IKE security association.
    * `is_ike_established` - Indicates whether IKE phase one is established.
    * `negotiated_authentication_algorithm` - The negotiated authentication algorithm.
    * `negotiated_dh_group` - The negotiated Diffie-Hellman group.
    * `negotiated_encryption_algorithm` - The negotiated encryption algorithm.
    * `remaining_lifetime_int` - The remaining lifetime before the key is refreshed.
    * `remaining_lifetime_last_retrieved` - The date and time we retrieved the remaining lifetime, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z`
* `phase_two_details` - (IPsec tunnel detail information specific to phase two. See [TunnelPhaseTwoDetails](https://docs.oracle.com/en-us/iaas/api/#/en/iaas/20160918/datatypes/TunnelPhaseTwoDetails).
	* `custom_authentication_algorithm` - Phase two authentication algorithm proposed during tunnel negotiation.
	* `custom_encryption_algorithm` - The proposed custom phase two encryption algorithm.
	* `dh_group` - The proposed Diffie-Hellman group.
	* `is_custom_phase_two_config` - Indicates whether custom phase two configuration is enabled. If this option is not enabled, default settings are proposed.
	* `is_pfs_enabled` - Indicates that PFS (perfect forward secrecy) is enabled.
	* `lifetime` - The total configured lifetime of the IKE security association.
    * `is_esp_established` - Indicates that ESP phase two is established.
    * `negotiated_authentication_algorithm` - The negotiated phase two authentication algorithm.
    * `negotiated_dh_group` - The negotiated Diffie-Hellman group.
    * `negotiated_encryption_algorithm` - The negotiated encryption algorithm.
    * `remaining_lifetime_int` - The remaining lifetime before the key is refreshed.
    * `remaining_lifetime_last_retrieved` - The date and time the remaining lifetime was last retrieved, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z`

