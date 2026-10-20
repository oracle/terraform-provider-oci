---
subcategory: "Functions"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_functions_functions_runtime"
sidebar_current: "docs-oci-datasource-functions-functions_runtime"
description: |-
  Provides details about a specific Functions Runtime in Oracle Cloud Infrastructure Functions service
---

# Data Source: oci_functions_functions_runtime
This data source provides details about a specific Functions Runtime resource in Oracle Cloud Infrastructure Functions service.

Fetches a FunctionsRuntime by ID. Returns a FunctionsRuntime response model.


## Example Usage

```hcl
data "oci_functions_functions_runtime" "test_functions_runtime" {
	#Required
	functions_runtime_id = var.functions_runtime_id
}
```

## Argument Reference

The following arguments are supported:

* `functions_runtime_id` - (Required) unique FunctionsRuntime identifier


## Attributes Reference

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
