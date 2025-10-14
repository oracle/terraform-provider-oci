---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_lifecycle_stage_reboot_management"
sidebar_current: "docs-oci-resource-os_management_hub-lifecycle_stage_reboot_management"
description: |-
  Provides the Lifecycle Stage Reboot Management resource in Oracle Cloud Infrastructure Os Management Hub service
---

# oci_os_management_hub_lifecycle_stage_reboot_management
This resource provides the Lifecycle Stage Reboot Management resource in Oracle Cloud Infrastructure Os Management Hub service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/os-management/latest/LifecycleStageRebootManagement

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/os_management_hub

Reboots all managed instances in the specified lifecycle stage.


## Example Usage

```hcl
resource "oci_os_management_hub_lifecycle_stage_reboot_management" "test_lifecycle_stage_reboot_management" {
	#Required
	lifecycle_stage_id = oci_os_management_hub_lifecycle_stage.test_lifecycle_stage.id

	#Optional
	reboot_timeout_in_mins = var.lifecycle_stage_reboot_management_reboot_timeout_in_mins
	work_request_details {

		#Optional
		description = var.lifecycle_stage_reboot_management_work_request_details_description
		display_name = var.lifecycle_stage_reboot_management_work_request_details_display_name
	}
}
```

## Argument Reference

The following arguments are supported:

* `lifecycle_stage_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the lifecycle stage.
* `reboot_timeout_in_mins` - (Optional) The number of minutes the service waits for the reboot to complete. If the instances in the stage don't reboot  within this time, the reboot job status is set to failed. 
* `work_request_details` - (Optional) Provides the name and description of the job.
	* `description` - (Optional) User-specified information about the job. Avoid entering confidential information.
	* `display_name` - (Optional) A user-friendly name for the job. The name does not have to be unique. Avoid entering confidential information.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Lifecycle Stage Reboot Management
	* `update` - (Defaults to 20 minutes), when updating the Lifecycle Stage Reboot Management
	* `delete` - (Defaults to 20 minutes), when destroying the Lifecycle Stage Reboot Management


## Import

LifecycleStageRebootManagement can be imported using the `id`, e.g.

```
$ terraform import oci_os_management_hub_lifecycle_stage_reboot_management.test_lifecycle_stage_reboot_management "id"
```

