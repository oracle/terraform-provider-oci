---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_subsetting_policy_subsetting_schema_relations"
sidebar_current: "docs-oci-datasource-data_safe-subsetting_policy_subsetting_schema_relations"
description: |-
  Provides the list of Subsetting Policy Subsetting Schema Relations in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_subsetting_policy_subsetting_schema_relations
This data source provides the list of Subsetting Policy Subsetting Schema Relations in Oracle Cloud Infrastructure Data Safe service.

Gets a list of referential relations present in the specified subsetting policy schemas based on the specified query parameters.


## Example Usage

```hcl
data "oci_data_safe_subsetting_policy_subsetting_schema_relations" "test_subsetting_policy_subsetting_schema_relations" {
	#Required
	subsetting_policy_id = oci_data_safe_subsetting_policy.test_subsetting_policy.id

	#Optional
	object = var.subsetting_policy_subsetting_schema_relation_object
	relation_type = var.subsetting_policy_subsetting_schema_relation_relation_type
	schema_name = var.subsetting_policy_subsetting_schema_relation_schema_name
}
```

## Argument Reference

The following arguments are supported:

* `object` - (Optional) A filter to return only items related to a specific object name.
* `relation_type` - (Optional) A filter to return columns based on their relationship with their parent columns. If set to APP_DEFINED, it returns all the child columns that have application-level (non-dictionary) relationship with their parents. If set to DB_DEFINED, it returns all the child columns that have database-level (dictionary-defined) relationship with their parents. 
* `schema_name` - (Optional) A filter to return only items related to specific schema name.
* `subsetting_policy_id` - (Required) The OCID of the subsetting policy.


## Attributes Reference

The following attributes are exported:

* `subsetting_schema_relation_collection` - The list of subsetting_schema_relation_collection.

### SubsettingPolicySubsettingSchemaRelation Reference

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

