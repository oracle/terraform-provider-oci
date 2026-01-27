---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_managed_instance_detach_software_sources_management"
sidebar_current: "docs-oci-resource-os_management_hub-managed_instance_detach_software_sources_management"
description: |-
  Provides the Managed Instance Detach Software Sources Management resource in Oracle Cloud Infrastructure Os Management Hub service
---

# oci_os_management_hub_managed_instance_detach_software_sources_management
This resource provides the Managed Instance Detach Software Sources Management resource in Oracle Cloud Infrastructure Os Management Hub service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/osmh/latest/ManagedInstance/DetachSoftwareSources

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/os_management_hub

Removes software sources from a managed instance.
Packages will no longer be able to be installed from these software sources.


## Example Usage

```hcl
resource "oci_os_management_hub_managed_instance_detach_software_sources_management" "test_managed_instance_detach_software_sources_management" {
	#Required
	managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance.id
	software_sources = var.managed_instance_detach_software_sources_management_software_sources

	#Optional
	work_request_details {

		#Optional
		description = var.managed_instance_detach_software_sources_management_work_request_details_description
		display_name = var.managed_instance_detach_software_sources_management_work_request_details_display_name
	}
}
```

## Argument Reference

The following arguments are supported:

* `managed_instance_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance.
* `software_sources` - (Required) The list of software source OCIDs to be attached/detached.
* `work_request_details` - (Optional) Provides the name and description of the job.
	* `description` - (Optional) User-specified information about the job. Avoid entering confidential information.
	* `display_name` - (Optional) A user-friendly name for the job. The name does not have to be unique. Avoid entering confidential information.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Managed Instance Detach Software Sources Management
	* `update` - (Defaults to 20 minutes), when updating the Managed Instance Detach Software Sources Management
	* `delete` - (Defaults to 20 minutes), when destroying the Managed Instance Detach Software Sources Management


## Import

ManagedInstanceDetachSoftwareSourcesManagement can be imported using the `id`, e.g.

```
$ terraform import oci_os_management_hub_managed_instance_detach_software_sources_management.test_managed_instance_detach_software_sources_management "id"
```

