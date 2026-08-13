---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_subsetting_policy_subsetting_schemas"
sidebar_current: "docs-oci-datasource-data_safe-subsetting_policy_subsetting_schemas"
description: |-
  Provides the list of Subsetting Policy Subsetting Schemas in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_subsetting_policy_subsetting_schemas
This data source provides the list of Subsetting Policy Subsetting Schemas in Oracle Cloud Infrastructure Data Safe service.

Gets a list of subsetting schemas present in the specified subsetting policy and based on the specified query parameters.


## Example Usage

```hcl
data "oci_data_safe_subsetting_policy_subsetting_schemas" "test_subsetting_policy_subsetting_schemas" {
	#Required
	subsetting_policy_id = oci_data_safe_subsetting_policy.test_subsetting_policy.id

	#Optional
	is_derived_schema = var.subsetting_policy_subsetting_schema_is_derived_schema
	schema_name = var.subsetting_policy_subsetting_schema_schema_name
}
```

## Argument Reference

The following arguments are supported:

* `is_derived_schema` - (Optional) A filter to return the schemas which are derived. A schema is derived if it is related to a input schema. 
* `schema_name` - (Optional) A filter to return only items related to specific schema name.
* `subsetting_policy_id` - (Required) The OCID of the subsetting policy.


## Attributes Reference

The following attributes are exported:

* `subsetting_schema_collection` - The list of subsetting_schema_collection.

### SubsettingPolicySubsettingSchema Reference

The following attributes are exported:

* `items` - An array of subsetting schema summary objects.
	* `is_derived` - Indicates if the schema is a derived schema and not directly came as input from the user. A schema is derived if it is related to a input schema
	* `schema_name` - The database schema present in subsetting policy which is to be subsetted.

