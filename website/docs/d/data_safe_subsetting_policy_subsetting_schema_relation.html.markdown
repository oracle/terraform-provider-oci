---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_subsetting_policy_subsetting_schema_relation"
sidebar_current: "docs-oci-datasource-data_safe-subsetting_policy_subsetting_schema_relation"
description: |-
  Provides details about a specific Subsetting Policy Subsetting Schema Relation in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_subsetting_policy_subsetting_schema_relation
This data source provides details about a specific Subsetting Policy Subsetting Schema Relation resource in Oracle Cloud Infrastructure Data Safe service.

Gets the details of the specified referential relation in the subsetting policy.

## Example Usage

```hcl
data "oci_data_safe_subsetting_policy_subsetting_schema_relation" "test_subsetting_policy_subsetting_schema_relation" {
	#Required
	subsetting_policy_id = oci_data_safe_subsetting_policy.test_subsetting_policy.id
	subsetting_schema_relation_key = var.subsetting_policy_subsetting_schema_relation_subsetting_schema_relation_key
}
```

## Argument Reference

The following arguments are supported:

* `subsetting_policy_id` - (Required) The OCID of the subsetting policy.
* `subsetting_schema_relation_key` - (Required) The unique key that identifies the subsetting relation. It's numeric and unique within a subsetting policy.


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

