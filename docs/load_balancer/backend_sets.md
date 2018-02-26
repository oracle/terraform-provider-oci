# oci\_load_balancer\_backend_set

## BackendSet Resource

### BackendSet Reference

The following attributes are exported:

* `backend` - 
	* `backup` - Whether the load balancer should treat this server as a backup unit. If `true`, the load balancer forwards no ingress traffic to this backend server unless all other backend servers not marked as "backup" fail the health check policy.  Example: `true` 
	* `drain` - Whether the load balancer should drain this server. Servers marked "drain" receive no new incoming traffic.  Example: `true` 
	* `ip_address` - The IP address of the backend server.  Example: `10.10.10.4` 
	* `name` - A read-only field showing the IP address and port that uniquely identify this backend server in the backend set.  Example: `10.10.10.4:8080` 
	* `offline` - Whether the load balancer should treat this server as offline. Offline servers receive no incoming traffic.  Example: `true` 
	* `port` - The communication port for the backend server.  Example: `8080` 
	* `weight` - The load balancing policy weight assigned to the server. Backend servers with a higher weight receive a larger proportion of incoming traffic. For example, a server weighted '3' receives 3 times the number of new connections as a server weighted '1'. For more information on load balancing policies, see [How Load Balancing Policies Work](https://docs.us-phoenix-1.oraclecloud.com/Content/Balance/Reference/lbpolicies.htm).  Example: `3` 
* `health_checker` - 
	* `interval_ms` - The interval between health checks, in milliseconds. The default is 30000 (30 seconds).  Example: `30000` 
	* `port` - The backend server port against which to run the health check. If the port is not specified, the load balancer uses the port information from the `Backend` object.  Example: `8080` 
	* `protocol` - The protocol the health check must use; either HTTP or TCP.  Example: `HTTP` 
	* `response_body_regex` - A regular expression for parsing the response body from the backend server.  Example: `^(500|40[1348])$` 
	* `retries` - The number of retries to attempt before a backend server is considered "unhealthy". Defaults to 3.  Example: `3` 
	* `return_code` - The status code a healthy backend server should return. If you configure the health check policy to use the HTTP protocol, you can use common HTTP status codes such as "200".  Example: `200` 
	* `timeout_in_millis` - The maximum time, in milliseconds, to wait for a reply to a health check. A health check is successful only if a reply returns within this timeout period. Defaults to 3000 (3 seconds).  Example: `6000` 
	* `url_path` - The path against which to run the health check.  Example: `/healthcheck` 
* `name` - A friendly name for the backend set. It must be unique and it cannot be changed.  Valid backend set names include only alphanumeric characters, dashes, and underscores. Backend set names cannot contain spaces. Avoid entering confidential information.  Example: `My_backend_set` 
* `policy` - The load balancer policy for the backend set. To get a list of available policies, use the [ListPolicies](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/LoadBalancerPolicy/ListPolicies) operation.  Example: `LEAST_CONNECTIONS` 
* `session_persistence_configuration` - 
	* `cookie_name` - The name of the cookie used to detect a session initiated by the backend server. Use '*' to specify that any cookie set by the backend causes the session to persist.  Example: `myCookieName` 
	* `disable_fallback` - Whether the load balancer is prevented from directing traffic from a persistent session client to a different backend server if the original server is unavailable. Defaults to false.  Example: `true` 
* `ssl_configuration` - 
	* `certificate_name` - A friendly name for the certificate bundle. It must be unique and it cannot be changed. Valid certificate bundle names include only alphanumeric characters, dashes, and underscores. Certificate bundle names cannot contain spaces. Avoid entering confidential information.  Example: `My_certificate_bundle` 
	* `verify_depth` - The maximum depth for peer certificate chain verification.  Example: `3` 
	* `verify_peer_certificate` - Whether the load balancer listener should verify peer certificates. Defaults to true.  Example: `true` 



### Create Operation
Adds a backend set to a load balancer.

The following arguments are supported:

* `backend` - (Optional) 
	* `backup` - (Optional) Whether the load balancer should treat this server as a backup unit. If `true`, the load balancer forwards no ingress traffic to this backend server unless all other backend servers not marked as "backup" fail the health check policy.  Example: `true` 
	* `drain` - (Optional) Whether the load balancer should drain this server. Servers marked "drain" receive no new incoming traffic.  Example: `true` 
	* `ip_address` - (Required) The IP address of the backend server.  Example: `10.10.10.4` 
	* `offline` - (Optional) Whether the load balancer should treat this server as offline. Offline servers receive no incoming traffic.  Example: `true` 
	* `port` - (Required) The communication port for the backend server.  Example: `8080` 
	* `weight` - (Optional) The load balancing policy weight assigned to the server. Backend servers with a higher weight receive a larger proportion of incoming traffic. For example, a server weighted '3' receives 3 times the number of new connections as a server weighted '1'. For more information on load balancing policies, see [How Load Balancing Policies Work](https://docs.us-phoenix-1.oraclecloud.com/Content/Balance/Reference/lbpolicies.htm).  Example: `3` 
* `health_checker` - (Required) 
	* `interval_ms` - (Optional) The interval between health checks, in milliseconds.  Example: `30000` 
	* `port` - (Optional) The backend server port against which to run the health check. If the port is not specified, the load balancer uses the port information from the `Backend` object.  Example: `8080` 
	* `protocol` - (Required) The protocol the health check must use; either HTTP or TCP.  Example: `HTTP` 
	* `response_body_regex` - (Optional) A regular expression for parsing the response body from the backend server.  Example: `^(500|40[1348])$` 
	* `retries` - (Optional) The number of retries to attempt before a backend server is considered "unhealthy".  Example: `3` 
	* `return_code` - (Optional) The status code a healthy backend server should return.  Example: `200` 
	* `timeout_in_millis` - (Optional) The maximum time, in milliseconds, to wait for a reply to a health check. A health check is successful only if a reply returns within this timeout period.  Example: `6000` 
	* `url_path` - (Optional) The path against which to run the health check.  Example: `/healthcheck` 
* `load_balancer_id` - (Required) The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the load balancer on which to add a backend set.
* `name` - (Required) A friendly name for the backend set. It must be unique and it cannot be changed.  Valid backend set names include only alphanumeric characters, dashes, and underscores. Backend set names cannot contain spaces. Avoid entering confidential information.  Example: `My_backend_set` 
* `policy` - (Required) The load balancer policy for the backend set. To get a list of available policies, use the [ListPolicies](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/LoadBalancerPolicy/ListPolicies) operation.  Example: `LEAST_CONNECTIONS` 
* `session_persistence_configuration` - (Optional) 
	* `cookie_name` - (Required) The name of the cookie used to detect a session initiated by the backend server. Use '*' to specify that any cookie set by the backend causes the session to persist.  Example: `myCookieName` 
	* `disable_fallback` - (Optional) Whether the load balancer is prevented from directing traffic from a persistent session client to a different backend server if the original server is unavailable. Defaults to false.  Example: `true` 
* `ssl_configuration` - (Optional) 
	* `certificate_name` - (Required) A friendly name for the certificate bundle. It must be unique and it cannot be changed. Valid certificate bundle names include only alphanumeric characters, dashes, and underscores. Certificate bundle names cannot contain spaces. Avoid entering confidential information.  Example: `My_certificate_bundle` 
	* `verify_depth` - (Optional) The maximum depth for peer certificate chain verification.  Example: `3` 
	* `verify_peer_certificate` - (Optional) Whether the load balancer listener should verify peer certificates.  Example: `true` 


### Update Operation
Updates a backend set.

The following arguments support updates:
	* `backup` - Whether the load balancer should treat this server as a backup unit. If `true`, the load balancer forwards no ingress traffic to this backend server unless all other backend servers not marked as "backup" fail the health check policy.  Example: `true` 
	* `drain` - Whether the load balancer should drain this server. Servers marked "drain" receive no new incoming traffic.  Example: `true` 
	* `offline` - Whether the load balancer should treat this server as offline. Offline servers receive no incoming traffic.  Example: `true` 
	* `weight` - The load balancing policy weight assigned to the server. Backend servers with a higher weight receive a larger proportion of incoming traffic. For example, a server weighted '3' receives 3 times the number of new connections as a server weighted '1'. For more information on load balancing policies, see [How Load Balancing Policies Work](https://docs.us-phoenix-1.oraclecloud.com/Content/Balance/Reference/lbpolicies.htm).  Example: `3` 
	* `interval_ms` - The interval between health checks, in milliseconds.  Example: `30000` 
	* `port` - The backend server port against which to run the health check. If the port is not specified, the load balancer uses the port information from the `Backend` object.  Example: `8080` 
	* `response_body_regex` - A regular expression for parsing the response body from the backend server.  Example: `^(500|40[1348])$` 
	* `retries` - The number of retries to attempt before a backend server is considered "unhealthy".  Example: `3` 
	* `return_code` - The status code a healthy backend server should return.  Example: `200` 
	* `timeout_in_millis` - The maximum time, in milliseconds, to wait for a reply to a health check. A health check is successful only if a reply returns within this timeout period.  Example: `6000` 
	* `url_path` - The path against which to run the health check.  Example: `/healthcheck` 
* `session_persistence_configuration` - 
	* `disable_fallback` - Whether the load balancer is prevented from directing traffic from a persistent session client to a different backend server if the original server is unavailable. Defaults to false.  Example: `true` 
* `ssl_configuration` - 
	* `verify_depth` - The maximum depth for peer certificate chain verification.  Example: `3` 
	* `verify_peer_certificate` - Whether the load balancer listener should verify peer certificates.  Example: `true` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```
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

	#Optional
	backend {
		#Required
		ip_address = "${var.backend_set_backend_ip_address}"
		port = "${var.backend_set_backend_port}"

		#Optional
		backup = "${var.backend_set_backend_backup}"
		drain = "${var.backend_set_backend_drain}"
		offline = "${var.backend_set_backend_offline}"
		weight = "${var.backend_set_backend_weight}"
	}
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

# oci\_load_balancer\_backend_sets

## BackendSet DataSource

Gets a list of backend_sets.

### List Operation
Lists all backend sets associated with a given load balancer.
The following arguments are supported:

* `load_balancer_id` - (Required) The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the load balancer associated with the backend sets to retrieve.


The following attributes are exported:

* `backend_sets` - The list of backend_sets.

### Example Usage

```
data "oci_load_balancer_backend_sets" "test_backend_sets" {
	#Required
	load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
}
```