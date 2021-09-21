---
subcategory: "Waf"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_waf_web_app_firewall_policy"
sidebar_current: "docs-oci-resource-waf-web_app_firewall_policy"
description: |-
  Provides the Web App Firewall Policy resource in Oracle Cloud Infrastructure Waf service
---

# oci_waf_web_app_firewall_policy
This resource provides the Web App Firewall Policy resource in Oracle Cloud Infrastructure Waf service.

Creates a new WebAppFirewallPolicy.


## Example Usage

```hcl
resource "oci_waf_web_app_firewall_policy" "test_web_app_firewall_policy" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	actions {
		#Required
		name = var.web_app_firewall_policy_actions_name
		type = var.web_app_firewall_policy_actions_type

		#Optional
		body {
			#Required
			text = var.web_app_firewall_policy_actions_body_text
			type = var.web_app_firewall_policy_actions_body_type
		}
		code = var.web_app_firewall_policy_actions_code
		headers {

			#Optional
			name = var.web_app_firewall_policy_actions_headers_name
			value = var.web_app_firewall_policy_actions_headers_value
		}
	}
	defined_tags = {"foo-namespace.bar-key"= "value"}
	display_name = var.web_app_firewall_policy_display_name
	freeform_tags = {"bar-key"= "value"}
	request_access_control {
		#Required
		default_action_name = var.web_app_firewall_policy_request_access_control_default_action_name

		#Optional
		rules {
			#Required
			action_name = var.web_app_firewall_policy_request_access_control_rules_action_name
			name = var.web_app_firewall_policy_request_access_control_rules_name
			type = var.web_app_firewall_policy_request_access_control_rules_type

			#Optional
			condition = var.web_app_firewall_policy_request_access_control_rules_condition
			condition_language = var.web_app_firewall_policy_request_access_control_rules_condition_language
		}
	}
	request_protection {

		#Optional
		rules {
			#Required
			action_name = var.web_app_firewall_policy_request_protection_rules_action_name
			name = var.web_app_firewall_policy_request_protection_rules_name
			protection_capabilities {
				#Required
				key = var.web_app_firewall_policy_request_protection_rules_protection_capabilities_key
				version = var.web_app_firewall_policy_request_protection_rules_protection_capabilities_version

				#Optional
				action_name = var.web_app_firewall_policy_request_protection_rules_protection_capabilities_action_name
				collaborative_action_threshold = var.web_app_firewall_policy_request_protection_rules_protection_capabilities_collaborative_action_threshold
				collaborative_weights {
					#Required
					key = var.web_app_firewall_policy_request_protection_rules_protection_capabilities_collaborative_weights_key
					weight = var.web_app_firewall_policy_request_protection_rules_protection_capabilities_collaborative_weights_weight
				}
				exclusions {

					#Optional
					args = var.web_app_firewall_policy_request_protection_rules_protection_capabilities_exclusions_args
					request_cookies = var.web_app_firewall_policy_request_protection_rules_protection_capabilities_exclusions_request_cookies
				}
			}
			type = var.web_app_firewall_policy_request_protection_rules_type

			#Optional
			condition = var.web_app_firewall_policy_request_protection_rules_condition
			condition_language = var.web_app_firewall_policy_request_protection_rules_condition_language
			protection_capability_settings {

				#Optional
				allowed_http_methods = var.web_app_firewall_policy_request_protection_rules_protection_capability_settings_allowed_http_methods
				max_http_request_header_length = var.web_app_firewall_policy_request_protection_rules_protection_capability_settings_max_http_request_header_length
				max_http_request_headers = var.web_app_firewall_policy_request_protection_rules_protection_capability_settings_max_http_request_headers
				max_number_of_arguments = var.web_app_firewall_policy_request_protection_rules_protection_capability_settings_max_number_of_arguments
				max_single_argument_length = var.web_app_firewall_policy_request_protection_rules_protection_capability_settings_max_single_argument_length
				max_total_argument_length = var.web_app_firewall_policy_request_protection_rules_protection_capability_settings_max_total_argument_length
			}
		}
	}
	request_rate_limiting {

		#Optional
		rules {
			#Required
			action_name = var.web_app_firewall_policy_request_rate_limiting_rules_action_name
			configurations {
				#Required
				period_in_seconds = var.web_app_firewall_policy_request_rate_limiting_rules_configurations_period_in_seconds
				requests_limit = var.web_app_firewall_policy_request_rate_limiting_rules_configurations_requests_limit

				#Optional
				action_duration_in_seconds = var.web_app_firewall_policy_request_rate_limiting_rules_configurations_action_duration_in_seconds
			}
			name = var.web_app_firewall_policy_request_rate_limiting_rules_name
			type = var.web_app_firewall_policy_request_rate_limiting_rules_type

			#Optional
			condition = var.web_app_firewall_policy_request_rate_limiting_rules_condition
			condition_language = var.web_app_firewall_policy_request_rate_limiting_rules_condition_language
		}
	}
	response_access_control {

		#Optional
		rules {
			#Required
			action_name = var.web_app_firewall_policy_response_access_control_rules_action_name
			name = var.web_app_firewall_policy_response_access_control_rules_name
			type = var.web_app_firewall_policy_response_access_control_rules_type

			#Optional
			condition = var.web_app_firewall_policy_response_access_control_rules_condition
			condition_language = var.web_app_firewall_policy_response_access_control_rules_condition_language
		}
	}
	response_protection {

		#Optional
		rules {
			#Required
			action_name = var.web_app_firewall_policy_response_protection_rules_action_name
			name = var.web_app_firewall_policy_response_protection_rules_name
			protection_capabilities {
				#Required
				key = var.web_app_firewall_policy_response_protection_rules_protection_capabilities_key
				version = var.web_app_firewall_policy_response_protection_rules_protection_capabilities_version

				#Optional
				action_name = var.web_app_firewall_policy_response_protection_rules_protection_capabilities_action_name
				collaborative_action_threshold = var.web_app_firewall_policy_response_protection_rules_protection_capabilities_collaborative_action_threshold
				collaborative_weights {
					#Required
					key = var.web_app_firewall_policy_response_protection_rules_protection_capabilities_collaborative_weights_key
					weight = var.web_app_firewall_policy_response_protection_rules_protection_capabilities_collaborative_weights_weight
				}
				exclusions {

					#Optional
					args = var.web_app_firewall_policy_response_protection_rules_protection_capabilities_exclusions_args
					request_cookies = var.web_app_firewall_policy_response_protection_rules_protection_capabilities_exclusions_request_cookies
				}
			}
			type = var.web_app_firewall_policy_response_protection_rules_type

			#Optional
			condition = var.web_app_firewall_policy_response_protection_rules_condition
			condition_language = var.web_app_firewall_policy_response_protection_rules_condition_language
			protection_capability_settings {

				#Optional
				allowed_http_methods = var.web_app_firewall_policy_response_protection_rules_protection_capability_settings_allowed_http_methods
				max_http_request_header_length = var.web_app_firewall_policy_response_protection_rules_protection_capability_settings_max_http_request_header_length
				max_http_request_headers = var.web_app_firewall_policy_response_protection_rules_protection_capability_settings_max_http_request_headers
				max_number_of_arguments = var.web_app_firewall_policy_response_protection_rules_protection_capability_settings_max_number_of_arguments
				max_single_argument_length = var.web_app_firewall_policy_response_protection_rules_protection_capability_settings_max_single_argument_length
				max_total_argument_length = var.web_app_firewall_policy_response_protection_rules_protection_capability_settings_max_total_argument_length
			}
		}
	}
	system_tags = var.web_app_firewall_policy_system_tags
}
```

