---
subcategory: "Load Balancer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_load_balancer_load_balancer_routing_policy"
sidebar_current: "docs-oci-resource-load_balancer-load_balancer_routing_policy"
description: |-
  Provides the Load Balancer Routing Policy resource in Oracle Cloud Infrastructure Load Balancer service
---

# oci_load_balancer_load_balancer_routing_policy
This resource provides the Load Balancer Routing Policy resource in Oracle Cloud Infrastructure Load Balancer service.

Adds a routing policy to a load balancer. For more information, see
[Managing Request Routing](https://docs.cloud.oracle.com/iaas/Content/Balance/Tasks/managingrequest.htm).


## Example Usage

```hcl
resource "oci_load_balancer_load_balancer_routing_policy" "test_load_balancer_routing_policy" {
	#Required
	condition_language_version = var.load_balancer_routing_policy_condition_language_version
	load_balancer_id = oci_load_balancer_load_balancer.test_load_balancer.id
	name = var.load_balancer_routing_policy_name
	rules {
		#Required
		actions {
			#Required
			name = var.load_balancer_routing_policy_rules_actions_name

			#Optional
			backend_set_name = oci_load_balancer_backend_set.test_backend_set.name
		}
		condition = var.load_balancer_routing_policy_rules_condition
		name = var.load_balancer_routing_policy_rules_name
	}
}
```

## Argument Reference

The following arguments are supported:

* `condition_language_version` - (Required) (Updatable) The version of the language in which `condition` of `rules` are composed. 
* `load_balancer_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the load balancer to add the routing policy rule list to.
* `name` - (Required) The name for this list of routing rules. It must be unique and it cannot be changed. Avoid entering confidential information.  Example: `example_routing_rules` 
* `rules` - (Required) (Updatable) The list of routing rules.
	* `actions` - (Required) (Updatable) A list of actions to be applied when conditions of the routing rule are met. 
		* `backend_set_name` - (Optional) (Updatable) Name of the backend set the listener will forward the traffic to.  Example: `backendSetForImages` 
		* `name` - (Required) (Updatable) The name can be one of these values: `FORWARD_TO_BACKENDSET`
	* `condition` - (Required) (Updatable) A routing rule to evaluate defined conditions against the incoming HTTP request and perform an action. 
	* `name` - (Required) (Updatable) A unique name for the routing policy rule. Avoid entering confidential information. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Load Balancer Routing Policy
	* `update` - (Defaults to 20 minutes), when updating the Load Balancer Routing Policy
	* `delete` - (Defaults to 20 minutes), when destroying the Load Balancer Routing Policy


## Import

LoadBalancerRoutingPolicies can be imported using the `id`, e.g.

```
$ terraform import oci_load_balancer_load_balancer_routing_policy.test_load_balancer_routing_policy "loadBalancers/{loadBalancerId}/routingPolicies/{routingPolicyName}" 
```

