---
subcategory: "Network Load Balancer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_load_balancer_network_load_balancers_policies"
sidebar_current: "docs-oci-datasource-network_load_balancer-network_load_balancers_policies"
description: |-
  Provides the list of Network Load Balancers Policies in Oracle Cloud Infrastructure Network Load Balancer service
---

# Data Source: oci_network_load_balancer_network_load_balancers_policies
This data source provides the list of Network Load Balancers Policies in Oracle Cloud Infrastructure Network Load Balancer service.

Lists the available network load balancer policies.

## Example Usage

```hcl
data "oci_network_load_balancer_network_load_balancers_policies" "test_network_load_balancers_policies" {
}
```

## Argument Reference

The following arguments are supported:



## Attributes Reference

The following attributes are exported:

* `network_load_balancers_policy_collection` - The list of network_load_balancers_policy_collection.

### NetworkLoadBalancersPolicy Reference

The following attributes are exported:

* `items` - Array of NetworkLoadBalancersPolicySummary objects.

