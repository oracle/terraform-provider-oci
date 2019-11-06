---
subcategory: "Load Balancer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_load_balancer_path_route_sets"
sidebar_current: "docs-oci-datasource-load_balancer-path_route_sets"
description: |-
  Provides the list of Path Route Sets in Oracle Cloud Infrastructure Load Balancer service
---

# Data Source: oci_load_balancer_path_route_sets
This data source provides the list of Path Route Sets in Oracle Cloud Infrastructure Load Balancer service.

Lists all path route sets associated with the specified load balancer.

## Example Usage

```hcl
data "oci_load_balancer_path_route_sets" "test_path_route_sets" {
	#Required
	load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
}
```

## Argument Reference

The following arguments are supported:

* `load_balancer_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the load balancer associated with the path route sets to retrieve. 


## Attributes Reference

The following attributes are exported:

* `path_route_sets` - The list of path_route_sets.

### PathRouteSet Reference

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

