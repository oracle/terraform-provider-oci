---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_software_source"
sidebar_current: "docs-oci-resource-os_management_hub-software_source"
description: |-
  Provides the Software Source resource in Oracle Cloud Infrastructure Os Management Hub service
---

# oci_os_management_hub_software_source
This resource provides the Software Source resource in Oracle Cloud Infrastructure Os Management Hub service.

Creates a new versioned or custom software source.


## Example Usage

```hcl
resource "oci_os_management_hub_software_source" "test_software_source" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.software_source_display_name
	software_source_type = var.software_source_software_source_type
	vendor_software_sources {
		#Required
		display_name = var.software_source_vendor_software_sources_display_name
		id = var.software_source_vendor_software_sources_id
	}

	#Optional
	custom_software_source_filter {

		#Optional
		module_stream_profile_filters {
			#Required
			filter_type = var.software_source_custom_software_source_filter_module_stream_profile_filters_filter_type
			module_name = var.software_source_custom_software_source_filter_module_stream_profile_filters_module_name

			#Optional
			profile_name = oci_os_management_hub_profile.test_profile.name
			stream_name = oci_streaming_stream.test_stream.name
		}
		package_filters {
			#Required
			filter_type = var.software_source_custom_software_source_filter_package_filters_filter_type

			#Optional
			package_name = var.software_source_custom_software_source_filter_package_filters_package_name
			package_name_pattern = var.software_source_custom_software_source_filter_package_filters_package_name_pattern
			package_version = var.software_source_custom_software_source_filter_package_filters_package_version
		}
		package_group_filters {
			#Required
			filter_type = var.software_source_custom_software_source_filter_package_group_filters_filter_type

			#Optional
			package_groups = var.software_source_custom_software_source_filter_package_group_filters_package_groups
		}
	}
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.software_source_description
	freeform_tags = {"Department"= "Finance"}
	is_automatically_updated = var.software_source_is_automatically_updated
	software_source_version = var.software_source_software_source_version
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the tenancy containing the software source.
* `custom_software_source_filter` - (Optional) (Updatable) Used to apply filters to a VendorSoftwareSource to create/update CustomSoftwareSources.
	* `module_stream_profile_filters` - (Optional) (Updatable) The list of module stream/profile filters.
		* `filter_type` - (Required) (Updatable) The type of the filter, which can be of two types - INCLUDE or EXCLUDE.
		* `module_name` - (Required) (Updatable) Module name.
		* `profile_name` - (Optional) (Updatable) Profile name.
		* `stream_name` - (Optional) (Updatable) Stream name.
	* `package_filters` - (Optional) (Updatable) The list of package filters.
		* `filter_type` - (Required) (Updatable) The type of the filter, which can be of two types - INCLUDE or EXCLUDE.
		* `package_name` - (Optional) (Updatable) The package name.
		* `package_name_pattern` - (Optional) (Updatable) The package name pattern.
		* `package_version` - (Optional) (Updatable) The package version, which is denoted by 'version-release', or 'epoch:version-release'.
	* `package_group_filters` - (Optional) (Updatable) The list of group filters.
		* `filter_type` - (Required) (Updatable) The type of the filter, which can be of two types - INCLUDE or EXCLUDE.
		* `package_groups` - (Optional) (Updatable) List of package group names.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) Information specified by the user about the software source.
* `display_name` - (Required) (Updatable) User friendly name for the software source.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `is_automatically_updated` - (Applicable when software_source_type=CUSTOM) (Updatable) Indicates whether service should automatically update the custom software source for the user.
* `software_source_type` - (Required) (Updatable) Type of the software source.
* `software_source_version` - (Required when software_source_type=VERSIONED) The version to assign to this custom software source.
* `vendor_software_sources` - (Required) (Updatable) List of vendor software sources.
	* `display_name` - (Required) (Updatable) User friendly name.
	* `id` - (Required) (Updatable) The OCID of the resource that is immutable on creation.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Software Source
	* `update` - (Defaults to 20 minutes), when updating the Software Source
	* `delete` - (Defaults to 20 minutes), when destroying the Software Source


## Import

SoftwareSources can be imported using the `id`, e.g.

```
$ terraform import oci_os_management_hub_software_source.test_software_source "id"
```

