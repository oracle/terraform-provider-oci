---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_user_assessment_user_access_analytics"
sidebar_current: "docs-oci-datasource-data_safe-user_assessment_user_access_analytics"
description: |-
  Provides the list of User Assessment User Access Analytics in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_user_assessment_user_access_analytics
This data source provides the list of User Assessment User Access Analytics in Oracle Cloud Infrastructure Data Safe service.

Gets a list of aggregated user access analytics in the specified target in a compartment.


## Example Usage

```hcl
data "oci_data_safe_user_assessment_user_access_analytics" "test_user_assessment_user_access_analytics" {
	#Required
	user_assessment_id = oci_data_safe_user_assessment.test_user_assessment.id
}
```

## Argument Reference

The following arguments are supported:

* `user_assessment_id` - (Required) The OCID of the user assessment.


## Attributes Reference

The following attributes are exported:

* `user_access_analytics_collection` - The list of user_access_analytics_collection.

### UserAssessmentUserAccessAnalytic Reference

The following attributes are exported:

* `items` - An array of user access analytics summary objects.
	* `user_assessment_user_access_analytic_count` - The total count of schemas a user can access
	* `user_name` - Name of the user.

