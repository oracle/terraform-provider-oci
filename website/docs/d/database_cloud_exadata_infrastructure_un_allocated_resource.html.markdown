---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_cloud_exadata_infrastructure_un_allocated_resource"
sidebar_current: "docs-oci-datasource-database-cloud_exadata_infrastructure_un_allocated_resource"
description: |-
  Provides details about a specific Cloud Exadata Infrastructure Un Allocated Resource in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_cloud_exadata_infrastructure_un_allocated_resource
This data source provides details about a specific Cloud Exadata Infrastructure Un Allocated Resource resource in Oracle Cloud Infrastructure Database service.

Gets un allocated resources information for the specified Cloud Exadata infrastructure.


## Example Usage

```hcl
data "oci_database_cloud_exadata_infrastructure_un_allocated_resource" "test_cloud_exadata_infrastructure_un_allocated_resource" {
	#Required
	cloud_exadata_infrastructure_id = oci_database_cloud_exadata_infrastructure.test_cloud_exadata_infrastructure.id
}
```

## Argument Reference

The following arguments are supported:

* `cloud_exadata_infrastructure_id` - (Required) The cloud Exadata infrastructure [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `cloud_autonomous_vm_clusters` - The list of Cloud Autonomous VM Clusters on the Infra and their associated unallocated resources details
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Cloud Exadata infrastructure.
	* `un_allocated_adb_storage_in_tbs` - Total unallocated autonomous data storage in the CAVM in TBs.
* `cloud_exadata_infrastructure_display_name` - The user-friendly name for the Cloud Exadata infrastructure. The name does not need to be unique.
* `cloud_exadata_infrastructure_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Cloud Exadata infrastructure.
* `exadata_storage_in_tbs` - Total unallocated exadata storage in the infrastructure in TBs.
* `local_storage_in_gbs` - The minimum amount of un allocated storage that is available across all nodes in the infrastructure.
* `memory_in_gbs` - The minimum amount of un allocated memory that is available across all nodes in the infrastructure.
* `ocpus` - The minimum amount of un allocated ocpus that is available across all nodes in the infrastructure.

