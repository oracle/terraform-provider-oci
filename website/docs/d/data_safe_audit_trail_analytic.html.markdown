---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_audit_trail_analytic"
sidebar_current: "docs-oci-datasource-data_safe-audit_trail_analytic"
description: |-
  Provides details about a specific Audit Trail Analytic in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_audit_trail_analytic
This data source provides details about a specific Audit Trail Analytic resource in Oracle Cloud Infrastructure Data Safe service.

Gets a list of audit trail aggregated details . A audit trail aggregation helps understand the overall  state of trails.
As an example, it helps understand how many trails are running or stopped. It is especially useful to create dashboards or to support analytics.

The parameter `accessLevel` specifies whether to return only those compartments for which the
requestor has INSPECT permissions on at least one resource directly
or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
Principal doesn't have access to even one of the child compartments. This is valid only when
`compartmentIdInSubtree` is set to `true`.

The parameter `compartmentIdInSubtree` applies when you perform AuditTrailAnalytics on the
`compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
To get a full list of all compartments and subcompartments in the tenancy (root compartment),
set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.


## Example Usage

```hcl
data "oci_data_safe_audit_trail_analytic" "test_audit_trail_analytic" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.audit_trail_analytic_access_level
	compartment_id_in_subtree = var.audit_trail_analytic_compartment_id_in_subtree
	group_by = var.audit_trail_analytic_group_by
	target_id = oci_cloud_guard_target.test_target.id
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `group_by` - (Optional) The group by parameter for summarize operation on audit trail.
* `target_id` - (Optional) A filter to return only items related to a specific target OCID.


## Attributes Reference

The following attributes are exported:

* `items` - Array of audit trail aggregration data.
	* `count` - Total count of aggregated metric.
	* `dimensions` - Details of aggregation dimensions used for summarizing audit trails.
		* `location` - The location represents the source of audit records that provides documentary evidence of the sequence of activities in the target database.
		* `state` - The current state of the audit trail.
		* `status` - The current sub-state of the audit trail..
		* `target_id` - The OCID of the Data Safe target for which the audit trail is created.

