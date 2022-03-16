---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_audit_archive_retrieval"
sidebar_current: "docs-oci-datasource-data_safe-audit_archive_retrieval"
description: |-
  Provides details about a specific Audit Archive Retrieval in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_audit_archive_retrieval
This data source provides details about a specific Audit Archive Retrieval resource in Oracle Cloud Infrastructure Data Safe service.

Gets the details of the specified archive retreival.

## Example Usage

```hcl
data "oci_data_safe_audit_archive_retrieval" "test_audit_archive_retrieval" {
	#Required
	audit_archive_retrieval_id = oci_data_safe_audit_archive_retrieval.test_audit_archive_retrieval.id
}
```

## Argument Reference

The following arguments are supported:

* `audit_archive_retrieval_id` - (Required) OCID of the archive retrieval.


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

