---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_image_shapes"
sidebar_current: "docs-oci-datasource-core-image_shapes"
description: |-
  Provides the list of Image Shapes in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_image_shapes
This data source provides the list of Image Shapes in Oracle Cloud Infrastructure Core service.

Lists the compatible shapes for the specified image.

## Example Usage

```hcl
data "oci_core_image_shapes" "test_image_shapes" {
	#Required
	image_id = oci_core_image.test_image.id
}
```

## Argument Reference

The following arguments are supported:

* `image_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the image.


## Attributes Reference

The following attributes are exported:

* `image_shape_compatibilities` - The list of image_shape_compatibilities.

### ImageShape Reference

The following attributes are exported:

* `image_id` - The image [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `memory_constraints` - 
	* `max_in_gbs` - The maximum amount of memory, in gigabytes.
	* `min_in_gbs` - The minimum amount of memory, in gigabytes.
* `ocpu_constraints` - 
	* `max` - The maximum number of OCPUs supported for this image and shape.
	* `min` - The minimum number of OCPUs supported for this image and shape.
* `shape` - The shape name.

