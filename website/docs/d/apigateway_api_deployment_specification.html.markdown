---
subcategory: "API Gateway"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apigateway_api_deployment_specification"
sidebar_current: "docs-oci-datasource-apigateway-api_deployment_specification"
description: |-
  Provides details about a specific Api Deployment Specification in Oracle Cloud Infrastructure API Gateway service
---

# Data Source: oci_apigateway_api_deployment_specification
This data source provides details about a specific Api Deployment Specification resource in Oracle Cloud Infrastructure API Gateway service.

Gets an API Deployment specification by identifier.

## Example Usage

```hcl
data "oci_apigateway_api_deployment_specification" "test_api_deployment_specification" {
	#Required
	api_id = oci_apigateway_api.test_api.id
}
```

## Argument Reference

The following arguments are supported:

* `api_id` - (Required) The ocid of the API.


## Attributes Reference

The following attributes are exported:

* `logging_policies` - Policies controlling the pushing of logs to Oracle Cloud Infrastructure Public Logging. 
	* `access_log` - Configures the logging policies for the access logs of an API Deployment. 
		* `is_enabled` - Enables pushing of access logs to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket.

			Oracle recommends using the Oracle Cloud Infrastructure Logging service to enable, retrieve, and query access logs for an API Deployment. If there is an active log object for the API Deployment and its category is set to 'access' in Oracle Cloud Infrastructure Logging service, the logs will not be uploaded to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket.

			Please note that the functionality to push to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket has been deprecated and will be removed in the future. 
	* `execution_log` - Configures the logging policies for the execution logs of an API Deployment. 
		* `is_enabled` - Enables pushing of execution logs to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket.

			Oracle recommends using the Oracle Cloud Infrastructure Logging service to enable, retrieve, and query execution logs for an API Deployment. If there is an active log object for the API Deployment and its category is set to 'execution' in Oracle Cloud Infrastructure Logging service, the logs will not be uploaded to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket.

			Please note that the functionality to push to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket has been deprecated and will be removed in the future. 
		* `log_level` - Specifies the log level used to control logging output of execution logs. Enabling logging at a given level also enables logging at all higher levels. 
