---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_vm_clusters"
sidebar_current: "docs-oci-datasource-database-autonomous_vm_clusters"
description: |-
  Provides the list of Autonomous Vm Clusters in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_autonomous_vm_clusters
This data source provides the list of Autonomous Vm Clusters in Oracle Cloud Infrastructure Database service.

Gets a list of Exadata Cloud@Customer Autonomous VM clusters in the specified compartment. To list Autonomous VM Clusters in the Oracle Cloud, see [ListCloudAutonomousVmClusters](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/CloudAutonomousVmCluster/ListCloudAutonomousVmClusters).


## Example Usage

```hcl
data "oci_database_autonomous_vm_clusters" "test_autonomous_vm_clusters" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.autonomous_vm_cluster_display_name
	exadata_infrastructure_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id
	state = var.autonomous_vm_cluster_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. The match is not case sensitive.
* `exadata_infrastructure_id` - (Optional) If provided, filters the results for the given Exadata Infrastructure.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state exactly.


## Attributes Reference

The following attributes are exported:

* `autonomous_vm_clusters` - The list of autonomous_vm_clusters.

### AutonomousVmCluster Reference

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

