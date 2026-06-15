---
subcategory: "Functions"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_functions_functions_runtimes"
sidebar_current: "docs-oci-datasource-functions-functions_runtimes"
description: |-
  Provides the list of Functions Runtimes in Oracle Cloud Infrastructure Functions service
---

# Data Source: oci_functions_functions_runtimes
This data source provides the list of Functions Runtimes in Oracle Cloud Infrastructure Functions service.

Fetches a wrapped list of all FunctionsRuntimes. Returns a FunctionsRuntimeCollection containing
an array of FunctionsRuntimeSummary response models.


## Example Usage

```hcl
data "oci_functions_functions_runtimes" "test_functions_runtimes" {

	#Optional
	functions_runtime_id = var.functions_runtime_id
	language = var.functions_runtime_language
	name = var.functions_runtime_name
	name_contains = var.functions_runtime_name_contains
	name_starts_with = var.functions_runtime_name_starts_with
	os = var.functions_runtime_os
	state = var.functions_runtime_state
}
```

## Argument Reference

The following arguments are supported:

* `functions_runtime_id` - (Optional) unique FunctionsRuntime identifier
* `language` - (Optional) A filter to return only resources that match the entire language name given.
* `name` - (Optional) A filter to return only resources that match the entire FunctionsRuntime name given.
* `name_contains` - (Optional) A filter to return only resources that contain the supplied filter text in the FunctionsRuntime name given.
* `name_starts_with` - (Optional) A filter to return only resources that start with the supplied filter text in the FunctionsRuntime name given.
* `os` - (Optional) A filter to return only resources that match the entire os name given.
* `state` - (Optional) A filter to return only resources where their lifecycleState matches the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `functions_runtime_collection` - The list of functions_runtime_collection.

### FunctionsRuntime Reference

The following attributes are exported:

* `current_functions_runtime_version_id` - The OCID of the current FunctionsRuntimeVersion for this FunctionsRuntime.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
* `id` - The OCID of the FunctionsRuntime that is immutable on creation.
* `language` - The programming language of the FunctionsRuntime. This is the language that the FunctionsRuntime provides for execution of customer payloads.
* `metadata` - Metadata for the FunctionsRuntime Resource.
* `name` - A brief descriptive name for the FunctionsRuntime. The FunctionsRuntime name must be unique, and not match any existing FunctionsRuntime.
* `os` - The operating system of the FunctionsRuntime. This is the OS that the FunctionsRuntime provides for execution of customer payloads.
* `state` - The current state of the FunctionsRuntime resource.
	* `ACTIVE`: The resource is currently active and operational.
	* `INACTIVE`: The resource is currently inactive and not operational.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `time_created` - The time when the FunctionsRuntime was created. An RFC3339 formatted datetime string.
* `time_decommissioned` - The time when the FunctionsRuntime will be decommissioned. An RFC3339 formatted datetime string.
* `time_deprecated` - The time when the FunctionsRuntime will be deprecated. An RFC3339 formatted datetime string.
* `time_updated` - The time when the FunctionsRuntime was updated. An RFC3339 formatted datetime string.
