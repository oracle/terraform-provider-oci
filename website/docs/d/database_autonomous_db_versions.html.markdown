---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_db_versions"
sidebar_current: "docs-oci-datasource-database-autonomous_db_versions"
description: |-
  Provides the list of Autonomous Db Versions in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_autonomous_db_versions
This data source provides the list of Autonomous Db Versions in Oracle Cloud Infrastructure Database service.

Gets a list of supported Autonomous Database versions.

## Example Usage

```hcl
data "oci_database_autonomous_db_versions" "test_autonomous_db_versions" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	db_workload = var.autonomous_db_version_db_workload
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `db_workload` - (Optional) A filter to return only autonomous database resources that match the specified workload type.


## Attributes Reference

The following attributes are exported:

* `autonomous_db_versions` - The list of autonomous_db_versions.

### AutonomousDbVersion Reference

The following attributes are exported:

* `db_workload` - The Autonomous Database workload type. The following values are valid:
	* OLTP - indicates an Autonomous Transaction Processing database
	* DW - indicates an Autonomous Data Warehouse database
	* AJD - indicates an Autonomous JSON Database
	* APEX - indicates an Autonomous Database with the Oracle APEX Application Development workload type. 
* `details` - A URL that points to a detailed description of the Autonomous Database version.
* `is_dedicated` - True if the database uses [dedicated Exadata infrastructure](https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html). 
* `is_default_for_free` - True if this version of the Oracle Database software's default is free.
* `is_default_for_paid` - True if this version of the Oracle Database software's default is paid.
* `is_free_tier_enabled` - True if this version of the Oracle Database software can be used for Always-Free Autonomous Databases.
* `is_paid_enabled` - True if this version of the Oracle Database software has payments enabled.
* `version` - A valid Oracle Database version for Autonomous Database.

