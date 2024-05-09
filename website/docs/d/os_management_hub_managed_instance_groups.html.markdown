---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_managed_instance_groups"
sidebar_current: "docs-oci-datasource-os_management_hub-managed_instance_groups"
description: |-
  Provides the list of Managed Instance Groups in Oracle Cloud Infrastructure Os Management Hub service
---

# Data Source: oci_os_management_hub_managed_instance_groups
This data source provides the list of Managed Instance Groups in Oracle Cloud Infrastructure Os Management Hub service.

Lists managed instance groups that match the specified compartment or managed instance group [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Filter the list against a variety of criteria including but not limited to name, status, architecture, and OS family.


## Example Usage

```hcl
data "oci_os_management_hub_managed_instance_groups" "test_managed_instance_groups" {

	#Optional
	arch_type = var.managed_instance_group_arch_type
	compartment_id = var.compartment_id
	display_name = var.managed_instance_group_display_name
	display_name_contains = var.managed_instance_group_display_name_contains
	is_managed_by_autonomous_linux = var.managed_instance_group_is_managed_by_autonomous_linux
	location = var.managed_instance_group_location
	location_not_equal_to = var.managed_instance_group_location_not_equal_to
	managed_instance_group_id = oci_os_management_hub_managed_instance_group.test_managed_instance_group.id
	os_family = var.managed_instance_group_os_family
	software_source_id = oci_os_management_hub_software_source.test_software_source.id
	state = var.managed_instance_group_state
}
```

## Argument Reference

The following arguments are supported:

* `arch_type` - (Optional) A filter to return only profiles that match the given archType.
* `compartment_id` - (Optional) (Updatable) The OCID of the compartment that contains the resources to list. This filter returns only resources contained within the specified compartment.
* `display_name` - (Optional) A filter to return resources that match the given display names.
* `display_name_contains` - (Optional) A filter to return resources that may partially match the given display name.
* `is_managed_by_autonomous_linux` - (Optional) Indicates whether to list only resources managed by the Autonomous Linux service. 
* `location` - (Optional) A filter to return only resources whose location matches the given value.
* `location_not_equal_to` - (Optional) A filter to return only resources whose location does not match the given value.
* `managed_instance_group_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance group. This filter returns resources associated with this group.
* `os_family` - (Optional) A filter to return only resources that match the given operating system family.
* `software_source_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source. This filter returns resources associated with this software source.
* `state` - (Optional) A filter to return only managed instance groups that are in the specified state.


## Attributes Reference

The following attributes are exported:

* `managed_instance_group_collection` - The list of managed_instance_group_collection.

### ManagedInstanceGroup Reference

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
* `software_source_ids` - The list of software source [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) that the managed instance group will use.
	* `description` - Software source description.
	* `display_name` - Software source name.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source.
	* `is_mandatory_for_autonomous_linux` - Indicates whether this is a required software source for Autonomous Linux instances. If true, the user can't unselect it.
	* `software_source_type` - Type of the software source.
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

