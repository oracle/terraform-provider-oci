---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_audit_profiles"
sidebar_current: "docs-oci-datasource-data_safe-audit_profiles"
description: |-
  Provides the list of Audit Profiles in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_audit_profiles
This data source provides the list of Audit Profiles in Oracle Cloud Infrastructure Data Safe service.

Gets a list of all audit profiles.

The ListAuditProfiles operation returns only the audit profiles in the specified `compartmentId`.
The list does not include any subcompartments of the compartmentId passed.

The parameter `accessLevel` specifies whether to return only those compartments for which the
requestor has INSPECT permissions on at least one resource directly
or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
Principal doesn't have access to even one of the child compartments. This is valid only when
`compartmentIdInSubtree` is set to `true`.

The parameter `compartmentIdInSubtree` applies when you perform ListAuditProfiles on the
`compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
To get a full list of all compartments and subcompartments in the tenancy (root compartment),
set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.


## Example Usage

```hcl
data "oci_data_safe_audit_profiles" "test_audit_profiles" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.audit_profile_access_level
	audit_collected_volume_greater_than_or_equal_to = var.audit_profile_audit_collected_volume_greater_than_or_equal_to
	audit_profile_id = oci_data_safe_audit_profile.test_audit_profile.id
	compartment_id_in_subtree = var.audit_profile_compartment_id_in_subtree
	display_name = var.audit_profile_display_name
	is_override_global_retention_setting = var.audit_profile_is_override_global_retention_setting
	is_paid_usage_enabled = var.audit_profile_is_paid_usage_enabled
	state = var.audit_profile_state
	target_id = oci_cloud_guard_target.test_target.id
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `audit_collected_volume_greater_than_or_equal_to` - (Optional) A filter to return only items that have count of audit records collected greater than or equal to the specified value.
* `audit_profile_id` - (Optional) A optional filter to return only resources that match the specified id.
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `display_name` - (Optional) A filter to return only resources that match the specified display name. 
* `is_override_global_retention_setting` - (Optional) A optional filter to return only resources that match the specified retention configured value.
* `is_paid_usage_enabled` - (Optional) Indicates if you want to continue audit record collection beyond the free limit of one million audit records per month per target database, incurring additional charges. The default value is inherited from the global settings. You can change at the global level or at the target level. 
* `state` - (Optional) A optional filter to return only resources that match the specified lifecycle state.
* `target_id` - (Optional) A filter to return only items related to a specific target OCID.


## Attributes Reference

The following attributes are exported:

* `audit_profile_collection` - The list of audit_profile_collection.

### AuditProfile Reference

The following attributes are exported:

* `audit_collected_volume` - Indicates number of audit records collected by Data Safe in the current calendar month.  Audit records for the Data Safe service account are excluded and are not counted towards your monthly free limit. 
* `audit_trails` - Indicates the list of available audit trails on the target.
	* `audit_collection_start_time` - The date from which the audit trail must start collecting data, in the format defined by RFC3339.
	* `audit_profile_id` - The OCID of the  parent audit.
	* `compartment_id` - The OCID of the compartment that contains the audit trail and is the same as the compartment of the audit profile resource. 
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
	* `description` - The description of the audit trail.
	* `display_name` - The display name of the audit trail.
	* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
	* `id` - The OCID of the audit trail.
	* `is_auto_purge_enabled` - Indicates if auto purge is enabled on the target database, which helps delete audit data in the target database every seven days so that the database's audit trail does not become too large. 
	* `lifecycle_details` - Details about the current state of the audit trail in Data Safe.
	* `peer_target_database_key` - The secondary id assigned for the peer database registered with Data Safe.
	* `purge_job_details` - The details of the audit trail purge job that ran on the "purgeJobTime".
	* `purge_job_status` - The current status of the audit trail purge job.
	* `purge_job_time` - The date and time of the last purge job, which deletes audit data in the target database every seven days so that the database's audit trail does not become too large. In the format defined by RFC3339. 
	* `state` - The current state of the audit trail.
	* `status` - The current sub-state of the audit trail.
	* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
	* `target_id` - The OCID of the Data Safe target for which the audit trail is created.
	* `time_created` - The date and time the audit trail was created, in the format defined by RFC3339.
	* `time_last_collected` - The date and time until when the audit events were collected from the target database by the Data Safe audit trail  collection process, in the format defined by RFC3339. 
	* `time_updated` - The date and time the audit trail was updated, in the format defined by RFC3339.
	* `trail_location` - An audit trail location represents the source of audit records that provides documentary evidence of the sequence of activities in the target database. 
	* `trail_source` - The underlying source of unified audit trail.
	* `work_request_id` - The OCID of the workrequest for audit trail which collects audit records.
* `compartment_id` - The OCID of the compartment that contains the audit.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the audit profile.
* `display_name` - The display name of the audit profile.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the audit profile.
* `is_override_global_retention_setting` - Indicates whether audit retention settings like online and offline months is set at the target level overriding the global audit retention settings. 
* `is_paid_usage_enabled` - Indicates if you want to continue collecting audit records beyond the free limit of one million audit records per month per target database, potentially incurring additional charges. The default value is inherited from the global settings.  You can change at the global level or at the target level. 
* `lifecycle_details` - Details about the current state of the audit profile in Data Safe.
* `offline_months` - Indicates the number of months the audit records will be stored offline in the Data Safe audit archive. Minimum: 0; Maximum: 72 months. If you have a requirement to store the audit data even longer in archive, please contact the Oracle Support. 
* `online_months` - Indicates the number of months the audit records will be stored online in Oracle Data Safe audit repository for immediate reporting and analysis.  Minimum: 1; Maximum:12 months 
* `state` - The current state of the audit profile.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `target_id` - The OCID of the Data Safe target for which the audit profile is created.
* `time_created` - The date and time the audit profile was created, in the format defined by RFC3339.
* `time_updated` - The date and time the audit profile was updated, in the format defined by RFC3339.

