---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_compliance_policies"
sidebar_current: "docs-oci-datasource-fleet_apps_management-compliance_policies"
description: |-
  Provides the list of Compliance Policies in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_compliance_policies
This data source provides the list of Compliance Policies in Oracle Cloud Infrastructure Fleet Apps Management service.

Gets a list of compliancePolicies.


## Example Usage

```hcl
data "oci_fleet_apps_management_compliance_policies" "test_compliance_policies" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.compliance_policy_display_name
	id = var.compliance_policy_id
	state = var.compliance_policy_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `id` - (Optional) unique CompliancePolicy identifier.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state. The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `compliance_policy_collection` - The list of compliance_policy_collection.

### CompliancePolicy Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment the CompliancePolicy belongs to.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Display name for the CompliancePolicy.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the CompliancePolicy.
* `lifecycle_details` - A message that describes the current state of the CompliancePolicy in more detail. For example, can be used to provide actionable information for a resource in the Failed state. 
* `product_id` - platformConfiguration OCID corresponding to the Product.
* `state` - The current state of the CompliancePolicy.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the CompliancePolicy was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the CompliancePolicy was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

