---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_user_assessment_profile_analytics"
sidebar_current: "docs-oci-datasource-data_safe-user_assessment_profile_analytics"
description: |-
  Provides the list of User Assessment Profile Analytics in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_user_assessment_profile_analytics
This data source provides the list of User Assessment Profile Analytics in Oracle Cloud Infrastructure Data Safe service.

Gets a list of aggregated user profile details in the specified compartment. This provides information about the 
overall profiles available. For example, the user profile details include how many users have the profile assigned
and do how many use password verification function. This data is especially useful content for dashboards or to support analytics.

When you perform the ListProfileAnalytics operation, if the parameter compartmentIdInSubtree is set to "true," and if the
parameter accessLevel is set to ACCESSIBLE, then the operation returns compartments in which the requestor has INSPECT
permissions on at least one resource, directly or indirectly (in subcompartments). If the operation is performed at the
root compartment and the requestor does not have access to at least one subcompartment of the compartment specified by
compartmentId, then "Not Authorized" is returned.

The parameter compartmentIdInSubtree applies when you perform ListProfileAnalytics on the compartmentId passed and when it is
set to true, the entire hierarchy of compartments can be returned.

To use ListProfileAnalytics to get a full list of all compartments and subcompartments in the tenancy from the root compartment,
set the parameter compartmentIdInSubtree to true and accessLevel to ACCESSIBLE.


## Example Usage

```hcl
data "oci_data_safe_user_assessment_profile_analytics" "test_user_assessment_profile_analytics" {
	#Required
	compartment_id = var.compartment_id
	user_assessment_id = oci_data_safe_user_assessment.test_user_assessment.id

	#Optional
	access_level = var.user_assessment_profile_analytic_access_level
	compartment_id_in_subtree = var.user_assessment_profile_analytic_compartment_id_in_subtree
	profile_name = oci_optimizer_profile.test_profile.name
	target_id = oci_cloud_guard_target.test_target.id
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `profile_name` - (Optional) A filter to return only items that match the specified profile name.
* `target_id` - (Optional) A filter to return only items related to a specific target OCID.
* `user_assessment_id` - (Required) The OCID of the user assessment.


## Attributes Reference

The following attributes are exported:

* `profile_aggregations` - The list of profile_aggregations.

### UserAssessmentProfileAnalytic Reference

The following attributes are exported:

* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `items` - The array of profile aggregation data.