## Argument Reference

The following arguments are supported:

* `actions` - (Optional) (Updatable) Predefined actions for use in multiple different rules. Not all actions are supported in every module. Some actions terminate further execution of modules and rules in a module and some do not. Actions names must be unique within this array. 
	* `body` - (Applicable when type=RETURN_HTTP_RESPONSE) (Updatable) Type of returned HTTP response body.
		* `text` - (Required) (Updatable) Static response body text.
		* `type` - (Required) (Updatable) Type of HttpResponseBody.
	* `code` - (Required when type=RETURN_HTTP_RESPONSE) (Updatable) Response code.

		The following response codes are valid values for this property:
		* 2xx

		200 OK 201 Created 202 Accepted 206 Partial Content
		* 3xx

		300 Multiple Choices 301 Moved Permanently 302 Found 303 See Other 307 Temporary Redirect
		* 4xx

		400 Bad Request 401 Unauthorized 403 Forbidden 404 Not Found 405 Method Not Allowed 408 Request Timeout 409 Conflict 411 Length Required 412 Precondition Failed 413 Payload Too Large 414 URI Too Long 415 Unsupported Media Type 416 Range Not Satisfiable 422 Unprocessable Entity 494 Request Header Too Large 495 Cert Error 496 No Cert 497 HTTP to HTTPS
		* 5xx

		500 Internal Server Error 501 Not Implemented 502 Bad Gateway 503 Service Unavailable 504 Gateway Timeout 507 Insufficient Storage

		Example: `200` 
	* `headers` - (Applicable when type=RETURN_HTTP_RESPONSE) (Updatable) Adds headers defined in this array for HTTP response.

		Hop-by-hop headers are not allowed to be set:
		* Connection
		* Keep-Alive
		* Proxy-Authenticate
		* Proxy-Authorization
		* TE
		* Trailer
		* Transfer-Encoding
		* Upgrade 
		* `name` - (Required when type=RETURN_HTTP_RESPONSE) (Updatable) The name of the header field.
		* `value` - (Required when type=RETURN_HTTP_RESPONSE) (Updatable) The value of the header field.
	* `name` - (Required) (Updatable) Action name. Can be used to reference the action.
	* `type` - (Required) (Updatable) 
		* **CHECK** is a non-terminating action that does not stop the execution of rules in current module, just emits a log message documenting result of rule execution.
		* **ALLOW** is a non-terminating action which upon matching rule skips all remaining rules in the current module.
		* **RETURN_HTTP_RESPONSE** is a terminating action which is executed immediately, returns a defined HTTP response. 
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Optional) (Updatable) WebAppFirewallPolicy display name, can be renamed.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `request_access_control` - (Optional) (Updatable) Module that allows inspection of HTTP request properties and to return a defined HTTP response. In this module, rules with the name 'Default Action' are not allowed, since this name is reserved for default action logs. 
	* `default_action_name` - (Required) (Updatable) References an default Action to take if no AccessControlRule was matched. Allowed action types:
		* **ALLOW** continues execution of other modules and their rules.
		* **RETURN_HTTP_RESPONSE** terminates further execution of modules and rules and returns defined HTTP response. 
	* `rules` - (Optional) (Updatable) Ordered list of AccessControlRules. Rules are executed in order of appearance in this array.
		* `action_name` - (Required) (Updatable) References action by name from actions defined in WebAppFirewallPolicy.
		* `condition` - (Optional) (Updatable) An expression that determines whether or not the rule action should be executed. 
		* `condition_language` - (Optional) (Updatable) The language used to parse condition from field `condition`. Available languages:
			* **JMESPATH** an extended JMESPath language syntax. 
		* `name` - (Required) (Updatable) Rule name. Must be unique within the module.
		* `type` - (Required) (Updatable) Type of WebAppFirewallPolicyRule.
