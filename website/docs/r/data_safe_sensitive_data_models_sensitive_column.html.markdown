---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_sensitive_data_models_sensitive_column"
sidebar_current: "docs-oci-resource-data_safe-sensitive_data_models_sensitive_column"
description: |-
  Provides the Sensitive Data Models Sensitive Column resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_sensitive_data_models_sensitive_column
This resource provides the Sensitive Data Models Sensitive Column resource in Oracle Cloud Infrastructure Data Safe service.

Creates a new sensitive column in the specified sensitive data model.


## Example Usage

```hcl
resource "oci_data_safe_sensitive_data_models_sensitive_column" "test_sensitive_data_models_sensitive_column" {
	#Required
	column_name = var.sensitive_data_models_sensitive_column_column_name
	object = var.sensitive_data_models_sensitive_column_object
	schema_name = var.sensitive_data_models_sensitive_column_schema_name
	sensitive_data_model_id = oci_data_safe_sensitive_data_model.test_sensitive_data_model.id

	#Optional
	app_defined_child_column_keys = var.sensitive_data_models_sensitive_column_app_defined_child_column_keys
	app_name = var.sensitive_data_models_sensitive_column_app_name
	data_type = var.sensitive_data_models_sensitive_column_data_type
	db_defined_child_column_keys = var.sensitive_data_models_sensitive_column_db_defined_child_column_keys
	object_type = var.sensitive_data_models_sensitive_column_object_type
	parent_column_keys = var.sensitive_data_models_sensitive_column_parent_column_keys
	relation_type = var.sensitive_data_models_sensitive_column_relation_type
	sensitive_type_id = oci_data_safe_sensitive_type.test_sensitive_type.id
	status = var.sensitive_data_models_sensitive_column_status
}
```

## Argument Reference

The following arguments are supported:

* `app_defined_child_column_keys` - (Optional) (Updatable) Unique keys identifying the columns that are application-level (non-dictionary) children of the sensitive column. This attribute can be used to establish relationship between columns in a sensitive data model. Note that the child columns must be added to the sensitive data model before their keys can be specified here. If this attribute is provided, the parentColumnKeys and relationType attributes of the child columns are automatically updated to reflect the relationship. 
* `app_name` - (Optional) The name of the application associated with the sensitive column. It's useful when the application name is different from the schema name. Otherwise, it can be ignored. If this attribute is not provided, it's automatically populated with the value provided for the schemaName attribute. 
* `column_name` - (Required) The name of the sensitive column.
* `data_type` - (Optional) (Updatable) The data type of the sensitive column.
* `db_defined_child_column_keys` - (Optional) (Updatable) Unique keys identifying the columns that are database-level (dictionary-defined) children of the sensitive column. This attribute can be used to establish relationship between columns in a sensitive data model. Note that the child columns must be added to the sensitive data model before their keys can be specified here. If this attribute is provided, the parentColumnKeys and relationType attributes of the child columns are automatically updated to reflect the relationship. 
* `object` - (Required) The database object that contains the sensitive column.
* `object_type` - (Optional) The type of the database object that contains the sensitive column.
* `parent_column_keys` - (Optional) (Updatable) Unique keys identifying the columns that are parents of the sensitive column. At present, it accepts only one parent column key. This attribute can be used to establish relationship between columns in a sensitive data model. Note that the parent column must be added to the sensitive data model before its key can be specified here. If this attribute is provided, the appDefinedChildColumnKeys or dbDefinedChildColumnKeys attribute of the parent column is automatically updated to reflect the relationship. 
* `relation_type` - (Optional) (Updatable) The type of referential relationship the sensitive column has with its parent. NONE indicates that the sensitive column does not have a parent. DB_DEFINED indicates that the relationship is defined in the database dictionary. APP_DEFINED indicates that the relationship is defined at the application level and not in the database dictionary. 
* `schema_name` - (Required) The database schema that contains the sensitive column.
* `sensitive_data_model_id` - (Required) The OCID of the sensitive data model.
* `sensitive_type_id` - (Optional) (Updatable) The OCID of the sensitive type to be associated with the sensitive column.
* `status` - (Optional) (Updatable) The status of the sensitive column. VALID means the column is considered sensitive. INVALID means the column is not considered sensitive. Tracking invalid columns in a sensitive data model helps ensure that an incremental data discovery job does not identify these columns as sensitive. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Sensitive Data Models Sensitive Column
	* `update` - (Defaults to 20 minutes), when updating the Sensitive Data Models Sensitive Column
	* `delete` - (Defaults to 20 minutes), when destroying the Sensitive Data Models Sensitive Column


## Import

SensitiveDataModelsSensitiveColumns can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_sensitive_data_models_sensitive_column.test_sensitive_data_models_sensitive_column "sensitiveDataModels/{sensitiveDataModelId}/sensitiveColumns/{sensitiveColumnKey}" 
```

