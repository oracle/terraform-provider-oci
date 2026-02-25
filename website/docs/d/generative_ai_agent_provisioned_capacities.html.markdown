---
subcategory: "Generative Ai Agent"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_agent_provisioned_capacities"
sidebar_current: "docs-oci-datasource-generative_ai_agent-provisioned_capacities"
description: |-
  Provides the list of Provisioned Capacities in Oracle Cloud Infrastructure Generative Ai Agent service
---

# Data Source: oci_generative_ai_agent_provisioned_capacities
This data source provides the list of Provisioned Capacities in Oracle Cloud Infrastructure Generative Ai Agent service.

Gets a list of provisioned capacities.


## Example Usage

```hcl
data "oci_generative_ai_agent_provisioned_capacities" "test_provisioned_capacities" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.provisioned_capacity_display_name
	provisioned_capacity_id = oci_generative_ai_agent_provisioned_capacity.test_provisioned_capacity.id
	state = var.provisioned_capacity_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly.
* `provisioned_capacity_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the provisioned capacity.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state. The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `provisioned_capacity_collection` - The list of provisioned_capacity_collection.

### ProvisionedCapacity Reference

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

