---
subcategory: "Waa"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_waa_web_app_acceleration_policy"
sidebar_current: "docs-oci-resource-waa-web_app_acceleration_policy"
description: |-
  Provides the Web App Acceleration Policy resource in Oracle Cloud Infrastructure Waa service
---

# oci_waa_web_app_acceleration_policy
This resource provides the Web App Acceleration Policy resource in Oracle Cloud Infrastructure Waa service.

Creates a new WebAppAccelerationPolicy.


## Example Usage

```hcl
resource "oci_waa_web_app_acceleration_policy" "test_web_app_acceleration_policy" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	display_name = var.web_app_acceleration_policy_display_name
	freeform_tags = {"bar-key"= "value"}
	response_caching_policy {

		#Optional
		is_response_header_based_caching_enabled = var.web_app_acceleration_policy_response_caching_policy_is_response_header_based_caching_enabled
	}
	response_compression_policy {

		#Optional
		gzip_compression {

			#Optional
			is_enabled = var.web_app_acceleration_policy_response_compression_policy_gzip_compression_is_enabled
		}
	}
	system_tags = var.web_app_acceleration_policy_system_tags
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Optional) (Updatable) WebAppAccelerationPolicy display name, can be renamed.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `response_caching_policy` - (Optional) (Updatable) An object that specifies an HTTP response caching policy. 
	* `is_response_header_based_caching_enabled` - (Optional) (Updatable) When false, responses will not be cached by the backend based on response headers.

		When true, responses that contain one of the supported cache control headers will be cached according to the values specified in the cache control headers.

		The "X-Accel-Expires" header field sets caching time of a response in seconds. The zero value disables caching for a response. If the value starts with the @ prefix, it sets an absolute time in seconds since Epoch, up to which the response may be cached.

		If the header does not include the "X-Accel-Expires" field, parameters of caching may be set in the header fields "Expires" or "Cache-Control".

		If the header includes the "Set-Cookie" field, such a response will not be cached.

		If the header includes the "Vary" field with the special value "*", such a response will not be cached. If the header includes the "Vary" field with another value, such a response will be cached taking into account the corresponding request header fields. 
* `response_compression_policy` - (Optional) (Updatable) An object that specifies a compression policy for HTTP response from ENABLEMENT POINT to the client.

	This compression policy can be used to enable support for HTTP response compression algorithms like gzip and configure the conditions of when a compression algorithm will be used.

	HTTP responses will only be compressed if the client indicates support for one of the enabled compression algorithms via the "Accept-Encoding" request header. 
	* `gzip_compression` - (Optional) (Updatable) An object that specifies the gzip compression policy. 
		* `is_enabled` - (Optional) (Updatable) When true, support for gzip compression is enabled. HTTP responses will be compressed with gzip only if the client indicates support for gzip via the "Accept-Encoding: gzip" request header.

			When false, support for gzip compression is disabled and HTTP responses will not be compressed with gzip even if the client indicates support for gzip. 
* `system_tags` - (Optional) (Updatable) Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - WebAppAccelerationPolicy display name, can be renamed.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the WebAppAccelerationPolicy.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in FAILED state. 
* `response_caching_policy` - An object that specifies an HTTP response caching policy. 
	* `is_response_header_based_caching_enabled` - When false, responses will not be cached by the backend based on response headers.

		When true, responses that contain one of the supported cache control headers will be cached according to the values specified in the cache control headers.

		The "X-Accel-Expires" header field sets caching time of a response in seconds. The zero value disables caching for a response. If the value starts with the @ prefix, it sets an absolute time in seconds since Epoch, up to which the response may be cached.

		If the header does not include the "X-Accel-Expires" field, parameters of caching may be set in the header fields "Expires" or "Cache-Control".

		If the header includes the "Set-Cookie" field, such a response will not be cached.

		If the header includes the "Vary" field with the special value "*", such a response will not be cached. If the header includes the "Vary" field with another value, such a response will be cached taking into account the corresponding request header fields. 
* `response_compression_policy` - An object that specifies a compression policy for HTTP response from ENABLEMENT POINT to the client.

	This compression policy can be used to enable support for HTTP response compression algorithms like gzip and configure the conditions of when a compression algorithm will be used.

	HTTP responses will only be compressed if the client indicates support for one of the enabled compression algorithms via the "Accept-Encoding" request header. 
	* `gzip_compression` - An object that specifies the gzip compression policy. 
		* `is_enabled` - When true, support for gzip compression is enabled. HTTP responses will be compressed with gzip only if the client indicates support for gzip via the "Accept-Encoding: gzip" request header.

			When false, support for gzip compression is disabled and HTTP responses will not be compressed with gzip even if the client indicates support for gzip. 
* `state` - The current state of the WebAppAccelerationPolicy.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the WebAppAccelerationPolicy was created. An RFC3339 formatted datetime string.
* `time_updated` - The time the WebAppAccelerationPolicy was updated. An RFC3339 formatted datetime string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Web App Acceleration Policy
	* `update` - (Defaults to 20 minutes), when updating the Web App Acceleration Policy
	* `delete` - (Defaults to 20 minutes), when destroying the Web App Acceleration Policy


## Import

WebAppAccelerationPolicies can be imported using the `id`, e.g.

```
$ terraform import oci_waa_web_app_acceleration_policy.test_web_app_acceleration_policy "id"
```

