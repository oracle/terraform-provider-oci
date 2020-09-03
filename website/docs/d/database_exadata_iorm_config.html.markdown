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

Gets `IORM` Setting for the requested Exadata DB System.
The default IORM Settings is pre-created in all the Exadata DB System.


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

* `db_plans` - Array of IORM Setting for all the database in this Exadata DB System 
	* `db_name` - Database Name. For default DbPlan, the dbName will always be `default` 
	* `flash_cache_limit` - Flash Cache limit, internally configured based on shares 
	* `share` - Relative priority of a database 
* `lifecycle_details` - Additional information about the current lifecycleState. 
* `objective` - Value for the IORM objective Default is "Auto" 
* `state` - The current config state of IORM settings for this Exadata System. 

