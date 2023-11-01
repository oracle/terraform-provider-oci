---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_security_policy"
sidebar_current: "docs-oci-datasource-data_safe-security_policy"
description: |-
  Provides details about a specific Security Policy in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_security_policy
This data source provides details about a specific Security Policy resource in Oracle Cloud Infrastructure Data Safe service.

Gets a security policy by the specified OCID of the security policy resource.

## Example Usage

```hcl
data "oci_data_safe_security_policy" "test_security_policy" {
	#Required
	security_policy_id = oci_data_safe_security_policy.test_security_policy.id
}
```

## Argument Reference

The following arguments are supported:

* `security_policy_id` - (Required) The OCID of the security policy resource.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the security policy.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the security policy.
* `display_name` - The display name of the security policy.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the security policy.
* `lifecycle_details` - Details about the current state of the security policy in Data Safe.
* `state` - The current state of the security policy.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time that the security policy was created, in the format defined by RFC3339.
* `time_updated` - The last date and time the security policy was updated, in the format defined by RFC3339.

