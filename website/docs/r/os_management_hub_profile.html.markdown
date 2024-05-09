---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_profile"
sidebar_current: "docs-oci-resource-os_management_hub-profile"
description: |-
  Provides the Profile resource in Oracle Cloud Infrastructure Os Management Hub service
---

# oci_os_management_hub_profile
This resource provides the Profile resource in Oracle Cloud Infrastructure Os Management Hub service.

Creates a registration profile. A profile defines the content applied to the instance when registering it with the service.


## Example Usage

```hcl
resource "oci_os_management_hub_profile" "test_profile" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.profile_display_name
	profile_type = var.profile_profile_type

	#Optional
	arch_type = var.profile_arch_type
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.profile_description
	freeform_tags = {"Department"= "Finance"}
	is_default_profile = var.profile_is_default_profile
	lifecycle_stage_id = oci_os_management_hub_lifecycle_stage.test_lifecycle_stage.id
	managed_instance_group_id = oci_os_management_hub_managed_instance_group.test_managed_instance_group.id
	management_station_id = oci_os_management_hub_management_station.test_management_station.id
	os_family = var.profile_os_family
	registration_type = var.profile_registration_type
	software_source_ids = var.profile_software_source_ids
	vendor_name = var.profile_vendor_name
}
```

## Argument Reference

The following arguments are supported:

* `arch_type` - (Required when profile_type=SOFTWARESOURCE | STATION) The architecture type.
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the registration profile.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) User-specified description of the registration profile.
* `display_name` - (Required) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `is_default_profile` - (Optional) (Updatable) Indicates if the profile is set as the default. There is exactly one default profile for a specified architecture, OS family, registration type, and vendor. When registering an instance with the corresonding characteristics, the default profile is used, unless another profile is specified. 
* `lifecycle_stage_id` - (Required when profile_type=LIFECYCLE) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the lifecycle stage that the instance will be associated with.
* `managed_instance_group_id` - (Required when profile_type=GROUP) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance group that the instance will join after registration.
* `management_station_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management station to associate with an instance once registered. Associating with a management station applies only to non-OCI instances.
* `os_family` - (Required when profile_type=SOFTWARESOURCE | STATION) The operating system family.
* `profile_type` - (Required) The type of profile.
* `registration_type` - (Optional) The type of instance to register.
* `software_source_ids` - (Applicable when profile_type=SOFTWARESOURCE) The list of software source [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) that the registration profile will use.
* `vendor_name` - (Required when profile_type=SOFTWARESOURCE | STATION) The vendor of the operating system for the instance.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `arch_type` - The architecture type.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the registration profile.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the registration profile.
* `display_name` - A user-friendly name for the profile.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
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
* `management_station_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management station to associate with an instance once registered. Associating with a management station applies only to non-OCI instances.
* `os_family` - The operating system family.
* `profile_type` - The type of profile.
* `registration_type` - The type of instance to register.
* `software_sources` - The list of software sources that the registration profile will use.
	* `description` - Software source description.
	* `display_name` - Software source name.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source.
	* `is_mandatory_for_autonomous_linux` - Indicates whether this is a required software source for Autonomous Linux instances. If true, the user can't unselect it.
	* `software_source_type` - Type of the software source.
* `state` - The current state of the registration profile.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the registration profile was created (in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) format).
* `vendor_name` - The vendor of the operating system for the instance.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Profile
	* `update` - (Defaults to 20 minutes), when updating the Profile
	* `delete` - (Defaults to 20 minutes), when destroying the Profile


## Import

Profiles can be imported using the `id`, e.g.

```
$ terraform import oci_os_management_hub_profile.test_profile "id"
```

