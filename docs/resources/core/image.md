# baremetal\_core\_image

Provide an image resource.

## Example Usage

```
resource "baremetal_core_image" "t" {
    compartment_id = "compartment_id"
    display_name = "display_name"
    instance_id = "instance_id"
}

```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment containing the instance you want to use as the basis for the image.
* `display_name` - (Optional) A user-friendly name for the image. It does not have to be unique, and it's changeable.
* `instance_id` - (Required) The OCID of the instance you want to use as the basis for the image.

## Attributes Reference
* `base_image_id` - The OCID of the image originally used to launch the instance.
* `compartment_id` - The OCID of the compartment containing the instance you want to use as the basis for the image.
* `create_image_allowed` - Whether instances launched with this image can be used to create new images.
* `display_name` - A user-friendly name for the image. It does not have to be unique, and it's changeable.
* `id` - The OCID of the image.
* `state` - The state of the image: [PROVISIONING, AVAILABLE, DISABLED, DELETED].
* `operating_system` - The image's operating system.
* `operating_system_version` - The image's operating system version.
* `time_created` - The date and time the image was created.
