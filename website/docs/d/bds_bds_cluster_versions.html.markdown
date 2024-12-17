---
subcategory: "Big Data Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bds_bds_cluster_versions"
sidebar_current: "docs-oci-datasource-bds-bds_cluster_versions"
description: |-
  Provides the list of Bds Cluster Versions in Oracle Cloud Infrastructure Big Data Service service
---

# Data Source: oci_bds_bds_cluster_versions
This data source provides the list of Bds Cluster Versions in Oracle Cloud Infrastructure Big Data Service service.

Returns a list of cluster versions with associated odh and bds versions.


## Example Usage

```hcl
data "oci_bds_bds_cluster_versions" "test_bds_cluster_versions" {
}
```

## Argument Reference

The following arguments are supported:



## Attributes Reference

The following attributes are exported:

* `bds_cluster_versions` - The list of bds_cluster_versions.

### BdsClusterVersion Reference

The following attributes are exported:

* `bds_version` - BDS version to be used for cluster creation
* `odh_version` - ODH version to be used for cluster creation

