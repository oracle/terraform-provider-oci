---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_managed_instance_group_remove_packages_management"
sidebar_current: "docs-oci-resource-os_management_hub-managed_instance_group_remove_packages_management"
description: |-
  Provides the Managed Instance Group Remove Packages Management resource in Oracle Cloud Infrastructure Os Management Hub service
---

# oci_os_management_hub_managed_instance_group_remove_packages_management
This resource provides the Managed Instance Group Remove Packages Management resource in Oracle Cloud Infrastructure Os Management Hub service.

Removes the specified packages from each managed instance in a managed instance group.


## Example Usage

```hcl
resource "oci_os_management_hub_managed_instance_group_remove_packages_management" "test_managed_instance_group_remove_packages_management" {
	#Required
	managed_instance_group_id = oci_os_management_hub_managed_instance_group.test_managed_instance_group.id
	package_names = var.managed_instance_group_remove_packages_management_package_names

	#Optional
	work_request_details {

		#Optional
		description = var.managed_instance_group_remove_packages_management_work_request_details_description
		display_name = var.managed_instance_group_remove_packages_management_work_request_details_display_name
	}
}
```

## Argument Reference

The following arguments are supported:

* `managed_instance_group_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance group.
* `package_names` - (Required) The list of package names.
* `work_request_details` - (Optional) Provides the name and description of the job.
	* `description` - (Optional) User-specified information about the job. Avoid entering confidential information.
	* `display_name` - (Optional) A user-friendly name for the job. The name does not have to be unique. Avoid entering confidential information.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Managed Instance Group Remove Packages Management
	* `update` - (Defaults to 20 minutes), when updating the Managed Instance Group Remove Packages Management
	* `delete` - (Defaults to 20 minutes), when destroying the Managed Instance Group Remove Packages Management


## Import

ManagedInstanceGroupRemovePackagesManagement can be imported using the `id`, e.g.

```
$ terraform import oci_os_management_hub_managed_instance_group_remove_packages_management.test_managed_instance_group_remove_packages_management "id"
```

