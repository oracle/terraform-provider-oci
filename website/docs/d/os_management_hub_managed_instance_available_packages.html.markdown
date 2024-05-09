---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_managed_instance_available_packages"
sidebar_current: "docs-oci-datasource-os_management_hub-managed_instance_available_packages"
description: |-
  Provides the list of Managed Instance Available Packages in Oracle Cloud Infrastructure Os Management Hub service
---

# Data Source: oci_os_management_hub_managed_instance_available_packages
This data source provides the list of Managed Instance Available Packages in Oracle Cloud Infrastructure Os Management Hub service.

Returns a list of packages that are available for installation on a managed instance.


## Example Usage

```hcl
data "oci_os_management_hub_managed_instance_available_packages" "test_managed_instance_available_packages" {
	#Required
	managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance.id

	#Optional
	compartment_id = var.compartment_id
	display_name = var.managed_instance_available_package_display_name
	display_name_contains = var.managed_instance_available_package_display_name_contains
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

* `available_package_collection` - The list of available_package_collection.

### ManagedInstanceAvailablePackage Reference

The following attributes are exported:

* `items` - List of available packages.
	* `architecture` - The architecture for which this package was built.
	* `display_name` - Package name.
	* `name` - Unique identifier for the package.
	* `package_classification` - Status of the software package.
	* `software_sources` - List of software sources that provide the software package.
		* `description` - Software source description.
		* `display_name` - Software source name.
		* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source.
		* `is_mandatory_for_autonomous_linux` - Indicates whether this is a required software source for Autonomous Linux instances. If true, the user can't unselect it.
		* `software_source_type` - Type of the software source.
	* `type` - Type of the package.
	* `version` - Version of the installed package.

