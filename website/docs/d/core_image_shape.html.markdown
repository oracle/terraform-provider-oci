---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_image_shape"
sidebar_current: "docs-oci-datasource-core-image_shape"
description: |-
  Provides details about a specific Image Shape in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_image_shape
This data source provides details about a specific Image Shape resource in Oracle Cloud Infrastructure Core service.

Retrieves an image shape compatibility entry.

## Example Usage

```hcl
data "oci_core_image_shape" "test_image_shape" {
	#Required
	image_id = "${oci_core_image.test_image.id}"
	shape_name = "${oci_core_shape.test_shape.name}"
}
```

## Argument Reference

The following arguments are supported:

* `image_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the image.
* `shape_name` - (Required) Shape name.


## Attributes Reference

The following attributes are exported:

* `image_id` - The image [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `ocpu_constraints` - 
	* `max` - The maximum number of OCPUs supported for this image and shape.
	* `min` - The minimum number of OCPUs supported for this image and shape.
* `shape` - The shape name.

