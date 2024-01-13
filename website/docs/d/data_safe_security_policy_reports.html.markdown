---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_security_policy_reports"
sidebar_current: "docs-oci-datasource-data_safe-security_policy_reports"
description: |-
  Provides the list of Security Policy Reports in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_security_policy_reports
This data source provides the list of Security Policy Reports in Oracle Cloud Infrastructure Data Safe service.

Retrieves a list of all security policy reports in Data Safe.

The ListSecurityPolicyReports operation returns only the security policy reports in the specified `compartmentId`.

The parameter `accessLevel` specifies whether to return only those compartments for which the
requestor has INSPECT permissions on at least one resource directly
or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
Principal doesn't have access to even one of the child compartments. This is valid only when
`compartmentIdInSubtree` is set to `true`.

The parameter `compartmentIdInSubtree` applies when you perform ListSecurityPolicyReports on the
`compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
To get a full list of all compartments and subcompartments in the tenancy (root compartment),
set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.


## Example Usage

```hcl
data "oci_data_safe_security_policy_reports" "test_security_policy_reports" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.security_policy_report_access_level
	compartment_id_in_subtree = var.security_policy_report_compartment_id_in_subtree
	display_name = var.security_policy_report_display_name
	security_policy_report_id = oci_data_safe_security_policy_report.test_security_policy_report.id
	state = var.security_policy_report_state
	target_id = oci_cloud_guard_target.test_target.id
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `display_name` - (Optional) A filter to return only resources that match the specified display name. 
* `security_policy_report_id` - (Optional) An optional filter to return only resources that match the specified OCID of the security policy report resource.
* `state` - (Optional) The current state of the security policy report.
* `target_id` - (Optional) A filter to return only items related to a specific target OCID.


## Attributes Reference

The following attributes are exported:

* `security_policy_report_collection` - The list of security_policy_report_collection.

### SecurityPolicyReport Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the security policy report.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the security policy report.
* `display_name` - The display name of the security policy report.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the security policy report.
* `lifecycle_details` - Details about the current state of the security policy report.
* `state` - The current state of the security policy report.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `target_id` - The OCID of the of the  target database.
* `time_created` - The date and time the security policy report was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_updated` - The date and time the security policy report was last updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).

