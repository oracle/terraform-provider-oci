---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_managed_instance_group_attach_software_sources_management"
sidebar_current: "docs-oci-resource-os_management_hub-managed_instance_group_attach_software_sources_management"
description: |-
  Provides the Managed Instance Group Attach Software Sources Management resource in Oracle Cloud Infrastructure Os Management Hub service
---

# oci_os_management_hub_managed_instance_group_attach_software_sources_management
This resource provides the Managed Instance Group Attach Software Sources Management resource in Oracle Cloud Infrastructure Os Management Hub service.

Attaches software sources to the specified managed instance group. The software sources must be compatible with the type of instances in the group.


## Example Usage

```hcl
resource "oci_os_management_hub_managed_instance_group_attach_software_sources_management" "test_managed_instance_group_attach_software_sources_management" {
	#Required
	managed_instance_group_id = oci_os_management_hub_managed_instance_group.test_managed_instance_group.id
	software_sources = var.managed_instance_group_attach_software_sources_management_software_sources

	#Optional
	work_request_details {

		#Optional
		description = var.managed_instance_group_attach_software_sources_management_work_request_details_description
		display_name = var.managed_instance_group_attach_software_sources_management_work_request_details_display_name
	}
}
```

## Argument Reference

The following arguments are supported:

* `managed_instance_group_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance group.
* `software_sources` - (Required) List of software source [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) to attach to the group.
* `work_request_details` - (Optional) Provides the name and description of the job.
	* `description` - (Optional) User-specified information about the job. Avoid entering confidential information.
	* `display_name` - (Optional) A user-friendly name for the job. The name does not have to be unique. Avoid entering confidential information.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Managed Instance Group Attach Software Sources Management
	* `update` - (Defaults to 20 minutes), when updating the Managed Instance Group Attach Software Sources Management
	* `delete` - (Defaults to 20 minutes), when destroying the Managed Instance Group Attach Software Sources Management


## Import

ManagedInstanceGroupAttachSoftwareSourcesManagement can be imported using the `id`, e.g.

```
$ terraform import oci_os_management_hub_managed_instance_group_attach_software_sources_management.test_managed_instance_group_attach_software_sources_management "id"
```

