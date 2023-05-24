---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_exadata_infrastructure_un_allocated_resource"
sidebar_current: "docs-oci-datasource-database-exadata_infrastructure_un_allocated_resource"
description: |-
  Provides details about a specific Exadata Infrastructure Un Allocated Resource in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_exadata_infrastructure_un_allocated_resource
This data source provides details about a specific Exadata Infrastructure Un Allocated Resource resource in Oracle Cloud Infrastructure Database service.

Gets un allocated resources information for the specified Exadata infrastructure. Applies to Exadata Cloud@Customer instances only.


## Example Usage

```hcl
data "oci_database_exadata_infrastructure_un_allocated_resource" "test_exadata_infrastructure_un_allocated_resource" {
	#Required
	exadata_infrastructure_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id

	#Optional
	db_servers = var.exadata_infrastructure_un_allocated_resource_db_servers
}
```

## Argument Reference

The following arguments are supported:

* `db_servers` - (Optional) The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Db servers.
* `exadata_infrastructure_id` - (Required) The Exadata infrastructure [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `autonomous_vm_clusters` - The list of Autonomous VM Clusters on the Infra and their associated unallocated resources details
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.
	* `un_allocated_adb_storage_in_tbs` - Total unallocated autonomous data storage in the AVM in TBs.
* `display_name` - The user-friendly name for the Exadata Cloud@Customer infrastructure. The name does not need to be unique.
* `exadata_storage_in_tbs` - Total unallocated exadata storage in the infrastructure in TBs.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.
* `local_storage_in_gbs` - The minimum amount of un allocated storage that is available across all nodes in the infrastructure.
* `memory_in_gbs` - The minimum amount of un allocated memory that is available across all nodes in the infrastructure.
* `ocpus` - The minimum amount of un allocated ocpus that is available across all nodes in the infrastructure.

