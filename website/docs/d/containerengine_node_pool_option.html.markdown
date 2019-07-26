---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_containerengine_node_pool_option"
sidebar_current: "docs-oci-datasource-containerengine-node_pool_option"
description: |-
  Provides details about a specific Node Pool Option in Oracle Cloud Infrastructure Container Engine service
---

# Data Source: oci_containerengine_node_pool_option
This data source provides details about a specific Node Pool Option resource in Oracle Cloud Infrastructure Container Engine service.

Get options available for node pools.

## Example Usage

```hcl
data "oci_containerengine_node_pool_option" "test_node_pool_option" {
	#Required
	node_pool_option_id = "${oci_containerengine_node_pool_option.test_node_pool_option.id}"
}
```

## Argument Reference

The following arguments are supported:

* `node_pool_option_id` - (Required) The id of the option set to retrieve. Use "all" get all options, or use a cluster ID to get options specific to the provided cluster.


## Attributes Reference

The following attributes are exported:

* `images` - Available image names.
* `kubernetes_versions` - Available Kubernetes versions.
* `shapes` - Available shapes for nodes.

