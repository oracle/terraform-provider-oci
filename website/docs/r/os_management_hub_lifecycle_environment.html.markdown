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

Creates a new lifecycle environment.


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
		display_name = var.lifecycle_environment_stages_display_name
		rank = var.lifecycle_environment_stages_rank

		#Optional
		defined_tags = {"Operations.CostCenter"= "42"}
		freeform_tags = {"Department"= "Finance"}
	}
	vendor_name = var.lifecycle_environment_vendor_name

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.lifecycle_environment_description
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `arch_type` - (Required) The CPU architecture of the managed instance(s) in the lifecycle environment.
* `compartment_id` - (Required) The OCID of the tenancy containing the lifecycle environment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) User specified information about the lifecycle environment.
* `display_name` - (Required) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `os_family` - (Required) The operating system type of the managed instance(s) in the lifecycle environment.
* `stages` - (Required) (Updatable) User specified list of ranked lifecycle stages to be created for the lifecycle environment.
	* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
	* `display_name` - (Required) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
	* `rank` - (Required) User specified rank for the lifecycle stage. Rank determines the hierarchy of the lifecycle stages for a given lifecycle environment. 
* `vendor_name` - (Required) The software source vendor name.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `arch_type` - The CPU architecture of the target instances.
* `compartment_id` - The OCID of the tenancy containing the lifecycle environment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - User specified information about the lifecycle environment.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of the resource that is immutable on creation.
* `managed_instance_ids` - The list of managed instance OCIDs specified in the lifecycle stage.
	* `display_name` - Managed instance name.
	* `id` - The OCID of the managed instance.
* `os_family` - The operating system type of the target instances.
* `stages` - User specified list of lifecycle stages to be created for the lifecycle environment.
	* `arch_type` - The CPU architecture of the target instances.
	* `compartment_id` - The OCID of the tenancy containing the lifecycle stage.
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
	* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
	* `id` - The lifecycle stage OCID that is immutable on creation.
	* `lifecycle_environment_id` - The OCID of the lifecycle environment for the lifecycle stage.
	* `managed_instance_ids` - The list of managed instances specified lifecycle stage.
		* `display_name` - Managed instance name.
		* `id` - The OCID of the managed instance.
	* `os_family` - The operating system type of the target instances.
	* `rank` - User specified rank for the lifecycle stage. Rank determines the hierarchy of the lifecycle stages for a given lifecycle environment. 
	* `software_source_id` - Identifying information for the specified software source.
		* `description` - Software source description.
		* `display_name` - Software source name.
		* `id` - The OCID of the software source.
		* `software_source_type` - Type of the software source.
	* `state` - The current state of the lifecycle stage.
	* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
	* `time_created` - The time the lifecycle stage was created. An RFC3339 formatted datetime string.
	* `time_modified` - The time the lifecycle stage was last modified. An RFC3339 formatted datetime string.
	* `vendor_name` - The software source vendor name.
* `state` - The current state of the lifecycle environment.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the lifecycle environment was created. An RFC3339 formatted datetime string.
* `time_modified` - The time the lifecycle environment was last modified. An RFC3339 formatted datetime string.
* `vendor_name` - The software source vendor name.

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

