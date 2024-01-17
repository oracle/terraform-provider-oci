---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_audit_trail"
sidebar_current: "docs-oci-resource-data_safe-audit_trail"
description: |-
  Provides the Audit Trail resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_audit_trail
This resource provides the Audit Trail resource in Oracle Cloud Infrastructure Data Safe service.

Updates one or more attributes of the specified audit trail.

## Example Usage

```hcl
resource "oci_data_safe_audit_trail" "test_audit_trail" {
	#Required
	audit_trail_id = oci_data_safe_audit_trail.test_audit_trail.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.audit_trail_description
	display_name = var.audit_trail_display_name
	freeform_tags = {"Department"= "Finance"}
	is_auto_purge_enabled = var.audit_trail_is_auto_purge_enabled
}
```

## Argument Reference

The following arguments are supported:

* `audit_trail_id` - (Required) The OCID of the audit trail.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) The description of the audit trail.
* `display_name` - (Optional) (Updatable) The display name of the audit trail. The name does not have to be unique, and it's changeable.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `is_auto_purge_enabled` - (Optional) (Updatable) Indicates if auto purge is enabled on the target database, which helps delete audit data in the target database every seven days so that the database's audit trail does not become too large. 
* `state` - (Optional) (Updatable) The target state for the Audit Trail. Could be set to `ACTIVE` or `INACTIVE`. 
* `resume_trigger` - (Optional) (Updatable) An optional property when incremented triggers Resume. Could be set to any integer value.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Audit Trail
	* `update` - (Defaults to 20 minutes), when updating the Audit Trail
	* `delete` - (Defaults to 20 minutes), when destroying the Audit Trail


## Import

AuditTrails can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_audit_trail.test_audit_trail "id"
```

