---
subcategory: "Functions"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_functions_functions_runtime_versions"
sidebar_current: "docs-oci-datasource-functions-functions_runtime_versions"
description: |-
  Provides the list of Functions Runtime Versions in Oracle Cloud Infrastructure Functions service
---

# Data Source: oci_functions_functions_runtime_versions
This data source provides the list of Functions Runtime Versions in Oracle Cloud Infrastructure Functions service.

Fetches a wrapped list of all FunctionsRuntimeVersions. Returns a FunctionsRuntimeVersionCollection containing
an array of FunctionsRuntimeVersionSummary response models.


## Example Usage

```hcl
data "oci_functions_functions_runtime_versions" "test_functions_runtime_versions" {

	#Optional
	display_name = var.functions_runtime_version_display_name
	functions_runtime_id = var.functions_runtime_id
	functions_runtime_name = var.functions_runtime_name
	functions_runtime_version_id = var.functions_runtime_version_id
	is_current_version = var.functions_runtime_version_is_current_version
	language_version = var.functions_runtime_version_language_version
	os_version = var.functions_runtime_version_os_version
	state = var.functions_runtime_version_state
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) A filter to return only resources that match the entire FunctionsRuntimeVersion name given.
* `functions_runtime_id` - (Optional) unique FunctionsRuntime identifier
* `functions_runtime_name` - (Optional) A filter to return only resources that match the entire FunctionsRuntime name given.
* `functions_runtime_version_id` - (Optional) unique FunctionsRuntimeVersion identifier
* `is_current_version` - (Optional) Matches the current version associated with a FunctionsRuntime.
* `language_version` - (Optional) A filter to return only resources that match the entire languageVersion name given.
* `os_version` - (Optional) A filter to return only resources that match the entire osVersion name given.
* `state` - (Optional) A filter to return only resources their lifecycleState matches the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `functions_runtime_version_collection` - The list of functions_runtime_version_collection.

### FunctionsRuntimeVersion Reference

The following attributes are exported:

* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
* `display_name` - The display name of the FunctionsRuntimeVersion.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
* `functions_runtime_id` - The OCID of the FunctionsRuntime this resource version belongs to.
* `id` - The OCID of the FunctionsRuntimeVersion that is immutable on creation.
* `language_version` - The version of the programming language of the FunctionsRuntime. This is the language version that the FunctionsRuntime provides for execution of customer payloads.
* `metadata` - Details of the change in the FunctionsRuntimeVersion of the FunctionsRuntime.
* `os_version` - The version of the operating system of the FunctionsRuntime. This is the OS version that the FunctionsRuntime provides for execution of customer payloads.
* `state` - The current state of the FunctionsRuntimeVersion resource.
	* `ACTIVE`: The resource is currently active and operational.
	* `INACTIVE`: The resource is currently inactive and not operational.
* `supported_architectures` - The list of supported architectures for the FunctionsRuntimeVersion.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `time_created` - The time when the FunctionsRuntimeVersion was created. An RFC3339 formatted datetime string.
* `time_updated` - The time when the FunctionsRuntimeVersion was updated. An RFC3339 formatted datetime string.
