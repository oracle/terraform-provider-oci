---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_unified_audit_policy"
sidebar_current: "docs-oci-resource-data_safe-unified_audit_policy"
description: |-
  Provides the Unified Audit Policy resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_unified_audit_policy
This resource provides the Unified Audit Policy resource in Oracle Cloud Infrastructure Data Safe service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/data-safe/latest/UnifiedAuditPolicy

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/datasafe

Creates the specified unified audit policy.


## Example Usage

```hcl
resource "oci_data_safe_unified_audit_policy" "test_unified_audit_policy" {
	#Required
	compartment_id = var.compartment_id
	conditions {
		#Required
		entity_selection = var.unified_audit_policy_conditions_entity_selection
		entity_type = var.unified_audit_policy_conditions_entity_type
		operation_status = var.unified_audit_policy_conditions_operation_status

		#Optional
		attribute_set_id = oci_data_safe_attribute_set.test_attribute_set.id
		role_names = var.unified_audit_policy_conditions_role_names
		user_names = var.unified_audit_policy_conditions_user_names
	}
	security_policy_id = oci_data_safe_security_policy.test_security_policy.id
	status = var.unified_audit_policy_status
	unified_audit_policy_definition_id = oci_data_safe_unified_audit_policy_definition.test_unified_audit_policy_definition.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.unified_audit_policy_description
	display_name = var.unified_audit_policy_display_name
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment in which to create the unified audit policy.
* `conditions` - (Required) (Updatable) Lists the audit policy provisioning conditions.
	* `attribute_set_id` - (Required when entity_type=ATTRIBUTE_SET) (Updatable) The OCID of the attribute set.
	* `entity_selection` - (Required) (Updatable) Specifies whether to include or exclude the specified users or roles.
	* `entity_type` - (Required) (Updatable) The type of users or roles that the unified audit policy is applied to.
	* `operation_status` - (Required) (Updatable) The operation status that the policy must be enabled for.
	* `role_names` - (Required when entity_type=ROLE) (Updatable) List of roles that the policy must be enabled for.
	* `user_names` - (Required when entity_type=USER) (Updatable) The list of users that the unified audit policy is enabled for.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm) Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) The description of the unified audit policy in Data Safe.
* `display_name` - (Optional) (Updatable) The display name of the unified audit policy in Data Safe. The name is modifiable and does not need to be unique.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `security_policy_id` - (Required) The OCID of the security policy corresponding to the unified audit policy.
* `status` - (Required) (Updatable) Indicates whether the unified audit policy has been enabled or disabled.
* `unified_audit_policy_definition_id` - (Required) The OCID of the associated unified audit policy definition.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the unified audit policy.
* `conditions` - Lists the audit policy provisioning conditions.
	* `attribute_set_id` - The OCID of the attribute set.
	* `entity_selection` - Specifies whether to include or exclude the specified users or roles.
	* `entity_type` - The type of users or roles that the unified audit policy is applied to.
	* `operation_status` - The operation status that the policy must be enabled for.
	* `role_names` - List of roles that the policy must be enabled for.
	* `user_names` - The list of users that the unified audit policy is enabled for.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm) Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the unified audit policy.
* `display_name` - The display name of the unified audit policy.
* `enabled_entities` - Indicates on whom the audit policy is enabled.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the unified audit policy.
* `is_seeded` - Indicates whether the unified audit policy is seeded or not.
* `lifecycle_details` - The details of the current state of the unified audit policy in Data Safe.
* `security_policy_id` - The OCID of the security policy corresponding to the unified audit policy.
* `state` - The current state of the unified audit policy.
* `status` - Indicates whether the policy has been enabled or disabled.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the the unified audit policy was created, in the format defined by RFC3339.
* `time_updated` - The last date and time the unified audit policy was updated, in the format defined by RFC3339.
* `unified_audit_policy_definition_id` - The OCID of the associated unified audit policy definition.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Unified Audit Policy
	* `update` - (Defaults to 20 minutes), when updating the Unified Audit Policy
	* `delete` - (Defaults to 20 minutes), when destroying the Unified Audit Policy


## Import

UnifiedAuditPolicies can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_unified_audit_policy.test_unified_audit_policy "id"
```

