---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_waas_waas_policy"
sidebar_current: "docs-oci-resource-waas-waas_policy"
description: |-
  Provides the Waas Policy resource in Oracle Cloud Infrastructure Waas service
---

# oci_waas_waas_policy
This resource provides the Waas Policy resource in Oracle Cloud Infrastructure Waas service.

Creates a new Web Application Acceleration and Security (WAAS) policy in the specified compartment. A WAAS policy must be established before creating Web Application Firewall (WAF) rules. To use WAF rules, your web application's origin servers must defined in the `WaasPolicy` schema.

A domain name must be specified when creating a WAAS policy. The domain name should be different from the origins specified in your `WaasPolicy`. Once domain name is entered and stored, it is unchangeable.

Use the record data returned in the `cname` field of the `WaasPolicy` object to create a CNAME record in your DNS configuration that will direct your domain's traffic through the WAF.

For the purposes of access control, you must provide the OCID of the compartment where you want the service to reside. For information about access control and compartments, see [Overview of the IAM Service](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm).

You must specify a display name and domain for the WAAS policy. The display name does not have to be unique and can be changed. The domain name should be different from every origin specified in `WaasPolicy`.

All Oracle Cloud Infrastructure resources, including WAAS policies, receive a unique, Oracle-assigned ID called an Oracle Cloud Identifier (OCID). When a resource is created, you can find its OCID in the response. You can also retrieve a resource's OCID by using a list API operation for that resource type, or by viewing the resource in the Console. Fore more information, see [Resource Identifiers](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).

**Note:** After sending the POST request, the new object's state will temporarily be `CREATING`. Ensure that the resource's state has changed to `ACTIVE` before use.

## Example Usage

