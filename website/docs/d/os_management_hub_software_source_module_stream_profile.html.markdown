---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_software_source_module_stream_profile"
sidebar_current: "docs-oci-datasource-os_management_hub-software_source_module_stream_profile"
description: |-
  Provides details about a specific Software Source Module Stream Profile in Oracle Cloud Infrastructure Os Management Hub service
---

# Data Source: oci_os_management_hub_software_source_module_stream_profile
This data source provides details about a specific Software Source Module Stream Profile resource in Oracle Cloud Infrastructure Os Management Hub service.

Gets information about the specified module stream profile in a software source.


## Example Usage

```hcl
data "oci_os_management_hub_software_source_module_stream_profile" "test_software_source_module_stream_profile" {
	#Required
	module_name = var.software_source_module_stream_profile_module_name
	profile_name = oci_os_management_hub_profile.test_profile.name
	software_source_id = oci_os_management_hub_software_source.test_software_source.id
	stream_name = oci_streaming_stream.test_stream.name
}
```

## Argument Reference

The following arguments are supported:

* `module_name` - (Required) The name of a module. 
* `profile_name` - (Required) The name of the profile of the containing module stream.
* `software_source_id` - (Required) The software source OCID.
* `stream_name` - (Required) The name of the stream of the containing module. 


## Attributes Reference

The following attributes are exported:

* `description` - A description of the contents of the module stream profile.
* `is_default` - Indicates if this profile is the default for its module stream.
* `module_name` - The name of the module that contains the stream profile.
* `name` - The name of the profile.
* `packages` - A list of packages that constitute the profile.  Each element in the list is the name of a package.  The name is suitable to use as an argument to other OS Management Hub APIs that interact directly with packages. 
* `stream_name` - The name of the stream that contains the profile.

