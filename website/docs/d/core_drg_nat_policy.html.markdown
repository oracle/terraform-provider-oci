---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_drg_nat_policy"
sidebar_current: "docs-oci-datasource-core-drg_nat_policy"
description: |-
  Provides details about a specific Drg Nat Policy in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_drg_nat_policy
This data source provides details about a specific Drg Nat Policy resource in Oracle Cloud Infrastructure Core service.

Gets the specified DRG NAT policy's information.

## Example Usage

```hcl
data "oci_core_drg_nat_policy" "test_drg_nat_policy" {
	#Required
	drg_nat_policy_id = oci_core_drg_nat_policy.test_drg_nat_policy.id
}
```

## Argument Reference

The following arguments are supported:

* `drg_nat_policy_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG NAT policy.


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

