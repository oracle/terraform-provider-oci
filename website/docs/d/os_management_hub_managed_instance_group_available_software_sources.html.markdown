---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_managed_instance_group_available_software_sources"
sidebar_current: "docs-oci-datasource-os_management_hub-managed_instance_group_available_software_sources"
description: |-
  Provides the list of Managed Instance Group Available Software Sources in Oracle Cloud Infrastructure Os Management Hub service
---

# Data Source: oci_os_management_hub_managed_instance_group_available_software_sources
This data source provides the list of Managed Instance Group Available Software Sources in Oracle Cloud Infrastructure Os Management Hub service.

Lists available software sources for a specified managed instance group. Filter the list against a variety of criteria including but not limited to the software source name. The results list only software sources that have not already been added to the group.


## Example Usage

```hcl
data "oci_os_management_hub_managed_instance_group_available_software_sources" "test_managed_instance_group_available_software_sources" {
	#Required
	managed_instance_group_id = oci_os_management_hub_managed_instance_group.test_managed_instance_group.id

	#Optional
	compartment_id = var.compartment_id
	display_name = var.managed_instance_group_available_software_source_display_name
	display_name_contains = var.managed_instance_group_available_software_source_display_name_contains
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The OCID of the compartment that contains the resources to list. This filter returns only resources contained within the specified compartment.
* `display_name` - (Optional) A filter to return resources that match the given display names.
* `display_name_contains` - (Optional) A filter to return resources that may partially match the given display name.
* `managed_instance_group_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance group.


## Attributes Reference

The following attributes are exported:

* `available_software_source_collection` - The list of available_software_source_collection.

### ManagedInstanceGroupAvailableSoftwareSource Reference

The following attributes are exported:

* `items` - List of available software sources.
	* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the software source.
	* `display_name` - User-friendly name for the software source.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source.

