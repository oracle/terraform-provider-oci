---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_db_versions"
sidebar_current: "docs-oci-datasource-database-db_versions"
description: |-
  Provides the list of Db Versions in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_db_versions
This data source provides the list of Db Versions in Oracle Cloud Infrastructure Database service.

Gets a list of supported Oracle database versions.

## Example Usage

```hcl
data "oci_database_db_versions" "test_db_versions" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	db_system_id = "${oci_database_db_system.test_db_system.id}"
	db_system_shape = "${var.db_version_db_system_shape}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `db_system_id` - (Optional) The DB system [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). If provided, filters the results to the set of database versions which are supported for the DB system.
* `db_system_shape` - (Optional) If provided, filters the results to the set of database versions which are supported for the given shape.


## Attributes Reference

The following attributes are exported:

* `db_versions` - The list of db_versions.

### DbVersion Reference

The following attributes are exported:

* `supports_pdb` - True if this version of the Oracle Database software supports pluggable databases.
* `version` - A valid Oracle Database version.

