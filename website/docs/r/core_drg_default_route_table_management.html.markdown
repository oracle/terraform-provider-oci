---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_drg_default_route_table_management"
sidebar_current: "docs-oci-resource-core-drg-default-route-table-management"
description: |-
  Provides details about managing a Drg's default DRG route table assignments in Oracle Cloud Infrastructure Core service
---

# oci_core_drg_default_route_table_management
This resource manages a DRG's per-attachment-type default DRG route table assignments — the same
operation exposed by the OCI console's DRG "Edit default assignments" screen, or
`oci network drg update --default-drg-route-tables`.

Each DRG auto-generates one default DRG route table per attachment type at DRG-creation time (VCN
gets its own; IPSEC_TUNNEL/VIRTUAL_CIRCUIT/REMOTE_PEERING_CONNECTION share one). This resource
reassigns which route table is used as an attachment type's default going forward — it does not
create a route table, and does not retroactively move any attachment already resolved to a
different route table.

This is intentionally a separate resource from `oci_core_drg` (which also exposes
`default_drg_route_tables` directly, for the case where the target route table's OCID is already
known and not created in the same configuration). Setting a network type's default to a DRG route
table created in the same `terraform apply` requires referencing that route table's resource
address; doing so via `oci_core_drg` itself would create a resource cycle, since
`oci_core_drg_route_table` must reference the DRG's `id` to be created in the first place. This
resource is addressed by `drg_id` instead, so the graph resolves as: create DRG -> create route
tables -> assign them as defaults, with no address referencing back to one that isn't finished
yet — same shape of problem `oci_core_drg_attachment_management` solves for
`drg_route_table_id`/`export_drg_route_distribution_id` on auto-created attachments.

Only one `oci_core_drg_default_route_table_management` resource should be declared per DRG — like
`oci_core_drg`'s own `default_drg_route_tables` argument, it manages all four attachment types'
assignments as a single `UpdateDrg` call.


## Example Usage

```hcl
resource "oci_core_drg_default_route_table_management" "test_drg_default_route_tables" {
  #Required
  drg_id = oci_core_drg.test_drg.id

  #Optional
  ipsec_tunnel              = oci_core_drg_route_table.test_drg_route_table.id
  remote_peering_connection = oci_core_drg_route_table.test_drg_route_table.id
  vcn                       = oci_core_drg_route_table.test_drg_route_table.id
  virtual_circuit           = oci_core_drg_route_table.test_drg_route_table.id
}
```

## Argument Reference

The following arguments are supported:

* `drg_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG whose default route table assignments this resource manages. Changing this creates a new resource.
* `ipsec_tunnel` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG route table to assign as the default for DRG attachments of type IPSEC_TUNNEL. Left unset, the current default for this type is unchanged.
* `remote_peering_connection` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG route table to assign as the default for DRG attachments of type REMOTE_PEERING_CONNECTION. Left unset, the current default for this type is unchanged.
* `vcn` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG route table to assign as the default for DRG attachments of type VCN. Left unset, the current default for this type is unchanged.
* `virtual_circuit` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG route table to assign as the default for DRG attachments of type VIRTUAL_CIRCUIT. Left unset, the current default for this type is unchanged.

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `drg_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG.
* `ipsec_tunnel` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG route table currently assigned as the default for DRG attachments of type IPSEC_TUNNEL.
* `remote_peering_connection` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG route table currently assigned as the default for DRG attachments of type REMOTE_PEERING_CONNECTION.
* `vcn` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG route table currently assigned as the default for DRG attachments of type VCN.
* `virtual_circuit` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG route table currently assigned as the default for DRG attachments of type VIRTUAL_CIRCUIT.
* `id` - Same value as `drg_id` — this resource has no identity of its own distinct from the DRG it manages.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the resource
	* `update` - (Defaults to 20 minutes), when updating the resource
	* `delete` - (Defaults to 20 minutes), when destroying the resource

Deleting this resource does not revert the DRG's default route table assignments — OCI has no
"unset" operation for a default, and this resource does not attempt to guess or restore whatever
the auto-generated default was before this resource started managing it. Terraform simply stops
tracking the assignment.

## Import

DrgDefaultRouteTableManagements can be imported using the `drg_id`, e.g.

```
$ terraform import oci_core_drg_default_route_table_management.test_drg_default_route_tables "drg_id"
```
