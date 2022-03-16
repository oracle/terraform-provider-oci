---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_audit_profile_collected_audit_volume"
sidebar_current: "docs-oci-datasource-data_safe-audit_profile_collected_audit_volume"
description: |-
  Provides details about a specific Audit Profile Collected Audit Volume in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_audit_profile_collected_audit_volume
This data source provides details about a specific Audit Profile Collected Audit Volume resource in Oracle Cloud Infrastructure Data Safe service.

Gets a list of all collected audit volume data points.

## Example Usage

```hcl
data "oci_data_safe_audit_profile_collected_audit_volume" "test_audit_profile_collected_audit_volume" {
	#Required
	audit_profile_id = oci_data_safe_audit_profile.test_audit_profile.id
	work_request_id = oci_containerengine_work_request.test_work_request.id

	#Optional
	month_in_consideration_greater_than = var.audit_profile_collected_audit_volume_month_in_consideration_greater_than
	month_in_consideration_less_than = var.audit_profile_collected_audit_volume_month_in_consideration_less_than
}
```

## Argument Reference

The following arguments are supported:

* `audit_profile_id` - (Required) The OCID of the audit.
* `month_in_consideration_greater_than` - (Optional) Specifying `monthInConsiderationGreaterThan` parameter will retrieve all items for which the event month is greater than the date and time specified, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).

	**Example:** 2016-12-19T00:00:00.000Z 
* `month_in_consideration_less_than` - (Optional) Specifying `monthInConsiderationLessThan` parameter will retrieve all items for which the event month is less than the date and time specified, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).

	**Example:** 2016-12-19T00:00:00.000Z 
* `work_request_id` - (Required) The OCID of the work request.


## Attributes Reference

The following attributes are exported:

* `items` - Array of collected audit volume summary.
	* `archived_volume` - The audit data volume collected by Data Safe and is available in archive storage.
	* `audit_profile_id` - The OCID of the audit profile resource.
	* `month_in_consideration` - Represents the month under consideration in which the aggregated audit data volume collected by Data Safe is displayed. This field will be the UTC start of the day of the first day of the month for which the aggregate count corresponds to, in the format defined by RFC3339.. For instance, the value of 01-01-2021T00:00:00Z represents Jan 2021. 
	* `online_volume` - The audit data volume collected by Data Safe and is available online in repository.

