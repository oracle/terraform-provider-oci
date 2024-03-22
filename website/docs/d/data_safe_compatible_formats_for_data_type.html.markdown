---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_compatible_formats_for_data_type"
sidebar_current: "docs-oci-datasource-data_safe-compatible_formats_for_data_type"
description: |-
  Provides details about a specific Compatible Formats For Data Type in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_compatible_formats_for_data_type
This data source provides details about a specific Compatible Formats For Data Type resource in Oracle Cloud Infrastructure Data Safe service.

Gets a list of basic masking formats compatible with the supported data types.
The data types are grouped into the following categories -
Character - Includes CHAR, NCHAR, VARCHAR2, and NVARCHAR2
Numeric - Includes NUMBER, FLOAT, RAW, BINARY_FLOAT, and BINARY_DOUBLE
Date - Includes DATE and TIMESTAMP
LOB - Includes BLOB, CLOB, and NCLOB
All - Includes all the supported data types


## Example Usage

```hcl
data "oci_data_safe_compatible_formats_for_data_type" "test_compatible_formats_for_data_type" {
}
```

## Argument Reference

The following arguments are supported:



## Attributes Reference

The following attributes are exported:

* `formats_for_data_type` - An array of lists of basic masking formats compatible with the supported data types.
	* `data_type` - The data type category, which can be one of the following - Character - Includes CHAR, NCHAR, VARCHAR2, and NVARCHAR2 Numeric - Includes NUMBER, FLOAT, RAW, BINARY_FLOAT, and BINARY_DOUBLE Date - Includes DATE and TIMESTAMP LOB - Includes BLOB, CLOB, and NCLOB All - Includes all the supported data types  
	* `masking_formats` - An array of the basic masking formats compatible with the data type category.
		* `description` - The description of the masking format.
		* `id` - The OCID of the masking format.
		* `name` - The name of the masking format.

