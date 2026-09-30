---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_subsetting_policy_subsetting_rule_processing_chain_object"
sidebar_current: "docs-oci-resource-data_safe-subsetting_policy_subsetting_rule_processing_chain_object"
description: |-
  Provides the Subsetting Policy Subsetting Rule Processing Chain Object resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_subsetting_policy_subsetting_rule_processing_chain_object
This resource provides the Subsetting Policy Subsetting Rule Processing Chain Object resource in Oracle Cloud Infrastructure Data Safe service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/data-safe/latest/ProcessingChainObject

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/datasafe

Updates one or more attributes of the specified processing chain object.


## Example Usage

```hcl
resource "oci_data_safe_subsetting_policy_subsetting_rule_processing_chain_object" "test_subsetting_policy_subsetting_rule_processing_chain_object" {
	#Required
	processing_chain_object_key = var.subsetting_policy_subsetting_rule_processing_chain_object_processing_chain_object_key
	subsetting_policy_id = oci_data_safe_subsetting_policy.test_subsetting_policy.id
	subsetting_rule_key = var.subsetting_policy_subsetting_rule_processing_chain_object_subsetting_rule_key

	#Optional
	is_enabled_for_processing = var.subsetting_policy_subsetting_rule_processing_chain_object_is_enabled_for_processing
}
```

## Argument Reference

The following arguments are supported:

* `is_enabled_for_processing` - (Required) (Updatable) Indicates if this object/edge is enabled for processing
* `processing_chain_object_key` - (Required) The unique key that identifies the processing chain object. It's numeric and unique within a processing chain.
* `subsetting_policy_id` - (Required) The OCID of the subsetting policy.
* `subsetting_rule_key` - (Required) The unique key that identifies the subsetting rule. It's numeric and unique within a subsetting policy.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `items` - An array of subsetting processing chain summary objects.
	* `approximate_row_count_before_subsetting` - The approximate count of rows in the subsetting table before subsetting
	* `child_columns` - Unique identifiers identifying the child columns in the relation.
	* `child_object_name` - The name of the child subsetting table
	* `child_schema_name` - The database schema that contains the child subsetting table
	* `estimated_row_count_after_subsetting` - The estimated count of rows in the subsetting table after subsetting
	* `is_enabled_for_processing` - Indicates if this object/edge is enabled for processing
	* `key` - The unique key that identifies a subsetting relation processed. The key is numeric and unique within a processing order
	* `parent_columns` - Unique identifiers identifying the parents columns in the relation.
	* `parent_object_name` - The name of the parent subsetting table
	* `parent_schema_name` - The database schema that contains the parent subsetting table
	* `propagation_impact` - The impact on the related table due to the processing of subsetting rule 
	* `subsetting_schema_relation_key` - The unique key that identifies a subsetting relation.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Subsetting Policy Subsetting Rule Processing Chain Object
	* `update` - (Defaults to 20 minutes), when updating the Subsetting Policy Subsetting Rule Processing Chain Object
	* `delete` - (Defaults to 20 minutes), when destroying the Subsetting Policy Subsetting Rule Processing Chain Object


## Import

SubsettingPolicySubsettingRuleProcessingChainObjects can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_subsetting_policy_subsetting_rule_processing_chain_object.test_subsetting_policy_subsetting_rule_processing_chain_object "subsettingPolicies/{subsettingPolicyId}/subsettingRules/{subsettingRuleKey}/processingChainObjects/{processingChainObjectKey}" 
```
