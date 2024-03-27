---
subcategory: "Network Load Balancer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_load_balancer_listeners"
sidebar_current: "docs-oci-datasource-network_load_balancer-listeners"
description: |-
  Provides the list of Listeners in Oracle Cloud Infrastructure Network Load Balancer service
---

# Data Source: oci_network_load_balancer_listeners
This data source provides the list of Listeners in Oracle Cloud Infrastructure Network Load Balancer service.

Lists all listeners associated with a given network load balancer.

## Example Usage

```hcl
data "oci_network_load_balancer_listeners" "test_listeners" {
	#Required
	network_load_balancer_id = oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id
}
```

## Argument Reference

The following arguments are supported:

* `network_load_balancer_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network load balancer to update.


## Attributes Reference

The following attributes are exported:

* `listener_collection` - The list of listener_collection.

### Listener Reference

The following attributes are exported:

* `default_backend_set_name` - The name of the associated backend set.  Example: `example_backend_set` 
* `ip_version` - IP version associated with the listener.
* `name` - A friendly name for the listener. It must be unique and it cannot be changed.  Example: `example_listener` 
* `port` - The communication port for the listener.  Example: `80` 
* `protocol` - The protocol on which the listener accepts connection requests. For public network load balancers, ANY protocol refers to TCP/UDP with the wildcard port. For private network load balancers, ANY protocol refers to TCP/UDP/ICMP (note that ICMP requires isPreserveSourceDestination to be set to true). "ListNetworkLoadBalancersProtocols" API is deprecated and it will not return the updated values. Use the allowed values for the protocol instead.  Example: `TCP` 

