---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_subsetting_policy_subsetting_schema_objects"
sidebar_current: "docs-oci-datasource-data_safe-subsetting_policy_subsetting_schema_objects"
description: |-
  Provides the list of Subsetting Policy Subsetting Schema Objects in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_subsetting_policy_subsetting_schema_objects
This data source provides the list of Subsetting Policy Subsetting Schema Objects in Oracle Cloud Infrastructure Data Safe service.

Gets a list of objects/tables present in the specified subsetting policy schemas based on the specified query parameters.


## Example Usage

```hcl
data "oci_data_safe_subsetting_policy_subsetting_schema_objects" "test_subsetting_policy_subsetting_schema_objects" {
	#Required
	subsetting_policy_id = oci_data_safe_subsetting_policy.test_subsetting_policy.id

	#Optional
	object = var.subsetting_policy_subsetting_schema_object_object
	schema_name = var.subsetting_policy_subsetting_schema_object_schema_name
}
```

## Argument Reference

The following arguments are supported:

* `object` - (Optional) A filter to return only items related to a specific object name.
* `schema_name` - (Optional) A filter to return only items related to specific schema name.
* `subsetting_policy_id` - (Required) The OCID of the subsetting policy.


## Attributes Reference

The following attributes are exported:

* `subsetting_schema_object_collection` - The list of subsetting_schema_object_collection.

### SubsettingPolicySubsettingSchemaObject Reference

The following attributes are exported:

* `items` - An array of subsetting table summary objects
	* `initial_row_count` - The initial number of rows in this object/table
	* `is_stats_stale` - Indicates if the table stats are stale. This can be used to judge the accuracy of initialRowCount
	* `key` - The unique key that identifies a subsetting table. The key is numeric and unique within a subsetting policy
	* `object` - The name of the database object
	* `object_type` - The type of the database object that contains the subsetting table
	* `schema_name` - The database schema that contains the subsetting table
	* `time_created` - The date and time the subsetting table was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 
	* `time_updated` - The date and time the subsetting table was last updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 

