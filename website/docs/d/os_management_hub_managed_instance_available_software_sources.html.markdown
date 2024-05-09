---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_managed_instance_available_software_sources"
sidebar_current: "docs-oci-datasource-os_management_hub-managed_instance_available_software_sources"
description: |-
  Provides the list of Managed Instance Available Software Sources in Oracle Cloud Infrastructure Os Management Hub service
---

# Data Source: oci_os_management_hub_managed_instance_available_software_sources
This data source provides the list of Managed Instance Available Software Sources in Oracle Cloud Infrastructure Os Management Hub service.

Returns a list of software sources that can be attached to the specified managed instance. Any software sources already attached to the instance are not included in the list.


## Example Usage

```hcl
data "oci_os_management_hub_managed_instance_available_software_sources" "test_managed_instance_available_software_sources" {
	#Required
	managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance.id

	#Optional
	compartment_id = var.compartment_id
	display_name = var.managed_instance_available_software_source_display_name
	display_name_contains = var.managed_instance_available_software_source_display_name_contains
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The OCID of the compartment that contains the resources to list. This filter returns only resources contained within the specified compartment.
* `display_name` - (Optional) A filter to return resources that match the given display names.
* `display_name_contains` - (Optional) A filter to return resources that may partially match the given display name.
* `managed_instance_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance.


## Attributes Reference

The following attributes are exported:

* `available_software_source_collection` - The list of available_software_source_collection.

### ManagedInstanceAvailableSoftwareSource Reference

The following attributes are exported:

* `items` - List of available software sources.
	* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the software source.
	* `display_name` - User-friendly name for the software source.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source.

