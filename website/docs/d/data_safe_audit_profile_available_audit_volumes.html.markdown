---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_audit_profile_available_audit_volumes"
sidebar_current: "docs-oci-datasource-data_safe-audit_profile_available_audit_volumes"
description: |-
  Provides the list of Audit Profile Available Audit Volumes in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_audit_profile_available_audit_volumes
This data source provides the list of Audit Profile Available Audit Volumes in Oracle Cloud Infrastructure Data Safe service.

Retrieves a list of audit trails, and associated audit event volume for each trail up to defined start date.

## Example Usage

```hcl
data "oci_data_safe_audit_profile_available_audit_volumes" "test_audit_profile_available_audit_volumes" {
	#Required
	audit_profile_id = oci_data_safe_audit_profile.test_audit_profile.id
	work_request_id = oci_containerengine_work_request.test_work_request.id

	#Optional
	month_in_consideration_greater_than = var.audit_profile_available_audit_volume_month_in_consideration_greater_than
	month_in_consideration_less_than = var.audit_profile_available_audit_volume_month_in_consideration_less_than
	trail_location = var.audit_profile_available_audit_volume_trail_location
}
```

## Argument Reference

The following arguments are supported:

* `audit_profile_id` - (Required) The OCID of the audit.
* `month_in_consideration_greater_than` - (Optional) Specifying `monthInConsiderationGreaterThan` parameter will retrieve all items for which the event month is greater than the date and time specified, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).

	**Example:** 2016-12-19T00:00:00.000Z 
* `month_in_consideration_less_than` - (Optional) Specifying `monthInConsiderationLessThan` parameter will retrieve all items for which the event month is less than the date and time specified, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).

	**Example:** 2016-12-19T00:00:00.000Z 
* `trail_location` - (Optional) The audit trail location.
* `work_request_id` - (Required) The OCID of the work request.


## Attributes Reference

The following attributes are exported:

* `available_audit_volume_collection` - The list of available_audit_volume_collection.

### AuditProfileAvailableAuditVolume Reference

The following attributes are exported:

* `items` - Array of available audit volume summary.
	* `audit_profile_id` - The OCID of the audit profile resource.
	* `audit_trail_id` - The OCID of the audit trail.
	* `database_unique_name` - Unique name of the database associated to the peer target database.
	* `month_in_consideration` - Represents the month under consideration for which aggregated audit data volume available at the target is computed. This field will be the UTC start of the day of the first day of the month for which the aggregate count corresponds to, in the format defined by RFC3339.. For instance, the value of 01-01-2021T00:00:00Z represents Jan 2021. 
	* `trail_location` - Audit trail location on the target database from where the audit data is being collected by Data Safe.
	* `volume` - Represents the aggregated audit data volume available in the audit trails on the target database which is yet to be collected by Data Safe for the specified month. 

