---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_managed_instance_installed_packages"
sidebar_current: "docs-oci-datasource-os_management_hub-managed_instance_installed_packages"
description: |-
  Provides the list of Managed Instance Installed Packages in Oracle Cloud Infrastructure Os Management Hub service
---

# Data Source: oci_os_management_hub_managed_instance_installed_packages
This data source provides the list of Managed Instance Installed Packages in Oracle Cloud Infrastructure Os Management Hub service.

Lists the packages that are installed on the managed instance.


## Example Usage

```hcl
data "oci_os_management_hub_managed_instance_installed_packages" "test_managed_instance_installed_packages" {
	#Required
	managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance.id

	#Optional
	compartment_id = var.compartment_id
	display_name = var.managed_instance_installed_package_display_name
	display_name_contains = var.managed_instance_installed_package_display_name_contains
	time_install_date_end = var.managed_instance_installed_package_time_install_date_end
	time_install_date_start = var.managed_instance_installed_package_time_install_date_start
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The OCID of the compartment that contains the resources to list. This filter returns only resources contained within the specified compartment.
* `display_name` - (Optional) A filter to return resources that match the given display names.
* `display_name_contains` - (Optional) A filter to return resources that may partially match the given display name.
* `managed_instance_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance.
* `time_install_date_end` - (Optional) A filter to return only packages that were installed on or before the date provided, in ISO 8601 format.  Example: 2017-07-14T02:40:00.000Z 
* `time_install_date_start` - (Optional) The install date after which to list all packages, in ISO 8601 format  Example: 2017-07-14T02:40:00.000Z 


## Attributes Reference

The following attributes are exported:

* `installed_package_collection` - The list of installed_package_collection.

### ManagedInstanceInstalledPackage Reference

The following attributes are exported:

* `items` - List of installed packages.
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
	* `time_installed` - The date and time the package was installed, as described in [RFC 3339](https://tools.ietf.org/rfc/rfc3339), section 14.29. 
	* `time_issued` - The date and time the package was issued by a providing erratum (in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) format). 
	* `type` - Type of the package.
	* `version` - Version of the installed package.

