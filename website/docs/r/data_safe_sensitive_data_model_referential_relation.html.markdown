---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_sensitive_data_model_referential_relation"
sidebar_current: "docs-oci-resource-data_safe-sensitive_data_model_referential_relation"
description: |-
  Provides the Sensitive Data Model Referential Relation resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_sensitive_data_model_referential_relation
This resource provides the Sensitive Data Model Referential Relation resource in Oracle Cloud Infrastructure Data Safe service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/data-safe/latest/SensitiveDataModelReferentialRelation

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/datasafe

Creates a new referential relation in the specified sensitive data model.


## Example Usage

```hcl
resource "oci_data_safe_sensitive_data_model_referential_relation" "test_sensitive_data_model_referential_relation" {
	#Required
	child {
		#Required
		app_name = var.sensitive_data_model_referential_relation_child_app_name
		column_group = var.sensitive_data_model_referential_relation_child_column_group
		object = var.sensitive_data_model_referential_relation_child_object
		object_type = var.sensitive_data_model_referential_relation_child_object_type
		schema_name = var.sensitive_data_model_referential_relation_child_schema_name

		#Optional
		sensitive_type_ids = var.sensitive_data_model_referential_relation_child_sensitive_type_ids
	}
	parent {
		#Required
		app_name = var.sensitive_data_model_referential_relation_parent_app_name
		column_group = var.sensitive_data_model_referential_relation_parent_column_group
		object = var.sensitive_data_model_referential_relation_parent_object
		object_type = var.sensitive_data_model_referential_relation_parent_object_type
		schema_name = var.sensitive_data_model_referential_relation_parent_schema_name

		#Optional
		sensitive_type_ids = var.sensitive_data_model_referential_relation_parent_sensitive_type_ids
	}
	relation_type = var.sensitive_data_model_referential_relation_relation_type
	sensitive_data_model_id = oci_data_safe_sensitive_data_model.test_sensitive_data_model.id

	#Optional
	is_sensitive = var.sensitive_data_model_referential_relation_is_sensitive
}
```

## Argument Reference

The following arguments are supported:

* `child` - (Required) columnsInfo object has details of column group with schema details.
	* `app_name` - (Required) The application name.
	* `column_group` - (Required) Group of columns in referential relation. Order needs to be maintained in the elements of the parent/child array listing. 
	* `object` - (Required) The database object that contains the columns.
	* `object_type` - (Required) The type of the database object that contains the sensitive column.
	* `schema_name` - (Required) The schema name.
	* `sensitive_type_ids` - (Optional) Sensitive type ocids of each column groups. Order needs to be maintained with the parent column group. For the DB defined referential relations identified during SDM creation, we cannot add sensitive types.  Instead use the sensitiveColumn POST API to mark the columns sensitive. 
* `is_sensitive` - (Optional) Add to sensitive data model if passed true. If false is passed, then the columns will not be added in the sensitive data model as sensitive columns and  if sensitive type OCIDs are assigned to the columns, then the sensitive type OCIDs will not be retained. 
* `parent` - (Required) columnsInfo object has details of column group with schema details.
	* `app_name` - (Required) The application name.
	* `column_group` - (Required) Group of columns in referential relation. Order needs to be maintained in the elements of the parent/child array listing. 
	* `object` - (Required) The database object that contains the columns.
	* `object_type` - (Required) The type of the database object that contains the sensitive column.
	* `schema_name` - (Required) The schema name.
	* `sensitive_type_ids` - (Optional) Sensitive type ocids of each column groups. Order needs to be maintained with the parent column group. For the DB defined referential relations identified during SDM creation, we cannot add sensitive types.  Instead use the sensitiveColumn POST API to mark the columns sensitive. 
* `relation_type` - (Required) The type of referential relationship the sensitive column has with its parent.  DB_DEFINED indicates that the relationship is defined in the database dictionary.  APP_DEFINED indicates that the relationship is defined at the application level and not in the database dictionary. 
* `sensitive_data_model_id` - (Required) The OCID of the sensitive data model.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Sensitive Data Model Referential Relation
	* `update` - (Defaults to 20 minutes), when updating the Sensitive Data Model Referential Relation
	* `delete` - (Defaults to 20 minutes), when destroying the Sensitive Data Model Referential Relation


## Import

SensitiveDataModelReferentialRelations can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_sensitive_data_model_referential_relation.test_sensitive_data_model_referential_relation "sensitiveDataModels/{sensitiveDataModelId}/referentialRelations/{referentialRelationKey}" 
```

