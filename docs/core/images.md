# oci\_core\_image

## Image Resource

### Image Reference

The following attributes are exported:

* `base_image_id` - The OCID of the image originally used to launch the instance.
* `compartment_id` - The OCID of the compartment containing the instance you want to use as the basis for the image. 
* `create_image_allowed` - Whether instances launched with this image can be used to create new images. For example, you cannot create an image of an Oracle Database instance.  Example: `true` 
* `display_name` - A user-friendly name for the image. It does not have to be unique, and it's changeable. Avoid entering confidential information. You cannot use an Oracle-provided image name as a custom image name.  Example: `My custom Oracle Linux image` 
* `id` - The OCID of the image.
* `operating_system` - The image's operating system.  Example: `Oracle Linux` 
* `operating_system_version` - The image's operating system version.  Example: `7.2` 
* `state` - 
* `time_created` - The date and time the image was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 



### Create Operation
Creates a boot disk image for the specified instance or imports an exported image from the Oracle Cloud Infrastructure Object Storage service.

When creating a new image, you must provide the OCID of the instance you want to use as the basis for the image, and
the OCID of the compartment containing that instance. For more information about images,
see [Managing Custom Images](https://docs.us-phoenix-1.oraclecloud.com/Content/Compute/Tasks/managingcustomimages.htm).

You may optionally specify a *display name* for the image, which is simply a friendly name or description.
It does not have to be unique, and you can change it. See [UpdateImage](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Image/UpdateImage).
Avoid entering confidential information.


The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment containing the instance you want to use as the basis for the image.
* `display_name` - (Optional) A user-friendly name for the image. It does not have to be unique, and it's changeable. Avoid entering confidential information.  You cannot use an Oracle-provided image name as a custom image name.  Example: `My Oracle Linux image` 
* `instance_id` - (Required) The OCID of the instance you want to use as the basis for the image.


### Update Operation
Updates the display name of the image. Avoid entering confidential information.

The following arguments support updates:
* `display_name` - A user-friendly name for the image. It does not have to be unique, and it's changeable. Avoid entering confidential information.  You cannot use an Oracle-provided image name as a custom image name.  Example: `My Oracle Linux image` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```
resource "oci_core_image" "test_image" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.image_display_name}"
	instance_id = "${oci_core_instance.test_instance.id}"
}
```

# oci\_core\_images

## Image DataSource

Gets a list of images.

### List Operation
Lists the available images in the specified compartment.
If you specify a value for the `sortBy` parameter, Oracle-provided images appear first in the list, followed by custom images.
For more
information about images, see
[Managing Custom Images](https://docs.us-phoenix-1.oraclecloud.com/Content/Compute/Tasks/managingcustomimages.htm).

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `operating_system` - (Optional) The image's operating system.  Example: `Oracle Linux` 
* `operating_system_version` - (Optional) The image's operating system version.  Example: `7.2` 
* `state` - (Optional) A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive. 


The following attributes are exported:

* `images` - The list of images.

### Example Usage

```
data "oci_core_images" "test_images" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.image_display_name}"
	operating_system = "${var.image_operating_system}"
	operating_system_version = "${var.image_operating_system_version}"
	state = "${var.image_state}"
}
```