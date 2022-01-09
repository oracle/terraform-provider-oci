---
subcategory: "Network Load Balancer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_load_balancer_network_load_balancers_protocols"
sidebar_current: "docs-oci-datasource-network_load_balancer-network_load_balancers_protocols"
description: |-
  Provides the list of Network Load Balancers Protocols in Oracle Cloud Infrastructure Network Load Balancer service
---

# Data Source: oci_network_load_balancer_network_load_balancers_protocols
This data source provides the list of Network Load Balancers Protocols in Oracle Cloud Infrastructure Network Load Balancer service.

This API has been deprecated so it won't return the updated list of supported protocls.
Lists all supported traffic protocols.


## Example Usage

```hcl
data "oci_network_load_balancer_network_load_balancers_protocols" "test_network_load_balancers_protocols" {
}
```

## Argument Reference

The following arguments are supported:



## Attributes Reference

The following attributes are exported:

* `network_load_balancers_protocol_collection` - The list of network_load_balancers_protocol_collection.

### NetworkLoadBalancersProtocol Reference

The following attributes are exported:

* `items` - Array of NetworkLoadBalancersProtocolSummary objects.

