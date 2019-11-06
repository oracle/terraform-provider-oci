---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_drg_attachment"
sidebar_current: "docs-oci-resource-core-drg_attachment"
description: |-
  Provides the Drg Attachment resource in Oracle Cloud Infrastructure Core service
---

# oci_core_drg_attachment
This resource provides the Drg Attachment resource in Oracle Cloud Infrastructure Core service.

Attaches the specified DRG to the specified VCN. A VCN can be attached to only one DRG at a time,
and vice versa. The response includes a `DrgAttachment` object with its own OCID. For more
information about DRGs, see
[Dynamic Routing Gateways (DRGs)](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingDRGs.htm).

You may optionally specify a *display name* for the attachment, otherwise a default is provided.
It does not have to be unique, and you can change it. Avoid entering confidential information.

For the purposes of access control, the DRG attachment is automatically placed into the same compartment
as the VCN. For more information about compartments and access control, see
[Overview of the IAM Service](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm).


## Example Usage

```hcl
resource "oci_core_drg_attachment" "test_drg_attachment" {
	#Required
	drg_id = "${oci_core_drg.test_drg.id}"
	vcn_id = "${oci_core_vcn.test_vcn.id}"

	#Optional
	display_name = "${var.drg_attachment_display_name}"
	route_table_id = "${oci_core_route_table.test_route_table.id}"
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique. Avoid entering confidential information.
* `drg_id` - (Required) The OCID of the DRG.
* `route_table_id` - (Optional) (Updatable) The OCID of the route table the DRG attachment will use.

	If you don't specify a route table here, the DRG attachment is created without an associated route table. The Networking service does NOT automatically associate the attached VCN's default route table with the DRG attachment.

	For information about why you would associate a route table with a DRG attachment, see:
	* [Transit Routing: Access to Multiple VCNs in Same Region](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/transitrouting.htm)
	* [Transit Routing: Private Access to Oracle Services](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/transitroutingoracleservices.htm) 
* `vcn_id` - (Required) The OCID of the VCN.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the DRG attachment.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `drg_id` - The OCID of the DRG.
* `id` - The DRG attachment's Oracle ID (OCID).
* `route_table_id` - The OCID of the route table the DRG attachment is using.

	For information about why you would associate a route table with a DRG attachment, see:
	* [Transit Routing: Access to Multiple VCNs in Same Region](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/transitrouting.htm)
	* [Transit Routing: Private Access to Oracle Services](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/transitroutingoracleservices.htm) 
* `state` - The DRG attachment's current state.
* `time_created` - The date and time the DRG attachment was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `vcn_id` - The OCID of the VCN.

## Import

DrgAttachments can be imported using the `id`, e.g.

```
$ terraform import oci_core_drg_attachment.test_drg_attachment "id"
```

