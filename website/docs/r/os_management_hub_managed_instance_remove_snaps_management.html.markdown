---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_managed_instance_remove_snaps_management"
sidebar_current: "docs-oci-resource-os_management_hub-managed_instance_remove_snaps_management"
description: |-
  Provides the Managed Instance Remove Snaps Management resource in Oracle Cloud Infrastructure Os Management Hub service
---

# oci_os_management_hub_managed_instance_remove_snaps_management
This resource provides the Managed Instance Remove Snaps Management resource in Oracle Cloud Infrastructure Os Management Hub service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/osmh/latest/ManagedInstance/RemoveSnaps

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/os_management_hub

Removes specified snaps from a managed instance.

## Example Usage

```hcl
resource "oci_os_management_hub_managed_instance_remove_snaps_management" "test_managed_instance_remove_snaps_management" {
	#Required
	managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance.id
	snap_details {
		#Required
		name = var.managed_instance_remove_snaps_management_snap_details_name

		#Optional
		revision = var.managed_instance_remove_snaps_management_snap_details_revision
	}

	#Optional
	work_request_details {

		#Optional
		description = var.managed_instance_remove_snaps_management_work_request_details_description
		display_name = var.managed_instance_remove_snaps_management_work_request_details_display_name
	}
}
```

## Argument Reference

The following arguments are supported:

* `managed_instance_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance.
* `snap_details` - (Required) The array of snaps to remove.
	* `name` - (Required) The name of the snap to remove. 
	* `revision` - (Optional) The revision to remove. Note: This option cannot be used when removing multiple snaps. 
* `work_request_details` - (Optional) Provides the name and description of the job.
	* `description` - (Optional) User-specified information about the job. Avoid entering confidential information.
	* `display_name` - (Optional) A user-friendly name for the job. The name does not have to be unique. Avoid entering confidential information.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Managed Instance Remove Snaps Management
	* `update` - (Defaults to 20 minutes), when updating the Managed Instance Remove Snaps Management
	* `delete` - (Defaults to 20 minutes), when destroying the Managed Instance Remove Snaps Management


## Import

ManagedInstanceRemoveSnapsManagement can be imported using the `id`, e.g.

```
$ terraform import oci_os_management_hub_managed_instance_remove_snaps_management.test_managed_instance_remove_snaps_management "id"
```

