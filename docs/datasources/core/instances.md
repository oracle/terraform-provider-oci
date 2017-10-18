
# oci\_core\_instances

**API:** [Instance Reference][af7539b8]

  [af7539b8]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Instance/ "InstanceReference"

Gets a list of compute instances.

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
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `limit` - (Optional) The maximum number of items to return in a paginated "List" call.
* `page` - (Optional) Length of the snapshot data to retrieve.

## Attributes Reference

The following attributes are exported:

* `instances` - The list of instances.

## Instance Reference
* `availability_domain` - The Availability Domain the instance is running in.
* `compartment_id` - The OCID of the compartment that contains the instance.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `id` - The OCID of the instance.
* `image_id` - The image used to boot the instance. You can enumerate all available images by calling [ListImages][d198fa10].
* `state` - The current state of the instance. Allowed values are: [PROVISIONING, RUNNING, STARTING, STOPPING, STOPPED, CREATING_IMAGE, TERMINATING, TERMINATED]
* `metadata` - Custom metadata that you provide.
* `extended_metadata` - Custom nested metadata that you provide.
* `region` - The region that contains the Availability Domain the instance is running in.
* `shape` - The shape of the instance. The shape determines the number of CPUs and the amount of memory allocated to the instance.
* `time_created` - The date and time the instance was created,  in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z`.

  [d198fa10]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Image/ListImages "ListImages"