* `request_protection` - (Optional) (Updatable) Module that allows to enable OCI-managed protection capabilities for incoming HTTP requests.
	* `rules` - (Optional) (Updatable) Ordered list of ProtectionRules. Rules are executed in order of appearance in this array. ProtectionRules in this array can only use protection cCapabilities of REQUEST_PROTECTION_CAPABILITY type. 
		* `action_name` - (Required) (Updatable) References action by name from actions defined in WebAppFirewallPolicy.
		* `condition` - (Optional) (Updatable) An expression that determines whether or not the rule action should be executed. 
		* `condition_language` - (Optional) (Updatable) The language used to parse condition from field `condition`. Available languages:
			* **JMESPATH** an extended JMESPath language syntax. 
		* `name` - (Required) (Updatable) Rule name. Must be unique within the module.
		* `protection_capabilities` - (Required) (Updatable) An ordered list that references OCI-managed protection capabilities. Referenced protection capabilities are executed in order of appearance. The array cannot contain entries with the same pair of capability key and version more than once. 
			* `action_name` - (Optional) (Updatable) Override action to take if capability was triggered, defined in Protection Rule for this capability. Only actions of type CHECK are allowed. 
			* `collaborative_action_threshold` - (Optional) (Updatable) The minimum sum of weights of associated collaborative protection capabilities that have triggered which must be reached in order for _this_ capability to trigger. This field is ignored for non-collaborative capabilities. 
			* `collaborative_weights` - (Optional) (Updatable) Explicit weight values to use for associated collaborative protection capabilities. 
				* `key` - (Required) (Updatable) Unique key of collaborative capability for which weight will be overridden.
				* `weight` - (Required) (Updatable) The value of weight to set.
			* `exclusions` - (Optional) (Updatable) Identifies specific HTTP message parameters to exclude from inspection by a protection capability. 
				* `args` - (Optional) (Updatable) List of URL query parameter values from form-urlencoded XML, JSON, AMP, or POST payloads to exclude from inspecting. Example: If we have query parameter 'argumentName=argumentValue' and args=['argumentName'], both 'argumentName' and 'argumentValue' will not be inspected. 
				* `request_cookies` - (Optional) (Updatable) List of HTTP request cookie values (by cookie name) to exclude from inspecting. Example: If we have cookie 'cookieName=cookieValue' and requestCookies=['cookieName'], both 'cookieName' and 'cookieValue' will not be inspected. 
			* `key` - (Required) (Updatable) Unique key of referenced protection capability.
			* `version` - (Required) (Updatable) Version of referenced protection capability.
		* `protection_capability_settings` - (Optional) (Updatable) Settings for protection capabilities 
			* `allowed_http_methods` - (Optional) (Updatable) List of allowed HTTP methods. Each value as a RFC7230 formated token string. Used in protection capability 911100: Restrict HTTP Request Methods. 
			* `max_http_request_header_length` - (Optional) (Updatable) Maximum allowed length of headers in an HTTP request. Used in protection capability: 9200024: Limit length of request header size. 
			* `max_http_request_headers` - (Optional) (Updatable) Maximum number of headers allowed in an HTTP request. Used in protection capability 9200014: Limit Number of Request Headers. 
			* `max_number_of_arguments` - (Optional) (Updatable) Maximum number of arguments allowed. Used in protection capability 920380: Number of Arguments Limits. 
			* `max_single_argument_length` - (Optional) (Updatable) Maximum allowed length of a single argument. Used in protection capability 920370: Limit argument value length. 
			* `max_total_argument_length` - (Optional) (Updatable) Maximum allowed total length of all arguments. Used in protection capability 920390: Limit arguments total length. 
		* `type` - (Required) (Updatable) Type of WebAppFirewallPolicyRule.
