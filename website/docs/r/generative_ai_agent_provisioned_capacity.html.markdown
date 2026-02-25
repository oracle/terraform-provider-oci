---
subcategory: "Generative Ai Agent"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_agent_provisioned_capacity"
sidebar_current: "docs-oci-resource-generative_ai_agent-provisioned_capacity"
description: |-
  Provides the Provisioned Capacity resource in Oracle Cloud Infrastructure Generative Ai Agent service
---

# oci_generative_ai_agent_provisioned_capacity
This resource provides the Provisioned Capacity resource in Oracle Cloud Infrastructure Generative Ai Agent service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/generative-ai-agents/latest/ProvisionedCapacity

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/generative_ai_agent

Creates a provisioned capacity.


## Example Usage

```hcl
resource "oci_generative_ai_agent_provisioned_capacity" "test_provisioned_capacity" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.provisioned_capacity_display_name
	number_of_units = var.provisioned_capacity_number_of_units

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.provisioned_capacity_description
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the endpoint in. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) An optional description of the provisioned capacity.
* `display_name` - (Required) (Updatable) The name of the provisioned capacity.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `number_of_units` - (Required) (Updatable) Provisioned Capacity Unit corresponds to the amount of characters processed per minute.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - An optional description of the provisioned capacity.
* `display_name` - The name of the provisioned capacity.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the provisioned capacity.
* `number_of_units` - Provisioned Capacity Unit corresponds to the amount of characters processed per minute.
* `state` - The current state of the provisioned capacity.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the provisioned capacity was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the provisioned capacity was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Provisioned Capacity
	* `update` - (Defaults to 20 minutes), when updating the Provisioned Capacity
	* `delete` - (Defaults to 20 minutes), when destroying the Provisioned Capacity


## Import

ProvisionedCapacities can be imported using the `id`, e.g.

```
$ terraform import oci_generative_ai_agent_provisioned_capacity.test_provisioned_capacity "id"
```

