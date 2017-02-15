# baremetal\_core\_cpe

Provide a CPE resource.

## Example Usage

```
resource "baremetal_core_cpe" "t" {
    compartment_id = "compartmentid"
    display_name = "displayname"
    ip_address = "123.123.123.123"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `ip_address` - (Required) The public IP address of the on-premise router.
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable.

## Attributes Reference
* `compartment_id` - The OCID of the compartment containing the CPE.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `id` - The CPE's Oracle ID (OCID).
* `ip_address` - The public IP address of the on-premise router.
* `time_created` - The date and time the image was created.
