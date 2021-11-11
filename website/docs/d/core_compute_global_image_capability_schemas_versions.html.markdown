---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_compute_global_image_capability_schemas_versions"
sidebar_current: "docs-oci-datasource-core-compute_global_image_capability_schemas_versions"
description: |-
  Provides the list of Compute Global Image Capability Schemas Versions in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_compute_global_image_capability_schemas_versions
This data source provides the list of Compute Global Image Capability Schemas Versions in Oracle Cloud Infrastructure Core service.

Lists Compute Global Image Capability Schema versions in the specified compartment.


## Example Usage

```hcl
data "oci_core_compute_global_image_capability_schemas_versions" "test_compute_global_image_capability_schemas_versions" {
	#Required
	compute_global_image_capability_schema_id = oci_core_compute_global_image_capability_schema.test_compute_global_image_capability_schema.id

	#Optional
	display_name = var.compute_global_image_capability_schemas_version_display_name
}
```

## Argument Reference

The following arguments are supported:

* `compute_global_image_capability_schema_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute global image capability schema
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 


## Attributes Reference

The following attributes are exported:

* `compute_global_image_capability_schema_versions` - The list of compute_global_image_capability_schema_versions.

### ComputeGlobalImageCapabilitySchemasVersion Reference

The following attributes are exported:

* `compute_global_image_capability_schema_id` - The ocid of the compute global image capability schema 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `name` - The name of the compute global image capability schema version 
* `schema_data` - The map of each capability name to its ImageCapabilityDescriptor.
	* `default_value` - the default value
	* `descriptor_type` - The image capability schema descriptor type for the capability 
	* `source` - 
	* `values` - the list of values for the enum
* `time_created` - The date and time the compute global image capability schema version was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

