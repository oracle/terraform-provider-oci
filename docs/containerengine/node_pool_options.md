# oci_containerengine_node_pool_option

## NodePoolOption Singular DataSource

### NodePoolOption Reference

The following attributes are exported:

* `images` - Available Kubernetes versions.
* `kubernetes_versions` - Available Kubernetes versions.
* `shapes` - Available shapes for nodes.



### Get Operation
Get options available for node pools.

The following arguments are supported:

* `node_pool_option_id` - (Required) The id of the option set to retrieve. Only "all" is supported.


### Example Usage

```hcl
data "oci_containerengine_node_pool_option" "test_node_pool_option" {
	#Required
	node_pool_option_id = "${oci_containerengine_node_pool_option.test_node_pool_option.id}"
}
```
