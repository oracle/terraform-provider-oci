# oci\_drg\_attachment

Provide a drg attachment resource.

## Example Usage

```
resource "oci_core_drg_attachment" "t" {
    compartment_id = "compartment_id"
    display_name = "display_name"
    drg_id = "drg_id"
    vcn_id = "vcn_id"
}
```

## Argument Reference

The following arguments are supported:

* `vcn_id` - (Required) The OCID of the VCN.
* `drg_id` - (Required) The OCID of the DRG.
* `display_name` - (Optional) The OCID of the compartment.

## Attributes Reference
* `compartment_id` - The OCID of the compartment containing the DRG attachment.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `drg_id` - The OCID of the DRG.
* `id` - The DRG attachment's Oracle ID (OCID).
* `state` - The DRG attachment's current state: [ATTACHING, ATTACHED, DETACHING, DETACHED].
* `time_created` - The date and time the image was created.
* `vcn_id` - The OCID of the VCN.

