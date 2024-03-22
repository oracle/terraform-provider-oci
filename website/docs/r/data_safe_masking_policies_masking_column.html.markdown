---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_masking_policies_masking_column"
sidebar_current: "docs-oci-resource-data_safe-masking_policies_masking_column"
description: |-
  Provides the Masking Policies Masking Column resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_masking_policies_masking_column
This resource provides the Masking Policies Masking Column resource in Oracle Cloud Infrastructure Data Safe service.

Creates a new masking column in the specified masking policy. Use this operation
to add parent columns only. It automatically adds the child columns from the
associated sensitive data model or target database. If you provide the
sensitiveTypeId attribute but not the maskingFormats attribute, it automatically
assigns the default masking format associated with the specified sensitive type.
Alternatively, if you provide the maskingFormats attribute, the specified masking
formats are assigned to the column.

Using the maskingFormats attribute, you can assign one or more masking formats
to a column. You need to specify a condition as part of each masking format. It
enables you to do <a href="https://docs.oracle.com/en/cloud/paas/data-safe/udscs/conditional-masking.html">conditional masking</a>
so that you can mask the column data values differently using different
masking conditions. A masking format can have one or more format entries. The
combined output of all the format entries is used for masking. It provides the
flexibility to define a masking format that can generate different parts of a data
value separately and then combine them to get the final data value for masking.

You can use the maskingColumnGroup attribute to group the columns that you would
like to mask together. It enables you to do <a href="https://docs.oracle.com/en/cloud/paas/data-safe/udscs/group-masking1.html#GUID-755056B9-9540-48C0-9491-262A44A85037">group or compound masking</a> that ensures that the
masked data across the columns in a group continue to retain the same logical relationship.


## Example Usage

```hcl
resource "oci_data_safe_masking_policies_masking_column" "test_masking_policies_masking_column" {
	#Required
	column_name = var.masking_policies_masking_column_column_name
	masking_policy_id = oci_data_safe_masking_policy.test_masking_policy.id
	object = var.masking_policies_masking_column_object
	schema_name = var.masking_policies_masking_column_schema_name

	#Optional
	is_masking_enabled = var.masking_policies_masking_column_is_masking_enabled
	masking_column_group = var.masking_policies_masking_column_masking_column_group
	masking_formats {
		#Required
		format_entries {
			#Required
			type = var.masking_policies_masking_column_masking_formats_format_entries_type

			#Optional
			column_name = var.masking_policies_masking_column_masking_formats_format_entries_column_name
			description = var.masking_policies_masking_column_masking_formats_format_entries_description
			end_date = var.masking_policies_masking_column_masking_formats_format_entries_end_date
			end_length = var.masking_policies_masking_column_masking_formats_format_entries_end_length
			end_value = var.masking_policies_masking_column_masking_formats_format_entries_end_value
			fixed_number = var.masking_policies_masking_column_masking_formats_format_entries_fixed_number
			fixed_string = var.masking_policies_masking_column_masking_formats_format_entries_fixed_string
			grouping_columns = var.masking_policies_masking_column_masking_formats_format_entries_grouping_columns
			length = var.masking_policies_masking_column_masking_formats_format_entries_length
			library_masking_format_id = oci_data_safe_library_masking_format.test_library_masking_format.id
			pattern = var.masking_policies_masking_column_masking_formats_format_entries_pattern
			post_processing_function = var.masking_policies_masking_column_masking_formats_format_entries_post_processing_function
			random_list = var.masking_policies_masking_column_masking_formats_format_entries_random_list
			regular_expression = var.masking_policies_masking_column_masking_formats_format_entries_regular_expression
			replace_with = var.masking_policies_masking_column_masking_formats_format_entries_replace_with
			schema_name = var.masking_policies_masking_column_masking_formats_format_entries_schema_name
			sql_expression = var.masking_policies_masking_column_masking_formats_format_entries_sql_expression
			start_date = var.masking_policies_masking_column_masking_formats_format_entries_start_date
			start_length = var.masking_policies_masking_column_masking_formats_format_entries_start_length
			start_position = var.masking_policies_masking_column_masking_formats_format_entries_start_position
			start_value = var.masking_policies_masking_column_masking_formats_format_entries_start_value
			table_name = oci_nosql_table.test_table.name
			user_defined_function = var.masking_policies_masking_column_masking_formats_format_entries_user_defined_function
		}

		#Optional
		condition = var.masking_policies_masking_column_masking_formats_condition
		description = var.masking_policies_masking_column_masking_formats_description
	}
	object_type = var.masking_policies_masking_column_object_type
	sensitive_type_id = oci_data_safe_sensitive_type.test_sensitive_type.id
}
```

