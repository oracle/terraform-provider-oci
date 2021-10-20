---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_log_analytics_import_custom_content"
sidebar_current: "docs-oci-resource-log_analytics-log_analytics_import_custom_content"
description: |-
  Provides the Log Analytics Import Custom Content resource in Oracle Cloud Infrastructure Log Analytics service
---

# oci_log_analytics_log_analytics_import_custom_content
This resource provides the Log Analytics Import Custom Content resource in Oracle Cloud Infrastructure Log Analytics service.

Imports the specified custom content from the input in zip format.


## Example Usage

```hcl
resource "oci_log_analytics_log_analytics_import_custom_content" "test_log_analytics_import_custom_content" {
	#Required
	import_custom_content_file = var.log_analytics_import_custom_content_import_custom_content_file
	namespace = var.log_analytics_import_custom_content_namespace

	#Optional
	expect = var.log_analytics_import_custom_content_expect
	is_overwrite = var.log_analytics_import_custom_content_is_overwrite
}
```

## Argument Reference

The following arguments are supported:

* `import_custom_content_file` - (Required) Path to the file to upload which contains the custom content.
* `is_overwrite` - (Optional) A flag indicating whether or not to overwrite existing content if a conflict is found during import content operation. 
* `namespace` - (Required) The Logging Analytics namespace used for the request.
* `expect` - (Optional) A value of `100-continue` requests preliminary verification of the request method, path, and headers before the request body is sent. If no error results from such verification, the server will send a 100 (Continue) interim response to indicate readiness for the request body. The only allowed value for this parameter is "100-Continue" (case-insensitive).


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `change_list` - LogAnalyticsImportCustomChangeList
	* `conflict_field_display_names` - A list of field display names with conflicts.
	* `conflict_parser_names` - A list of parser names with conflicts.
	* `conflict_source_names` - A list of source names with conflicts.
	* `created_field_display_names` - An array of created field display names.
	* `created_parser_names` - An array of created parser names.
	* `created_source_names` - An array of created source names.
	* `updated_field_display_names` - An array of updated field display names.
	* `updated_parser_names` - An array of updated parser names.
	* `updated_source_names` - An array of updated source names.
* `content_name` - The content name.
* `field_names` - The field names.
* `parser_names` - The parser names.
* `source_names` - The source names.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Log Analytics Import Custom Content
	* `update` - (Defaults to 20 minutes), when updating the Log Analytics Import Custom Content
	* `delete` - (Defaults to 20 minutes), when destroying the Log Analytics Import Custom Content


## Import

Import is not supported for LogAnalyticsImportCustomContent
