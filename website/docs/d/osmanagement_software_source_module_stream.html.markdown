---
subcategory: "OS Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_osmanagement_software_source_module_stream"
sidebar_current: "docs-oci-datasource-osmanagement-software_source_module_stream"
description: |-
  Provides details about a specific Software Source Module Stream in Oracle Cloud Infrastructure OS Management service
---

# Data Source: oci_osmanagement_software_source_module_stream
This data source provides details about a specific Software Source Module Stream resource in Oracle Cloud Infrastructure OS Management service.

Retrieve a detailed description of a module stream from a software source.


## Example Usage

```hcl
data "oci_osmanagement_software_source_module_stream" "test_software_source_module_stream" {
	#Required
	module_name = var.software_source_module_stream_module_name
	software_source_id = var.software_source.id
	stream_name = var.software_source_module_stream_name
}
```

## Argument Reference

The following arguments are supported:

* `module_name` - (Required) The name of the module
* `software_source_id` - (Required) The OCID of the software source.
* `stream_name` - (Required) The name of the stream of the containing module


## Attributes Reference

The following attributes are exported:

* `architecture` - The architecture for which the packages in this module stream were built
* `description` - A description of the contents of the module stream
* `is_default` - Indicates if this stream is the default for its module.
* `module_name` - The name of the module that contains the stream
* `packages` - A list of packages that are contained by the stream.  Each element in the list is the name of a package.  The name is suitable to use as an argument to other OS Management APIs that interact directly with packages. 
* `profiles` - A list of profiles that are part of the stream.  Each element in the list is the name of a profile.  The name is suitable to use as an argument to other OS Management APIs that interact directly with module stream profiles.  However, it is not URL encoded. 
* `software_source_id` - The OCID of the software source that provides this module stream.
* `stream_name` - The name of the stream

