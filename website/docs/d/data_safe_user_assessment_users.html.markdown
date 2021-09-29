---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_user_assessment_users"
sidebar_current: "docs-oci-datasource-data_safe-user_assessment_users"
description: |-
  Provides the list of User Assessment Users in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_user_assessment_users
This data source provides the list of User Assessment Users in Oracle Cloud Infrastructure Data Safe service.

Gets a list of users of the specified user assessment. The result contains the database user details for each user, such
as user type, account status, last login time, user creation time, authentication type, user profile, and the date and time
of the latest password change. It also contains the user category derived from these user details as well as privileges
granted to each user.


## Example Usage

```hcl
data "oci_data_safe_user_assessment_users" "test_user_assessment_users" {
	#Required
	user_assessment_id = oci_data_safe_user_assessment.test_user_assessment.id

	#Optional
	access_level = var.user_assessment_user_access_level
	account_status = var.user_assessment_user_account_status
	authentication_type = var.user_assessment_user_authentication_type
	compartment_id_in_subtree = var.user_assessment_user_compartment_id_in_subtree
	target_id = oci_cloud_guard_target.test_target.id
	time_last_login_greater_than_or_equal_to = var.user_assessment_user_time_last_login_greater_than_or_equal_to
	time_last_login_less_than = var.user_assessment_user_time_last_login_less_than
	time_password_last_changed_greater_than_or_equal_to = var.user_assessment_user_time_password_last_changed_greater_than_or_equal_to
	time_password_last_changed_less_than = var.user_assessment_user_time_password_last_changed_less_than
	time_user_created_greater_than_or_equal_to = var.user_assessment_user_time_user_created_greater_than_or_equal_to
	time_user_created_less_than = var.user_assessment_user_time_user_created_less_than
	user_category = var.user_assessment_user_user_category
	user_key = var.user_assessment_user_user_key
	user_name = oci_identity_user.test_user.name
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `account_status` - (Optional) A filter to return only items that match the specified account status.
* `authentication_type` - (Optional) A filter to return only items that match the specified authentication type.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `target_id` - (Optional) A filter to return only items that match the specified target.
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

* `users` - The list of users.

### UserAssessmentUser Reference

The following attributes are exported:

* `account_status` - The user account status.
* `admin_roles` - The admin roles granted to the user.
* `authentication_type` - The user authentication method.
* `key` - The unique user key. This is a system-generated identifier. Use ListUsers to get the user key for a user.
* `target_id` - The OCID of the target database.
* `time_last_login` - The date and time when the user last logged in, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_password_changed` - The date and time when the user password was last changed, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_user_created` - The date and time when the user was created in the database, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `user_category` - The user category based on the privileges and other details of the user.
* `user_name` - The database user name.
* `user_profile` - The user profile name.
* `user_types` - The user type, which can be a combination of the following:

	'Admin Privileged': The user has administrative privileges. 'Application': The user is an Oracle E-Business Suite Applications (EBS) or Fusion Applications (FA) user. 'Privileged': The user is a privileged user. 'Schema': The user is EXPIRED & LOCKED / EXPIRED / LOCKED, or a schema-only account (authentication type is NONE). 'Non-privileged': The user is a non-privileged user. 

