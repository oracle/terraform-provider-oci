---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_work_request_rerun_management"
sidebar_current: "docs-oci-resource-os_management_hub-work_request_rerun_management"
description: |-
  Provides the Work Request Rerun Management resource in Oracle Cloud Infrastructure Os Management Hub service
---

# oci_os_management_hub_work_request_rerun_management
This resource provides the Work Request Rerun Management resource in Oracle Cloud Infrastructure Os Management Hub service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/os-management/latest/WorkRequestRerunManagement

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/os_management_hub

Reruns a failed work for the specified work request [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Rerunning restarts the work on failed targets.


## Example Usage

```hcl
resource "oci_os_management_hub_work_request_rerun_management" "test_work_request_rerun_management" {
	#Required
	work_request_id = oci_containerengine_work_request.test_work_request.id

	#Optional
	managed_instances = var.work_request_rerun_management_managed_instances
	work_request_details {

		#Optional
		description = var.work_request_rerun_management_work_request_details_description
		display_name = var.work_request_rerun_management_work_request_details_display_name
	}
}
```

## Argument Reference

The following arguments are supported:

* `managed_instances` - (Optional) List of managed instance [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) to affected by the rerun of the work request.
* `work_request_details` - (Optional) Provides the name and description of the job.
	* `description` - (Optional) User-specified information about the job. Avoid entering confidential information.
	* `display_name` - (Optional) A user-friendly name for the job. The name does not have to be unique. Avoid entering confidential information.
* `work_request_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the work request.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Work Request Rerun Management
	* `update` - (Defaults to 20 minutes), when updating the Work Request Rerun Management
	* `delete` - (Defaults to 20 minutes), when destroying the Work Request Rerun Management


## Import

WorkRequestRerunManagement can be imported using the `id`, e.g.

```
$ terraform import oci_os_management_hub_work_request_rerun_management.test_work_request_rerun_management "id"
```

