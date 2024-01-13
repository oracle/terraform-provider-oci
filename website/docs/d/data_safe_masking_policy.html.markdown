---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_masking_policy"
sidebar_current: "docs-oci-datasource-data_safe-masking_policy"
description: |-
  Provides details about a specific Masking Policy in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_masking_policy
This data source provides details about a specific Masking Policy resource in Oracle Cloud Infrastructure Data Safe service.

Gets the details of the specified masking policy.

## Example Usage

```hcl
data "oci_data_safe_masking_policy" "test_masking_policy" {
	#Required
	masking_policy_id = oci_data_safe_masking_policy.test_masking_policy.id
}
```

## Argument Reference

The following arguments are supported:

* `masking_policy_id` - (Required) The OCID of the masking policy.


## Attributes Reference

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

