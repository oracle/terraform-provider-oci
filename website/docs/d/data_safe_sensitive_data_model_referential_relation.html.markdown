---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_sensitive_data_model_referential_relation"
sidebar_current: "docs-oci-datasource-data_safe-sensitive_data_model_referential_relation"
description: |-
  Provides details about a specific Sensitive Data Model Referential Relation in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_sensitive_data_model_referential_relation
This data source provides details about a specific Sensitive Data Model Referential Relation resource in Oracle Cloud Infrastructure Data Safe service.

Gets the details of the specified referential relation.

## Example Usage

```hcl
data "oci_data_safe_sensitive_data_model_referential_relation" "test_sensitive_data_model_referential_relation" {
	#Required
	referential_relation_key = var.sensitive_data_model_referential_relation_referential_relation_key
	sensitive_data_model_id = oci_data_safe_sensitive_data_model.test_sensitive_data_model.id
}
```

## Argument Reference

The following arguments are supported:

* `referential_relation_key` - (Required) The unique key that identifies the referential relation. It's numeric and unique within a sensitive data model. 
* `sensitive_data_model_id` - (Required) The OCID of the sensitive data model.


## Attributes Reference

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

