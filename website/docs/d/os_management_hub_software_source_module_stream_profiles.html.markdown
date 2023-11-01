---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_software_source_module_stream_profiles"
sidebar_current: "docs-oci-datasource-os_management_hub-software_source_module_stream_profiles"
description: |-
  Provides the list of Software Source Module Stream Profiles in Oracle Cloud Infrastructure Os Management Hub service
---

# Data Source: oci_os_management_hub_software_source_module_stream_profiles
This data source provides the list of Software Source Module Stream Profiles in Oracle Cloud Infrastructure Os Management Hub service.

Lists module stream profiles from the specified software source OCID. Filter the list against a variety of 
criteria including but not limited to its module name, stream name, and (profile) name.


## Example Usage

```hcl
data "oci_os_management_hub_software_source_module_stream_profiles" "test_software_source_module_stream_profiles" {
	#Required
	software_source_id = oci_os_management_hub_software_source.test_software_source.id

	#Optional
	module_name = var.software_source_module_stream_profile_module_name
	name = var.software_source_module_stream_profile_name
	stream_name = oci_streaming_stream.test_stream.name
}
```

## Argument Reference

The following arguments are supported:

* `module_name` - (Optional) The name of a module. This parameter is required if a streamName is specified. 
* `name` - (Optional) The name of the entity to be queried.
* `software_source_id` - (Required) The software source OCID.
* `stream_name` - (Optional) The name of the stream of the containing module.  This parameter is required if a profileName is specified. 


## Attributes Reference

The following attributes are exported:

* `module_stream_profile_collection` - The list of module_stream_profile_collection.

### SoftwareSourceModuleStreamProfile Reference

The following attributes are exported:

* `description` - A description of the contents of the module stream profile.
* `is_default` - Indicates if this profile is the default for its module stream.
* `module_name` - The name of the module that contains the stream profile.
* `name` - The name of the profile.
* `packages` - A list of packages that constitute the profile.  Each element in the list is the name of a package.  The name is suitable to use as an argument to other OS Management Hub APIs that interact directly with packages. 
* `stream_name` - The name of the stream that contains the profile.

