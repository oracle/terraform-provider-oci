---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_drg_nat_policies"
sidebar_current: "docs-oci-datasource-core-drg_nat_policies"
description: |-
  Provides the list of DrgNatPolicies in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_drg_nat_policies
This data source provides the list of DrgNatPolicies in Oracle Cloud Infrastructure Core service.

The list of DrgNatPolicies in the compartment.


## Example Usage

```hcl
data "oci_core_drg_nat_policies" "test_drg_nat_policies" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.drg_nat_policy_display_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly.


## Attributes Reference

The following attributes are exported:

* `drg_nat_policies` - The list of drg_nat_policies.

### DrgNatPolicy Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the DrgNatPolicy.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
* `id` - The DrgNatPolicy's Oracle ID ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)).
* `state` - The DrgNatPolicy's current state.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `time_created` - The date and time the DrgNatPolicy was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
