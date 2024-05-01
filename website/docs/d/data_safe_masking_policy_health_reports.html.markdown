---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_masking_policy_health_reports"
sidebar_current: "docs-oci-datasource-data_safe-masking_policy_health_reports"
description: |-
  Provides the list of Masking Policy Health Reports in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_masking_policy_health_reports
This data source provides the list of Masking Policy Health Reports in Oracle Cloud Infrastructure Data Safe service.

Gets a list of masking policy health reports based on the specified query parameters.

## Example Usage

```hcl
data "oci_data_safe_masking_policy_health_reports" "test_masking_policy_health_reports" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.masking_policy_health_report_access_level
	compartment_id_in_subtree = var.masking_policy_health_report_compartment_id_in_subtree
	display_name = var.masking_policy_health_report_display_name
	masking_policy_health_report_id = oci_data_safe_masking_policy_health_report.test_masking_policy_health_report.id
	masking_policy_id = oci_data_safe_masking_policy.test_masking_policy.id
	state = var.masking_policy_health_report_state
	target_id = oci_cloud_guard_target.test_target.id
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `display_name` - (Optional) A filter to return only resources that match the specified display name. 
* `masking_policy_health_report_id` - (Optional) A filter to return only the resources that match the specified masking policy health report OCID.
* `masking_policy_id` - (Optional) A filter to return only the resources that match the specified masking policy OCID.
* `state` - (Optional) A filter to return only the resources that match the specified lifecycle states.
* `target_id` - (Optional) A filter to return only items related to a specific target OCID.


## Attributes Reference

The following attributes are exported:

* `masking_policy_health_report_collection` - The list of masking_policy_health_report_collection.

### MaskingPolicyHealthReport Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the health report.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the masking health report.
* `display_name` - The display name of the health report.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the health report.
* `masking_policy_id` - The OCID of the masking policy.
* `state` - The current state of the health report.
* `target_id` - The OCID of the target database for which this report was created.
* `time_created` - The date and time the report was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 
* `time_updated` - The date and time the report was last updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339)  

