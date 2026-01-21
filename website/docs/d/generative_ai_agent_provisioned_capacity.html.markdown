---
subcategory: "Generative Ai Agent"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_agent_provisioned_capacity"
sidebar_current: "docs-oci-datasource-generative_ai_agent-provisioned_capacity"
description: |-
  Provides details about a specific Provisioned Capacity in Oracle Cloud Infrastructure Generative Ai Agent service
---

# Data Source: oci_generative_ai_agent_provisioned_capacity
This data source provides details about a specific Provisioned Capacity resource in Oracle Cloud Infrastructure Generative Ai Agent service.

Gets information about a provisioned capacity.


## Example Usage

```hcl
data "oci_generative_ai_agent_provisioned_capacity" "test_provisioned_capacity" {
	#Required
	provisioned_capacity_id = oci_generative_ai_agent_provisioned_capacity.test_provisioned_capacity.id
}
```

## Argument Reference

The following arguments are supported:

* `provisioned_capacity_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the provisioned capacity.


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

