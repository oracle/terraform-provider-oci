---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_masking_policy"
sidebar_current: "docs-oci-resource-data_safe-masking_policy"
description: |-
  Provides the Masking Policy resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_masking_policy
This resource provides the Masking Policy resource in Oracle Cloud Infrastructure Data Safe service.

Creates a new masking policy and associates it with a sensitive data model or a target database.

To use a sensitive data model as the source of masking columns, set the columnSource attribute to
SENSITIVE_DATA_MODEL and provide the sensitiveDataModelId attribute. After creating a masking policy,
you can use the AddMaskingColumnsFromSdm operation to automatically add all the columns from
the associated sensitive data model. In this case, the target database associated with the
sensitive data model is used for column and masking format validations.

You can also create a masking policy without using a sensitive data model. In this case,
you need to associate your masking policy with a target database by setting the columnSource
attribute to TARGET and providing the targetId attribute. The specified target database
is used for column and masking format validations.

After creating a masking policy, you can use the CreateMaskingColumn or PatchMaskingColumns
operation to manually add columns to the policy. You need to add the parent columns only,
and it automatically adds the child columns (in referential relationship with the parent columns)
from the associated sensitive data model or target database.


## Example Usage

```hcl
resource "oci_data_safe_masking_policy" "test_masking_policy" {
	#Required
	column_source {
		#Required
		column_source = var.masking_policy_column_source_column_source

		#Optional
		sensitive_data_model_id = oci_data_safe_sensitive_data_model.test_sensitive_data_model.id
		target_id = oci_cloud_guard_target.test_target.id
	}
	compartment_id = var.compartment_id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.masking_policy_description
	display_name = var.masking_policy_display_name
	freeform_tags = {"Department"= "Finance"}
	is_drop_temp_tables_enabled = var.masking_policy_is_drop_temp_tables_enabled
	is_redo_logging_enabled = var.masking_policy_is_redo_logging_enabled
	is_refresh_stats_enabled = var.masking_policy_is_refresh_stats_enabled
	parallel_degree = var.masking_policy_parallel_degree
	post_masking_script = var.masking_policy_post_masking_script
	pre_masking_script = var.masking_policy_pre_masking_script
	recompile = var.masking_policy_recompile
}
```

## Argument Reference

The following arguments are supported:

* `column_source` - (Required) (Updatable) Details to associate a column source with a masking policy.
	* `column_source` - (Required) (Updatable) The source of masking columns.
	* `sensitive_data_model_id` - (Required when column_source=SENSITIVE_DATA_MODEL) (Updatable) The OCID of the sensitive data model to be associated as the column source with the masking policy.
	* `target_id` - (Required when column_source=TARGET) (Updatable) The OCID of the target database to be associated as the column source with the masking policy.
* `compartment_id` - (Required) (Updatable) The OCID of the compartment where the masking policy should be created.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) The description of the masking policy.
* `display_name` - (Optional) (Updatable) The display name of the masking policy. The name does not have to be unique, and it's changeable.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `is_drop_temp_tables_enabled` - (Optional) (Updatable) Indicates if the temporary tables created during a masking operation should be dropped after masking. It's enabled by default. Set this attribute to false to preserve the temporary tables. Masking creates temporary tables that map the original sensitive  data values to mask values. By default, these temporary tables are dropped after masking. But, in some cases, you may want  to preserve this information to track how masking changed your data. Note that doing so compromises security. These tables  must be dropped before the database is available for unprivileged users.  
* `is_redo_logging_enabled` - (Optional) (Updatable) Indicates if redo logging is enabled during a masking operation. It's disabled by default. Set this attribute to true to enable redo logging. By default, masking disables redo logging and flashback logging to purge any original unmasked  data from logs. However, in certain circumstances when you only want to test masking, rollback changes, and retry masking, you could enable logging and use a flashback database to retrieve the original unmasked data after it has been masked.  
* `is_refresh_stats_enabled` - (Optional) (Updatable) Indicates if statistics gathering is enabled. It's enabled by default. Set this attribute to false to disable statistics gathering. The masking process gathers statistics on masked database tables after masking completes. 
* `parallel_degree` - (Optional) (Updatable) Specifies options to enable parallel execution when running data masking. Allowed values are 'NONE' (no parallelism), 'DEFAULT' (the Oracle Database computes the optimum degree of parallelism) or an integer value to be used as the degree of parallelism. Parallel execution helps effectively use multiple CPUs and improve masking performance. Refer to the Oracle Database parallel execution framework when choosing an explicit degree of parallelism. 
* `post_masking_script` - (Optional) (Updatable) A post-masking script, which can contain SQL and PL/SQL statements. It's executed after the core masking script generated using the masking policy. It's usually used to perform additional transformation or cleanup work after masking. 
* `pre_masking_script` - (Optional) (Updatable) A pre-masking script, which can contain SQL and PL/SQL statements. It's executed before  the core masking script generated using the masking policy. It's usually used to perform any preparation or prerequisite work before masking data. 
* `recompile` - (Optional) (Updatable) Specifies how to recompile invalid objects post data masking. Allowed values are 'SERIAL' (recompile in serial),  'PARALLEL' (recompile in parallel), 'NONE' (do not recompile). If it's set to PARALLEL, the value of parallelDegree attribute is used. Use the built-in UTL_RECOMP package to recompile any remaining invalid objects after masking completes. 
* `add_masking_columns_from_sdm_trigger` - (Optional) (Updatable) An optional property when incremented triggers Add Masking Columns From Sdm. Could be set to any integer value.
* `generate_health_report_trigger` - (Optional) (Updatable) An optional property when incremented triggers Generate Health Report. Could be set to any integer value.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Masking Policy
	* `update` - (Defaults to 20 minutes), when updating the Masking Policy
	* `delete` - (Defaults to 20 minutes), when destroying the Masking Policy


## Import

MaskingPolicies can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_masking_policy.test_masking_policy "id"
```

