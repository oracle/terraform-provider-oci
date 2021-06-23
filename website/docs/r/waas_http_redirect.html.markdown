---
subcategory: "Web Application Acceleration and Security"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_waas_http_redirect"
sidebar_current: "docs-oci-resource-waas-http_redirect"
description: |-
  Provides the Http Redirect resource in Oracle Cloud Infrastructure Web Application Acceleration and Security service
---

# oci_waas_http_redirect
This resource provides the Http Redirect resource in Oracle Cloud Infrastructure Web Application Acceleration and Security service.

Creates a new HTTP Redirect on the WAF edge.

## Example Usage

```hcl
resource "oci_waas_http_redirect" "test_http_redirect" {
	#Required
	compartment_id = var.compartment_id
	domain = var.http_redirect_domain
	target {
		#Required
		host = var.http_redirect_target_host
		path = var.http_redirect_target_path
		protocol = var.http_redirect_target_protocol
		query = var.http_redirect_target_query

		#Optional
		port = var.http_redirect_target_port
	}

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.http_redirect_display_name
	freeform_tags = {"Department"= "Finance"}
	response_code = var.http_redirect_response_code
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the HTTP Redirects compartment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) The user-friendly name of the HTTP Redirect. The name can be changed and does not need to be unique.
* `domain` - (Required) The domain from which traffic will be redirected.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `response_code` - (Optional) (Updatable) The response code returned for the redirect to the client. For more information, see [RFC 7231](https://tools.ietf.org/html/rfc7231#section-6.4).
* `target` - (Required) (Updatable) The redirect target object including all the redirect data.
	* `host` - (Required) (Updatable) The host portion of the redirect.
	* `path` - (Required) (Updatable) The path component of the target URL (e.g., "/path/to/resource" in "https://target.example.com/path/to/resource?redirected"), which can be empty, static, or request-copying, or request-prefixing. Use of \ is not permitted except to escape a following \, {, or }. An empty value is treated the same as static "/". A static value must begin with a leading "/", optionally followed by other path characters. A request-copying value must exactly match "{path}", and will be replaced with the path component of the request URL (including its initial "/"). A request-prefixing value must start with "/" and end with a non-escaped "{path}", which will be replaced with the path component of the request URL (including its initial "/"). Only one such replacement token is allowed.
	* `port` - (Optional) (Updatable) Port number of the target destination of the redirect, default to match protocol
	* `protocol` - (Required) (Updatable) The protocol used for the target, http or https.
	* `query` - (Required) (Updatable) The query component of the target URL (e.g., "?redirected" in "https://target.example.com/path/to/resource?redirected"), which can be empty, static, or request-copying. Use of \ is not permitted except to escape a following \, {, or }. An empty value results in a redirection target URL with no query component. A static value must begin with a leading "?", optionally followed by other query characters. A request-copying value must exactly match "{query}", and will be replaced with the query component of the request URL (including a leading "?" if and only if the request URL includes a query component).


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Http Redirect
	* `update` - (Defaults to 20 minutes), when updating the Http Redirect
	* `delete` - (Defaults to 20 minutes), when destroying the Http Redirect


## Import

HttpRedirects can be imported using the `id`, e.g.

```
$ terraform import oci_waas_http_redirect.test_http_redirect "id"
```

