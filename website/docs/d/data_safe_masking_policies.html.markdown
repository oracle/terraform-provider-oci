---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_masking_policies"
sidebar_current: "docs-oci-datasource-data_safe-masking_policies"
description: |-
  Provides the list of Masking Policies in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_masking_policies
This data source provides the list of Masking Policies in Oracle Cloud Infrastructure Data Safe service.

Gets a list of masking policies based on the specified query parameters.

## Example Usage

```hcl
data "oci_data_safe_masking_policies" "test_masking_policies" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.masking_policy_access_level
	compartment_id_in_subtree = var.masking_policy_compartment_id_in_subtree
	display_name = var.masking_policy_display_name
	masking_policy_id = oci_data_safe_masking_policy.test_masking_policy.id
	sensitive_data_model_id = oci_data_safe_sensitive_data_model.test_sensitive_data_model.id
	state = var.masking_policy_state
	target_id = oci_cloud_guard_target.test_target.id
	time_created_greater_than_or_equal_to = var.masking_policy_time_created_greater_than_or_equal_to
	time_created_less_than = var.masking_policy_time_created_less_than
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
* `target_id` - (Optional) A filter to return only items related to a specific target OCID.
* `time_created_greater_than_or_equal_to` - (Optional) A filter to return only the resources that were created after the specified date and time, as defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Using TimeCreatedGreaterThanOrEqualToQueryParam parameter retrieves all resources created after that date.

	**Example:** 2016-12-19T16:39:57.600Z 
* `time_created_less_than` - (Optional) Search for resources that were created before a specific date. Specifying this parameter corresponding `timeCreatedLessThan` parameter will retrieve all resources created before the specified created date, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by RFC 3339.

	**Example:** 2016-12-19T16:39:57.600Z 


## Attributes Reference

The following attributes are exported:

* `masking_policy_collection` - The list of masking_policy_collection.

### MaskingPolicy Reference

The following attributes are exported:

* `column_source` - The source of masking columns.
	* `column_source` - The source of masking columns.
	* `sensitive_data_model_id` - The OCID of the sensitive data model that's used as the source of masking columns.
	* `target_id` - The OCID of the target database that's used as the source of masking columns.
* `compartment_id` - The OCID of the compartment that contains the masking policy.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the masking policy.
* `display_name` - The display name of the masking policy.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the masking policy.
* `is_drop_temp_tables_enabled` - Indicates if the temporary tables created during a masking operation should be dropped after masking. It's enabled by default. Set this attribute to false to preserve the temporary tables. Masking creates temporary tables that map the original sensitive  data values to mask values. By default, these temporary tables are dropped after masking. But, in some cases, you may want  to preserve this information to track how masking changed your data. Note that doing so compromises security. These tables  must be dropped before the database is available for unprivileged users.  
* `is_redo_logging_enabled` - Indicates if redo logging is enabled during a masking operation. It's disabled by default. Set this attribute to true to enable redo logging. By default, masking disables redo logging and flashback logging to purge any original unmasked  data from logs. However, in certain circumstances when you only want to test masking, rollback changes, and retry masking, you could enable logging and use a flashback database to retrieve the original unmasked data after it has been masked.  
* `is_refresh_stats_enabled` - Indicates if statistics gathering is enabled. It's enabled by default. Set this attribute to false to disable statistics gathering. The masking process gathers statistics on masked database tables after masking completes. 
* `parallel_degree` - Specifies options to enable parallel execution when running data masking. Allowed values are 'NONE' (no parallelism), 'DEFAULT' (the Oracle Database computes the optimum degree of parallelism) or an integer value to be used as the degree of parallelism. Parallel execution helps effectively use multiple CPUs and improve masking performance. Refer to the Oracle Database parallel execution framework when choosing an explicit degree of parallelism. 
* `post_masking_script` - A post-masking script, which can contain SQL and PL/SQL statements. It's executed after the core masking script generated using the masking policy. It's usually used to perform additional transformation or cleanup work after masking. 
* `pre_masking_script` - A pre-masking script, which can contain SQL and PL/SQL statements. It's executed before  the core masking script generated using the masking policy. It's usually used to perform any preparation or prerequisite work before masking data. 
* `recompile` - Specifies how to recompile invalid objects post data masking. Allowed values are 'SERIAL' (recompile in serial),  'PARALLEL' (recompile in parallel), 'NONE' (do not recompile). If it's set to PARALLEL, the value of parallelDegree attribute is used. Use the built-in UTL_RECOMP package to recompile any remaining invalid objects after masking completes. 
* `state` - The current state of the masking policy.
* `time_created` - The date and time the masking policy was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  
* `time_updated` - The date and time the masking policy was last updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339)  

