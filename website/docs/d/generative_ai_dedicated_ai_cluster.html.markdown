---
subcategory: "Generative Ai"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_dedicated_ai_cluster"
sidebar_current: "docs-oci-datasource-generative_ai-dedicated_ai_cluster"
description: |-
  Provides details about a specific Dedicated Ai Cluster in Oracle Cloud Infrastructure Generative Ai service
---

# Data Source: oci_generative_ai_dedicated_ai_cluster
This data source provides details about a specific Dedicated Ai Cluster resource in Oracle Cloud Infrastructure Generative Ai service.

Gets information about a dedicated AI cluster.

## Example Usage

```hcl
data "oci_generative_ai_dedicated_ai_cluster" "test_dedicated_ai_cluster" {
	#Required
	dedicated_ai_cluster_id = oci_generative_ai_dedicated_ai_cluster.test_dedicated_ai_cluster.id
}
```

## Argument Reference

The following arguments are supported:

* `dedicated_ai_cluster_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the dedicated AI cluster.


## Attributes Reference

The following attributes are exported:

* `capacity` - The total capacity for a dedicated AI cluster.
	* `capacity_type` - The type of the dedicated AI cluster capacity.
	* `total_endpoint_capacity` - The total number of endpoints that can be hosted on this dedicated AI cluster.
	* `used_endpoint_capacity` - The number of endpoints hosted on this dedicated AI cluster.
* `compartment_id` - The compartment OCID to create the dedicated AI cluster in.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - An optional description of the dedicated AI cluster.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the dedicated AI cluster.
* `lifecycle_details` - A message describing the current state with detail that can provide actionable information.
* `state` - The current state of the dedicated AI cluster.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the dedicated AI cluster was created, in the format defined by RFC 3339
* `time_updated` - The date and time the dedicated AI cluster was updated, in the format defined by RFC 3339
* `type` - The dedicated AI cluster type indicating whether this is a fine-tuning/training processor or hosting/inference processor.
* `unit_count` - The number of dedicated units in this AI cluster.
* `unit_shape` - The shape of dedicated unit in this AI cluster. The underlying hardware configuration is hidden from customers.

