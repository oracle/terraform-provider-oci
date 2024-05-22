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
	default_backend_set_name = oci_load_balancer_backend_set.test_backend_set.name
	load_balancer_id = oci_load_balancer_load_balancer.test_load_balancer.id
	name = var.listener_name
	port = var.listener_port
	protocol = var.listener_protocol

	#Optional
	connection_configuration {
		#Required
		idle_timeout_in_seconds = var.listener_connection_configuration_idle_timeout_in_seconds

		#Optional
		backend_tcp_proxy_protocol_version = var.listener_connection_configuration_backend_tcp_proxy_protocol_version
	}
	hostname_names = [oci_load_balancer_hostname.test_hostname.name]
	path_route_set_name = oci_load_balancer_path_route_set.test_path_route_set.name
	routing_policy_name = oci_load_balancer_load_balancer_routing_policy.test_load_balancer_routing_policy.name
	rule_set_names = [oci_load_balancer_rule_set.test_rule_set.name]
	ssl_configuration {
        #Optional
		certificate_name = oci_load_balancer_certificate.test_certificate.name
		has_session_resumption = var.listener_ssl_configuration_has_session_resumption
		certificate_ids = var.listener_ssl_configuration_certificate_ids
		cipher_suite_name = var.listener_ssl_configuration_cipher_suite_name
		protocols = var.listener_ssl_configuration_protocols
		server_order_preference = var.listener_ssl_configuration_server_order_preference
		trusted_certificate_authority_ids = var.listener_ssl_configuration_trusted_certificate_authority_ids
		verify_depth = var.listener_ssl_configuration_verify_depth
		verify_peer_certificate = var.listener_ssl_configuration_verify_peer_certificate
	}
}
```

## Argument Reference

The following arguments are supported:

* `connection_configuration` - (Optional) (Updatable) Configuration details for the connection between the client and backend servers. 
	* `backend_tcp_proxy_protocol_version` - (Required when `protocol` = `TCP`) (Updatable) The backend TCP Proxy Protocol version.  Example: `1` 
	* `idle_timeout_in_seconds` - (Required) (Updatable) The maximum idle time, in seconds, allowed between two successive receive or two successive send operations between the client and backend servers. A send operation does not reset the timer for receive operations. A receive operation does not reset the timer for send operations.

		For more information, see [Connection Configuration](https://docs.cloud.oracle.com/iaas/Content/Balance/Reference/connectionreuse.htm#ConnectionConfiguration).

		Example: `1200` 
* `default_backend_set_name` - (Required) (Updatable) The name of the associated backend set.  Example: `example_backend_set` 
* `hostname_names` - (Optional) (Updatable) An array of hostname resource names.
* `load_balancer_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the load balancer on which to add a listener.
* `name` - (Required) A friendly name for the listener. It must be unique and it cannot be changed. Avoid entering confidential information.  Example: `example_listener` 
* `path_route_set_name` - (Optional) (Updatable) Deprecated. Please use `routingPolicies` instead.

	The name of the set of path-based routing rules, [PathRouteSet](https://docs.cloud.oracle.com/iaas/api/#/en/loadbalancer/20170115/PathRouteSet/), applied to this listener's traffic.

	Example: `example_path_route_set` 
* `port` - (Required) (Updatable) The communication port for the listener.  Example: `80` 
* `protocol` - (Required) (Updatable) The protocol on which the listener accepts connection requests. To get a list of valid protocols, use the [ListProtocols](https://docs.cloud.oracle.com/iaas/api/#/en/loadbalancer/20170115/LoadBalancerProtocol/ListProtocols) operation.  Example: `HTTP` 
* `routing_policy_name` - (Optional) (Updatable) The name of the routing policy applied to this listener's traffic.  Example: `example_routing_policy` 
* `rule_set_names` - (Optional) (Updatable) The names of the [rule sets](https://docs.cloud.oracle.com/iaas/api/#/en/loadbalancer/20170115/RuleSet/) to apply to the listener.  Example: ["example_rule_set"] 
* `ssl_configuration` - (Optional) (Updatable) The load balancer's SSL handling configuration details.

	**Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API. 
	* `certificate_ids` - (Optional) (Updatable) Ids for Oracle Cloud Infrastructure certificates service certificates. Currently only a single Id may be passed.  Example: `[ocid1.certificate.oc1.us-ashburn-1.amaaaaaaav3bgsaa5o2q7rh5nfmkkukfkogasqhk6af2opufhjlqg7m6jqzq]` 
	* `certificate_name` - (Optional) (Updatable) A friendly name for the certificate bundle. It must be unique and it cannot be changed. Valid certificate bundle names include only alphanumeric characters, dashes, and underscores. Certificate bundle names cannot contain spaces. Avoid entering confidential information.  Example: `example_certificate_bundle` 
	* `has_session_resumption` - (Optional) (Updatable) Whether the load balancer listener should resume an encrypted session by reusing the cryptographic parameters of a previous TLS session, without having to perform a full handshake again. If "true", the service resumes the previous TLS encrypted session. If "false", the service starts a new TLS encrypted session. Enabling session resumption improves performance but provides a lower level of security. Disabling session resumption improves security but reduces performance.  Example: `true` 
	* `cipher_suite_name` - (Optional) (Updatable) The name of the cipher suite to use for HTTPS or SSL connections.

		If this field is not specified, the default is `oci-default-ssl-cipher-suite-v1`.

		**Notes:**
		*  You must ensure compatibility between the specified SSL protocols and the ciphers configured in the cipher suite. Clients cannot perform an SSL handshake if there is an incompatible configuration.
		*  You must ensure compatibility between the ciphers configured in the cipher suite and the configured certificates. For example, RSA-based ciphers require RSA certificates and ECDSA-based ciphers require ECDSA certificates.
		*  If the cipher configuration is not modified after load balancer creation, the `GET` operation returns `oci-default-ssl-cipher-suite-v1` as the value of this field in the SSL configuration for existing listeners that predate this feature.
		*  If the cipher configuration was modified using Oracle operations after load balancer creation, the `GET` operation returns `oci-customized-ssl-cipher-suite` as the value of this field in the SSL configuration for existing listeners that predate this feature.
		*  The `GET` operation returns `oci-wider-compatible-ssl-cipher-suite-v1` as the value of this field in the SSL configuration for existing backend sets that predate this feature.
		*  If the `GET` operation on a listener returns `oci-customized-ssl-cipher-suite` as the value of this field, you must specify an appropriate predefined or custom cipher suite name when updating the resource.
		*  The `oci-customized-ssl-cipher-suite` Oracle reserved cipher suite name is not accepted as valid input for this field.

		example: `example_cipher_suite` 
	* `protocols` - (Optional) (Updatable) A list of SSL protocols the load balancer must support for HTTPS or SSL connections.

		The load balancer uses SSL protocols to establish a secure connection between a client and a server. A secure connection ensures that all data passed between the client and the server is private.

		The Load Balancing service supports the following protocols:
		*  TLSv1
		*  TLSv1.1
		*  TLSv1.2
        *  TLSv1.3

		If this field is not specified, TLSv1.2 is the default.

		**Warning:** All SSL listeners created on a given port must use the same set of SSL protocols.

		**Notes:**
		*  The handshake to establish an SSL connection fails if the client supports none of the specified protocols.
		*  You must ensure compatibility between the specified SSL protocols and the ciphers configured in the cipher suite.
		*  For all existing load balancer listeners and backend sets that predate this feature, the `GET` operation displays a list of SSL protocols currently used by those resources.

		example: `["TLSv1.1", "TLSv1.2"]` 
	* `server_order_preference` - (Optional) (Updatable) When this attribute is set to ENABLED, the system gives preference to the server ciphers over the client ciphers.

		**Note:** This configuration is applicable only when the load balancer is acting as an SSL/HTTPS server. This field is ignored when the `SSLConfiguration` object is associated with a backend set. 
	* `trusted_certificate_authority_ids` - (Optional) (Updatable) Ids for Oracle Cloud Infrastructure certificates service CA or CA bundles for the load balancer to trust.  Example: `[ocid1.cabundle.oc1.us-ashburn-1.amaaaaaaav3bgsaagl4zzyqdop5i2vuwoqewdvauuw34llqa74otq2jdsfyq]` 
	* `verify_depth` - (Optional) (Updatable) The maximum depth for peer certificate chain verification.  Example: `3` 
	* `verify_peer_certificate` - (Optional) (Updatable) Whether the load balancer listener should verify peer certificates.  Example: `true` 


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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Listener
	* `update` - (Defaults to 20 minutes), when updating the Listener
	* `delete` - (Defaults to 20 minutes), when destroying the Listener


## Import

Listeners can be imported using the `id`, e.g.

```
$ terraform import oci_load_balancer_listener.test_listener "loadBalancers/{loadBalancerId}/listeners/{listenerName}" 
```

