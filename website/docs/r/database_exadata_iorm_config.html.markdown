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

Update `IORM` Settings for the requested Exadata DB System.


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
	* `db_name` - (Required) (Updatable) Database Name. For updating default DbPlan, pass in dbName as `default` 
	* `share` - (Required) (Updatable) Relative priority of a database 
* `db_system_id` - (Required) (Updatable) The DB system [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `objective` - (Optional) (Updatable) Value for the IORM objective Default is "Auto" 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `db_plans` - Array of IORM Setting for all the database in this Exadata DB System 
	* `db_name` - Database Name. For default DbPlan, the dbName will always be `default` 
	* `flash_cache_limit` - Flash Cache limit, internally configured based on shares 
	* `share` - Relative priority of a database 
* `lifecycle_details` - Additional information about the current lifecycleState. 
* `objective` - Value for the IORM objective Default is "Auto" 
* `state` - The current config state of IORM settings for this Exadata System. 

## Import

Import is not supported for this resource.

