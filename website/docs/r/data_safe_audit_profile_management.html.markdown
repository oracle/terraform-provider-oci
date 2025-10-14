---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_audit_profile_management"
sidebar_current: "docs-oci-resource-data_safe-audit_profile_management"
description: |-
  Provides the Audit Profile Management resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_audit_profile_management
This resource provides the Audit Profile Management resource in Oracle Cloud Infrastructure Data Safe service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/data-safe/latest/AuditProfileManagement

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/datasafe

Create a new audit profile resource for a target group. For a target database, it will update the auto created audit profile by fetching the Audit profile.

## Example Usage

```hcl
resource "oci_data_safe_audit_profile_management" "test_audit_profile_management" {
	#Required
	compartment_id = var.compartment_id
	target_id = oci_cloud_guard_target.test_target.id
	target_type = var.audit_profile_target_type

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.audit_profile_description
	display_name = var.audit_profile_display_name
	freeform_tags = {"Department"= "Finance"}
	is_override_global_paid_usage = var.audit_profile_is_override_global_paid_usage
	is_paid_usage_enabled = var.audit_profile_is_paid_usage_enabled
	offline_months = var.audit_profile_offline_months
	online_months = var.audit_profile_online_months
	change_retention_trigger= var.retention_trigger
    is_override_global_retention_setting = true
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment where you want to create the audit profile.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm) Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) The description of the audit profile.
* `display_name` - (Optional) (Updatable) The display name of the audit profile. The name does not have to be unique, and it's updatable.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `is_override_global_paid_usage` - (Optional) Indicates whether audit paid usage settings specified at the target database level override both the global and the target database group level paid usage settings. Enabling paid usage continues the collection of audit records beyond the free limit of one million audit records per month per target database, potentially incurring additional charges. For more information, see [Data Safe Price List](https://www.oracle.com/cloud/price-list/#data-safe). 
* `is_paid_usage_enabled` - (Optional) (Updatable) Indicates if you want to continue collecting audit records beyond the free limit of one million audit records per month per target database, potentially incurring additional charges. The default value is inherited from the global settings.  You can change at the global level or at the target level. 
* `offline_months` - (Optional) Number of months the audit records will be stored offline in the offline archive. Minimum: 0; Maximum: 72 months. If you have a requirement to store the audit data even longer in the offline archive, please contact the Oracle Support. 
* `online_months` - (Optional) Number of months the audit records will be stored online in the audit repository for immediate reporting and analysis. Minimum: 1; Maximum: 12 months 
* `target_id` - (Required) The OCID of the target database or target database group for which the audit profile is created.
* `target_type` - (Required) The resource type that is represented by the audit profile.
* `change_retention_trigger` - (Optional) (Updatable) An optional property when incremented triggers Change Retention. Could be set to any integer value.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `audit_collected_volume` - Number of audit records collected in the current calendar month.  Audit records for the Data Safe service account are excluded and are not counted towards your monthly free limit. 
* `audit_trails` - Contains the list of available audit trails on the target database.
	* `audit_collection_start_time` - The date from which the audit trail must start collecting data, in the format defined by RFC3339.
	* `audit_profile_id` - The OCID of the  parent audit.
	* `can_update_last_archive_time_on_target` - Indicates if the Datasafe updates last archive time on target database. If isAutoPurgeEnabled field is enabled, this field must be true. 
	* `compartment_id` - The OCID of the compartment that contains the audit trail and is the same as the compartment of the audit profile resource. 
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm) Example: `{"Operations.CostCenter": "42"}` 
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
* `compartment_id` - The OCID of the compartment that contains the audit profile.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm) Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the audit profile.
* `display_name` - The display name of the audit profile.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the audit profile.
* `is_override_global_retention_setting` - Indicates whether audit retention settings like online and offline months set at the  target level override both the global settings and the target group level audit retention settings. 
* `is_paid_usage_enabled` - Indicates if you want to continue collecting audit records beyond the free limit of one million audit records per month per target database, potentially incurring additional charges. The default value is inherited from the global settings.  You can change at the global level or at the target level. 
* `lifecycle_details` - Details about the current state of the audit profile in Data Safe.
* `offline_months` - Number of months the audit records will be stored offline in the offline archive. Minimum: 0; Maximum: 72 months. If you have a requirement to store the audit data even longer in the offline archive, please contact the Oracle Support. 
* `offline_months_source` - The name or the OCID of the resource from which the offline month retention setting is sourced. For example, a global setting or a target database group OCID.
* `online_months` - Number of months the audit records will be stored online in the audit repository for immediate reporting and analysis.  Minimum: 1; Maximum: 12 months 
* `online_months_source` - The name or the OCID of the resource from which the online month retention setting is sourced. For example, a global setting or a target database group OCID.
* `paid_usage_source` - The name or the OCID of the resource from which the paid usage setting is sourced. For example, a global setting or a target database group OCID.
* `state` - The current state of the audit profile.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `target_id` - The OCID of the target database or target database group for which the audit profile is created.
* `target_type` - The resource type that is represented by the audit profile.
* `time_created` - The date and time the audit profile was created, in the format defined by RFC3339.
* `time_updated` - The date and time the audit profile was updated, in the format defined by RFC3339.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Audit Profile
	* `update` - (Defaults to 20 minutes), when updating the Audit Profile
	* `delete` - (Defaults to 20 minutes), when destroying the Audit Profile


## Import

AuditProfiles can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_audit_profile_management_.test_audit_profile_management "id"
```

