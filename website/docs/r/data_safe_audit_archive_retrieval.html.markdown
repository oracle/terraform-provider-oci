---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_audit_archive_retrieval"
sidebar_current: "docs-oci-resource-data_safe-audit_archive_retrieval"
description: |-
  Provides the Audit Archive Retrieval resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_audit_archive_retrieval
This resource provides the Audit Archive Retrieval resource in Oracle Cloud Infrastructure Data Safe service.

Creates a work request to retrieve archived audit data. This asynchronous process will usually take over an hour to complete.
Save the id from the response of this operation. Call GetAuditArchiveRetrieval operation after an hour, passing the id to know the status of
this operation.


## Example Usage

```hcl
resource "oci_data_safe_audit_archive_retrieval" "test_audit_archive_retrieval" {
	#Required
	compartment_id = var.compartment_id
	end_date = var.audit_archive_retrieval_end_date
	start_date = var.audit_archive_retrieval_start_date
	target_id = oci_cloud_guard_target.test_target.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.audit_archive_retrieval_description
	display_name = var.audit_archive_retrieval_display_name
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment that contains the archival retrieval.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) Description of the archive retrieval.
* `display_name` - (Optional) (Updatable) The display name of the archive retrieval. The name does not have to be unique, and is changeable.
* `end_date` - (Required) End month of the archive retrieval, in the format defined by RFC3339.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `start_date` - (Required) Start month of the archive retrieval, in the format defined by RFC3339.
* `target_id` - (Required) The OCID of the target associated with the archive retrieval.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Audit Archive Retrieval
	* `update` - (Defaults to 20 minutes), when updating the Audit Archive Retrieval
	* `delete` - (Defaults to 20 minutes), when destroying the Audit Archive Retrieval


## Import

AuditArchiveRetrievals can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_audit_archive_retrieval.test_audit_archive_retrieval "id"
```

