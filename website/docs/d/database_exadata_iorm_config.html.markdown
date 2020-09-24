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

Gets the IORM configuration settings for the specified Exadata DB system.
All Exadata DB systems have default IORM settings.


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