* `request_rate_limiting` - (Optional) (Updatable) Module that allows inspection of HTTP connection properties and to limit requests frequency for a given key.
	* `rules` - (Optional) (Updatable) Ordered list of RequestRateLimitingRules. Rules are executed in order of appearance in this array. 
		* `action_name` - (Required) (Updatable) References action by name from actions defined in WebAppFirewallPolicy.
		* `condition` - (Optional) (Updatable) An expression that determines whether or not the rule action should be executed. 
		* `condition_language` - (Optional) (Updatable) The language used to parse condition from field `condition`. Available languages:
			* **JMESPATH** an extended JMESPath language syntax. 
		* `configurations` - (Required) (Updatable) Rate Limiting Configurations. Each configuration counts requests towards its own `requestsLimit`. 
			* `action_duration_in_seconds` - (Optional) (Updatable) Duration of block action application in seconds when `requestsLimit` is reached. Optional and can be 0 (no block duration).
			* `period_in_seconds` - (Required) (Updatable) Evaluation period in seconds.
			* `requests_limit` - (Required) (Updatable) Requests allowed per evaluation period.
		* `name` - (Required) (Updatable) Rule name. Must be unique within the module.
		* `type` - (Required) (Updatable) Type of WebAppFirewallPolicyRule.
* `response_access_control` - (Optional) (Updatable) Module that allows inspection of HTTP response properties and to return a defined HTTP response.
	* `rules` - (Optional) (Updatable) Ordered list of AccessControlRules. Rules are executed in order of appearance in this array.
		* `action_name` - (Required) (Updatable) References action by name from actions defined in WebAppFirewallPolicy.
		* `condition` - (Optional) (Updatable) An expression that determines whether or not the rule action should be executed. 
		* `condition_language` - (Optional) (Updatable) The language used to parse condition from field `condition`. Available languages:
			* **JMESPATH** an extended JMESPath language syntax. 
		* `name` - (Required) (Updatable) Rule name. Must be unique within the module.
		* `type` - (Required) (Updatable) Type of WebAppFirewallPolicyRule.
