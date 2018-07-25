---
layout: "oci"
page_title: "OCI: oci_containerengine_cluster_option"
sidebar_current: "docs-oci-datasource-containerengine-cluster_option"
description: |-
  Provides details about a specific ClusterOption
---

# Data Source: oci_containerengine_cluster_option
The ClusterOption data source provides details about a specific ClusterOption

Get options available for clusters.

## Example Usage

```hcl
data "oci_containerengine_cluster_option" "test_cluster_option" {
	#Required
	cluster_option_id = "${oci_containerengine_cluster_option.test_cluster_option.id}"
}
```

## Argument Reference

The following arguments are supported:

* `cluster_option_id` - (Required) The id of the option set to retrieve. Only "all" is supported.


## Attributes Reference

The following attributes are exported:

* `kubernetes_versions` - Available Kubernetes versions.

