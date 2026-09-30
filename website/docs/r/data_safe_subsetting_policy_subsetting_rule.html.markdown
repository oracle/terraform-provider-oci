---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_subsetting_policy_subsetting_rule"
sidebar_current: "docs-oci-resource-data_safe-subsetting_policy_subsetting_rule"
description: |-
  Provides the Subsetting Policy Subsetting Rule resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_subsetting_policy_subsetting_rule
This resource provides the Subsetting Policy Subsetting Rule resource in Oracle Cloud Infrastructure Data Safe service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/data-safe/latest/SubsettingRule

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/datasafe

Details to create a new subsetting rule

## Example Usage

```hcl
resource "oci_data_safe_subsetting_policy_subsetting_rule" "test_subsetting_policy_subsetting_rule" {
	#Required
	scope {
		#Required
		schema_name = var.subsetting_policy_subsetting_rule_scope_schema_name
		scope_type = var.subsetting_policy_subsetting_rule_scope_scope_type

		#Optional
		object = var.subsetting_policy_subsetting_rule_scope_object
	}
	subset_rule_entry {
		#Required
		rule_type = var.subsetting_policy_subsetting_rule_subset_rule_entry_rule_type

		#Optional
		condition = var.subsetting_policy_subsetting_rule_subset_rule_entry_condition
		partitions_list = var.subsetting_policy_subsetting_rule_subset_rule_entry_partitions_list
		percent = var.subsetting_policy_subsetting_rule_subset_rule_entry_percent
		sub_partitions_list = var.subsetting_policy_subsetting_rule_subset_rule_entry_sub_partitions_list
	}
	subsetting_policy_id = oci_data_safe_subsetting_policy.test_subsetting_policy.id

	#Optional
	description = var.subsetting_policy_subsetting_rule_description
	display_name = var.subsetting_policy_subsetting_rule_display_name
	peer_tables_action = var.subsetting_policy_subsetting_rule_peer_tables_action
	related_tables_propagation = var.subsetting_policy_subsetting_rule_related_tables_propagation
	rule_combination_mode = var.subsetting_policy_subsetting_rule_rule_combination_mode
}
```

## Argument Reference

The following arguments are supported:

* `description` - (Optional) (Updatable) The description of the subset rule
* `display_name` - (Optional) (Updatable) The display name of the subset rule
* `peer_tables_action` - (Optional) (Updatable) Strategy to be applied while processing peer tables
* `related_tables_propagation` - (Optional) (Updatable) Strategy to be applied while propagating subsetting rule to related tables
* `rule_combination_mode` - (Optional) (Updatable) Specifies how this rule combines with other rules. UNION evaluates this rule independently and adds matching rows to the result set. SERIAL applies this rule sequentially to filter rows selected by a compatible preceding rule. 
* `scope` - (Required) (Updatable) The scope of the subset rule
	* `object` - (Required when scope_type=SPECIFIC) (Updatable) The name of the specific object (e.g., table) to be subsetted
	* `schema_name` - (Required) (Updatable) The name of the schema containing the specific object to be subsetted
	* `scope_type` - (Required) (Updatable) Scope of a subsetting rule
* `subset_rule_entry` - (Required) (Updatable) The details of the subset rule
	* `condition` - (Required when rule_type=CONDITION) (Updatable) The SQL WHERE clause condition used to filter rows for the subset
	* `partitions_list` - (Applicable when rule_type=PARTITION) (Updatable) A list of partition names which are to be part of the subset data 
	* `percent` - (Required when rule_type=PERCENT) (Updatable) The percentage of rows to retain in the subset (between 0 and 100)
	* `rule_type` - (Required) (Updatable) type of subset rule
	* `sub_partitions_list` - (Applicable when rule_type=PARTITION) (Updatable) A list of sub-partition names which are to be part of the subset data. The sub-partition names should have the partition name also, separated by a dot 
* `subsetting_policy_id` - (Required) The OCID of the subsetting policy.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `description` - The description of the subset rule
* `display_name` - The display name of the subset rule
* `key` - The unique key that identifies a subsetting rule. The key is numeric and unique within a subsetting policy
* `peer_tables_action` - Strategy to be applied while processing peer tables
* `related_tables_propagation` - Strategy to be applied while propagating subsetting rule to related tables
* `rule_combination_mode` - Specifies how this rule combines with other rules. UNION evaluates this rule independently and adds matching rows to the result set. SERIAL applies this rule sequentially to filter rows selected by a compatible preceding rule. 
* `scope` - The scope of the subset rule
	* `object` - The name of the specific object (e.g., table) to be subsetted
	* `schema_name` - The name of the schema containing the specific object to be subsetted
	* `scope_type` - Scope of a subsetting rule
* `subset_rule_entry` - The details of the subset rule
	* `condition` - The SQL WHERE clause condition used to filter rows for the subset
	* `partitions_list` - A list of partition names which are to be part of the subset data 
	* `percent` - The percentage of rows to retain in the subset (between 0 and 100)
	* `rule_type` - type of subset rule
	* `sub_partitions_list` - A list of sub-partition names which are to be part of the subset data. The sub-partition names should have the partition name also, separated by a dot 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Subsetting Policy Subsetting Rule
	* `update` - (Defaults to 20 minutes), when updating the Subsetting Policy Subsetting Rule
	* `delete` - (Defaults to 20 minutes), when destroying the Subsetting Policy Subsetting Rule


## Import

SubsettingPolicySubsettingRules can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_subsetting_policy_subsetting_rule.test_subsetting_policy_subsetting_rule "subsettingPolicies/{subsettingPolicyId}/subsettingRules/{subsettingRuleKey}" 
```

