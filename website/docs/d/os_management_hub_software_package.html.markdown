---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_software_package"
sidebar_current: "docs-oci-datasource-os_management_hub-software_package"
description: |-
  Provides details about a specific Software Package in Oracle Cloud Infrastructure Os Management Hub service
---

# Data Source: oci_os_management_hub_software_package
This data source provides details about a specific Software Package resource in Oracle Cloud Infrastructure Os Management Hub service.

Returns information about the specified software package based on its fully qualified name.

## Example Usage

```hcl
data "oci_os_management_hub_software_package" "test_software_package" {
	#Required
	software_package_name = oci_os_management_hub_software_package.test_software_package.name
}
```

## Argument Reference

The following arguments are supported:

* `software_package_name` - (Required) The name of the software package.


## Attributes Reference

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

