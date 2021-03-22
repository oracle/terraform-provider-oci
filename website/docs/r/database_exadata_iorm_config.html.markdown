---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_exadata_iorm_config"
sidebar_current: "docs-oci-resource-database-exadata_iorm_config"
description: |-
  Provides the Exadata Iorm Config resource in Oracle Cloud Infrastructure Database service
---

# oci_database_exadata_iorm_config
This resource provides the Exadata Iorm Config resource in Oracle Cloud Infrastructure Database service.

Updates IORM settings for the specified Exadata DB system.

**Note:** Deprecated for Exadata Cloud Service systems. Use the [new resource model APIs](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/exaflexsystem.htm#exaflexsystem_topic-resource_model) instead.

For Exadata Cloud Service instances, support for this API will end on May 15th, 2021. See [Switching an Exadata DB System to the New Resource Model and APIs](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/exaflexsystem_topic-resource_model_conversion.htm) for details on converting existing Exadata DB systems to the new resource model.

The [UpdateCloudVmClusterIormConfig](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/CloudVmCluster/UpdateCloudVmClusterIormConfig/) API is used for Exadata systems using the
new resource model.


## Example Usage

```hcl
resource "oci_database_exadata_iorm_config" "test_exadata_iorm_config" {
	#Required
	db_plans {
		#Required
		db_name = var.exadata_iorm_config_db_plans_db_name
		share = var.exadata_iorm_config_db_plans_share
	}
	db_system_id = oci_database_db_system.test_db_system.id

	#Optional
	objective = "AUTO"
}
```

## Argument Reference

The following arguments are supported:

* `db_plans` - (Required) (Updatable) Array of IORM Setting for all the database in this Exadata DB System 
	* `db_name` - (Required) (Updatable) The database name. For the default `DbPlan`, the `dbName` is `default`. 
	* `share` - (Required) (Updatable) The relative priority of this database. 
* `db_system_id` - (Required) (Updatable) The DB system [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `objective` - (Optional) (Updatable) Value for the IORM objective Default is "Auto" 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `db_plans` - An array of IORM settings for all the database in the Exadata DB system. 
	* `db_name` - The database name. For the default `DbPlan`, the `dbName` is `default`. 
	* `flash_cache_limit` - The flash cache limit for this database. This value is internally configured based on the share value assigned to the database. 
	* `share` - The relative priority of this database. 
* `lifecycle_details` - Additional information about the current `lifecycleState`. 
* `objective` - The current value for the IORM objective. The default is `AUTO`. 
* `state` - The current state of IORM configuration for the Exadata DB system. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Exadata Iorm Config
	* `update` - (Defaults to 20 minutes), when updating the Exadata Iorm Config
	* `delete` - (Defaults to 20 minutes), when destroying the Exadata Iorm Config


## Import

Import is not supported for this resource.

