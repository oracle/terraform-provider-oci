---
subcategory: "Opsi"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opsi_chargeback_plan"
sidebar_current: "docs-oci-resource-opsi-chargeback_plan"
description: |-
  Provides the Chargeback Plan resource in Oracle Cloud Infrastructure Opsi service
---

# oci_opsi_chargeback_plan
This resource provides the Chargeback Plan resource in Oracle Cloud Infrastructure Opsi service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/operations-insights/latest/ChargebackPlan

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/osi

Create a chargeback plan resource for the resource in Ops Insights.


## Example Usage

```hcl
resource "oci_opsi_chargeback_plan" "test_chargeback_plan" {
	#Required
	compartment_id = var.compartment_id
	entity_source = var.chargeback_plan_entity_source
	plan_name = var.chargeback_plan_plan_name
	plan_type = var.chargeback_plan_plan_type
	plan_description = var.chargeback_plan_plan_description
	plan_custom_items {
		name  = var.chargeback_plan_plan_custom_items_name
		value = var.chargeback_plan_plan_custom_items_value

		#Optional
		is_customizable = var.chargeback_plan_plan_custom_items_is_customizable
	}
	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
* `entity_source` - (Required) Source of the chargeback plan.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
* `plan_custom_items` - (Required) (Updatable) List of chargeback plan customizations. At least one item is required.
	* `is_customizable` - (Optional) (Updatable) Indicates whether the chargeback plan customization item can be customized.
	* `name` - (Required) (Updatable) Name of chargeback plan customization item. Example items for Exadata Insights Chargeback are statistic, percentile, infrastructureCost, additionalServerCost etc.
	* `value` - (Required) (Updatable) Value of chargeback plan customization item.
* `plan_description` - (Required) (Updatable) Description of OPSI Chargeback Plan.
* `plan_name` - (Required) (Updatable) Name for the OPSI Chargeback plan.
* `plan_type` - (Required) Chargeback Plan type of the chargeback entity. For an Exadata it can be WEIGHTED_ALLOCATION, EQUAL_ALLOCATION, UNUSED_ALLOCATION.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
* `entity_source` - Source of the chargeback plan.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
* `id` - [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of OPSI Chargeback plan resource.
* `is_customizable` - Indicates whether the chargeback plan can be customized.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `plan_category` - Chargeback Plan category of the chargeback entity. It can be OOB, or CUSTOM.
* `plan_custom_items` - List of chargeback plan customizations.
	* `is_customizable` - Indicates whether the chargeback plan customization item can be customized.
	* `name` - Name of chargeback plan customization item. Example items for Exadata Insights Chargeback are statistic, percentile, infrastructureCost, additionalServerCost etc.
	* `value` - Value of chargeback plan customization item.
* `plan_description` - Description of OPSI Chargeback Plan.
* `plan_name` - Name for the OPSI Chargeback plan.
* `plan_type` - Chargeback Plan type of the chargeback entity. For an Exadata it can be WEIGHTED_ALLOCATION, EQUAL_ALLOCATION, UNUSED_ALLOCATION.
* `state` - Chargeback Plan lifecycle states
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `time_created` - The date and time the chargeback plan was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_updated` - The time chargeback plan was updated. An RFC3339 formatted datetime string

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
* `create` - (Defaults to 20 minutes), when creating the Chargeback Plan
* `update` - (Defaults to 20 minutes), when updating the Chargeback Plan
* `delete` - (Defaults to 20 minutes), when destroying the Chargeback Plan


## Import

ChargebackPlans can be imported using the `id`, e.g.

```
$ terraform import oci_opsi_chargeback_plan.test_chargeback_plan "id"
```

