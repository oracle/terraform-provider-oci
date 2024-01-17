---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_masking_policies_masking_columns"
sidebar_current: "docs-oci-datasource-data_safe-masking_policies_masking_columns"
description: |-
  Provides the list of Masking Policies Masking Columns in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_masking_policies_masking_columns
This data source provides the list of Masking Policies Masking Columns in Oracle Cloud Infrastructure Data Safe service.

Gets a list of masking columns present in the specified masking policy and based on the specified query parameters. 


## Example Usage

```hcl
data "oci_data_safe_masking_policies_masking_columns" "test_masking_policies_masking_columns" {
	#Required
	masking_policy_id = oci_data_safe_masking_policy.test_masking_policy.id

	#Optional
	column_name = var.masking_policies_masking_column_column_name
	data_type = var.masking_policies_masking_column_data_type
	is_masking_enabled = var.masking_policies_masking_column_is_masking_enabled
	is_seed_required = var.masking_policies_masking_column_is_seed_required
	masking_column_group = var.masking_policies_masking_column_masking_column_group
	masking_column_lifecycle_state = var.masking_policies_masking_column_masking_column_lifecycle_state
	object = var.masking_policies_masking_column_object
	object_type = var.masking_policies_masking_column_object_type
	schema_name = var.masking_policies_masking_column_schema_name
	sensitive_type_id = oci_data_safe_sensitive_type.test_sensitive_type.id
	time_created_greater_than_or_equal_to = var.masking_policies_masking_column_time_created_greater_than_or_equal_to
	time_created_less_than = var.masking_policies_masking_column_time_created_less_than
	time_updated_greater_than_or_equal_to = var.masking_policies_masking_column_time_updated_greater_than_or_equal_to
	time_updated_less_than = var.masking_policies_masking_column_time_updated_less_than
}
```

## Argument Reference

The following arguments are supported:

* `column_name` - (Optional) A filter to return only a specific column based on column name.
* `data_type` - (Optional) A filter to return only resources that match the specified data types.
* `is_masking_enabled` - (Optional) A filter to return the masking column resources based on the value of their isMaskingEnabled attribute. A value of true returns only those columns for which masking is enabled. A value of false returns only those columns for which masking is disabled. Omitting this parameter returns all the masking columns in a masking policy. 
* `is_seed_required` - (Optional) A filter to return masking columns based on whether the assigned masking formats need a seed value for masking. A value of true returns those masking columns that are using Deterministic Encryption or Deterministic Substitution masking format. 
* `masking_column_group` - (Optional) A filter to return only the resources that match the specified masking column group.
* `masking_column_lifecycle_state` - (Optional) A filter to return only the resources that match the specified lifecycle states.
* `masking_policy_id` - (Required) The OCID of the masking policy.
* `object` - (Optional) A filter to return only items related to a specific object name.
* `object_type` - (Optional) A filter to return only items related to a specific object type.
* `schema_name` - (Optional) A filter to return only items related to specific schema name.
* `sensitive_type_id` - (Optional) A filter to return only items related to a specific sensitive type OCID.
* `time_created_greater_than_or_equal_to` - (Optional) A filter to return only the resources that were created after the specified date and time, as defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Using TimeCreatedGreaterThanOrEqualToQueryParam parameter retrieves all resources created after that date.

	**Example:** 2016-12-19T16:39:57.600Z 
* `time_created_less_than` - (Optional) Search for resources that were created before a specific date. Specifying this parameter corresponding `timeCreatedLessThan` parameter will retrieve all resources created before the specified created date, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by RFC 3339.

	**Example:** 2016-12-19T16:39:57.600Z 
* `time_updated_greater_than_or_equal_to` - (Optional) Search for resources that were updated after a specific date. Specifying this parameter corresponding `timeUpdatedGreaterThanOrEqualTo` parameter will retrieve all resources updated after the specified created date, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by RFC 3339. 
* `time_updated_less_than` - (Optional) Search for resources that were updated before a specific date. Specifying this parameter corresponding `timeUpdatedLessThan` parameter will retrieve all resources updated before the specified created date, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by RFC 3339. 


## Attributes Reference

The following attributes are exported:

* `masking_column_collection` - The list of masking_column_collection.

### MaskingPoliciesMaskingColumn Reference

The following attributes are exported:

