---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_subsetting_policy"
sidebar_current: "docs-oci-resource-data_safe-subsetting_policy"
description: |-
  Provides the Subsetting Policy resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_subsetting_policy
This resource provides the Subsetting Policy resource in Oracle Cloud Infrastructure Data Safe service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/data-safe/latest/SubsettingPolicy

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/datasafe

Creates a new subsetting policy and associates it with a sensitive data model or a target database.

To use a sensitive data model as the source of subsetting schemas, set the schemaSource attribute to
SENSITIVE_DATA_MODEL and provide the sensitiveDataModelId attribute. In this case, the target database associated with the
sensitive data model is used for subsetting rules validations.

You can also create a subsetting policy without using a sensitive data model. In this case,
you need to associate your subsetting policy with a target database by setting the schemaSource
attribute to TARGET and providing the targetId attribute. The specified target database
is used for subsetting rules validations.

After creating a subsetting policy, you can use the CreateSubsettingRule
operation to manually add subsetting rules to the policy.


## Example Usage

```hcl
resource "oci_data_safe_subsetting_policy" "test_subsetting_policy" {
	#Required
	compartment_id = var.compartment_id
	schema_source {
		#Required
		schema_source = var.subsetting_policy_schema_source_schema_source

		#Optional
		schemas_for_subsetting = var.subsetting_policy_schema_source_schemas_for_subsetting
		sensitive_data_model_id = oci_data_safe_sensitive_data_model.test_sensitive_data_model.id
		target_id = oci_cloud_guard_target.test_target.id
	}

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.subsetting_policy_description
	display_name = var.subsetting_policy_display_name
	freeform_tags = {"Department"= "Finance"}
	is_redo_logging_enabled = var.subsetting_policy_is_redo_logging_enabled
	is_refresh_stats_enabled = var.subsetting_policy_is_refresh_stats_enabled
	masking_policy_id = oci_data_safe_masking_policy.test_masking_policy.id
	parallel_degree = var.subsetting_policy_parallel_degree
	post_subsetting_script = var.subsetting_policy_post_subsetting_script
	pre_subsetting_script = var.subsetting_policy_pre_subsetting_script
	recompile = var.subsetting_policy_recompile
	unrelated_tables_action = var.subsetting_policy_unrelated_tables_action
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment where the subsetting policy should be created
* `check_type` - (Optional) The health check type used when `generate_health_report_trigger` is incremented. Allowed values are `ALL` and `TABLESPACE_CHECK`
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm) Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) The description of the subsetting policy
* `display_name` - (Optional) (Updatable) The display name of the subsetting policy. The name does not have to be unique, and it's changeable
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `is_redo_logging_enabled` - (Optional) (Updatable) Indicates if redo logging is enabled during a subsetting operation. It's disabled by default. Set this attribute to true to enable redo logging. By default, subsetting disables redo logging and flashback logging to purge any original  data from logs. However, in certain circumstances when you only want to test subsetting, rollback changes, and retry subsetting, you could enable logging and use a flashback database to retrieve the original data after it has been subsetted 
* `is_refresh_stats_enabled` - (Optional) (Updatable) Indicates if statistics gathering is enabled. It's enabled by default. Set this attribute to false to disable statistics gathering. The subsetting process gathers statistics on subsetted database tables after subsetting completes 
* `masking_policy_id` - (Optional) (Updatable) The OCID of the masking policy to associate with this subsetting policy
* `parallel_degree` - (Optional) (Updatable) Specifies options to enable parallel execution when running data subsetting. Allowed values are 'NONE' (no parallelism), 'DEFAULT' (the Oracle Database computes the optimum degree of parallelism) or an integer value to be used as the degree of parallelism. Parallel execution helps effectively use multiple CPUs and improve subsetting performance. Refer to the Oracle Database parallel execution framework when choosing an explicit degree of parallelism 
* `post_subsetting_script` - (Optional) (Updatable) A post-subsetting script, which can contain SQL and PL/SQL statements. It's executed after the subsetting process. It's usually used to perform additional transformation or cleanup work after subsetting data. 
* `pre_subsetting_script` - (Optional) (Updatable) A pre-subsetting script, which can contain SQL and PL/SQL statements. It's executed before  the subsetting process. It's usually used to perform any preparation or prerequisite work before subsetting data. 
* `recompile` - (Optional) (Updatable) Specifies how to recompile invalid objects post data subsetting. Allowed values are 'SERIAL' (recompile in serial),  'PARALLEL' (recompile in parallel), 'NONE' (do not recompile). If it's set to PARALLEL, the value of parallelDegree attribute is used. Use the built-in UTL_RECOMP package to recompile any remaining invalid objects after subsetting completes 
* `schema_source` - (Required) (Updatable) The source of subsetting schemas
	* `schema_source` - (Required) (Updatable) The source of subsetting schemas
	* `schemas_for_subsetting` - (Applicable when schema_source=TARGET) (Updatable) The schemas to be subsetted
	* `sensitive_data_model_id` - (Required when schema_source=SENSITIVE_DATA_MODEL) (Updatable) The OCID of the sensitive data model that's used as the source of subsetting schemas
	* `target_id` - (Required when schema_source=TARGET) (Updatable) The OCID of the target database that's used as the source of subsetting schemas
* `unrelated_tables_action` - (Optional) (Updatable) Strategy to be applied for tables which are not impacted by any of the subsetting rules
* `tablespace` - (Optional) The tablespace used by the health report when `generate_health_report_trigger` is incremented
* `target_credentials` - (Optional) Credentials used by the health report when `generate_health_report_trigger` is incremented. Required when the trigger is changed.
* `target_id` - (Optional) The target database override used by the health report when `generate_health_report_trigger` is incremented. If omitted, the target associated with the policy is used.
* `generate_health_report_trigger` - (Optional) An optional property that triggers Generate Health Report when incremented. The trigger runs after a policy update; use `oci_data_safe_subsetting_policy_health_report_management` when the report must be generated after creating separate subsetting rules.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Subsetting Policy
	* `update` - (Defaults to 20 minutes), when updating the Subsetting Policy
	* `delete` - (Defaults to 20 minutes), when destroying the Subsetting Policy


## Import

SubsettingPolicies can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_subsetting_policy.test_subsetting_policy "id"
```
