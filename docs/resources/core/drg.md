# oci\_core\_drg

Provide a drg attachment resource.

## Example Usage

```
resource "oci_core_drg" "t" {
    compartment_id = "compartment_id"
    display_name = "display_name"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment to contain the DRG.
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable.

## Attributes Reference
* `compartment_id` - The OCID of the compartment containing the DRG.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `id` - The DRG's Oracle ID (OCID).
* `state` - The DRG's current state: [PROVISIONING, AVAILABLE, TERMINATING, TERMINATED].
* `time_created` - The date and time the image was created.

