---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_user_assessment_user_analytics"
sidebar_current: "docs-oci-datasource-data_safe-user_assessment_user_analytics"
description: |-
  Provides the list of User Assessment User Analytics in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_user_assessment_user_analytics
This data source provides the list of User Assessment User Analytics in Oracle Cloud Infrastructure Data Safe service.

Gets a list of aggregated user details from the specified user assessment. This provides information about the overall state.
of database user security.  For example, the user details include how many users have the DBA role and how many users are in
the critical category. This data is especially useful content for dashboards or to support analytics.

When you perform the ListUserAnalytics operation, if the parameter compartmentIdInSubtree is set to "true," and if the
parameter accessLevel is set to ACCESSIBLE, then the operation returns compartments in which the requestor has INSPECT
permissions on at least one resource, directly or indirectly (in subcompartments). If the operation is performed at the
root compartment. If the requestor does not have access to at least one subcompartment of the compartment specified by
compartmentId, then "Not Authorized" is returned.

The parameter compartmentIdInSubtree applies when you perform ListUserAnalytics on the compartmentId passed and when it is
set to true, the entire hierarchy of compartments can be returned.

To use ListUserAnalytics to get a full list of all compartments and subcompartments in the tenancy from the root compartment,
set the parameter compartmentIdInSubtree to true and accessLevel to ACCESSIBLE.


## Example Usage

```hcl
data "oci_data_safe_user_assessment_user_analytics" "test_user_assessment_user_analytics" {
	#Required
	user_assessment_id = oci_data_safe_user_assessment.test_user_assessment.id

	#Optional
	access_level = var.user_assessment_user_analytic_access_level
	account_status = var.user_assessment_user_analytic_account_status
	authentication_type = var.user_assessment_user_analytic_authentication_type
	compartment_id_in_subtree = var.user_assessment_user_analytic_compartment_id_in_subtree
	target_id = oci_cloud_guard_target.test_target.id
	time_last_login_greater_than_or_equal_to = var.user_assessment_user_analytic_time_last_login_greater_than_or_equal_to
	time_last_login_less_than = var.user_assessment_user_analytic_time_last_login_less_than
	time_password_last_changed_greater_than_or_equal_to = var.user_assessment_user_analytic_time_password_last_changed_greater_than_or_equal_to
	time_password_last_changed_less_than = var.user_assessment_user_analytic_time_password_last_changed_less_than
	time_user_created_greater_than_or_equal_to = var.user_assessment_user_analytic_time_user_created_greater_than_or_equal_to
	time_user_created_less_than = var.user_assessment_user_analytic_time_user_created_less_than
	user_category = var.user_assessment_user_analytic_user_category
	user_key = var.user_assessment_user_analytic_user_key
	user_name = oci_identity_user.test_user.name
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `account_status` - (Optional) A filter to return only items that match the specified account status.
* `authentication_type` - (Optional) A filter to return only items that match the specified authentication type.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `target_id` - (Optional) A filter to return only items related to a specific target OCID.
* `time_last_login_greater_than_or_equal_to` - (Optional) A filter to return users whose last login time in the database is greater than or equal to the date and time specified, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).

	**Example:** 2016-12-19T16:39:57.600Z 
* `time_last_login_less_than` - (Optional) A filter to return users whose last login time in the database is less than the date and time specified, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). **Example:** 2016-12-19T16:39:57.600Z 
* `time_password_last_changed_greater_than_or_equal_to` - (Optional) A filter to return users whose last password change in the database is greater than or equal to the date and time specified, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).

	**Example:** 2016-12-19T16:39:57.600Z 
* `time_password_last_changed_less_than` - (Optional) A filter to return users whose last password change in the database is less than the date and time specified, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).

	**Example:** 2016-12-19T16:39:57.600Z 
* `time_user_created_greater_than_or_equal_to` - (Optional) A filter to return users whose creation time in the database is greater than or equal to the date and time specified, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). **Example:** 2016-12-19T16:39:57.600Z 
* `time_user_created_less_than` - (Optional) A filter to return users whose creation time in the database is less than the date and time specified, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). **Example:** 2016-12-19T16:39:57.600Z 
* `user_assessment_id` - (Required) The OCID of the user assessment.
* `user_category` - (Optional) A filter to return only items that match the specified user category.
* `user_key` - (Optional) A filter to return only items that match the specified user key.
* `user_name` - (Optional) A filter to return only items that match the specified user name.


## Attributes Reference

The following attributes are exported:

* `user_aggregations` - The list of user_aggregations.

### UserAssessmentUserAnalytic Reference

The following attributes are exported:

* `items` - The array of user aggregation data.