* `response_protection` - (Optional) (Updatable) Module that allows to enable OCI-managed protection capabilities for HTTP responses.
	* `rules` - (Optional) (Updatable) Ordered list of ProtectionRules. Rules are executed in order of appearance in this array. ProtectionRules in this array can only use protection capabilities of RESPONSE_PROTECTION_CAPABILITY type. 
		* `action_name` - (Required) (Updatable) References action by name from actions defined in WebAppFirewallPolicy.
		* `condition` - (Optional) (Updatable) An expression that determines whether or not the rule action should be executed. 
		* `condition_language` - (Optional) (Updatable) The language used to parse condition from field `condition`. Available languages:
			* **JMESPATH** an extended JMESPath language syntax. 
		* `name` - (Required) (Updatable) Rule name. Must be unique within the module.
		* `protection_capabilities` - (Required) (Updatable) An ordered list that references OCI-managed protection capabilities. Referenced protection capabilities are executed in order of appearance. The array cannot contain entries with the same pair of capability key and version more than once. 
			* `action_name` - (Optional) (Updatable) Override action to take if capability was triggered, defined in Protection Rule for this capability. Only actions of type CHECK are allowed. 
			* `collaborative_action_threshold` - (Optional) (Updatable) The minimum sum of weights of associated collaborative protection capabilities that have triggered which must be reached in order for _this_ capability to trigger. This field is ignored for non-collaborative capabilities. 
			* `collaborative_weights` - (Optional) (Updatable) Explicit weight values to use for associated collaborative protection capabilities. 
				* `key` - (Required) (Updatable) Unique key of collaborative capability for which weight will be overridden.
				* `weight` - (Required) (Updatable) The value of weight to set.
			* `exclusions` - (Optional) (Updatable) Identifies specific HTTP message parameters to exclude from inspection by a protection capability. 
				* `args` - (Optional) (Updatable) List of URL query parameter values from form-urlencoded XML, JSON, AMP, or POST payloads to exclude from inspecting. Example: If we have query parameter 'argumentName=argumentValue' and args=['argumentName'], both 'argumentName' and 'argumentValue' will not be inspected. 
				* `request_cookies` - (Optional) (Updatable) List of HTTP request cookie values (by cookie name) to exclude from inspecting. Example: If we have cookie 'cookieName=cookieValue' and requestCookies=['cookieName'], both 'cookieName' and 'cookieValue' will not be inspected. 
			* `key` - (Required) (Updatable) Unique key of referenced protection capability.
			* `version` - (Required) (Updatable) Version of referenced protection capability.
		* `protection_capability_settings` - (Optional) (Updatable) Settings for protection capabilities 
			* `allowed_http_methods` - (Optional) (Updatable) List of allowed HTTP methods. Each value as a RFC7230 formated token string. Used in protection capability 911100: Restrict HTTP Request Methods. 
			* `max_http_request_header_length` - (Optional) (Updatable) Maximum allowed length of headers in an HTTP request. Used in protection capability: 9200024: Limit length of request header size. 
			* `max_http_request_headers` - (Optional) (Updatable) Maximum number of headers allowed in an HTTP request. Used in protection capability 9200014: Limit Number of Request Headers. 
			* `max_number_of_arguments` - (Optional) (Updatable) Maximum number of arguments allowed. Used in protection capability 920380: Number of Arguments Limits. 
			* `max_single_argument_length` - (Optional) (Updatable) Maximum allowed length of a single argument. Used in protection capability 920370: Limit argument value length. 
			* `max_total_argument_length` - (Optional) (Updatable) Maximum allowed total length of all arguments. Used in protection capability 920390: Limit arguments total length. 
		* `type` - (Required) (Updatable) Type of WebAppFirewallPolicyRule.
