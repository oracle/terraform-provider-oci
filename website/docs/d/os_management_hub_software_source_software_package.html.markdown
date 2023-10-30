---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_software_source_software_package"
sidebar_current: "docs-oci-datasource-os_management_hub-software_source_software_package"
description: |-
  Provides details about a specific Software Source Software Package in Oracle Cloud Infrastructure Os Management Hub service
---

# Data Source: oci_os_management_hub_software_source_software_package
This data source provides details about a specific Software Source Software Package resource in Oracle Cloud Infrastructure Os Management Hub service.

Gets information about the specified software package.


## Example Usage

```hcl
data "oci_os_management_hub_software_source_software_package" "test_software_source_software_package" {
	#Required
	software_package_name = var.software_source_software_package_software_package_name
	software_source_id = oci_os_management_hub_software_source.test_software_source.id
}
```

## Argument Reference

The following arguments are supported:

* `software_package_name` - (Required) The name of the software package.
* `software_source_id` - (Required) The software source OCID.


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
	* `time_modified` - The date and time of the last modification to this file, as described in [RFC 3339](https://tools.ietf.org/rfc/rfc3339), section 14.29. 
	* `type` - Type of the file.
* `is_latest` - Indicates whether this package is the latest version.
* `last_modified_date` - Date of the last update to the package.
* `name` - Unique identifier for the package. NOTE - This is not an OCID.
* `size_in_bytes` - Size of the package in bytes.
* `software_sources` - List of software sources that provide the software package.
	* `description` - Software source description.
	* `display_name` - Software source name.
	* `id` - The OCID of the software source.
	* `software_source_type` - Type of the software source.
* `type` - Type of the package.
* `version` - Version of the package.

