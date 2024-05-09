---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_lifecycle_stage_promote_software_source_management"
sidebar_current: "docs-oci-resource-os_management_hub-lifecycle_stage_promote_software_source_management"
description: |-
  Provides the Lifecycle Stage Promote Software Source Management resource in Oracle Cloud Infrastructure Os Management Hub service
---

# oci_os_management_hub_lifecycle_stage_promote_software_source_management
This resource provides the Lifecycle Stage Promote Software Source Management resource in Oracle Cloud Infrastructure Os Management Hub service.

Updates the versioned custom software source content to the specified lifecycle stage.
A versioned custom software source OCID (softwareSourceId) is required when promoting content to the first lifecycle stage. You must promote content to the first stage before promoting to subsequent stages, otherwise the service returns an error.
The softwareSourceId is optional when promoting content to the second, third, forth, or fifth stages. If you provide a softwareSourceId, the service validates that it matches the softwareSourceId of the previous stage. If it does not match, the service returns an error. If you don't provide a softwareSourceId, the service promotes the versioned software source from the previous lifecycle stage. If the previous lifecycle stage has no software source, the service returns an error.


## Example Usage

```hcl
resource "oci_os_management_hub_lifecycle_stage_promote_software_source_management" "test_lifecycle_stage_promote_software_source_management" {
	#Required
	lifecycle_stage_id = oci_os_management_hub_lifecycle_stage.test_lifecycle_stage.id

	#Optional
	software_source_id = oci_os_management_hub_software_source.test_software_source.id
	work_request_details {

		#Optional
		description = var.lifecycle_stage_promote_software_source_management_work_request_details_description
		display_name = var.lifecycle_stage_promote_software_source_management_work_request_details_display_name
	}
}
```

## Argument Reference

The following arguments are supported:

* `lifecycle_stage_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the lifecycle stage.
* `software_source_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source. This filter returns resources associated with this software source.
* `work_request_details` - (Optional) Provides the name and description of the job.
	* `description` - (Optional) User-specified information about the job. Avoid entering confidential information.
	* `display_name` - (Optional) A user-friendly name for the job. The name does not have to be unique. Avoid entering confidential information.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Lifecycle Stage Promote Software Source Management
	* `update` - (Defaults to 20 minutes), when updating the Lifecycle Stage Promote Software Source Management
	* `delete` - (Defaults to 20 minutes), when destroying the Lifecycle Stage Promote Software Source Management


## Import

LifecycleStagePromoteSoftwareSourceManagement can be imported using the `id`, e.g.

```
$ terraform import oci_os_management_hub_lifecycle_stage_promote_software_source_management.test_lifecycle_stage_promote_software_source_management "id"
```

