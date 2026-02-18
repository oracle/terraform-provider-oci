---
subcategory: "Container Engine"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_containerengine_cluster_public_api_endpoint_decommission_status"
sidebar_current: "docs-oci-datasource-containerengine-cluster_public_api_endpoint_decommission_status"
description: |-
  Provides details about a specific Cluster Public Api Endpoint Decommission Status in Oracle Cloud Infrastructure Container Engine service
---

# Data Source: oci_containerengine_cluster_public_api_endpoint_decommission_status
This data source provides details about a specific Cluster Public Api Endpoint Decommission Status resource in Oracle Cloud Infrastructure Container Engine service.

Get cluster public api endpoint decommission status.

## Example Usage

```hcl
data "oci_containerengine_cluster_public_api_endpoint_decommission_status" "test_cluster_public_api_endpoint_decommission_status" {
	#Required
	cluster_id = oci_containerengine_cluster.test_cluster.id
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Required) The OCID of the cluster.


## Attributes Reference

The following attributes are exported:

* `status` - The current public api endpoint decommission status of the cluster.
* `time_decommission_rollback_deadline` - The date and time of rollback deadline for public api endpoint decommission.  Once the date is passed, rollback is not able to be launched. 

