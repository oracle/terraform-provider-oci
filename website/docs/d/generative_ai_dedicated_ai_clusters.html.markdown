---
subcategory: "Generative AI"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_dedicated_ai_clusters"
sidebar_current: "docs-oci-datasource-generative_ai-dedicated_ai_clusters"
description: |-
  Provides the list of Dedicated Ai Clusters in Oracle Cloud Infrastructure Generative AI service
---

# Data Source: oci_generative_ai_dedicated_ai_clusters
This data source provides the list of Dedicated Ai Clusters in Oracle Cloud Infrastructure Generative AI service.

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


