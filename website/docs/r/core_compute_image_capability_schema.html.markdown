---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_compute_image_capability_schema"
sidebar_current: "docs-oci-resource-core-compute_image_capability_schema"
description: |-
  Provides the Compute Image Capability Schema resource in Oracle Cloud Infrastructure Core service
---

# oci_core_compute_image_capability_schema
This resource provides the Compute Image Capability Schema resource in Oracle Cloud Infrastructure Core service.

Creates compute image capability schema.


## Example Usage

```hcl
resource "oci_core_compute_image_capability_schema" "test_compute_image_capability_schema" {
	#Required
	compartment_id = var.compartment_id
	compute_global_image_capability_schema_version_name = var.compute_image_capability_schema_compute_global_image_capability_schema_version_name
	image_id = oci_core_image.test_image.id
	schema_data {
		#Required
		descriptor_type = var.compute_image_capability_schema_schema_data_descriptor_type
		source = var.compute_image_capability_schema_schema_data_source

		#Optional
		default_value = var.compute_image_capability_schema_schema_data_default_value
		values = var.compute_image_capability_schema_schema_data_values
	}

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.compute_image_capability_schema_display_name
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment that contains the resource.
* `compute_global_image_capability_schema_version_name` - (Required) The name of the compute global image capability schema version 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `image_id` - (Required) The ocid of the image 
* `schema_data` - (Required) (Updatable) The map of each capability name to its ImageCapabilitySchemaDescriptor.
	* `default_value` - (Optional) (Updatable) the default value
	* `descriptor_type` - (Required) (Updatable) The image capability schema descriptor type for the capability 
	* `source` - (Required) (Updatable) 
	* `values` - (Required when descriptor_type=enuminteger | enumstring) (Updatable) the list of values for the enum


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Compute Image Capability Schema
	* `update` - (Defaults to 20 minutes), when updating the Compute Image Capability Schema
	* `delete` - (Defaults to 20 minutes), when destroying the Compute Image Capability Schema


## Import

ComputeImageCapabilitySchemas can be imported using the `id`, e.g.

```
$ terraform import oci_core_compute_image_capability_schema.test_compute_image_capability_schema "id"
```

