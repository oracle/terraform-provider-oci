---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_ipsec_connection_tunnel_routes"
sidebar_current: "docs-oci-datasource-core-ipsec_connection_tunnel_routes"
description: |-
  Provides the list of Ipsec Connection Tunnel Routes in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_ipsec_connection_tunnel_routes
This data source provides the list of Ipsec Connection Tunnel Routes in Oracle Cloud Infrastructure Core service.

The routes advertised to the on-premises network and the routes received from the on-premises network.


## Example Usage

```hcl
data "oci_core_ipsec_connection_tunnel_routes" "test_ipsec_connection_tunnel_routes" {
	#Required
	ipsec_id = oci_core_ipsec.test_ipsec.id
	tunnel_id = oci_core_tunnel.test_tunnel.id

	#Optional
	advertiser = var.ipsec_connection_tunnel_route_advertiser
}
```

## Argument Reference

The following arguments are supported:

* `advertiser` - (Optional) Specifies the advertiser of the routes. If set to `ORACLE`, this returns only the routes advertised by Oracle. When set to `CUSTOMER`, this returns only the routes advertised by the CPE. 
* `ipsec_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IPSec connection.
* `tunnel_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the tunnel.


## Attributes Reference

The following attributes are exported:

* `tunnel_routes` - The list of tunnel_routes.

### IpsecConnectionTunnelRoute Reference

The following attributes are exported:

* `advertiser` - The source of the route advertisement.
* `age` - The age of the route.
* `as_path` - A list of ASNs in AS_Path.
* `is_best_path` - Indicates this is the best route.
* `prefix` - The BGP network layer reachability information.