* `request_policies` - Global behavior applied to all requests received by the API.
	* `authentication` - Information on how to authenticate incoming requests.
		* `audiences` - The list of intended recipients for the token.
		* `cache_key` - A list of keys from "parameters" attribute value whose values will be added to the cache key. 
		* `function_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Functions function resource. 
		* `is_anonymous_access_allowed` - Whether an unauthenticated user may access the API. Must be "true" to enable ANONYMOUS route authorization. 
		* `issuers` - A list of parties that could have issued the token.
		* `max_clock_skew_in_seconds` - The maximum expected time difference between the system clocks of the token issuer and the API Gateway. 
		* `parameters` - A map where key is a user defined string and value is a context expressions whose values will be sent to the custom auth function. Values should contain an expression. Example: `{"foo": "request.header[abc]"}` 
		* `public_keys` - A set of Public Keys that will be used to verify the JWT signature.
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
		* `validation_failure_policy` - Policy for defining behaviour on validation failure.
			* `client_details` - Client App Credential details.
				* `client_id` - Client ID for the OAuth2/OIDC app.
				* `client_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Vault Service secret resource. 
				* `client_secret_version_number` - The version number of the client secret to use.
				* `type` - To specify where the Client App details should be taken from.
			* `fallback_redirect_path` - The path to be used as fallback after OAuth2.
			* `logout_path` - The path to be used as logout.
			* `max_expiry_duration_in_hours` - The duration for which the OAuth2 success token should be cached before it is fetched again. 
			* `response_code` - HTTP response code, can include context variables.
			* `response_header_transformations` - A set of transformations to apply to HTTP headers that pass through the gateway. 
				* `filter_headers` - Filter HTTP headers as they pass through the gateway.  The gateway applies filters after other transformations, so any headers set or renamed must also be listed here when using an ALLOW type policy. 
					* `items` - The list of headers. 
						* `name` - The case-insensitive name of the header.  This name must be unique across transformation policies. 
					* `type` - BLOCK drops any headers that are in the list of items, so it acts as an exclusion list.  ALLOW permits only the headers in the list and removes all others, so it acts as an inclusion list. 
				* `rename_headers` - Rename HTTP headers as they pass through the gateway. 
					* `items` - The list of headers.
						* `from` - The original case-insensitive name of the header.  This name must be unique across transformation policies. 
						* `to` - The new name of the header.  This name must be unique across transformation policies. 
				* `set_headers` - Set HTTP headers as they pass through the gateway. 
					* `items` - The list of headers.
						* `if_exists` - If a header with the same name already exists in the request, OVERWRITE will overwrite the value, APPEND will append to the existing value, or SKIP will keep the existing value. 
						* `name` - The case-insensitive name of the header.  This name must be unique across transformation policies. 
						* `values` - A list of new values.  Each value can be a constant or may include one or more expressions enclosed within ${} delimiters. 
			* `response_message` - HTTP response message.
			* `response_type` - Response Type.
			* `scopes` - List of scopes.
			* `source_uri_details` - Auth endpoint details.
				* `type` - Type of the Uri detail.
				* `uri` - The discovery URI for the auth server.
			* `type` - Type of the Validation failure Policy.
			* `use_cookies_for_intermediate_steps` - Defines whether or not to use cookies for OAuth2 intermediate steps. 
			* `use_cookies_for_session` - Defines whether or not to use cookies for session maintenance. 
			* `use_pkce` - Defines whether or not to support PKCE. 
		* `validation_policy` - Authentication Policies for the Token Authentication types.
			* `additional_validation_policy` - Additional JWT validation checks.
				* `audiences` - The list of intended recipients for the token.
				* `issuers` - A list of parties that could have issued the token.
				* `verify_claims` - A list of claims which should be validated to consider the token valid.
					* `is_required` - Whether the claim is required to be present in the JWT or not. If set to "false", the claim values will be matched only if the claim is present in the JWT. 
					* `key` - Name of the claim.
					* `values` - The list of acceptable values for a given claim. If this value is "null" or empty and "isRequired" set to "true", then the presence of this claim in the JWT is validated. 
			* `client_details` - Client App Credential details.
				* `client_id` - Client ID for the OAuth2/OIDC app.
				* `client_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Vault Service secret resource. 
				* `client_secret_version_number` - The version number of the client secret to use.
				* `type` - To specify where the Client App details should be taken from.
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
			* `max_cache_duration_in_hours` - The duration for which the introspect URL response should be cached before it is fetched again. 
			* `source_uri_details` - Auth endpoint details.
				* `type` - Type of the Uri detail.
				* `uri` - The discovery URI for the auth server.
			* `type` - Type of the token validation policy.
			* `uri` - The uri from which to retrieve the key. It must be accessible without authentication. 
		* `verify_claims` - A list of claims which should be validated to consider the token valid.
			* `is_required` - Whether the claim is required to be present in the JWT or not. If set to "false", the claim values will be matched only if the claim is present in the JWT. 
			* `key` - Name of the claim.
			* `values` - The list of acceptable values for a given claim. If this value is "null" or empty and "isRequired" set to "true", then the presence of this claim in the JWT is validated. 
	* `cors` - Enable CORS (Cross-Origin-Resource-Sharing) request handling. 
		* `allowed_headers` - The list of headers that will be allowed from the client via the Access-Control-Allow-Headers header. '*' will allow all headers. 
		* `allowed_methods` - The list of allowed HTTP methods that will be returned for the preflight OPTIONS request in the Access-Control-Allow-Methods header. '*' will allow all methods. 
		* `allowed_origins` - The list of allowed origins that the CORS handler will use to respond to CORS requests. The gateway will send the Access-Control-Allow-Origin header with the best origin match for the circumstances. '*' will match any origins, and 'null' will match queries from 'file:' origins. All other origins must be qualified with the scheme, full hostname, and port if necessary. 
		* `exposed_headers` - The list of headers that the client will be allowed to see from the response as indicated by the Access-Control-Expose-Headers header. '*' will expose all headers. 
		* `is_allow_credentials_enabled` - Whether to send the Access-Control-Allow-Credentials header to allow CORS requests with cookies. 
		* `max_age_in_seconds` - The time in seconds for the client to cache preflight responses. This is sent as the Access-Control-Max-Age if greater than 0. 
	* `dynamic_authentication` - Policy on how to authenticate requests when multiple authentication options are configured for a deployment. For an incoming request, the value of selector specified under selectionSource will be matched against the keys specified for each authentication server. The authentication server whose key matches the value of selector will be used for authentication.
		* `authentication_servers` - List of authentication servers to choose from during dynamic authentication.
			* `authentication_server_detail` - Information on how to authenticate incoming requests.
				* `audiences` - The list of intended recipients for the token.
				* `function_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Functions function resource. 
				* `is_anonymous_access_allowed` - Whether an unauthenticated user may access the API. Must be "true" to enable ANONYMOUS route authorization. 
				* `issuers` - A list of parties that could have issued the token.
				* `max_clock_skew_in_seconds` - The maximum expected time difference between the system clocks of the token issuer and the API Gateway. 
				* `public_keys` - A set of Public Keys that will be used to verify the JWT signature.
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
				* `validation_failure_policy` - Policy for defining behaviour on validation failure.
					* `client_details` - Client App Credential details.
						* `client_id` - Client ID for the OAuth2/OIDC app.
						* `client_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Vault Service secret resource. 
						* `client_secret_version_number` - The version number of the client secret to use.
						* `type` - To specify where the Client App details should be taken from.
					* `fallback_redirect_path` - The path to be used as fallback after OAuth2.
					* `logout_path` - The path to be used as logout.
					* `max_expiry_duration_in_hours` - The duration for which the OAuth2 success token should be cached before it is fetched again. 
					* `response_code` - HTTP response code, can include context variables.
					* `response_header_transformations` - A set of transformations to apply to HTTP headers that pass through the gateway. 
						* `filter_headers` - Filter HTTP headers as they pass through the gateway.  The gateway applies filters after other transformations, so any headers set or renamed must also be listed here when using an ALLOW type policy. 
							* `items` - The list of headers. 
								* `name` - The case-insensitive name of the header.  This name must be unique across transformation policies. 
							* `type` - BLOCK drops any headers that are in the list of items, so it acts as an exclusion list.  ALLOW permits only the headers in the list and removes all others, so it acts as an inclusion list. 
						* `rename_headers` - Rename HTTP headers as they pass through the gateway. 
							* `items` - The list of headers.
								* `from` - The original case-insensitive name of the header.  This name must be unique across transformation policies. 
								* `to` - The new name of the header.  This name must be unique across transformation policies. 
						* `set_headers` - Set HTTP headers as they pass through the gateway. 
							* `items` - The list of headers.
								* `if_exists` - If a header with the same name already exists in the request, OVERWRITE will overwrite the value, APPEND will append to the existing value, or SKIP will keep the existing value. 
								* `name` - The case-insensitive name of the header.  This name must be unique across transformation policies. 
								* `values` - A list of new values.  Each value can be a constant or may include one or more expressions enclosed within ${} delimiters. 
					* `response_message` - HTTP response message.
					* `response_type` - Response Type.
					* `scopes` - List of scopes.
					* `source_uri_details` - Auth endpoint details.
						* `type` - Type of the Uri detail.
						* `uri` - The discovery URI for the auth server.
					* `type` - Type of the Validation failure Policy.
					* `use_cookies_for_intermediate_steps` - Defines whether or not to use cookies for OAuth2 intermediate steps. 
					* `use_cookies_for_session` - Defines whether or not to use cookies for session maintenance. 
					* `use_pkce` - Defines whether or not to support PKCE. 
				* `validation_policy` - Authentication Policies for the Token Authentication types.
					* `additional_validation_policy` - Additional JWT validation checks.
						* `audiences` - The list of intended recipients for the token.
						* `issuers` - A list of parties that could have issued the token.
						* `verify_claims` - A list of claims which should be validated to consider the token valid.
							* `is_required` - Whether the claim is required to be present in the JWT or not. If set to "false", the claim values will be matched only if the claim is present in the JWT. 
							* `key` - Name of the claim.
							* `values` - The list of acceptable values for a given claim. If this value is "null" or empty and "isRequired" set to "true", then the presence of this claim in the JWT is validated. 
					* `client_details` - Client App Credential details.
						* `client_id` - Client ID for the OAuth2/OIDC app.
						* `client_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Vault Service secret resource. 
						* `client_secret_version_number` - The version number of the client secret to use.
						* `type` - To specify where the Client App details should be taken from.
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
					* `max_cache_duration_in_hours` - The duration for which the introspect URL response should be cached before it is fetched again. 
					* `source_uri_details` - Auth endpoint details.
						* `type` - Type of the Uri detail.
						* `uri` - The discovery URI for the auth server.
					* `type` - Type of the token validation policy.
					* `uri` - The uri from which to retrieve the key. It must be accessible without authentication. 
				* `verify_claims` - A list of claims which should be validated to consider the token valid.
					* `is_required` - Whether the claim is required to be present in the JWT or not. If set to "false", the claim values will be matched only if the claim is present in the JWT. 
					* `key` - Name of the claim.
					* `values` - The list of acceptable values for a given claim. If this value is "null" or empty and "isRequired" set to "true", then the presence of this claim in the JWT is validated. 
			* `key` - Base policy for defining how to match the context variable in an incoming request with selection keys when dynamically routing and dynamically authenticating requests.
				* `expression` - A selection key string containing a wildcard to match with the context variable in an incoming request. If the context variable matches the string, the request is sent to the route or authentication server associated with the selection key. Valid wildcards are '*' (zero or more characters) and '+' (one or more characters). The string can only contain one wildcard, and the wildcard must be at the start or the end of the string.
				* `is_default` - Specifies whether to use the route or authentication server associated with this selection key as the default. The default is used if the value of a context variable in an incoming request does not match any of the other selection key values when dynamically routing and dynamically authenticating requests.
				* `name` - Name assigned to the branch.
				* `type` - Type of the selection key.
				* `values` - The set of selection keys to match with the context variable in an incoming request. If the context variable exactly matches one of the keys in the set, the request is sent to the route or authentication server associated with the set.
		* `selection_source` - The type of selector to use when dynamically routing and dynamically authenticating requests.
			* `selector` - String describing the context variable used as selector.
			* `type` - Type of the Selection source to use.
	* `mutual_tls` - Properties used to configure client mTLS verification when API Consumer makes connection to the gateway. 
		* `allowed_sans` - Allowed list of CN or SAN which will be used for verification of certificate.
		* `is_verified_certificate_required` - Determines whether to enable client verification when API Consumer makes connection to the gateway.
	* `rate_limiting` - Limit the number of requests that should be handled for the specified window using a specfic key.
		* `rate_in_requests_per_second` - The maximum number of requests per second to allow.
		* `rate_key` - The key used to group requests together.
	* `usage_plans` - Usage plan policies for this deployment
		* `token_locations` - A list of context variables specifying where API tokens may be located in a request. Example locations:
			* "request.headers[token]"
			* "request.query[token]"
			* "request.auth[Token]"
			* "request.path[TOKEN]" 
