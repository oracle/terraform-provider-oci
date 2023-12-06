---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_compute_image_capability_schema"
sidebar_current: "docs-oci-datasource-core-compute_image_capability_schema"
description: |-
  Provides details about a specific Compute Image Capability Schema in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_compute_image_capability_schema
This data source provides details about a specific Compute Image Capability Schema resource in Oracle Cloud Infrastructure Core service.

Gets the specified Compute Image Capability Schema


## Example Usage

```hcl
data "oci_core_compute_image_capability_schema" "test_compute_image_capability_schema" {
	#Required
	compute_image_capability_schema_id = oci_core_compute_image_capability_schema.test_compute_image_capability_schema.id

	#Optional
	is_merge_enabled = var.compute_image_capability_schema_is_merge_enabled
}
```

## Argument Reference

The following arguments are supported:

* `compute_image_capability_schema_id` - (Required) The id of the compute image capability schema or the image ocid
* `is_merge_enabled` - (Optional) Merge the image capability schema with the global image capability schema 


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the compute global image capability schema 
* `compute_global_image_capability_schema_id` - The ocid of the compute global image capability schema 
* `compute_global_image_capability_schema_version_name` - The name of the compute global image capability schema version 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The compute image capability schema [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `image_id` - The OCID of the image associated with this compute image capability schema 
* `schema_data` - A mapping of each capability name to its ImageCapabilityDescriptor.
	* `default_value` - the default value
	* `descriptor_type` - The image capability schema descriptor type for the capability 
	* `source` - 
	* `values` - the list of values for the enum
* `time_created` - The date and time the compute image capability schema was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

