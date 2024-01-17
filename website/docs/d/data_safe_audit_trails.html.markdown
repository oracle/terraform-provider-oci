---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_audit_trails"
sidebar_current: "docs-oci-datasource-data_safe-audit_trails"
description: |-
  Provides the list of Audit Trails in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_audit_trails
This data source provides the list of Audit Trails in Oracle Cloud Infrastructure Data Safe service.

Gets a list of all audit trails.
The ListAuditTrails operation returns only the audit trails in the specified `compartmentId`.
The list does not include any subcompartments of the compartmentId passed.

The parameter `accessLevel` specifies whether to return only those compartments for which the
requestor has INSPECT permissions on at least one resource directly
or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
Principal doesn't have access to even one of the child compartments. This is valid only when
`compartmentIdInSubtree` is set to `true`.

The parameter `compartmentIdInSubtree` applies when you perform ListAuditTrails on the
`compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
To get a full list of all compartments and subcompartments in the tenancy (root compartment),
set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.


## Example Usage

```hcl
data "oci_data_safe_audit_trails" "test_audit_trails" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.audit_trail_access_level
	audit_trail_id = oci_data_safe_audit_trail.test_audit_trail.id
	compartment_id_in_subtree = var.audit_trail_compartment_id_in_subtree
	display_name = var.audit_trail_display_name
	state = var.audit_trail_state
	status = var.audit_trail_status
	target_id = oci_cloud_guard_target.test_target.id
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `audit_trail_id` - (Optional) A optional filter to return only resources that match the specified id.
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `display_name` - (Optional) A filter to return only resources that match the specified display name. 
* `state` - (Optional) A optional filter to return only resources that match the specified lifecycle state.
* `status` - (Optional) A optional filter to return only resources that match the specified sub-state of audit trail.
* `target_id` - (Optional) A filter to return only items related to a specific target OCID.


## Attributes Reference

The following attributes are exported:

* `audit_trail_collection` - The list of audit_trail_collection.

### AuditTrail Reference

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

