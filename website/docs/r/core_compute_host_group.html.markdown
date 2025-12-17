---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_compute_host_group"
sidebar_current: "docs-oci-resource-core-compute_host_group"
description: |-
  Provides the Compute Host Group resource in Oracle Cloud Infrastructure Core service
---

# oci_core_compute_host_group
This resource provides the Compute Host Group resource in Oracle Cloud Infrastructure Core service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/iaas/latest/ComputeHostGroup

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/

Creates a new compute host group in the specified compartment and availability domain.

## Example Usage

```hcl
resource "oci_core_compute_host_group" "test_compute_host_group" {
	#Required
	availability_domain = var.compute_host_group_availability_domain
	compartment_id = var.compartment_id
	display_name = var.compute_host_group_display_name
	is_targeted_placement_required = var.compute_host_group_is_targeted_placement_required

	#Optional
	configurations {

		#Optional
		firmware_bundle_id = oci_core_firmware_bundle.test_firmware_bundle.id
		recycle_level = var.compute_host_group_configurations_recycle_level
		state = var.compute_host_group_configurations_state
		target = var.compute_host_group_configurations_target
	}
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Required) The availability domain of a host group.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) (Updatable) The OCID of the compartment that contains host group. 
* `configurations` - (Optional) (Updatable) A list of HostGroupConfiguration objects
	* `firmware_bundle_id` - (Optional) (Updatable) The OCID for firmware bundle
	* `recycle_level` - (Optional) (Updatable) Preferred recycle level for hosts associated with the reservation config.
		* `SKIP_RECYCLE` - Skips host wipe.
		* `FULL_RECYCLE` - Does not skip host wipe. This is the default behavior. 
	* `state` - (Optional) (Updatable) The state of the host group configuration.
	* `target` - (Optional) (Updatable) Either the platform name or compute shape that the configuration is targeting
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Required) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `is_targeted_placement_required` - (Required) (Updatable) A flag that allows customers to restrict placement for hosts attached to the group. If true, the only way to place on hosts is to target the specific host group.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `availability_domain` - The availability domain of a host group.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - The OCID of the compartment that contains host group. 
* `configurations` - A list of HostGroupConfiguration objects
	* `firmware_bundle_id` - The OCID for firmware bundle
	* `recycle_level` - Preferred recycle level for hosts associated with the reservation config.
		* `SKIP_RECYCLE` - Skips host wipe.
		* `FULL_RECYCLE` - Does not skip host wipe. This is the default behavior. 
	* `state` - The state of the host group configuration.
	* `target` - Either the platform name or compute shape that the configuration is targeting
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the Customer-unique host group 
* `is_targeted_placement_required` - A flag that allows customers to restrict placement for hosts attached to the group. If true, the only way to place on hosts is to target the specific host group.
* `state` - The lifecycle state of the host group 
* `system_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `time_created` - The date and time the host group was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the host group was updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Compute Host Group
	* `update` - (Defaults to 20 minutes), when updating the Compute Host Group
	* `delete` - (Defaults to 20 minutes), when destroying the Compute Host Group


## Import

ComputeHostGroups can be imported using the `id`, e.g.

```
$ terraform import oci_core_compute_host_group.test_compute_host_group "id"
```

