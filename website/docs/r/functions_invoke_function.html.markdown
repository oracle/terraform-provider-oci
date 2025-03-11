---
subcategory: "Functions"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_functions_invoke_function"
sidebar_current: "docs-oci-resource-functions-invoke_function"
description: |-
  Provides the Invoke Function resource in Oracle Cloud Infrastructure Functions service
---

# oci_functions_invoke_function
This resource provides the Invoke Function resource in Oracle Cloud Infrastructure Functions service.

Invokes a function

## Example Usage

```hcl
resource "oci_functions_invoke_function" "test_invoke_function" {
	#Required
	function_id = oci_functions_function.test_function.id

	#Optional
	invoke_function_body = var.invoke_function_invoke_function_body
	fn_intent = var.invoke_function_fn_intent
	fn_invoke_type = var.invoke_function_fn_invoke_type
	is_dry_run = var.invoke_function_is_dry_run
	base64_encode_content = false
}
```

## Argument Reference

The following arguments are supported:

* `base64_encode_content` - (Optional) Encodes the response returned, if any, in base64. It is recommended to set this to `true` to avoid corrupting the returned response, if any, in Terraform state. The default value is `false`.
* `invoke_function_body` - (Optional) The body of the function invocation. Note: The maximum size of the request is limited. This limit is currently 6MB and the endpoint will not accept requests that are bigger than this limit. Cannot be defined if `input_body_source_path` or `invoke_function_body_base64_encoded` is defined.
* `fn_intent` - (Optional) An optional intent header that indicates to the FDK the way the event should be interpreted. E.g. 'httprequest', 'cloudevent'. 
* `fn_invoke_type` - (Optional) Indicates whether Oracle Functions should execute the request and return the result ('sync') of the execution,  or whether Oracle Functions should return as soon as processing has begun ('detached') and leave result handling to the function. 
* `function_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of this function. 
* `is_dry_run` - (Optional) Indicates that the request is a dry run, if set to "true". A dry run request does not execute the function. 

* `invoke_function_body_base64_encoded` - (Optional) The Base64 encoded body of the function invocation. Base64 encoded input avoids corruption in Terraform state. Cannot be defined if `invoke_function_body` or `input_body_source_path` is defined. Note: The maximum size of the request is limited. This limit is currently 6MB and the endpoint will not accept requests that are bigger than this limit. 
* `input_body_source_path` - (Optional) An absolute path to a file on the local system that contains the input to be provided to the function. Cannot be defined if `invoke_function_body` or `invoke_function_body_base64_encoded` is defined. Note: The maximum size of the request is limited. This limit is currently 6MB and the endpoint will not accept requests that are bigger than this limit.

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `content` - Content of the response string, if any. If `base64_encode_content` is set to `true`, then this content will be base64 encoded.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Invoke Function
	* `update` - (Defaults to 20 minutes), when updating the Invoke Function
	* `delete` - (Defaults to 20 minutes), when destroying the Invoke Function


## Import

Import is not supported for this resource.
