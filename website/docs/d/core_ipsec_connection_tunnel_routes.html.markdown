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

The routes advertised to the Customer and the routes received from the Customer.


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

* `advertiser` - (Optional) Specifies the advertiser of the routes. If set to ORACLE, then returns only the routes advertised by ORACLE, else if set to CUSTOMER, then returns only the routes advertised by the CUSTOMER. 
* `ipsec_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IPSec connection.
* `tunnel_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the tunnel.


## Attributes Reference

The following attributes are exported:

* `tunnel_routes` - The list of tunnel_routes.

### IpsecConnectionTunnelRoute Reference

The following attributes are exported:

* `advertiser` - Route advertiser
* `age` - The age of the route
* `as_path` - List of ASNs in AS Path
* `is_best_path` - Is this the best route
* `prefix` - BGP Network Layer Reachability Information