* `system_tags` - (Optional) (Updatable) Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `actions` - Predefined actions for use in multiple different rules. Not all actions are supported in every module. Some actions terminate further execution of modules and rules in a module and some do not. Actions names must be unique within this array. 
	* `body` - Type of returned HTTP response body.
		* `text` - Static response body text.
		* `type` - Type of HttpResponseBody.
	* `code` - Response code.

		The following response codes are valid values for this property:
		* 2xx

		200 OK 201 Created 202 Accepted 206 Partial Content
		* 3xx

		300 Multiple Choices 301 Moved Permanently 302 Found 303 See Other 307 Temporary Redirect
		* 4xx

		400 Bad Request 401 Unauthorized 403 Forbidden 404 Not Found 405 Method Not Allowed 408 Request Timeout 409 Conflict 411 Length Required 412 Precondition Failed 413 Payload Too Large 414 URI Too Long 415 Unsupported Media Type 416 Range Not Satisfiable 422 Unprocessable Entity 494 Request Header Too Large 495 Cert Error 496 No Cert 497 HTTP to HTTPS
		* 5xx

		500 Internal Server Error 501 Not Implemented 502 Bad Gateway 503 Service Unavailable 504 Gateway Timeout 507 Insufficient Storage

		Example: `200` 
	* `headers` - Adds headers defined in this array for HTTP response.

		Hop-by-hop headers are not allowed to be set:
		* Connection
		* Keep-Alive
		* Proxy-Authenticate
		* Proxy-Authorization
		* TE
		* Trailer
		* Transfer-Encoding
		* Upgrade 
		* `name` - The name of the header field.
		* `value` - The value of the header field.
	* `name` - Action name. Can be used to reference the action.
	* `type` - 
		* **CHECK** is a non-terminating action that does not stop the execution of rules in current module, just emits a log message documenting result of rule execution.
		* **ALLOW** is a non-terminating action which upon matching rule skips all remaining rules in the current module.
		* **RETURN_HTTP_RESPONSE** is a terminating action which is executed immediately, returns a defined HTTP response. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - WebAppFirewallPolicy display name, can be renamed.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the WebAppFirewallPolicy.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in FAILED state. 
* `request_access_control` - Module that allows inspection of HTTP request properties and to return a defined HTTP response. In this module, rules with the name 'Default Action' are not allowed, since this name is reserved for default action logs. 
	* `default_action_name` - References an default Action to take if no AccessControlRule was matched. Allowed action types:
		* **ALLOW** continues execution of other modules and their rules.
		* **RETURN_HTTP_RESPONSE** terminates further execution of modules and rules and returns defined HTTP response. 
	* `rules` - Ordered list of AccessControlRules. Rules are executed in order of appearance in this array.
		* `action_name` - References action by name from actions defined in WebAppFirewallPolicy.
		* `condition` - An expression that determines whether or not the rule action should be executed. 
		* `condition_language` - The language used to parse condition from field `condition`. Available languages:
			* **JMESPATH** an extended JMESPath language syntax. 
		* `name` - Rule name. Must be unique within the module.
		* `type` - Type of WebAppFirewallPolicyRule.
