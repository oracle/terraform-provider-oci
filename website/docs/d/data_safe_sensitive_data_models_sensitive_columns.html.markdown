---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_sensitive_data_models_sensitive_columns"
sidebar_current: "docs-oci-datasource-data_safe-sensitive_data_models_sensitive_columns"
description: |-
  Provides the list of Sensitive Data Models Sensitive Columns in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_sensitive_data_models_sensitive_columns
This data source provides the list of Sensitive Data Models Sensitive Columns in Oracle Cloud Infrastructure Data Safe service.

Gets a list of sensitive columns present in the specified sensitive data model based on the specified query parameters.


## Example Usage

```hcl
data "oci_data_safe_sensitive_data_models_sensitive_columns" "test_sensitive_data_models_sensitive_columns" {
	#Required
	sensitive_data_model_id = oci_data_safe_sensitive_data_model.test_sensitive_data_model.id

	#Optional
	column_group = var.sensitive_data_models_sensitive_column_column_group
	column_name = var.sensitive_data_models_sensitive_column_column_name
	data_type = var.sensitive_data_models_sensitive_column_data_type
	is_case_in_sensitive = var.sensitive_data_models_sensitive_column_is_case_in_sensitive
	object = var.sensitive_data_models_sensitive_column_object
	object_type = var.sensitive_data_models_sensitive_column_object_type
	parent_column_key = var.sensitive_data_models_sensitive_column_parent_column_key
	relation_type = var.sensitive_data_models_sensitive_column_relation_type
	schema_name = var.sensitive_data_models_sensitive_column_schema_name
	sensitive_column_lifecycle_state = var.sensitive_data_models_sensitive_column_sensitive_column_lifecycle_state
	sensitive_type_id = oci_data_safe_sensitive_type.test_sensitive_type.id
	status = var.sensitive_data_models_sensitive_column_status
	time_created_greater_than_or_equal_to = var.sensitive_data_models_sensitive_column_time_created_greater_than_or_equal_to
	time_created_less_than = var.sensitive_data_models_sensitive_column_time_created_less_than
	time_updated_greater_than_or_equal_to = var.sensitive_data_models_sensitive_column_time_updated_greater_than_or_equal_to
	time_updated_less_than = var.sensitive_data_models_sensitive_column_time_updated_less_than
}
```

## Argument Reference

The following arguments are supported:

* `column_group` - (Optional) A filter to return only the sensitive columns that belong to the specified column group.
* `column_name` - (Optional) A filter to return only a specific column based on column name.
* `data_type` - (Optional) A filter to return only the resources that match the specified data types.
* `is_case_in_sensitive` - (Optional) A boolean flag indicating whether the search should be case-insensitive. The search is case-sensitive by default. Set this parameter to true to do case-insensitive search. 
* `object` - (Optional) A filter to return only items related to a specific object name.
* `object_type` - (Optional) A filter to return only items related to a specific object type.
* `parent_column_key` - (Optional) A filter to return only the sensitive columns that are children of one of the columns identified by the specified keys.
* `relation_type` - (Optional) A filter to return sensitive columns based on their relationship with their parent columns. If set to NONE, it returns the sensitive columns that do not have any parent. The response includes the parent columns as well as the independent columns that are not in any relationship. If set to APP_DEFINED, it returns all the child columns that have application-level (non-dictionary) relationship with their parents. If set to DB_DEFINED, it returns all the child columns that have database-level (dictionary-defined) relationship with their parents. 
* `schema_name` - (Optional) A filter to return only items related to specific schema name.
* `sensitive_column_lifecycle_state` - (Optional) Filters the sensitive column resources with the given lifecycle state values.
* `sensitive_data_model_id` - (Required) The OCID of the sensitive data model.
* `sensitive_type_id` - (Optional) A filter to return only the sensitive columns that are associated with one of the sensitive types identified by the specified OCIDs.
* `status` - (Optional) A filter to return only the sensitive columns that match the specified status.
* `time_created_greater_than_or_equal_to` - (Optional) A filter to return only the resources that were created after the specified date and time, as defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Using TimeCreatedGreaterThanOrEqualToQueryParam parameter retrieves all resources created after that date.

	**Example:** 2016-12-19T16:39:57.600Z 
* `time_created_less_than` - (Optional) Search for resources that were created before a specific date. Specifying this parameter corresponding `timeCreatedLessThan` parameter will retrieve all resources created before the specified created date, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by RFC 3339.

	**Example:** 2016-12-19T16:39:57.600Z 
* `time_updated_greater_than_or_equal_to` - (Optional) Search for resources that were updated after a specific date. Specifying this parameter corresponding `timeUpdatedGreaterThanOrEqualTo` parameter will retrieve all resources updated after the specified created date, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by RFC 3339. 
* `time_updated_less_than` - (Optional) Search for resources that were updated before a specific date. Specifying this parameter corresponding `timeUpdatedLessThan` parameter will retrieve all resources updated before the specified created date, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by RFC 3339. 


## Attributes Reference

The following attributes are exported:

* `sensitive_column_collection` - The list of sensitive_column_collection.

### SensitiveDataModelsSensitiveColumn Reference

The following attributes are exported:

* `app_defined_child_column_keys` - Unique keys identifying the columns that are application-level (non-dictionary) children of the sensitive column.
* `app_name` - The name of the application associated with the sensitive column. It's useful when the application name is different from the schema name. Otherwise, it can be ignored. 
* `column_groups` - The composite key groups to which the sensitive column belongs. If the column is part of a composite key, it's assigned a column group. It helps identify and manage referential relationships that involve composite keys. 
* `column_name` - The name of the sensitive column.
* `data_type` - The data type of the sensitive column.
* `db_defined_child_column_keys` - Unique keys identifying the columns that are database-level (dictionary-defined) children of the sensitive column.
* `estimated_data_value_count` - The estimated number of data values the column has in the associated database.
* `key` - The unique key that identifies the sensitive column. It's numeric and unique within a sensitive data model.
* `lifecycle_details` - Details about the current state of the sensitive column.
* `object` - The database object that contains the sensitive column.
* `object_type` - The type of the database object that contains the sensitive column.
* `parent_column_keys` - Unique keys identifying the columns that are parents of the sensitive column. At present, it tracks a single parent only.
* `relation_type` - The type of referential relationship the sensitive column has with its parent. NONE indicates that the sensitive column does not have a parent. DB_DEFINED indicates that the relationship is defined in the database dictionary. APP_DEFINED indicates that the relationship is defined at the application level and not in the database dictionary. 
* `sample_data_values` - Original data values collected for the sensitive column from the associated database. Sample data helps review the column and ensure that it actually contains sensitive data. Note that sample data is retrieved by a data discovery job only if the isSampleDataCollectionEnabled attribute is set to true. At present, only one data value is collected per sensitive column. 
* `schema_name` - The database schema that contains the sensitive column.
* `sensitive_data_model_id` - The OCID of the sensitive data model that contains the sensitive column.
* `sensitive_type_id` - The OCID of the sensitive type associated with the sensitive column.
* `source` - The source of the sensitive column. DISCOVERY indicates that the column was added to the sensitive data model using a data discovery job. MANUAL indicates that the column was added manually. 
* `state` - The current state of the sensitive column.
* `status` - The status of the sensitive column. VALID means the column is considered sensitive. INVALID means the column is not considered sensitive. Tracking invalid columns in a sensitive data model helps ensure that an incremental data discovery job does not identify these columns as sensitive again. 
* `time_created` - The date and time, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339), the sensitive column was created in the sensitive data model. 
* `time_updated` - The date and time, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339), the sensitive column was last updated in the sensitive data model. 

