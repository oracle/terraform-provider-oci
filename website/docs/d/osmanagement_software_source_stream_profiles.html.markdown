---
subcategory: "OS Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_osmanagement_software_source_stream_profiles"
sidebar_current: "docs-oci-datasource-osmanagement-software_source_stream_profiles"
description: |-
  Provides the list of Software Source Stream Profiles in Oracle Cloud Infrastructure OS Management service
---

# Data Source: oci_osmanagement_software_source_stream_profiles
This data source provides the list of Software Source Stream Profiles in Oracle Cloud Infrastructure OS Management service.

Retrieve a list of module stream profiles from a software source.
Filters may be applied to select a subset of module stream profiles
based on the filter criteria.

The "moduleName", "streamName", and "profileName" attributes combine
to form a set of filters on the list of module stream profiles.  If
a "moduleName" is provided, only profiles that belong to that module
are returned.  If both a "moduleName" and "streamName" are given,
only profiles belonging to that module stream are returned.  Finally,
if all three are given then only the particular profile indicated
by the triple is returned.  It is not valid to supply a "streamName"
without a "moduleName".  It is also not valid to supply a "profileName"
without a "streamName".


## Example Usage

```hcl
data "oci_osmanagement_software_source_stream_profiles" "test_software_source_stream_profiles" {
	#Required
	software_source_id = var.software_source.id

	#Optional
	compartment_id = var.compartment_id
	module_name = var.software_source_module_name
	profile_name = var.software_source_module_stream_profile_name
	stream_name = var.software_source_module_stream_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which to list resources. This parameter is optional and in some cases may have no effect.
* `module_name` - (Optional) The name of a module.  This parameter is required if a streamName is specified. 
* `profile_name` - (Optional) The name of the profile of the containing module stream
* `software_source_id` - (Required) The OCID of the software source.
* `stream_name` - (Optional) The name of the stream of the containing module.  This parameter is required if a profileName is specified. 


## Attributes Reference

The following attributes are exported:

* `module_stream_profiles` - The list of module_stream_profiles.

### SoftwareSourceStreamProfile Reference

The following attributes are exported:

* `module_name` - The name of the module that contains the stream profile
* `profile_name` - The name of the profile
* `stream_name` - The name of the stream that contains the profile

