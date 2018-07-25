---
layout: "oci"
page_title: "OCI: oci_containerengine_node_pool_option"
sidebar_current: "docs-oci-datasource-containerengine-node_pool_option"
description: |-
  Provides details about a specific NodePoolOption
---

# Data Source: oci_containerengine_node_pool_option
The NodePoolOption data source provides details about a specific NodePoolOption

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

* `node_pool_option_id` - (Required) The id of the option set to retrieve. Only "all" is supported.


## Attributes Reference

The following attributes are exported:

* `images` - Available Kubernetes versions.
* `kubernetes_versions` - Available Kubernetes versions.
* `shapes` - Available shapes for nodes.

