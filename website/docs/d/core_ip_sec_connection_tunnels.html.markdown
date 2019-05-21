---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_ipsec_connection_tunnels"
sidebar_current: "docs-oci-datasource-core-ipsec_connection_tunnels"
description: |-
  Provides the list of Ip Sec Connection Tunnels in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_ipsec_connection_tunnels
This data source provides the list of Ip Sec Connection Tunnels in Oracle Cloud Infrastructure Core service.

Gets the lists of tunnel information for the specified IPSec connection.


## Example Usage

```hcl
data "oci_core_ipsec_connection_tunnels" "test_ip_sec_connection_tunnels" {
	#Required
	ipsec_id = "${oci_core_ipsec.test_ipsec.id}"
}
```

## Argument Reference

The following arguments are supported:

* `ipsec_id` - (Required) The OCID of the IPSec connection.


## Attributes Reference

The following attributes are exported:

* `ip_sec_connection_tunnels` - The list of ip_sec_connection_tunnels.

### IpSecConnectionTunnel Reference

The following attributes are exported:

* `bgp_session_info` - Information needed to establish a BGP Session on an interface. 
	* `bgp_state` - the state of the BGP. 
	* `customer_bgp_asn` - This is the value of the remote Bgp ASN in asplain format, as a string. Example: 1587232876 (4 byte ASN) or 12345 (2 byte ASN) 
	* `customer_interface_ip` - This is the IPv4 Address used in the BGP peering session for the non-Oracle router. Example: 10.0.0.2/31 
	* `oracle_bgp_asn` - This is the value of the Oracle Bgp ASN in asplain format, as a string. Example: 1587232876 (4 byte ASN) or 12345 (2 byte ASN) 
	* `oracle_interface_ip` - This is the IPv4 Address used in the BGP peering session for the Oracle router. Example: 10.0.0.1/31 
* `compartment_id` - The OCID of the compartment containing the tunnel.
* `cpe_ip` - The IP address of Cpe headend.  Example: `129.146.17.50` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `id` - The tunnel's Oracle ID (OCID).
* `routing` - the routing strategy used for this tunnel, either static route or BGP dynamic routing
* `state` - The IPSec connection's tunnel's lifecycle state.
* `status` - The tunnel's current state.
* `time_created` - The date and time the IPSec connection tunnel was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `time_status_updated` - When the status of the tunnel last changed, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `vpn_ip` - The IP address of Oracle's VPN headend.  Example: `129.146.17.50` 

