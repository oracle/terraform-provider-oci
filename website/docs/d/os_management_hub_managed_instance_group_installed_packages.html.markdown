---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_managed_instance_group_installed_packages"
sidebar_current: "docs-oci-datasource-os_management_hub-managed_instance_group_installed_packages"
description: |-
  Provides the list of Managed Instance Group Installed Packages in Oracle Cloud Infrastructure Os Management Hub service
---

# Data Source: oci_os_management_hub_managed_instance_group_installed_packages
This data source provides the list of Managed Instance Group Installed Packages in Oracle Cloud Infrastructure Os Management Hub service.

Lists installed packages on the specified managed instances group. Filter the list against a variety 
of criteria including but not limited to the package name.


## Example Usage

```hcl
data "oci_os_management_hub_managed_instance_group_installed_packages" "test_managed_instance_group_installed_packages" {
	#Required
	managed_instance_group_id = oci_os_management_hub_managed_instance_group.test_managed_instance_group.id

	#Optional
	compartment_id = var.compartment_id
	display_name = var.managed_instance_group_installed_package_display_name
	display_name_contains = var.managed_instance_group_installed_package_display_name_contains
	time_install_date_end = var.managed_instance_group_installed_package_time_install_date_end
	time_install_date_start = var.managed_instance_group_installed_package_time_install_date_start
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The OCID of the compartment that contains the resources to list. This filter returns only resources contained within the specified compartment.
* `display_name` - (Optional) A filter to return resources that match the given display names.
* `display_name_contains` - (Optional) A filter to return resources that may partially match the given display name.
* `managed_instance_group_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance group.
* `time_install_date_end` - (Optional) A filter to return only packages that were installed on or before the date provided, in ISO 8601 format.  Example: 2017-07-14T02:40:00.000Z 
* `time_install_date_start` - (Optional) The install date after which to list all packages, in ISO 8601 format  Example: 2017-07-14T02:40:00.000Z 


## Attributes Reference

The following attributes are exported:

* `managed_instance_group_installed_package_collection` - The list of managed_instance_group_installed_package_collection.

### ManagedInstanceGroupInstalledPackage Reference

The following attributes are exported:

* `items` - List of installed packages.
	* `architecture` - The architecture of the package that is installed on the managed instance group. 
	* `name` - The name of the package that is installed on the managed instance group. 

