---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_cloud_vm_cluster_iorm_config"
sidebar_current: "docs-oci-datasource-database-cloud_vm_cluster_iorm_config"
description: |-
  Provides details about a specific Cloud Vm Cluster Iorm Config in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_cloud_vm_cluster_iorm_config
This data source provides details about a specific Cloud Vm Cluster Iorm Config resource in Oracle Cloud Infrastructure Database service.

Gets the IORM configuration settings for the specified Cloud Vm Cluster.
All Exadata service instances have default IORM settings.

The [GetCloudVmClusterIormConfig](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/CloudVmCluster/GetCloudVmClusterIormConfig/) API is used for this operation with Cloud Vm Cluster.


## Example Usage

```hcl
data "oci_database_cloud_vm_cluster_iorm_config" "test_cloud_vm_cluster_iorm_config" {
	#Required
	cloud_vm_cluster_id = oci_database_cloud_vm_cluster.test_cloud_vm_cluster.id
}
```

## Argument Reference

The following arguments are supported:

* `cloud_vm_cluster_id` - (Required) The cluster [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `db_plans` - An array of IORM settings for all the database in the cloud vm cluster. 
	* `db_name` - The database name. For the default `DbPlan`, the `dbName` is `default`. 
	* `flash_cache_limit` - The flash cache limit for this database. This value is internally configured based on the share value assigned to the database. 
	* `share` - The relative priority of this database. 
* `lifecycle_details` - Additional information about the current `lifecycleState`. 
* `objective` - The current value for the IORM objective. The default is `AUTO`. 
* `state` - The current state of IORM configuration for the cloud vm cluster. 
