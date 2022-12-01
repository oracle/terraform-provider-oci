---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_masking_reports"
sidebar_current: "docs-oci-datasource-data_safe-masking_reports"
description: |-
  Provides the list of Masking Reports in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_masking_reports
This data source provides the list of Masking Reports in Oracle Cloud Infrastructure Data Safe service.

Gets a list of masking reports based on the specified query parameters.

## Example Usage

```hcl
data "oci_data_safe_masking_reports" "test_masking_reports" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.masking_report_access_level
	compartment_id_in_subtree = var.masking_report_compartment_id_in_subtree
	masking_policy_id = oci_data_safe_masking_policy.test_masking_policy.id
	target_id = oci_cloud_guard_target.test_target.id
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `masking_policy_id` - (Optional) A filter to return only the resources that match the specified masking policy OCID.
* `target_id` - (Optional) A filter to return only items related to a specific target OCID.


## Attributes Reference

The following attributes are exported:

* `masking_report_collection` - The list of masking_report_collection.

### MaskingReport Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the masking report.
* `id` - The OCID of the masking report.
* `masking_policy_id` - The OCID of the masking policy used.
* `masking_work_request_id` - The OCID of the masking work request that resulted in this masking report.
* `state` - The current state of the masking report.
* `target_id` - The OCID of the target database masked.
* `time_created` - The date and time the masking report was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 
* `time_masking_finished` - The date and time data masking finished, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339)
* `time_masking_started` - The date and time data masking started, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339)
* `total_masked_columns` - The total number of masked columns.
* `total_masked_objects` - The total number of unique objects (tables and editioning views) that contain the masked columns.
* `total_masked_schemas` - The total number of unique schemas that contain the masked columns.
* `total_masked_sensitive_types` - The total number of unique sensitive types associated with the masked columns.
* `total_masked_values` - The total number of masked values.

