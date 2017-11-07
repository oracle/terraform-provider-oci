# oci\_core\_image

[Image Reference][9da3c3c9]

  [9da3c3c9]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Image/ "ImageReference"

Provide an image resource.

## Example Usage

```
resource "oci_core_image" "t" {
    compartment_id = "compartment_id"
    display_name = "display_name"
    instance_id = "instance_id"
}

```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment containing the instance you want to use as the basis for the image.
* `display_name` - (Optional) A user-friendly name for the image. It does not have to be unique, and it's changeable. Avoid entering confidential information. You **cannot** use an Oracle-provided image name as a custom image name.
* `instance_id` - (Required) The OCID of the instance you want to use as the basis for the image.

## Attributes Reference
* `base_image_id` - The OCID of the image originally used to launch the instance.
* `compartment_id` - The OCID of the compartment containing the instance you want to use as the basis for the image.
* `create_image_allowed` - Whether instances launched with this image can be used to create new images. Example: `true`
* `display_name` - A user-friendly name for the image. It does not have to be unique, and it's changeable. Avoid entering confidential information.
* `id` - The OCID of the image.
* `state` - The state of the image. Allowed values are: [PROVISIONING, IMPORTING, AVAILABLE, EXPORTING, DISABLED, DELETED].
* `operating_system` - The image's operating system.
* `operating_system_version` - The image's operating system version.
* `time_created` - The date and time the image was created, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z`.
