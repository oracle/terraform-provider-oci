---
subcategory: "API Gateway"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apigateway_deployment"
sidebar_current: "docs-oci-resource-apigateway-deployment"
description: |-
  Provides the Deployment resource in Oracle Cloud Infrastructure API Gateway service
---

# oci_apigateway_deployment
This resource provides the Deployment resource in Oracle Cloud Infrastructure API Gateway service.

Creates a new deployment.


## Example Usage

```hcl
resource "oci_apigateway_deployment" "test_deployment" {
	#Required
	compartment_id = "${var.compartment_id}"
	gateway_id = "${oci_apigateway_gateway.test_gateway.id}"
	path_prefix = "${var.deployment_path_prefix}"
	specification {

		#Optional
		logging_policies {

			#Optional
			access_log {

				#Optional
				is_enabled = "${var.deployment_specification_logging_policies_access_log_is_enabled}"
			}
			execution_log {

				#Optional
				is_enabled = "${var.deployment_specification_logging_policies_execution_log_is_enabled}"
				log_level = "${var.deployment_specification_logging_policies_execution_log_log_level}"
			}
		}
		request_policies {

			#Optional
			authentication {
				#Required
				type = "${var.deployment_specification_request_policies_authentication_type}"

				#Optional
				audiences = "${var.deployment_specification_request_policies_authentication_audiences}"
				function_id = "${oci_functions_function.test_function.id}"
				is_anonymous_access_allowed = "${var.deployment_specification_request_policies_authentication_is_anonymous_access_allowed}"
				issuers = "${var.deployment_specification_request_policies_authentication_issuers}"
				max_clock_skew_in_seconds = "${var.deployment_specification_request_policies_authentication_max_clock_skew_in_seconds}"
				public_keys {
					#Required
					type = "${var.deployment_specification_request_policies_authentication_public_keys_type}"

					#Optional
					is_ssl_verify_disabled = "${var.deployment_specification_request_policies_authentication_public_keys_is_ssl_verify_disabled}"
					keys {
						#Required
						format = "${var.deployment_specification_request_policies_authentication_public_keys_keys_format}"

						#Optional
						alg = "${var.deployment_specification_request_policies_authentication_public_keys_keys_alg}"
						e = "${var.deployment_specification_request_policies_authentication_public_keys_keys_e}"
						key = "${var.deployment_specification_request_policies_authentication_public_keys_keys_key}"
						key_ops = "${var.deployment_specification_request_policies_authentication_public_keys_keys_key_ops}"
						kid = "${var.deployment_specification_request_policies_authentication_public_keys_keys_kid}"
						kty = "${var.deployment_specification_request_policies_authentication_public_keys_keys_kty}"
						n = "${var.deployment_specification_request_policies_authentication_public_keys_keys_n}"
						use = "${var.deployment_specification_request_policies_authentication_public_keys_keys_use}"
					}
					max_cache_duration_in_hours = "${var.deployment_specification_request_policies_authentication_public_keys_max_cache_duration_in_hours}"
					uri = "${var.deployment_specification_request_policies_authentication_public_keys_uri}"
				}
				token_auth_scheme = "${var.deployment_specification_request_policies_authentication_token_auth_scheme}"
				token_header = "${var.deployment_specification_request_policies_authentication_token_header}"
				token_query_param = "${var.deployment_specification_request_policies_authentication_token_query_param}"
				verify_claims {

					#Optional
					is_required = "${var.deployment_specification_request_policies_authentication_verify_claims_is_required}"
					key = "${var.deployment_specification_request_policies_authentication_verify_claims_key}"
					values = "${var.deployment_specification_request_policies_authentication_verify_claims_values}"
				}
			}
			cors {
				#Required
				allowed_origins = "${var.deployment_specification_request_policies_cors_allowed_origins}"

				#Optional
				allowed_headers = "${var.deployment_specification_request_policies_cors_allowed_headers}"
				allowed_methods = "${var.deployment_specification_request_policies_cors_allowed_methods}"
				exposed_headers = "${var.deployment_specification_request_policies_cors_exposed_headers}"
				is_allow_credentials_enabled = "${var.deployment_specification_request_policies_cors_is_allow_credentials_enabled}"
				max_age_in_seconds = "${var.deployment_specification_request_policies_cors_max_age_in_seconds}"
			}
			rate_limiting {
				#Required
				rate_in_requests_per_second = "${var.deployment_specification_request_policies_rate_limiting_rate_in_requests_per_second}"
				rate_key = "${var.deployment_specification_request_policies_rate_limiting_rate_key}"
			}
		}
		routes {
			#Required
			backend {
				#Required
				type = "${var.deployment_specification_routes_backend_type}"

				#Optional
				body = "${var.deployment_specification_routes_backend_body}"
				connect_timeout_in_seconds = "${var.deployment_specification_routes_backend_connect_timeout_in_seconds}"
				function_id = "${oci_functions_function.test_function.id}"
				headers {

					#Optional
					name = "${var.deployment_specification_routes_backend_headers_name}"
					value = "${var.deployment_specification_routes_backend_headers_value}"
				}
				is_ssl_verify_disabled = "${var.deployment_specification_routes_backend_is_ssl_verify_disabled}"
				read_timeout_in_seconds = "${var.deployment_specification_routes_backend_read_timeout_in_seconds}"
				send_timeout_in_seconds = "${var.deployment_specification_routes_backend_send_timeout_in_seconds}"
				status = "${var.deployment_specification_routes_backend_status}"
				url = "${var.deployment_specification_routes_backend_url}"
			}
			path = "${var.deployment_specification_routes_path}"

			#Optional
			logging_policies {

				#Optional
				access_log {

					#Optional
					is_enabled = "${var.deployment_specification_routes_logging_policies_access_log_is_enabled}"
				}
				execution_log {

					#Optional
					is_enabled = "${var.deployment_specification_routes_logging_policies_execution_log_is_enabled}"
					log_level = "${var.deployment_specification_routes_logging_policies_execution_log_log_level}"
				}
			}
			methods = "${var.deployment_specification_routes_methods}"
			request_policies {

				#Optional
				authorization {

					#Optional
					allowed_scope = "${var.deployment_specification_routes_request_policies_authorization_allowed_scope}"
					type = "${var.deployment_specification_routes_request_policies_authorization_type}"
				}
				cors {
					#Required
					allowed_origins = "${var.deployment_specification_routes_request_policies_cors_allowed_origins}"

					#Optional
					allowed_headers = "${var.deployment_specification_routes_request_policies_cors_allowed_headers}"
					allowed_methods = "${var.deployment_specification_routes_request_policies_cors_allowed_methods}"
					exposed_headers = "${var.deployment_specification_routes_request_policies_cors_exposed_headers}"
					is_allow_credentials_enabled = "${var.deployment_specification_routes_request_policies_cors_is_allow_credentials_enabled}"
					max_age_in_seconds = "${var.deployment_specification_routes_request_policies_cors_max_age_in_seconds}"
				}
			}
		}
	}

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = "${var.deployment_display_name}"
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the resource is created. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable.  Avoid entering confidential information.  Example: `My new resource` 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `gateway_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource. 
* `path_prefix` - (Required) A path on which to deploy all routes contained in the API deployment specification. For more information, see [Deploying an API on an API Gateway by Creating an API  Deployment](https://docs.cloud.oracle.com/iaas/Content/APIGateway/Tasks/apigatewaycreatingdeployment.htm). 
* `specification` - (Required) (Updatable) 
	* `logging_policies` - (Optional) (Updatable) 
		* `access_log` - (Optional) (Updatable) 
			* `is_enabled` - (Optional) (Updatable) Enables pushing of access logs to Oracle Cloud Infrastructure Public Logging.
		* `execution_log` - (Optional) (Updatable) 
			* `is_enabled` - (Optional) (Updatable) Enables pushing of execution logs to Oracle Cloud Infrastructure Public Logging.
			* `log_level` - (Optional) (Updatable) Specifies the logging level, which affects the log entries pushed to Oracle Cloud Infrastructure Public Logging if `isEnabled` is set to True. 
	* `request_policies` - (Optional) (Updatable) 
		* `authentication` - (Optional) (Updatable) 
			* `audiences` - (Required when type=JWT_AUTHENTICATION) (Updatable) The list of intended recipients for the token.
			* `function_id` - (Required when type=CUSTOM_AUTHENTICATION) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Functions function resource. 
			* `is_anonymous_access_allowed` - (Optional) (Updatable) Whether an unauthenticated user may access the API. Must be "true" to enable ANONYMOUS route authorization. 
			* `issuers` - (Required when type=JWT_AUTHENTICATION) (Updatable) A list of parties that could have issued the token.
			* `max_clock_skew_in_seconds` - (Applicable when type=JWT_AUTHENTICATION) (Updatable) The maximum expected time difference between the system clocks of the token issuer and the API Gateway. 
			* `public_keys` - (Required when type=JWT_AUTHENTICATION) (Updatable) 
				* `is_ssl_verify_disabled` - (Applicable when type=REMOTE_JWKS) (Updatable) Defines whether or not to uphold SSL verification. 
				* `keys` - (Applicable when type=STATIC_KEYS) (Updatable) The set of static public keys.
					* `alg` - (Required when format=JSON_WEB_KEY) (Updatable) The algorithm intended for use with this key.
					* `e` - (Required when format=JSON_WEB_KEY) (Updatable) The base64 url encoded exponent of the RSA public key represented by this key. 
					* `format` - (Required) (Updatable) The format of the public key.
					* `key` - (Required when format=PEM) (Updatable) The content of the PEM-encoded public key.
					* `key_ops` - (Applicable when format=JSON_WEB_KEY) (Updatable) The operations for which this key is to be used.
					* `kid` - (Required when type=STATIC_KEYS) (Updatable) A unique key ID. This key will be used to verify the signature of a JWT with matching "kid". 
					* `kty` - (Required when format=JSON_WEB_KEY) (Updatable) The key type.
					* `n` - (Required when format=JSON_WEB_KEY) (Updatable) The base64 url encoded modulus of the RSA public key represented by this key. 
					* `use` - (Applicable when format=JSON_WEB_KEY) (Updatable) The intended use of the public key.
				* `max_cache_duration_in_hours` - (Applicable when type=REMOTE_JWKS) (Updatable) The duration for which the JWKS should be cached before it is fetched again. 
				* `type` - (Required) (Updatable) Type of the public key set.
				* `uri` - (Required when type=REMOTE_JWKS) (Updatable) The uri from which to retrieve the key. It must be accessible without authentication. 
			* `token_auth_scheme` - (Applicable when type=JWT_AUTHENTICATION) (Updatable) The authentication scheme that is to be used when authenticating the token. This must to be provided if "tokenHeader" is specified. 
			* `token_header` - (Optional) (Updatable) The name of the header containing the authentication token.
			* `token_query_param` - (Optional) (Updatable) The name of the query parameter containing the authentication token.
			* `type` - (Required) (Updatable) Type of the authentication policy to use.
			* `verify_claims` - (Applicable when type=JWT_AUTHENTICATION) (Updatable) A list of claims which should be validated to consider the token valid.
				* `is_required` - (Applicable when type=JWT_AUTHENTICATION) (Updatable) Whether the claim is required to be present in the JWT or not. If set to "false", the claim values will be matched only if the claim is present in the JWT. 
				* `key` - (Required when type=JWT_AUTHENTICATION) (Updatable) Name of the claim.
				* `values` - (Applicable when type=JWT_AUTHENTICATION) (Updatable) The list of acceptable values for a given claim. If this value is "null" or empty and "isRequired" set to "true", then the presence of this claim in the JWT is validated. 
		* `cors` - (Optional) (Updatable) 
			* `allowed_headers` - (Optional) (Updatable) The list of headers that will be allowed from the client via the Access-Control-Allow-Headers header. '*' will allow all headers. 
			* `allowed_methods` - (Optional) (Updatable) The list of allowed HTTP methods that will be returned for the preflight OPTIONS request in the Access-Control-Allow-Methods header. '*' will allow all methods. 
			* `allowed_origins` - (Required) (Updatable) The list of allowed origins that the CORS handler will use to respond to CORS requests. The gateway will send the Access-Control-Allow-Origin header with the best origin match for the circumstances. '*' will match any origins, and 'null' will match queries from 'file:' origins. All other origins must be qualified with the scheme, full hostname, and port if necessary. 
			* `exposed_headers` - (Optional) (Updatable) The list of headers that the client will be allowed to see from the response as indicated by the Access-Control-Expose-Headers header. '*' will expose all headers. 
			* `is_allow_credentials_enabled` - (Optional) (Updatable) Whether to send the Access-Control-Allow-Credentials header to allow CORS requests with cookies. 
			* `max_age_in_seconds` - (Optional) (Updatable) The time in seconds for the client to cache preflight responses. This is sent as the Access-Control-Max-Age if greater than 0. 
		* `rate_limiting` - (Optional) (Updatable) 
			* `rate_in_requests_per_second` - (Required) (Updatable) The maximum number of requests per second to allow.
			* `rate_key` - (Required) (Updatable) The key used to group requests together.
	* `routes` - (Required) (Updatable) A list of routes that this API exposes.
		* `backend` - (Required) (Updatable) 
			* `body` - (Applicable when type=STOCK_RESPONSE_BACKEND) (Updatable) The body of the stock response from the mock backend.
			* `connect_timeout_in_seconds` - (Applicable when type=HTTP_BACKEND) (Updatable) Defines a timeout for establishing a connection with a proxied server. 
			* `function_id` - (Required when type=ORACLE_FUNCTIONS_BACKEND) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Functions function resource. 
			* `headers` - (Applicable when type=STOCK_RESPONSE_BACKEND) (Updatable) The headers of the stock response from the mock backend.
				* `name` - (Applicable when type=STOCK_RESPONSE_BACKEND) (Updatable) Name of the header.
				* `value` - (Applicable when type=STOCK_RESPONSE_BACKEND) (Updatable) Value of the header.
			* `is_ssl_verify_disabled` - (Applicable when type=HTTP_BACKEND) (Updatable) Defines whether or not to uphold SSL verification. 
			* `read_timeout_in_seconds` - (Applicable when type=HTTP_BACKEND) (Updatable) Defines a timeout for reading a response from the proxied server. 
			* `send_timeout_in_seconds` - (Applicable when type=HTTP_BACKEND) (Updatable) Defines a timeout for transmitting a request to the proxied server. 
			* `status` - (Required when type=STOCK_RESPONSE_BACKEND) (Updatable) The status code of the stock response from the mock backend.
			* `type` - (Required) (Updatable) Type of the API backend.
			* `url` - (Required when type=HTTP_BACKEND) (Updatable) 
		* `logging_policies` - (Optional) (Updatable) 
			* `access_log` - (Optional) (Updatable) 
				* `is_enabled` - (Optional) (Updatable) Enables pushing of access logs to Oracle Cloud Infrastructure Public Logging.
			* `execution_log` - (Optional) (Updatable) 
				* `is_enabled` - (Optional) (Updatable) Enables pushing of execution logs to Oracle Cloud Infrastructure Public Logging.
				* `log_level` - (Optional) (Updatable) Specifies the logging level, which affects the log entries pushed to Oracle Cloud Infrastructure Public Logging if `isEnabled` is set to True. 
		* `methods` - (Optional) (Updatable) A list of allowed methods on this route. 
		* `path` - (Required) (Updatable) A URL path pattern that must be matched on this route. The path pattern may contain a subset of RFC 6570 identifiers to allow wildcard and parameterized matching. 
		* `request_policies` - (Optional) (Updatable) 
			* `authorization` - (Optional) (Updatable) 
				* `allowed_scope` - (Required when type=ANY_OF) (Updatable) A user whose scope includes any of these access ranges is allowed on this route. Access ranges are case-sensitive. 
				* `type` - (Optional) (Updatable) Indicates how authorization should be applied. For a type of ANY_OF, an "allowedScope" property must also be specfied. Otherwise, only a type is required. For a type of ANONYMOUS, an authenticated API must have the "isAnonymousAccessAllowed" property set to "true" in the authentication policy. 
			* `cors` - (Optional) (Updatable) 
				* `allowed_headers` - (Optional) (Updatable) The list of headers that will be allowed from the client via the Access-Control-Allow-Headers header. '*' will allow all headers. 
				* `allowed_methods` - (Optional) (Updatable) The list of allowed HTTP methods that will be returned for the preflight OPTIONS request in the Access-Control-Allow-Methods header. '*' will allow all methods. 
				* `allowed_origins` - (Required) (Updatable) The list of allowed origins that the CORS handler will use to respond to CORS requests. The gateway will send the Access-Control-Allow-Origin header with the best origin match for the circumstances. '*' will match any origins, and 'null' will match queries from 'file:' origins. All other origins must be qualified with the scheme, full hostname, and port if necessary. 
				* `exposed_headers` - (Optional) (Updatable) The list of headers that the client will be allowed to see from the response as indicated by the Access-Control-Expose-Headers header. '*' will expose all headers. 
				* `is_allow_credentials_enabled` - (Optional) (Updatable) Whether to send the Access-Control-Allow-Credentials header to allow CORS requests with cookies. 
				* `max_age_in_seconds` - (Optional) (Updatable) The time in seconds for the client to cache preflight responses. This is sent as the Access-Control-Max-Age if greater than 0. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the resource is created. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.  Avoid entering confidential information.  Example: `My new resource` 
* `endpoint` - The endpoint to access this deployment on the gateway.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `gateway_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource. 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource. 
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in a Failed state. 
* `path_prefix` - A path on which to deploy all routes contained in the API deployment specification. For more information, see [Deploying an API on an API Gateway by Creating an API  Deployment](https://docs.cloud.oracle.com/iaas/Content/APIGateway/Tasks/apigatewaycreatingdeployment.htm). 
* `specification` - 
	* `logging_policies` - 
		* `access_log` - 
			* `is_enabled` - Enables pushing of access logs to Oracle Cloud Infrastructure Public Logging.
		* `execution_log` - 
			* `is_enabled` - Enables pushing of execution logs to Oracle Cloud Infrastructure Public Logging.
			* `log_level` - Specifies the logging level, which affects the log entries pushed to Oracle Cloud Infrastructure Public Logging if `isEnabled` is set to True. 
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
				* `is_enabled` - Enables pushing of access logs to Oracle Cloud Infrastructure Public Logging.
			* `execution_log` - 
				* `is_enabled` - Enables pushing of execution logs to Oracle Cloud Infrastructure Public Logging.
				* `log_level` - Specifies the logging level, which affects the log entries pushed to Oracle Cloud Infrastructure Public Logging if `isEnabled` is set to True. 
		* `methods` - A list of allowed methods on this route. 
		* `path` - A URL path pattern that must be matched on this route. The path pattern may contain a subset of RFC 6570 identifiers to allow wildcard and parameterized matching. 
		* `request_policies` - 
			* `authorization` - 
				* `allowed_scope` - A user whose scope includes any of these access ranges is allowed on this route. Access ranges are case-sensitive. 
				* `type` - Indicates how authorization should be applied. For a type of ANY_OF, an "allowedScope" property must also be specfied. Otherwise, only a type is required. For a type of ANONYMOUS, an authenticated API must have the "isAnonymousAccessAllowed" property set to "true" in the authentication policy. 
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

## Import

Deployments can be imported using the `id`, e.g.

```
$ terraform import oci_apigateway_deployment.test_deployment "id"
```

