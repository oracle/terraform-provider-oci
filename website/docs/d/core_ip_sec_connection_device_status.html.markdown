---
layout: "oci"
page_title: "OCI: oci_core_ipsec_status"
sidebar_current: "docs-oci-datasource-core-ipsec_status"
description: |-
  Provides details about a specific IpSecConnectionDeviceStatus
---

# Data Source: oci_core_ipsec_status
The IpSecConnectionDeviceStatus data source provides details about a specific IpSecConnectionDeviceStatus

Gets the status of the specified IPSec connection (whether it's up or down).


## Example Usage

```hcl
data "oci_core_ipsec_status" "test_ip_sec_connection_device_status" {
	#Required
	ipsec_id = "${oci_core_ipsec.test_ipsec.id}"
}
```

## Argument Reference

The following arguments are supported:

* `ipsec_id` - (Required) The OCID of the IPSec connection.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the IPSec connection.
* `id` - The IPSec connection's Oracle ID (OCID).
* `time_created` - The date and time the IPSec connection was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `tunnels` - Two [TunnelStatus](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/TunnelStatus/) objects.
	* `ip_address` - The IP address of Oracle's VPN headend.  Example: `129.146.17.50` 
	* `state` - The tunnel's current state.
	* `time_created` - The date and time the IPSec connection was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
	* `time_state_modified` - When the state of the tunnel last changed, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 

