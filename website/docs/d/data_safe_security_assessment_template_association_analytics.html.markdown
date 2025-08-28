---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_security_assessment_template_association_analytics"
sidebar_current: "docs-oci-datasource-data_safe-security_assessment_template_association_analytics"
description: |-
  Provides the list of Security Assessment Template Association Analytics in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_security_assessment_template_association_analytics
This data source provides the list of Security Assessment Template Association Analytics in Oracle Cloud Infrastructure Data Safe service.

Gets a list of template association details in the specified compartment. This provides information about the
overall template usage, by returning the count of the target databases/target groups using the templates.

If the template baseline is created for a target group which contains several targets, we will have each individual target 
listed there as targetId field together with targetDatabaseGroupId. And if the template baseline is created for an individual target,
it will have targetId field only.

By leveraging the targetId filter, you will be able to know all the template or template baseline that this target has something to do with.
No matter if they are directly applied or created for this target, or they are for the target group the target belongs to.

When you perform the ListTemplateAssociationAnalytics operation, if the parameter compartmentIdInSubtree is set to "true," and if the
parameter accessLevel is set to ACCESSIBLE, then the operation returns statistics from the compartments in which the requestor has INSPECT
permissions on at least one resource, directly or indirectly (in subcompartments). If the operation is performed at the
root compartment and the requestor does not have access to at least one subcompartment of the compartment specified by
compartmentId, then "Not Authorized" is returned.


## Example Usage

```hcl
data "oci_data_safe_security_assessment_template_association_analytics" "test_security_assessment_template_association_analytics" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.security_assessment_template_association_analytic_access_level
	compartment_id_in_subtree = var.security_assessment_template_association_analytic_compartment_id_in_subtree
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
* `target_database_group_id` - (Optional) A filter to return the target database group that matches the specified OCID.
* `target_id` - (Optional) A filter to return only items related to a specific target OCID.
* `template_assessment_id` - (Optional) The OCID of the security assessment of type TEMPLATE.
* `template_baseline_assessment_id` - (Optional) The OCID of the security assessment of type TEMPLATE_BASELINE.


## Attributes Reference

The following attributes are exported:

* `template_association_analytics_collection` - The list of template_association_analytics_collection.

### SecurityAssessmentTemplateAssociationAnalytic Reference

The following attributes are exported:

* `items` - The array of template association analytics summary.
	* `dimensions` - The scope of template association analytics data.
		* `target_database_group_id` - The OCID of the target database group that the group assessment is created for.  This field will be in the response if the template was applied on a target group. 
		* `target_id` - The OCID of the target database. If the template was applied on a target group, this field will be the OCID of the target members of the target group. If the template was applied on an individual target, this field will contain that targetId. 
		* `template_assessment_id` - The OCID of the security assessment of type TEMPLATE.
		* `template_baseline_assessment_id` - The OCID of the security assessment of type TEMPLATE_BASELINE.
	* `metric_name` - The name of the aggregation metric.
	* `security_assessment_template_association_analytic_count` - The total count for the aggregation metric.