* `child_columns` - An array of child columns that are in referential relationship with the masking column.
* `column_name` - The name of the database column. Note that the same name is used for the masking column.  There is no separate displayName attribute for the masking column.  
* `data_type` - The data type of the masking column.
* `is_masking_enabled` - Indicates whether data masking is enabled for the masking column.
* `key` - The unique key that identifies the masking column. It's numeric and unique within a masking policy.
* `lifecycle_details` - Details about the current state of the masking column.
* `masking_column_group` - The group of the masking column. All the columns in a group are masked together to ensure  that the masked data across these columns continue to retain the same logical relationship.  For more details, check <a href=https://docs.oracle.com/en/cloud/paas/data-safe/udscs/group-masking1.html#GUID-755056B9-9540-48C0-9491-262A44A85037>Group Masking in the Data Safe documentation.</a>  
* `masking_formats` - An array of masking formats assigned to the masking column.
	* `condition` - A condition that must be true for applying the masking format. It can be any valid  SQL construct that can be used in a SQL predicate. It enables you to do  <a href="https://docs.oracle.com/en/cloud/paas/data-safe/udscs/conditional-masking.html">conditional masking</a>  so that you can mask the column data values differently using different masking  formats and the associated conditions. 
	* `description` - The description of the masking format.
	* `format_entries` - An array of format entries. The combined output of all the format entries is  used for masking the column data values. 
		* `column_name` - The name of the substitution column.
		* `description` - The description of the format entry.
		* `end_date` - The upper bound of the range within which all the original column values fall. The end date must be greater than or equal to the start date.  
		* `end_length` - The maximum number of characters the generated strings should have. It can  be any integer greater than zero, but it must be greater than or equal to  the start length.  
		* `end_value` - The upper bound of the range within which random decimal numbers should be generated. It must be greater than or equal to the start value. It supports  input of double type.  
		* `fixed_number` - The constant number to be used for masking.
		* `fixed_string` - The constant string to be used for masking.
		* `grouping_columns` - One or more reference columns to be used to group column values so that they can be shuffled within their own group. The grouping columns and  the column to be masked must belong to the same table.  
		* `length` - The number of characters that should be there in the substring. It should be an integer and greater than zero.  
		* `library_masking_format_id` - The OCID of the library masking format.
		* `pattern` - The pattern that should be used to mask data.
		* `post_processing_function` - The post processing function in SCHEMA_NAME.PACKAGE_NAME.FUNCTION_NAME format. It can be a standalone or packaged function, so PACKAGE_NAME is optional.  
		* `random_list` - A comma-separated list of values to be used to replace column values. The list can be of strings, numbers, or dates. The data type of each value in the list must be compatible with the data type of the column. The number of entries in the list cannot be more than 999.  
		* `regular_expression` - The regular expression to be used for masking. For data with characters in the ASCII character set, providing a regular expression is optional. However, it  is required if the data contains multi-byte characters. If not provided, an  error is returned when a multi-byte character is found.

			In the case of ASCII characters, if a regular expression is not provided,  Deterministic Encryption can encrypt variable-length column values while  preserving their original format.

			If a regular expression is provided, the column values in all the rows must match  the regular expression. Deterministic Encryption supports a subset of the regular  expression language. It supports encryption of fixed-length strings, and does not  support * or + syntax of regular expressions. The encrypted values also match the  regular expression, which helps to ensure that the original format is preserved.  If an original value does not match the regular expression, Deterministic Encryption  might not produce a one-to-one mapping. All non-confirming values are mapped to a  single encrypted value, thereby producing a many-to-one mapping.  
		* `replace_with` - The value that should be used to replace the data matching the regular  expression. It can be a fixed string, fixed number, null value, or  SQL expression.  
		* `schema_name` - The name of the schema that contains the substitution column.
		* `sql_expression` - The SQL expression to be used to generate the masked values. It can  consist of one or more values, operators, and SQL functions that  evaluate to a value. It can also contain substitution columns from  the same table. Specify the substitution columns within percent (%)  symbols.  
		* `start_date` - The lower bound of the range within which all the original column values fall. The start date must be less than or equal to the end date.  
		* `start_length` - The minimum number of characters the generated strings should have. It can  be any integer greater than zero, but it must be less than or equal to the  end length.  
		* `start_position` - The starting position in the original string from where the substring should be extracted. It can be either a positive or a negative integer. If It's negative, the counting starts from the end of the string.  
		* `start_value` - The lower bound of the range within which random decimal numbers should  be generated. It must be less than or equal to the end value. It supports  input of double type. 
		* `table_name` - The name of the table that contains the substitution column.
		* `type` - The type of the format entry.
		* `user_defined_function` - The user-defined function in SCHEMA_NAME.PACKAGE_NAME.FUNCTION_NAME format.  It can be a standalone or packaged function, so PACKAGE_NAME is optional.  
* `masking_policy_id` - The OCID of the masking policy that contains the masking column.
* `object` - The name of the object (table or editioning view) that contains the database column.
* `object_type` - The type of the object that contains the database column.
* `schema_name` - The name of the schema that contains the database column.
* `sensitive_type_id` - The OCID of the sensitive type associated with the masking column.
* `state` - The current state of the masking column.
* `time_created` - The date and time the masking column was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  
* `time_updated` - The date and time the masking column was last updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  

