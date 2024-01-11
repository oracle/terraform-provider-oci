---
subcategory: "Generative Ai"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_dedicated_ai_clusters"
sidebar_current: "docs-oci-datasource-generative_ai-dedicated_ai_clusters"
description: |-
  Provides the list of Dedicated Ai Clusters in Oracle Cloud Infrastructure Generative Ai service
---

# Data Source: oci_generative_ai_dedicated_ai_clusters
This data source provides the list of Dedicated Ai Clusters in Oracle Cloud Infrastructure Generative Ai service.

Lists the dedicated AI clusters in a specific compartment.

## Example Usage

```hcl
data "oci_generative_ai_dedicated_ai_clusters" "test_dedicated_ai_clusters" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.dedicated_ai_cluster_display_name
	id = var.dedicated_ai_cluster_id
	state = var.dedicated_ai_cluster_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly.
* `id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the dedicated AI cluster.
* `state` - (Optional) A filter to return only the dedicated AI clusters that their lifecycle state matches the given lifecycle state.


## Attributes Reference

The following attributes are exported:

* `dedicated_ai_cluster_collection` - The list of dedicated_ai_cluster_collection.

### DedicatedAiCluster Reference

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

