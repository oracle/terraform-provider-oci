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

Attaches the specified DRG to the specified network resource. A VCN can be attached to only one DRG
at a time, but a DRG can be attached to more than one VCN. The response includes a `DrgAttachment`
object with its own [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). For more information about DRGs, see
[Dynamic Routing Gateways (DRGs)](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingDRGs.htm).

You may optionally specify a *display name* for the attachment, otherwise a default is provided.
It does not have to be unique, and you can change it. Avoid entering confidential information.

For the purposes of access control, the DRG attachment is automatically placed into the currently selected compartment.
For more information about compartments and access control, see
[Overview of the IAM Service](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm).


## Example Usage

```hcl
resource "oci_core_drg_attachment" "test_drg_attachment" {
	#Required
	drg_id = oci_core_drg.test_drg.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.drg_attachment_display_name
	drg_route_table_id = oci_core_drg_route_table.test_drg_route_table.id
	freeform_tags = {"Department"= "Finance"}
	network_details {
		#Required
		id = oci_core_vcn.test_vcn.id
		type = var.drg_attachment_network_details_type

		#Optional
		route_table_id = oci_core_route_table.test_route_table.id
	}
}
```

## Argument Reference

The following arguments are supported:

* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `drg_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG.
* `drg_route_table_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG route table that is assigned to this attachment.

	The DRG route table manages traffic inside the DRG.
* `export_drg_route_distribution_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the export route distribution used to specify how routes in the assigned DRG route table are advertised to the attachment. If this value is null, no routes are advertised through this attachment.
    This field cannot be set by the user while creating the resource and gets a default value on creation. This can be only be updated to its default value. If this fields needs to be set to null, remove_export_drg_route_distribution_trigger needs to be used.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `network_details` - (Optional) (Updatable) 
	* `id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network attached to the DRG. 
	* `route_table_id` - (Optional) (Updatable) This is the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the route table that is used to route the traffic as it enters a VCN through this attachment.

		For information about why you would associate a route table with a DRG attachment, see [Advanced Scenario: Transit Routing](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/transitrouting.htm). For information about why you would associate a route table with a DRG attachment, see:
		* [Transit Routing: Access to Multiple VCNs in Same Region](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/transitrouting.htm)
		* [Transit Routing: Private Access to Oracle Services](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/transitroutingoracleservices.htm) 
	* `type` - (Required) (Updatable) The type can be one of these values: `IPSEC_TUNNEL`, `REMOTE_PEERING_CONNECTION`, `VCN`, `VIRTUAL_CIRCUIT`
* `route_table_id` - (Deprecated) (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the route table used by the DRG attachment.

	If you don't specify a route table here, the DRG attachment is created without an associated route table. The Networking service does NOT automatically associate the attached VCN's default route table with the DRG attachment. For information about why you would associate a route table with a DRG attachment, see:
	* [Transit Routing: Access to Multiple VCNs in Same Region](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/transitrouting.htm)
	* [Transit Routing: Private Access to Oracle Services](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/transitroutingoracleservices.htm) 
	This field is deprecated. Instead, use the networkDetails field to specify the VCN route table for this attachment. 
* `vcn_id` - (Deprecated) (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN. This field is deprecated. Instead, use the `networkDetails` field to specify the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the attached resource. 
* `remove_export_drg_route_distribution_trigger` - (Optional) (Updatable) An optional property when set to true during update disables the export of route Distribution by setting export_drg_route_distribution_id to null.

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the DRG attachment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `drg_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG.
* `drg_route_table_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG route table that is assigned to this attachment.

	The DRG route table manages traffic inside the DRG. 
* `export_drg_route_distribution_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the export route distribution used to specify how routes in the assigned DRG route table are advertised to the attachment. If this value is null, no routes are advertised through this attachment. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The DRG attachment's Oracle ID ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)).
* `is_cross_tenancy` - Indicates whether the DRG attachment and attached network live in a different tenancy than the DRG.  Example: `false` 
* `network_details` - 
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network attached to the DRG. 
	* `ipsec_connection_id` - The IPSec connection that contains the attached IPSec tunnel.
	* `route_table_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the route table the DRG attachment is using.

		For information about why you would associate a route table with a DRG attachment, see:
		* [Transit Routing: Access to Multiple VCNs in Same Region](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/transitrouting.htm)
		* [Transit Routing: Private Access to Oracle Services](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/transitroutingoracleservices.htm)
	* `type` - The type can be one of these values: `IPSEC_TUNNEL`, `REMOTE_PEERING_CONNECTION`, `VCN`, `VIRTUAL_CIRCUIT`
* `route_table_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the route table the DRG attachment is using.

	For information about why you would associate a route table with a DRG attachment, see:
	* [Transit Routing: Access to Multiple VCNs in Same Region](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/transitrouting.htm)
	* [Transit Routing: Private Access to Oracle Services](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/transitroutingoracleservices.htm)

	This field is deprecated. Instead, use the `networkDetails` field to view the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the attached resource. 
* `state` - The DRG attachment's current state.
* `time_created` - The date and time the DRG attachment was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `vcn_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN. This field is deprecated. Instead, use the `networkDetails` field to view the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the attached resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Drg Attachment
	* `update` - (Defaults to 20 minutes), when updating the Drg Attachment
	* `delete` - (Defaults to 20 minutes), when destroying the Drg Attachment


## Import

DrgAttachments can be imported using the `id`, e.g.

```
$ terraform import oci_core_drg_attachment.test_drg_attachment "id"
```

