---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_user_assessment_profiles"
sidebar_current: "docs-oci-datasource-data_safe-user_assessment_profiles"
description: |-
  Provides the list of User Assessment Profiles in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_user_assessment_profiles
This data source provides the list of User Assessment Profiles in Oracle Cloud Infrastructure Data Safe service.

Gets a list of user profiles containing the profile details along with the target id and user counts.

The ListProfiles operation returns only the profiles belonging to a certain target. If compartment type user assessment
id is provided, then profile information for all the targets belonging to the pertaining compartment is returned.
The list does not include any subcompartments of the compartment under consideration.

The parameter 'accessLevel' specifies whether to return only those compartments for which the requestor has 
INSPECT permissions on at least one resource directly or indirectly (ACCESSIBLE) (the resource can be in a 
subcompartment) or to return Not Authorized if Principal doesn't have access to even one of the child compartments.
This is valid only when 'compartmentIdInSubtree' is set to 'true'.

The parameter 'compartmentIdInSubtree' applies when you perform ListUserProfiles on the 'compartmentId' belonging
to the assessmentId passed and when it is set to true, the entire hierarchy of compartments can be returned.
To get a full list of all compartments and subcompartments in the tenancy (root compartment), set the parameter
'compartmentIdInSubtree' to true and 'accessLevel' to ACCESSIBLE.


## Example Usage

