---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_managed_instance_group_attach_managed_instances_management"
sidebar_current: "docs-oci-resource-os_management_hub-managed_instance_group_attach_managed_instances_management"
description: |-
  Provides the Managed Instance Group Attach Managed Instances Management resource in Oracle Cloud Infrastructure Os Management Hub service
---

# oci_os_management_hub_managed_instance_group_attach_managed_instances_management
This resource provides the Managed Instance Group Attach Managed Instances Management resource in Oracle Cloud Infrastructure Os Management Hub service.

Adds managed instances to the specified managed instance group. After adding instances to the group, any operation applied to the group will be applied to all instances in the group.


## Example Usage

```hcl
resource "oci_os_management_hub_managed_instance_group_attach_managed_instances_management" "test_managed_instance_group_attach_managed_instances_management" {
	#Required
	managed_instance_group_id = oci_os_management_hub_managed_instance_group.test_managed_instance_group.id
	managed_instances = var.managed_instance_group_attach_managed_instances_management_managed_instances

	#Optional
	work_request_details {

		#Optional
		description = var.managed_instance_group_attach_managed_instances_management_work_request_details_description
		display_name = var.managed_instance_group_attach_managed_instances_management_work_request_details_display_name
	}
}
```

## Argument Reference

The following arguments are supported:

* `managed_instance_group_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance group.
* `managed_instances` - (Required) List of managed instance [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) to attach to the group.
* `work_request_details` - (Optional) Provides the name and description of the job.
	* `description` - (Optional) User-specified information about the job. Avoid entering confidential information.
	* `display_name` - (Optional) A user-friendly name for the job. The name does not have to be unique. Avoid entering confidential information.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Managed Instance Group Attach Managed Instances Management
	* `update` - (Defaults to 20 minutes), when updating the Managed Instance Group Attach Managed Instances Management
	* `delete` - (Defaults to 20 minutes), when destroying the Managed Instance Group Attach Managed Instances Management


## Import

ManagedInstanceGroupAttachManagedInstancesManagement can be imported using the `id`, e.g.

```
$ terraform import oci_os_management_hub_managed_instance_group_attach_managed_instances_management.test_managed_instance_group_attach_managed_instances_management "id"
```

