---
subcategory: "Load Balancer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_load_balancer_load_balancer_routing_policies"
sidebar_current: "docs-oci-datasource-load_balancer-load_balancer_routing_policies"
description: |-
  Provides the list of Load Balancer Routing Policies in Oracle Cloud Infrastructure Load Balancer service
---

# Data Source: oci_load_balancer_load_balancer_routing_policies
This data source provides the list of Load Balancer Routing Policies in Oracle Cloud Infrastructure Load Balancer service.

Lists all routing policies associated with the specified load balancer.

## Example Usage

```hcl
data "oci_load_balancer_load_balancer_routing_policies" "test_load_balancer_routing_policies" {
	#Required
	load_balancer_id = oci_load_balancer_load_balancer.test_load_balancer.id
}
```

## Argument Reference

The following arguments are supported:

* `load_balancer_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the load balancer associated with the routing policies. 


## Attributes Reference

The following attributes are exported:

* `routing_policies` - The list of routing_policies.

### LoadBalancerRoutingPolicy Reference

The following attributes are exported:

* `condition_language_version` - The version of the language in which `condition` of `rules` are composed. 
* `name` - The unique name for this list of routing rules. Avoid entering confidential information.  Example: `example_routing_policy` 
* `rules` - The ordered list of routing rules.
	* `actions` - A list of actions to be applied when conditions of the routing rule are met. 
		* `backend_set_name` - Name of the backend set the listener will forward the traffic to.  Example: `backendSetForImages` 
		* `name` - The name can be one of these values: `FORWARD_TO_BACKENDSET`
	* `condition` - A routing rule to evaluate defined conditions against the incoming HTTP request and perform an action. 
	* `name` - A unique name for the routing policy rule. Avoid entering confidential information. 

