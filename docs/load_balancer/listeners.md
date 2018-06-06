# oci_load_balancer_listener

## Listener Resource

### Listener Reference

The following attributes are exported:

* `connection_configuration` -
	* `idle_timeout_in_seconds` - The maximum idle time, in seconds, allowed between two successive receive or two successive send operations between the client and backend servers. A send operation does not reset the timer for receive operations. A receive operation does not reset the timer for send operations.  The default values are:  *  300 seconds for TCP  *  60 seconds for HTTP and WebSocket protocols.  Note: The protocol is set at the listener.  Modify this parameter if the client or backend server stops transmitting data for more than the default time. Some examples include:  *  The client sends a database query to the backend server and the database takes over 300 seconds to execute.    Therefore, the backend server does not transmit any data within 300 seconds.  *  The client uploads data using the HTTP protocol. During the upload, the backend does not transmit any data    to the client for more than 60 seconds.  *  The client downloads data using the HTTP protocol.  After the initial request, it stops transmitting data to    the backend server for more than 60 seconds.  *  The client starts transmitting data after establishing a WebSocket connection, but the backend server does    not transmit data for more than 60 seconds.  *  The backend server starts transmitting data after establishing a WebSocket connection, but the client does    not transmit data for more than 60 seconds.  The maximum value is 7200 seconds. Contact My Oracle Support to file a service request if you want to increase this limit for your tenancy. For more information, see [Service Limits](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/servicelimits.htm).  Example: `1200`
* `default_backend_set_name` - The name of the associated backend set.  Example: `example_backend_set` 
* `hostname_names` - An array of hostname resource names.
* `load_balancer_id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the load balancer on which to add a listener.
* `name` - A friendly name for the listener. It must be unique and it cannot be changed. Avoid entering confidential information.  Example: `My listener` 
* `path_route_set_name` - The name of the set of path-based routing rules, [PathRouteSet](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/PathRouteSet/), applied to this listener's traffic.  Example: `path-route-set-001` 
* `port` - The communication port for the listener.  Example: `80` 
* `protocol` - The protocol on which the listener accepts connection requests. To get a list of valid protocols, use the [ListProtocols](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/LoadBalancerProtocol/ListProtocols) operation.  Example: `HTTP` 
* `ssl_configuration` - 
	* `certificate_name` - A friendly name for the certificate bundle. It must be unique and it cannot be changed. Valid certificate bundle names include only alphanumeric characters, dashes, and underscores. Certificate bundle names cannot contain spaces. Avoid entering confidential information.  Example: `example_certificate_bundle` 
	* `verify_depth` - The maximum depth for peer certificate chain verification.  Example: `3` 
	* `verify_peer_certificate` - Whether the load balancer listener should verify peer certificates.  Example: `true` 



### Create Operation
Adds a listener to a load balancer.

The following arguments are supported:

* `connection_configuration` - (Optional) 
	* `idle_timeout_in_seconds` - (Required) The maximum idle time, in seconds, allowed between two successive receive or two successive send operations between the client and backend servers. A send operation does not reset the timer for receive operations. A receive operation does not reset the timer for send operations.  For more information, see [Connection Configuration](https://docs.us-phoenix-1.oraclecloud.com/Content/Balance/Reference/connectionreuse.htm#ConnectionConfiguration).  Example: `1200` 
* `default_backend_set_name` - (Required) The name of the associated backend set.  Example: `example_backend_set` 
* `hostname_names` - (Optional) An array of hostname resource names.
* `load_balancer_id` - (Required) The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the load balancer on which to add a listener.
* `name` - (Required) A friendly name for the listener. It must be unique and it cannot be changed. Avoid entering confidential information.  Example: `example_listener` 
* `path_route_set_name` - (Optional) The name of the set of path-based routing rules, [PathRouteSet](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/PathRouteSet/), applied to this listener's traffic.  Example: `example_path_route_set` 
* `port` - (Required) The communication port for the listener.  Example: `80` 
* `protocol` - (Required) The protocol on which the listener accepts connection requests. To get a list of valid protocols, use the [ListProtocols](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/LoadBalancerProtocol/ListProtocols) operation.  Example: `HTTP` 
* `ssl_configuration` - (Optional) 
	* `certificate_name` - (Required) A friendly name for the certificate bundle. It must be unique and it cannot be changed. Valid certificate bundle names include only alphanumeric characters, dashes, and underscores. Certificate bundle names cannot contain spaces. Avoid entering confidential information.  Example: `example_certificate_bundle` 
	* `verify_depth` - (Optional) The maximum depth for peer certificate chain verification.  Example: `3` 
	* `verify_peer_certificate` - (Optional) Whether the load balancer listener should verify peer certificates.  Example: `true` 


### Update Operation
Updates a listener for a given load balancer.

The following arguments support updates:
* `connection_configuration` - 
	* `idle_timeout_in_seconds` - The maximum idle time, in seconds, allowed between two successive receive or two successive send operations between the client and backend servers. A send operation does not reset the timer for receive operations. A receive operation does not reset the timer for send operations.  For more information, see [Connection Configuration](https://docs.us-phoenix-1.oraclecloud.com/Content/Balance/Reference/connectionreuse.htm#ConnectionConfiguration).  Example: `1200` 
* `default_backend_set_name` - The name of the associated backend set.  Example: `example_backend_set` 
* `hostname_names` - An array of hostname resource names.
* `path_route_set_name` - The name of the set of path-based routing rules, [PathRouteSet](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/PathRouteSet/), applied to this listener's traffic.  Example: `example_path_route_set` 
* `port` - The communication port for the listener.  Example: `80` 
* `protocol` - The protocol on which the listener accepts connection requests. To get a list of valid protocols, use the [ListProtocols](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/LoadBalancerProtocol/ListProtocols) operation.  Example: `HTTP` 
* `ssl_configuration` - 
	* `certificate_name` - A friendly name for the certificate bundle. It must be unique and it cannot be changed. Valid certificate bundle names include only alphanumeric characters, dashes, and underscores. Certificate bundle names cannot contain spaces. Avoid entering confidential information.  Example: `example_certificate_bundle` 
	* `verify_depth` - The maximum depth for peer certificate chain verification.  Example: `3` 
	* `verify_peer_certificate` - Whether the load balancer listener should verify peer certificates.  Example: `true` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```hcl
resource "oci_load_balancer_listener" "test_listener" {
	#Required
	default_backend_set_name = "${var.listener_default_backend_set_name}"
	load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
	name = "${var.listener_name}"
	port = "${var.listener_port}"
	protocol = "${var.listener_protocol}"

	#Optional
  	connection_configuration {
		#Required
		idle_timeout_in_seconds = "${var.listener_connection_configuration_idle_timeout_in_seconds}"
	}
	hostname_names = hostname_names = ["${oci_load_balancer_hostname.test_hostname.name}"]
	path_route_set_name = "${var.listener_path_route_set_name}"
	ssl_configuration {
		#Required
		certificate_name = "${var.listener_ssl_configuration_certificate_name}"

		#Optional
		verify_depth = "${var.listener_ssl_configuration_verify_depth}"
		verify_peer_certificate = "${var.listener_ssl_configuration_verify_peer_certificate}"
	}
}
```

