---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_lifecycle_environments"
sidebar_current: "docs-oci-datasource-os_management_hub-lifecycle_environments"
description: |-
  Provides the list of Lifecycle Environments in Oracle Cloud Infrastructure Os Management Hub service
---

# Data Source: oci_os_management_hub_lifecycle_environments
This data source provides the list of Lifecycle Environments in Oracle Cloud Infrastructure Os Management Hub service.

Lists lifecycle environments that match the specified compartment or lifecycle environment OCID. Filter the list
against a variety of criteria including but not limited to its name, status, architecture, and OS family.


## Example Usage

```hcl
data "oci_os_management_hub_lifecycle_environments" "test_lifecycle_environments" {

	#Optional
	arch_type = var.lifecycle_environment_arch_type
	compartment_id = var.compartment_id
	display_name = var.lifecycle_environment_display_name
	display_name_contains = var.lifecycle_environment_display_name_contains
	lifecycle_environment_id = oci_os_management_hub_lifecycle_environment.test_lifecycle_environment.id
	location = var.lifecycle_environment_location
	location_not_equal_to = var.lifecycle_environment_location_not_equal_to
	os_family = var.lifecycle_environment_os_family
	state = var.lifecycle_environment_state
}
```

## Argument Reference

The following arguments are supported:

* `arch_type` - (Optional) A filter to return only profiles that match the given archType.
* `compartment_id` - (Optional) (Updatable) The OCID of the compartment that contains the resources to list. This filter returns only resources contained within the specified compartment.
* `display_name` - (Optional) A filter to return resources that match the given display names.
* `display_name_contains` - (Optional) A filter to return resources that may partially match the given display name.
* `lifecycle_environment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the lifecycle environment.
* `location` - (Optional) A filter to return only resources whose location matches the given value.
* `location_not_equal_to` - (Optional) A filter to return only resources whose location does not match the given value.
* `os_family` - (Optional) A filter to return only resources that match the given operating system family.
* `state` - (Optional) A filter to return only the lifecycle environments that match the display name given.


## Attributes Reference

The following attributes are exported:

* `lifecycle_environment_collection` - The list of lifecycle_environment_collection.

### LifecycleEnvironment Reference

The following attributes are exported:

* `arch_type` - The CPU architecture of the managed instances in the lifecycle environment.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the lifecycle environment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - User-specified information about the lifecycle environment.
* `display_name` - The user-friendly name for the lifecycle environment.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the lifecycle environment.
* `location` - The location of managed instances attached to the lifecycle environment.
* `os_family` - The operating system of the managed instances in the lifecycle environment.
* `stages` - User-specified list of lifecycle stages used within the lifecycle environment.
	* `arch_type` - The CPU architecture of the managed instances in the lifecycle stage.
	* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the lifecycle stage.
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
	* `display_name` - The user-friendly name for the lifecycle stage.
	* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the lifecycle stage.
	* `lifecycle_environment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the lifecycle environment that contains the lifecycle stage.
    * `lifecycle_environment_display_name` - The user-friendly name for the lifecycle environment. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	* `location` - The location of managed instances associated with the lifecycle stage.
    * `managed_instances` - The number of managed instances associated with the lifecycle stage.
	* `os_family` - The operating system of the managed instances in the lifecycle stage.
	* `rank` - User-specified rank for the lifecycle stage. Rank determines the hierarchy of the lifecycle stages within the lifecycle environment. 
	* `software_source_id` - Provides identifying information for the specified software source.
		* `description` - Software source description.
		* `display_name` - Software source name.
		* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source.
		* `is_mandatory_for_autonomous_linux` - Indicates whether this is a required software source for Autonomous Linux instances. If true, the user can't unselect it.
		* `software_source_type` - Type of the software source.
	* `state` - The current state of the lifecycle stage.
	* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
	* `time_created` - The time the lifecycle stage was created (in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) format).
	* `time_modified` - The time the lifecycle stage was last modified (in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) format).
	* `vendor_name` - The vendor of the operating system used by the managed instances in the lifecycle stage.
* `state` - The current state of the lifecycle environment.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the lifecycle environment was created (in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) format).
* `time_modified` - The time the lifecycle environment was last modified (in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) format).
* `vendor_name` - The vendor of the operating system used by the managed instances in the lifecycle environment.

