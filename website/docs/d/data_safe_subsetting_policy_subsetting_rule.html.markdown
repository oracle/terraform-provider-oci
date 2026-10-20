---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_subsetting_policy_subsetting_rule"
sidebar_current: "docs-oci-datasource-data_safe-subsetting_policy_subsetting_rule"
description: |-
  Provides details about a specific Subsetting Policy Subsetting Rule in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_subsetting_policy_subsetting_rule
This data source provides details about a specific Subsetting Policy Subsetting Rule resource in Oracle Cloud Infrastructure Data Safe service.

Gets the details of the specified subsetting rule.

## Example Usage

```hcl
data "oci_data_safe_subsetting_policy_subsetting_rule" "test_subsetting_policy_subsetting_rule" {
	#Required
	subsetting_policy_id = oci_data_safe_subsetting_policy.test_subsetting_policy.id
	subsetting_rule_key = var.subsetting_policy_subsetting_rule_subsetting_rule_key
}
```

## Argument Reference

The following arguments are supported:

* `subsetting_policy_id` - (Required) The OCID of the subsetting policy.
* `subsetting_rule_key` - (Required) The unique key that identifies the subsetting rule. It's numeric and unique within a subsetting policy.


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

