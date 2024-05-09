---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_managed_instance_group_modules"
sidebar_current: "docs-oci-datasource-os_management_hub-managed_instance_group_modules"
description: |-
  Provides the list of Managed Instance Group Modules in Oracle Cloud Infrastructure Os Management Hub service
---

# Data Source: oci_os_management_hub_managed_instance_group_modules
This data source provides the list of Managed Instance Group Modules in Oracle Cloud Infrastructure Os Management Hub service.

Retrieve a list of module streams, along with a summary of their
status, from a managed instance group.  Filters may be applied to select
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
data "oci_os_management_hub_managed_instance_group_modules" "test_managed_instance_group_modules" {
	#Required
	managed_instance_group_id = oci_os_management_hub_managed_instance_group.test_managed_instance_group.id

	#Optional
	compartment_id = var.compartment_id
	name = var.managed_instance_group_module_name
	name_contains = var.managed_instance_group_module_name_contains
	stream_name = oci_streaming_stream.test_stream.name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The OCID of the compartment that contains the resources to list. This filter returns only resources contained within the specified compartment.
* `managed_instance_group_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance group.
* `name` - (Optional) The resource name. 
* `name_contains` - (Optional) A filter to return resources that may partially match the name given.
* `stream_name` - (Optional) The name of the module stream. This parameter is required if a profile name is specified.


## Attributes Reference

The following attributes are exported:

* `managed_instance_group_module_collection` - The list of managed_instance_group_module_collection.

### ManagedInstanceGroupModule Reference

The following attributes are exported:

* `items` - List of module streams.
	* `enabled_stream` - The name of the module stream that is enabled for the group.
	* `installed_profiles` - The list of installed profiles under the currently enabled module stream.
	* `name` - The name of the module.
	* `software_source_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source that provides this module stream.

