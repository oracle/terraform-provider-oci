---
subcategory: "Big Data Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bds_bds_instance_replace_node_action"
sidebar_current: "docs-oci-resource-bds-bds_instance_replace_node_action"
description: |-
  Provides the Bds Instance Replace Node Action resource in Oracle Cloud Infrastructure Big Data Service service
---

# oci_bds_bds_instance_replace_node_action

Invokes the Big Data Service replace node action for a cluster node.

Use this action resource to replace a node identified by `node_host_name`. You can optionally provide a specific node backup to restore from, a target shape for the replacement node, and either `cluster_admin_password` or `secret_id` for authentication.

When `node_backup_id` is omitted, the service uses the latest available node backup. If no suitable backup is available, or the original node is already in a failed or terminated state, the service attempts to recover from the last saved boot volume state.

## Example Usage

```hcl
resource "oci_bds_bds_instance_replace_node_action" "test_bds_instance_replace_node_action" {
  # Required
  bds_instance_id = var.bds_instance_id
  node_host_name  = var.node_host_name

  # Provide one of cluster_admin_password or secret_id
  cluster_admin_password = "T3JhY2xlVGVhbVVTQSExMjM="

  # Optional
  node_backup_id = var.node_backup_id
  shape          = var.shape
}
```

## Argument Reference

The following arguments are supported:

* `bds_instance_id` - (Required) The OCID of the Big Data Service cluster.
* `node_host_name` - (Required) Host name of the node to replace.
* `cluster_admin_password` - (Optional) Base64-encoded cluster admin password. Use this or `secret_id`.
* `node_backup_id` - (Optional) The OCID of the node backup to use for replacement.
* `secret_id` - (Optional) The OCID of the secret that stores the cluster admin password. Use this or `cluster_admin_password`.
* `shape` - (Optional) The shape to use for the replacement node. If not specified, the existing node shape is used.

**IMPORTANT**
This is an action resource. Any change forces Terraform to create the action resource again and invoke the replace node workflow.

## Attributes Reference

The following attributes are exported:

* `id` - Terraform identifier for the completed replace node action.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:

* `create` - (Defaults to 60 minutes), when creating the Bds Instance Replace Node Action