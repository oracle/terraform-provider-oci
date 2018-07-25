---
layout: "oci"
page_title: "OCI: oci_load_balancer_backends"
sidebar_current: "docs-oci-datasource-load_balancer-backends"
description: |-
  Provides a list of Backends
---

# Data Source: oci_load_balancer_backends
The `oci_load_balancer_backends` data source allows access to the list of OCI backends

Lists the backend servers for a given load balancer and backend set.

## Example Usage

```hcl
data "oci_load_balancer_backends" "test_backends" {
	#Required
	backendset_name = "${var.backend_backendset_name}"
	load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
}
```

## Argument Reference

The following arguments are supported:

* `backendset_name` - (Required) The name of the backend set associated with the backend servers.  Example: `example_backend_set` 
* `load_balancer_id` - (Required) The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the load balancer associated with the backend set and servers.


## Attributes Reference

The following attributes are exported:

* `backends` - The list of backends.

### Backend Reference

The following attributes are exported:

* `backendset_name` - The name of the backend set to add the backend server to.  Example: `My_backend_set`
* `backup` - Whether the load balancer should treat this server as a backup unit. If `true`, the load balancer forwards no ingress traffic to this backend server unless all other backend servers not marked as "backup" fail the health check policy.  Example: `false` 
* `drain` - Whether the load balancer should drain this server. Servers marked "drain" receive no new incoming traffic.  Example: `false` 
* `ip_address` - The IP address of the backend server.  Example: `10.0.0.3` 
* `name` - A read-only field showing the IP address and port that uniquely identify this backend server in the backend set.  Example: `10.0.0.3:8080` 
* `offline` - Whether the load balancer should treat this server as offline. Offline servers receive no incoming traffic.  Example: `false` 
* `port` - The communication port for the backend server.  Example: `8080` 
* `weight` - The load balancing policy weight assigned to the server. Backend servers with a higher weight receive a larger proportion of incoming traffic. For example, a server weighted '3' receives 3 times the number of new connections as a server weighted '1'. For more information on load balancing policies, see [How Load Balancing Policies Work](https://docs.us-phoenix-1.oraclecloud.com/Content/Balance/Reference/lbpolicies.htm).  Example: `3` 

