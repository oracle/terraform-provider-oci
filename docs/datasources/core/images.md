# oci\_core\_images

Gets a list of images.

## Example Usage

```
data "oci_core_images" "t" {
  compartment_id = "compartment_id"
  limit = 1
  page = "page"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable.
* `operating_system` - (Optional) The image's operating system.
* `operating_system_version` - (Optional) The image's operating system version.
* `limit` - (Optional) The maximum number of items to return in a paginated "List" call.
* `page` - (Optional) The page to fetch.

## Attributes Reference

The following attributes are exported:

* `images` - The list of images.

## Image reference
* `base_image_id` - The OCID of the image originally used to launch the instance.
* `compartment_id` - The OCID of the compartment containing the instance you want to use as the basis for the image.
* `create_image_allowed` - Whether instances launched with this image can be used to create new images.
* `display_name` - A user-friendly name for the image. It does not have to be unique, and it's changeable.
* `id` - The OCID of the image.
* `state` - The state of the image: [PROVISIONING, AVAILABLE, DISABLED, DELETED].
* `operating_system` - The image's operating system.
* `operating_system_version` - The image's operating system version.
* `time_created` - The date and time the image was created.
