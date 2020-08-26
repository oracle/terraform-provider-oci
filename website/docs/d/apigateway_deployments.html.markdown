---
subcategory: "API Gateway"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apigateway_deployments"
sidebar_current: "docs-oci-datasource-apigateway-deployments"
description: |-
  Provides the list of Deployments in Oracle Cloud Infrastructure API Gateway service
---

# Data Source: oci_apigateway_deployments
This data source provides the list of Deployments in Oracle Cloud Infrastructure API Gateway service.

Returns a list of deployments.


## Example Usage

```hcl
data "oci_apigateway_deployments" "test_deployments" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.deployment_display_name}"
	gateway_id = "${oci_apigateway_gateway.test_gateway.id}"
	state = "${var.deployment_state}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ocid of the compartment in which to list resources.
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable.  Example: `My new resource` 
* `gateway_id` - (Optional) Filter deployments by the gateway ocid.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state.  Example: `SUCCEEDED` 


## Attributes Reference

The following attributes are exported:

* `deployment_collection` - The list of deployment_collection.

### Deployment Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the resource is created. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `endpoint` - The endpoint to access this deployment on the gateway.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `gateway_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource. 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource. 
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in a Failed state. 
* `path_prefix` - A path on which to deploy all routes contained in the API deployment specification. For more information, see [Deploying an API on an API Gateway by Creating an API Deployment](https://docs.cloud.oracle.com/iaas/Content/APIGateway/Tasks/apigatewaycreatingdeployment.htm). 
* `specification` - 
	* `logging_policies` - 
		* `access_log` - 
			* `is_enabled` - Enables pushing of access logs to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket.

				Oracle recommends using the Oracle Cloud Infrastructure Logging service to enable, retrieve, and query access logs for an API Deployment. If there is an active log object for the API Deployment and its category is set to 'access' in Oracle Cloud Infrastructure Logging service, the logs will not be uploaded to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket.

				Please note that the functionality to push to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket has been deprecated and will be removed in the future. 
		* `execution_log` - 
			* `is_enabled` - Enables pushing of execution logs to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket.

				Oracle recommends using the Oracle Cloud Infrastructure Logging service to enable, retrieve, and query execution logs for an API Deployment. If there is an active log object for the API Deployment and its category is set to 'execution' in Oracle Cloud Infrastructure Logging service, the logs will not be uploaded to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket.

				Please note that the functionality to push to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket has been deprecated and will be removed in the future. 
			* `log_level` - Specifies the log level used to control logging output of execution logs. Enabling logging at a given level also enables logging at all higher levels. 
	* `request_policies` - 
		* `authentication` - 
			* `audiences` - The list of intended recipients for the token.
			* `function_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Functions function resource. 
			* `is_anonymous_access_allowed` - Whether an unauthenticated user may access the API. Must be "true" to enable ANONYMOUS route authorization. 
			* `issuers` - A list of parties that could have issued the token.
			* `max_clock_skew_in_seconds` - The maximum expected time difference between the system clocks of the token issuer and the API Gateway. 
			* `public_keys` - 
				* `is_ssl_verify_disabled` - Defines whether or not to uphold SSL verification. 
				* `keys` - The set of static public keys.
					* `alg` - The algorithm intended for use with this key.
					* `e` - The base64 url encoded exponent of the RSA public key represented by this key. 
					* `format` - The format of the public key.
					* `key` - The content of the PEM-encoded public key.
					* `key_ops` - The operations for which this key is to be used.
					* `kid` - A unique key ID. This key will be used to verify the signature of a JWT with matching "kid". 
					* `kty` - The key type.
					* `n` - The base64 url encoded modulus of the RSA public key represented by this key. 
					* `use` - The intended use of the public key.
				* `max_cache_duration_in_hours` - The duration for which the JWKS should be cached before it is fetched again. 
				* `type` - Type of the public key set.
				* `uri` - The uri from which to retrieve the key. It must be accessible without authentication. 
			* `token_auth_scheme` - The authentication scheme that is to be used when authenticating the token. This must to be provided if "tokenHeader" is specified. 
			* `token_header` - The name of the header containing the authentication token.
			* `token_query_param` - The name of the query parameter containing the authentication token.
			* `type` - Type of the authentication policy to use.
			* `verify_claims` - A list of claims which should be validated to consider the token valid.
				* `is_required` - Whether the claim is required to be present in the JWT or not. If set to "false", the claim values will be matched only if the claim is present in the JWT. 
				* `key` - Name of the claim.
				* `values` - The list of acceptable values for a given claim. If this value is "null" or empty and "isRequired" set to "true", then the presence of this claim in the JWT is validated. 
		* `cors` - 
			* `allowed_headers` - The list of headers that will be allowed from the client via the Access-Control-Allow-Headers header. '*' will allow all headers. 
			* `allowed_methods` - The list of allowed HTTP methods that will be returned for the preflight OPTIONS request in the Access-Control-Allow-Methods header. '*' will allow all methods. 
			* `allowed_origins` - The list of allowed origins that the CORS handler will use to respond to CORS requests. The gateway will send the Access-Control-Allow-Origin header with the best origin match for the circumstances. '*' will match any origins, and 'null' will match queries from 'file:' origins. All other origins must be qualified with the scheme, full hostname, and port if necessary. 
			* `exposed_headers` - The list of headers that the client will be allowed to see from the response as indicated by the Access-Control-Expose-Headers header. '*' will expose all headers. 
			* `is_allow_credentials_enabled` - Whether to send the Access-Control-Allow-Credentials header to allow CORS requests with cookies. 
			* `max_age_in_seconds` - The time in seconds for the client to cache preflight responses. This is sent as the Access-Control-Max-Age if greater than 0. 
		* `rate_limiting` - 
			* `rate_in_requests_per_second` - The maximum number of requests per second to allow.
			* `rate_key` - The key used to group requests together.
	* `routes` - A list of routes that this API exposes.
		* `backend` - 
			* `body` - The body of the stock response from the mock backend.
			* `connect_timeout_in_seconds` - Defines a timeout for establishing a connection with a proxied server. 
			* `function_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Functions function resource. 
			* `headers` - The headers of the stock response from the mock backend.
				* `name` - Name of the header.
				* `value` - Value of the header.
			* `is_ssl_verify_disabled` - Defines whether or not to uphold SSL verification. 
			* `read_timeout_in_seconds` - Defines a timeout for reading a response from the proxied server. 
			* `send_timeout_in_seconds` - Defines a timeout for transmitting a request to the proxied server. 
			* `status` - The status code of the stock response from the mock backend.
			* `type` - Type of the API backend.
			* `url` - 
		* `logging_policies` - 
			* `access_log` - 
				* `is_enabled` - Enables pushing of access logs to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket.

					Oracle recommends using the Oracle Cloud Infrastructure Logging service to enable, retrieve, and query access logs for an API Deployment. If there is an active log object for the API Deployment and its category is set to 'access' in Oracle Cloud Infrastructure Logging service, the logs will not be uploaded to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket.

					Please note that the functionality to push to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket has been deprecated and will be removed in the future. 
			* `execution_log` - 
				* `is_enabled` - Enables pushing of execution logs to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket.

					Oracle recommends using the Oracle Cloud Infrastructure Logging service to enable, retrieve, and query execution logs for an API Deployment. If there is an active log object for the API Deployment and its category is set to 'execution' in Oracle Cloud Infrastructure Logging service, the logs will not be uploaded to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket.

					Please note that the functionality to push to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket has been deprecated and will be removed in the future. 
				* `log_level` - Specifies the log level used to control logging output of execution logs. Enabling logging at a given level also enables logging at all higher levels. 
		* `methods` - A list of allowed methods on this route. 
		* `path` - A URL path pattern that must be matched on this route. The path pattern may contain a subset of RFC 6570 identifiers to allow wildcard and parameterized matching. 
		* `request_policies` - 
			* `authorization` - 
				* `allowed_scope` - A user whose scope includes any of these access ranges is allowed on this route. Access ranges are case-sensitive. 
				* `type` - Indicates how authorization should be applied. For a type of ANY_OF, an "allowedScope" property must also be specified. Otherwise, only a type is required. For a type of ANONYMOUS, an authenticated API must have the "isAnonymousAccessAllowed" property set to "true" in the authentication policy. 
			* `cors` - 
				* `allowed_headers` - The list of headers that will be allowed from the client via the Access-Control-Allow-Headers header. '*' will allow all headers. 
				* `allowed_methods` - The list of allowed HTTP methods that will be returned for the preflight OPTIONS request in the Access-Control-Allow-Methods header. '*' will allow all methods. 
				* `allowed_origins` - The list of allowed origins that the CORS handler will use to respond to CORS requests. The gateway will send the Access-Control-Allow-Origin header with the best origin match for the circumstances. '*' will match any origins, and 'null' will match queries from 'file:' origins. All other origins must be qualified with the scheme, full hostname, and port if necessary. 
				* `exposed_headers` - The list of headers that the client will be allowed to see from the response as indicated by the Access-Control-Expose-Headers header. '*' will expose all headers. 
				* `is_allow_credentials_enabled` - Whether to send the Access-Control-Allow-Credentials header to allow CORS requests with cookies. 
				* `max_age_in_seconds` - The time in seconds for the client to cache preflight responses. This is sent as the Access-Control-Max-Age if greater than 0. 
* `state` - The current state of the deployment.
* `time_created` - The time this resource was created. An RFC3339 formatted datetime string.
* `time_updated` - The time this resource was last updated. An RFC3339 formatted datetime string.

