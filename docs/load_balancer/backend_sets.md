# oci_load_balancer_backend_set
`oci_load_balancer_backendset` is the old name for `oci_load_balancer_backend_set`. Both names are supported but `oci_load_balancer_backend_set` is used in the docs.

## BackendSet Resource

### BackendSet Reference

The following attributes are exported:

* `backend` - 
	* `backup` - Whether the load balancer should treat this server as a backup unit. If `true`, the load balancer forwards no ingress traffic to this backend server unless all other backend servers not marked as "backup" fail the health check policy.  Example: `false` 
	* `drain` - Whether the load balancer should drain this server. Servers marked "drain" receive no new incoming traffic.  Example: `false` 
	* `ip_address` - The IP address of the backend server.  Example: `10.0.0.3` 
	* `name` - A read-only field showing the IP address and port that uniquely identify this backend server in the backend set.  Example: `10.0.0.3:8080` 
	* `offline` - Whether the load balancer should treat this server as offline. Offline servers receive no incoming traffic.  Example: `false` 
	* `port` - The communication port for the backend server.  Example: `8080` 
	* `weight` - The load balancing policy weight assigned to the server. Backend servers with a higher weight receive a larger proportion of incoming traffic. For example, a server weighted '3' receives 3 times the number of new connections as a server weighted '1'. For more information on load balancing policies, see [How Load Balancing Policies Work](https://docs.us-phoenix-1.oraclecloud.com/Content/Balance/Reference/lbpolicies.htm).  Example: `3` 
* `health_checker` - 
	* `interval_ms` - The interval between health checks, in milliseconds. The default is 30000 (30 seconds).  Example: `30000` 
	* `port` - The backend server port against which to run the health check. If the port is not specified, the load balancer uses the port information from the `Backend` object.  Example: `8080` 
	* `protocol` - The protocol the health check must use; either HTTP or TCP.  Example: `HTTP` 
	* `response_body_regex` - A regular expression for parsing the response body from the backend server.  Example: `^((?!false).|\s)*$` 
	* `retries` - The number of retries to attempt before a backend server is considered "unhealthy". Defaults to 3.  Example: `3` 
	* `return_code` - The status code a healthy backend server should return. If you configure the health check policy to use the HTTP protocol, you can use common HTTP status codes such as "200".  Example: `200` 
	* `timeout_in_millis` - The maximum time, in milliseconds, to wait for a reply to a health check. A health check is successful only if a reply returns within this timeout period. Defaults to 3000 (3 seconds).  Example: `3000` 
	* `url_path` - The path against which to run the health check.  Example: `/healthcheck` 
* `name` - A friendly name for the backend set. It must be unique and it cannot be changed.  Valid backend set names include only alphanumeric characters, dashes, and underscores. Backend set names cannot contain spaces. Avoid entering confidential information.  Example: `example_backend_set` 
* `policy` - The load balancer policy for the backend set. To get a list of available policies, use the [ListPolicies](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/LoadBalancerPolicy/ListPolicies) operation.  Example: `LEAST_CONNECTIONS` 
* `session_persistence_configuration` - 
	* `cookie_name` - The name of the cookie used to detect a session initiated by the backend server. Use '*' to specify that any cookie set by the backend causes the session to persist.  Example: `example_cookie` 
	* `disable_fallback` - Whether the load balancer is prevented from directing traffic from a persistent session client to a different backend server if the original server is unavailable. Defaults to false.  Example: `false` 
* `ssl_configuration` - 
	* `certificate_name` - A friendly name for the certificate bundle. It must be unique and it cannot be changed. Valid certificate bundle names include only alphanumeric characters, dashes, and underscores. Certificate bundle names cannot contain spaces. Avoid entering confidential information.  Example: `example_certificate_bundle` 
	* `verify_depth` - The maximum depth for peer certificate chain verification.  Example: `3` 
	* `verify_peer_certificate` - Whether the load balancer listener should verify peer certificates. Defaults to true.  Example: `true` 



### Create Operation
Adds a backend set to a load balancer.

The following arguments are supported:
 
* `health_checker` - (Required) 
	* `interval_ms` - (Optional) The interval between health checks, in milliseconds.  Example: `10000` 
	* `port` - (Optional) The backend server port against which to run the health check. If the port is not specified, the load balancer uses the port information from the `Backend` object.  Example: `8080` 
	* `protocol` - (Required) The protocol the health check must use; either HTTP or TCP.  Example: `HTTP` 
	* `response_body_regex` - (Optional) A regular expression for parsing the response body from the backend server.  Example: `^((?!false).|\s)*$` 
	* `retries` - (Optional) The number of retries to attempt before a backend server is considered "unhealthy".  Example: `3` 
	* `return_code` - (Optional) The status code a healthy backend server should return.  Example: `200` 
	* `timeout_in_millis` - (Optional) The maximum time, in milliseconds, to wait for a reply to a health check. A health check is successful only if a reply returns within this timeout period.  Example: `3000` 
	* `url_path` - (Required) The path against which to run the health check.  Example: `/healthcheck` 
* `load_balancer_id` - (Required) The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the load balancer on which to add a backend set.
* `name` - (Required) A friendly name for the backend set. It must be unique and it cannot be changed.  Valid backend set names include only alphanumeric characters, dashes, and underscores. Backend set names cannot contain spaces. Avoid entering confidential information.  Example: `example_backend_set` 
* `policy` - (Required) The load balancer policy for the backend set. To get a list of available policies, use the [ListPolicies](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/LoadBalancerPolicy/ListPolicies) operation.  Example: `LEAST_CONNECTIONS` 
* `session_persistence_configuration` - (Optional) 
	* `cookie_name` - (Required) The name of the cookie used to detect a session initiated by the backend server. Use '*' to specify that any cookie set by the backend causes the session to persist.  Example: `example_cookie` 
	* `disable_fallback` - (Optional) Whether the load balancer is prevented from directing traffic from a persistent session client to a different backend server if the original server is unavailable. Defaults to false.  Example: `false` 
* `ssl_configuration` - (Optional) 
	* `certificate_name` - (Required) A friendly name for the certificate bundle. It must be unique and it cannot be changed. Valid certificate bundle names include only alphanumeric characters, dashes, and underscores. Certificate bundle names cannot contain spaces. Avoid entering confidential information.  Example: `example_certificate_bundle` 
	* `verify_depth` - (Optional) The maximum depth for peer certificate chain verification.  Example: `3` 
	* `verify_peer_certificate` - (Optional) Whether the load balancer listener should verify peer certificates.  Example: `true` 


### Update Operation
Updates a backend set.

The following arguments support updates:
* `health_checker` - 
	* `interval_ms` - The interval between health checks, in milliseconds.  Example: `10000` 
	* `port` - The backend server port against which to run the health check. If the port is not specified, the load balancer uses the port information from the `Backend` object.  Example: `8080` 
	* `protocol` - The protocol the health check must use; either HTTP or TCP.  Example: `HTTP` 
	* `response_body_regex` - A regular expression for parsing the response body from the backend server.  Example: `^((?!false).|\s)*$` 
	* `retries` - The number of retries to attempt before a backend server is considered "unhealthy".  Example: `3` 
	* `return_code` - The status code a healthy backend server should return.  Example: `200` 
	* `timeout_in_millis` - The maximum time, in milliseconds, to wait for a reply to a health check. A health check is successful only if a reply returns within this timeout period.  Example: `3000` 
	* `url_path` - The path against which to run the health check.  Example: `/healthcheck` 
* `policy` - The load balancer policy for the backend set. To get a list of available policies, use the [ListPolicies](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/LoadBalancerPolicy/ListPolicies) operation.  Example: `LEAST_CONNECTIONS` 
* `session_persistence_configuration` - 
	* `cookie_name` - The name of the cookie used to detect a session initiated by the backend server. Use '*' to specify that any cookie set by the backend causes the session to persist.  Example: `example_cookie` 
	* `disable_fallback` - Whether the load balancer is prevented from directing traffic from a persistent session client to a different backend server if the original server is unavailable. Defaults to false.  Example: `false` 
* `ssl_configuration` - 
	* `certificate_name` - A friendly name for the certificate bundle. It must be unique and it cannot be changed. Valid certificate bundle names include only alphanumeric characters, dashes, and underscores. Certificate bundle names cannot contain spaces. Avoid entering confidential information.  Example: `example_certificate_bundle` 
	* `verify_depth` - The maximum depth for peer certificate chain verification.  Example: `3` 
	* `verify_peer_certificate` - Whether the load balancer listener should verify peer certificates.  Example: `true` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```hcl
resource "oci_load_balancer_backend_set" "test_backend_set" {
	#Required
	health_checker {
		#Required
		protocol = "${var.backend_set_health_checker_protocol}"

		#Optional
		interval_ms = "${var.backend_set_health_checker_interval_ms}"
		port = "${var.backend_set_health_checker_port}"
		response_body_regex = "${var.backend_set_health_checker_response_body_regex}"
		retries = "${var.backend_set_health_checker_retries}"
		return_code = "${var.backend_set_health_checker_return_code}"
		timeout_in_millis = "${var.backend_set_health_checker_timeout_in_millis}"
		url_path = "${var.backend_set_health_checker_url_path}"
	}
	load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
	name = "${var.backend_set_name}"
	policy = "${var.backend_set_policy}"

	session_persistence_configuration {
		#Required
		cookie_name = "${var.backend_set_session_persistence_configuration_cookie_name}"

		#Optional
		disable_fallback = "${var.backend_set_session_persistence_configuration_disable_fallback}"
	}
	ssl_configuration {
		#Required
		certificate_name = "${var.backend_set_ssl_configuration_certificate_name}"

		#Optional
		verify_depth = "${var.backend_set_ssl_configuration_verify_depth}"
		verify_peer_certificate = "${var.backend_set_ssl_configuration_verify_peer_certificate}"
	}
}
```

# oci_load_balancer_backend_sets
`oci_load_balancer_backendsets` is the old name for `oci_load_balancer_backend_sets`. Both names are supported but `oci_load_balancer_backend_sets` is used in the docs.

## BackendSet DataSource

Gets a list of backend_sets.

### List Operation
Lists all backend sets associated with a given load balancer.
The following arguments are supported:

* `load_balancer_id` - (Required) The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the load balancer associated with the backend sets to retrieve.


The following attributes are exported:

* `backendsets` - The list of backendsets.

### Example Usage

```hcl
data "oci_load_balancer_backend_sets" "test_backend_sets" {
	#Required
	load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
}
```