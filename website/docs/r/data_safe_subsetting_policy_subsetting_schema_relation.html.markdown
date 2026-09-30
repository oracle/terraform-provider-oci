---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_subsetting_policy_subsetting_schema_relation"
sidebar_current: "docs-oci-resource-data_safe-subsetting_policy_subsetting_schema_relation"
description: |-
  Provides the Subsetting Policy Subsetting Schema Relation resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_subsetting_policy_subsetting_schema_relation
This resource provides the Subsetting Policy Subsetting Schema Relation resource in Oracle Cloud Infrastructure Data Safe service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/data-safe/latest/SubsettingSchemaRelation

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/datasafe

Details to create a new referential relation.

## Example Usage

```hcl
resource "oci_data_safe_subsetting_policy_subsetting_schema_relation" "test_subsetting_policy_subsetting_schema_relation" {
	#Required
	child_columns = var.subsetting_policy_subsetting_schema_relation_child_columns
	child_object_name = oci_objectstorage_object.test_object.name
	child_schema_name = var.subsetting_policy_subsetting_schema_relation_child_schema_name
	parent_columns = var.subsetting_policy_subsetting_schema_relation_parent_columns
	parent_object_name = oci_objectstorage_object.test_object.name
	parent_schema_name = var.subsetting_policy_subsetting_schema_relation_parent_schema_name
	subsetting_policy_id = oci_data_safe_subsetting_policy.test_subsetting_policy.id

	#Optional
	child_object_key = var.subsetting_policy_subsetting_schema_relation_child_object_key
	parent_object_key = var.subsetting_policy_subsetting_schema_relation_parent_object_key
}
```

## Argument Reference

The following arguments are supported:

* `child_columns` - (Required) Unique identifiers identifying the child columns in the relation.
* `child_object_key` - (Optional) The key that identifies the child subsetting table in this relation.
* `child_object_name` - (Required) The name of the child subsetting table
* `child_schema_name` - (Required) The database schema that contains the child subsetting table
* `parent_columns` - (Required) Unique identifiers identifying the parents columns in the relation.
* `parent_object_key` - (Optional) The key that identifies the parent subsetting table in this relation.
* `parent_object_name` - (Required) The name of the parent subsetting table
* `parent_schema_name` - (Required) The database schema that contains the parent subsetting table
* `subsetting_policy_id` - (Required) The OCID of the subsetting policy.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `child_columns` - Unique identifiers identifying the child columns in the relation.
* `child_object_key` - The key that identifies the child subsetting table in this relation.
* `child_object_name` - The name of the child subsetting table
* `child_schema_name` - The database schema that contains the child subsetting table
* `key` - The unique key that identifies a relation between subsetting tables. The key is numeric and unique within a subsetting policy.
* `parent_columns` - Unique identifiers identifying the parents columns in the relation.
* `parent_object_key` - The key that identifies the parent subsetting table in this relation.
* `parent_object_name` - The name of the parent subsetting table
* `parent_schema_name` - The database schema that contains the parent subsetting table
* `relation_type` - The type of referential relationship the column has with its parent. NONE indicates that the sensitive column does not have a parent. DB_DEFINED indicates that the relationship is defined in the database dictionary. APP_DEFINED indicates that the relationship is defined at the application level and not in the database dictionary. 
* `time_created` - The date and time the subsetting relation was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 
* `time_updated` - The date and time the subsetting relation was last updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Subsetting Policy Subsetting Schema Relation
	* `update` - (Defaults to 20 minutes), when updating the Subsetting Policy Subsetting Schema Relation
	* `delete` - (Defaults to 20 minutes), when destroying the Subsetting Policy Subsetting Schema Relation


## Import

SubsettingPolicySubsettingSchemaRelations can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_subsetting_policy_subsetting_schema_relation.test_subsetting_policy_subsetting_schema_relation "subsettingPolicies/{subsettingPolicyId}/subsettingSchemaRelations/{subsettingSchemaRelationKey}" 
```

