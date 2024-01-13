---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_security_assessment_finding_analytics"
sidebar_current: "docs-oci-datasource-data_safe-security_assessment_finding_analytics"
description: |-
  Provides the list of Security Assessment Finding Analytics in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_security_assessment_finding_analytics
This data source provides the list of Security Assessment Finding Analytics in Oracle Cloud Infrastructure Data Safe service.

Gets a list of findings aggregated details in the specified compartment. This provides information about the overall state
of security assessment findings. You can use groupBy to get the count of findings under a certain risk level and with a certain findingKey, 
and as well as get the list of the targets that match the condition.
This data is especially useful content for the statistic chart or to support analytics.

When you perform the ListFindingAnalytics operation, if the parameter compartmentIdInSubtree is set to "true," and if the
parameter accessLevel is set to ACCESSIBLE, then the operation returns statistics from the compartments in which the requestor has INSPECT
permissions on at least one resource, directly or indirectly (in subcompartments). If the operation is performed at the
root compartment and the requestor does not have access to at least one subcompartment of the compartment specified by
compartmentId, then "Not Authorized" is returned.


## Example Usage

```hcl
data "oci_data_safe_security_assessment_finding_analytics" "test_security_assessment_finding_analytics" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.security_assessment_finding_analytic_access_level
	compartment_id_in_subtree = var.security_assessment_finding_analytic_compartment_id_in_subtree
	finding_key = var.security_assessment_finding_analytic_finding_key
	group_by = var.security_assessment_finding_analytic_group_by
	is_top_finding = var.security_assessment_finding_analytic_is_top_finding
	severity = var.security_assessment_finding_analytic_severity
	top_finding_status = var.security_assessment_finding_analytic_top_finding_status
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `finding_key` - (Optional) The unique key that identifies the finding. It is a string and unique within a security assessment.
* `group_by` - (Optional) Attribute by which the finding analytics data should be grouped.
* `is_top_finding` - (Optional) A filter to return only the findings that are marked as top findings.
* `severity` - (Optional) A filter to return only findings of a particular risk level.
* `top_finding_status` - (Optional) An optional filter to return only the top finding that match the specified status.


## Attributes Reference

The following attributes are exported:

* `finding_analytics_collection` - The list of finding_analytics_collection.

### SecurityAssessmentFindingAnalytic Reference

The following attributes are exported:

* `items` - The array of the summary objects of the analytics data of findings or top findings.
	* `dimensions` - The scope of analytics data.
		* `key` - Each finding in security assessment has an associated key (think of key as a finding's name). For a given finding, the key will be the same across targets. The user can use these keys to filter the findings. 
		* `severity` - The severity (risk level) of the finding.
		* `target_id` - The OCID of the target database.
		* `title` - The short title of the finding.
		* `top_finding_category` - The category of the top finding.
		* `top_finding_status` - The status of the top finding.  All findings will have "severity" to indicate the risk level, but only top findings will have "status".  Possible status: Pass / Risk (Low, Medium, High)/ Evaluate / Advisory / Deferred Instead of having "Low, Medium, High" in severity, "Risk" will include these three situations in status. 
	* `metric_name` - The name of the aggregation metric.
	* `security_assessment_finding_analytic_count` - The total count for the aggregation metric.

