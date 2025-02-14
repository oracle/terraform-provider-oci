---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_profile_version"
sidebar_current: "docs-oci-datasource-os_management_hub-profile_version"
description: |-
  Provides details about a specific Profile Version in Oracle Cloud Infrastructure Os Management Hub service
---

# Data Source: oci_os_management_hub_profile_version
This data source provides details about a specific Profile Version resource in Oracle Cloud Infrastructure Os Management Hub service.

Returns information about the version of the specified registration profile.

## Example Usage

```hcl
data "oci_os_management_hub_profile_version" "test_profile_version" {
	#Required
	profile_id = oci_os_management_hub_profile.test_profile.id
	profile_version = var.profile_version_profile_version
}
```

## Argument Reference

The following arguments are supported:

* `profile_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the registration profile.
* `profile_version` - (Required) The version of the registration profile.


## Attributes Reference

The following attributes are exported:

* `arch_type` - The architecture type.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the registration profile.
* `description` - The description of the registration profile.
* `display_name` - A user-friendly name for the profile.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the registration profile.
* `is_default_profile` - Indicates if the profile is set as the default. There is exactly one default profile for a specified architecture, OS family, registration type, and vendor. When registering an instance with the corresonding characteristics, the default profile is used, unless another profile is specified. 
* `is_service_provided_profile` - Indicates if the profile was created by the service. OS Management Hub provides a limited set of standardized profiles that can be used to register Autonomous Linux or Windows instances. 
* `lifecycle_environment` - Provides identifying information for the specified lifecycle environment.
	* `display_name` - Lifecycle environment name.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the lifecycle environment.
* `lifecycle_stage` - Provides identifying information for the specified lifecycle stage.
	* `display_name` - Lifecycle stage name.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the lifecycle stage.
* `managed_instance_group` - Provides identifying information for the specified managed instance group.
	* `display_name` - Managed instance group name.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance group.
* `management_station_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management station to associate with an  instance once registered. Management stations are only used with non-OCI instances. 
* `os_family` - The operating system family.
* `profile_type` - The type of profile.
* `profile_version` - The version of the profile.
* `registration_type` - The type of instance to register.
* `software_sources` - The list of software sources that the registration profile will use.
	* `description` - Software source description.
	* `display_name` - Software source name.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source.
	* `is_mandatory_for_autonomous_linux` - Indicates whether this is a required software source for Autonomous Linux instances. If true, the user can't unselect it.
	* `software_source_type` - Type of the software source.
* `state` - The current state of the registration profile.
* `time_created` - The time the registration profile was created (in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) format).
* `time_modified` - The time the registration profile was last modified (in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) format).
* `vendor_name` - The vendor of the operating system for the instance.

