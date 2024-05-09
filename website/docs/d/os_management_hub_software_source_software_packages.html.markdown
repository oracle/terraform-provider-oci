---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_software_source_software_packages"
sidebar_current: "docs-oci-datasource-os_management_hub-software_source_software_packages"
description: |-
  Provides the list of Software Source Software Packages in Oracle Cloud Infrastructure Os Management Hub service
---

# Data Source: oci_os_management_hub_software_source_software_packages
This data source provides the list of Software Source Software Packages in Oracle Cloud Infrastructure Os Management Hub service.

Lists software packages in the specified software source.  Filter the list against a variety of criteria 
including but not limited to its name.


## Example Usage

```hcl
data "oci_os_management_hub_software_source_software_packages" "test_software_source_software_packages" {
	#Required
	software_source_id = oci_os_management_hub_software_source.test_software_source.id

	#Optional
	display_name = var.software_source_software_package_display_name
	display_name_contains = var.software_source_software_package_display_name_contains
	is_latest = var.software_source_software_package_is_latest
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) A filter to return resources that match the given user-friendly name.
* `display_name_contains` - (Optional) A filter to return resources that may partially match the given display name.
* `is_latest` - (Optional) Indicates whether to list only the latest versions of packages, module streams, and stream profiles.
* `software_source_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source.


## Attributes Reference

The following attributes are exported:

* `software_package_collection` - The list of software_package_collection.

### SoftwareSourceSoftwarePackage Reference

The following attributes are exported:

* `architecture` - The architecture for which this software was built
* `checksum` - Checksum of the package.
* `checksum_type` - Type of the checksum.
* `dependencies` - List of dependencies for the software package.
	* `dependency` - The software package's dependency.
	* `dependency_modifier` - The modifier for the dependency.
	* `dependency_type` - The type of the dependency.
* `description` - Description of the package.
* `display_name` - Package name.
* `files` - List of files for the software package.
	* `checksum` - Checksum of the file.
	* `checksum_type` - Type of the checksum.
	* `path` - File path.
	* `size_in_bytes` - Size of the file in bytes.
	* `time_modified` - The date and time the file was last modified (in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) format). 
	* `type` - Type of the file.
* `is_latest` - Indicates whether this package is the latest version.
* `last_modified_date` - The date and time the package was last modified (in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) format).
* `name` - Unique identifier for the package. Note that this is not an OCID.
* `os_families` - The OS families the package belongs to.
* `size_in_bytes` - Size of the package in bytes.
* `software_sources` - List of software sources that provide the software package. This property is deprecated and it will be removed in a future API release.
	* `description` - Software source description.
	* `display_name` - Software source name.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source.
	* `is_mandatory_for_autonomous_linux` - Indicates whether this is a required software source for Autonomous Linux instances. If true, the user can't unselect it.
	* `software_source_type` - Type of the software source.
* `type` - Type of the package.
* `version` - Version of the package.

