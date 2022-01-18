---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_vm_cluster"
sidebar_current: "docs-oci-datasource-database-autonomous_vm_cluster"
description: |-
  Provides details about a specific Autonomous Vm Cluster in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_autonomous_vm_cluster
This data source provides details about a specific Autonomous Vm Cluster resource in Oracle Cloud Infrastructure Database service.

Gets information about the specified Autonomous VM cluster for an Exadata Cloud@Customer system. To get information about an Autonomous VM Cluster in the Oracle cloud, see [GetCloudAutonomousVmCluster](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/CloudAutonomousVmCluster/GetCloudAutonomousVmCluster). 


## Example Usage

```hcl
data "oci_database_autonomous_vm_cluster" "test_autonomous_vm_cluster" {
	#Required
	autonomous_vm_cluster_id = oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster.id
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_vm_cluster_id` - (Required) The autonomous VM cluster [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `available_cpus` - The numnber of CPU cores available.
* `available_data_storage_size_in_tbs` - The data storage available in TBs
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `cpus_enabled` - The number of enabled CPU cores.
* `data_storage_size_in_gb` - The total data storage allocated in GBs.
* `data_storage_size_in_tbs` - The total data storage allocated in TBs
* `db_node_storage_size_in_gbs` - The local node storage allocated in GBs.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - The user-friendly name for the Autonomous VM cluster. The name does not need to be unique.
* `exadata_infrastructure_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous VM cluster.
* `is_local_backup_enabled` - If true, database backup on local Exadata storage is configured for the Autonomous VM cluster. If false, database backup on local Exadata storage is not available in the Autonomous VM cluster. 
* `license_model` - The Oracle license model that applies to the Autonomous VM cluster. The default is LICENSE_INCLUDED. 
* `lifecycle_details` - Additional information about the current lifecycle state.
* `memory_size_in_gbs` - The memory allocated in GBs.
* `ocpus_enabled` - The number of enabled OCPU cores.
* `state` - The current state of the Autonomous VM cluster.
* `time_created` - The date and time that the Autonomous VM cluster was created.
* `time_zone` - The time zone to use for the Autonomous VM cluster. For details, see [DB System Time Zones](https://docs.cloud.oracle.com/iaas/Content/Database/References/timezones.htm).
* `vm_cluster_network_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster network.

