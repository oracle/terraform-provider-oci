---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_software_source_module_streams"
sidebar_current: "docs-oci-datasource-os_management_hub-software_source_module_streams"
description: |-
  Provides the list of Software Source Module Streams in Oracle Cloud Infrastructure Os Management Hub service
---

# Data Source: oci_os_management_hub_software_source_module_streams
This data source provides the list of Software Source Module Streams in Oracle Cloud Infrastructure Os Management Hub service.

Lists module streams from the specified software source [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 
Filter the list against a variety of criteria including but not limited to its module name and (stream) name.


## Example Usage

```hcl
data "oci_os_management_hub_software_source_module_streams" "test_software_source_module_streams" {
	#Required
	software_source_id = oci_os_management_hub_software_source.test_software_source.id

	#Optional
	is_latest = var.software_source_module_stream_is_latest
	module_name = var.software_source_module_stream_module_name
	module_name_contains = var.software_source_module_stream_module_name_contains
	name = var.software_source_module_stream_name
}
```

## Argument Reference

The following arguments are supported:

* `is_latest` - (Optional) Indicates whether to list only the latest versions of packages, module streams, and stream profiles.
* `module_name` - (Optional) The name of a module. This parameter is required if a streamName is specified. 
* `module_name_contains` - (Optional) A filter to return resources that may partially match the module name given.
* `name` - (Optional) The name of the entity to be queried.
* `software_source_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source.


## Attributes Reference

The following attributes are exported:

* `module_stream_collection` - The list of module_stream_collection.

### SoftwareSourceModuleStream Reference

The following attributes are exported:

* `arch_type` - The architecture for which the packages in this module stream were built.
* `description` - A description of the contents of the module stream.
* `is_default` - Indicates if this stream is the default for its module.
* `is_latest` - Indicates whether this module stream is the latest.
* `module_name` - The name of the module that contains the stream.
* `name` - The name of the stream.
* `packages` - A list of packages that are contained by the stream.  Each element in the list is the name of a package.  The name is suitable to use as an argument to other OS Management Hub APIs that interact directly with packages. 
* `profiles` - A list of profiles that are part of the stream.  Each element in the list is the name of a profile.  The name is suitable to use as an argument to other OS Management Hub APIs that interact directly with module stream profiles.  However, it is not URL encoded. 
* `software_source_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source that provides this module stream.

