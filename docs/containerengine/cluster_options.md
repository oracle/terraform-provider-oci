# oci_containerengine_cluster_option

## ClusterOption Singular DataSource

### ClusterOption Reference

The following attributes are exported:

* `kubernetes_versions` - Available Kubernetes versions.



### Get Operation
Get options available for clusters.

The following arguments are supported:

* `cluster_option_id` - (Required) The id of the option set to retrieve. Only "all" is supported.


### Example Usage

```hcl
data "oci_containerengine_cluster_option" "test_cluster_option" {
	#Required
	cluster_option_id = "${oci_containerengine_cluster_option.test_cluster_option.id}"
}
```