* `request_protection` - Module that allows to enable OCI-managed protection capabilities for incoming HTTP requests.
	* `rules` - Ordered list of ProtectionRules. Rules are executed in order of appearance in this array. ProtectionRules in this array can only use protection cCapabilities of REQUEST_PROTECTION_CAPABILITY type. 
		* `action_name` - References action by name from actions defined in WebAppFirewallPolicy.
		* `condition` - An expression that determines whether or not the rule action should be executed. 
		* `condition_language` - The language used to parse condition from field `condition`. Available languages:
			* **JMESPATH** an extended JMESPath language syntax. 
		* `name` - Rule name. Must be unique within the module.
		* `protection_capabilities` - An ordered list that references OCI-managed protection capabilities. Referenced protection capabilities are executed in order of appearance. The array cannot contain entries with the same pair of capability key and version more than once. 
			* `action_name` - Override action to take if capability was triggered, defined in Protection Rule for this capability. Only actions of type CHECK are allowed. 
			* `collaborative_action_threshold` - The minimum sum of weights of associated collaborative protection capabilities that have triggered which must be reached in order for _this_ capability to trigger. This field is ignored for non-collaborative capabilities. 
			* `collaborative_weights` - Explicit weight values to use for associated collaborative protection capabilities. 
				* `key` - Unique key of collaborative capability for which weight will be overridden.
				* `weight` - The value of weight to set.
			* `exclusions` - Identifies specific HTTP message parameters to exclude from inspection by a protection capability. 
				* `args` - List of URL query parameter values from form-urlencoded XML, JSON, AMP, or POST payloads to exclude from inspecting. Example: If we have query parameter 'argumentName=argumentValue' and args=['argumentName'], both 'argumentName' and 'argumentValue' will not be inspected. 
				* `request_cookies` - List of HTTP request cookie values (by cookie name) to exclude from inspecting. Example: If we have cookie 'cookieName=cookieValue' and requestCookies=['cookieName'], both 'cookieName' and 'cookieValue' will not be inspected. 
			* `key` - Unique key of referenced protection capability.
			* `version` - Version of referenced protection capability.
		* `protection_capability_settings` - Settings for protection capabilities 
			* `allowed_http_methods` - List of allowed HTTP methods. Each value as a RFC7230 formated token string. Used in protection capability 911100: Restrict HTTP Request Methods. 
			* `max_http_request_header_length` - Maximum allowed length of headers in an HTTP request. Used in protection capability: 9200024: Limit length of request header size. 
			* `max_http_request_headers` - Maximum number of headers allowed in an HTTP request. Used in protection capability 9200014: Limit Number of Request Headers. 
			* `max_number_of_arguments` - Maximum number of arguments allowed. Used in protection capability 920380: Number of Arguments Limits. 
			* `max_single_argument_length` - Maximum allowed length of a single argument. Used in protection capability 920370: Limit argument value length. 
			* `max_total_argument_length` - Maximum allowed total length of all arguments. Used in protection capability 920390: Limit arguments total length. 
		* `type` - Type of WebAppFirewallPolicyRule.
* `request_rate_limiting` - Module that allows inspection of HTTP connection properties and to limit requests frequency for a given key.
	* `rules` - Ordered list of RequestRateLimitingRules. Rules are executed in order of appearance in this array. 
		* `action_name` - References action by name from actions defined in WebAppFirewallPolicy.
		* `condition` - An expression that determines whether or not the rule action should be executed. 
		* `condition_language` - The language used to parse condition from field `condition`. Available languages:
			* **JMESPATH** an extended JMESPath language syntax. 
		* `configurations` - Rate Limiting Configurations. Each configuration counts requests towards its own `requestsLimit`. 
			* `action_duration_in_seconds` - Duration of block action application in seconds when `requestsLimit` is reached. Optional and can be 0 (no block duration).
			* `period_in_seconds` - Evaluation period in seconds.
			* `requests_limit` - Requests allowed per evaluation period.
		* `name` - Rule name. Must be unique within the module.
		* `type` - Type of WebAppFirewallPolicyRule.
