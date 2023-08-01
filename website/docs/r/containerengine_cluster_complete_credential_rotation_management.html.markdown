---
subcategory: "Container Engine"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_containerengine_cluster_complete_credential_rotation_management"
sidebar_current: "docs-oci-resource-containerengine-cluster_complete_credential_rotation_management"
description: |-
  Provides the Cluster Complete Credential Rotation Management resource in Oracle Cloud Infrastructure Container Engine service
---

# oci_containerengine_cluster_complete_credential_rotation_management
This resource provides the Cluster Complete Credential Rotation Management resource in Oracle Cloud Infrastructure Container Engine service.

Complete cluster credential rotation. Retire old credentials from kubernetes components.

## Example Usage

```hcl
resource "oci_containerengine_cluster_complete_credential_rotation_management" "test_cluster_complete_credential_rotation_management" {
	#Required
	cluster_id = oci_containerengine_cluster.test_cluster.id
	depends_on = [oci_containerengine_cluster_start_credential_rotation_management.test_cluster_start_credential_rotation_management]
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Required) The OCID of the cluster.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Cluster Complete Credential Rotation Management
	* `update` - (Defaults to 20 minutes), when updating the Cluster Complete Credential Rotation Management
	* `delete` - (Defaults to 20 minutes), when destroying the Cluster Complete Credential Rotation Management


## Import

Import is not supported for this resource.

