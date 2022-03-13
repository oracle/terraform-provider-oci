---
subcategory: "OS Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_osmanagement_managed_instance_stream_profiles"
sidebar_current: "docs-oci-datasource-osmanagement-managed_instance_stream_profiles"
description: |-
  Provides the list of Managed Instance Stream Profiles in Oracle Cloud Infrastructure OS Management service
---

# Data Source: oci_osmanagement_managed_instance_stream_profiles
This data source provides the list of Managed Instance Stream Profiles in Oracle Cloud Infrastructure OS Management service.

Retrieve a list of module stream profiles, along with a summary of their
of their status, from a managed instance.  Filters may be applied to
select a subset of profiles based on the filter criteria.

The "moduleName", "streamName", and "profileName" attributes combine
to form a set of filters on the list of module stream profiles.  If
a "modulName" is provided, only profiles that belong to that module
are returned.  If both a "moduleName" and "streamName" are given,
only profiles belonging to that module stream are returned.  Finally,
if all three are given then only the particular profile indicated
by the triple is returned.  It is not valid to supply a "streamName"
without a "moduleName".  It is also not valid to supply a "profileName"
without a "streamName".

The "status" attribute filters against the state of a module stream
profile.  Valid values are "INSTALLED" and "AVAILABLE".  If the
attribute is set to "INSTALLED", only module stream profiles that
are installed are included in the result set.  If the attribute is
set to "AVAILABLE", only module stream profiles that are not
installed are included in the result set.  If the attribute is not
defined, the request is not subject to this filter.

When sorting by display name, the result set is sorted first by
module name, then by stream name, and finally by profile name.


## Example Usage

```hcl
data "oci_osmanagement_managed_instance_stream_profiles" "test_managed_instance_stream_profiles" {
	#Required
	managed_instance_id = var.managed_instance_id

	#Optional
	compartment_id = var.compartment_id
	module_name = var.managed_instance_module_name
	profile_name = var.managed_instance_module_stream_profile_name
	profile_status = var.managed_instance_profile_status
	stream_name = var.managed_instance_module_stream_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which to list resources. This parameter is optional and in some cases may have no effect.
* `managed_instance_id` - (Required) OCID for the managed instance
* `module_name` - (Optional) The name of a module.  This parameter is required if a streamName is specified. 
* `profile_name` - (Optional) The name of the profile of the containing module stream
* `profile_status` - (Optional) The status of the profile.

	A profile with the "INSTALLED" status indicates that the profile has been installed.

	A profile with the "AVAILABLE" status indicates that the profile is not installed, but can be. 
* `stream_name` - (Optional) The name of the stream of the containing module.  This parameter is required if a profileName is specified. 


## Attributes Reference

The following attributes are exported:

* `module_stream_profile_on_managed_instances` - The list of module_stream_profile_on_managed_instances.

### ManagedInstanceStreamProfile Reference

The following attributes are exported:

* `module_name` - The name of the module that contains the stream profile
* `profile_name` - The name of the profile
* `status` - The status of the profile.

	A profile with the "INSTALLED" status indicates that the profile has been installed.

	A profile with the "AVAILABLE" status indicates that the profile is not installed, but can be. 
* `stream_name` - The name of the stream that contains the profile
* `time_modified` - The date and time of the last status change for this profile, as described in [RFC 3339](https://tools.ietf.org/rfc/rfc3339), section 14.29. 