* `routes` - A list of routes that this API exposes.
	* `backend` - The backend to forward requests to. 
		* `allowed_post_logout_uris` - A list of allowed post-logout URLs to which a request can be redirected after revoke access
		* `body` - The body of the stock response from the mock backend.
		* `connect_timeout_in_seconds` - Defines a timeout for establishing a connection with a proxied server. 
		* `function_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Functions function resource. 
		* `headers` - The headers of the stock response from the mock backend.
			* `name` - Name of the header.
			* `value` - Value of the header.
		* `is_ssl_verify_disabled` - Defines whether or not to uphold SSL verification. 
		* `post_logout_state` - Defines a state that should be shared on redirecting to postLogout URL. 
		* `read_timeout_in_seconds` - Defines a timeout for reading a response from the proxied server. 
		* `routing_backends` - List of backends to chose from for Dynamic Routing.
			* `backend` - The backend to forward requests to.
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
				* `url` - The url of the proxied server.
			* `key` - Information around the values for selector of an authentication/ routing branch.
				* `expression` - String describing the expression with wildcards.
				* `is_default` - Information regarding whether this is the default branch.
				* `name` - Name assigned to the branch.
				* `type` - Information regarding type of the selection key.
				* `values` - Information regarding the set of values of selector for which this branch should be selected.
		* `selection_source` - Information around selector used for branching among routes/ authentication servers while dynamic routing/ authentication.
			* `selector` - String describing the context variable used as selector.
			* `type` - Type of the Selection source to use.
		* `send_timeout_in_seconds` - Defines a timeout for transmitting a request to the proxied server. 
		* `status` - The status code of the stock response from the mock backend.
		* `type` - Type of the API backend.
		* `url` - The url of the HTTP Backend
	* `logging_policies` - Policies controlling the pushing of logs to Oracle Cloud Infrastructure Public Logging. 
		* `access_log` - Configures the logging policies for the access logs of an API Deployment. 
			* `is_enabled` - Enables pushing of access logs to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket.

				Oracle recommends using the Oracle Cloud Infrastructure Logging service to enable, retrieve, and query access logs for an API Deployment. If there is an active log object for the API Deployment and its category is set to 'access' in Oracle Cloud Infrastructure Logging service, the logs will not be uploaded to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket.

				Please note that the functionality to push to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket has been deprecated and will be removed in the future. 
		* `execution_log` - Configures the logging policies for the execution logs of an API Deployment. 
			* `is_enabled` - Enables pushing of execution logs to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket.

				Oracle recommends using the Oracle Cloud Infrastructure Logging service to enable, retrieve, and query execution logs for an API Deployment. If there is an active log object for the API Deployment and its category is set to 'execution' in Oracle Cloud Infrastructure Logging service, the logs will not be uploaded to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket.

				Please note that the functionality to push to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket has been deprecated and will be removed in the future. 
			* `log_level` - Specifies the log level used to control logging output of execution logs. Enabling logging at a given level also enables logging at all higher levels. 
	* `methods` - A list of allowed methods on this route. 
	* `path` - A URL path pattern that must be matched on this route. The path pattern may contain a subset of RFC 6570 identifiers to allow wildcard and parameterized matching. 
	* `request_policies` - Behavior applied to any requests received by the API on this route. 
		* `authorization` - If authentication has been performed, validate whether the request scope (if any) applies to this route. If no RouteAuthorizationPolicy is defined for a route, a policy with a type of AUTHENTICATION_ONLY is applied. 
			* `allowed_scope` - A user whose scope includes any of these access ranges is allowed on this route. Access ranges are case-sensitive. 
			* `type` - Indicates how authorization should be applied. For a type of ANY_OF, an "allowedScope" property must also be specified. Otherwise, only a type is required. For a type of ANONYMOUS, an authenticated API must have the "isAnonymousAccessAllowed" property set to "true" in the authentication policy. 
		* `body_validation` - Validate the payload body of the incoming API requests on a specific route.
			* `content` - The content of the request body.
				* `media_type` - The media type is a [media type range](https://tools.ietf.org/html/rfc7231#appendix-D) subset restricted to the following schema

					( / (  "*" "/" "*" ) / ( type "/" "*" ) / ( type "/" subtype ) )

					For requests that match multiple media types, only the most specific media type is applicable. e.g. `text/plain` overrides `text/*` 
				* `validation_type` - Validation type defines the content validation method.

					Make the validation to first parse the body as the respective format. 
			* `required` - Determines if the request body is required in the request.
			* `validation_mode` - Validation behavior mode.

				In `ENFORCING` mode, upon a validation failure, the request will be rejected with a 4xx response and not sent to the backend.

				In `PERMISSIVE` mode, the result of the validation will be exposed as metrics while the request will follow the normal path.

				`DISABLED` type turns the validation off. 
		* `cors` - Enable CORS (Cross-Origin-Resource-Sharing) request handling. 
			* `allowed_headers` - The list of headers that will be allowed from the client via the Access-Control-Allow-Headers header. '*' will allow all headers. 
			* `allowed_methods` - The list of allowed HTTP methods that will be returned for the preflight OPTIONS request in the Access-Control-Allow-Methods header. '*' will allow all methods. 
			* `allowed_origins` - The list of allowed origins that the CORS handler will use to respond to CORS requests. The gateway will send the Access-Control-Allow-Origin header with the best origin match for the circumstances. '*' will match any origins, and 'null' will match queries from 'file:' origins. All other origins must be qualified with the scheme, full hostname, and port if necessary. 
			* `exposed_headers` - The list of headers that the client will be allowed to see from the response as indicated by the Access-Control-Expose-Headers header. '*' will expose all headers. 
			* `is_allow_credentials_enabled` - Whether to send the Access-Control-Allow-Credentials header to allow CORS requests with cookies. 
			* `max_age_in_seconds` - The time in seconds for the client to cache preflight responses. This is sent as the Access-Control-Max-Age if greater than 0. 
		* `header_transformations` - A set of transformations to apply to HTTP headers that pass through the gateway. 
			* `filter_headers` - Filter HTTP headers as they pass through the gateway.  The gateway applies filters after other transformations, so any headers set or renamed must also be listed here when using an ALLOW type policy. 
				* `items` - The list of headers. 
					* `name` - The case-insensitive name of the header.  This name must be unique across transformation policies. 
				* `type` - BLOCK drops any headers that are in the list of items, so it acts as an exclusion list.  ALLOW permits only the headers in the list and removes all others, so it acts as an inclusion list. 
			* `rename_headers` - Rename HTTP headers as they pass through the gateway. 
				* `items` - The list of headers.
					* `from` - The original case-insensitive name of the header.  This name must be unique across transformation policies. 
					* `to` - The new name of the header.  This name must be unique across transformation policies. 
			* `set_headers` - Set HTTP headers as they pass through the gateway. 
				* `items` - The list of headers.
					* `if_exists` - If a header with the same name already exists in the request, OVERWRITE will overwrite the value, APPEND will append to the existing value, or SKIP will keep the existing value. 
					* `name` - The case-insensitive name of the header.  This name must be unique across transformation policies. 
					* `values` - A list of new values.  Each value can be a constant or may include one or more expressions enclosed within ${} delimiters. 
		* `header_validations` - Validate the HTTP headers on the incoming API requests on a specific route.
			* `headers` - The List of Headers
				* `name` - Parameter name.
				* `required` - Determines if the header is required in the request.
			* `validation_mode` - Validation behavior mode.

				In `ENFORCING` mode, upon a validation failure, the request will be rejected with a 4xx response and not sent to the backend.

				In `PERMISSIVE` mode, the result of the validation will be exposed as metrics while the request will follow the normal path.

				`DISABLED` type turns the validation off. 
		* `query_parameter_transformations` - A set of transformations to apply to query parameters that pass through the gateway. 
			* `filter_query_parameters` - Filter parameters from the query string as they pass through the gateway.  The gateway applies filters after other transformations, so any parameters set or renamed must also be listed here when using an ALLOW type policy. 
				* `items` - The list of query parameters. 
					* `name` - The case-sensitive name of the query parameter. 
				* `type` - BLOCK drops any query parameters that are in the list of items, so it acts as an exclusion list.  ALLOW permits only the parameters in the list and removes all others, so it acts as an inclusion list. 
			* `rename_query_parameters` - Rename parameters on the query string as they pass through the gateway. 
				* `items` - The list of query parameters. 
					* `from` - The original case-sensitive name of the query parameter.  This name must be unique across transformation policies. 
					* `to` - The new name of the query parameter.  This name must be unique across transformation policies. 
			* `set_query_parameters` - Set parameters on the query string as they pass through the gateway. 
				* `items` - The list of query parameters. 
					* `if_exists` - If a query parameter with the same name already exists in the request, OVERWRITE will overwrite the value, APPEND will append to the existing value, or SKIP will keep the existing value. 
					* `name` - The case-sensitive name of the query parameter.  This name must be unique across transformation policies. 
					* `values` - A list of new values.  Each value can be a constant or may include one or more expressions enclosed within ${} delimiters. 
		* `query_parameter_validations` - Validate the URL query parameters on the incoming API requests on a specific route.
			* `parameters` - The List of Query Parameters
				* `name` - Parameter name.
				* `required` - Determines if the parameter is required in the request.
			* `validation_mode` - Validation behavior mode.

				In `ENFORCING` mode, upon a validation failure, the request will be rejected with a 4xx response and not sent to the backend.

				In `PERMISSIVE` mode, the result of the validation will be exposed as metrics while the request will follow the normal path.

				`DISABLED` type turns the validation off. 
		* `response_cache_lookup` - Base policy for Response Cache lookup. 
			* `cache_key_additions` - A list of context expressions whose values will be added to the base cache key. Values should contain an expression enclosed within ${} delimiters. Only the request context is available. 
			* `is_enabled` - Whether this policy is currently enabled. 
			* `is_private_caching_enabled` - Set true to allow caching responses where the request has an Authorization header. Ensure you have configured your  cache key additions to get the level of isolation across authenticated requests that you require.

				When false, any request with an Authorization header will not be stored in the Response Cache.

				If using the CustomAuthenticationPolicy then the tokenHeader/tokenQueryParam are also subject to this check. 
			* `type` - Type of the Response Cache Store Policy.
	* `response_policies` - Behavior applied to any responses sent by the API for requests on this route. 
		* `header_transformations` - A set of transformations to apply to HTTP headers that pass through the gateway. 
			* `filter_headers` - Filter HTTP headers as they pass through the gateway.  The gateway applies filters after other transformations, so any headers set or renamed must also be listed here when using an ALLOW type policy. 
				* `items` - The list of headers. 
					* `name` - The case-insensitive name of the header.  This name must be unique across transformation policies. 
				* `type` - BLOCK drops any headers that are in the list of items, so it acts as an exclusion list.  ALLOW permits only the headers in the list and removes all others, so it acts as an inclusion list. 
			* `rename_headers` - Rename HTTP headers as they pass through the gateway. 
				* `items` - The list of headers.
					* `from` - The original case-insensitive name of the header.  This name must be unique across transformation policies. 
					* `to` - The new name of the header.  This name must be unique across transformation policies. 
			* `set_headers` - Set HTTP headers as they pass through the gateway. 
				* `items` - The list of headers.
					* `if_exists` - If a header with the same name already exists in the request, OVERWRITE will overwrite the value, APPEND will append to the existing value, or SKIP will keep the existing value. 
					* `name` - The case-insensitive name of the header.  This name must be unique across transformation policies. 
					* `values` - A list of new values.  Each value can be a constant or may include one or more expressions enclosed within ${} delimiters. 
		* `response_cache_store` - Base policy for how a response from a backend is cached in the Response Cache. 
			* `time_to_live_in_seconds` - Sets the number of seconds for a response from a backend being stored in the Response Cache before it expires. 
			* `type` - Type of the Response Cache Store Policy.

