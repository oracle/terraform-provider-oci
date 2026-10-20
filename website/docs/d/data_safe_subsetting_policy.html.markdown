---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_subsetting_policy"
sidebar_current: "docs-oci-datasource-data_safe-subsetting_policy"
description: |-
  Provides details about a specific Subsetting Policy in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_subsetting_policy
This data source provides details about a specific Subsetting Policy resource in Oracle Cloud Infrastructure Data Safe service.

Gets the details of the specified subsetting policy.

## Example Usage

```hcl
data "oci_data_safe_subsetting_policy" "test_subsetting_policy" {
	#Required
	subsetting_policy_id = oci_data_safe_subsetting_policy.test_subsetting_policy.id
}
```

## Argument Reference

The following arguments are supported:

* `subsetting_policy_id` - (Required) The OCID of the subsetting policy.


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