```hcl
data "oci_data_safe_user_assessment_profiles" "test_user_assessment_profiles" {
	#Required
	compartment_id = var.compartment_id
	user_assessment_id = oci_data_safe_user_assessment.test_user_assessment.id

	#Optional
	access_level = var.user_assessment_profile_access_level
	compartment_id_in_subtree = var.user_assessment_profile_compartment_id_in_subtree
	failed_login_attempts_greater_than_or_equal = var.user_assessment_profile_failed_login_attempts_greater_than_or_equal
	failed_login_attempts_less_than = var.user_assessment_profile_failed_login_attempts_less_than
	inactive_account_time_greater_than_or_equal = var.user_assessment_profile_inactive_account_time_greater_than_or_equal
	inactive_account_time_less_than = var.user_assessment_profile_inactive_account_time_less_than
	is_user_created = var.user_assessment_profile_is_user_created
	password_lock_time_greater_than_or_equal = var.user_assessment_profile_password_lock_time_greater_than_or_equal
	password_lock_time_less_than = var.user_assessment_profile_password_lock_time_less_than
	password_verification_function = var.user_assessment_profile_password_verification_function
	profile_name = oci_optimizer_profile.test_profile.name
	sessions_per_user_greater_than_or_equal = var.user_assessment_profile_sessions_per_user_greater_than_or_equal
	sessions_per_user_less_than = var.user_assessment_profile_sessions_per_user_less_than
	target_id = oci_cloud_guard_target.test_target.id
	user_count_greater_than_or_equal = var.user_assessment_profile_user_count_greater_than_or_equal
	user_count_less_than = var.user_assessment_profile_user_count_less_than
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `failed_login_attempts_greater_than_or_equal` - (Optional) An optional filter to return the profiles having allow failed login attempts number greater than or equal to the provided value. String value is used for accommodating the "UNLIMITED" and "DEFAULT" values. 
* `failed_login_attempts_less_than` - (Optional) An optional filter to return the profiles having failed login attempts number less than the provided value. String value is used for accommodating the "UNLIMITED" and "DEFAULT" values. 
* `inactive_account_time_greater_than_or_equal` - (Optional) An optional filter to return the profiles allowing inactive account time in days greater than or equal to the provided value. String value is used for accommodating the "UNLIMITED" and "DEFAULT" values. 
* `inactive_account_time_less_than` - (Optional) An optional filter to return the profiles  allowing inactive account time in days less than the provided value. String value is used for accommodating the "UNLIMITED" and "DEFAULT" values. 
* `is_user_created` - (Optional) An optional filter to return the user created profiles.
* `password_lock_time_greater_than_or_equal` - (Optional) An optional filter to return the profiles having password lock number greater than or equal to the provided value. String value is used for accommodating the "UNLIMITED" and "DEFAULT" values. 
* `password_lock_time_less_than` - (Optional) An optional filter to return the profiles having password lock number less than the provided value. String value is used for accommodating the "UNLIMITED" and "DEFAULT" values. 
* `password_verification_function` - (Optional) An optional filter to filter the profiles based on password verification function.
* `profile_name` - (Optional) A filter to return only items that match the specified profile name.
* `sessions_per_user_greater_than_or_equal` - (Optional) An optional filter to return the profiles permitting the user to spawn multiple sessions having count. greater than or equal to the provided value. String value is used for accommodating the "UNLIMITED" and "DEFAULT" values. 
* `sessions_per_user_less_than` - (Optional) An optional filter to return the profiles permitting the user to spawn multiple sessions having count less than the provided value. String value is used for accommodating the "UNLIMITED" and "DEFAULT" values. 
* `target_id` - (Optional) A filter to return only items related to a specific target OCID.
* `user_assessment_id` - (Required) The OCID of the user assessment.
* `user_count_greater_than_or_equal` - (Optional) An optional filter to return the profiles having user count greater than or equal to the provided value. 
* `user_count_less_than` - (Optional) An optional filter to return the profiles having user count less than the provided value. 


## Attributes Reference

The following attributes are exported:

* `profiles` - The list of profiles.

### UserAssessmentProfile Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the user assessment.
* `composite_limit` - Specify the total resource cost for a session, expressed in service units. Oracle Database calculates the total service units as a weighted sum of CPU_PER_SESSION, CONNECT_TIME, LOGICAL_READS_PER_SESSION, and PRIVATE_SGA. 
* `connect_time` - Specify the total elapsed time limit for a session, expressed in minutes.
* `cpu_per_call` - Specify the CPU time limit for a call (a parse, execute, or fetch), expressed in hundredths of seconds.
* `cpu_per_session` - Specify the CPU time limit for a session, expressed in hundredth of seconds.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `failed_login_attempts` - Maximum times the user is allowed in fail login before the user account is locked.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `idle_time` - Specify the permitted periods of continuous inactive time during a  session, expressed in minutes.
* `inactive_account_time` - The permitted periods of continuous inactive time during a session, expressed in minutes. Long-running queries and other operations are not subjected to this limit. 
* `is_user_created` - Represents if the profile is created by user.
* `logical_reads_per_call` - Specify the permitted the number of data blocks read for a call to process a SQL statement (a parse, execute, or fetch).
* `logical_reads_per_session` - Specify the permitted number of data blocks read in a session, including blocks read from memory and disk.
* `password_grace_time` - Number of grace days for user to change password.
* `password_life_time` - Number of days the password is valid before expiry.
* `password_lock_time` - Number of days the user account remains locked after failed login.
* `password_reuse_max` - Number of day after the user can use the already used password.
* `password_reuse_time` - Number of days before which a password cannot be reused.
* `password_rollover_time` - Number of days the password rollover is allowed. Minimum value can be 1/24 day (1 hour) to 60 days.
* `password_verification_function` - Name of the PL/SQL that can be used for password verification.
* `password_verification_function_details` - Details about the PL/SQL that can be used for password verification.
* `private_sga` - Specify the amount of private space a session can allocate in the shared pool of the system global area (SGA), expressed in bytes. 
* `profile_name` - The name of the profile.
* `sessions_per_user` - Specify the number of concurrent sessions to which you want to limit the user.
* `target_id` - The OCID of the target database.
* `user_assessment_id` - The OCID of the latest user assessment corresponding to the target under consideration. A compartment  type assessment can also be passed to profiles from all the targets from the corresponding compartment. 
* `user_count` - The number of users that have a given profile.

