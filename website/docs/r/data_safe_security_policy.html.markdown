---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_security_policy"
sidebar_current: "docs-oci-resource-data_safe-security_policy"
description: |-
  Provides the Security Policy resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_security_policy
This resource provides the Security Policy resource in Oracle Cloud Infrastructure Data Safe service.

Updates the security policy.

## Example Usage

```hcl
resource "oci_data_safe_security_policy" "test_security_policy" {
	#Required
	security_policy_id = oci_data_safe_security_policy.test_security_policy.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.security_policy_description
	display_name = var.security_policy_display_name
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) (Updatable) The OCID of the compartment containing the security policy.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) The description of the security policy.
* `display_name` - (Optional) (Updatable) The display name of the security policy. The name does not have to be unique, and it is changeable.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `security_policy_id` - (Required) The OCID of the security policy resource.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Security Policy
	* `update` - (Defaults to 20 minutes), when updating the Security Policy
	* `delete` - (Defaults to 20 minutes), when destroying the Security Policy


## Import

SecurityPolicies can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_security_policy.test_security_policy "id"
```

