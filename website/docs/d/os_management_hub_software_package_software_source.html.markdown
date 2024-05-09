---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_software_package_software_source"
sidebar_current: "docs-oci-datasource-os_management_hub-software_package_software_source"
description: |-
  Provides the list of Software Package Software Source in Oracle Cloud Infrastructure Os Management Hub service
---

# Data Source: oci_os_management_hub_software_package_software_source
This data source provides the list of Software Package Software Source in Oracle Cloud Infrastructure Os Management Hub service.

Lists the software sources in the tenancy that contain the software package. Filter the list against a
variety of criteria including but not limited to its name, type, architecture, and OS family.


## Example Usage

```hcl
data "oci_os_management_hub_software_package_software_source" "test_software_package_software_source" {
	#Required
	compartment_id = var.compartment_id
	software_package_name = oci_os_management_hub_software_package.test_software_package.name

	#Optional
	arch_type = var.software_package_software_source_arch_type
	availability = var.software_package_software_source_availability
	availability_anywhere = var.software_package_software_source_availability_anywhere
	availability_at_oci = var.software_package_software_source_availability_at_oci
	display_name = var.software_package_software_source_display_name
	display_name_contains = var.software_package_software_source_display_name_contains
	os_family = var.software_package_software_source_os_family
	software_source_type = var.software_package_software_source_software_source_type
	state = var.software_package_software_source_state
}
```

## Argument Reference

The following arguments are supported:

* `arch_type` - (Optional) A filter to return only instances whose architecture type matches the given architecture.
* `availability` - (Optional) The availabilities of the software source in a non-OCI environment for a tenancy.
* `availability_anywhere` - (Optional) The availabilities of the software source. Use this query parameter to filter across availabilities in different environments.
* `availability_at_oci` - (Optional) The availabilities of the software source in an Oracle Cloud Infrastructure environment for a tenancy.
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. This parameter is required and returns only resources contained within the specified compartment.
* `display_name` - (Optional) A filter to return resources that match the given user-friendly name.
* `display_name_contains` - (Optional) A filter to return resources that may partially match the given display name.
* `os_family` - (Optional) A filter to return only resources that match the given operating system family.
* `software_package_name` - (Required) The name of the software package.
* `software_source_type` - (Optional) The type of the software source.
* `state` - (Optional) A filter to return only software sources whose state matches the given state.


## Attributes Reference

The following attributes are exported:

* `software_source_collection` - The list of software_source_collection.

### SoftwarePackageSoftwareSource Reference

The following attributes are exported:

* `items` - List of software sources.
	* `arch_type` - The architecture type supported by the software source.
	* `availability` - Availability of the software source (for non-OCI environments).
	* `availability_at_oci` - Availability of the software source (for Oracle Cloud Infrastructure environments).
	* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the software source.
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
	* `description` - Description of the software source. For custom software sources, this is user-specified.
	* `display_name` - User-friendly name for the software source.
	* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source.
	* `is_mandatory_for_autonomous_linux` - Indicates whether the software source is required for the Autonomous Linux service.
	* `os_family` - The OS family the software source belongs to.
	* `package_count` - Number of packages the software source contains.
	* `repo_id` - The repository ID for the software source.
	* `size` - The size of the software source in gigabytes (GB).
	* `software_source_type` - Type of software source.
	* `software_source_version` - The version to assign to this custom software source.
	* `state` - The current state of the software source.
	* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
	* `time_created` - The date and time the software source was created (in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) format). 
	* `time_updated` - The date and time the software source was updated (in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) format). 
	* `url` - URL for the repository. For vendor software sources, this is the URL to the regional yum server. For custom software sources, this is 'custom/<repoId>'.
	* `vendor_name` - Name of the vendor providing the software source.
	* `vendor_software_sources` - List of vendor software sources that are used for the basis of the versioned custom software source.
		* `display_name` - User-friendly name.
		* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource that is immutable on creation.

