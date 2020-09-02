---
subcategory: "Load Balancer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_load_balancer_listener"
sidebar_current: "docs-oci-resource-load_balancer-listener"
description: |-
  Provides the Listener resource in Oracle Cloud Infrastructure Load Balancer service
---

# oci_load_balancer_listener
This resource provides the Listener resource in Oracle Cloud Infrastructure Load Balancer service.

Adds a listener to a load balancer.

## Example Usage

```hcl
resource "oci_load_balancer_listener" "test_listener" {
	#Required
	default_backend_set_name = "${oci_load_balancer_backend_set.test_backend_set.name}"
	load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
	name = "${var.listener_name}"
	port = "${var.listener_port}"
	protocol = "${var.listener_protocol}"

	#Optional
	connection_configuration {
		#Required
		idle_timeout_in_seconds = "${var.listener_connection_configuration_idle_timeout_in_seconds}"

		#Optional
		backend_tcp_proxy_protocol_version = "${var.listener_connection_configuration_backend_tcp_proxy_protocol_version}"
	}
	hostname_names = ["${oci_load_balancer_hostname.test_hostname.name}"]
	path_route_set_name = "${oci_load_balancer_path_route_set.test_path_route_set.name}"
	rule_set_names = ["${oci_load_balancer_rule_set.test_rule_set.name}"]
	ssl_configuration {
		#Required
		certificate_name = "${oci_load_balancer_certificate.test_certificate.name}"

		#Optional
		verify_depth = "${var.listener_ssl_configuration_verify_depth}"
		verify_peer_certificate = "${var.listener_ssl_configuration_verify_peer_certificate}"
        protocols = ["TLSv1.1", "TLSv1.2"]
        cipher_suite_name = "${oci_load_balancer_ssl_cipher_suite.example_ssl_cipher_suite.name}"
        server_order_preference = ENABLED
	}
}
```

## Argument Reference

The following arguments are supported:

* `connection_configuration` - (Optional) (Updatable) 
	* `backend_tcp_proxy_protocol_version` - (Required when `protocol` = `TCP`) (Updatable) The backend TCP Proxy Protocol version.  Example: `1` 
	* `idle_timeout_in_seconds` - (Required) (Updatable) The maximum idle time, in seconds, allowed between two successive receive or two successive send operations between the client and backend servers. A send operation does not reset the timer for receive operations. A receive operation does not reset the timer for send operations.

		For more information, see [Connection Configuration](https://docs.cloud.oracle.com/iaas/Content/Balance/Reference/connectionreuse.htm#ConnectionConfiguration).

		Example: `1200` 
* `default_backend_set_name` - (Required) (Updatable) The name of the associated backend set.  Example: `example_backend_set` 
* `hostname_names` - (Optional) (Updatable) An array of hostname resource names.
* `load_balancer_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the load balancer on which to add a listener.
* `name` - (Required) A friendly name for the listener. It must be unique and it cannot be changed. Avoid entering confidential information.  Example: `example_listener` 
* `path_route_set_name` - (Optional) (Updatable) The name of the set of path-based routing rules, [PathRouteSet](https://docs.cloud.oracle.com/iaas/api/#/en/loadbalancer/20170115/PathRouteSet/), applied to this listener's traffic.  Example: `example_path_route_set` 
* `port` - (Required) (Updatable) The communication port for the listener.  Example: `80` 
* `protocol` - (Required) (Updatable) The protocol on which the listener accepts connection requests. To get a list of valid protocols, use the [ListProtocols](https://docs.cloud.oracle.com/iaas/api/#/en/loadbalancer/20170115/LoadBalancerProtocol/ListProtocols) operation.  Example: `HTTP` 
* `rule_set_names` - (Optional) (Updatable) The names of the [rule sets](https://docs.cloud.oracle.com/iaas/api/#/en/loadbalancer/20170115/RuleSet/) to apply to the listener.  Example: ["example_rule_set"] 
* `ssl_configuration` - (Optional) (Updatable) 
	* `certificate_name` - (Required) (Updatable) A friendly name for the certificate bundle. It must be unique and it cannot be changed. Valid certificate bundle names include only alphanumeric characters, dashes, and underscores. Certificate bundle names cannot contain spaces. Avoid entering confidential information.  Example: `example_certificate_bundle` 
	* `verify_depth` - (Optional) (Updatable) The maximum depth for peer certificate chain verification.  Example: `3` 
	* `verify_peer_certificate` - (Optional) (Updatable) Whether the load balancer listener should verify peer certificates.  Example: `true` 
	* `protocols` - (Optional) (Updatable) A list of SSL protocols the load balancer must support for HTTPS or SSL connections. The load balancer uses SSL protocols to establish a secure connection between a client and a server. A secure connection ensures that all data passed between the client and the server is private. The Load Balancing service supports the following protocols:  TLSv1  TLSv1.1  TLSv1.2  If this field is not specified, TLSv1.2 is the default.  Example: `["TLSv1.1", "TLSv1.2"]`
    * `cipher_suite_name` - (Optional) (Updatable) The name of the cipher suite to use for HTTPS or SSL connections. If this field is not specified, the default is `oci-default-ssl-cipher-suite-v1`. Example: `example_cipher_suite`
    * `server_order_preference` - (Optional) (Updatable) When this attribute is set to ENABLED, the system gives preference to the server ciphers over the client ciphers.
** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Import

Listeners can be imported using the `id`, e.g.

```
$ terraform import oci_load_balancer_listener.test_listener "loadBalancers/{loadBalancerId}/listeners/{listenerName}" 
```

