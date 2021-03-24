---
subcategory: "Network Load Balancer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_load_balancer_backend"
sidebar_current: "docs-oci-datasource-network_load_balancer-backend"
description: |-
  Provides details about a specific Backend in Oracle Cloud Infrastructure Network Load Balancer service
---

# Data Source: oci_network_load_balancer_backend
This data source provides details about a specific Backend resource in Oracle Cloud Infrastructure Network Load Balancer service.

Retrieves the configuration information for the specified backend server.

## Example Usage

```hcl
data "oci_network_load_balancer_backend" "test_backend" {
	#Required
	backend_name = oci_network_load_balancer_backend.test_backend.name
	backend_set_name = oci_network_load_balancer_backend_set.test_backend_set.name
	network_load_balancer_id = oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id
}
```

## Argument Reference

The following arguments are supported:

* `backend_name` - (Required) The name of the backend server to retrieve. This is specified as <ip>:<port>, or as <ip> <OCID>:<port>.  Example: `10.0.0.3:8080` or `ocid1.privateip..oc1.<var>&lt;unique_ID&gt;</var>:8080` 
* `backend_set_name` - (Required) The name of the backend set that includes the backend server.  Example: `example_backend_set` 
* `network_load_balancer_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network load balancer to update.


## Attributes Reference

The following attributes are exported:

* `ip_address` - The IP address of the backend server. Example: `10.0.0.3` 
* `is_backup` - Whether the network load balancer should treat this server as a backup unit. If `true`, then the network load balancer forwards no ingress traffic to this backend server unless all other backend servers not marked as "isBackup" fail the health check policy.  Example: `false` 
* `is_drain` - Whether the network load balancer should drain this server. Servers marked "isDrain" receive no  incoming traffic.  Example: `false` 
* `is_offline` - Whether the network load balancer should treat this server as offline. Offline servers receive no incoming traffic.  Example: `false` 
* `name` - A read-only field showing the IP address/IP OCID and port that uniquely identify this backend server in the backend set.  Example: `10.0.0.3:8080`, or `ocid1.privateip..oc1.<var>&lt;unique_ID&gt;</var>:443` or `10.0.0.3:0` 
* `port` - The communication port for the backend server.  Example: `8080` 
* `target_id` - The IP OCID/Instance OCID associated with the backend server. Example: `ocid1.privateip..oc1.<var>&lt;unique_ID&gt;</var>` 
* `weight` - The network load balancing policy weight assigned to the server. Backend servers with a higher weight receive a larger proportion of incoming traffic. For example, a server weighted '3' receives three times the number of new connections as a server weighted '1'. For more information about load balancing policies, see [How Network Load Balancing Policies Work](https://docs.cloud.oracle.com/iaas/Content/Balance/Reference/lbpolicies.htm).  Example: `3` 

