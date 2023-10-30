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

Lists managed instance groups that match the specified compartment or managed instance group OCID. Filter the 
list against a variety of criteria including but not limited to its name, status, architecture, and OS family.


## Example Usage

```hcl
data "oci_os_management_hub_managed_instance_groups" "test_managed_instance_groups" {

	#Optional
	arch_type = var.managed_instance_group_arch_type
	compartment_id = var.compartment_id
	display_name = var.managed_instance_group_display_name
	display_name_contains = var.managed_instance_group_display_name_contains
	managed_instance_group_id = oci_os_management_hub_managed_instance_group.test_managed_instance_group.id
	os_family = var.managed_instance_group_os_family
	software_source_id = oci_os_management_hub_software_source.test_software_source.id
	state = var.managed_instance_group_state
}
```

## Argument Reference

The following arguments are supported:

* `arch_type` - (Optional) A filter to return only profiles that match the given archType.
* `compartment_id` - (Optional) The OCID of the compartment that contains the resources to list.
* `display_name` - (Optional) A filter to return resources that match the given display names.
* `display_name_contains` - (Optional) A filter to return resources that may partially match the given display name.
* `managed_instance_group_id` - (Optional) The OCID of the managed instance group for which to list resources.
* `os_family` - (Optional) A filter to return only profiles that match the given osFamily.
* `software_source_id` - (Optional) The OCID for the software source.
* `state` - (Optional) A filter to return only resources their lifecycle state matches the given lifecycle state.


## Attributes Reference

The following attributes are exported:

* `managed_instance_group_collection` - The list of managed_instance_group_collection.

### ManagedInstanceGroup Reference

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
* `software_source_ids` - The list of software sources that the managed instance group will use.
	* `description` - Software source description.
	* `display_name` - Software source name.
	* `id` - The OCID of the software source.
	* `software_source_type` - Type of the software source.
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

