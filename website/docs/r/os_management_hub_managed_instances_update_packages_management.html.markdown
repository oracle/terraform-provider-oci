---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_managed_instances_update_packages_management"
sidebar_current: "docs-oci-resource-os_management_hub-managed_instances_update_packages_management"
description: |-
  Provides the Managed Instances Update Packages Management resource in Oracle Cloud Infrastructure Os Management Hub service
---

# oci_os_management_hub_managed_instances_update_packages_management
This resource provides the Managed Instances Update Packages Management resource in Oracle Cloud Infrastructure Os Management Hub service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/osmh/latest/ManagedInstance/UpdatePackages

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/os_management_hub

Install all of the available package updates for all of the managed instances in a compartment. This applies only to standalone non-Windows instances. This will not update instances that belong to a group or lifecycle environment.


## Example Usage

```hcl
resource "oci_os_management_hub_managed_instances_update_packages_management" "test_managed_instances_update_packages_management" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	update_types = var.managed_instances_update_packages_management_update_types
	work_request_details {

		#Optional
		description = var.managed_instances_update_packages_management_work_request_details_description
		display_name = var.managed_instances_update_packages_management_work_request_details_display_name
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `update_types` - (Optional) The types of updates to be applied.
* `work_request_details` - (Optional) Provides the name and description of the job.
	* `description` - (Optional) User-specified information about the job. Avoid entering confidential information.
	* `display_name` - (Optional) A user-friendly name for the job. The name does not have to be unique. Avoid entering confidential information.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Managed Instances Update Packages Management
	* `update` - (Defaults to 20 minutes), when updating the Managed Instances Update Packages Management
	* `delete` - (Defaults to 20 minutes), when destroying the Managed Instances Update Packages Management


## Import

ManagedInstancesUpdatePackagesManagement can be imported using the `id`, e.g.

```
$ terraform import oci_os_management_hub_managed_instances_update_packages_management.test_managed_instances_update_packages_management "id"
```

