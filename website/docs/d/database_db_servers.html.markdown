---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_db_servers"
sidebar_current: "docs-oci-datasource-database-db_servers"
description: |-
  Provides the list of Db Servers in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_db_servers
This data source provides the list of Db Servers in Oracle Cloud Infrastructure Database service.

Lists the Exadata DB servers in the ExadataInfrastructureId and specified compartment.


## Example Usage

```hcl
data "oci_database_db_servers" "test_db_servers" {
	#Required
	compartment_id = var.compartment_id
	exadata_infrastructure_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id

	#Optional
	display_name = var.db_server_display_name
	state = var.db_server_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. The match is not case sensitive.
* `exadata_infrastructure_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ExadataInfrastructure.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state exactly.


## Attributes Reference

The following attributes are exported:

* `db_servers` - The list of db_servers.

### DbServer Reference

The following attributes are exported:

* `autonomous_virtual_machine_ids` - The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Virtual Machines associated with the Db server. 
* `autonomous_vm_cluster_ids` - The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous VM Clusters associated with the Db server. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `compute_model` - The compute model of the Autonomous Database. This is required if using the `computeCount` parameter. If using `cpuCoreCount` then it is an error to specify `computeModel` to a non-null value. ECPU compute model is the recommended model and OCPU compute model is legacy.
* `cpu_core_count` - The number of CPU cores enabled on the Db server.
* `db_node_ids` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Db nodes associated with the Db server. 
* `db_node_storage_size_in_gbs` - The allocated local node storage in GBs on the Db server.
* `db_server_patching_details` - The scheduling details for the quarterly maintenance window. Patching and system updates take place during the maintenance window. 
	* `estimated_patch_duration` - Estimated time, in minutes, to patch one database server.
	* `patching_status` - The status of the patching operation.
	* `time_patching_ended` - The time when the patching operation ended.
	* `time_patching_started` - The time when the patching operation started.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - The user-friendly name for the Db server. The name does not need to be unique.
* `exadata_infrastructure_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exacc Db server.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `max_cpu_count` - The total number of CPU cores available.
* `max_db_node_storage_in_gbs` - The total local node storage available in GBs.
* `max_memory_in_gbs` - The total memory available in GBs.
* `memory_size_in_gbs` - The allocated memory in GBs on the Db server.
* `shape` - The shape of the Db server. The shape determines the amount of CPU, storage, and memory resources available. 
* `state` - The current state of the Db server.
* `time_created` - The date and time that the Db Server was created.
* `vm_cluster_ids` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM Clusters associated with the Db server. 

