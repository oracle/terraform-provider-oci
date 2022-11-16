---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_cloud_vm_cluster_iorm_config"
sidebar_current: "docs-oci-resource-database-cloud_vm_cluster_iorm_config"
description: |-
  Provides the Cloud Vm Cluster Iorm Config resource in Oracle Cloud Infrastructure Database service
---

# oci_database_cloud_vm_cluster_iorm_config
This resource provides the Cloud Vm Cluster Iorm Config resource in Oracle Cloud Infrastructure Database service.

Updates IORM settings for the specified Cloud Vm Cluster.

The [UpdateCloudVmClusterIormConfig](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/CloudVmCluster/UpdateCloudVmClusterIormConfig/) API is used for Cloud Vm Cluster.


## Example Usage

```hcl
resource "oci_database_cloud_vm_cluster_iorm_config" "test_cloud_vm_cluster_iorm_config" {
	#Required
	db_plans {
		#Required
		db_name = var.cloud_vm_cluster_iorm_config_db_plans_db_name
		share = var.cloud_vm_cluster_iorm_config_db_plans_share
	}
	cloud_vm_cluster_id = oci_database_cloud_vm_cluster.test_cloud_vm_cluster.id

	#Optional
	objective = "AUTO"
}
```

## Argument Reference

The following arguments are supported:

* `db_plans` - (Required) (Updatable) Array of IORM Setting for all the database in this Cloud Vm Cluster 
	* `db_name` - (Required) (Updatable) The database name. For the default `DbPlan`, the `dbName` is `default`. 
	* `share` - (Required) (Updatable) The relative priority of this database. 
* `cloud_vm_cluster_id` - (Required) The Cluster [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `objective` - (Optional) (Updatable) Value for the IORM objective Default is "Auto" 

## Attributes Reference

The following attributes are exported:

* `db_plans` - An array of IORM settings for all the database in the Cloud Vm Cluster. 
	* `db_name` - The database name. For the default `DbPlan`, the `dbName` is `default`. 
	* `flash_cache_limit` - The flash cache limit for this database. This value is internally configured based on the share value assigned to the database. 
	* `share` - The relative priority of this database. 
* `lifecycle_details` - Additional information about the current `lifecycleState`. 
* `objective` - The current value for the IORM objective. The default is `AUTO`. 
<<<<<<< ours
* `state` - The current state of IORM configuration for the Exadata DB system. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Cloud Vm Cluster Iorm Config
	* `update` - (Defaults to 20 minutes), when updating the Cloud Vm Cluster Iorm Config
	* `delete` - (Defaults to 20 minutes), when destroying the Cloud Vm Cluster Iorm Config


## Import

CloudVmClusterIormConfigs can be imported using the `id`, e.g.

```
$ terraform import oci_database_cloud_vm_cluster_iorm_config.test_cloud_vm_cluster_iorm_config "cloudVmClusters/{cloudVmClusterId}/CloudVmClusterIormConfig" 
```

=======
* `state` - The current state of IORM configuration for the Cloud Vm Cluster. 
>>>>>>> theirs
