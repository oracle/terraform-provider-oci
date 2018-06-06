# oci_load_balancer_path_route_set

## PathRouteSet Resource

### PathRouteSet Reference

The following attributes are exported:

* `name` - The unique name for this set of path route rules. Avoid entering confidential information.  Example: `example_path_route_set` 
* `path_routes` - The set of path route rules.
	* `backend_set_name` - The name of the target backend set for requests where the incoming URI matches the specified path.  Example: `example_backend_set` 
	* `path` - The path string to match against the incoming URI path. Path strings are case-insensitive. Asterisk (*) wildcards are not supported. Regular expressions are not supported.  Example: `/example/video/123` 
	* `path_match_type` - The type of matching to apply to incoming URIs.
		* `match_type` - Specifies how the load balancing service compares a [PathRoute](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/requests/PathRoute) object's `path` string against the incoming URI.  
		    *  **EXACT_MATCH** - Looks for a `path` string that exactly matches the incoming URI path.  
		    *  **FORCE_LONGEST_PREFIX_MATCH** - Looks for the `path` string with the best, longest match of the beginning portion of the incoming URI path.  
		    *  **PREFIX_MATCH** - Looks for a `path` string that matches the beginning portion of the incoming URI path.  
		    *  **SUFFIX_MATCH** - Looks for a `path` string that matches the ending portion of the incoming URI path.  
		    * For a full description of how the system handles `matchType` in a path route set containing multiple rules, see [Managing Request Routing](https://docs.us-phoenix-1.oraclecloud.com/Content/Balance/Tasks/managingrequest.htm). 



### Create Operation
Adds a path route set to a load balancer. For more information, see
[Managing Request Routing](https://docs.us-phoenix-1.oraclecloud.com/Content/Balance/Tasks/managingrequest.htm).


The following arguments are supported:

* `load_balancer_id` - (Required) The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the load balancer to add the path route set to.
* `name` - (Required) The name for this set of path route rules. It must be unique and it cannot be changed. Avoid entering confidential information.  Example: `example_path_route_set` 
* `path_routes` - (Required) The set of path route rules.
	* `backend_set_name` - (Required) The name of the target backend set for requests where the incoming URI matches the specified path.  Example: `example_backend_set` 
	* `path` - (Required) The path string to match against the incoming URI path. Path strings are case-insensitive. Asterisk (*) wildcards are not supported. Regular expressions are not supported.  Example: `/example/video/123` 
	* `path_match_type` - (Required) The type of matching to apply to incoming URIs.
		* `match_type` - (Required) Specifies how the load balancing service compares a [PathRoute](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/requests/PathRoute) object's `path` string against the incoming URI.  
		    *  **EXACT_MATCH** - Looks for a `path` string that exactly matches the incoming URI path.  
		    *  **FORCE_LONGEST_PREFIX_MATCH** - Looks for the `path` string with the best, longest match of the beginning    portion of the incoming URI path.  
		    *  **PREFIX_MATCH** - Looks for a `path` string that matches the beginning portion of the incoming URI path.  
		    *  **SUFFIX_MATCH** - Looks for a `path` string that matches the ending portion of the incoming URI path.  
		    * For a full description of how the system handles `matchType` in a path route set containing multiple rules, see [Managing Request Routing](https://docs.us-phoenix-1.oraclecloud.com/Content/Balance/Tasks/managingrequest.htm). 


### Update Operation
Overwrites an existing path route set on the specified load balancer. Use this operation to add, delete, or alter
path route rules in a path route set.

To add a new path route rule to a path route set, the `pathRoutes` in the
[UpdatePathRouteSetDetails](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/requests/UpdatePathRouteSetDetails) object must include
both the new path route rule to add and the existing path route rules to retain.


The following arguments support updates:
* `path_routes` - The set of path route rules.
	* `backend_set_name` - The name of the target backend set for requests where the incoming URI matches the specified path.  Example: `example_backend_set` 
	* `path` - The path string to match against the incoming URI path. Path strings are case-insensitive. Asterisk (*) wildcards are not supported. Regular expressions are not supported.  Example: `/example/video/123` 
	* `path_match_type` - The type of matching to apply to incoming URIs.
		* `match_type` - Specifies how the load balancing service compares a [PathRoute](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/requests/PathRoute) object's `path` string against the incoming URI.  
		    *  **EXACT_MATCH** - Looks for a `path` string that exactly matches the incoming URI path.  
		    *  **FORCE_LONGEST_PREFIX_MATCH** - Looks for the `path` string with the best, longest match of the beginning    portion of the incoming URI path.  
		    *  **PREFIX_MATCH** - Looks for a `path` string that matches the beginning portion of the incoming URI path.  
		    *  **SUFFIX_MATCH** - Looks for a `path` string that matches the ending portion of the incoming URI path.  
		    * For a full description of how the system handles `matchType` in a path route set containing multiple rules, see [Managing Request Routing](https://docs.us-phoenix-1.oraclecloud.com/Content/Balance/Tasks/managingrequest.htm). 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```hcl
resource "oci_load_balancer_path_route_set" "test_path_route_set" {
	#Required
	load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
	name = "${var.path_route_set_name}"
	path_routes {
		#Required
		backend_set_name = "${var.path_route_set_path_routes_backend_set_name}"
		path = "${var.path_route_set_path_routes_path}"
		path_match_type {
			#Required
			match_type = "${var.path_route_set_path_routes_path_match_type_match_type}"
		}
	}
}
```

# oci_load_balancer_path_route_sets

## PathRouteSet DataSource

Gets a list of path_route_sets.

### List Operation
Lists all path route sets associated with the specified load balancer.
The following arguments are supported:

* `load_balancer_id` - (Required) The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the load balancer associated with the path route sets to retrieve. 


The following attributes are exported:

* `path_route_sets` - The list of path_route_sets.

### Example Usage

```hcl
data "oci_load_balancer_path_route_sets" "test_path_route_sets" {
	#Required
	load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
}
```