```hcl
resource "oci_waas_waas_policy" "test_waas_policy" {
	#Required
	compartment_id = "${var.compartment_id}"
	domain = "${var.waas_policy_domain}"

	#Optional
	additional_domains = "${var.waas_policy_additional_domains}"
	defined_tags = "${var.waas_policy_defined_tags}"
	display_name = "${var.waas_policy_display_name}"
	freeform_tags = "${var.waas_policy_freeform_tags}"
	origins {
		#Required
		uri = "${var.waas_policy_origins_uri}"

		#Optional
		custom_headers {
			#Required
			name = "${var.waas_policy_origins_custom_headers_name}"
			value = "${var.waas_policy_origins_custom_headers_value}"
		}
		http_port = "${var.waas_policy_origins_http_port}"
		https_port = "${var.waas_policy_origins_https_port}"
	}
	policy_config {

		#Optional
		certificate_id = "${oci_waas_certificate.test_certificate.id}"
		is_https_enabled = "${var.waas_policy_policy_config_is_https_enabled}"
		is_https_forced = "${var.waas_policy_policy_config_is_https_forced}"
	}
	waf_config {

		#Optional
		access_rules {
			#Required
			action = "${var.waas_policy_waf_config_access_rules_action}"
			criteria {
				#Required
				condition = "${var.waas_policy_waf_config_access_rules_criteria_condition}"
				value = "${var.waas_policy_waf_config_access_rules_criteria_value}"
			}
			name = "${var.waas_policy_waf_config_access_rules_name}"

			#Optional
			block_action = "${var.waas_policy_waf_config_access_rules_block_action}"
			block_error_page_code = "${var.waas_policy_waf_config_access_rules_block_error_page_code}"
			block_error_page_description = "${var.waas_policy_waf_config_access_rules_block_error_page_description}"
			block_error_page_message = "${var.waas_policy_waf_config_access_rules_block_error_page_message}"
			block_response_code = "${var.waas_policy_waf_config_access_rules_block_response_code}"
		}
		address_rate_limiting {
			#Required
			is_enabled = "${var.waas_policy_waf_config_address_rate_limiting_is_enabled}"

			#Optional
			allowed_rate_per_address = "${var.waas_policy_waf_config_address_rate_limiting_allowed_rate_per_address}"
			block_response_code = "${var.waas_policy_waf_config_address_rate_limiting_block_response_code}"
			max_delayed_count_per_address = "${var.waas_policy_waf_config_address_rate_limiting_max_delayed_count_per_address}"
		}
		captchas {
			#Required
			failure_message = "${var.waas_policy_waf_config_captchas_failure_message}"
			session_expiration_in_seconds = "${var.waas_policy_waf_config_captchas_session_expiration_in_seconds}"
			submit_label = "${var.waas_policy_waf_config_captchas_submit_label}"
			title = "${var.waas_policy_waf_config_captchas_title}"
			url = "${var.waas_policy_waf_config_captchas_url}"

			#Optional
			footer_text = "${var.waas_policy_waf_config_captchas_footer_text}"
			header_text = "${var.waas_policy_waf_config_captchas_header_text}"
		}
		device_fingerprint_challenge {
			#Required
			is_enabled = "${var.waas_policy_waf_config_device_fingerprint_challenge_is_enabled}"

			#Optional
			action = "${var.waas_policy_waf_config_device_fingerprint_challenge_action}"
			action_expiration_in_seconds = "${var.waas_policy_waf_config_device_fingerprint_challenge_action_expiration_in_seconds}"
			challenge_settings {

				#Optional
				block_action = "${var.waas_policy_waf_config_device_fingerprint_challenge_challenge_settings_block_action}"
				block_error_page_code = "${var.waas_policy_waf_config_device_fingerprint_challenge_challenge_settings_block_error_page_code}"
				block_error_page_description = "${var.waas_policy_waf_config_device_fingerprint_challenge_challenge_settings_block_error_page_description}"
				block_error_page_message = "${var.waas_policy_waf_config_device_fingerprint_challenge_challenge_settings_block_error_page_message}"
				block_response_code = "${var.waas_policy_waf_config_device_fingerprint_challenge_challenge_settings_block_response_code}"
				captcha_footer = "${var.waas_policy_waf_config_device_fingerprint_challenge_challenge_settings_captcha_footer}"
				captcha_header = "${var.waas_policy_waf_config_device_fingerprint_challenge_challenge_settings_captcha_header}"
				captcha_submit_label = "${var.waas_policy_waf_config_device_fingerprint_challenge_challenge_settings_captcha_submit_label}"
				captcha_title = "${var.waas_policy_waf_config_device_fingerprint_challenge_challenge_settings_captcha_title}"
			}
			failure_threshold = "${var.waas_policy_waf_config_device_fingerprint_challenge_failure_threshold}"
			failure_threshold_expiration_in_seconds = "${var.waas_policy_waf_config_device_fingerprint_challenge_failure_threshold_expiration_in_seconds}"
			max_address_count = "${var.waas_policy_waf_config_device_fingerprint_challenge_max_address_count}"
			max_address_count_expiration_in_seconds = "${var.waas_policy_waf_config_device_fingerprint_challenge_max_address_count_expiration_in_seconds}"
		}
		human_interaction_challenge {
			#Required
			is_enabled = "${var.waas_policy_waf_config_human_interaction_challenge_is_enabled}"

			#Optional
			action = "${var.waas_policy_waf_config_human_interaction_challenge_action}"
			action_expiration_in_seconds = "${var.waas_policy_waf_config_human_interaction_challenge_action_expiration_in_seconds}"
			challenge_settings {

				#Optional
				block_action = "${var.waas_policy_waf_config_human_interaction_challenge_challenge_settings_block_action}"
				block_error_page_code = "${var.waas_policy_waf_config_human_interaction_challenge_challenge_settings_block_error_page_code}"
				block_error_page_description = "${var.waas_policy_waf_config_human_interaction_challenge_challenge_settings_block_error_page_description}"
				block_error_page_message = "${var.waas_policy_waf_config_human_interaction_challenge_challenge_settings_block_error_page_message}"
				block_response_code = "${var.waas_policy_waf_config_human_interaction_challenge_challenge_settings_block_response_code}"
				captcha_footer = "${var.waas_policy_waf_config_human_interaction_challenge_challenge_settings_captcha_footer}"
				captcha_header = "${var.waas_policy_waf_config_human_interaction_challenge_challenge_settings_captcha_header}"
				captcha_submit_label = "${var.waas_policy_waf_config_human_interaction_challenge_challenge_settings_captcha_submit_label}"
				captcha_title = "${var.waas_policy_waf_config_human_interaction_challenge_challenge_settings_captcha_title}"
			}
			failure_threshold = "${var.waas_policy_waf_config_human_interaction_challenge_failure_threshold}"
			failure_threshold_expiration_in_seconds = "${var.waas_policy_waf_config_human_interaction_challenge_failure_threshold_expiration_in_seconds}"
			interaction_threshold = "${var.waas_policy_waf_config_human_interaction_challenge_interaction_threshold}"
			recording_period_in_seconds = "${var.waas_policy_waf_config_human_interaction_challenge_recording_period_in_seconds}"
			set_http_header {
				#Required
				name = "${var.waas_policy_waf_config_human_interaction_challenge_set_http_header_name}"
				value = "${var.waas_policy_waf_config_human_interaction_challenge_set_http_header_value}"
			}
		}
		js_challenge {
			#Required
			is_enabled = "${var.waas_policy_waf_config_js_challenge_is_enabled}"

			#Optional
			action = "${var.waas_policy_waf_config_js_challenge_action}"
			action_expiration_in_seconds = "${var.waas_policy_waf_config_js_challenge_action_expiration_in_seconds}"
			challenge_settings {

				#Optional
				block_action = "${var.waas_policy_waf_config_js_challenge_challenge_settings_block_action}"
				block_error_page_code = "${var.waas_policy_waf_config_js_challenge_challenge_settings_block_error_page_code}"
				block_error_page_description = "${var.waas_policy_waf_config_js_challenge_challenge_settings_block_error_page_description}"
				block_error_page_message = "${var.waas_policy_waf_config_js_challenge_challenge_settings_block_error_page_message}"
				block_response_code = "${var.waas_policy_waf_config_js_challenge_challenge_settings_block_response_code}"
				captcha_footer = "${var.waas_policy_waf_config_js_challenge_challenge_settings_captcha_footer}"
				captcha_header = "${var.waas_policy_waf_config_js_challenge_challenge_settings_captcha_header}"
				captcha_submit_label = "${var.waas_policy_waf_config_js_challenge_challenge_settings_captcha_submit_label}"
				captcha_title = "${var.waas_policy_waf_config_js_challenge_challenge_settings_captcha_title}"
			}
			failure_threshold = "${var.waas_policy_waf_config_js_challenge_failure_threshold}"
			set_http_header {
				#Required
				name = "${var.waas_policy_waf_config_js_challenge_set_http_header_name}"
				value = "${var.waas_policy_waf_config_js_challenge_set_http_header_value}"
			}
		}
		origin = "${var.waas_policy_waf_config_origin}"
		protection_settings {

			#Optional
			allowed_http_methods = "${var.waas_policy_waf_config_protection_settings_allowed_http_methods}"
			block_action = "${var.waas_policy_waf_config_protection_settings_block_action}"
			block_error_page_code = "${var.waas_policy_waf_config_protection_settings_block_error_page_code}"
			block_error_page_description = "${var.waas_policy_waf_config_protection_settings_block_error_page_description}"
			block_error_page_message = "${var.waas_policy_waf_config_protection_settings_block_error_page_message}"
			block_response_code = "${var.waas_policy_waf_config_protection_settings_block_response_code}"
			is_response_inspected = "${var.waas_policy_waf_config_protection_settings_is_response_inspected}"
			max_argument_count = "${var.waas_policy_waf_config_protection_settings_max_argument_count}"
			max_name_length_per_argument = "${var.waas_policy_waf_config_protection_settings_max_name_length_per_argument}"
			max_response_size_in_ki_b = "${var.waas_policy_waf_config_protection_settings_max_response_size_in_ki_b}"
			max_total_name_length_of_arguments = "${var.waas_policy_waf_config_protection_settings_max_total_name_length_of_arguments}"
			media_types = "${var.waas_policy_waf_config_protection_settings_media_types}"
			recommendations_period_in_days = "${var.waas_policy_waf_config_protection_settings_recommendations_period_in_days}"
		}
		whitelists {
			#Required
			addresses = "${var.waas_policy_waf_config_whitelists_addresses}"
			name = "${var.waas_policy_waf_config_whitelists_name}"
		}
	}
}
```

## Argument Reference

The following arguments are supported:

* `additional_domains` - (Optional) (Updatable) An array of additional domains for the specified web application.
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to create the WAAS policy.
* `defined_tags` - (Optional) (Updatable) A key-value pair with a defined schema that restricts the values of tags. These predefined keys are scoped to namespaces.
* `display_name` - (Optional) (Updatable) A user-friendly name for the WAAS policy. The name is can be changed and does not need to be unique.
* `domain` - (Required) The web application domain that the WAAS policy protects.
* `freeform_tags` - (Optional) (Updatable) A simple key-value pair without any defined schema.
* `origins` - (Optional) (Updatable) A map of host to origin for the web application. The key should be a customer friendly name for the host, ex. primary, secondary, etc.
	* `custom_headers` - (Optional) (Updatable) A list of HTTP headers to forward to your origin.
		* `name` - (Required) (Updatable) The name of the header.
		* `value` - (Required) (Updatable) The value of the header.
	* `http_port` - (Optional) (Updatable) The HTTP port on the origin that the web application listens on. If unspecified, defaults to `80`.
	* `https_port` - (Optional) (Updatable) The HTTPS port on the origin that the web application listens on. If unspecified, defaults to `443`.
	* `uri` - (Required) (Updatable) The URI of the origin. Does not support paths. Port numbers should be specified in the `httpPort` and `httpsPort` fields.
