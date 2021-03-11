---
subcategory: "Load Balancer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_load_balancer_load_balancer_routing_policy"
sidebar_current: "docs-oci-datasource-load_balancer-load_balancer_routing_policy"
description: |-
  Provides details about a specific Load Balancer Routing Policy in Oracle Cloud Infrastructure Load Balancer service
---

# Data Source: oci_load_balancer_load_balancer_routing_policy
This data source provides details about a specific Load Balancer Routing Policy resource in Oracle Cloud Infrastructure Load Balancer service.

Gets the specified routing policy.

## Example Usage

```hcl
data "oci_load_balancer_load_balancer_routing_policy" "test_load_balancer_routing_policy" {
	#Required
	load_balancer_id = oci_load_balancer_load_balancer.test_load_balancer.id
	routing_policy_name = oci_load_balancer_routing_policy.test_routing_policy.name
}
```

## Argument Reference

The following arguments are supported:

* `load_balancer_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the specified load balancer.
* `routing_policy_name` - (Required) The name of the routing policy to retrieve.  Example: `example_routing_policy` 


## Attributes Reference

The following attributes are exported:

* `condition_language_version` - The version of the language in which `condition` of `rules` are composed. 
* `name` - The unique name for this list of routing rules. Avoid entering confidential information.  Example: `example_routing_policy` 
* `rules` - The ordered list of routing rules.
	* `actions` - A list of actions to be applied when conditions of the routing rule are met. 
		* `backend_set_name` - Name of the backend set the listener will forward the traffic to.  Example: `backendSetForImages` 
		* `name` - The name can be one of these values: `FORWARD_TO_BACKENDSET`
	* `condition` - A routing rule to evaluate defined conditions against the incoming HTTP request and perform an action. 
	* `name` - A unique name for the routing policy rule. Avoid entering confidential information. 

