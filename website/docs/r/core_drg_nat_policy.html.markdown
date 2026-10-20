---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_drg_nat_policy"
sidebar_current: "docs-oci-resource-core-drg_nat_policy"
description: |-
  Provides the Drg Nat Policy resource in Oracle Cloud Infrastructure Core service
---

# oci_core_drg_nat_policy
This resource provides the Drg Nat Policy resource in Oracle Cloud Infrastructure Core service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/iaas/latest/DrgNatPolicy

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/

Creates a new DRG NAT policy. Assign the DRG NAT policy to a DRG attachment
using the `UpdateDrgAttachment` or `CreateDrgAttachment` operations.


## Example Usage

```hcl
resource "oci_core_drg_nat_policy" "test_drg_nat_policy" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.drg_nat_policy_display_name
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the DRG NAT policy. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the DRG NAT policy.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The DRG NAT policy's Oracle ID ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)).
* `state` - The DRG NAT policy's current state.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the DRG NAT policy was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Drg Nat Policy
	* `update` - (Defaults to 20 minutes), when updating the Drg Nat Policy
	* `delete` - (Defaults to 20 minutes), when destroying the Drg Nat Policy


## Import

DrgNatPolicies can be imported using the `id`, e.g.

```
$ terraform import oci_core_drg_nat_policy.test_drg_nat_policy "id"
```

