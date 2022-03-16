---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_sensitive_data_models_sensitive_column"
sidebar_current: "docs-oci-datasource-data_safe-sensitive_data_models_sensitive_column"
description: |-
  Provides details about a specific Sensitive Data Models Sensitive Column in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_sensitive_data_models_sensitive_column
This data source provides details about a specific Sensitive Data Models Sensitive Column resource in Oracle Cloud Infrastructure Data Safe service.

Gets the details of the specified sensitive column.

## Example Usage

```hcl
data "oci_data_safe_sensitive_data_models_sensitive_column" "test_sensitive_data_models_sensitive_column" {
	#Required
	sensitive_column_key = var.sensitive_data_models_sensitive_column_sensitive_column_key
	sensitive_data_model_id = oci_data_safe_sensitive_data_model.test_sensitive_data_model.id
}
```

## Argument Reference

The following arguments are supported:

* `sensitive_column_key` - (Required) The unique key that identifies the sensitive column. It's numeric and unique within a sensitive data model.
* `sensitive_data_model_id` - (Required) The OCID of the sensitive data model.


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

