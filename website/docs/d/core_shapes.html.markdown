---
layout: "oci"
page_title: "OCI: oci_core_shapes"
sidebar_current: "docs-oci-datasource-core-shapes"
description: |-
  Provides a list of Shapes
---

# Data Source: oci_core_shapes
The Shapes data source allows access to the list of OCI shapes

Lists the shapes that can be used to launch an instance within the specified compartment. You can
filter the list by compatibility with a specific image.


## Example Usage

```hcl
data "oci_core_shapes" "test_shapes" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	availability_domain = "${var.shape_availability_domain}"
	image_id = "${oci_core_image.test_image.id}"
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) The name of the Availability Domain.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) The OCID of the compartment.
* `image_id` - (Optional) The OCID of an image.


## Attributes Reference

The following attributes are exported:

* `shapes` - The list of shapes.

### Shape Reference

The following attributes are exported:

* `name` - The name of the shape. You can enumerate all available shapes by calling [ListShapes](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Shape/ListShapes). 

