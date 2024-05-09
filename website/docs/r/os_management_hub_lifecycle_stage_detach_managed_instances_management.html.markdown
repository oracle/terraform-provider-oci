---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_lifecycle_stage_detach_managed_instances_management"
sidebar_current: "docs-oci-resource-os_management_hub-lifecycle_stage_detach_managed_instances_management"
description: |-
  Provides the Lifecycle Stage Detach Managed Instances Management resource in Oracle Cloud Infrastructure Os Management Hub service
---

# oci_os_management_hub_lifecycle_stage_detach_managed_instances_management
This resource provides the Lifecycle Stage Detach Managed Instances Management resource in Oracle Cloud Infrastructure Os Management Hub service.

Detaches (removes) a managed instance from a lifecycle stage.


## Example Usage

```hcl
resource "oci_os_management_hub_lifecycle_stage_detach_managed_instances_management" "test_lifecycle_stage_detach_managed_instances_management" {
	#Required
	lifecycle_stage_id = oci_os_management_hub_lifecycle_stage.test_lifecycle_stage.id

	#Required
	managed_instance_details {
		#Required
		managed_instances = var.lifecycle_stage_detach_managed_instances_management_managed_instance_details_managed_instances

		#Optional
		work_request_details {

			#Optional
			description = var.lifecycle_stage_detach_managed_instances_management_managed_instance_details_work_request_details_description
			display_name = var.lifecycle_stage_detach_managed_instances_management_managed_instance_details_work_request_details_display_name
		}
	}
}
```

## Argument Reference

The following arguments are supported:

* `lifecycle_stage_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the lifecycle stage.
* `managed_instance_details` - (Required) The details about the managed instances.
	* `managed_instances` - (Required) The list of managed instance OCIDs to be attached/detached.
	* `work_request_details` - (Optional) Provides the name and description of the job.
		* `description` - (Optional) User-specified information about the job. Avoid entering confidential information.
		* `display_name` - (Optional) A user-friendly name for the job. The name does not have to be unique. Avoid entering confidential information.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Lifecycle Stage Detach Managed Instances Management
	* `update` - (Defaults to 20 minutes), when updating the Lifecycle Stage Detach Managed Instances Management
	* `delete` - (Defaults to 20 minutes), when destroying the Lifecycle Stage Detach Managed Instances Management


## Import

LifecycleStageDetachManagedInstancesManagement can be imported using the `id`, e.g.

```
$ terraform import oci_os_management_hub_lifecycle_stage_detach_managed_instances_management.test_lifecycle_stage_detach_managed_instances_management "id"
```

