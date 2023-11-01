---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_managed_instance_group"
sidebar_current: "docs-oci-resource-os_management_hub-managed_instance_group"
description: |-
  Provides the Managed Instance Group resource in Oracle Cloud Infrastructure Os Management Hub service
---

# oci_os_management_hub_managed_instance_group
This resource provides the Managed Instance Group resource in Oracle Cloud Infrastructure Os Management Hub service.

Creates a new managed instance group.


## Example Usage

```hcl
resource "oci_os_management_hub_managed_instance_group" "test_managed_instance_group" {
	#Required
	arch_type = var.managed_instance_group_arch_type
	compartment_id = var.compartment_id
	display_name = var.managed_instance_group_display_name
	os_family = var.managed_instance_group_os_family
	software_source_ids {
	}
	vendor_name = var.managed_instance_group_vendor_name

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.managed_instance_group_description
	freeform_tags = {"Department"= "Finance"}
	managed_instance_ids = var.managed_instance_group_managed_instance_ids
}
```

## Argument Reference

The following arguments are supported:

* `arch_type` - (Required) The CPU architecture type of the managed instance(s) that this managed instance group will contain. 
* `compartment_id` - (Required) The OCID of the tenancy containing the managed instance group.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) Details about the managed instance group.
* `display_name` - (Required) (Updatable) A user-friendly name for the managed instance group. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `managed_instance_ids` - (Optional) The list of managed instance OCIDs to be added to the managed instance group.
* `os_family` - (Required) The operating system type of the managed instance(s) that this managed instance group will contain. 
* `software_source_ids` - (Required) The list of software source OCIDs available to the managed instances in the managed instance group.
* `vendor_name` - (Required) The software source vendor name.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `arch_type` - The CPU architecture of the instances in the managed instance group.
* `compartment_id` - The OCID of the tenancy containing the managed instance group.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - Details describing the managed instance group.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The managed instance group OCID that is immutable on creation.
* `managed_instance_count` - The number of Managed Instances in the managed instance group.
* `managed_instance_ids` - The list of managed instances OCIDs attached to the managed instance group.
* `os_family` - The operating system type of the instances in the managed instance group.
* `pending_job_count` - The number of scheduled jobs pending against the managed instance group.
* `software_source_ids` - The list of software source OCIDs that the managed instance group will use.
* `software_sources` - The list of software sources that the managed instance group will use.
	* `description` - Software source description.
	* `display_name` - Software source name.
	* `id` - The OCID of the software source.
	* `software_source_type` - Type of the software source.
* `state` - The current state of the managed instance group.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the managed instance group was created. An RFC3339 formatted datetime string.
* `time_modified` - The time the managed instance group was last modified. An RFC3339 formatted datetime string.
* `vendor_name` - The software source vendor name.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Managed Instance Group
	* `update` - (Defaults to 20 minutes), when updating the Managed Instance Group
	* `delete` - (Defaults to 20 minutes), when destroying the Managed Instance Group


## Import

ManagedInstanceGroups can be imported using the `id`, e.g.

```
$ terraform import oci_os_management_hub_managed_instance_group.test_managed_instance_group "id"
```

