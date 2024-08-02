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

Returns information about the specified software source.

## Example Usage

```hcl
data "oci_os_management_hub_software_source" "test_software_source" {
	#Required
	software_source_id = oci_os_management_hub_software_source.test_software_source.id
}
```

## Argument Reference

The following arguments are supported:

* `software_source_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source.


## Attributes Reference

The following attributes are exported:

* `arch_type` - The architecture type supported by the software source.
* `availability` - Availability of the software source (for non-OCI environments).
* `availability_at_oci` - Availability of the software source (for Oracle Cloud Infrastructure environments).
* `checksum_type` - The yum repository checksum type used by this software source.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the software source.
* `custom_software_source_filter` - Provides the information used to apply filters to a vendor software source to create or update a custom software source.
	* `module_stream_profile_filters` - The list of module stream/profile filters.
		* `filter_type` - The type of the filter.
		* `module_name` - Module name.
		* `profile_name` - Profile name.
		* `stream_name` - Stream name.
	* `package_filters` - The list of package filters.
		* `filter_type` - The type of the filter.
		* `package_name` - The package name.
		* `package_name_pattern` - The package name pattern.
		* `package_version` - The package version, which is denoted by 'version-release', or 'epoch:version-release'.
	* `package_group_filters` - The list of group filters.
		* `filter_type` - The type of the filter.
		* `package_groups` - List of package group names.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - User-specified description for the software source.
* `display_name` - User-friendly name for the software source.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `gpg_key_fingerprint` - Fingerprint of the GPG key for this software source.
* `gpg_key_id` - ID of the GPG key for this software source.
* `gpg_key_url` - URL of the GPG key for this software source.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source.
* `is_auto_resolve_dependencies` - Indicates whether the service should automatically resolve package dependencies when including specific packages in the software source.
* `is_automatically_updated` - Indicates whether the service should automatically update the custom software source to use the latest package versions available. The service reviews packages levels once a day.
* `is_created_from_package_list` - Indicates whether the service should create the software source from a list of packages provided by the user.
* `is_latest_content_only` - Indicates whether the software source will include only the latest versions of content from vendor software sources, while accounting for other constraints set in the custom or versioned custom software source (such as a package list or filters).
	* For a module filter that does not specify a stream, this will include all available streams, and within each stream only the latest version of packages.
	* For a module filter that does specify a stream, this will include only the latest version of packages for the specified stream.
	* For a package filter that does not specify a version, this will include only the latest available version of the package.
	* For a package filter that does specify a version, this will include only the specified version of the package (the isLatestContentOnly attribute is ignored).
	* For a package list, this will include only the specified version of packages and modules in the list (the isLatestContentOnly attribute is ignored). 
* `is_mandatory_for_autonomous_linux` - Indicates whether the software source is required for the Autonomous Linux service.
* `origin_software_source_id` - This property applies only to replicated vendor software sources. This is the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the vendor software source in the root compartment.
* `os_family` - The OS family the software source belongs to.
* `package_count` - Number of packages the software source contains.
* `packages` - The packages in the software source.
* `repo_id` - The repository ID for the software source.
* `size` - The size of the software source in gigabytes (GB).
* `software_source_type` - Type of software source.
* `software_source_version` - The version to assign to this custom software source.
* `state` - The current state of the software source.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the software source was created (in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) format). 
* `url` - URL for the repository. For vendor software sources, this is the URL to the regional yum server. For custom software sources, this is 'custom/<repoId>'.
* `vendor_name` - Name of the vendor providing the software source.
* `vendor_software_sources` - List of vendor software sources that are used for the basis of the versioned custom software source.
	* `display_name` - User-friendly name.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource that is immutable on creation.

