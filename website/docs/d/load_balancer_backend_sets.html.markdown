---
subcategory: "Load Balancer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_load_balancer_backend_sets"
sidebar_current: "docs-oci-datasource-load_balancer-backend_sets"
description: |-
  Provides the list of Backend Sets in Oracle Cloud Infrastructure Load Balancer service
---

# Data Source: oci_load_balancer_backend_sets
This data source provides the list of Backend Sets in Oracle Cloud Infrastructure Load Balancer service.

Lists all backend sets associated with a given load balancer.

## Supported Aliases

* `oci_load_balancer_backendsets`

## Example Usage

```hcl
data "oci_load_balancer_backend_sets" "test_backend_sets" {
	#Required
	load_balancer_id = oci_load_balancer_load_balancer.test_load_balancer.id
}
```

## Argument Reference

The following arguments are supported:

* `load_balancer_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the load balancer associated with the backend sets to retrieve.


## Attributes Reference

The following attributes are exported:

* `backendsets` - The list of backendsets.

### BackendSet Reference

The following attributes are exported:

* `backend` - 
	* `backup` - Whether the load balancer should treat this server as a backup unit. If `true`, the load balancer forwards no ingress traffic to this backend server unless all other backend servers not marked as "backup" fail the health check policy.

		**Note:** You cannot add a backend server marked as `backup` to a backend set that uses the IP Hash policy.

		Example: `false` 
	* `drain` - Whether the load balancer should drain this server. Servers marked "drain" receive no new incoming traffic.  Example: `false` 
	* `ip_address` - The IP address of the backend server.  Example: `10.0.0.3` 
	* `max_connections` - The maximum number of simultaneous connections the load balancer can make to the backend. If this is not set then the maximum number of simultaneous connections the load balancer can make to the backend is unlimited.  Example: `300` 
	* `name` - A read-only field showing the IP address and port that uniquely identify this backend server in the backend set.  Example: `10.0.0.3:8080` 
	* `offline` - Whether the load balancer should treat this server as offline. Offline servers receive no incoming traffic.  Example: `false` 
	* `port` - The communication port for the backend server.  Example: `8080` 
	* `weight` - The load balancing policy weight assigned to the server. Backend servers with a higher weight receive a larger proportion of incoming traffic. For example, a server weighted '3' receives 3 times the number of new connections as a server weighted '1'. For more information on load balancing policies, see [How Load Balancing Policies Work](https://docs.cloud.oracle.com/iaas/Content/Balance/Reference/lbpolicies.htm).  Example: `3` 
* `backend_max_connections` - The maximum number of simultaneous connections the load balancer can make to any backend in the backend set unless the backend has its own maxConnections setting. If this is not set then the number of simultaneous connections the load balancer can make to any backend in the backend set unless the backend has its own maxConnections setting is unlimited.  Example: `300` 
* `health_checker` - The health check policy configuration. For more information, see [Editing Health Check Policies](https://docs.cloud.oracle.com/iaas/Content/Balance/Tasks/editinghealthcheck.htm). 
	* `interval_ms` - The interval between health checks, in milliseconds. The default is 10000 (10 seconds).  Example: `10000` 
	* `is_force_plain_text` - Specifies if health checks should always be done using plain text instead of depending on whether or not the associated backend set is using SSL.

		If "true", health checks will be done using plain text even if the associated backend set is configured to use SSL.

		If "false", health checks will be done using SSL encryption if the associated backend set is configured to use SSL. If the backend set is not so configured the health checks will be done using plain text.

		Example: `false`
	* `port` - The backend server port against which to run the health check. If the port is not specified, the load balancer uses the port information from the `Backend` object.  Example: `8080` 
	* `protocol` - The protocol the health check must use; either HTTP or TCP.  Example: `HTTP` 
	* `response_body_regex` - A regular expression for parsing the response body from the backend server.  Example: `^((?!false).|\s)*$` 
	* `retries` - The number of retries to attempt before a backend server is considered "unhealthy". This number also applies when recovering a server to the "healthy" state. Defaults to 3.  Example: `3` 
	* `return_code` - The status code a healthy backend server should return. If you configure the health check policy to use the HTTP protocol, you can use common HTTP status codes such as "200".  Example: `200` 
	* `timeout_in_millis` - The maximum time, in milliseconds, to wait for a reply to a health check. A health check is successful only if a reply returns within this timeout period. Defaults to 3000 (3 seconds).  Example: `3000` 
	* `url_path` - The path against which to run the health check.  Example: `/healthcheck` 
* `lb_cookie_session_persistence_configuration` - The configuration details for implementing load balancer cookie session persistence (LB cookie stickiness).

	Session persistence enables the Load Balancing service to direct all requests that originate from a single logical client to a single backend web server. For more information, see [Session Persistence](https://docs.cloud.oracle.com/iaas/Content/Balance/Reference/sessionpersistence.htm).

	When you configure LB cookie stickiness, the load balancer inserts a cookie into the response. The parameters configured in the cookie enable session stickiness. This method is useful when you have applications and Web backend services that cannot generate their own cookies.

	Path route rules take precedence to determine the target backend server. The load balancer verifies that session stickiness is enabled for the backend server and that the cookie configuration (domain, path, and cookie hash) is valid for the target. The system ignores invalid cookies.

	To disable LB cookie stickiness on a running load balancer, use the [UpdateBackendSet](https://docs.cloud.oracle.com/iaas/api/#/en/loadbalancer/20170115/BackendSet/UpdateBackendSet) operation and specify `null` for the `LBCookieSessionPersistenceConfigurationDetails` object.

	Example: `LBCookieSessionPersistenceConfigurationDetails: null`

	**Note:** `SessionPersistenceConfigurationDetails` (application cookie stickiness) and `LBCookieSessionPersistenceConfigurationDetails` (LB cookie stickiness) are mutually exclusive. An error results if you try to enable both types of session persistence.

	**Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API. 
	* `cookie_name` - The name of the cookie inserted by the load balancer. If this field is not configured, the cookie name defaults to "X-Oracle-BMC-LBS-Route".  Example: `example_cookie`

		**Notes:**
		*  Ensure that the cookie name used at the backend application servers is different from the cookie name used at the load balancer. To minimize the chance of name collision, Oracle recommends that you use a prefix such as "X-Oracle-OCI-" for this field.
		*  If a backend server and the load balancer both insert cookies with the same name, the client or browser behavior can vary depending on the domain and path values associated with the cookie. If the name, domain, and path values of the `Set-cookie` generated by a backend server and the `Set-cookie` generated by the load balancer are all the same, the client or browser treats them as one cookie and returns only one of the cookie values in subsequent requests. If both `Set-cookie` names are the same, but the domain and path names are different, the client or browser treats them as two different cookies. 
	* `disable_fallback` - Whether the load balancer is prevented from directing traffic from a persistent session client to a different backend server if the original server is unavailable. Defaults to false.  Example: `false` 
	* `domain` - The domain in which the cookie is valid. The `Set-cookie` header inserted by the load balancer contains a domain attribute with the specified value.

		This attribute has no default value. If you do not specify a value, the load balancer does not insert the domain attribute into the `Set-cookie` header.

		**Notes:**
		*  [RFC 6265 - HTTP State Management Mechanism](https://www.ietf.org/rfc/rfc6265.txt) describes client and browser behavior when the domain attribute is present or not present in the `Set-cookie` header.

		If the value of the `Domain` attribute is `example.com` in the `Set-cookie` header, the client includes the same cookie in the `Cookie` header when making HTTP requests to `example.com`, `www.example.com`, and `www.abc.example.com`. If the `Domain` attribute is not present, the client returns the cookie only for the domain to which the original request was made.
		*  Ensure that this attribute specifies the correct domain value. If the `Domain` attribute in the `Set-cookie` header does not include the domain to which the original request was made, the client or browser might reject the cookie. As specified in RFC 6265, the client accepts a cookie with the `Domain` attribute value `example.com` or `www.example.com` sent from `www.example.com`. It does not accept a cookie with the `Domain` attribute `abc.example.com` or `www.abc.example.com` sent from `www.example.com`.

		Example: `example.com` 
	* `is_http_only` - Whether the `Set-cookie` header should contain the `HttpOnly` attribute. If `true`, the `Set-cookie` header inserted by the load balancer contains the `HttpOnly` attribute, which limits the scope of the cookie to HTTP requests. This attribute directs the client or browser to omit the cookie when providing access to cookies through non-HTTP APIs. For example, it restricts the cookie from JavaScript channels.  Example: `true` 
	* `is_secure` - Whether the `Set-cookie` header should contain the `Secure` attribute. If `true`, the `Set-cookie` header inserted by the load balancer contains the `Secure` attribute, which directs the client or browser to send the cookie only using a secure protocol.

		**Note:** If you set this field to `true`, you cannot associate the corresponding backend set with an HTTP listener.

		Example: `true` 
	* `max_age_in_seconds` - The amount of time the cookie remains valid. The `Set-cookie` header inserted by the load balancer contains a `Max-Age` attribute with the specified value.

		The specified value must be at least one second. There is no default value for this attribute. If you do not specify a value, the load balancer does not include the `Max-Age` attribute in the `Set-cookie` header. In most cases, the client or browser retains the cookie until the current session ends, as defined by the client.

		Example: `3600` 
	* `path` - The path in which the cookie is valid. The `Set-cookie header` inserted by the load balancer contains a `Path` attribute with the specified value.

		Clients include the cookie in an HTTP request only if the path portion of the request-uri matches, or is a subdirectory of, the cookie's `Path` attribute.

		The default value is `/`.

		Example: `/example` 
* `name` - A friendly name for the backend set. It must be unique and it cannot be changed.

	Valid backend set names include only alphanumeric characters, dashes, and underscores. Backend set names cannot contain spaces. Avoid entering confidential information.

	Example: `example_backend_set` 
* `policy` - The load balancer policy for the backend set. To get a list of available policies, use the [ListPolicies](https://docs.cloud.oracle.com/iaas/api/#/en/loadbalancer/20170115/LoadBalancerPolicy/ListPolicies) operation.  Example: `LEAST_CONNECTIONS` 
* `session_persistence_configuration` - The configuration details for implementing session persistence based on a user-specified cookie name (application cookie stickiness).

	Session persistence enables the Load Balancing service to direct any number of requests that originate from a single logical client to a single backend web server. For more information, see [Session Persistence](https://docs.cloud.oracle.com/iaas/Content/Balance/Reference/sessionpersistence.htm).

	With application cookie stickiness, the load balancer enables session persistence only when the response from a backend application server includes a `Set-cookie` header with the user-specified cookie name.

	To disable application cookie stickiness on a running load balancer, use the [UpdateBackendSet](https://docs.cloud.oracle.com/iaas/api/#/en/loadbalancer/20170115/BackendSet/UpdateBackendSet) operation and specify `null` for the `SessionPersistenceConfigurationDetails` object.

	Example: `SessionPersistenceConfigurationDetails: null`

	**Note:** `SessionPersistenceConfigurationDetails` (application cookie stickiness) and `LBCookieSessionPersistenceConfigurationDetails` (LB cookie stickiness) are mutually exclusive. An error results if you try to enable both types of session persistence.

	**Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API. 
	* `cookie_name` - The name of the cookie used to detect a session initiated by the backend server. Use '*' to specify that any cookie set by the backend causes the session to persist.  Example: `example_cookie` 
	* `disable_fallback` - Whether the load balancer is prevented from directing traffic from a persistent session client to a different backend server if the original server is unavailable. Defaults to false.  Example: `false` 
* `ssl_configuration` - A listener's SSL handling configuration.

	To use SSL, a listener must be associated with a [certificate bundle](https://docs.cloud.oracle.com/iaas/api/#/en/loadbalancer/20170115/Certificate/).

	**Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API. 
	* `certificate_ids` - Ids for Oracle Cloud Infrastructure certificates service certificates. Currently only a single Id may be passed.  Example: `[ocid1.certificate.oc1.us-ashburn-1.amaaaaaaav3bgsaa5o2q7rh5nfmkkukfkogasqhk6af2opufhjlqg7m6jqzq]` 
	* `certificate_name` - A friendly name for the certificate bundle. It must be unique and it cannot be changed. Valid certificate bundle names include only alphanumeric characters, dashes, and underscores. Certificate bundle names cannot contain spaces. Avoid entering confidential information.  Example: `example_certificate_bundle` 
	* `cipher_suite_name` - The name of the cipher suite to use for HTTPS or SSL connections.

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
	* `protocols` - A list of SSL protocols the load balancer must support for HTTPS or SSL connections.

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
	* `server_order_preference` - When this attribute is set to ENABLED, the system gives preference to the server ciphers over the client ciphers.

		**Note:** This configuration is applicable only when the load balancer is acting as an SSL/HTTPS server. This field is ignored when the `SSLConfiguration` object is associated with a backend set. 
	* `trusted_certificate_authority_ids` - Ids for Oracle Cloud Infrastructure certificates service CA or CA bundles for the load balancer to trust.  Example: `[ocid1.cabundle.oc1.us-ashburn-1.amaaaaaaav3bgsaagl4zzyqdop5i2vuwoqewdvauuw34llqa74otq2jdsfyq]` 
	* `verify_depth` - The maximum depth for peer certificate chain verification.  Example: `3` 
	* `verify_peer_certificate` - Whether the load balancer listener should verify peer certificates.  Example: `true` 

