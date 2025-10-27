---
subcategory: "Container Engine"
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
	node_pool_option_id = oci_containerengine_node_pool_option.test_node_pool_option.id

	#Optional
	compartment_id = var.compartment_id
	node_pool_k8s_version = var.node_pool_option_node_pool_k8s_version
	node_pool_os_arch = var.node_pool_option_node_pool_os_arch
	node_pool_os_type = var.node_pool_option_node_pool_os_type
	should_list_all_patch_versions = var.node_pool_option_should_list_all_patch_versions
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The OCID of the compartment.
* `node_pool_k8s_version` - (Optional) Filter node pool options by Kubernetes version.
* `node_pool_option_id` - (Required) The id of the option set to retrieve. Use "all" get all options, or use a cluster ID to get options specific to the provided cluster.
* `node_pool_os_arch` - (Optional) Filter node pool options by OS architecture.
* `node_pool_os_type` - (Optional) Filter node pool options by OS type.
* `should_list_all_patch_versions` - (Optional) Option to show all kubernetes patch versions


## Attributes Reference

The following attributes are exported:

* `images` - Deprecated. See sources. When creating a node pool, only image names contained in this property can be passed to the `node_image_name` property.
* `kubernetes_versions` - Available Kubernetes versions.
* `shapes` - Available shapes for nodes.
* `sources` - Available source of the node.
	* `image_id` - The OCID of the image.
	* `source_name` - The user-friendly name of the entity corresponding to the OCID. 
	* `source_type` - The source type of this option. `IMAGE` means the OCID is of an image. 

