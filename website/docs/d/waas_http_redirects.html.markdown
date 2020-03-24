---
subcategory: "Web Application Acceleration and Security"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_waas_http_redirects"
sidebar_current: "docs-oci-datasource-waas-http_redirects"
description: |-
  Provides the list of Http Redirects in Oracle Cloud Infrastructure Web Application Acceleration and Security service
---

# Data Source: oci_waas_http_redirects
This data source provides the list of Http Redirects in Oracle Cloud Infrastructure Web Application Acceleration and Security service.

Gets a list of HTTP Redirects.

## Example Usage

```hcl
data "oci_waas_http_redirects" "test_http_redirects" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	display_names = "${var.http_redirect_display_names}"
	ids = "${var.http_redirect_ids}"
	states = "${var.http_redirect_states}"
	time_created_greater_than_or_equal_to = "${var.http_redirect_time_created_greater_than_or_equal_to}"
	time_created_less_than = "${var.http_redirect_time_created_less_than}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. This number is generated when the compartment is created.
* `display_names` - (Optional) Filter redirects using a display name.
* `ids` - (Optional) Filter redirects using a list of redirect OCIDs.
* `states` - (Optional) Filter redirects using a list of lifecycle states.
* `time_created_greater_than_or_equal_to` - (Optional) A filter that matches redirects created on or after the specified date and time.
* `time_created_less_than` - (Optional) A filter that matches redirects created before the specified date-time. Default to 1 day before now.


## Attributes Reference

The following attributes are exported:

* `http_redirects` - The list of http_redirects.

### HttpRedirect Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the HTTP Redirect's compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The user-friendly name of the HTTP Redirect. The name can be changed and does not need to be unique.
* `domain` - The domain from which traffic will be redirected.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the HTTP Redirect.
* `response_code` - The response code returned for the redirect to the client. For more information, see [RFC 7231](https://tools.ietf.org/html/rfc7231#section-6.4).
* `state` - The current lifecycle state of the HTTP Redirect.
* `target` - The redirect target object including all the redirect data.
	* `host` - The host portion of the redirect.
	* `path` - The path component of the target URL (e.g., "/path/to/resource" in "https://target.example.com/path/to/resource?redirected"), which can be empty, static, or request-copying, or request-prefixing. Use of \ is not permitted except to escape a following \, {, or }. An empty value is treated the same as static "/". A static value must begin with a leading "/", optionally followed by other path characters. A request-copying value must exactly match "{path}", and will be replaced with the path component of the request URL (including its initial "/"). A request-prefixing value must start with "/" and end with a non-escaped "{path}", which will be replaced with the path component of the request URL (including its initial "/"). Only one such replacement token is allowed.
	* `port` - Port number of the target destination of the redirect, default to match protocol
	* `protocol` - The protocol used for the target, http or https.
	* `query` - The query component of the target URL (e.g., "?redirected" in "https://target.example.com/path/to/resource?redirected"), which can be empty, static, or request-copying. Use of \ is not permitted except to escape a following \, {, or }. An empty value results in a redirection target URL with no query component. A static value must begin with a leading "?", optionally followed by other query characters. A request-copying value must exactly match "{query}", and will be replaced with the query component of the request URL (including a leading "?" if and only if the request URL includes a query component).
* `time_created` - The date and time the policy was created, expressed in RFC 3339 timestamp format.

