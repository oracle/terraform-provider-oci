
# oci\_core\_instances

Gets a list of instances.

## Example Usage

```
data "oci_core_instances" "s" {
  compartment_id = "compartmentid"
  availability_domain = "availabilityid"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `availability_domain` - (Optional) The name of the Availability Domain.
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable.
* `limit` - (Optional) The maximum number of items to return in a paginated "List" call.
* `page` - (Optional) Length of the snapshot data to retrieve.

## Attributes Reference

The following attributes are exported:

* `instances` - The list of instances.

## Instance Reference
* `availability_domain` - The Availability Domain the instance is running in.
* `compartment_id` - The OCID of the compartment that contains the instance.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `id` - The OCID of the instance.
* `image_id` - The image used to boot the instance. You can enumerate all available images by calling ListImages.
* `state` - The current state of the instance: [PROVISIONING, RUNNING, STARTING, STOPPING, STOPPED, CREATING_IMAGE, TERMINATING, TERMINATED]
* `metadata` - Custom metadata that you provide.
* `extended_metadata` - Custom nested metadata that you provide.
* `region` - The region that contains the Availability Domain the instance is running in.
* `shape` - The shape of the instance. The shape determines the number of CPUs and the amount of memory allocated to the instance.
* `time_created` - The date and time the instance was created.
