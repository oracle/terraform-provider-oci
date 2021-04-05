---
subcategory: "Functions"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_functions_function"
sidebar_current: "docs-oci-datasource-functions-function"
description: |-
  Provides details about a specific Function in Oracle Cloud Infrastructure Functions service
---

# Data Source: oci_functions_function
This data source provides details about a specific Function resource in Oracle Cloud Infrastructure Functions service.

Retrieves a function.

## Example Usage

```hcl
data "oci_functions_function" "test_function" {
	#Required
	function_id = oci_functions_function.test_function.id
}
```

## Argument Reference

The following arguments are supported:

* `function_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of this function. 


## Attributes Reference

The following attributes are exported:

* `application_id` - The OCID of the application the function belongs to.
* `compartment_id` - The OCID of the compartment that contains the function. 
* `config` - Function configuration. Overrides application configuration. Keys must be ASCII strings consisting solely of letters, digits, and the '_' (underscore) character, and must not begin with a digit. Values should be limited to printable unicode characters.  Example: `{"MY_FUNCTION_CONFIG": "ConfVal"}`

	The maximum size for all configuration keys and values is limited to 4KB. This is measured as the sum of octets necessary to represent each key and value in UTF-8. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The display name of the function. The display name is unique within the application containing the function. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the function. 
* `image` - The qualified name of the Docker image to use in the function, including the image tag. The image should be in the Oracle Cloud Infrastructure Registry that is in the same region as the function itself. Example: `phx.ocir.io/ten/functions/function:0.0.1` 
* `image_digest` - The image digest for the version of the image that will be pulled when invoking this function. If no value is specified, the digest currently associated with the image in the Oracle Cloud Infrastructure Registry will be used. Example: `sha256:ca0eeb6fb05351dfc8759c20733c91def84cb8007aa89a5bf606bc8b315b9fc7` 
* `invoke_endpoint` - The base https invoke URL to set on a client in order to invoke a function. This URL will never change over the lifetime of the function and can be cached. 
* `memory_in_mbs` - Maximum usable memory for the function (MiB).
* `state` - The current state of the function. 
* `time_created` - The time the function was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.  Example: `2018-09-12T22:47:12.613Z` 
* `time_updated` - The time the function was updated, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.  Example: `2018-09-12T22:47:12.613Z` 
* `timeout_in_seconds` - Timeout for executions of the function. Value in seconds.
* `trace_config` - Define the tracing configuration for a function. 
	* `is_enabled` - Define if tracing is enabled for the resource. 

