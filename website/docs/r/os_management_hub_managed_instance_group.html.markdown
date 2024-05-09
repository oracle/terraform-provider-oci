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
	vendor_name = var.managed_instance_group_vendor_name

	#Optional
	autonomous_settings {

		#Optional
		is_data_collection_authorized = var.managed_instance_group_autonomous_settings_is_data_collection_authorized
	}
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.managed_instance_group_description
	freeform_tags = {"Department"= "Finance"}
	location = var.managed_instance_group_location
	managed_instance_ids = var.managed_instance_group_managed_instance_ids
	notification_topic_id = oci_ons_notification_topic.test_notification_topic.id
	software_source_ids {
	}
}
```

## Argument Reference

The following arguments are supported:

* `arch_type` - (Required) The CPU architecture type of the managed instances that will be attached to this group. 
* `autonomous_settings` - (Optional) (Updatable) Updatable settings for the Autonomous Linux service.
	* `is_data_collection_authorized` - (Optional) (Updatable) Indicates whether Autonomous Linux will collect crash files.
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the managed instance group.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) User-specified description of the managed instance group. Avoid entering confidential information.
* `display_name` - (Required) (Updatable) A user-friendly name for the managed instance group. Does not have to be unique and you can change the name later. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `location` - (Optional) The location of managed instances attached to the group. If no location is provided, the default is on premises. 
* `managed_instance_ids` - (Optional) The list of managed instance [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) to be added to the group.
* `notification_topic_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the Oracle Notifications service (ONS) topic. ONS is the channel used to send notifications to the customer.
* `os_family` - (Required) The operating system type of the managed instances that will be attached to this group. 
* `software_source_ids` - (Optional) The list of software source [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) available to the managed instances in the group.
* `vendor_name` - (Required) The vendor of the operating system that will be used by the managed instances in the group. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `arch_type` - The CPU architecture of the instances in the managed instance group.
* `autonomous_settings` - Settings for the Autonomous Linux service.
	* `is_data_collection_authorized` - Indicates whether Autonomous Linux will collect crash files. This setting can be changed by the user.
	* `scheduled_job_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the restricted scheduled job associated with this instance. This value cannot be deleted by the user.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the managed instance group.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - User-specified information about the managed instance group.
* `display_name` - A user-friendly name for the managed instance group.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance group.
* `is_managed_by_autonomous_linux` - Indicates whether the Autonomous Linux service manages the group.
* `location` - The location of managed instances attached to the group.
* `managed_instance_count` - The number of managed instances in the group.
* `managed_instance_ids` - The list of managed instance [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) attached to the managed instance group.
* `notification_topic_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the Oracle Notifications service (ONS) topic. ONS is the channel used to send notifications to the customer.
* `os_family` - The operating system type of the instances in the managed instance group.
* `pending_job_count` - The number of scheduled jobs pending against the managed instance group.
* `software_source_ids` - The list of software source OCIDs that the managed instance group will use.
* `software_sources` - The list of software sources that the managed instance group will use.
	* `description` - Software source description.
	* `display_name` - Software source name.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source.
	* `is_mandatory_for_autonomous_linux` - Indicates whether this is a required software source for Autonomous Linux instances. If true, the user can't unselect it.
	* `software_source_type` - Type of the software source.
* `state` - The current state of the managed instance group.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the managed instance group was created (in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) format).
* `time_modified` - The time the managed instance group was last modified (in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) format).
* `vendor_name` - The vendor of the operating system used by the managed instances in the group.

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

