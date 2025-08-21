---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_security_assessment_template_analytics"
sidebar_current: "docs-oci-datasource-data_safe-security_assessment_template_analytics"
description: |-
  Provides the list of Security Assessment Template Analytics in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_security_assessment_template_analytics
This data source provides the list of Security Assessment Template Analytics in Oracle Cloud Infrastructure Data Safe service.

Gets a list of template aggregated details in the specified compartment. This provides information about the
overall template usage, by returning the count of the target databases/target groups using the templates. It also provides information
about the statistics for the template baseline and the comparison related. If the comparison is done, it will show if there is any drift,
and how many checks have drifts.
The dimension field - isGroup identifies if the targetId belongs to a target group or a individual target.
The dimension field - isCompared identifies if the comparison between the latest assessment and the template baseline assessment is done or not.
The dimension field - isCompliant identifies if the latest assessment is compliant with the template baseline assessment or not.
The dimension field - totalChecksFailed identifies how many checks in the template have drifts in the comparison.

When you perform the ListTemplateAnalytics operation, if the parameter compartmentIdInSubtree is set to "true," and if the
parameter accessLevel is set to ACCESSIBLE, then the operation returns statistics from the compartments in which the requestor has INSPECT
permissions on at least one resource, directly or indirectly (in subcompartments). If the operation is performed at the
root compartment and the requestor does not have access to at least one subcompartment of the compartment specified by
compartmentId, then "Not Authorized" is returned.


## Example Usage

```hcl
data "oci_data_safe_security_assessment_template_analytics" "test_security_assessment_template_analytics" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.security_assessment_template_analytic_access_level
	compartment_id_in_subtree = var.security_assessment_template_analytic_compartment_id_in_subtree
	is_compared = var.security_assessment_template_analytic_is_compared
	is_compliant = var.security_assessment_template_analytic_is_compliant
	is_group = var.security_assessment_template_analytic_is_group
	target_database_group_id = oci_data_safe_target_database_group.test_target_database_group.id
	target_id = oci_cloud_guard_target.test_target.id
	template_assessment_id = oci_data_safe_template_assessment.test_template_assessment.id
	template_baseline_assessment_id = oci_data_safe_template_baseline_assessment.test_template_baseline_assessment.id
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `is_compared` - (Optional) A filter to return only the statistics where the comparison between the latest assessment and the template baseline assessment is done.
* `is_compliant` - (Optional) A filter to return only the statistics where the latest assessment is compliant with the template baseline assessment.
* `is_group` - (Optional) A filter to return only the target group related information if the OCID belongs to a target group.
* `target_database_group_id` - (Optional) A filter to return the target database group that matches the specified OCID.
* `target_id` - (Optional) A filter to return only items related to a specific target OCID.
* `template_assessment_id` - (Optional) The OCID of the security assessment of type TEMPLATE.
* `template_baseline_assessment_id` - (Optional) The OCID of the security assessment of type TEMPLATE_BASELINE.


## Attributes Reference

The following attributes are exported:

* `template_analytics_collection` - The list of template_analytics_collection.

### SecurityAssessmentTemplateAnalytic Reference

The following attributes are exported:

* `items` - The array of template analytics summary.
	* `dimensions` - The scope of analytics data.
		* `is_compared` - Indicates whether or not the comparison between the latest assessment and the template baseline assessment is done. If the value is false, it means the comparison is not done yet. 
		* `is_compliant` - Indicates whether or not the latest assessment is compliant with the template baseline assessment. If the value is false, it means there is drift in the comparison report and the totalChecksFailed field will have a non-zero value. 
		* `is_group` - Indicates whether or not the template security assessment is applied to a target group.  If the value is false, it means the template security assessment is applied to a individual target. 
		* `target_database_group_id` - The OCID of the target database group that the group assessment is created for.  This field will be in the response if the template was applied on a target group. 
		* `target_id` - The OCID of the target database. This field will be in the response if the template was applied on an individual target.
		* `template_assessment_id` - The OCID of the security assessment of type TEMPLATE.
		* `template_baseline_assessment_id` - The OCID of the security assessment of type TEMPLATE_BASELINE.
		* `time_last_compared` - The date and time when the comparison was made upon the template baseline. Conforms to the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
		* `total_checks` - The number of checks inside the template assessment.
		* `total_checks_failed` - Indicates how many checks in the template have drifts in the comparison report. This field is only present if isCompliant is false. 
		* `total_non_compliant_targets` - The number of the target(s) that have drifts in the comparison report. This field is only present if isCompared is true. 
		* `total_targets` - The number of the target(s) inside the target group for which the template baseline assessment was created for. If the isGroup field is false, the value will be 1, representing the single target. 
	* `metric_name` - The name of the aggregation metric.
	* `security_assessment_template_analytic_count` - The total count for the aggregation metric.

