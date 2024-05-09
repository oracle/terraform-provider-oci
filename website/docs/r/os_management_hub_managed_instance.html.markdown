---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_managed_instance"
sidebar_current: "docs-oci-resource-os_management_hub-managed_instance"
description: |-
  Provides the Managed Instance resource in Oracle Cloud Infrastructure Os Management Hub service
---

# oci_os_management_hub_managed_instance
This resource provides the Managed Instance resource in Oracle Cloud Infrastructure Os Management Hub service.

Updates the specified managed instance information, such as description, ONS topic, and associated management station.


## Example Usage

```hcl
resource "oci_os_management_hub_managed_instance" "test_managed_instance" {
	#Required
	managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance.id

	#Optional
	autonomous_settings {

		#Optional
		is_data_collection_authorized = var.managed_instance_autonomous_settings_is_data_collection_authorized
	}
	description = var.managed_instance_description
	notification_topic_id = oci_ons_notification_topic.test_notification_topic.id
	primary_management_station_id = oci_os_management_hub_management_station.test_management_station.id
	secondary_management_station_id = oci_os_management_hub_management_station.test_management_station.id
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_settings` - (Optional) (Updatable) Updatable settings for the Autonomous Linux service.
	* `is_data_collection_authorized` - (Optional) (Updatable) Indicates whether Autonomous Linux will collect crash files.
* `description` - (Optional) (Updatable) User-specified description of the managed instance. Avoid entering confidential information.
* `managed_instance_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance.
* `notification_topic_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the Oracle Notifications service (ONS) topic. ONS is the channel used to send notifications to the customer. 
* `primary_management_station_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management station for the instance to use as primary management station. 
* `secondary_management_station_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management station for the instance to use as secondary management station. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `architecture` - The CPU architecture type of the managed instance.
* `autonomous_settings` - Settings for the Autonomous Linux service.
	* `is_data_collection_authorized` - Indicates whether Autonomous Linux will collect crash files. This setting can be changed by the user.
	* `scheduled_job_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the restricted scheduled job associated with this instance. This value cannot be deleted by the user.
* `bug_updates_available` - Number of bug fix type updates available for installation.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the managed instance. 
* `description` - User-specified description for the managed instance.
* `display_name` - User-friendly name for the managed instance.
* `enhancement_updates_available` - Number of enhancement type updates available for installation.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance. 
* `installed_packages` - Number of packages installed on the instance.
* `installed_windows_updates` - Number of Windows updates installed on the instance.
* `is_managed_by_autonomous_linux` - Indicates whether the Autonomous Linux service manages the instance.
* `is_management_station` - Indicates whether this managed instance is acting as an on-premises management station.
* `is_reboot_required` - Indicates whether a reboot is required to complete installation of updates.
* `ksplice_effective_kernel_version` - The ksplice effective kernel version.
* `lifecycle_environment` - Id and name of a resource to simplify the display for the user.
	* `display_name` - User-friendly name.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource that is immutable on creation.
* `lifecycle_stage` - Id and name of a resource to simplify the display for the user.
	* `display_name` - User-friendly name.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource that is immutable on creation.
* `location` - The location of the managed instance.
* `managed_instance_group` - Id and name of a resource to simplify the display for the user.
	* `display_name` - User-friendly name.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource that is immutable on creation.
* `notification_topic_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the Oracle Notifications service (ONS) topic. ONS is the channel used to send notifications to the customer. 
* `os_family` - The operating system type of the managed instance.
* `os_kernel_version` - Operating system kernel version.
* `os_name` - Operating system name.
* `os_version` - Operating system version.
* `other_updates_available` - Number of non-classified (other) updates available for installation.
* `primary_management_station_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management station for the instance to use as primary management station. 
* `profile` - The profile that was used to register this instance with the service.
* `scheduled_job_count` - Number of scheduled jobs associated with this instance.
* `secondary_management_station_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management station for the instance to use as secondary managment station. 
* `security_updates_available` - Number of security type updates available for installation.
* `software_sources` - The list of software sources currently attached to the managed instance.
	* `description` - Software source description.
	* `display_name` - Software source name.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source.
	* `is_mandatory_for_autonomous_linux` - Indicates whether this is a required software source for Autonomous Linux instances. If true, the user can't unselect it.
	* `software_source_type` - Type of the software source.
* `status` - Current status of the managed instance.
* `tenancy_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the tenancy that the managed instance resides in. 
* `time_created` - The date and time the instance was created (in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) format). 
* `time_last_boot` - Time that the instance last booted (in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) format). 
* `time_last_checkin` - Time that the instance last checked in with the service (in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) format). 
* `time_updated` - The date and time the instance was last updated (in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) format). 
* `updates_available` - Number of updates available for installation.
* `work_request_count` - Number of work requests associated with this instance.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Managed Instance
	* `update` - (Defaults to 20 minutes), when updating the Managed Instance
	* `delete` - (Defaults to 20 minutes), when destroying the Managed Instance


## Import

ManagedInstances can be imported using the `id`, e.g.

```
$ terraform import oci_os_management_hub_managed_instance.test_managed_instance "id"
```

