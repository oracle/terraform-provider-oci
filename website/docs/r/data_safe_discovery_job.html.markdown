---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_discovery_job"
sidebar_current: "docs-oci-resource-data_safe-discovery_job"
description: |-
  Provides the Discovery Job resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_discovery_job
This resource provides the Discovery Job resource in Oracle Cloud Infrastructure Data Safe service.

Performs incremental data discovery for the specified sensitive data model. It uses the target database associated
with the sensitive data model.
After performing data discovery, you can use ListDiscoveryJobResults to view the discovery results, PatchDiscoveryJobResults
to specify the action you want perform on these results, and then ApplyDiscoveryJobResults to process the results
and apply them to the sensitive data model.


## Example Usage

```hcl
resource "oci_data_safe_discovery_job" "test_discovery_job" {
	#Required
	compartment_id = var.compartment_id
	sensitive_data_model_id = oci_data_safe_sensitive_data_model.test_sensitive_data_model.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	discovery_type = var.discovery_job_discovery_type
	display_name = var.discovery_job_display_name
	freeform_tags = {"Department"= "Finance"}
	is_app_defined_relation_discovery_enabled = var.discovery_job_is_app_defined_relation_discovery_enabled
	is_include_all_schemas = var.discovery_job_is_include_all_schemas
	is_include_all_sensitive_types = var.discovery_job_is_include_all_sensitive_types
	is_sample_data_collection_enabled = var.discovery_job_is_sample_data_collection_enabled
	schemas_for_discovery = var.discovery_job_schemas_for_discovery
	sensitive_type_ids_for_discovery = var.discovery_job_sensitive_type_ids_for_discovery
	tables_for_discovery {
		#Required
		schema_name = var.discovery_job_tables_for_discovery_schema_name

		#Optional
		table_names = var.discovery_job_tables_for_discovery_table_names
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment where the discovery job resource should be created.
* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `discovery_type` - (Optional) The type of the discovery job. It defines the job's scope. NEW identifies new sensitive columns in the target database that are not in the sensitive data model. DELETED identifies columns that are present in the sensitive data model but have been deleted from the target database. MODIFIED identifies columns that are present in the target database as well as the sensitive data model but some of their attributes have been modified. ALL covers all the above three scenarios and reports new, deleted and modified columns. 
* `display_name` - (Optional) A user-friendly name for the discovery job. Does not have to be unique, and it is changeable. Avoid entering confidential information.
* `freeform_tags` - (Optional) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `is_app_defined_relation_discovery_enabled` - (Optional) Indicates if the discovery job should identify potential application-level (non-dictionary) referential relationships between columns. Note that data discovery automatically identifies and adds database-level (dictionary-defined) relationships. This option helps identify application-level relationships that are not defined in the database dictionary, which in turn, helps identify additional sensitive columns and preserve referential integrity during data masking. It's disabled by default and should be used only if there is a need to identify application-level relationships. 
* `is_include_all_schemas` - (Optional) Indicates if all the schemas should be scanned by the discovery job. If it is set to true, sensitive data is discovered in all schemas (except for schemas maintained by Oracle). If both attributes are not provided, the configuration from the sensitive data model is used. 
* `is_include_all_sensitive_types` - (Optional) Indicates if all the existing sensitive types should be used by the discovery job. If it's set to true, the sensitiveTypeIdsForDiscovery attribute is ignored and all sensitive types are used for data discovery. If both attributes are not provided, the configuration from the sensitive data model is used. 
* `is_sample_data_collection_enabled` - (Optional) Indicates if the discovery job should collect and store sample data values for the discovered columns. Sample data helps review the discovered columns and ensure that they actually contain sensitive data. As it collects original data from the target database, it's disabled by default and should be used only if it's acceptable to store sample data in Data Safe's repository in Oracle Cloud. Note that sample data values are not collected for columns with the following data types: LONG, LOB, RAW, XMLTYPE and BFILE. 
* `schemas_for_discovery` - (Optional) The schemas to be scanned by the discovery job. If not provided, the schemasForDiscovery attribute of the sensitive data model is used to get the list of schemas. 
* `sensitive_data_model_id` - (Required) The OCID of the sensitive data model.
* `sensitive_type_ids_for_discovery` - (Optional) The OCIDs of the sensitive types to be used by the discovery job. If not provided, the sensitiveTypeIdsForDiscovery attribute of the sensitive data model is used to get the list of sensitive types. 
* `tables_for_discovery` - (Optional) The data discovery jobs will scan the tables specified here, including both schemas and tables. In the absence  of explicit input, the list of tables is obtained from the tablesForDiscovery attribute of the sensitive data model. 
	* `schema_name` - (Required) This contains the name of the schema.
	* `table_names` - (Optional) This contains an optional list of the table names.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the discovery job.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `discovery_type` - The type of the discovery job. It defines the job's scope. NEW identifies new sensitive columns in the target database that are not in the sensitive data model. DELETED identifies columns that are present in the sensitive data model but have been deleted from the target database. MODIFIED identifies columns that are present in the target database as well as the sensitive data model but some of their attributes have been modified. ALL covers all the above three scenarios and reports new, deleted and modified columns. 
* `display_name` - The display name of the discovery job.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the discovery job.
* `is_app_defined_relation_discovery_enabled` - Indicates if the discovery job should identify potential application-level (non-dictionary) referential relationships between columns. Note that data discovery automatically identifies and adds database-level (dictionary-defined) relationships. This option helps identify application-level relationships that are not defined in the database dictionary, which in turn, helps identify additional sensitive columns and preserve referential integrity during data masking. It's disabled by default and should be used only if there is a need to identify application-level relationships. 
* `is_include_all_schemas` - Indicates if all the schemas in the associated target database are used for data discovery. If it is set to true, sensitive data is discovered in all schemas (except for schemas maintained by Oracle). 
* `is_include_all_sensitive_types` - Indicates if all the existing sensitive types are used for data discovery. If it's set to true, the sensitiveTypeIdsForDiscovery attribute is ignored and all sensitive types are used. 
* `is_sample_data_collection_enabled` - Indicates if the discovery job should collect and store sample data values for the discovered columns. Sample data helps review the discovered columns and ensure that they actually contain sensitive data. As it collects original data from the target database, it's disabled by default and should be used only if it's acceptable to store sample data in Data Safe's repository in Oracle Cloud. Note that sample data values are not collected for columns with the following data types: LONG, LOB, RAW, XMLTYPE and BFILE. 
* `schemas_for_discovery` - The schemas used for data discovery.
* `sensitive_data_model_id` - The OCID of the sensitive data model associated with the discovery job.
* `sensitive_type_ids_for_discovery` - The OCIDs of the sensitive types used for data discovery.
* `state` - The current state of the discovery job.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `tables_for_discovery` - The data discovery jobs will scan the tables specified here, including both schemas and tables. 
	* `schema_name` - This contains the name of the schema.
	* `table_names` - This contains an optional list of the table names.
* `target_id` - The OCID of the target database associated with the discovery job.
* `time_finished` - The date and time the discovery job finished, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339)..
* `time_started` - The date and time the discovery job started, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `total_columns_scanned` - The total number of columns scanned by the discovery job.
* `total_deleted_sensitive_columns` - The total number of deleted sensitive columns identified by the discovery job.
* `total_modified_sensitive_columns` - The total number of modified sensitive columns identified by the discovery job.
* `total_new_sensitive_columns` - The total number of new sensitive columns identified by the discovery job.
* `total_objects_scanned` - The total number of objects (tables and editioning views) scanned by the discovery job.
* `total_schemas_scanned` - The total number of schemas scanned by the discovery job.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Discovery Job
	* `update` - (Defaults to 20 minutes), when updating the Discovery Job
	* `delete` - (Defaults to 20 minutes), when destroying the Discovery Job


## Import

DiscoveryJobs can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_discovery_job.test_discovery_job "id"
```

