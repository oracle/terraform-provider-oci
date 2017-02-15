# baremetal\_core\_volume\_attachment

Provides a volue attachment resource

## Example Usage

```
resource "baremetal_core_volume_attachment" "t" {
    attachment_type = "attachment_type"
    compartment_id = "compartment_id"
    instance_id = "instance_id"
    volume_id = "volume_id"
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) The OCID of the compartment.
* `instance_id` - (Required) The OCID of the instance.
* `volume_id` - (Required) The OCID of the volume.
* `type` - (Required) The type of volume. The only supported value is "iscsi".


## Attributes Reference
* `attachment_type` - The type of volume attachment.
* `availability_domain` - The Availability Domain of an instance.
* `compartment_id` - The OCID of the compartment.
* `display_name` - A user-friendly name. Does not have to be unique, and it cannot be changed.
* `id` - The OCID of the volume attachment.
* `instance_id` - The OCID of the instance the volume is attached to.
* `state` - The current state of the volume attachment: [ATTACHING, ATTACHED, DETACHING, DETACHED].
* `time_created` - The date and time the volume was created
* `volume_id` - The OCID of the volume.


