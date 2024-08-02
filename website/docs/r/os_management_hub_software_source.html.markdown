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
	software_source_type = var.software_source_software_source_type

	#Optional
	custom_software_source_filter {

		#Optional
		module_stream_profile_filters {

			#Optional
			filter_type = var.software_source_custom_software_source_filter_module_stream_profile_filters_filter_type
			module_name = var.software_source_custom_software_source_filter_module_stream_profile_filters_module_name
			profile_name = oci_os_management_hub_profile.test_profile.name
			stream_name = oci_streaming_stream.test_stream.name
		}
		package_filters {

			#Optional
			filter_type = var.software_source_custom_software_source_filter_package_filters_filter_type
			package_name = var.software_source_custom_software_source_filter_package_filters_package_name
			package_name_pattern = var.software_source_custom_software_source_filter_package_filters_package_name_pattern
			package_version = var.software_source_custom_software_source_filter_package_filters_package_version
		}
		package_group_filters {

			#Optional
			filter_type = var.software_source_custom_software_source_filter_package_group_filters_filter_type
			package_groups = var.software_source_custom_software_source_filter_package_group_filters_package_groups
		}
	}
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.software_source_description
	display_name = var.software_source_display_name
	freeform_tags = {"Department"= "Finance"}
	is_auto_resolve_dependencies = var.software_source_is_auto_resolve_dependencies
	is_automatically_updated = var.software_source_is_automatically_updated
	is_created_from_package_list = var.software_source_is_created_from_package_list
	is_latest_content_only = var.software_source_is_latest_content_only
	origin_software_source_id = oci_os_management_hub_software_source.test_software_source.id
	packages = var.software_source_packages
	software_source_version = var.software_source_software_source_version
	vendor_software_sources {

		#Optional
		display_name = var.software_source_vendor_software_sources_display_name
		id = var.software_source_vendor_software_sources_id
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the software source.
* `custom_software_source_filter` - (Applicable when software_source_type=CUSTOM | VERSIONED) (Updatable) Provides the information used to apply filters to a vendor software source to create or update a custom software source.
	* `module_stream_profile_filters` - (Applicable when software_source_type=CUSTOM | VERSIONED) (Updatable) The list of module stream/profile filters.
		* `filter_type` - (Required when software_source_type=CUSTOM | VERSIONED) (Updatable) The type of the filter.
		* `module_name` - (Required when software_source_type=CUSTOM | VERSIONED) (Updatable) Module name.
		* `profile_name` - (Applicable when software_source_type=CUSTOM | VERSIONED) (Updatable) Profile name.
		* `stream_name` - (Applicable when software_source_type=CUSTOM | VERSIONED) (Updatable) Stream name.
	* `package_filters` - (Applicable when software_source_type=CUSTOM | VERSIONED) (Updatable) The list of package filters.
		* `filter_type` - (Required when software_source_type=CUSTOM | VERSIONED) (Updatable) The type of the filter.
		* `package_name` - (Applicable when software_source_type=CUSTOM | VERSIONED) (Updatable) The package name.
		* `package_name_pattern` - (Applicable when software_source_type=CUSTOM | VERSIONED) (Updatable) The package name pattern.
		* `package_version` - (Applicable when software_source_type=CUSTOM | VERSIONED) (Updatable) The package version, which is denoted by 'version-release', or 'epoch:version-release'.
	* `package_group_filters` - (Applicable when software_source_type=CUSTOM | VERSIONED) (Updatable) The list of group filters.
		* `filter_type` - (Required when software_source_type=CUSTOM | VERSIONED) (Updatable) The type of the filter.
		* `package_groups` - (Applicable when software_source_type=CUSTOM | VERSIONED) (Updatable) List of package group names.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) User-specified description for the software source. Avoid entering confidential information.
* `display_name` - (Optional) (Updatable) User-friendly name for the software source. Does not have to be unique and you can change the name later. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `is_auto_resolve_dependencies` - (Applicable when software_source_type=CUSTOM | VERSIONED) (Updatable) Indicates whether the service should automatically resolve package dependencies when including specific packages in the software source.
* `is_automatically_updated` - (Applicable when software_source_type=CUSTOM) (Updatable) Indicates whether the service should automatically update the custom software source to use the latest package versions available. The service reviews packages levels once a day.
* `is_created_from_package_list` - (Applicable when software_source_type=CUSTOM | VERSIONED) Indicates whether the service should create the software source from a list of packages provided by the user.
* `is_latest_content_only` - (Applicable when software_source_type=CUSTOM | VERSIONED) (Updatable) Indicates whether the software source will include only the latest versions of content from vendor software sources, while accounting for other constraints set in the custom or versioned custom software source (such as a package list or filters).
	* For a module filter that does not specify a stream, this will include all available streams, and within each stream only the latest version of packages.
	* For a module filter that does specify a stream, this will include only the latest version of packages for the specified stream.
	* For a package filter that does not specify a version, this will include only the latest available version of the package.
	* For a package filter that does specify a version, this will include only the specified version of the package (the isLatestContentOnly attribute is ignored).
	* For a package list, this will include only the specified version of packages and modules in the list (the isLatestContentOnly attribute is ignored). 
* `origin_software_source_id` - (Required when software_source_type=VENDOR) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the vendor software source in the root compartment that is being replicated.
* `packages` - (Applicable when software_source_type=CUSTOM | VERSIONED) A property used for compatibility only. It doesn't provide a complete list of packages. See [AddPackagesToSoftwareSourceDetails](https://docs.cloud.oracle.com/iaas/api/#/en/osmh/latest/datatypes/AddPackagesToSoftwareSourceDetails) for providing the list of packages used to create the software source when isCreatedFromPackageList is set to true.
* `software_source_type` - (Required) (Updatable) Type of software source.
* `software_source_version` - (Required when software_source_type=VERSIONED) The version to assign to this custom software source.
* `vendor_software_sources` - (Required when software_source_type=CUSTOM | VERSIONED) (Updatable) List of vendor software sources.
	* `display_name` - (Required when software_source_type=CUSTOM | VERSIONED) (Updatable) User-friendly name.
	* `id` - (Required when software_source_type=CUSTOM | VERSIONED) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource that is immutable on creation.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

