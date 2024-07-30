---
subcategory: "Generative AI"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_dedicated_ai_cluster"
sidebar_current: "docs-oci-datasource-generative_ai-dedicated_ai_cluster"
description: |-
  Provides details about a specific Dedicated Ai Cluster in Oracle Cloud Infrastructure Generative AI service
---

# Data Source: oci_generative_ai_dedicated_ai_cluster
This data source provides details about a specific Dedicated Ai Cluster resource in Oracle Cloud Infrastructure Generative AI service.

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

* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the dedicated AI cluster.

