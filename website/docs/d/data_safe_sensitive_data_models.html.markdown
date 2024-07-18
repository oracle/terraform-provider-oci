---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_sensitive_data_models"
sidebar_current: "docs-oci-datasource-data_safe-sensitive_data_models"
description: |-
  Provides the list of Sensitive Data Models in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_sensitive_data_models
This data source provides the list of Sensitive Data Models in Oracle Cloud Infrastructure Data Safe service.

Gets a list of sensitive data models based on the specified query parameters.


## Example Usage

```hcl
data "oci_data_safe_sensitive_data_models" "test_sensitive_data_models" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.sensitive_data_model_access_level
	compartment_id_in_subtree = var.sensitive_data_model_compartment_id_in_subtree
	display_name = var.sensitive_data_model_display_name
	sensitive_data_model_id = oci_data_safe_sensitive_data_model.test_sensitive_data_model.id
	state = var.sensitive_data_model_state
	target_id = oci_cloud_guard_target.test_target.id
	time_created_greater_than_or_equal_to = var.sensitive_data_model_time_created_greater_than_or_equal_to
	time_created_less_than = var.sensitive_data_model_time_created_less_than
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `display_name` - (Optional) A filter to return only resources that match the specified display name. 
* `sensitive_data_model_id` - (Optional) A filter to return only the resources that match the specified sensitive data model OCID.
* `state` - (Optional) A filter to return only the resources that match the specified lifecycle state.
* `target_id` - (Optional) A filter to return only items related to a specific target OCID.
* `time_created_greater_than_or_equal_to` - (Optional) A filter to return only the resources that were created after the specified date and time, as defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Using TimeCreatedGreaterThanOrEqualToQueryParam parameter retrieves all resources created after that date.

	**Example:** 2016-12-19T16:39:57.600Z 
* `time_created_less_than` - (Optional) Search for resources that were created before a specific date. Specifying this parameter corresponding `timeCreatedLessThan` parameter will retrieve all resources created before the specified created date, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by RFC 3339.

	**Example:** 2016-12-19T16:39:57.600Z 


## Attributes Reference

The following attributes are exported:

* `sensitive_data_model_collection` - The list of sensitive_data_model_collection.

### SensitiveDataModel Reference

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
* `tables_for_discovery` - The data discovery jobs will scan the tables specified here, including both schemas and tables. For instance, the input could be in the format: [{schemaName: "HR", tableName: ["T1", "T2"]}, {schemaName:  "OE", tableName : ["T3", "T4"]}]. 
	* `schema_name` - This contains the name of the schema.
	* `table_names` - This contains an optional list of the table names.
* `target_id` - The OCID of the reference target database associated with the sensitive data model. All operations such as performing data discovery and adding columns manually are done in the context of the associated target database. 
* `time_created` - The date and time the sensitive data model was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_updated` - The date and time the sensitive data model was last updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).

