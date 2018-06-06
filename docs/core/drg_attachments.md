# oci_core_drg_attachment

## DrgAttachment Resource

### DrgAttachment Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the DRG attachment.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `drg_id` - The OCID of the DRG.
* `id` - The DRG attachment's Oracle ID (OCID).
* `state` - The DRG attachment's current state.
* `time_created` - The date and time the DRG attachment was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `vcn_id` - The OCID of the VCN.



### Create Operation
Attaches the specified DRG to the specified VCN. A VCN can be attached to only one DRG at a time,
and vice versa. The response includes a `DrgAttachment` object with its own OCID. For more
information about DRGs, see
[Dynamic Routing Gateways (DRGs)](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingDRGs.htm).

You may optionally specify a *display name* for the attachment, otherwise a default is provided.
It does not have to be unique, and you can change it. Avoid entering confidential information.

For the purposes of access control, the DRG attachment is automatically placed into the same compartment
as the VCN. For more information about compartments and access control, see
[Overview of the IAM Service](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/overview.htm).


The following arguments are supported:

* `display_name` - (Optional) A user-friendly name. Does not have to be unique. Avoid entering confidential information.
* `drg_id` - (Required) The OCID of the DRG.
* `vcn_id` - (Required) The OCID of the VCN.


### Update Operation
Updates the display name for the specified `DrgAttachment`.
Avoid entering confidential information.


The following arguments support updates:
* `display_name` - A user-friendly name. Does not have to be unique. Avoid entering confidential information.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```hcl
resource "oci_core_drg_attachment" "test_drg_attachment" {
	#Required
	drg_id = "${oci_core_drg.test_drg.id}"
	vcn_id = "${oci_core_vcn.test_vcn.id}"

	#Optional
	display_name = "${var.drg_attachment_display_name}"
}
```

# oci_core_drg_attachments

## DrgAttachment DataSource

Gets a list of drg_attachments.

### List Operation
Lists the `DrgAttachment` objects for the specified compartment. You can filter the
results by VCN or DRG.

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `drg_id` - (Optional) The OCID of the DRG.
* `vcn_id` - (Optional) The OCID of the VCN.


The following attributes are exported:

* `drg_attachments` - The list of drg_attachments.

### Example Usage

```hcl
data "oci_core_drg_attachments" "test_drg_attachments" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	drg_id = "${oci_core_drg.test_drg.id}"
	vcn_id = "${oci_core_vcn.test_vcn.id}"
}
```