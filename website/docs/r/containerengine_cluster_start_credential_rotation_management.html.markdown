---
subcategory: "Container Engine"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_containerengine_cluster_start_credential_rotation_management"
sidebar_current: "docs-oci-resource-containerengine-cluster_start_credential_rotation_management"
description: |-
  Provides the Cluster Start Credential Rotation Management resource in Oracle Cloud Infrastructure Container Engine service
---

# oci_containerengine_cluster_start_credential_rotation_management
This resource provides the Cluster Start Credential Rotation Management resource in Oracle Cloud Infrastructure Container Engine service.

Start cluster credential rotation by adding new credentials, old credentials will still work after this operation.

## Example Usage

```hcl
resource "oci_containerengine_cluster_start_credential_rotation_management" "test_cluster_start_credential_rotation_management" {
	#Required
	auto_completion_delay_duration = var.cluster_start_credential_rotation_management_auto_completion_delay_duration
	cluster_id = oci_containerengine_cluster.test_cluster.id
}
```

## Argument Reference

The following arguments are supported:

* `auto_completion_delay_duration` - (Required) The duration in days(in ISO 8601 notation eg. P5D) after which the old credentials should be retired. Maximum delay duration is 14 days.
* `cluster_id` - (Required) The OCID of the cluster.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Cluster Start Credential Rotation Management
	* `update` - (Defaults to 20 minutes), when updating the Cluster Start Credential Rotation Management
	* `delete` - (Defaults to 20 minutes), when destroying the Cluster Start Credential Rotation Management


## Import

Import is not supported for this resource.

