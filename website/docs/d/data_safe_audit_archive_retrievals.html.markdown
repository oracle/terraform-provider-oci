---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_audit_archive_retrievals"
sidebar_current: "docs-oci-datasource-data_safe-audit_archive_retrievals"
description: |-
  Provides the list of Audit Archive Retrievals in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_audit_archive_retrievals
This data source provides the list of Audit Archive Retrievals in Oracle Cloud Infrastructure Data Safe service.

Returns the list of audit archive retrieval.


## Example Usage

```hcl
data "oci_data_safe_audit_archive_retrievals" "test_audit_archive_retrievals" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.audit_archive_retrieval_access_level
	audit_archive_retrieval_id = oci_data_safe_audit_archive_retrieval.test_audit_archive_retrieval.id
	compartment_id_in_subtree = var.audit_archive_retrieval_compartment_id_in_subtree
	display_name = var.audit_archive_retrieval_display_name
	state = var.audit_archive_retrieval_state
	target_id = oci_cloud_guard_target.test_target.id
	time_of_expiry = var.audit_archive_retrieval_time_of_expiry
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `audit_archive_retrieval_id` - (Optional) OCID of the archive retrieval.
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `display_name` - (Optional) A filter to return only resources that match the specified display name. 
* `state` - (Optional) A filter to return only resources that matches the specified lifecycle state.
* `target_id` - (Optional) The OCID of the target associated with the archive retrieval.
* `time_of_expiry` - (Optional) The date time when retrieved archive data will be deleted from Data Safe and unloaded back into archival.


## Attributes Reference

The following attributes are exported:

* `audit_archive_retrieval_collection` - The list of audit_archive_retrieval_collection.

### AuditArchiveRetrieval Reference

The following attributes are exported:

* `audit_event_count` - Total count of audit events to be retrieved from the archive for the specified date range.
* `compartment_id` - The OCID of the compartment that contains archive retrieval.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - Description of the archive retrieval.
* `display_name` - The display name of the archive retrieval. The name does not have to be unique, and is changeable.
* `end_date` - End month of the archive retrieval, in the format defined by RFC3339.
* `error_info` - The Error details of a failed archive retrieval.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the archive retrieval.
* `lifecycle_details` - Details about the current state of the archive retrieval.
* `start_date` - Start month of the archive retrieval, in the format defined by RFC3339.
* `state` - The current state of the archive retrieval.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `target_id` - The OCID of the target associated with the archive retrieval.
* `time_completed` - The date time when archive retrieval request was fulfilled, in the format defined by RFC3339.
* `time_of_expiry` - The date time when retrieved archive data will be deleted from Data Safe and unloaded back into archival.
* `time_requested` - The date time when archive retrieval was requested, in the format defined by RFC3339.

