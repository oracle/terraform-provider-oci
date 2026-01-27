---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_dynamic_set_remove_packages_management"
sidebar_current: "docs-oci-resource-os_management_hub-dynamic_set_remove_packages_management"
description: |-
  Provides the Dynamic Set Remove Packages Management resource in Oracle Cloud Infrastructure Os Management Hub service
---

# oci_os_management_hub_dynamic_set_remove_packages_management
This resource provides the Dynamic Set Remove Packages Management resource in Oracle Cloud Infrastructure Os Management Hub service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/osmh/latest/DynamicSet/RemovePackages

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/os_management_hub

Removes specified software packages from all managed instances in the dynamic set.

## Example Usage

```hcl
resource "oci_os_management_hub_dynamic_set_remove_packages_management" "test_dynamic_set_remove_packages_management" {
	#Required
	dynamic_set_id = oci_os_management_hub_dynamic_set.test_dynamic_set.id
	package_names = var.dynamic_set_remove_packages_management_package_names

	#Optional
	managed_instances = var.dynamic_set_remove_packages_management_managed_instances
	work_request_details {

		#Optional
		description = var.dynamic_set_remove_packages_management_work_request_details_description
		display_name = var.dynamic_set_remove_packages_management_work_request_details_display_name
	}
}
```

## Argument Reference

The following arguments are supported:

* `dynamic_set_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the dynamic set. This filter returns resources associated with this dynamic set.
* `managed_instances` - (Optional) The list of managed instance OCIDs to be attached/detached.
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
	* `create` - (Defaults to 20 minutes), when creating the Dynamic Set Remove Packages Management
	* `update` - (Defaults to 20 minutes), when updating the Dynamic Set Remove Packages Management
	* `delete` - (Defaults to 20 minutes), when destroying the Dynamic Set Remove Packages Management


## Import

DynamicSetRemovePackagesManagement can be imported using the `id`, e.g.

```
$ terraform import oci_os_management_hub_dynamic_set_remove_packages_management.test_dynamic_set_remove_packages_management "id"
```

