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

Gets the configuration information for the specified IPSec connection. For each tunnel, the
response includes the IP address of Oracle's VPN headend and the shared secret.


## Example Usage

```hcl
data "oci_core_ipsec_config" "test_ip_sec_connection_device_config" {
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
* `time_created` - The date and time the IPSec connection was created.
* `tunnels` - Two [TunnelConfig](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/TunnelConfig/) objects.
	* `ip_address` - The IP address of Oracle's VPN headend.  Example: `129.146.17.50` 
	* `shared_secret` - The shared secret of the IPSec tunnel.  Example: `EXAMPLEToUis6j1cp8GdVQxcmdfMO0yXMLilZTbYCMDGu4V8o` 
	* `time_created` - The date and time the IPSec connection was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 

