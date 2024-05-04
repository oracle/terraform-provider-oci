---
subcategory: "Load Balancer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_load_balancer_rule_set"
sidebar_current: "docs-oci-resource-load_balancer-rule_set"
description: |-
  Provides the Rule Set resource in Oracle Cloud Infrastructure Load Balancer service
---

# oci_load_balancer_rule_set
This resource provides the Rule Set resource in Oracle Cloud Infrastructure Load Balancer service.

Creates a new rule set associated with the specified load balancer. For more information, see
[Managing Rule Sets](https://docs.cloud.oracle.com/iaas/Content/Balance/Tasks/managingrulesets.htm).


## Example Usage

```hcl
resource "oci_load_balancer_rule_set" "test_rule_set" {
	#Required
	items {
		#Required
		action = var.rule_set_items_action

		#Optional
		allowed_methods = var.rule_set_items_allowed_methods
		are_invalid_characters_allowed = var.rule_set_items_are_invalid_characters_allowed
		conditions {
			#Required
			attribute_name = var.rule_set_items_conditions_attribute_name
			attribute_value = var.rule_set_items_conditions_attribute_value

			#Optional
			operator = var.rule_set_items_conditions_operator
		}
		default_max_connections = var.rule_set_items_default_max_connections
		description = var.rule_set_items_description
		header = var.rule_set_items_header
		http_large_header_size_in_kb = var.rule_set_items_http_large_header_size_in_kb
		ip_max_connections {

			#Optional
			ip_addresses = var.rule_set_items_ip_max_connections_ip_addresses
			max_connections = var.rule_set_items_ip_max_connections_max_connections
		}
		prefix = var.rule_set_items_prefix
		redirect_uri {

			#Optional
			host = var.rule_set_items_redirect_uri_host
			path = var.rule_set_items_redirect_uri_path
			port = var.rule_set_items_redirect_uri_port
			protocol = var.rule_set_items_redirect_uri_protocol
			query = var.rule_set_items_redirect_uri_query
		}
		response_code = var.rule_set_items_response_code
		status_code = var.rule_set_items_status_code
		suffix = var.rule_set_items_suffix
		value = var.rule_set_items_value
	}
	load_balancer_id = oci_load_balancer_load_balancer.test_load_balancer.id
	name = var.rule_set_name
}
```

## Argument Reference

The following arguments are supported:

* `items` - (Required) (Updatable) An array of rules that compose the rule set. For more information, see [Managing Rule Sets](https://docs.cloud.oracle.com/iaas/Content/Balance/Tasks/managingrulesets.htm)
	* `action` - (Required) (Updatable) The action can be one of these values: `ADD_HTTP_REQUEST_HEADER`, `ADD_HTTP_RESPONSE_HEADER`, `ALLOW`, `CONTROL_ACCESS_USING_HTTP_METHODS`, `EXTEND_HTTP_REQUEST_HEADER_VALUE`, `EXTEND_HTTP_RESPONSE_HEADER_VALUE`, `HTTP_HEADER`, `IP_BASED_MAX_CONNECTIONS`, `REDIRECT`, `REMOVE_HTTP_REQUEST_HEADER`, `REMOVE_HTTP_RESPONSE_HEADER`
	* `allowed_methods` - (Required when action=CONTROL_ACCESS_USING_HTTP_METHODS) (Updatable) The list of HTTP methods allowed for this listener.

		By default, you can specify only the standard HTTP methods defined in the [HTTP Method Registry](http://www.iana.org/assignments/http-methods/http-methods.xhtml). You can also see a list of supported standard HTTP methods in the Load Balancing service documentation at [Managing Rule Sets](https://docs.cloud.oracle.com/iaas/Content/Balance/Tasks/managingrulesets.htm).

		Your backend application must be able to handle the methods specified in this list.

		The list of HTTP methods is extensible. If you need to configure custom HTTP methods, contact [My Oracle Support](http://support.oracle.com/) to remove the restriction for your tenancy.

		Example: ["GET", "PUT", "POST", "PROPFIND"] 
	* `are_invalid_characters_allowed` - (Applicable when action=HTTP_HEADER) (Updatable) Indicates whether or not invalid characters in client header fields will be allowed. Valid names are composed of English letters, digits, hyphens and underscores. If "true", invalid characters are allowed in the HTTP header. If "false", invalid characters are not allowed in the HTTP header 
	* `conditions` - (Required when action=ALLOW | REDIRECT) (Updatable) 
		* `attribute_name` - (Required) (Updatable) The attribute_name can be one of these values: `PATH`, `SOURCE_IP_ADDRESS`, `SOURCE_VCN_ID`, `SOURCE_VCN_IP_ADDRESS`
		* `attribute_value` - (Required) (Updatable) Depends on `attribute_name`:
		    - when `attribute_name` = `SOURCE_IP_ADDRESS` | IPv4 or IPv6 address range to which the source IP address of incoming packet would be matched against
            - when `attribute_name` = `SOURCE_VCN_IP_ADDRESS` | IPv4 address range to which the original client IP address (in customer VCN) of incoming packet would be matched against
            - when `attribute_name` = `SOURCE_VCN_ID` | OCID of the customer VCN to which the service gateway embedded VCN ID of incoming packet would be matched against 
		* `operator` - (Required when attribute_name=PATH) (Updatable) A string that specifies how to compare the PathMatchCondition object's `attributeValue` string to the incoming URI.
			*  **EXACT_MATCH** - The incoming URI path must exactly and completely match the `attributeValue` string.
			*  **FORCE_LONGEST_PREFIX_MATCH** - The system looks for the `attributeValue` string with the best, longest match of the beginning portion of the incoming URI path.
			*  **PREFIX_MATCH** - The beginning portion of the incoming URI path must exactly match the `attributeValue` string.
			*  **SUFFIX_MATCH** - The ending portion of the incoming URI path must exactly match the `attributeValue` string. 
	* `default_max_connections` - (Applicable when action=IP_BASED_MAX_CONNECTIONS) (Updatable) The maximum number of connections that the any IP can make to a listener unless the IP is mentioned in maxConnections. If no defaultMaxConnections is specified the default is unlimited. 
	* `description` - (Applicable when action=ALLOW) (Updatable) A brief description of the access control rule. Avoid entering confidential information.

		example: `192.168.0.0/16 and 2001:db8::/32 are trusted clients. Whitelist them.` 
	* `header` - (Required when action=ADD_HTTP_REQUEST_HEADER | ADD_HTTP_RESPONSE_HEADER | EXTEND_HTTP_REQUEST_HEADER_VALUE | EXTEND_HTTP_RESPONSE_HEADER_VALUE | REMOVE_HTTP_REQUEST_HEADER | REMOVE_HTTP_RESPONSE_HEADER) (Updatable) A header name that conforms to RFC 7230.  Example: `example_header_name` 
	* `http_large_header_size_in_kb` - (Applicable when action=HTTP_HEADER) (Updatable) The maximum size of each buffer used for reading http client request header. This value indicates the maximum size allowed for each buffer. The allowed values for buffer size are 8, 16, 32 and 64. 
	* `ip_max_connections` - (Applicable when action=IP_BASED_MAX_CONNECTIONS) (Updatable) An array of IPs that have a maxConnection setting different than the default and what that maxConnection setting is 
		* `ip_addresses` - (Required when action=IP_BASED_MAX_CONNECTIONS) (Updatable) Each element in the list should be valid IPv4 or IPv6 CIDR Block address. Example: '["129.213.176.0/24", "150.136.187.0/24", "2002::1234:abcd:ffff:c0a8:101/64"]' 
		* `max_connections` - (Required when action=IP_BASED_MAX_CONNECTIONS) (Updatable) The maximum number of simultaneous connections that the specified IPs can make to the Listener. IPs without a maxConnections setting can make either defaultMaxConnections simultaneous connections to a listener or, if no defaultMaxConnections is specified, an unlimited number of simultaneous connections to a listener. 
	* `prefix` - (Applicable when action=EXTEND_HTTP_REQUEST_HEADER_VALUE | EXTEND_HTTP_RESPONSE_HEADER_VALUE) (Updatable) A string to prepend to the header value. The resulting header value must still conform to RFC 7230. With the following exceptions:
		*  value cannot contain `$`
		*  value cannot contain patterns like `{variable_name}`. They are reserved for future extensions. Currently, such values are invalid.

		Example: `example_prefix_value` 
	* `redirect_uri` - (Applicable when action=REDIRECT) (Updatable) An object that defines the redirect URI applied to the original request. The object property values compose the redirect URI.

		**NOTE:** The Load Balancing service cannot automatically detect or avoid infinite redirects. Be sure to provide meaningful, complete, and correct field values. If any component field of this object has no value, the system retains the value from the incoming HTTP request URI.

		For example, if you specify only the protocol field `https`, and the incoming request URI is `http://example.com:8080`, the resulting runtime redirect URI is `https://example.com:8080`. The system retains the host and port from the incoming URI and does not automatically change the port setting from `8080` to `443`.

		Be sure to configure valid percent-encoding (URL encoding) when needed.

		In addition to static string values, you can use the following tokens to construct the redirect URI. These tokens extract values from the incoming HTTP request URI.
		*  {protocol} : The protocol from the incoming HTTP request URI.
		*  {host}     : The domain name from the incoming HTTP request URI.
		*  {port}     : The port from the incoming HTTP request URI.
		*  {path}     : The path from the incoming HTTP request URI.
		*  {query}    : The query string from the incoming HTTP request URI.

		The tokens are case sensitive. For example, `{host}` is a valid token, but `{HOST}` is not.

		You can retain the literal characters of a token when you specify values for the path and query properties of the redirect URI. Use a backslash (\\) as the escape character for the \\, {, and } characters. For example, if the incoming HTTP request URI is `/video`, the path property value:

		`/example{path}123\{path\}`

		appears in the constructed redirect URI as:

		`/example/video123{path}` 
		* `host` - (Applicable when action=REDIRECT) (Updatable) The valid domain name (hostname) or IP address to use in the redirect URI.

			When this value is null, not set, or set to `{host}`, the service preserves the original domain name from the incoming HTTP request URI.

			All RedirectUri tokens are valid for this property. You can use any token more than once.

			Curly braces are valid in this property only to surround tokens, such as `{host}`

			Examples:
			*  **example.com** appears as `example.com` in the redirect URI.
			*  **in{host}** appears as `inexample.com` in the redirect URI if `example.com` is the hostname in the incoming HTTP request URI.
			*  **{port}{host}** appears as `8081example.com` in the redirect URI if `example.com` is the hostname and the port is `8081` in the incoming HTTP request URI. 
		* `path` - (Applicable when action=REDIRECT) (Updatable) The HTTP URI path to use in the redirect URI.

			When this value is null, not set, or set to `{path}`, the service preserves the original path from the incoming HTTP request URI. To omit the path from the redirect URI, set this value to an empty string, "".

			All RedirectUri tokens are valid for this property. You can use any token more than once.

			The path string must begin with `/` if it does not begin with the `{path}` token.

			Examples:
			*  __/example/video/123__ appears as `/example/video/123` in the redirect URI.
			*  __/example{path}__ appears as `/example/video/123` in the redirect URI if `/video/123` is the path in the incoming HTTP request URI.
			*  __{path}/123__ appears as `/example/video/123` in the redirect URI if `/example/video` is the path in the incoming HTTP request URI.
			*  __{path}123__ appears as `/example/video123` in the redirect URI if `/example/video` is the path in the incoming HTTP request URI.
			*  __/{host}/123__ appears as `/example.com/123` in the redirect URI if `example.com` is the hostname in the incoming HTTP request URI.
			*  __/{host}/{port}__ appears as `/example.com/123` in the redirect URI if `example.com` is the hostname and `123` is the port in the incoming HTTP request URI.
			*  __/{query}__ appears as `/lang=en` in the redirect URI if the query is `lang=en` in the incoming HTTP request URI. 
		* `port` - (Applicable when action=REDIRECT) (Updatable) The communication port to use in the redirect URI.

			Valid values include integers from 1 to 65535.

			When this value is null, the service preserves the original port from the incoming HTTP request URI.

			Example: `8081` 
		* `protocol` - (Applicable when action=REDIRECT) (Updatable) The HTTP protocol to use in the redirect URI.

			When this value is null, not set, or set to `{protocol}`, the service preserves the original protocol from the incoming HTTP request URI. Allowed values are:
			*  HTTP
			*  HTTPS
			*  {protocol}

			`{protocol}` is the only valid token for this property. It can appear only once in the value string.

			Example: `HTTPS` 
		* `query` - (Applicable when action=REDIRECT) (Updatable) The query string to use in the redirect URI.

			When this value is null, not set, or set to `{query}`, the service preserves the original query parameters from the incoming HTTP request URI.

			All `RedirectUri` tokens are valid for this property. You can use any token more than once.

			If the query string does not begin with the `{query}` token, it must begin with the question mark (?) character.

			You can specify multiple query parameters as a single string. Separate each query parameter with an ampersand (&) character. To omit all incoming query parameters from the redirect URI, set this value to an empty string, "".

			If the specified query string results in a redirect URI ending with `?` or `&`, the last character is truncated. For example, if the incoming URI is `http://host.com:8080/documents` and the query property value is `?lang=en&{query}`, the redirect URI is `http://host.com:8080/documents?lang=en`. The system truncates the final ampersand (&) because the incoming URI included no value to replace the {query} token.

			Examples:
			* **lang=en&time_zone=PST** appears as `lang=en&time_zone=PST` in the redirect URI.
			* **{query}** appears as `lang=en&time_zone=PST` in the redirect URI if `lang=en&time_zone=PST` is the query string in the incoming HTTP request. If the incoming HTTP request has no query parameters, the `{query}` token renders as an empty string.
			* **lang=en&{query}&time_zone=PST** appears as `lang=en&country=us&time_zone=PST` in the redirect URI if `country=us` is the query string in the incoming HTTP request. If the incoming HTTP request has no query parameters, this value renders as `lang=en&time_zone=PST`.
			*  **protocol={protocol}&hostname={host}** appears as `protocol=http&hostname=example.com` in the redirect URI if the protocol is `HTTP` and the hostname is `example.com` in the incoming HTTP request.
			*  **port={port}&hostname={host}** appears as `port=8080&hostname=example.com` in the redirect URI if the port is `8080` and the hostname is `example.com` in the incoming HTTP request URI. 
	* `response_code` - (Applicable when action=REDIRECT) (Updatable) The HTTP status code to return when the incoming request is redirected.

		The status line returned with the code is mapped from the standard HTTP specification. Valid response codes for redirection are:
		*  301
		*  302
		*  303
		*  307
		*  308

		The default value is `302` (Found).

		Example: `301` 
	* `status_code` - (Applicable when action=CONTROL_ACCESS_USING_HTTP_METHODS) (Updatable) The HTTP status code to return when the requested HTTP method is not in the list of allowed methods. The associated status line returned with the code is mapped from the standard HTTP specification. The default value is `405 (Method Not Allowed)`.  Example: 403 
	* `suffix` - (Applicable when action=EXTEND_HTTP_REQUEST_HEADER_VALUE | EXTEND_HTTP_RESPONSE_HEADER_VALUE) (Updatable) A string to append to the header value. The resulting header value must still conform to RFC 7230. With the following exceptions:
		*  value cannot contain `$`
		*  value cannot contain patterns like `{variable_name}`. They are reserved for future extensions. Currently, such values are invalid.

		Example: `example_suffix_value` 
	* `value` - (Required when action=ADD_HTTP_REQUEST_HEADER | ADD_HTTP_RESPONSE_HEADER) (Updatable) A header value that conforms to RFC 7230. With the following exceptions:
		*  value cannot contain `$`
		*  value cannot contain patterns like `{variable_name}`. They are reserved for future extensions. Currently, such values are invalid.

		Example: `example_value` 
* `load_balancer_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the specified load balancer.
* `name` - (Required) The name for this set of rules. It must be unique and it cannot be changed. Avoid entering confidential information.  Example: `example_rule_set` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `items` - An array of rules that compose the rule set.
	* `action` - The action can be one of these values: `ADD_HTTP_REQUEST_HEADER`, `ADD_HTTP_RESPONSE_HEADER`, `ALLOW`, `CONTROL_ACCESS_USING_HTTP_METHODS`, `EXTEND_HTTP_REQUEST_HEADER_VALUE`, `EXTEND_HTTP_RESPONSE_HEADER_VALUE`, `HTTP_HEADER`, `IP_BASED_MAX_CONNECTIONS`, `REDIRECT`, `REMOVE_HTTP_REQUEST_HEADER`, `REMOVE_HTTP_RESPONSE_HEADER`
	* `allowed_methods` - The list of HTTP methods allowed for this listener.

		By default, you can specify only the standard HTTP methods defined in the [HTTP Method Registry](http://www.iana.org/assignments/http-methods/http-methods.xhtml). You can also see a list of supported standard HTTP methods in the Load Balancing service documentation at [Managing Rule Sets](https://docs.cloud.oracle.com/iaas/Content/Balance/Tasks/managingrulesets.htm).

		Your backend application must be able to handle the methods specified in this list.

		The list of HTTP methods is extensible. If you need to configure custom HTTP methods, contact [My Oracle Support](http://support.oracle.com/) to remove the restriction for your tenancy.

		Example: ["GET", "PUT", "POST", "PROPFIND"] 
	* `are_invalid_characters_allowed` - Indicates whether or not invalid characters in client header fields will be allowed. Valid names are composed of English letters, digits, hyphens and underscores. If "true", invalid characters are allowed in the HTTP header. If "false", invalid characters are not allowed in the HTTP header 
	* `conditions` - 
		* `attribute_name` - (Required) (Updatable) The attribute_name can be one of these values: `PATH`, `SOURCE_IP_ADDRESS`, `SOURCE_VCN_ID`, `SOURCE_VCN_IP_ADDRESS`
		* `attribute_value` - (Required) (Updatable) Depends on `attribute_name`:
		    - when `attribute_name` = `SOURCE_IP_ADDRESS` | IPv4 or IPv6 address range to which the source IP address of incoming packet would be matched against
            - when `attribute_name` = `SOURCE_VCN_IP_ADDRESS` | IPv4 address range to which the original client IP address (in customer VCN) of incoming packet would be matched against
            - when `attribute_name` = `SOURCE_VCN_ID` | OCID of the customer VCN to which the service gateway embedded VCN ID of incoming packet would be matched against 
		* `operator` - A string that specifies how to compare the PathMatchCondition object's `attributeValue` string to the incoming URI.
			*  **EXACT_MATCH** - The incoming URI path must exactly and completely match the `attributeValue` string.
			*  **FORCE_LONGEST_PREFIX_MATCH** - The system looks for the `attributeValue` string with the best, longest match of the beginning portion of the incoming URI path.
			*  **PREFIX_MATCH** - The beginning portion of the incoming URI path must exactly match the `attributeValue` string.
			*  **SUFFIX_MATCH** - The ending portion of the incoming URI path must exactly match the `attributeValue` string. 
	* `default_max_connections` - The maximum number of connections that the any IP can make to a listener unless the IP is mentioned in maxConnections. If no defaultMaxConnections is specified the default is unlimited. 
	* `description` - A brief description of the access control rule. Avoid entering confidential information.

		example: `192.168.0.0/16 and 2001:db8::/32 are trusted clients. Whitelist them.` 
	* `header` - A header name that conforms to RFC 7230.  Example: `example_header_name` 
	* `http_large_header_size_in_kb` - The maximum size of each buffer used for reading http client request header. This value indicates the maximum size allowed for each buffer. The allowed values for buffer size are 8, 16, 32 and 64. 
	* `ip_max_connections` - An array of IPs that have a maxConnection setting different than the default and what that maxConnection setting is 
		* `ip_addresses` - Each element in the list should be valid IPv4 or IPv6 CIDR Block address. Example: '["129.213.176.0/24", "150.136.187.0/24", "2002::1234:abcd:ffff:c0a8:101/64"]' 
		* `max_connections` - The maximum number of simultaneous connections that the specified IPs can make to the Listener. IPs without a maxConnections setting can make either defaultMaxConnections simultaneous connections to a listener or, if no defaultMaxConnections is specified, an unlimited number of simultaneous connections to a listener. 
	* `prefix` - A string to prepend to the header value. The resulting header value must still conform to RFC 7230. With the following exceptions:
		*  value cannot contain `$`
		*  value cannot contain patterns like `{variable_name}`. They are reserved for future extensions. Currently, such values are invalid.

		Example: `example_prefix_value` 
	* `redirect_uri` - An object that defines the redirect URI applied to the original request. The object property values compose the redirect URI.

		**NOTE:** The Load Balancing service cannot automatically detect or avoid infinite redirects. Be sure to provide meaningful, complete, and correct field values. If any component field of this object has no value, the system retains the value from the incoming HTTP request URI.

		For example, if you specify only the protocol field `https`, and the incoming request URI is `http://example.com:8080`, the resulting runtime redirect URI is `https://example.com:8080`. The system retains the host and port from the incoming URI and does not automatically change the port setting from `8080` to `443`.

		Be sure to configure valid percent-encoding (URL encoding) when needed.

		In addition to static string values, you can use the following tokens to construct the redirect URI. These tokens extract values from the incoming HTTP request URI.
		*  {protocol} : The protocol from the incoming HTTP request URI.
		*  {host}     : The domain name from the incoming HTTP request URI.
		*  {port}     : The port from the incoming HTTP request URI.
		*  {path}     : The path from the incoming HTTP request URI.
		*  {query}    : The query string from the incoming HTTP request URI.

		The tokens are case sensitive. For example, `{host}` is a valid token, but `{HOST}` is not.

		You can retain the literal characters of a token when you specify values for the path and query properties of the redirect URI. Use a backslash (\\) as the escape character for the \\, {, and } characters. For example, if the incoming HTTP request URI is `/video`, the path property value:

		`/example{path}123\{path\}`

		appears in the constructed redirect URI as:

		`/example/video123{path}` 
		* `host` - The valid domain name (hostname) or IP address to use in the redirect URI.

			When this value is null, not set, or set to `{host}`, the service preserves the original domain name from the incoming HTTP request URI.

			All RedirectUri tokens are valid for this property. You can use any token more than once.

			Curly braces are valid in this property only to surround tokens, such as `{host}`

			Examples:
			*  **example.com** appears as `example.com` in the redirect URI.
			*  **in{host}** appears as `inexample.com` in the redirect URI if `example.com` is the hostname in the incoming HTTP request URI.
			*  **{port}{host}** appears as `8081example.com` in the redirect URI if `example.com` is the hostname and the port is `8081` in the incoming HTTP request URI. 
		* `path` - The HTTP URI path to use in the redirect URI.

			When this value is null, not set, or set to `{path}`, the service preserves the original path from the incoming HTTP request URI. To omit the path from the redirect URI, set this value to an empty string, "".

			All RedirectUri tokens are valid for this property. You can use any token more than once.

			The path string must begin with `/` if it does not begin with the `{path}` token.

			Examples:
			*  __/example/video/123__ appears as `/example/video/123` in the redirect URI.
			*  __/example{path}__ appears as `/example/video/123` in the redirect URI if `/video/123` is the path in the incoming HTTP request URI.
			*  __{path}/123__ appears as `/example/video/123` in the redirect URI if `/example/video` is the path in the incoming HTTP request URI.
			*  __{path}123__ appears as `/example/video123` in the redirect URI if `/example/video` is the path in the incoming HTTP request URI.
			*  __/{host}/123__ appears as `/example.com/123` in the redirect URI if `example.com` is the hostname in the incoming HTTP request URI.
			*  __/{host}/{port}__ appears as `/example.com/123` in the redirect URI if `example.com` is the hostname and `123` is the port in the incoming HTTP request URI.
			*  __/{query}__ appears as `/lang=en` in the redirect URI if the query is `lang=en` in the incoming HTTP request URI. 
		* `port` - The communication port to use in the redirect URI.

			Valid values include integers from 1 to 65535.

			When this value is null, the service preserves the original port from the incoming HTTP request URI.

			Example: `8081` 
		* `protocol` - The HTTP protocol to use in the redirect URI.

			When this value is null, not set, or set to `{protocol}`, the service preserves the original protocol from the incoming HTTP request URI. Allowed values are:
			*  HTTP
			*  HTTPS
			*  {protocol}

			`{protocol}` is the only valid token for this property. It can appear only once in the value string.

			Example: `HTTPS` 
		* `query` - The query string to use in the redirect URI.

			When this value is null, not set, or set to `{query}`, the service preserves the original query parameters from the incoming HTTP request URI.

			All `RedirectUri` tokens are valid for this property. You can use any token more than once.

			If the query string does not begin with the `{query}` token, it must begin with the question mark (?) character.

			You can specify multiple query parameters as a single string. Separate each query parameter with an ampersand (&) character. To omit all incoming query parameters from the redirect URI, set this value to an empty string, "".

			If the specified query string results in a redirect URI ending with `?` or `&`, the last character is truncated. For example, if the incoming URI is `http://host.com:8080/documents` and the query property value is `?lang=en&{query}`, the redirect URI is `http://host.com:8080/documents?lang=en`. The system truncates the final ampersand (&) because the incoming URI included no value to replace the {query} token.

			Examples:
			* **lang=en&time_zone=PST** appears as `lang=en&time_zone=PST` in the redirect URI.
			* **{query}** appears as `lang=en&time_zone=PST` in the redirect URI if `lang=en&time_zone=PST` is the query string in the incoming HTTP request. If the incoming HTTP request has no query parameters, the `{query}` token renders as an empty string.
			* **lang=en&{query}&time_zone=PST** appears as `lang=en&country=us&time_zone=PST` in the redirect URI if `country=us` is the query string in the incoming HTTP request. If the incoming HTTP request has no query parameters, this value renders as `lang=en&time_zone=PST`.
			*  **protocol={protocol}&hostname={host}** appears as `protocol=http&hostname=example.com` in the redirect URI if the protocol is `HTTP` and the hostname is `example.com` in the incoming HTTP request.
			*  **port={port}&hostname={host}** appears as `port=8080&hostname=example.com` in the redirect URI if the port is `8080` and the hostname is `example.com` in the incoming HTTP request URI. 
	* `response_code` - The HTTP status code to return when the incoming request is redirected.

		The status line returned with the code is mapped from the standard HTTP specification. Valid response codes for redirection are:
		*  301
		*  302
		*  303
		*  307
		*  308

		The default value is `302` (Found).

		Example: `301` 
	* `status_code` - The HTTP status code to return when the requested HTTP method is not in the list of allowed methods. The associated status line returned with the code is mapped from the standard HTTP specification. The default value is `405 (Method Not Allowed)`.  Example: 403 
	* `suffix` - A string to append to the header value. The resulting header value must still conform to RFC 7230. With the following exceptions:
		*  value cannot contain `$`
		*  value cannot contain patterns like `{variable_name}`. They are reserved for future extensions. Currently, such values are invalid.

		Example: `example_suffix_value` 
	* `value` - A header value that conforms to RFC 7230. With the following exceptions:
		*  value cannot contain `$`
		*  value cannot contain patterns like `{variable_name}`. They are reserved for future extensions. Currently, such values are invalid.

		Example: `example_value` 
* `name` - The name for this set of rules. It must be unique and it cannot be changed. Avoid entering confidential information.  Example: `example_rule_set` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Rule Set
	* `update` - (Defaults to 20 minutes), when updating the Rule Set
	* `delete` - (Defaults to 20 minutes), when destroying the Rule Set


## Import

RuleSets can be imported using the `id`, e.g.

```
$ terraform import oci_load_balancer_rule_set.test_rule_set "loadBalancers/{loadBalancerId}/ruleSets/{ruleSetName}" 
```

