---
subcategory: "Apm Synthetics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apm_synthetics_scripts"
sidebar_current: "docs-oci-datasource-apm_synthetics-scripts"
description: |-
  Provides the list of Scripts in Oracle Cloud Infrastructure Apm Synthetics service
---

# Data Source: oci_apm_synthetics_scripts
This data source provides the list of Scripts in Oracle Cloud Infrastructure Apm Synthetics service.

Returns a list of scripts.


## Example Usage

```hcl
data "oci_apm_synthetics_scripts" "test_scripts" {
	#Required
	apm_domain_id = oci_apm_synthetics_apm_domain.test_apm_domain.id

	#Optional
	content_type = var.script_content_type
	display_name = var.script_display_name
}
```

## Argument Reference

The following arguments are supported:

* `apm_domain_id` - (Required) The APM domain ID the request is intended for. 
* `content_type` - (Optional) A filter to return only resources that match the content type given.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.


## Attributes Reference

The following attributes are exported:

* `script_collection` - The list of script_collection.

### Script Reference

The following attributes are exported:

* `content` - The content of the script. It may contain custom-defined tags that can be used for setting dynamic parameters. The format to set dynamic parameters is: `<ORAP><ON>param name</ON><OV>param value</OV><OS>isParamValueSecret(true/false)</OS></ORAP>`. Param value and isParamValueSecret are optional, the default value for isParamValueSecret is false. Examples: With mandatory param name : `<ORAP><ON>param name</ON></ORAP>` With parameter name and value : `<ORAP><ON>param name</ON><OV>param value</OV></ORAP>` Note that the content is valid if it matches the given content type. For example, if the content type is SIDE, then the content should be in Side script format. If the content type is JS, then the content should be in JavaScript format. 
* `content_file_name` - File name of the uploaded script content.
* `content_size_in_bytes` - Size of the script content.
* `content_type` - Content type of the script.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Unique name that can be edited. The name should not contain any confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the script. scriptId is mandatory for creation of SCRIPTED_BROWSER and SCRIPTED_REST monitor types. For other monitor types, it should be set to null. 
* `monitor_status_count_map` - Details of the monitor count per state. Example: `{ "total" : 5, "enabled" : 3 , "disabled" : 2, "invalid" : 0 }` 
	* `disabled` - Number of disabled monitors using the script.
	* `enabled` - Number of enabled monitors using the script.
	* `invalid` - Number of invalid monitors using the script.
	* `total` - Total number of monitors using the script.
* `parameters` - List of script parameters. Example: `[{"scriptParameter": {"paramName": "userid", "paramValue":"testuser", "isSecret": false}, "isOverwritten": false}]` 
	* `is_overwritten` - If parameter value is default or overwritten. 
	* `script_parameter` - Details of the script parameters, paramName must be from the script content and these details can be used to overwrite the default parameter present in the script content. 
		* `is_secret` - If the parameter value is secret and should be kept confidential, then set isSecret to true.
		* `param_name` - Name of the parameter.
		* `param_value` - Value of the parameter.
* `time_created` - The time the resource was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2020-02-12T22:47:12.613Z` 
* `time_updated` - The time the resource was updated, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2020-02-13T22:47:12.613Z` 
* `time_uploaded` - The time when the script was uploaded.

