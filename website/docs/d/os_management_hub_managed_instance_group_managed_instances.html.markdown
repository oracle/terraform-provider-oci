---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_managed_instance_group_managed_instances"
sidebar_current: "docs-oci-datasource-os_management_hub-managed_instance_group_managed_instances"
description: |-
  Provides the list of Managed Instance Group Managed Instances in Oracle Cloud Infrastructure Os Management Hub service
---

# Data Source: oci_os_management_hub_managed_instance_group_managed_instances
This data source provides the list of Managed Instance Group Managed Instances in Oracle Cloud Infrastructure Os Management Hub service.

Lists all managed instances for a specified managed instance group. Filter the list against a variety of criteria including but not limited to the managed instance name. The results list all managed instances that have already been added to the group.


## Example Usage

```hcl
data "oci_os_management_hub_managed_instance_group_managed_instances" "test_managed_instance_group_managed_instances" {
	#Required
	managed_instance_group_id = oci_os_management_hub_managed_instance_group.test_managed_instance_group.id

	#Optional
	compartment_id = var.compartment_id
	compartment_id_in_subtree = var.managed_instance_group_managed_instance_compartment_id_in_subtree
	display_name = var.managed_instance_group_managed_instance_display_name
	display_name_contains = var.managed_instance_group_managed_instance_display_name_contains
	managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The OCID of the compartment that contains the resources to list. This filter returns only resources contained within the specified compartment.
* `compartment_id_in_subtree` - (Optional) Indicates whether to include subcompartments in the returned results. Default is false.
* `display_name` - (Optional) A filter to return resources that match the given display names.
* `display_name_contains` - (Optional) A filter to return resources that may partially match the given display name.
* `managed_instance_group_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance group.
* `managed_instance_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance. This filter returns resources associated with this managed instance.


## Attributes Reference

The following attributes are exported:

* `managed_instance_collection` - The list of managed_instance_collection.

### ManagedInstanceGroupManagedInstance Reference

The following attributes are exported:

* `items` - List of managed instances.
	* `agent_version` - The version of osmh-agent running on the managed instance
	* `architecture` - The CPU architecture type of the managed instance.
	* `autonomous_settings` - Settings for the Autonomous Linux service.
		* `is_data_collection_authorized` - Indicates whether Autonomous Linux will collect crash files. This setting can be changed by the user.
		* `scheduled_job_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the restricted scheduled job associated with this instance. This value cannot be deleted by the user.
	* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the managed instance. 
	* `description` - User-specified description of the managed instance.
	* `display_name` - User-friendly name for the managed instance.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance. 
	* `is_managed_by_autonomous_linux` - Indicates whether Autonomous Linux manages this instance.
	* `is_management_station` - Whether this managed instance is acting as an on-premises management station.
	* `is_reboot_required` - Indicates whether a reboot is required to complete installation of updates.
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
	* `status` - Current status of the managed instance.
	* `tenancy_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the tenancy this managed instance resides in. 
	* `time_last_boot` - Time that the instance last booted (in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) format). 
	* `updates_available` - Number of updates available for installation.

