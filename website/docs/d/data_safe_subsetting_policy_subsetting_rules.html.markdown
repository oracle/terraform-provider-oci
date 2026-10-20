---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_subsetting_policy_subsetting_rules"
sidebar_current: "docs-oci-datasource-data_safe-subsetting_policy_subsetting_rules"
description: |-
  Provides the list of Subsetting Policy Subsetting Rules in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_subsetting_policy_subsetting_rules
This data source provides the list of Subsetting Policy Subsetting Rules in Oracle Cloud Infrastructure Data Safe service.

Gets a list of subsetting rules present in the specified subsetting policy and based on the specified query parameters.
A subsetting rule is the criteria that tells Data Safe which rows to retain from the selected starting table for a subsetting operation.
It is the entry point for the subset. Data Safe uses this rule, along with the propagation setting,
to determine the related rows that must also be retained across parent and child tables.


## Example Usage

```hcl
data "oci_data_safe_subsetting_policy_subsetting_rules" "test_subsetting_policy_subsetting_rules" {
	#Required
	subsetting_policy_id = oci_data_safe_subsetting_policy.test_subsetting_policy.id

	#Optional
	object = var.subsetting_policy_subsetting_rule_object
	schema_name = var.subsetting_policy_subsetting_rule_schema_name
}
```

## Argument Reference

The following arguments are supported:

* `object` - (Optional) A filter to return only items related to a specific object name.
* `schema_name` - (Optional) A filter to return only items related to specific schema name.
* `subsetting_policy_id` - (Required) The OCID of the subsetting policy.


## Attributes Reference

The following attributes are exported:

* `subsetting_rule_collection` - The list of subsetting_rule_collection.

### SubsettingPolicySubsettingRule Reference

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

