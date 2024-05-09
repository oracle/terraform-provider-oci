---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_lifecycle_environment"
sidebar_current: "docs-oci-resource-os_management_hub-lifecycle_environment"
description: |-
  Provides the Lifecycle Environment resource in Oracle Cloud Infrastructure Os Management Hub service
---

# oci_os_management_hub_lifecycle_environment
This resource provides the Lifecycle Environment resource in Oracle Cloud Infrastructure Os Management Hub service.

Creates a lifecycle environment. A lifecycle environment is a user-defined pipeline to deliver curated, versioned content in a prescribed, methodical manner.


## Example Usage

```hcl
resource "oci_os_management_hub_lifecycle_environment" "test_lifecycle_environment" {
	#Required
	arch_type = var.lifecycle_environment_arch_type
	compartment_id = var.compartment_id
	display_name = var.lifecycle_environment_display_name
	os_family = var.lifecycle_environment_os_family
	stages {
		#Required
		display_name = var.lifecycle_environment_stages_display_name_1
		rank = var.lifecycle_environment_stages_rank_1

		#Optional
		defined_tags = {"Operations.CostCenter"= "42"}
		freeform_tags = {"Department"= "Finance"}
	}
	stages {
		#Required
		display_name = var.lifecycle_environment_stages_display_name_2
		rank = var.lifecycle_environment_stages_rank_2

		#Optional
		defined_tags = {"Operations.CostCenter"= "42"}
		freeform_tags = {"Department"= "Finance"}
	}
	vendor_name = var.lifecycle_environment_vendor_name

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.lifecycle_environment_description
	freeform_tags = {"Department"= "Finance"}
	location = var.lifecycle_environment_location
}
```

## Argument Reference

The following arguments are supported:

* `arch_type` - (Required) The CPU architecture of the managed instances in the lifecycle environment.
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the lifecycle environment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) User-specified information about the lifecycle environment. Avoid entering confidential information.
* `display_name` - (Required) (Updatable) A user-friendly name for the lifecycle environment. Does not have to be unique and you can change the name later. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `location` - (Optional) The location of managed instances attached to the lifecycle environment. If no location is provided, the default is 'ON_PREMISE.' 
* `os_family` - (Required) The operating system of the managed instances in the lifecycle environment.
* `stages` - (Required) (Updatable) User-specified list of ranked lifecycle stages used within the lifecycle environment.
	* `compartment_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the lifecycle stage.
	* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
	* `display_name` - (Required) (Updatable) A user-friendly name for the lifecycle stage. Does not have to be unique and you can change the name later. Avoid entering confidential information.
	* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
	* `rank` - (Required) User-specified rank for the lifecycle stage. Rank determines the hierarchy of the lifecycle stages within the lifecycle environment. 
* `vendor_name` - (Required) The vendor of the operating system used by the managed instances in the lifecycle environment.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `arch_type` - The CPU architecture of the managed instances in the lifecycle environment.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the lifecycle environment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - User-specified information about the lifecycle environment.
* `display_name` - The user-friendly name for the lifecycle environment.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the lifecycle environment.
* `location` - The location of managed instances attached to the lifecycle environment.
* `managed_instance_ids` - List of managed instance [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) assigned to the lifecycle stage.
	* `display_name` - Managed instance name.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance.
* `os_family` - The operating system of the managed instances in the lifecycle environment.
* `stages` - User-specified list of lifecycle stages used within the lifecycle environment.
	* `arch_type` - The CPU architecture of the managed instances in the lifecycle stage.
	* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the lifecycle stage.
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
	* `display_name` - The user-friendly name for the lifecycle stage.
	* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the lifecycle stage.
	* `lifecycle_environment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the lifecycle environment that contains the lifecycle stage.
	* `location` - The location of managed instances associated with the lifecycle stage.
	* `managed_instance_ids` - The list of managed instances associated with the lifecycle stage.
		* `display_name` - Managed instance name.
		* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance.
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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Lifecycle Environment
	* `update` - (Defaults to 20 minutes), when updating the Lifecycle Environment
	* `delete` - (Defaults to 20 minutes), when destroying the Lifecycle Environment


## Import

LifecycleEnvironments can be imported using the `id`, e.g.

```
$ terraform import oci_os_management_hub_lifecycle_environment.test_lifecycle_environment "id"
```

