---
subcategory: "Functions"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_functions_functions_runtime_version"
sidebar_current: "docs-oci-datasource-functions-functions_runtime_version"
description: |-
  Provides details about a specific Functions Runtime Version in Oracle Cloud Infrastructure Functions service
---

# Data Source: oci_functions_functions_runtime_version
This data source provides details about a specific Functions Runtime Version resource in Oracle Cloud Infrastructure Functions service.

Fetches a FunctionsRuntimeVersion by ID. Returns a FunctionsRuntimeVersion response model.


## Example Usage

```hcl
data "oci_functions_functions_runtime_version" "test_functions_runtime_version" {
	#Required
	functions_runtime_version_id = var.functions_runtime_version_id
}
```

## Argument Reference

The following arguments are supported:

* `functions_runtime_version_id` - (Required) unique FunctionsRuntimeVersion identifier


## Attributes Reference

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