* `response_access_control` - Module that allows inspection of HTTP response properties and to return a defined HTTP response.
	* `rules` - Ordered list of AccessControlRules. Rules are executed in order of appearance in this array.
		* `action_name` - References action by name from actions defined in WebAppFirewallPolicy.
		* `condition` - An expression that determines whether or not the rule action should be executed. 
		* `condition_language` - The language used to parse condition from field `condition`. Available languages:
			* **JMESPATH** an extended JMESPath language syntax. 
		* `name` - Rule name. Must be unique within the module.
		* `type` - Type of WebAppFirewallPolicyRule.
* `response_protection` - Module that allows to enable OCI-managed protection capabilities for HTTP responses.
	* `rules` - Ordered list of ProtectionRules. Rules are executed in order of appearance in this array. ProtectionRules in this array can only use protection capabilities of RESPONSE_PROTECTION_CAPABILITY type. 
		* `action_name` - References action by name from actions defined in WebAppFirewallPolicy.
		* `condition` - An expression that determines whether or not the rule action should be executed. 
		* `condition_language` - The language used to parse condition from field `condition`. Available languages:
			* **JMESPATH** an extended JMESPath language syntax. 
		* `name` - Rule name. Must be unique within the module.
		* `protection_capabilities` - An ordered list that references OCI-managed protection capabilities. Referenced protection capabilities are executed in order of appearance. The array cannot contain entries with the same pair of capability key and version more than once. 
			* `action_name` - Override action to take if capability was triggered, defined in Protection Rule for this capability. Only actions of type CHECK are allowed. 
			* `collaborative_action_threshold` - The minimum sum of weights of associated collaborative protection capabilities that have triggered which must be reached in order for _this_ capability to trigger. This field is ignored for non-collaborative capabilities. 
			* `collaborative_weights` - Explicit weight values to use for associated collaborative protection capabilities. 
				* `key` - Unique key of collaborative capability for which weight will be overridden.
				* `weight` - The value of weight to set.
			* `exclusions` - Identifies specific HTTP message parameters to exclude from inspection by a protection capability. 
				* `args` - List of URL query parameter values from form-urlencoded XML, JSON, AMP, or POST payloads to exclude from inspecting. Example: If we have query parameter 'argumentName=argumentValue' and args=['argumentName'], both 'argumentName' and 'argumentValue' will not be inspected. 
				* `request_cookies` - List of HTTP request cookie values (by cookie name) to exclude from inspecting. Example: If we have cookie 'cookieName=cookieValue' and requestCookies=['cookieName'], both 'cookieName' and 'cookieValue' will not be inspected. 
			* `key` - Unique key of referenced protection capability.
			* `version` - Version of referenced protection capability.
		* `protection_capability_settings` - Settings for protection capabilities 
			* `allowed_http_methods` - List of allowed HTTP methods. Each value as a RFC7230 formated token string. Used in protection capability 911100: Restrict HTTP Request Methods. 
			* `max_http_request_header_length` - Maximum allowed length of headers in an HTTP request. Used in protection capability: 9200024: Limit length of request header size. 
			* `max_http_request_headers` - Maximum number of headers allowed in an HTTP request. Used in protection capability 9200014: Limit Number of Request Headers. 
			* `max_number_of_arguments` - Maximum number of arguments allowed. Used in protection capability 920380: Number of Arguments Limits. 
			* `max_single_argument_length` - Maximum allowed length of a single argument. Used in protection capability 920370: Limit argument value length. 
			* `max_total_argument_length` - Maximum allowed total length of all arguments. Used in protection capability 920390: Limit arguments total length. 
		* `type` - Type of WebAppFirewallPolicyRule.
* `state` - The current state of the WebAppFirewallPolicy.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the WebAppFirewallPolicy was created. An RFC3339 formatted datetime string.
* `time_updated` - The time the WebAppFirewallPolicy was updated. An RFC3339 formatted datetime string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Web App Firewall Policy
	* `update` - (Defaults to 20 minutes), when updating the Web App Firewall Policy
	* `delete` - (Defaults to 20 minutes), when destroying the Web App Firewall Policy


## Import

WebAppFirewallPolicies can be imported using the `id`, e.g.

```
$ terraform import oci_waf_web_app_firewall_policy.test_web_app_firewall_policy "id"
```

