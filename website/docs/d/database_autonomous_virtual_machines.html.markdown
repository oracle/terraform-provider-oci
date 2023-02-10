---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_virtual_machines"
sidebar_current: "docs-oci-datasource-database-autonomous_virtual_machines"
description: |-
  Provides the list of Autonomous Virtual Machines in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_autonomous_virtual_machines
This data source provides the list of Autonomous Virtual Machines in Oracle Cloud Infrastructure Database service.

Lists the Autonomous Virtual Machines in the specified Autonomous VM Cluster and Compartment.


## Example Usage

```hcl
data "oci_database_autonomous_virtual_machines" "test_autonomous_virtual_machines" {
	#Required
	autonomous_vm_cluster_id = oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster.id
	compartment_id = var.compartment_id

	#Optional
	state = var.autonomous_virtual_machine_state
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_vm_cluster_id` - (Required) The Autonomous Virtual machine [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `state` - (Optional) A filter to return only resources that match the given lifecycle state exactly.


## Attributes Reference

The following attributes are exported:

* `autonomous_virtual_machines` - The list of autonomous_virtual_machines.

### AutonomousVirtualMachine Reference

The following attributes are exported:

* `autonomous_vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous VM Cluster associated with the Autonomous Virtual Machine.
* `client_ip_address` - Client IP Address.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `cpu_core_count` - The number of CPU cores enabled on the Autonomous Virtual Machine.
* `db_node_storage_size_in_gbs` - The allocated local node storage in GBs on the Autonomous Virtual Machine.
* `db_server_display_name` - The display name of the dbServer associated with the Autonomous Virtual Machine.
* `db_server_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Db server associated with the Autonomous Virtual Machine.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Virtual Machine.
* `memory_size_in_gbs` - The allocated memory in GBs on the Autonomous Virtual Machine.
* `state` - The current state of the Autonomous Virtual Machine.
* `vm_name` - The name of the Autonomous Virtual Machine.

