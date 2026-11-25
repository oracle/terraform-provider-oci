---
subcategory: "Distributed Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_distributed_database_distributed_autonomous_databases_add_gds_control_node"
sidebar_current: "docs-oci-resource-distributed_database-distributed_autonomous_databases_add_gds_control_node"
description: |-
  Provides the Distributed Autonomous Databases Add Gds Control Node resource in Oracle Cloud Infrastructure Distributed Database service
---

# oci_distributed_database_distributed_autonomous_databases_add_gds_control_node
This resource provides the Distributed Autonomous Databases Add Gds Control Node resource in Oracle Cloud Infrastructure Distributed Database service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/distributed_database

Add new Global database services control(GDS CTL) node for the Globally distributed autonomous database.


## Example Usage

```hcl
resource "oci_distributed_database_distributed_autonomous_databases_add_gds_control_node" "test_distributed_autonomous_databases_add_gds_control_node" {
	#Required
	distributed_autonomous_database_id = oci_distributed_database_distributed_autonomous_database.test_distributed_autonomous_database.id
	public_ssh_key = var.distributed_autonomous_databases_add_gds_control_node_public_ssh_key

	#Optional
	subnet_id = oci_core_subnet.test_subnet.id
}
```

## Argument Reference

The following arguments are supported:

* `distributed_autonomous_database_id` - (Required) Globally distributed autonomous database identifier.
* `public_ssh_key` - (Required) The public sshKey for Global database services control(GDS CTL) node.
* `subnet_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet on which the Global database services control(GDS CTL) node should be created.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Distributed Autonomous Databases Add Gds Control Node
	* `update` - (Defaults to 20 minutes), when updating the Distributed Autonomous Databases Add Gds Control Node
	* `delete` - (Defaults to 20 minutes), when destroying the Distributed Autonomous Databases Add Gds Control Node


## Import

Import is not supported for this resource.

