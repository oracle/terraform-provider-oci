---
subcategory: "OS Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_osmanagement_software_source_module_stream_profile"
sidebar_current: "docs-oci-datasource-osmanagement-software_source_module_stream_profile"
description: |-
  Provides details about a specific Software Source Module Stream Profile in Oracle Cloud Infrastructure OS Management service
---

# Data Source: oci_osmanagement_software_source_module_stream_profile
This data source provides details about a specific Software Source Module Stream Profile resource in Oracle Cloud Infrastructure OS Management service.

Retrieve a detailed description of a module stream profile from a software source.


## Example Usage

```hcl
data "oci_osmanagement_software_source_module_stream_profile" "test_software_source_module_stream_profile" {
	#Required
	module_name = var.software_source_module_name
	profile_name = var.software_source_module_stream_profile_name
	software_source_id = var.software_source.id
	stream_name = var.software_source_module_stream_name
}
```

## Argument Reference

The following arguments are supported:

* `module_name` - (Required) The name of the module
* `profile_name` - (Required) The name of the profile of the containing module stream
* `software_source_id` - (Required) The OCID of the software source.
* `stream_name` - (Required) The name of the stream of the containing module


## Attributes Reference

The following attributes are exported:

* `description` - A description of the contents of the module stream profile
* `is_default` - Indicates if this profile is the default for its module stream.
* `module_name` - The name of the module that contains the stream profile
* `packages` - A list of packages that constitute the profile.  Each element in the list is the name of a package.  The name is suitable to use as an argument to other OS Management APIs that interact directly with packages. 
* `profile_name` - The name of the profile
* `stream_name` - The name of the stream that contains the profile

