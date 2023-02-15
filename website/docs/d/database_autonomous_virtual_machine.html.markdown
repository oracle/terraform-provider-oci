---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_virtual_machine"
sidebar_current: "docs-oci-datasource-database-autonomous_virtual_machine"
description: |-
  Provides details about a specific Autonomous Virtual Machine in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_autonomous_virtual_machine
This data source provides details about a specific Autonomous Virtual Machine resource in Oracle Cloud Infrastructure Database service.

Gets the details of specific Autonomous Virtual Machine.


## Example Usage

```hcl
data "oci_database_autonomous_virtual_machine" "test_autonomous_virtual_machine" {
	#Required
	autonomous_virtual_machine_id = oci_database_autonomous_virtual_machine.test_autonomous_virtual_machine.id
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_virtual_machine_id` - (Required) The Autonomous Virtual machine [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

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

