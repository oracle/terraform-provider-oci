---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_compliance_policy"
sidebar_current: "docs-oci-datasource-fleet_apps_management-compliance_policy"
description: |-
  Provides details about a specific Compliance Policy in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_compliance_policy
This data source provides details about a specific Compliance Policy resource in Oracle Cloud Infrastructure Fleet Apps Management service.

Gets information about a CompliancePolicy.

## Example Usage

```hcl
data "oci_fleet_apps_management_compliance_policy" "test_compliance_policy" {
	#Required
	compliance_policy_id = oci_fleet_apps_management_compliance_policy.test_compliance_policy.id
}
```

## Argument Reference

The following arguments are supported:

* `compliance_policy_id` - (Required) unique CompliancePolicy identifier.


## Attributes Reference

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

