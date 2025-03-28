---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_sensitive_data_model_referential_relations"
sidebar_current: "docs-oci-datasource-data_safe-sensitive_data_model_referential_relations"
description: |-
  Provides the list of Sensitive Data Model Referential Relations in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_sensitive_data_model_referential_relations
This data source provides the list of Sensitive Data Model Referential Relations in Oracle Cloud Infrastructure Data Safe service.

Gets a list of referential relations present in the specified sensitive data model based on the specified query parameters.


## Example Usage

```hcl
data "oci_data_safe_sensitive_data_model_referential_relations" "test_sensitive_data_model_referential_relations" {
	#Required
	sensitive_data_model_id = oci_data_safe_sensitive_data_model.test_sensitive_data_model.id

	#Optional
	column_name = var.sensitive_data_model_referential_relation_column_name
	is_sensitive = var.sensitive_data_model_referential_relation_is_sensitive
	object = var.sensitive_data_model_referential_relation_object
	relation_type = var.sensitive_data_model_referential_relation_relation_type
	schema_name = var.sensitive_data_model_referential_relation_schema_name
}
```

## Argument Reference

The following arguments are supported:

* `column_name` - (Optional) A filter to return only a specific column based on column name.
* `is_sensitive` - (Optional) Returns referential relations containing sensitive columns when true. Returns referential relations containing non sensitive columns when false. 
* `object` - (Optional) A filter to return only items related to a specific object name.
* `relation_type` - (Optional) A filter to return sensitive columns based on their relationship with their parent columns. If set to NONE, it returns the sensitive columns that do not have any parent. The response includes the parent columns as well as the independent columns that are not in any relationship. If set to APP_DEFINED, it returns all the child columns that have application-level (non-dictionary) relationship with their parents. If set to DB_DEFINED, it returns all the child columns that have database-level (dictionary-defined) relationship with their parents. 
* `schema_name` - (Optional) A filter to return only items related to specific schema name.
* `sensitive_data_model_id` - (Required) The OCID of the sensitive data model.


## Attributes Reference

The following attributes are exported:

* `referential_relation_collection` - The list of referential_relation_collection.

### SensitiveDataModelReferentialRelation Reference

The following attributes are exported:

* `child` - columnsInfo object has details of column group with schema details.
	* `app_name` - The application name.
	* `column_group` - Group of columns in referential relation. Order needs to be maintained in the elements of the parent/child array listing. 
	* `object` - The database object that contains the columns.
	* `object_type` - The type of the database object that contains the sensitive column.
	* `schema_name` - The schema name.
	* `sensitive_type_ids` - Sensitive type ocids of each column groups. Order needs to be maintained with the parent column group. For the DB defined referential relations identified during SDM creation, we cannot add sensitive types.  Instead use the sensitiveColumn POST API to mark the columns sensitive. 
* `is_sensitive` - Determines if the columns present in the referential relation is present in the sensitive data model
* `key` - The unique key that identifies the referential relation. It's numeric and unique within a sensitive data model.
* `parent` - columnsInfo object has details of column group with schema details.
	* `app_name` - The application name.
	* `column_group` - Group of columns in referential relation. Order needs to be maintained in the elements of the parent/child array listing. 
	* `object` - The database object that contains the columns.
	* `object_type` - The type of the database object that contains the sensitive column.
	* `schema_name` - The schema name.
	* `sensitive_type_ids` - Sensitive type ocids of each column groups. Order needs to be maintained with the parent column group. For the DB defined referential relations identified during SDM creation, we cannot add sensitive types.  Instead use the sensitiveColumn POST API to mark the columns sensitive. 
* `relation_type` - The type of referential relationship the sensitive column has with its parent. NONE indicates that the sensitive column does not have a parent. DB_DEFINED indicates that the relationship is defined in the database dictionary. APP_DEFINED indicates that the relationship is defined at the application level and not in the database dictionary. 
* `sensitive_data_model_id` - The OCID of the sensitive data model that contains the sensitive column.
* `state` - The current state of the referential relation.

