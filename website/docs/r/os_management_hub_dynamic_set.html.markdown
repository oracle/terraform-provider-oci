---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_dynamic_set"
sidebar_current: "docs-oci-resource-os_management_hub-dynamic_set"
description: |-
  Provides the Dynamic Set resource in Oracle Cloud Infrastructure Os Management Hub service
---

# oci_os_management_hub_dynamic_set
This resource provides the Dynamic Set resource in Oracle Cloud Infrastructure Os Management Hub service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/osmh/latest/DynamicSet

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/os_management_hub

Creates a new dynamic set.


## Example Usage

```hcl
resource "oci_os_management_hub_dynamic_set" "test_dynamic_set" {
	#Required
	compartment_id = var.compartment_id
	match_type = var.dynamic_set_match_type
	matching_rule {

		#Optional
		architectures = var.dynamic_set_matching_rule_architectures
		display_names = var.dynamic_set_matching_rule_display_names
		is_reboot_required = var.dynamic_set_matching_rule_is_reboot_required
		locations = var.dynamic_set_matching_rule_locations
		managed_instance_group_ids = var.dynamic_set_matching_rule_managed_instance_group_ids
		managed_instance_ids = var.dynamic_set_matching_rule_managed_instance_ids
		os_families = var.dynamic_set_matching_rule_os_families
		os_names = var.dynamic_set_matching_rule_os_names
		statuses = var.dynamic_set_matching_rule_statuses
		tags {
			#Required
			type = var.dynamic_set_matching_rule_tags_type

			#Optional
			key = var.dynamic_set_matching_rule_tags_key
			namespace = var.dynamic_set_matching_rule_tags_namespace
			value = var.dynamic_set_matching_rule_tags_value
		}
	}
	target_compartments {
		#Required
		compartment_id = var.compartment_id
		does_include_children = var.dynamic_set_target_compartments_does_include_children
	}

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.dynamic_set_description
	display_name = var.dynamic_set_display_name
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the dynamic set. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) User-specified description for the dynamic set.
* `display_name` - (Optional) (Updatable) User-friendly name for the dynamic set.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `match_type` - (Required) (Updatable) Include either any or all attributes.
* `matching_rule` - (Required) (Updatable) An object that defines the set of rules that identifies the target instances in a dynamic set.
	* `architectures` - (Optional) (Updatable) The list of managed instance architectures.
	* `display_names` - (Optional) (Updatable) The list of managed instance display names.
	* `is_reboot_required` - (Optional) (Updatable) Indicates if the managed instance needs to be rebooted.
	* `locations` - (Optional) (Updatable) The list of managed instance locations.
	* `managed_instance_group_ids` - (Optional) (Updatable) The list of managed instance group IDs.
	* `managed_instance_ids` - (Optional) (Updatable) The list of managed instance ids.
	* `os_families` - (Optional) (Updatable) The list of managed instance OS families.
	* `os_names` - (Optional) (Updatable) The list of managed instance OS names.
	* `statuses` - (Optional) (Updatable) The list of managed instance statuses.
	* `tags` - (Optional) (Updatable) The list of the managed instance tags.
		* `key` - (Optional) (Updatable) The key of the tag. 
		* `namespace` - (Applicable when type=DEFINED) (Updatable) The namespace of the tag. 
		* `type` - (Required) (Updatable) The type of the tag. Common values include `defined` or `freeform`. 
		* `value` - (Optional) (Updatable) The value associated with the tag key. 
* `target_compartments` - (Required) (Updatable) The list of compartment details.
	* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	* `does_include_children` - (Required) (Updatable) Indicates if the child compartments are included in the matching rule. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the dynamic set. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - User-specified description for the dynamic set.
* `display_name` - User-friendly name for the dynamic set.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the dynamic set. 
* `match_type` - Include either any or all attributes.
* `matching_rule` - An object that defines the set of rules that identifies the target instances in a dynamic set.
	* `architectures` - The list of managed instance architectures.
	* `display_names` - The list of managed instance display names.
	* `is_reboot_required` - Indicates if the managed instance needs to be rebooted.
	* `locations` - The list of managed instance locations.
	* `managed_instance_group_ids` - The list of managed instance group IDs.
	* `managed_instance_ids` - The list of managed instance ids.
	* `os_families` - The list of managed instance OS families.
	* `os_names` - The list of managed instance OS names.
	* `statuses` - The list of managed instance statuses.
	* `tags` - The list of the managed instance tags.
		* `key` - The key of the tag. 
		* `namespace` - The namespace of the tag. 
		* `type` - The type of the tag. Common values include `defined` or `freeform`. 
		* `value` - The value associated with the tag key. 
* `scheduled_job_count` - Number of scheduled jobs currently targeting this dynamic set.
* `state` - The current state of the event.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `target_compartments` - The list of compartment details.
	* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	* `does_include_children` - Indicates if the child compartments are included in the matching rule. 
* `time_created` - The date and time the dynamic set was created (in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) format). 
* `time_updated` - The date and time the dynamic set was last updated (in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) format). 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Dynamic Set
	* `update` - (Defaults to 20 minutes), when updating the Dynamic Set
	* `delete` - (Defaults to 20 minutes), when destroying the Dynamic Set


## Import

DynamicSets can be imported using the `id`, e.g.

```
$ terraform import oci_os_management_hub_dynamic_set.test_dynamic_set "id"
```

