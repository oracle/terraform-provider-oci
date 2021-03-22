---
subcategory: "Load Balancer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_load_balancer_path_route_set"
sidebar_current: "docs-oci-resource-load_balancer-path_route_set"
description: |-
  Provides the Path Route Set resource in Oracle Cloud Infrastructure Load Balancer service
---

# oci_load_balancer_path_route_set
This resource provides the Path Route Set resource in Oracle Cloud Infrastructure Load Balancer service.

Adds a path route set to a load balancer. For more information, see
[Managing Request Routing](https://docs.cloud.oracle.com/iaas/Content/Balance/Tasks/managingrequest.htm).


## Example Usage

```hcl
resource "oci_load_balancer_path_route_set" "test_path_route_set" {
	#Required
	load_balancer_id = oci_load_balancer_load_balancer.test_load_balancer.id
	name = var.path_route_set_name
	path_routes {
		#Required
		backend_set_name = oci_load_balancer_backend_set.test_backend_set.name
		path = var.path_route_set_path_routes_path
		path_match_type {
			#Required
			match_type = var.path_route_set_path_routes_path_match_type_match_type
		}
	}
}
```

## Argument Reference

The following arguments are supported:

* `load_balancer_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the load balancer to add the path route set to.
* `name` - (Required) The name for this set of path route rules. It must be unique and it cannot be changed. Avoid entering confidential information.  Example: `example_path_route_set` 
* `path_routes` - (Required) (Updatable) The set of path route rules.
	* `backend_set_name` - (Required) (Updatable) The name of the target backend set for requests where the incoming URI matches the specified path.  Example: `example_backend_set` 
	* `path` - (Required) (Updatable) The path string to match against the incoming URI path.
		*  Path strings are case-insensitive.
		*  Asterisk (*) wildcards are not supported.
		*  Regular expressions are not supported.

		Example: `/example/video/123` 
	* `path_match_type` - (Required) (Updatable) The type of matching to apply to incoming URIs.
		* `match_type` - (Required) (Updatable) Specifies how the load balancing service compares a [PathRoute](https://docs.cloud.oracle.com/iaas/api/#/en/loadbalancer/20170115/requests/PathRoute) object's `path` string against the incoming URI.
			*  **EXACT_MATCH** - Looks for a `path` string that exactly matches the incoming URI path.
			*  **FORCE_LONGEST_PREFIX_MATCH** - Looks for the `path` string with the best, longest match of the beginning portion of the incoming URI path.
			*  **PREFIX_MATCH** - Looks for a `path` string that matches the beginning portion of the incoming URI path.
			*  **SUFFIX_MATCH** - Looks for a `path` string that matches the ending portion of the incoming URI path.

			For a full description of how the system handles `matchType` in a path route set containing multiple rules, see [Managing Request Routing](https://docs.cloud.oracle.com/iaas/Content/Balance/Tasks/managingrequest.htm). 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `name` - The unique name for this set of path route rules. Avoid entering confidential information.  Example: `example_path_route_set` 
* `path_routes` - The set of path route rules.
	* `backend_set_name` - The name of the target backend set for requests where the incoming URI matches the specified path.  Example: `example_backend_set` 
	* `path` - The path string to match against the incoming URI path.
		*  Path strings are case-insensitive.
		*  Asterisk (*) wildcards are not supported.
		*  Regular expressions are not supported.

		Example: `/example/video/123` 
	* `path_match_type` - The type of matching to apply to incoming URIs.
		* `match_type` - Specifies how the load balancing service compares a [PathRoute](https://docs.cloud.oracle.com/iaas/api/#/en/loadbalancer/20170115/requests/PathRoute) object's `path` string against the incoming URI.
			*  **EXACT_MATCH** - Looks for a `path` string that exactly matches the incoming URI path.
			*  **FORCE_LONGEST_PREFIX_MATCH** - Looks for the `path` string with the best, longest match of the beginning portion of the incoming URI path.
			*  **PREFIX_MATCH** - Looks for a `path` string that matches the beginning portion of the incoming URI path.
			*  **SUFFIX_MATCH** - Looks for a `path` string that matches the ending portion of the incoming URI path.

			For a full description of how the system handles `matchType` in a path route set containing multiple rules, see [Managing Request Routing](https://docs.cloud.oracle.com/iaas/Content/Balance/Tasks/managingrequest.htm). 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Path Route Set
	* `update` - (Defaults to 20 minutes), when updating the Path Route Set
	* `delete` - (Defaults to 20 minutes), when destroying the Path Route Set


## Import

PathRouteSets can be imported using the `id`, e.g.

```
$ terraform import oci_load_balancer_path_route_set.test_path_route_set "loadBalancers/{loadBalancerId}/pathRouteSets/{pathRouteSetName}" 
```

