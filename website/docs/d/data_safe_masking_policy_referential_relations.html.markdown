---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_masking_policy_referential_relations"
sidebar_current: "docs-oci-datasource-data_safe-masking_policy_referential_relations"
description: |-
  Provides the list of Masking Policy Referential Relations in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_masking_policy_referential_relations
This data source provides the list of Masking Policy Referential Relations in Oracle Cloud Infrastructure Data Safe service.

Gets a list of referential relations present in the specified masking policy based on the specified query parameters.


## Example Usage

```hcl
data "oci_data_safe_masking_policy_referential_relations" "test_masking_policy_referential_relations" {
	#Required
	masking_policy_id = oci_data_safe_masking_policy.test_masking_policy.id

	#Optional
	column_name = var.masking_policy_referential_relation_column_name
	object = var.masking_policy_referential_relation_object
	relation_type = var.masking_policy_referential_relation_relation_type
	schema_name = var.masking_policy_referential_relation_schema_name
}
```

## Argument Reference

The following arguments are supported:

* `column_name` - (Optional) A filter to return only a specific column based on column name.
* `masking_policy_id` - (Required) The OCID of the masking policy.
* `object` - (Optional) A filter to return only items related to a specific object name.
* `relation_type` - (Optional) A filter to return columns based on their relationship with their parent columns. If set to NONE, it returns the columns that do not have any parent. The response includes the parent columns as well as the independent columns that are not in any relationship. If set to APP_DEFINED, it returns all the child columns that have application-level (non-dictionary) relationship with their parents. If set to DB_DEFINED, it returns all the child columns that have database-level (dictionary-defined) relationship with their parents. 
* `schema_name` - (Optional) A filter to return only items related to specific schema name.


## Attributes Reference

The following attributes are exported:

* `masking_policy_referential_relation_collection` - The list of masking_policy_referential_relation_collection.

### MaskingPolicyReferentialRelation Reference

The following attributes are exported:

* `items` - An array of referential relation summary objects.
	* `child` - maskingPolicyColumnsInfo object has details of column group with schema details.
		* `object` - The name of the object (table or editioning view) that contains the database column(s).
		* `object_type` - The type of the database object that contains the masking policy.
		* `referential_column_group` - Group of columns in referential relation. Order needs to be maintained in the elements of the parent/child array listing. 
		* `schema_name` - The name of the schema that contains the database column(s).
	* `masking_format` - The masking format associated with the parent column.
	* `masking_policy_id` - The OCID of the masking policy that contains the column.
	* `parent` - maskingPolicyColumnsInfo object has details of column group with schema details.
		* `object` - The name of the object (table or editioning view) that contains the database column(s).
		* `object_type` - The type of the database object that contains the masking policy.
		* `referential_column_group` - Group of columns in referential relation. Order needs to be maintained in the elements of the parent/child array listing. 
		* `schema_name` - The name of the schema that contains the database column(s).
	* `relation_type` - The type of referential relationship the column has with its parent. DB_DEFINED indicates that the relationship is defined in the database dictionary. APP_DEFINED indicates that the relationship is defined at the application level and not in the database dictionary. 

