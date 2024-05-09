---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_managed_instances"
sidebar_current: "docs-oci-datasource-os_management_hub-managed_instances"
description: |-
  Provides the list of Managed Instances in Oracle Cloud Infrastructure Os Management Hub service
---

# Data Source: oci_os_management_hub_managed_instances
This data source provides the list of Managed Instances in Oracle Cloud Infrastructure Os Management Hub service.

Lists managed instances that match the specified compartment or managed instance OCID. Filter the list against a variety of criteria including but not limited to its name, status, architecture, and OS version.

## Example Usage

```hcl
data "oci_os_management_hub_managed_instances" "test_managed_instances" {

	#Optional
	advisory_name = var.managed_instance_advisory_name
	arch_type = var.managed_instance_arch_type
	compartment_id = var.compartment_id
	display_name = var.managed_instance_display_name
	display_name_contains = var.managed_instance_display_name_contains
	group = var.managed_instance_group
	group_not_equal_to = var.managed_instance_group_not_equal_to
	is_attached_to_group_or_lifecycle_stage = var.managed_instance_is_attached_to_group_or_lifecycle_stage
	is_managed_by_autonomous_linux = var.managed_instance_is_managed_by_autonomous_linux
	is_management_station = var.managed_instance_is_management_station
	is_profile_attached = var.managed_instance_is_profile_attached
	lifecycle_environment = var.managed_instance_lifecycle_environment
	lifecycle_environment_not_equal_to = var.managed_instance_lifecycle_environment_not_equal_to
	lifecycle_stage = var.managed_instance_lifecycle_stage
	lifecycle_stage_not_equal_to = var.managed_instance_lifecycle_stage_not_equal_to
	location = var.managed_instance_location
	location_not_equal_to = var.managed_instance_location_not_equal_to
	managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance.id
	os_family = var.managed_instance_os_family
	profile = var.managed_instance_profile
	profile_not_equal_to = var.managed_instance_profile_not_equal_to
	software_source_id = oci_os_management_hub_software_source.test_software_source.id
	status = var.managed_instance_status
}
```

## Argument Reference

The following arguments are supported:

* `advisory_name` - (Optional) The assigned erratum name. It's unique and not changeable.  Example: `ELSA-2020-5804` 
* `arch_type` - (Optional) A filter to return only instances whose architecture type matches the given architecture.
* `compartment_id` - (Optional) The OCID of the compartment that contains the resources to list. This filter returns only resources contained within the specified compartment.
* `display_name` - (Optional) A filter to return resources that match the given display names.
* `display_name_contains` - (Optional) A filter to return resources that may partially match the given display name.
* `group` - (Optional) A filter to return only managed instances that are attached to the specified group.
* `group_not_equal_to` - (Optional) A filter to return only managed instances that are NOT attached to the specified group.
* `is_attached_to_group_or_lifecycle_stage` - (Optional) A filter to return only managed instances that are attached to the specified group or lifecycle environment.
* `is_managed_by_autonomous_linux` - (Optional) Indicates whether to list only resources managed by the Autonomous Linux service. 
* `is_management_station` - (Optional) A filter to return only managed instances that are acting as management stations.
* `is_profile_attached` - (Optional) A filter to return only managed instances with a registration profile attached.
* `lifecycle_environment` - (Optional) A filter to return only managed instances in a specific lifecycle environment.
* `lifecycle_environment_not_equal_to` - (Optional) A filter to return only managed instances that aren't in a specific lifecycle environment.
* `lifecycle_stage` - (Optional) A filter to return only managed instances that are associated with the specified lifecycle environment.
* `lifecycle_stage_not_equal_to` - (Optional) A filter to return only managed instances that are NOT associated with the specified lifecycle environment.
* `location` - (Optional) A filter to return only resources whose location matches the given value.
* `location_not_equal_to` - (Optional) A filter to return only resources whose location does not match the given value.
* `managed_instance_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance. This filter returns resources associated with this managed instance.
* `os_family` - (Optional) A filter to return only resources that match the given operating system family.
* `profile` - (Optional) A multi filter to return only managed instances that match the given profile ids.
* `profile_not_equal_to` - (Optional) A multi filter to return only managed instances that don't contain the given profile [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `software_source_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source. This filter returns resources associated with this software source.
* `status` - (Optional) A filter to return only managed instances whose status matches the status provided.


## Attributes Reference

The following attributes are exported:

* `managed_instance_collection` - The list of managed_instance_collection.

### ManagedInstance Reference

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

