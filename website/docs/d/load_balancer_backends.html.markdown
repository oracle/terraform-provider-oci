---
subcategory: "Load Balancer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_load_balancer_backends"
sidebar_current: "docs-oci-datasource-load_balancer-backends"
description: |-
  Provides the list of Backends in Oracle Cloud Infrastructure Load Balancer service
---

# Data Source: oci_load_balancer_backends
This data source provides the list of Backends in Oracle Cloud Infrastructure Load Balancer service.

Lists the backend servers for a given load balancer and backend set.

## Example Usage

```hcl
data "oci_load_balancer_backends" "test_backends" {
	#Required
	backendset_name = oci_load_balancer_backend_set.test_backend_set.name
	load_balancer_id = oci_load_balancer_load_balancer.test_load_balancer.id
}
```

## Argument Reference

The following arguments are supported:

* `backendset_name` - (Required) The name of the backend set associated with the backend servers.  Example: `example_backend_set` 
* `load_balancer_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the load balancer associated with the backend set and servers.


## Attributes Reference

The following attributes are exported:

* `backends` - The list of backends.

### Backend Reference

The following attributes are exported:

* `backup` - Whether the load balancer should treat this server as a backup unit. If `true`, the load balancer forwards no ingress traffic to this backend server unless all other backend servers not marked as "backup" fail the health check policy.

	**Note:** You cannot add a backend server marked as `backup` to a backend set that uses the IP Hash policy.

	Example: `false` 
* `drain` - Whether the load balancer should drain this server. Servers marked "drain" receive no new incoming traffic.  Example: `false` 
* `ip_address` - The IP address of the backend server.  Example: `10.0.0.3` 
* `max_connections` - The maximum number of simultaneous connections the load balancer can make to the backend. If this is not set then the maximum number of simultaneous connections the load balancer can make to the backend is unlimited.  Example: `300` 
* `name` - A read-only field showing the IP address and port that uniquely identify this backend server in the backend set.  Example: `10.0.0.3:8080` 
* `offline` - Whether the load balancer should treat this server as offline. Offline servers receive no incoming traffic.  Example: `false` 
* `port` - The communication port for the backend server.  Example: `8080` 
* `weight` - The load balancing policy weight assigned to the server. Backend servers with a higher weight receive a larger proportion of incoming traffic. For example, a server weighted '3' receives 3 times the number of new connections as a server weighted '1'. For more information on load balancing policies, see [How Load Balancing Policies Work](https://docs.cloud.oracle.com/iaas/Content/Balance/Reference/lbpolicies.htm).  Example: `3` 

