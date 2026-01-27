---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_managed_instance_install_snaps_management"
sidebar_current: "docs-oci-resource-os_management_hub-managed_instance_install_snaps_management"
description: |-
  Provides the Managed Instance Install Snaps Management resource in Oracle Cloud Infrastructure Os Management Hub service
---

# oci_os_management_hub_managed_instance_install_snaps_management
This resource provides the Managed Instance Install Snaps Management resource in Oracle Cloud Infrastructure Os Management Hub service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/osmh/latest/ManagedInstance/InstallSnaps

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/os_management_hub

Installs specified snaps on a managed instance.

## Example Usage

```hcl
resource "oci_os_management_hub_managed_instance_install_snaps_management" "test_managed_instance_install_snaps_management" {
	#Required
	managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance.id

	#Optional
	snap_details {
		#Required
		name = var.managed_instance_install_snaps_management_snap_details_name

		#Optional
		channel = var.managed_instance_install_snaps_management_snap_details_channel
		is_signed = var.managed_instance_install_snaps_management_snap_details_is_signed
		mode = var.managed_instance_install_snaps_management_snap_details_mode
		revision = var.managed_instance_install_snaps_management_snap_details_revision
	}
	work_request_details {

		#Optional
		description = var.managed_instance_install_snaps_management_work_request_details_description
		display_name = var.managed_instance_install_snaps_management_work_request_details_display_name
	}
}
```

## Argument Reference

The following arguments are supported:

* `managed_instance_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance.
* `snap_details` - (Optional) The array of snaps to install.
	* `channel` - (Optional) The snap channel to install from (e.g. stable, 1.2/edge, beta, candidate, or a custom channel). 
	* `is_signed` - (Optional) If false, allows installing snaps not signed by the Snap Store. E.g., snaps from local file. Use with caution. 
	* `mode` - (Optional) The confinement mode for the snap. 
	* `name` - (Required) The name of the snap to install. 
	* `revision` - (Optional) The snap revision to install. 
* `work_request_details` - (Optional) Provides the name and description of the job.
	* `description` - (Optional) User-specified information about the job. Avoid entering confidential information.
	* `display_name` - (Optional) A user-friendly name for the job. The name does not have to be unique. Avoid entering confidential information.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Managed Instance Install Snaps Management
	* `update` - (Defaults to 20 minutes), when updating the Managed Instance Install Snaps Management
	* `delete` - (Defaults to 20 minutes), when destroying the Managed Instance Install Snaps Management


## Import

ManagedInstanceInstallSnapsManagement can be imported using the `id`, e.g.

```
$ terraform import oci_os_management_hub_managed_instance_install_snaps_management.test_managed_instance_install_snaps_management "id"
```

