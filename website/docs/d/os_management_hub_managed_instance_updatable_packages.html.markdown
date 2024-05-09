---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_managed_instance_updatable_packages"
sidebar_current: "docs-oci-datasource-os_management_hub-managed_instance_updatable_packages"
description: |-
  Provides the list of Managed Instance Updatable Packages in Oracle Cloud Infrastructure Os Management Hub service
---

# Data Source: oci_os_management_hub_managed_instance_updatable_packages
This data source provides the list of Managed Instance Updatable Packages in Oracle Cloud Infrastructure Os Management Hub service.

Returns a list of updatable packages for a managed instance.


## Example Usage

```hcl
data "oci_os_management_hub_managed_instance_updatable_packages" "test_managed_instance_updatable_packages" {
	#Required
	managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance.id

	#Optional
	advisory_name = var.managed_instance_updatable_package_advisory_name
	classification_type = var.managed_instance_updatable_package_classification_type
	compartment_id = var.compartment_id
	display_name = var.managed_instance_updatable_package_display_name
	display_name_contains = var.managed_instance_updatable_package_display_name_contains
}
```

## Argument Reference

The following arguments are supported:

* `advisory_name` - (Optional) The assigned erratum name. It's unique and not changeable.  Example: `ELSA-2020-5804` 
* `classification_type` - (Optional) A filter to return only packages that match the given update classification type.
* `compartment_id` - (Optional) The OCID of the compartment that contains the resources to list. This filter returns only resources contained within the specified compartment.
* `display_name` - (Optional) A filter to return resources that match the given display names.
* `display_name_contains` - (Optional) A filter to return resources that may partially match the given display name.
* `managed_instance_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance.


## Attributes Reference

The following attributes are exported:

* `updatable_package_collection` - The list of updatable_package_collection.

### ManagedInstanceUpdatablePackage Reference

The following attributes are exported:

* `items` - List of updatable packages.
	* `architecture` - The architecture for which this package was built.
	* `display_name` - Package name.
	* `errata` - List of errata applicable to this update.
	* `installed_version` - The version of the package that is currently installed on the instance.
	* `name` - Unique identifier for the package.
	* `package_classification` - Status of the software package.
	* `related_cves` - List of CVEs applicable to this erratum.
	* `software_sources` - List of software sources that provide the software package.
		* `description` - Software source description.
		* `display_name` - Software source name.
		* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source.
		* `is_mandatory_for_autonomous_linux` - Indicates whether this is a required software source for Autonomous Linux instances. If true, the user can't unselect it.
		* `software_source_type` - Type of the software source.
	* `type` - Type of the package.
	* `update_type` - The type of update.
	* `version` - Version of the installed package.

