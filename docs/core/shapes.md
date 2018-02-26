
# oci\_core\_shapes

## Shape DataSource

Gets a list of shapes.

### List Operation
Lists the shapes that can be used to launch an instance within the specified compartment. You can
filter the list by compatibility with a specific image.

The following arguments are supported:

* `availability_domain` - (Optional) The name of the Availability Domain.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) The OCID of the compartment.
* `image_id` - (Optional) The OCID of an image.


The following attributes are exported:

* `shapes` - The list of shapes.

### Example Usage

```
data "oci_core_shapes" "test_shapes" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	availability_domain = "${var.shape_availability_domain}"
	image_id = "${var.shape_image_id}"
}
```