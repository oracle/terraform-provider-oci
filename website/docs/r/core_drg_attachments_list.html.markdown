---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_drg_attachments_list"
sidebar_current: "docs-oci-resource-core-drg_attachments_list"
description: |-
  Provides the Drg Attachments List resource in Oracle Cloud Infrastructure Core service
---

# oci_core_drg_attachments_list
This resource provides the Drg Attachments List resource in Oracle Cloud Infrastructure Core service.

Returns a complete list of DRG attachments that belong to a particular DRG.


## Example Usage

```hcl
resource "oci_core_drg_attachments_list" "test_drg_attachments_list" {
	#Required
	drg_id = oci_core_drg.test_drg.id

	#Optional
	attachment_type = var.drg_attachments_list_attachment_type
	is_cross_tenancy = var.drg_attachments_list_is_cross_tenancy
}
```

## Argument Reference

The following arguments are supported:

* `attachment_type` - (Optional) The type for the network resource attached to the DRG.
* `drg_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG.
* `is_cross_tenancy` - (Optional) Whether the DRG attachment lives in a different tenancy than the DRG.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `drg_all_attachments` - The list of drg_attachments.

### Drg_All_Attachments Reference

The following attributes are exported:

* `id` - The Oracle-assigned ID of the DRG attachment 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Drg Attachments List
	* `update` - (Defaults to 20 minutes), when updating the Drg Attachments List
	* `delete` - (Defaults to 20 minutes), when destroying the Drg Attachments List


## Import

Import is not supported for this resource.

