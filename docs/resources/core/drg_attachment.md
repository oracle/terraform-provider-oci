# oci\_drg\_attachment

[DrgAttachment Reference][6b5d9217]

  [6b5d9217]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/DrgAttachment/ "DrgAttachmentReference"

Provide a DRG attachment resource.

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
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `drg_id` - The OCID of the DRG.
* `id` - The DRG attachment's Oracle ID (OCID).
* `state` - The DRG attachment's current state: [ATTACHING, ATTACHED, DETACHING, DETACHED].
* `time_created` - The date and time the DRG attachment was created, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z`.
* `vcn_id` - The OCID of the VCN.
