---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_exadata_iorm_config"
sidebar_current: "docs-oci-datasource-database-exadata_iorm_config"
description: |-
  Provides details about a specific Exadata Iorm Config in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_exadata_iorm_config
This data source provides details about a specific Exadata Iorm Config resource in Oracle Cloud Infrastructure Database service.

Gets the IORM configuration settings for the specified cloud Exadata DB system.
All Exadata service instances have default IORM settings.

**Note:** Deprecated for Exadata Cloud Service systems. Use the [new resource model APIs](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/exaflexsystem.htm#exaflexsystem_topic-resource_model) instead.

For Exadata Cloud Service instances, support for this API will end on May 15th, 2021. See [Switching an Exadata DB System to the New Resource Model and APIs](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/exaflexsystem_topic-resource_model_conversion.htm) for details on converting existing Exadata DB systems to the new resource model.

The [GetCloudVmClusterIormConfig](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/CloudVmCluster/GetCloudVmClusterIormConfig/) API is used for this operation with Exadata systems using the
new resource model.


## Example Usage

```hcl
data "oci_database_exadata_iorm_config" "test_exadata_iorm_config" {
	#Required
	db_system_id = oci_database_db_system.test_db_system.id
}
```

## Argument Reference

The following arguments are supported:

* `db_system_id` - (Required) The DB system [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `db_plans` - An array of IORM settings for all the database in the Exadata DB system. 
	* `db_name` - The database name. For the default `DbPlan`, the `dbName` is `default`. 
	* `flash_cache_limit` - The flash cache limit for this database. This value is internally configured based on the share value assigned to the database. 
	* `share` - The relative priority of this database. 
* `lifecycle_details` - Additional information about the current `lifecycleState`. 
* `objective` - The current value for the IORM objective. The default is `AUTO`. 
* `state` - The current state of IORM configuration for the Exadata DB system. 

