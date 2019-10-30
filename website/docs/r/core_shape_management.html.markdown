---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_shape_management"
sidebar_current: "docs-oci-resource-core-shape_management"
description: |-
  Provides details about a specific Shape in Oracle Cloud Infrastructure Core service
---

# oci_core_shape_management
This resource provides the Shape Management resource in Oracle Cloud Infrastructure Core service.

Add/Remove the specified shape from the compatible shapes list for the image.

## Example Usage

```hcl
"oci_core_shape_management" "test_shape" {
	#Required
	compartment_id = "${var.compartment_id}"
	image_id = "${oci_core_image.test_image.id}"
	shape_name = "${var.shape_name}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment containing the image.
* `image_id` - (Required) The OCID of the Image to which the shape should be added.
* `shape_name` - (Required) The compatible shape that is to be added to the compatible shapes list for the image. 

## Attributes Reference

The following attributes are exported:

* `id` - The image's Oracle ID (OCID).
* `image_id` - The OCID of the image containing the shape.
* `shape_name` - The compatible Shape for the image.  