* `policy_config` - (Optional) (Updatable) 
	* `certificate_id` - (Optional) (Updatable) The OCID of the SSL certificate to use if HTTPS is supported.
	* `is_https_enabled` - (Optional) (Updatable) Enable or disable HTTPS support. If true, a `certificateId` is required. If unspecified, defaults to `false`.
	* `is_https_forced` - (Optional) (Updatable) Force HTTP to HTTPS redirection. If unspecified, defaults to `false`.
* `waf_config` - (Optional) (Updatable) 
	* `access_rules` - (Optional) (Updatable) The access rules applied to the Web Application Firewall. Access rules allow custom content access policies to be defined and `ALLOW`, `DETECT`, or `BLOCK` actions to be taken on a request when specified criteria are met.
		* `action` - (Required) (Updatable) The action to take when the access criteria are met for a rule. If unspecified, defaults to `ALLOW`.
		* `block_action` - (Optional) (Updatable) The method used to block requests if `action` is set to `BLOCK` and the access criteria are met. If unspecified, defaults to `SET_RESPONSE_CODE`.
		* `block_error_page_code` - (Optional) (Updatable) The error code to show on the error page when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_ERROR_PAGE`, and the access criteria are met. If unspecified, defaults to 'Access rules'.
		* `block_error_page_description` - (Optional) (Updatable) The description text to show on the error page when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_ERROR_PAGE`, and the access criteria are met. If unspecified, defaults to 'Access blocked by website owner. Please contact support.'
		* `block_error_page_message` - (Optional) (Updatable) The message to show on the error page when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_ERROR_PAGE`, and the access criteria are met. If unspecified, defaults to 'Access to the website is blocked.'
		* `block_response_code` - (Optional) (Updatable) The response status code to return when `action` is set to `BLOCK`, `blockAction` is set to `SET_RESPONSE_CODE`, and the access criteria are met. If unspecified, defaults to `403`.
		* `criteria` - (Required) (Updatable) The list of access rule criteria.
			* `condition` - (Required) (Updatable) The criteria the access rule uses to determine if action should be taken on a request.
				* **URL_IS:** Matches if the concatenation of request URL path and query is identical to the contents of the `value` field.
				* **URL_IS_NOT:** Matches if the concatenation of request URL path and query is not identical to the contents of the `value` field.
				* **URL_STARTS_WITH:** Matches if the concatenation of request URL path and query starts with the contents of the `value` field.
				* **URL_PART_ENDS_WITH:** Matches if the concatenation of request URL path and query ends with the contents of the `value` field.
				* **URL_PART_CONTAINS:** Matches if the concatenation of request URL path and query contains the contents of the `value` field.
				* **URL_REGEX:** Matches if the request is described by the regular expression in the `value` field.
				* **IP_IS:** Matches if the request originates from an IP address in the `value` field.
				* **IP_IS_NOT:** Matches if the request does not originate from an IP address in the `value` field.
				* **HTTP_HEADER_CONTAINS:** The HTTP_HEADER_CONTAINS criteria is defined using a compound value separated by a colon: a header field name and a header field value. `host:test.example.com` is an example of a criteria value where `host` is the header field name and `test.example.com` is the header field value. A request matches when the header field name is a case insensitive match and the header field value is a case insensitive, substring match. *Example:* With a criteria value of `host:test.example.com`, where `host` is the name of the field and `test.example.com` is the value of the host field, a request with the header values, `Host: www.test.example.com` will match, where as a request with header values of `host: www.example.com` or `host: test.sub.example.com` will not match.
				* **COUNTRY_IS:** Matches if the request originates from a country in the `value` field. Country codes are in ISO 3166-1 alpha-2 format. For a list of codes, see [ISO's website](https://www.iso.org/obp/ui/#search/code/).
				* **COUNTRY_IS_NOT:** Matches if the request does not originate from a country in the `value` field. Country codes are in ISO 3166-1 alpha-2 format. For a list of codes, see [ISO's website](https://www.iso.org/obp/ui/#search/code/).
				* **USER_AGENT_IS:** Matches if the requesting user agent is identical to the contents of the `value` field. Example: `Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:35.0) Gecko/20100101 Firefox/35.0`
				* **USER_AGENT_IS_NOT:** Matches if the requesting user agent is not identical to the contents of the `value` field. Example: `Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:35.0) Gecko/20100101 Firefox/35.0`
			* `value` - (Required) (Updatable) The criteria value.
		* `name` - (Required) (Updatable) The unique name of the access rule.
	* `address_rate_limiting` - (Optional) (Updatable) The settings used to limit the number of requests from an IP address.
		* `allowed_rate_per_address` - (Optional) (Updatable) The number of allowed requests per second from one IP address. If unspecified, defaults to `1`.
		* `block_response_code` - (Optional) (Updatable) The response status code returned when a request is blocked. If unspecified, defaults to `503`.
		* `is_enabled` - (Required) (Updatable) Enables or disables the address rate limiting Web Application Firewall feature.
		* `max_delayed_count_per_address` - (Optional) (Updatable) The maximum number of requests allowed to be queued before subsequent requests are dropped. If unspecified, defaults to `10`.
	* `captchas` - (Optional) (Updatable) A list of CAPTCHA challenge settings. CAPTCHAs challenge requests to ensure a human is attempting to reach the specified URL and not a bot.
		* `failure_message` - (Required) (Updatable) The text to show when incorrect CAPTCHA text is entered. If unspecified, defaults to `The CAPTCHA was incorrect. Try again.`
		* `footer_text` - (Optional) (Updatable) The text to show in the footer when showing a CAPTCHA challenge. If unspecified, defaults to 'Enter the letters and numbers as they are shown in the image above.'
		* `header_text` - (Optional) (Updatable) The text to show in the header when showing a CAPTCHA challenge. If unspecified, defaults to 'We have detected an increased number of attempts to access this website. To help us keep this site secure, please let us know that you are not a robot by entering the text from the image below.'
		* `session_expiration_in_seconds` - (Required) (Updatable) The amount of time before the CAPTCHA expires, in seconds. If unspecified, defaults to `300`.
		* `submit_label` - (Required) (Updatable) The text to show on the label of the CAPTCHA challenge submit button. If unspecified, defaults to `Yes, I am human`.
		* `title` - (Required) (Updatable) The title used when displaying a CAPTCHA challenge. If unspecified, defaults to `Are you human?`
		* `url` - (Required) (Updatable) The unique URL path at which to show the CAPTCHA challenge.
	* `device_fingerprint_challenge` - (Optional) (Updatable) The device fingerprint challenge settings. Blocks bots based on unique device fingerprint information.
		* `action` - (Optional) (Updatable) The action to take on requests from detected bots. If unspecified, defaults to `DETECT`.
		* `action_expiration_in_seconds` - (Optional) (Updatable) The number of seconds between challenges for the same IP address. If unspecified, defaults to `60`.
		* `challenge_settings` - (Optional) (Updatable) 
			* `block_action` - (Optional) (Updatable) The method used to block requests that fail the challenge, if `action` is set to `BLOCK`. If unspecified, defaults to `SHOW_ERROR_PAGE`.
			* `block_error_page_code` - (Optional) (Updatable) The error code to show on the error page when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_ERROR_PAGE` and the request is blocked. If unspecified, defaults to `403`.
			* `block_error_page_description` - (Optional) (Updatable) The description text to show on the error page when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_ERROR_PAGE`, and the request is blocked. If unspecified, defaults to `Access blocked by website owner. Please contact support.`
			* `block_error_page_message` - (Optional) (Updatable) The message to show on the error page when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_ERROR_PAGE`, and the request is blocked. If unspecified, defaults to `Access to the website is blocked`.
			* `block_response_code` - (Optional) (Updatable) The response status code to return when `action` is set to `BLOCK`, `blockAction` is set to `SET_RESPONSE_CODE` or `SHOW_ERROR_PAGE`, and the request is blocked. If unspecified, defaults to `403`.
			* `captcha_footer` - (Optional) (Updatable) The text to show in the footer when showing a CAPTCHA challenge when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_CAPTCHA`, and the request is blocked. If unspecified, default to `Enter the letters and numbers as they are shown in image above`.
			* `captcha_header` - (Optional) (Updatable) The text to show in the header when showing a CAPTCHA challenge when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_CAPTCHA`, and the request is blocked. If unspecified, defaults to `We have detected an increased number of attempts to access this webapp. To help us keep this webapp secure, please let us know that you are not a robot by entering the text from captcha below.`
			* `captcha_submit_label` - (Optional) (Updatable) The text to show on the label of the CAPTCHA challenge submit button when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_CAPTCHA`, and the request is blocked. If unspecified, defaults to `Yes, I am human`.
			* `captcha_title` - (Optional) (Updatable) The title used when showing a CAPTCHA challenge when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_CAPTCHA`, and the request is blocked. If unspecified, defaults to `Are you human?`
		* `failure_threshold` - (Optional) (Updatable) The number of failed requests allowed before taking action. If unspecified, defaults to `10`.
		* `failure_threshold_expiration_in_seconds` - (Optional) (Updatable) The number of seconds before the failure threshold resets. If unspecified, defaults to `60`.
		* `is_enabled` - (Required) (Updatable) Enables or disables the device fingerprint challenge Web Application Firewall feature.
		* `max_address_count` - (Optional) (Updatable) The maximum number of IP addresses permitted with the same device fingerprint. If unspecified, defaults to `20`.
		* `max_address_count_expiration_in_seconds` - (Optional) (Updatable) The number of seconds before the maximum addresses count resets. If unspecified, defaults to `60`.
	* `human_interaction_challenge` - (Optional) (Updatable) The human interaction challenge settings. Detects natural human interactions such as mouse movements, time on site, and page scrolling to identify bots.
		* `action` - (Optional) (Updatable) The action to take against requests from detected bots. If unspecified, defaults to `DETECT`.
		* `action_expiration_in_seconds` - (Optional) (Updatable) The number of seconds between challenges for the same IP address. If unspecified, defaults to `60`.
		* `challenge_settings` - (Optional) (Updatable) 
			* `block_action` - (Optional) (Updatable) The method used to block requests that fail the challenge, if `action` is set to `BLOCK`. If unspecified, defaults to `SHOW_ERROR_PAGE`.
			* `block_error_page_code` - (Optional) (Updatable) The error code to show on the error page when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_ERROR_PAGE` and the request is blocked. If unspecified, defaults to `403`.
			* `block_error_page_description` - (Optional) (Updatable) The description text to show on the error page when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_ERROR_PAGE`, and the request is blocked. If unspecified, defaults to `Access blocked by website owner. Please contact support.`
			* `block_error_page_message` - (Optional) (Updatable) The message to show on the error page when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_ERROR_PAGE`, and the request is blocked. If unspecified, defaults to `Access to the website is blocked`.
			* `block_response_code` - (Optional) (Updatable) The response status code to return when `action` is set to `BLOCK`, `blockAction` is set to `SET_RESPONSE_CODE` or `SHOW_ERROR_PAGE`, and the request is blocked. If unspecified, defaults to `403`.
			* `captcha_footer` - (Optional) (Updatable) The text to show in the footer when showing a CAPTCHA challenge when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_CAPTCHA`, and the request is blocked. If unspecified, default to `Enter the letters and numbers as they are shown in image above`.
			* `captcha_header` - (Optional) (Updatable) The text to show in the header when showing a CAPTCHA challenge when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_CAPTCHA`, and the request is blocked. If unspecified, defaults to `We have detected an increased number of attempts to access this webapp. To help us keep this webapp secure, please let us know that you are not a robot by entering the text from captcha below.`
			* `captcha_submit_label` - (Optional) (Updatable) The text to show on the label of the CAPTCHA challenge submit button when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_CAPTCHA`, and the request is blocked. If unspecified, defaults to `Yes, I am human`.
			* `captcha_title` - (Optional) (Updatable) The title used when showing a CAPTCHA challenge when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_CAPTCHA`, and the request is blocked. If unspecified, defaults to `Are you human?`
		* `failure_threshold` - (Optional) (Updatable) The number of failed requests before taking action. If unspecified, defaults to `10`.
		* `failure_threshold_expiration_in_seconds` - (Optional) (Updatable) The number of seconds before the failure threshold resets. If unspecified, defaults to  `60`.
		* `interaction_threshold` - (Optional) (Updatable) The number of interactions required to pass the challenge. If unspecified, defaults to `3`.
		* `is_enabled` - (Required) (Updatable) Enables or disables the human interaction challenge Web Application Firewall feature.
		* `recording_period_in_seconds` - (Optional) (Updatable) The number of seconds to record the interactions from the user. If unspecified, defaults to `15`.
		* `set_http_header` - (Optional) (Updatable) Adds an additional HTTP header to requests that fail the challenge before being passed to the origin. Only applicable when the `action` is set to `DETECT`.
			* `name` - (Required) (Updatable) The name of the header.
			* `value` - (Required) (Updatable) The value of the header.
	* `js_challenge` - (Optional) (Updatable) The JavaScript challenge settings. Blocks bots by challenging requests from browsers that have no JavaScript support.
		* `action` - (Optional) (Updatable) The action to take against requests from detected bots. If unspecified, defaults to `DETECT`.
		* `action_expiration_in_seconds` - (Optional) (Updatable) The number of seconds between challenges from the same IP address. If unspecified, defaults to `60`.
		* `challenge_settings` - (Optional) (Updatable) 
			* `block_action` - (Optional) (Updatable) The method used to block requests that fail the challenge, if `action` is set to `BLOCK`. If unspecified, defaults to `SHOW_ERROR_PAGE`.
			* `block_error_page_code` - (Optional) (Updatable) The error code to show on the error page when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_ERROR_PAGE` and the request is blocked. If unspecified, defaults to `403`.
			* `block_error_page_description` - (Optional) (Updatable) The description text to show on the error page when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_ERROR_PAGE`, and the request is blocked. If unspecified, defaults to `Access blocked by website owner. Please contact support.`
			* `block_error_page_message` - (Optional) (Updatable) The message to show on the error page when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_ERROR_PAGE`, and the request is blocked. If unspecified, defaults to `Access to the website is blocked`.
			* `block_response_code` - (Optional) (Updatable) The response status code to return when `action` is set to `BLOCK`, `blockAction` is set to `SET_RESPONSE_CODE` or `SHOW_ERROR_PAGE`, and the request is blocked. If unspecified, defaults to `403`.
			* `captcha_footer` - (Optional) (Updatable) The text to show in the footer when showing a CAPTCHA challenge when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_CAPTCHA`, and the request is blocked. If unspecified, default to `Enter the letters and numbers as they are shown in image above`.
			* `captcha_header` - (Optional) (Updatable) The text to show in the header when showing a CAPTCHA challenge when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_CAPTCHA`, and the request is blocked. If unspecified, defaults to `We have detected an increased number of attempts to access this webapp. To help us keep this webapp secure, please let us know that you are not a robot by entering the text from captcha below.`
			* `captcha_submit_label` - (Optional) (Updatable) The text to show on the label of the CAPTCHA challenge submit button when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_CAPTCHA`, and the request is blocked. If unspecified, defaults to `Yes, I am human`.
			* `captcha_title` - (Optional) (Updatable) The title used when showing a CAPTCHA challenge when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_CAPTCHA`, and the request is blocked. If unspecified, defaults to `Are you human?`
		* `failure_threshold` - (Optional) (Updatable) The number of failed requests before taking action. If unspecified, defaults to `10`.
		* `is_enabled` - (Required) (Updatable) Enables or disables the JavaScript challenge Web Application Firewall feature.
		* `set_http_header` - (Optional) (Updatable) Adds an additional HTTP header to requests that fail the challenge before being passed to the origin. Only applicable when the `action` is set to `DETECT`.
			* `name` - (Required) (Updatable) The name of the header.
			* `value` - (Required) (Updatable) The value of the header.
	* `origin` - (Optional) (Updatable) The key in the map of origins referencing the origin used for the Web Application Firewall. The origin must already be included in `Origins`. Required when creating the `WafConfig` resource, but is not required upon updating the configuration.
	* `protection_settings` - (Optional) (Updatable) The settings applied to protection rules.
		* `allowed_http_methods` - (Optional) (Updatable) The list of allowed HTTP methods. If unspecified, default to `[OPTIONS, GET, HEAD, POST]`. This setting only applies if a corresponding protection rule is enabled, such as the "Restrict HTTP Request Methods" rule (key: 911100).
		* `block_action` - (Optional) (Updatable) If `action` is set to `BLOCK`, this specifies how the traffic is blocked when detected as malicious by a protection rule. If unspecified, defaults to `SET_RESPONSE_CODE`.
		* `block_error_page_code` - (Optional) (Updatable) The error code to show on the error page when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_ERROR_PAGE`, and the traffic is detected as malicious by a protection rule. If unspecified, defaults to `403`.
		* `block_error_page_description` - (Optional) (Updatable) The description text to show on the error page when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_ERROR_PAGE`, and the traffic is detected as malicious by a protection rule. If unspecified, defaults to `Access blocked by website owner. Please contact support.`
		* `block_error_page_message` - (Optional) (Updatable) The message to show on the error page when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_ERROR_PAGE`, and the traffic is detected as malicious by a protection rule. If unspecified, defaults to 'Access to the website is blocked.'
		* `block_response_code` - (Optional) (Updatable) The response code returned when `action` is set to `BLOCK`, `blockAction` is set to `SET_RESPONSE_CODE`, and the traffic is detected as malicious by a protection rule. If unspecified, defaults to `403`.
		* `is_response_inspected` - (Optional) (Updatable) Inspects the response body of origin responses. Can be used to detect leakage of sensitive data. If unspecified, defaults to `false`.

			**Note:** Only origin responses with a Content-Type matching a value in `mediaTypes` will be inspected.
		* `max_argument_count` - (Optional) (Updatable) The maximum number of arguments allowed to be passed to your application before an action is taken. Arguements are query parameters or body parameters in a PUT or POST request. If unspecified, defaults to `255`. This setting only applies if a corresponding protection rule is enabled, such as the "Number of Arguments Limits" rule (key: 960335).  Example: If `maxArgumentCount` to `2` for the Max Number of Arguments protection rule (key: 960335), the following requests would be blocked: `GET /myapp/path?query=one&query=two&query=three` `POST /myapp/path` with Body `{"argument1":"one","argument2":"two","argument3":"three"}`
		* `max_name_length_per_argument` - (Optional) (Updatable) The maximum length allowed for each argument name, in characters. Arguements are query parameters or body parameters in a PUT or POST request. If unspecified, defaults to `400`. This setting only applies if a corresponding protection rule is enabled, such as the "Values Limits" rule (key: 960208).
		* `max_response_size_in_ki_b` - (Optional) (Updatable) The maximum response size to be fully inspected, in binary kilobytes (KiB). Anything over this limit will be partially inspected. If unspecified, defaults to `1024`.
		* `max_total_name_length_of_arguments` - (Optional) (Updatable) The maximum length allowed for the sum of the argument name and value, in characters. Arguements are query parameters or body parameters in a PUT or POST request. If unspecified, defaults to `64000`. This setting only applies if a corresponding protection rule is enabled, such as the "Total Arguments Limits" rule (key: 960341).
		* `media_types` - (Optional) (Updatable) The list of media types to allow for inspection, if `isResponseInspected` is enabled. Only responses with MIME types in this list will be inspected. If unspecified, defaults to `["text/html", "text/plain", "text/xml"]`.

			Supported MIME types include:
			* text/html
			* text/plain
			* text/asp
			* text/css
			* text/x-script
			* application/json
			* text/webviewhtml
			* text/x-java-source
			* application/x-javascript
			* application/javascript
			* application/ecmascript
			* text/javascript
			* text/ecmascript
			* text/x-script.perl
			* text/x-script.phyton
			* application/plain
			* application/xml
			* text/xml
		* `recommendations_period_in_days` - (Optional) (Updatable) The length of time to analyze traffic traffic, in days. After the analysis period, `WafRecommendations` will be populated. If unspecified, defaults to `10`.

			Use `GET /waasPolicies/{waasPolicyId}/wafRecommendations` to view WAF recommendations.
	* `whitelists` - (Optional) (Updatable) A list of IP addresses that bypass the Web Application Firewall.
		* `addresses` - (Required) (Updatable) A set of IP addresses or CIDR notations to include in the whitelist.
		* `name` - (Required) (Updatable) The unique name of the whitelist.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `additional_domains` - An array of additional domains for this web application.
* `cname` - The CNAME record to add to your DNS configuration to route traffic for the domain, and all additional domains, through the WAF.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the WAAS policy's compartment.
* `defined_tags` - A key-value pair with a defined schema that restricts the values of tags. These predefined keys are scoped to namespaces.
* `display_name` - The user-friendly name of the WAAS policy. The name can be changed and does not need to be unique.
* `domain` - The web application domain that the WAAS policy protects.
* `freeform_tags` - A simple key-value pair without any defined schema.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the WAAS policy.
* `origins` - A map of host servers (origins) and their keys for the web application. Origin keys are used to associate origins to specific protection rules. The key should be a user-friendly name for the host. **Examples:** `primary` or `secondary`.
	* `custom_headers` - A list of HTTP headers to forward to your origin.
		* `name` - The name of the header.
		* `value` - The value of the header.
	* `http_port` - The HTTP port on the origin that the web application listens on. If unspecified, defaults to `80`.
	* `https_port` - The HTTPS port on the origin that the web application listens on. If unspecified, defaults to `443`.
	* `uri` - The URI of the origin. Does not support paths. Port numbers should be specified in the `httpPort` and `httpsPort` fields.
* `policy_config` - 
	* `certificate_id` - The OCID of the SSL certificate to use if HTTPS is supported.
	* `is_https_enabled` - Enable or disable HTTPS support. If true, a `certificateId` is required. If unspecified, defaults to `false`.
	* `is_https_forced` - Force HTTP to HTTPS redirection. If unspecified, defaults to `false`.
* `state` - The current lifecycle state of the WAAS policy.
* `time_created` - The date and time the policy was created, expressed in RFC 3339 timestamp format.
* `waf_config` - 
	* `access_rules` - The access rules applied to the Web Application Firewall. Used for defining custom access policies with the combination of `ALLOW`, `DETECT`, and `BLOCK` rules, based on different criteria.
		* `action` - The action to take when the access criteria are met for a rule. If unspecified, defaults to `ALLOW`.
		* `block_action` - The method used to block requests if `action` is set to `BLOCK` and the access criteria are met. If unspecified, defaults to `SET_RESPONSE_CODE`.
		* `block_error_page_code` - The error code to show on the error page when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_ERROR_PAGE`, and the access criteria are met. If unspecified, defaults to 'Access rules'.
		* `block_error_page_description` - The description text to show on the error page when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_ERROR_PAGE`, and the access criteria are met. If unspecified, defaults to 'Access blocked by website owner. Please contact support.'
		* `block_error_page_message` - The message to show on the error page when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_ERROR_PAGE`, and the access criteria are met. If unspecified, defaults to 'Access to the website is blocked.'
		* `block_response_code` - The response status code to return when `action` is set to `BLOCK`, `blockAction` is set to `SET_RESPONSE_CODE`, and the access criteria are met. If unspecified, defaults to `403`.
		* `criteria` - The list of access rule criteria.
			* `condition` - The criteria the access rule uses to determine if action should be taken on a request.
				* **URL_IS:** Matches if the concatenation of request URL path and query is identical to the contents of the `value` field.
				* **URL_IS_NOT:** Matches if the concatenation of request URL path and query is not identical to the contents of the `value` field.
				* **URL_STARTS_WITH:** Matches if the concatenation of request URL path and query starts with the contents of the `value` field.
				* **URL_PART_ENDS_WITH:** Matches if the concatenation of request URL path and query ends with the contents of the `value` field.
				* **URL_PART_CONTAINS:** Matches if the concatenation of request URL path and query contains the contents of the `value` field.
				* **URL_REGEX:** Matches if the request is described by the regular expression in the `value` field.
				* **IP_IS:** Matches if the request originates from an IP address in the `value` field.
				* **IP_IS_NOT:** Matches if the request does not originate from an IP address in the `value` field.
				* **HTTP_HEADER_CONTAINS:** The HTTP_HEADER_CONTAINS criteria is defined using a compound value separated by a colon: a header field name and a header field value. `host:test.example.com` is an example of a criteria value where `host` is the header field name and `test.example.com` is the header field value. A request matches when the header field name is a case insensitive match and the header field value is a case insensitive, substring match. *Example:* With a criteria value of `host:test.example.com`, where `host` is the name of the field and `test.example.com` is the value of the host field, a request with the header values, `Host: www.test.example.com` will match, where as a request with header values of `host: www.example.com` or `host: test.sub.example.com` will not match.
				* **COUNTRY_IS:** Matches if the request originates from a country in the `value` field. Country codes are in ISO 3166-1 alpha-2 format. For a list of codes, see [ISO's website](https://www.iso.org/obp/ui/#search/code/).
				* **COUNTRY_IS_NOT:** Matches if the request does not originate from a country in the `value` field. Country codes are in ISO 3166-1 alpha-2 format. For a list of codes, see [ISO's website](https://www.iso.org/obp/ui/#search/code/).
				* **USER_AGENT_IS:** Matches if the requesting user agent is identical to the contents of the `value` field. Example: `Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:35.0) Gecko/20100101 Firefox/35.0`
				* **USER_AGENT_IS_NOT:** Matches if the requesting user agent is not identical to the contents of the `value` field. Example: `Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:35.0) Gecko/20100101 Firefox/35.0`
			* `value` - The criteria value.
		* `name` - The unique name of the access rule.
	* `address_rate_limiting` - The IP address rate limiting settings used to limit the number of requests from an address.
		* `allowed_rate_per_address` - The number of allowed requests per second from one IP address. If unspecified, defaults to `1`.
		* `block_response_code` - The response status code returned when a request is blocked. If unspecified, defaults to `503`.
		* `is_enabled` - Enables or disables the address rate limiting Web Application Firewall feature.
		* `max_delayed_count_per_address` - The maximum number of requests allowed to be queued before subsequent requests are dropped. If unspecified, defaults to `10`.
	* `captchas` - A list of CAPTCHA challenge settings. These are used to challenge requests with a CAPTCHA to block bots.
		* `failure_message` - The text to show when incorrect CAPTCHA text is entered. If unspecified, defaults to `The CAPTCHA was incorrect. Try again.`
		* `footer_text` - The text to show in the footer when showing a CAPTCHA challenge. If unspecified, defaults to 'Enter the letters and numbers as they are shown in the image above.'
		* `header_text` - The text to show in the header when showing a CAPTCHA challenge. If unspecified, defaults to 'We have detected an increased number of attempts to access this website. To help us keep this site secure, please let us know that you are not a robot by entering the text from the image below.'
		* `session_expiration_in_seconds` - The amount of time before the CAPTCHA expires, in seconds. If unspecified, defaults to `300`.
		* `submit_label` - The text to show on the label of the CAPTCHA challenge submit button. If unspecified, defaults to `Yes, I am human`.
		* `title` - The title used when displaying a CAPTCHA challenge. If unspecified, defaults to `Are you human?`
		* `url` - The unique URL path at which to show the CAPTCHA challenge.
	* `device_fingerprint_challenge` - The device fingerprint challenge settings. Used to detect unique devices based on the device fingerprint information collected in order to block bots.
		* `action` - The action to take on requests from detected bots. If unspecified, defaults to `DETECT`.
		* `action_expiration_in_seconds` - The number of seconds between challenges for the same IP address. If unspecified, defaults to `60`.
		* `challenge_settings` - 
			* `block_action` - The method used to block requests that fail the challenge, if `action` is set to `BLOCK`. If unspecified, defaults to `SHOW_ERROR_PAGE`.
			* `block_error_page_code` - The error code to show on the error page when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_ERROR_PAGE` and the request is blocked. If unspecified, defaults to `403`.
			* `block_error_page_description` - The description text to show on the error page when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_ERROR_PAGE`, and the request is blocked. If unspecified, defaults to `Access blocked by website owner. Please contact support.`
			* `block_error_page_message` - The message to show on the error page when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_ERROR_PAGE`, and the request is blocked. If unspecified, defaults to `Access to the website is blocked`.
			* `block_response_code` - The response status code to return when `action` is set to `BLOCK`, `blockAction` is set to `SET_RESPONSE_CODE` or `SHOW_ERROR_PAGE`, and the request is blocked. If unspecified, defaults to `403`.
			* `captcha_footer` - The text to show in the footer when showing a CAPTCHA challenge when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_CAPTCHA`, and the request is blocked. If unspecified, default to `Enter the letters and numbers as they are shown in image above`.
			* `captcha_header` - The text to show in the header when showing a CAPTCHA challenge when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_CAPTCHA`, and the request is blocked. If unspecified, defaults to `We have detected an increased number of attempts to access this webapp. To help us keep this webapp secure, please let us know that you are not a robot by entering the text from captcha below.`
			* `captcha_submit_label` - The text to show on the label of the CAPTCHA challenge submit button when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_CAPTCHA`, and the request is blocked. If unspecified, defaults to `Yes, I am human`.
			* `captcha_title` - The title used when showing a CAPTCHA challenge when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_CAPTCHA`, and the request is blocked. If unspecified, defaults to `Are you human?`
		* `failure_threshold` - The number of failed requests allowed before taking action. If unspecified, defaults to `10`.
		* `failure_threshold_expiration_in_seconds` - The number of seconds before the failure threshold resets. If unspecified, defaults to `60`.
		* `is_enabled` - Enables or disables the device fingerprint challenge Web Application Firewall feature.
		* `max_address_count` - The maximum number of IP addresses permitted with the same device fingerprint. If unspecified, defaults to `20`.
		* `max_address_count_expiration_in_seconds` - The number of seconds before the maximum addresses count resets. If unspecified, defaults to `60`.
	* `human_interaction_challenge` - The human interaction challenge settings. Used to look for natural human interactions such as mouse movements, time on site, and page scrolling to identify bots.
		* `action` - The action to take against requests from detected bots. If unspecified, defaults to `DETECT`.
		* `action_expiration_in_seconds` - The number of seconds between challenges for the same IP address. If unspecified, defaults to `60`.
		* `challenge_settings` - 
			* `block_action` - The method used to block requests that fail the challenge, if `action` is set to `BLOCK`. If unspecified, defaults to `SHOW_ERROR_PAGE`.
			* `block_error_page_code` - The error code to show on the error page when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_ERROR_PAGE` and the request is blocked. If unspecified, defaults to `403`.
			* `block_error_page_description` - The description text to show on the error page when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_ERROR_PAGE`, and the request is blocked. If unspecified, defaults to `Access blocked by website owner. Please contact support.`
			* `block_error_page_message` - The message to show on the error page when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_ERROR_PAGE`, and the request is blocked. If unspecified, defaults to `Access to the website is blocked`.
			* `block_response_code` - The response status code to return when `action` is set to `BLOCK`, `blockAction` is set to `SET_RESPONSE_CODE` or `SHOW_ERROR_PAGE`, and the request is blocked. If unspecified, defaults to `403`.
			* `captcha_footer` - The text to show in the footer when showing a CAPTCHA challenge when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_CAPTCHA`, and the request is blocked. If unspecified, default to `Enter the letters and numbers as they are shown in image above`.
			* `captcha_header` - The text to show in the header when showing a CAPTCHA challenge when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_CAPTCHA`, and the request is blocked. If unspecified, defaults to `We have detected an increased number of attempts to access this webapp. To help us keep this webapp secure, please let us know that you are not a robot by entering the text from captcha below.`
			* `captcha_submit_label` - The text to show on the label of the CAPTCHA challenge submit button when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_CAPTCHA`, and the request is blocked. If unspecified, defaults to `Yes, I am human`.
			* `captcha_title` - The title used when showing a CAPTCHA challenge when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_CAPTCHA`, and the request is blocked. If unspecified, defaults to `Are you human?`
		* `failure_threshold` - The number of failed requests before taking action. If unspecified, defaults to `10`.
		* `failure_threshold_expiration_in_seconds` - The number of seconds before the failure threshold resets. If unspecified, defaults to  `60`.
		* `interaction_threshold` - The number of interactions required to pass the challenge. If unspecified, defaults to `3`.
		* `is_enabled` - Enables or disables the human interaction challenge Web Application Firewall feature.
		* `recording_period_in_seconds` - The number of seconds to record the interactions from the user. If unspecified, defaults to `15`.
		* `set_http_header` - Adds an additional HTTP header to requests that fail the challenge before being passed to the origin. Only applicable when the `action` is set to `DETECT`.
			* `name` - The name of the header.
			* `value` - The value of the header.
	* `js_challenge` - The JavaScript challenge settings. Used to challenge requests with a JavaScript challenge and take the action if a browser has no JavaScript support in order to block bots.
		* `action` - The action to take against requests from detected bots. If unspecified, defaults to `DETECT`.
		* `action_expiration_in_seconds` - The number of seconds between challenges from the same IP address. If unspecified, defaults to `60`.
		* `challenge_settings` - 
			* `block_action` - The method used to block requests that fail the challenge, if `action` is set to `BLOCK`. If unspecified, defaults to `SHOW_ERROR_PAGE`.
			* `block_error_page_code` - The error code to show on the error page when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_ERROR_PAGE` and the request is blocked. If unspecified, defaults to `403`.
			* `block_error_page_description` - The description text to show on the error page when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_ERROR_PAGE`, and the request is blocked. If unspecified, defaults to `Access blocked by website owner. Please contact support.`
			* `block_error_page_message` - The message to show on the error page when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_ERROR_PAGE`, and the request is blocked. If unspecified, defaults to `Access to the website is blocked`.
			* `block_response_code` - The response status code to return when `action` is set to `BLOCK`, `blockAction` is set to `SET_RESPONSE_CODE` or `SHOW_ERROR_PAGE`, and the request is blocked. If unspecified, defaults to `403`.
			* `captcha_footer` - The text to show in the footer when showing a CAPTCHA challenge when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_CAPTCHA`, and the request is blocked. If unspecified, default to `Enter the letters and numbers as they are shown in image above`.
			* `captcha_header` - The text to show in the header when showing a CAPTCHA challenge when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_CAPTCHA`, and the request is blocked. If unspecified, defaults to `We have detected an increased number of attempts to access this webapp. To help us keep this webapp secure, please let us know that you are not a robot by entering the text from captcha below.`
			* `captcha_submit_label` - The text to show on the label of the CAPTCHA challenge submit button when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_CAPTCHA`, and the request is blocked. If unspecified, defaults to `Yes, I am human`.
			* `captcha_title` - The title used when showing a CAPTCHA challenge when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_CAPTCHA`, and the request is blocked. If unspecified, defaults to `Are you human?`
		* `failure_threshold` - The number of failed requests before taking action. If unspecified, defaults to `10`.
		* `is_enabled` - Enables or disables the JavaScript challenge Web Application Firewall feature.
		* `set_http_header` - Adds an additional HTTP header to requests that fail the challenge before being passed to the origin. Only applicable when the `action` is set to `DETECT`.
			* `name` - The name of the header.
			* `value` - The value of the header.
	* `origin` - The key in the map of origins referencing the origin used for the Web Application Firewall. The origin must already be included in `Origins`. Required when creating the `WafConfig` resource, but not on update.
	* `protection_settings` - The settings to apply to protection rules.
		* `allowed_http_methods` - The list of allowed HTTP methods. If unspecified, default to `[OPTIONS, GET, HEAD, POST]`. This setting only applies if a corresponding protection rule is enabled, such as the "Restrict HTTP Request Methods" rule (key: 911100).
		* `block_action` - If `action` is set to `BLOCK`, this specifies how the traffic is blocked when detected as malicious by a protection rule. If unspecified, defaults to `SET_RESPONSE_CODE`.
		* `block_error_page_code` - The error code to show on the error page when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_ERROR_PAGE`, and the traffic is detected as malicious by a protection rule. If unspecified, defaults to `403`.
		* `block_error_page_description` - The description text to show on the error page when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_ERROR_PAGE`, and the traffic is detected as malicious by a protection rule. If unspecified, defaults to `Access blocked by website owner. Please contact support.`
		* `block_error_page_message` - The message to show on the error page when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_ERROR_PAGE`, and the traffic is detected as malicious by a protection rule. If unspecified, defaults to 'Access to the website is blocked.'
		* `block_response_code` - The response code returned when `action` is set to `BLOCK`, `blockAction` is set to `SET_RESPONSE_CODE`, and the traffic is detected as malicious by a protection rule. If unspecified, defaults to `403`.
		* `is_response_inspected` - Inspects the response body of origin responses. Can be used to detect leakage of sensitive data. If unspecified, defaults to `false`.

			**Note:** Only origin responses with a Content-Type matching a value in `mediaTypes` will be inspected.
		* `max_argument_count` - The maximum number of arguments allowed to be passed to your application before an action is taken. Arguements are query parameters or body parameters in a PUT or POST request. If unspecified, defaults to `255`. This setting only applies if a corresponding protection rule is enabled, such as the "Number of Arguments Limits" rule (key: 960335).  Example: If `maxArgumentCount` to `2` for the Max Number of Arguments protection rule (key: 960335), the following requests would be blocked: `GET /myapp/path?query=one&query=two&query=three` `POST /myapp/path` with Body `{"argument1":"one","argument2":"two","argument3":"three"}`
		* `max_name_length_per_argument` - The maximum length allowed for each argument name, in characters. Arguements are query parameters or body parameters in a PUT or POST request. If unspecified, defaults to `400`. This setting only applies if a corresponding protection rule is enabled, such as the "Values Limits" rule (key: 960208).
		* `max_response_size_in_ki_b` - The maximum response size to be fully inspected, in binary kilobytes (KiB). Anything over this limit will be partially inspected. If unspecified, defaults to `1024`.
		* `max_total_name_length_of_arguments` - The maximum length allowed for the sum of the argument name and value, in characters. Arguements are query parameters or body parameters in a PUT or POST request. If unspecified, defaults to `64000`. This setting only applies if a corresponding protection rule is enabled, such as the "Total Arguments Limits" rule (key: 960341).
		* `media_types` - The list of media types to allow for inspection, if `isResponseInspected` is enabled. Only responses with MIME types in this list will be inspected. If unspecified, defaults to `["text/html", "text/plain", "text/xml"]`.

			Supported MIME types include:
			* text/html
			* text/plain
			* text/asp
			* text/css
			* text/x-script
			* application/json
			* text/webviewhtml
			* text/x-java-source
			* application/x-javascript
			* application/javascript
			* application/ecmascript
			* text/javascript
			* text/ecmascript
			* text/x-script.perl
			* text/x-script.phyton
			* application/plain
			* application/xml
			* text/xml
		* `recommendations_period_in_days` - The length of time to analyze traffic traffic, in days. After the analysis period, `WafRecommendations` will be populated. If unspecified, defaults to `10`.

			Use `GET /waasPolicies/{waasPolicyId}/wafRecommendations` to view WAF recommendations.
	* `whitelists` - A list of IP addresses that bypass the Web Application Firewall.
		* `addresses` - A set of IP addresses or CIDR notations to include in the whitelist.
		* `name` - The unique name of the whitelist.

## Import

WaasPolicies can be imported using the `id`, e.g.

```
$ terraform import oci_waas_waas_policy.test_waas_policy "id"
```

