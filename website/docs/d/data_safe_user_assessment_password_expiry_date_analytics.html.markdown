---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_user_assessment_password_expiry_date_analytics"
sidebar_current: "docs-oci-datasource-data_safe-user_assessment_password_expiry_date_analytics"
description: |-
  Provides the list of User Assessment Password Expiry Date Analytics in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_user_assessment_password_expiry_date_analytics
This data source provides the list of User Assessment Password Expiry Date Analytics in Oracle Cloud Infrastructure Data Safe service.

Gets a list of count of the users with password expiry dates in next 30 days, between next 30-90 days, and beyond 90 days based on specified user assessment.
It internally uses the aforementioned userAnalytics api.

When you perform the ListPasswordExpiryDateAnalytics operation, if the parameter compartmentIdInSubtree is set to "true," and if the
parameter accessLevel is set to ACCESSIBLE, then the operation returns compartments in which the requestor has READ
permissions on at least one resource, directly or indirectly (in subcompartments). If the operation is performed at the
root compartment and the requestor does not have access to at least one subcompartment of the compartment specified by
compartmentId, then "Not Authorized" is returned.

To use ListPasswordExpiryDateAnalytics to get a full list of all compartments and subcompartments in the tenancy from the root compartment,
set the parameter compartmentIdInSubtree to true and accessLevel to ACCESSIBLE.


## Example Usage

```hcl
data "oci_data_safe_user_assessment_password_expiry_date_analytics" "test_user_assessment_password_expiry_date_analytics" {
	#Required
	user_assessment_id = oci_data_safe_user_assessment.test_user_assessment.id

	#Optional
	access_level = var.user_assessment_password_expiry_date_analytic_access_level
	compartment_id_in_subtree = var.user_assessment_password_expiry_date_analytic_compartment_id_in_subtree
	time_password_expiry_less_than = var.user_assessment_password_expiry_date_analytic_time_password_expiry_less_than
	user_category = var.user_assessment_password_expiry_date_analytic_user_category
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `time_password_expiry_less_than` - (Optional) A filter to return users whose password expiry date in the database is less than the date and time specified, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). **Example:** 2016-12-19T16:39:57.600Z 
* `user_assessment_id` - (Required) The OCID of the user assessment.
* `user_category` - (Optional) A filter to return only items that match the specified user category.


## Attributes Reference

The following attributes are exported:

* `user_aggregations` - The list of user_aggregations.

### UserAssessmentPasswordExpiryDateAnalytic Reference

The following attributes are exported:

* `items` - List containing maps as values. Example: `{"Operations": [ {"CostCenter": "42"} ] }` 

