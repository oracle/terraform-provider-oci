---
subcategory: "Opsi"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opsi_chargeback_plans"
sidebar_current: "docs-oci-datasource-opsi-chargeback_plans"
description: |-
  Provides the list of Chargeback Plans in Oracle Cloud Infrastructure Opsi service
---

# Data Source: oci_opsi_chargeback_plans
This data source provides the list of Chargeback Plans in Oracle Cloud Infrastructure Opsi service.

Gets a list of Ops Insights chargeback plans.

## Example Usage

```hcl
data "oci_opsi_chargeback_plans" "test_chargeback_plans" {

	#Optional
	chargebackplan_id = oci_opsi_chargebackplan.test_chargebackplan.id
	compartment_id = var.compartment_id
	compartment_id_in_subtree = var.chargeback_plan_compartment_id_in_subtree
}
```

## Argument Reference

The following arguments are supported:

* `chargebackplan_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Ops Insights chargeback plan.
* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `compartment_id_in_subtree` - (Optional) A flag to search all resources within a given compartment and all sub-compartments.


## Attributes Reference

The following attributes are exported:

* `chargeback_plan_collection` - The list of chargeback_plan_collection.

### ChargebackPlan Reference

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

