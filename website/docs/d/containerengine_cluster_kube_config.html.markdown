---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_containerengine_cluster_kube_config"
sidebar_current: "docs-oci-datasource-containerengine-cluster_kube_config"
description: |-
  Provides details about a specific Cluster Kube Config in Oracle Cloud Infrastructure Container Engine service
---

# Data Source: oci_containerengine_cluster_kube_config
This data source provides details about a specific Cluster Kube Config resource in Oracle Cloud Infrastructure Container Engine service.

Create the Kubeconfig YAML for a cluster.

## Example Usage

```hcl
data "oci_containerengine_cluster_kube_config" "test_cluster_kube_config" {
	#Required
	cluster_id = "${oci_containerengine_cluster.test_cluster.id}"

	#Optional
	expiration = "${var.cluster_kube_config_expiration}"
	token_version = "${var.cluster_kube_config_token_version}"
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Required) The OCID of the cluster.
* `expiration` - (Optional) The desired expiration, in seconds, to use for the kubeconfig token. Important Note, expiration field is only honored for token version 1.0.0 
* `token_version` - (Optional) The version of the kubeconfig token. Supported values 1.0.0 and 2.0.0 


## Attributes Reference

The following attributes are exported:

* `content` - content of the Kubeconfig YAML for the cluster.
