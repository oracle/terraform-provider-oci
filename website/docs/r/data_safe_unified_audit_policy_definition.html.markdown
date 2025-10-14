---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_unified_audit_policy_definition"
sidebar_current: "docs-oci-resource-data_safe-unified_audit_policy_definition"
description: |-
  Provides the Unified Audit Policy Definition resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_unified_audit_policy_definition
This resource provides the Unified Audit Policy Definition resource in Oracle Cloud Infrastructure Data Safe service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/data-safe/latest/UnifiedAuditPolicyDefinition

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/datasafe

Updates the unified audit policy definition.

## Example Usage

```hcl
resource "oci_data_safe_unified_audit_policy_definition" "test_unified_audit_policy_definition" {
	#Required
	unified_audit_policy_definition_id = oci_data_safe_unified_audit_policy_definition.test_unified_audit_policy_definition.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.unified_audit_policy_definition_description
	display_name = var.unified_audit_policy_definition_display_name
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) (Updatable) The OCID of the compartment containing the unified audit policy definition.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm) Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) The description of the audit policy.
* `display_name` - (Optional) (Updatable) The display name of the audit policy. The name does not have to be unique, and it is changeable.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `unified_audit_policy_definition_id` - (Required) The OCID of the unified audit policy definition resource.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `audit_policy_category` - The category to which the unified audit policy belongs in the target database.
* `compartment_id` - The OCID of the compartment containing the unified audit policy definition.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm) Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the unified audit policy definition.
* `display_name` - The display name of the unified audit policy definition.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the unified audit policy definition.
* `is_seeded` - Signifies whether the unified audit policy definition is seeded or not.
* `lifecycle_details` - Details about the current state of the unified audit policy definition.
* `policy_definition_statement` - The definition of the unified audit policy to be provisioned in the target database.
* `policy_name` - The unified audit policy name in the target database.
* `state` - The current state of the unified audit policy definition.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the unified audit policy was created, in the format defined by RFC3339.
* `time_updated` - The last date and time the unified audit policy was updated, in the format defined by RFC3339.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Unified Audit Policy Definition
	* `update` - (Defaults to 20 minutes), when updating the Unified Audit Policy Definition
	* `delete` - (Defaults to 20 minutes), when destroying the Unified Audit Policy Definition


## Import

UnifiedAuditPolicyDefinitions can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_unified_audit_policy_definition.test_unified_audit_policy_definition "id"
```

