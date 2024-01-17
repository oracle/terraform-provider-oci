---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_sensitive_data_model"
sidebar_current: "docs-oci-resource-data_safe-sensitive_data_model"
description: |-
  Provides the Sensitive Data Model resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_sensitive_data_model
This resource provides the Sensitive Data Model resource in Oracle Cloud Infrastructure Data Safe service.

Creates a new sensitive data model. If schemas and sensitive types are provided, it automatically runs data discovery
and adds the discovered columns to the sensitive data model. Otherwise, it creates an empty sensitive data model
that can be updated later.


## Example Usage

```hcl
resource "oci_data_safe_sensitive_data_model" "test_sensitive_data_model" {
	#Required
	compartment_id = var.compartment_id
	target_id = oci_cloud_guard_target.test_target.id

	#Optional
	app_suite_name = var.sensitive_data_model_app_suite_name
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.sensitive_data_model_description
	display_name = var.sensitive_data_model_display_name
	freeform_tags = {"Department"= "Finance"}
	is_app_defined_relation_discovery_enabled = var.sensitive_data_model_is_app_defined_relation_discovery_enabled
	is_include_all_schemas = var.sensitive_data_model_is_include_all_schemas
	is_include_all_sensitive_types = var.sensitive_data_model_is_include_all_sensitive_types
	is_sample_data_collection_enabled = var.sensitive_data_model_is_sample_data_collection_enabled
	schemas_for_discovery = var.sensitive_data_model_schemas_for_discovery
	sensitive_type_ids_for_discovery = var.sensitive_data_model_sensitive_type_ids_for_discovery
}
```

## Argument Reference

The following arguments are supported:

* `app_suite_name` - (Optional) (Updatable) The application suite name identifying a collection of applications. It's useful only if maintaining a sensitive data model for a suite of applications.
* `compartment_id` - (Required) (Updatable) The OCID of the compartment where the sensitive data model should be created.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) The description of the sensitive data model.
* `display_name` - (Optional) (Updatable) The display name of the sensitive data model. The name does not have to be unique, and it's changeable.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `is_app_defined_relation_discovery_enabled` - (Optional) (Updatable) Indicates if data discovery jobs should identify potential application-level (non-dictionary) referential relationships between columns. Note that data discovery automatically identifies and adds database-level (dictionary-defined) relationships. This option helps identify application-level relationships that are not defined in the database dictionary, which in turn, helps identify additional sensitive columns and preserve referential integrity during data masking. It's disabled by default and should be used only if there is a need to identify application-level relationships. 
* `is_include_all_schemas` - (Optional) Indicates if all the schemas in the associated target database should be scanned by data discovery jobs. If it is set to true, sensitive data is discovered in all schemas (except for schemas maintained by Oracle). 
* `is_include_all_sensitive_types` - (Optional) Indicates if all the existing sensitive types should be used by data discovery jobs. If it's set to true, the sensitiveTypeIdsForDiscovery attribute is ignored and all sensitive types are used for data discovery. 
* `is_sample_data_collection_enabled` - (Optional) (Updatable) Indicates if data discovery jobs should collect and store sample data values for the discovered columns. Sample data helps review the discovered columns and ensure that they actually contain sensitive data. As it collects original data from the target database, it's disabled by default and should be used only if it's acceptable to store sample data in Data Safe's repository in Oracle Cloud. Note that sample data values are not collected for columns with the following data types: LONG, LOB, RAW, XMLTYPE and BFILE. 
* `schemas_for_discovery` - (Optional) (Updatable) The schemas to be scanned by data discovery jobs.
* `sensitive_type_ids_for_discovery` - (Optional) (Updatable) The OCIDs of the sensitive types to be used by data discovery jobs. If OCID of a sensitive category is provided, all its child sensitive types are used for data discovery. 
* `target_id` - (Required) (Updatable) The OCID of the reference target database to be associated with the sensitive data model. All operations such as performing data discovery and adding columns manually are done in the context of the associated target database. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `app_suite_name` - The application suite name identifying a collection of applications. The default value is GENERIC. It's useful only if maintaining a sensitive data model for a suite of applications. 
* `compartment_id` - The OCID of the compartment that contains the sensitive data model.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the sensitive data model.
* `display_name` - The display name of the sensitive data model.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the sensitive data model.
* `is_app_defined_relation_discovery_enabled` - Indicates if data discovery jobs should identify potential application-level (non-dictionary) referential relationships between columns. Note that data discovery automatically identifies and adds database-level (dictionary-defined) relationships. This option helps identify application-level relationships that are not defined in the database dictionary, which in turn, helps identify additional sensitive columns and preserve referential integrity during data masking. It's disabled by default and should be used only if there is a need to identify application-level relationships. 
* `is_include_all_schemas` - Indicates if all the schemas in the associated target database should be scanned by data discovery jobs. If it is set to true, sensitive data is discovered in all schemas (except for schemas maintained by Oracle). 
* `is_include_all_sensitive_types` - Indicates if all the existing sensitive types should be used by data discovery jobs.If it's set to true, the sensitiveTypeIdsForDiscovery attribute is ignored and all sensitive types are used for data discovery. 
* `is_sample_data_collection_enabled` - Indicates if data discovery jobs should collect and store sample data values for the discovered columns. Sample data helps review the discovered columns and ensure that they actually contain sensitive data. As it collects original data from the target database, it's disabled by default and should be used only if it's acceptable to store sample data in Data Safe's repository in Oracle Cloud. Note that sample data values are not collected for columns with the following data types: LONG, LOB, RAW, XMLTYPE and BFILE. 
* `schemas_for_discovery` - The schemas to be scanned by data discovery jobs.
* `sensitive_type_ids_for_discovery` - The OCIDs of the sensitive types to be used by data discovery jobs.
* `state` - The current state of the sensitive data model.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `target_id` - The OCID of the reference target database associated with the sensitive data model. All operations such as performing data discovery and adding columns manually are done in the context of the associated target database. 
* `time_created` - The date and time the sensitive data model was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_updated` - The date and time the sensitive data model was last updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Sensitive Data Model
	* `update` - (Defaults to 20 minutes), when updating the Sensitive Data Model
	* `delete` - (Defaults to 20 minutes), when destroying the Sensitive Data Model


## Import

SensitiveDataModels can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_sensitive_data_model.test_sensitive_data_model "id"
```

