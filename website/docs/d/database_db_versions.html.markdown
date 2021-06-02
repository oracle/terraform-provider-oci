---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_db_versions"
sidebar_current: "docs-oci-datasource-database-db_versions"
description: |-
  Provides the list of Db Versions in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_db_versions
This data source provides the list of Db Versions in Oracle Cloud Infrastructure Database service.

Gets a list of supported Oracle Database versions.

## Example Usage

```hcl
data "oci_database_db_versions" "test_db_versions" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	db_system_id = oci_database_db_system.test_db_system.id
	db_system_shape = var.db_version_db_system_shape
	is_database_software_image_supported = var.db_version_is_database_software_image_supported
	is_upgrade_supported = var.db_version_is_upgrade_supported
	storage_management = var.db_version_storage_management
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `db_system_id` - (Optional) The DB system [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). If provided, filters the results to the set of database versions which are supported for the DB system.
* `db_system_shape` - (Optional) If provided, filters the results to the set of database versions which are supported for the given shape.
* `is_database_software_image_supported` - (Optional) If true, filters the results to the set of Oracle Database versions that are supported for Oracle Cloud Infrastructure database software images.
* `is_upgrade_supported` - (Optional) If provided, filters the results to the set of database versions which are supported for Upgrade.
* `storage_management` - (Optional) The DB system storage management option. Used to list database versions available for that storage manager. Valid values are:
	* ASM - Automatic storage management
	* LVM - Logical volume management 


## Attributes Reference

The following attributes are exported:

* `db_versions` - The list of db_versions.

### DbVersion Reference

The following attributes are exported:

* `is_latest_for_major_version` - True if this version of the Oracle Database software is the latest version for a release.
* `is_preview_db_version` - True if this version of the Oracle Database software is the preview version.
* `is_upgrade_supported` - True if this version of the Oracle Database software is supported for Upgrade.
* `supports_pdb` - True if this version of the Oracle Database software supports pluggable databases.
* `version` - A valid Oracle Database version.

