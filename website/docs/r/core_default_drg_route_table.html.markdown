---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_default_drg_route_table"
sidebar_current: "docs-oci-resource-core-default_drg_route_table"
description: |-
  Provides the Default Drg Route Table resource in Oracle Cloud Infrastructure Core service
---

# oci_core_default_drg_route_table
This resource provides the Default Drg Route Table resource in Oracle Cloud Infrastructure Core service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/iaas/latest/DrgRouteTable

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/

Manages an Oracle-created default DRG route table. Use this resource when you want Terraform to adopt and manage
the default DRG route table that already exists for a DRG attachment type instead of creating a new DRG route table.

## Example Usage

```hcl
resource "oci_core_default_drg_route_table" "default_drg_route_table" {
	#Required
	manage_default_resource_id = oci_core_drg.test_drg.default_drg_route_tables[0].vcn

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = "managed-default-drg-route-table"
	freeform_tags = {"Department"= "Finance"}
	import_drg_route_distribution_id = oci_core_drg_route_distribution.test_drg_route_distribution.id
	is_ecmp_enabled = true
}
```

## Argument Reference

The following arguments are supported:

* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
* `import_drg_route_distribution_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the import route distribution used to specify how incoming route advertisements through referenced attachments are inserted into the DRG route table.
* `is_ecmp_enabled` - (Optional) (Updatable) If you want traffic to be routed using ECMP across your virtual circuits or IPSec tunnels to your on-premises networks, enable ECMP on the DRG route table.
* `manage_default_resource_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the default DRG route table to manage.
* `remove_import_trigger` - (Optional) (Updatable) An optional property when flipped disables the import of route Distribution by setting import_drg_route_distribution_id to null.

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment the DRG is in. The DRG route table is always in the same compartment as the DRG.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG route table.
* `import_drg_route_distribution_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the import route distribution used to specify how incoming route advertisements from referenced attachments are inserted into the DRG route table.
* `is_ecmp_enabled` - If you want traffic to be routed using ECMP across your virtual circuits or IPSec tunnels to your on-premises network, enable ECMP on the DRG route table to which these attachments import routes.
* `manage_default_resource_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the default DRG route table being managed.
* `state` - The DRG route table's current state.
* `time_created` - The date and time the DRG route table was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when adopting the Default Drg Route Table
	* `update` - (Defaults to 20 minutes), when updating the Default Drg Route Table
	* `delete` - (Defaults to 20 minutes), when removing the Default Drg Route Table from Terraform management

## Import

DefaultDrgRouteTables can be imported using the `id`, e.g.

```
$ terraform import oci_core_default_drg_route_table.default_vcn_route_table "id"
```
