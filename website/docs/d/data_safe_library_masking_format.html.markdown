---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_library_masking_format"
sidebar_current: "docs-oci-datasource-data_safe-library_masking_format"
description: |-
  Provides details about a specific Library Masking Format in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_library_masking_format
This data source provides details about a specific Library Masking Format resource in Oracle Cloud Infrastructure Data Safe service.

Gets the details of the specified library masking format.

## Example Usage

```hcl
data "oci_data_safe_library_masking_format" "test_library_masking_format" {
	#Required
	library_masking_format_id = oci_data_safe_library_masking_format.test_library_masking_format.id
}
```

## Argument Reference

The following arguments are supported:

* `library_masking_format_id` - (Required) The OCID of the library masking format.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the library masking format.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the library masking format.
* `display_name` - The display name of the library masking format.
* `format_entries` - An array of format entries. The combined output of all the format entries is used for masking.
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
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the library masking format.
* `sensitive_type_ids` - An array of OCIDs of the sensitive types compatible with the library masking format.
* `source` - Specifies whether the library masking format is user-defined or predefined.
* `state` - The current state of the library masking format.
* `time_created` - The date and time the library masking format was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339)
* `time_updated` - The date and time the library masking format was updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339)

