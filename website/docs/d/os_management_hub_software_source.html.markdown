---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_software_source"
sidebar_current: "docs-oci-datasource-os_management_hub-software_source"
description: |-
  Provides details about a specific Software Source in Oracle Cloud Infrastructure Os Management Hub service
---

# Data Source: oci_os_management_hub_software_source
This data source provides details about a specific Software Source resource in Oracle Cloud Infrastructure Os Management Hub service.

Gets information about the specified software source.

## Example Usage

```hcl
data "oci_os_management_hub_software_source" "test_software_source" {
	#Required
	software_source_id = oci_os_management_hub_software_source.test_software_source.id
}
```

## Argument Reference

The following arguments are supported:

* `software_source_id` - (Required) The software source OCID.


## Attributes Reference

The following attributes are exported:

* `arch_type` - The architecture type supported by the software source.
* `availability` - Possible availabilities of a software source.
* `checksum_type` - The yum repository checksum type used by this software source.
* `compartment_id` - The OCID of the tenancy containing the software source.
* `custom_software_source_filter` - Used to apply filters to a VendorSoftwareSource to create/update CustomSoftwareSources.
	* `module_stream_profile_filters` - The list of module stream/profile filters.
		* `filter_type` - The type of the filter, which can be of two types - INCLUDE or EXCLUDE.
		* `module_name` - Module name.
		* `profile_name` - Profile name.
		* `stream_name` - Stream name.
	* `package_filters` - The list of package filters.
		* `filter_type` - The type of the filter, which can be of two types - INCLUDE or EXCLUDE.
		* `package_name` - The package name.
		* `package_name_pattern` - The package name pattern.
		* `package_version` - The package version, which is denoted by 'version-release', or 'epoch:version-release'.
	* `package_group_filters` - The list of group filters.
		* `filter_type` - The type of the filter, which can be of two types - INCLUDE or EXCLUDE.
		* `package_groups` - List of package group names.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - Information specified by the user about the software source.
* `display_name` - User friendly name for the software source.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `gpg_key_fingerprint` - Fingerprint of the GPG key for this software source.
* `gpg_key_id` - ID of the GPG key for this software source.
* `gpg_key_url` - URL of the GPG key for this software source.
* `id` - OCID for the software source.
* `is_automatically_updated` - Indicates whether service should automatically update the custom software source for the user.
* `os_family` - The OS family the software source belongs to.
* `package_count` - Number of packages.
* `repo_id` - The Repo ID for the software source.
* `software_source_type` - Type of the software source.
* `software_source_version` - The version to assign to this custom software source.
* `state` - The current state of the software source.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the software source was created, as described in [RFC 3339](https://tools.ietf.org/rfc/rfc3339), section 14.29. 
* `url` - URL for the repository.
* `vendor_name` - Name of the vendor providing the software source.
* `vendor_software_sources` - List of vendor software sources.
	* `display_name` - User friendly name.
	* `id` - The OCID of the resource that is immutable on creation.

