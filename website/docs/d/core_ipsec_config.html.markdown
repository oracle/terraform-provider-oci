---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_ipsec_config"
sidebar_current: "docs-oci-datasource-core-ip_sec_connection_device_config"
description: |-
  Provides details about a specific Ip Sec Connection Device Config in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_ipsec_config
This data source provides details about a specific Ip Sec Connection Device Config resource in Oracle Cloud Infrastructure Core service.

Deprecated. To get tunnel information, instead use:

* [GetIPSecConnectionTunnel](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/IPSecConnectionTunnel/GetIPSecConnectionTunnel)
* [GetIPSecConnectionTunnelSharedSecret](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/IPSecConnectionTunnelSharedSecret/GetIPSecConnectionTunnelSharedSecret)


## Example Usage

```hcl
data "oci_core_ipsec_config" "test_ip_sec_connection_device_config" {
	#Required
	ipsec_id = oci_core_ipsec.test_ipsec.id
}
```

## Argument Reference

The following arguments are supported:

* `ipsec_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IPSec connection.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the IPSec connection.
* `id` - The IPSec connection's Oracle ID ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)).
* `time_created` - The date and time the IPSec connection was created.
* `tunnels` - Two [TunnelConfig](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/TunnelConfig/) objects.
	* `ip_address` - The IP address of Oracle's VPN headend.  Example: `203.0.113.50 ` 
	* `shared_secret` - The shared secret of the IPSec tunnel. 
	* `time_created` - The date and time the IPSec connection was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

