---
subcategory: "OS Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_osmanagement_managed_instance_module_streams"
sidebar_current: "docs-oci-datasource-osmanagement-managed_instance_module_streams"
description: |-
  Provides the list of Managed Instance Module Streams in Oracle Cloud Infrastructure OS Management service
---

# Data Source: oci_osmanagement_managed_instance_module_streams
This data source provides the list of Managed Instance Module Streams in Oracle Cloud Infrastructure OS Management service.

Retrieve a list of module streams, along with a summary of their
status, from a managed instance.  Filters may be applied to select
a subset of module streams based on the filter criteria.

The 'moduleName' attribute filters against the name of a module.
It accepts strings of the format "<module>".  If this attribute
is defined, only streams that belong to the specified module are
included in the result set.  If it is not defined, the request is
not subject to this filter.

The "status" attribute filters against the state of a module stream.
Valid values are "ENABLED", "DISABLED", and "ACTIVE".  If the
attribute is set to "ENABLED", only module streams that are enabled
are included in the result set.  If the attribute is set to "DISABLED",
only module streams that are not enabled are included in the result
set.  If the attribute is set to "ACTIVE", only module streams that
are active are included in the result set.  If the attribute is not
defined, the request is not subject to this filter.

When sorting by the display name, the result set is sorted first
by the module name and then by the stream name.


## Example Usage

```hcl
data "oci_osmanagement_managed_instance_module_streams" "test_managed_instance_module_streams" {
	#Required
	managed_instance_id = var.managed_instance_id

	#Optional
	compartment_id = var.compartment_id
	module_name = var.managed_instance_module_name
	stream_name = var.managed_instance_module_stream_name
	stream_status = var.managed_instance_module_stream_status
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which to list resources. This parameter is optional and in some cases may have no effect.
* `managed_instance_id` - (Required) OCID for the managed instance
* `module_name` - (Optional) The name of a module.  This parameter is required if a streamName is specified. 
* `stream_name` - (Optional) The name of the stream of the containing module.  This parameter is required if a profileName is specified. 
* `stream_status` - (Optional) The status of the stream

	A stream with the "ENABLED" status can be used as a source for installing profiles.  Streams with this status are also "ACTIVE".

	A stream with the "DISABLED" status cannot be the source for installing profiles.  To install profiles and packages from this stream, it must be enabled.

	A stream with the "ACTIVE" status can be used as a source for installing profiles.  The packages that comprise the stream are also used when a matching package is installed directly.  In general, a stream can have this status if it is the default stream for the module and no stream has been explicitly enabled. 


## Attributes Reference

The following attributes are exported:

* `module_stream_on_managed_instances` - The list of module_stream_on_managed_instances.

### ManagedInstanceModuleStream Reference

The following attributes are exported:

* `module_name` - The name of the module that contains the stream. 
* `profiles` - The set of profiles that the module stream contains.
	* `module_name` - The name of the module that contains the stream profile
	* `profile_name` - The name of the profile
	* `status` - The status of the profile.

		A profile with the "INSTALLED" status indicates that the profile has been installed.

		A profile with the "AVAILABLE" status indicates that the profile is not installed, but can be. 
	* `stream_name` - The name of the stream that contains the profile
	* `time_modified` - The date and time of the last status change for this profile, as described in [RFC 3339](https://tools.ietf.org/rfc/rfc3339), section 14.29. 
* `software_source_id` - The OCID of the software source that provides this module stream.
* `status` - The status of the stream

	A stream with the "ENABLED" status can be used as a source for installing profiles.  Streams with this status are also "ACTIVE".

	A stream with the "DISABLED" status cannot be the source for installing profiles.  To install profiles and packages from this stream, it must be enabled.

	A stream with the "ACTIVE" status can be used as a source for installing profiles.  The packages that comprise the stream are also used when a matching package is installed directly.  In general, a stream can have this status if it is the default stream for the module and no stream has been explicitly enabled. 
* `stream_name` - The name of the stream. 
* `time_modified` - The date and time of the last status change for this profile, as described in [RFC 3339](https://tools.ietf.org/rfc/rfc3339), section 14.29. 

