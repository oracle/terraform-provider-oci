---
subcategory: "Container Engine"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: "
sidebar_current: "docs-oci_containerengine_cluster_public_api_endpoint_decommission_manager"
description: |-
  Provides the Cluster Public Api Endpoint Decommission Manager resource in Oracle Cloud Infrastructure Container Engine service
---

# oci_containerengine_cluster_public_api_endpoint_decommission_manager
This resource provides the Cluster Public Api Endpoint Decommission Managerresource in Oracle Cloud Infrastructure Container Engine service.

Request or Rollback Public Api Endpoint Decommission for a cluster.

## Example Usage

```hcl
resource "oci_containerengine_cluster_public_api_endpoint_decommission_manager" "decommission_manager"{
    #Required
    cluster_id = oci_containerengine_cluster.test_cluster.id
    is_public_api_endpoint_decommissioned = true

    #Optional
    rollback_deadline_delay = "P3D"
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Required) The OCID of the cluster.
* `is_public_api_endpoint_decommissioned` - (Required)(Updatable) Controls if a public API endpoint decommission or a rollback will happen . true is for raising public api endpoint decommission, false is for rollback public api endpoint decommission
* `rollback_deadline_delay` - (Optional)(Updatable) Extend rollback deadline for this cluster. 

## Attributes Reference

## Timeouts

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Cluster Public Api Endpoint Decommission Manager
	* `update` - (Defaults to 20 minutes), when updating the Cluster Public Api Endpoint Decommission Manager
	* `delete` - (Defaults to 20 minutes), when destroying the Cluster Public Api Endpoint Decommission Manager


## Import

Import is not supported for this resource.