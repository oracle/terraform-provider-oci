---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_software_sources"
sidebar_current: "docs-oci-datasource-os_management_hub-software_sources"
description: |-
  Provides the list of Software Sources in Oracle Cloud Infrastructure Os Management Hub service
---

# Data Source: oci_os_management_hub_software_sources
This data source provides the list of Software Sources in Oracle Cloud Infrastructure Os Management Hub service.

Lists software sources that match the specified tenancy or software source [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Filter the list against a
variety of criteria including but not limited to its name, status, architecture, and OS family.


## Example Usage

```hcl
data "oci_os_management_hub_software_sources" "test_software_sources" {

	#Optional
	arch_type = var.software_source_arch_type
	availability = var.software_source_availability
	availability_anywhere = var.software_source_availability_anywhere
	availability_at_oci = var.software_source_availability_at_oci
	compartment_id = var.compartment_id
	display_name = var.software_source_display_name
	display_name_contains = var.software_source_display_name_contains
	display_name_not_equal_to = var.software_source_display_name_not_equal_to
	is_mandatory_for_autonomous_linux = var.software_source_is_mandatory_for_autonomous_linux
	os_family = var.software_source_os_family
	software_source_id = oci_os_management_hub_software_source.test_software_source.id
	software_source_type = var.software_source_software_source_type
	state = var.software_source_state
	vendor_name = var.software_source_vendor_name
}
```

## Argument Reference

The following arguments are supported:

* `arch_type` - (Optional) A filter to return only instances whose architecture type matches the given architecture.
* `availability` - (Optional) The availabilities of the software source in a non-OCI environment for a tenancy.
* `availability_anywhere` - (Optional) The availabilities of the software source. Use this query parameter to filter across availabilities in different environments.
* `availability_at_oci` - (Optional) The availabilities of the software source in an Oracle Cloud Infrastructure environment for a tenancy.
* `compartment_id` - (Optional) (Updatable) The OCID of the compartment that contains the resources to list. This filter returns only resources contained within the specified compartment.
* `display_name` - (Optional) A filter to return resources that match the given user-friendly name.
* `display_name_contains` - (Optional) A filter to return resources that may partially match the given display name.
* `display_name_not_equal_to` - (Optional) A multi filter to return resources that do not contains the given display names.
* `is_mandatory_for_autonomous_linux` - (Applicable when software_source_type=VENDOR) Indicates whether the software source is mandatory for the Autonomous Linux service.
* `os_family` - (Optional) A filter to return only resources that match the given operating system family.
* `software_source_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the software source.
* `software_source_type` - (Optional) The type of the software source.
* `state` - (Optional) A filter to return only software sources whose state matches the given state.
* `vendor_name` - (Applicable when software_source_type=VENDOR) A filter to return only resources that match the given vendor name.


## Attributes Reference

The following attributes are exported:

* `software_source_collection` - The list of software_source_collection.

### SoftwareSource Reference

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

