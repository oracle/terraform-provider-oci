---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_compute_image_capability_schemas"
sidebar_current: "docs-oci-datasource-core-compute_image_capability_schemas"
description: |-
  Provides the list of Compute Image Capability Schemas in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_compute_image_capability_schemas
This data source provides the list of Compute Image Capability Schemas in Oracle Cloud Infrastructure Core service.

Lists Compute Image Capability Schema in the specified compartment. You can also query by a specific imageId.


## Example Usage

```hcl
data "oci_core_compute_image_capability_schemas" "test_compute_image_capability_schemas" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.compute_image_capability_schema_display_name
	image_id = oci_core_image.test_image.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) A filter to return only resources that match the given compartment OCID exactly. 
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `image_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of an image.


## Attributes Reference

The following attributes are exported:

* `compute_image_capability_schemas` - The list of compute_image_capability_schemas.

### ComputeImageCapabilitySchema Reference

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

