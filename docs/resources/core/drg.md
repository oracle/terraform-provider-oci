# oci\_core\_drg

[Drg Reference][56529582]

  [56529582]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Drg/ "DrgReference"

Provide a Dynamic Routing Gateway (DRG) resource.

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
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.

## Attributes Reference
* `compartment_id` - The OCID of the compartment containing the DRG.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `id` - The DRG's Oracle ID (OCID).
* `state` - The DRG's current state. Allowed values are: [PROVISIONING, AVAILABLE, TERMINATING, TERMINATED].
* `time_created` - The date and time the DRG was created, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z`.
