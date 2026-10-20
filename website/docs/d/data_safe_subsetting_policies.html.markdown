---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_subsetting_policies"
sidebar_current: "docs-oci-datasource-data_safe-subsetting_policies"
description: |-
  Provides the list of Subsetting Policies in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_subsetting_policies
This data source provides the list of Subsetting Policies in Oracle Cloud Infrastructure Data Safe service.

Gets a list of subsetting policies based on the specified query parameters.

## Example Usage

```hcl
data "oci_data_safe_subsetting_policies" "test_subsetting_policies" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.subsetting_policy_access_level
	compartment_id_in_subtree = var.subsetting_policy_compartment_id_in_subtree
	display_name = var.subsetting_policy_display_name
	masking_policy_id = oci_data_safe_masking_policy.test_masking_policy.id
	sensitive_data_model_id = oci_data_safe_sensitive_data_model.test_sensitive_data_model.id
	state = var.subsetting_policy_state
	subsetting_policy_id = oci_data_safe_subsetting_policy.test_subsetting_policy.id
	target_id = oci_cloud_guard_target.test_target.id
	time_created_greater_than_or_equal_to = var.subsetting_policy_time_created_greater_than_or_equal_to
	time_created_less_than = var.subsetting_policy_time_created_less_than
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `display_name` - (Optional) A filter to return only resources that match the specified display name. 
* `masking_policy_id` - (Optional) A filter to return only the resources that match the specified masking policy OCID.
* `sensitive_data_model_id` - (Optional) A filter to return only the resources that match the specified sensitive data model OCID.
* `state` - (Optional) A filter to return only the resources that match the specified lifecycle states.
* `subsetting_policy_id` - (Optional) A filter to return only the resources that match the specified subsetting policy OCID.
* `target_id` - (Optional) A filter to return only items related to a specific target OCID.
* `time_created_greater_than_or_equal_to` - (Optional) A filter to return only the resources that were created after the specified date and time, as defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Using TimeCreatedGreaterThanOrEqualToQueryParam parameter retrieves all resources created after that date.

	**Example:** 2016-12-19T16:39:57.600Z 
* `time_created_less_than` - (Optional) Search for resources that were created before a specific date. Specifying this parameter corresponding `timeCreatedLessThan` parameter will retrieve all resources created before the specified created date, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by RFC 3339.

	**Example:** 2016-12-19T16:39:57.600Z 


## Attributes Reference

The following attributes are exported:

* `subsetting_policy_collection` - The list of subsetting_policy_collection.

### SubsettingPolicy Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the subsetting policy
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm) Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the subsetting policy
* `display_name` - The display name of the subsetting policy
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the subsetting policy
* `is_redo_logging_enabled` - Indicates if redo logging is enabled during a subsetting operation. It's disabled by default. Set this attribute to true to enable redo logging. By default, subsetting disables redo logging and flashback logging to purge any original   data from logs. However, in certain circumstances when you only want to test subsetting, rollback changes, and retry subsetting, you could enable logging and use a flashback database to retrieve the original data after it has been subsetted. 
* `is_refresh_stats_enabled` - Indicates if statistics gathering is enabled. It's enabled by default. Set this attribute to false to disable statistics gathering. The subsetting process gathers statistics on database tables after subsetting completes 
* `masking_policy_id` - The OCID of the masking policy associated with this subsetting policy
* `parallel_degree` - Specifies options to enable parallel execution when running data subsetting. Allowed values are 'NONE' (no parallelism), 'DEFAULT' (the Oracle Database computes the optimum degree of parallelism) or an integer value to be used as the degree of parallelism. Parallel execution helps effectively use multiple CPUs and improve subsetting performance. Refer to the Oracle Database parallel execution framework when choosing an explicit degree of parallelism 
* `post_subsetting_script` - A post-subsetting script, which can contain SQL and PL/SQL statements. It's executed after the core subsetting script generated using the subsetting policy. It's usually used to perform additional transformation or cleanup work after subsetting. 
* `pre_subsetting_script` - A pre-subsetting script, which can contain SQL and PL/SQL statements. It's executed before  the core subsetting script generated using the subsetting policy. It's usually used to perform any preparation or prerequisite work before subsetting data 
* `recompile` - Specifies how to recompile invalid objects post data subsetting. Allowed values are 'SERIAL' (recompile in serial),  'PARALLEL' (recompile in parallel), 'NONE' (do not recompile). If it's set to PARALLEL, the value of parallelDegree attribute is used. Use the built-in UTL_RECOMP package to recompile any remaining invalid objects after subsetting completes 
* `schema_source` - The source of subsetting schemas
	* `derived_schemas` - The schemas which are related to the input list of schemas in 'schemasForSubsetting'. These schemas can also be impacted from the subsetting process due to their relations with the schemas in 'schemasForSubsetting' 
	* `schema_source` - The source of subsetting schemas
	* `schemas_for_subsetting` - The schemas to be subsetted
	* `sensitive_data_model_id` - The OCID of the sensitive data model that's used as the source of subsetting schemas
	* `target_id` - The OCID of the target database that's used as the source of subsetting schemas
* `state` - The current state of the subsetting policy
* `time_created` - The date and time the subsetting policy was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339) 
* `time_updated` - The date and time the subsetting policy was last updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339) 
* `unrelated_tables_action` - Strategy to be applied for tables which are not impacted by any of the subsetting rules