## Argument Reference

The following arguments are supported:

* `column_name` - (Required) The name of the database column. This attribute cannot be updated for an existing  masking column. Note that the same name is used for the masking column. There  is no separate displayName attribute for the masking column.  
* `is_masking_enabled` - (Optional) (Updatable) Indicates whether data masking is enabled for the masking column. Set it to false if  you don't want to mask the column.  
* `masking_column_group` - (Optional) (Updatable) The group of the masking column. It's a masking group identifier and can be any string  of acceptable length. All the columns in a group are masked together to ensure that  the masked data across these columns continue to retain the same logical relationship.  For more details, check  <a href=https://docs.oracle.com/en/cloud/paas/data-safe/udscs/group-masking1.html#GUID-755056B9-9540-48C0-9491-262A44A85037>Group Masking in the Data Safe documentation.</a>  
* `masking_formats` - (Optional) (Updatable) The masking formats to be assigned to the masking column. You can specify a condition  as part of each masking format. It enables you to do  <a href="https://docs.oracle.com/en/cloud/paas/data-safe/udscs/conditional-masking.html">conditional masking</a>  so that you can mask the column data values differently using different masking  formats and the associated conditions. A masking format can have one or more format  entries. The combined output of all the format entries is used for masking. It  provides the flexibility to define a masking format that can generate different parts  of a data value separately and then combine them to get the final data value for masking.  
	* `condition` - (Optional) (Updatable) A condition that must be true for applying the masking format. It can be any valid  SQL construct that can be used in a SQL predicate. It enables you to do  <a href="https://docs.oracle.com/en/cloud/paas/data-safe/udscs/conditional-masking.html">conditional masking</a>  so that you can mask the column data values differently using different masking  formats and the associated conditions. 
	* `description` - (Optional) (Updatable) The description of the masking format.
	* `format_entries` - (Required) (Updatable) An array of format entries. The combined output of all the format entries is  used for masking the column data values. 
		* `column_name` - (Required when type=DETERMINISTIC_SUBSTITUTION | RANDOM_SUBSTITUTION) (Updatable) The name of the substitution column.
		* `description` - (Optional) (Updatable) The description of the format entry.
		* `end_date` - (Required when type=DETERMINISTIC_ENCRYPTION_DATE | RANDOM_DATE) (Updatable) The upper bound of the range within which all the original column values fall. The end date must be greater than or equal to the start date.  
		* `end_length` - (Required when type=RANDOM_DIGITS | RANDOM_STRING) (Updatable) The maximum number of characters the generated strings should have. It can  be any integer greater than zero, but it must be greater than or equal to  the start length.  
		* `end_value` - (Required when type=RANDOM_DECIMAL_NUMBER | RANDOM_NUMBER) (Updatable) The upper bound of the range within which random decimal numbers should be generated. It must be greater than or equal to the start value. It supports  input of double type.  
		* `fixed_number` - (Required when type=FIXED_NUMBER) (Updatable) The constant number to be used for masking.
		* `fixed_string` - (Required when type=FIXED_STRING) (Updatable) The constant string to be used for masking.
		* `grouping_columns` - (Applicable when type=SHUFFLE) (Updatable) One or more reference columns to be used to group column values so that they can be shuffled within their own group. The grouping columns and  the column to be masked must belong to the same table.  
		* `length` - (Required when type=SUBSTRING) (Updatable) The number of characters that should be there in the substring. It should be an integer and greater than zero.  
		* `library_masking_format_id` - (Required when type=LIBRARY_MASKING_FORMAT) (Updatable) The OCID of the library masking format.
		* `pattern` - (Required when type=PATTERN) (Updatable) The pattern that should be used to mask data.
		* `post_processing_function` - (Required when type=POST_PROCESSING_FUNCTION) (Updatable) The post processing function in SCHEMA_NAME.PACKAGE_NAME.FUNCTION_NAME format. It can be a standalone or packaged function, so PACKAGE_NAME is optional.  
		* `random_list` - (Required when type=RANDOM_LIST) (Updatable) A comma-separated list of values to be used to replace column values. The list can be of strings, numbers, or dates. The data type of each value in the list must be compatible with the data type of the column. The number of entries in the list cannot be more than 999.  
		* `regular_expression` - (Required when type=DETERMINISTIC_ENCRYPTION | REGULAR_EXPRESSION) (Updatable) The regular expression to be used for masking. For data with characters in the ASCII character set, providing a regular expression is optional. However, it  is required if the data contains multi-byte characters. If not provided, an  error is returned when a multi-byte character is found.

			In the case of ASCII characters, if a regular expression is not provided,  Deterministic Encryption can encrypt variable-length column values while  preserving their original format.

			If a regular expression is provided, the column values in all the rows must match  the regular expression. Deterministic Encryption supports a subset of the regular  expression language. It supports encryption of fixed-length strings, and does not  support * or + syntax of regular expressions. The encrypted values also match the  regular expression, which helps to ensure that the original format is preserved.  If an original value does not match the regular expression, Deterministic Encryption  might not produce a one-to-one mapping. All non-confirming values are mapped to a  single encrypted value, thereby producing a many-to-one mapping.  
		* `replace_with` - (Required when type=REGULAR_EXPRESSION) (Updatable) The value that should be used to replace the data matching the regular  expression. It can be a fixed string, fixed number, null value, or  SQL expression.  
		* `schema_name` - (Required when type=DETERMINISTIC_SUBSTITUTION | RANDOM_SUBSTITUTION) (Updatable) The name of the schema that contains the substitution column.
		* `sql_expression` - (Required when type=SQL_EXPRESSION) (Updatable) The SQL expression to be used to generate the masked values. It can  consist of one or more values, operators, and SQL functions that  evaluate to a value. It can also contain substitution columns from  the same table. Specify the substitution columns within percent (%)  symbols.  
		* `start_date` - (Required when type=DETERMINISTIC_ENCRYPTION_DATE | RANDOM_DATE) (Updatable) The lower bound of the range within which all the original column values fall. The start date must be less than or equal to the end date.  
		* `start_length` - (Required when type=RANDOM_DIGITS | RANDOM_STRING) (Updatable) The minimum number of characters the generated strings should have. It can  be any integer greater than zero, but it must be less than or equal to the  end length.  
		* `start_position` - (Required when type=SUBSTRING) (Updatable) The starting position in the original string from where the substring should be extracted. It can be either a positive or a negative integer. If It's negative, the counting starts from the end of the string.  
		* `start_value` - (Required when type=RANDOM_DECIMAL_NUMBER | RANDOM_NUMBER) (Updatable) The lower bound of the range within which random decimal numbers should  be generated. It must be less than or equal to the end value. It supports  input of double type. 
		* `table_name` - (Required when type=DETERMINISTIC_SUBSTITUTION | RANDOM_SUBSTITUTION) (Updatable) The name of the table that contains the substitution column.
		* `type` - (Required) (Updatable) The type of the format entry.
		* `user_defined_function` - (Required when type=USER_DEFINED_FUNCTION) (Updatable) The user-defined function in SCHEMA_NAME.PACKAGE_NAME.FUNCTION_NAME format.  It can be a standalone or packaged function, so PACKAGE_NAME is optional.  
* `masking_policy_id` - (Required) The OCID of the masking policy.
* `object` - (Required) The name of the object (table or editioning view) that contains the database column. This attribute cannot be updated for an existing masking column.
* `object_type` - (Optional) (Updatable) The type of the object that contains the database column.
* `schema_name` - (Required) The name of the schema that contains the database column. This attribute cannot be updated for an existing masking column.
* `sensitive_type_id` - (Optional) (Updatable) The OCID of the sensitive type to be associated with the masking column. Note that  if the maskingFormats attribute isn't provided while creating a masking column,   the default masking format associated with the specified sensitive type is assigned  to the masking column.  


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Masking Policies Masking Column
	* `update` - (Defaults to 20 minutes), when updating the Masking Policies Masking Column
	* `delete` - (Defaults to 20 minutes), when destroying the Masking Policies Masking Column


## Import

MaskingPoliciesMaskingColumns can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_masking_policies_masking_column.test_masking_policies_masking_column "maskingPolicies/{maskingPolicyId}/maskingColumns/{maskingColumnKey}" 
```

