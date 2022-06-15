---
subcategory: "Waa"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_waa_web_app_acceleration_policies"
sidebar_current: "docs-oci-datasource-waa-web_app_acceleration_policies"
description: |-
  Provides the list of Web App Acceleration Policies in Oracle Cloud Infrastructure Waa service
---

# Data Source: oci_waa_web_app_acceleration_policies
This data source provides the list of Web App Acceleration Policies in Oracle Cloud Infrastructure Waa service.

Gets a list of all WebAppAccelerationPolicies in a compartment.


## Example Usage

```hcl
data "oci_waa_web_app_acceleration_policies" "test_web_app_acceleration_policies" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.web_app_acceleration_policy_display_name
	id = var.web_app_acceleration_policy_id
	state = var.web_app_acceleration_policy_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `id` - (Optional) A filter to return only the WebAppAccelerationPolicy with the given [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `state` - (Optional) A filter to return only resources that match the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `web_app_acceleration_policy_collection` - The list of web_app_acceleration_policy_collection.

### WebAppAccelerationPolicy Reference

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

