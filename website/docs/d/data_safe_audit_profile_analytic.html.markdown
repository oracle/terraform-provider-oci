---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_audit_profile_analytic"
sidebar_current: "docs-oci-datasource-data_safe-audit_profile_analytic"
description: |-
  Provides details about a specific Audit Profile Analytic in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_audit_profile_analytic
This data source provides details about a specific Audit Profile Analytic resource in Oracle Cloud Infrastructure Data Safe service.

Gets a list of audit profile aggregated details . A audit profile  aggregation helps understand the overall  state of audit profile profiles.
As an example, it helps understand how many audit profiles have paid usage. It is especially useful to create dashboards or to support analytics.

The parameter `accessLevel` specifies whether to return only those compartments for which the
requestor has INSPECT permissions on at least one resource directly
or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
Principal doesn't have access to even one of the child compartments. This is valid only when
`compartmentIdInSubtree` is set to `true`.

The parameter `compartmentIdInSubtree` applies when you perform AuditProfileAnalytics on the
`compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
To get a full list of all compartments and subcompartments in the tenancy (root compartment),
set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.


## Example Usage

```hcl
data "oci_data_safe_audit_profile_analytic" "test_audit_profile_analytic" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.audit_profile_analytic_access_level
	compartment_id_in_subtree = var.audit_profile_analytic_compartment_id_in_subtree
	group_by = var.audit_profile_analytic_group_by
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `group_by` - (Optional) The group by parameter for summarize operation on audit.


## Attributes Reference

The following attributes are exported:

* `items` - Array of audit profile aggregration data.
	* `count` - Total count of aggregated metric.
	* `dimensions` - Details of aggregation dimensions used for summarizing audit profiles.
		* `is_paid_usage_enabled` - Indicates if you want to continue collecting audit records beyond the free limit of one million audit records per month per target database, potentially incurring additional charges. The default value is inherited from the global settings.  You can change at the global level or at the target level. 

