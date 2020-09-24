---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_migration"
sidebar_current: "docs-oci-resource-database-migration"
description: |-
  Provides the Migration resource in Oracle Cloud Infrastructure Database service
---

# oci_database_migration
This resource provides the Migration resource in Oracle Cloud Infrastructure Database service.

Migration Exadata dbSystem resource model to cloud Exadata infrastructure model. All related resources will be migrated.


## Example Usage

```hcl
resource "oci_database_migration" "test_migration" {
	#Required
	db_system_id = "${oci_database_db_system.test_db_system.id}"
}
```

## Argument Reference

The following arguments are supported:

* `db_system_id` - (Required) The DB system [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `additional_migrations` - The details of addtional resources related to the migration.
	* `cloud_exadata_infrastructure_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Exadata infrastructure.
	* `cloud_vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud VM cluster.
	* `db_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB system.
* `cloud_exadata_infrastructure_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Exadata infrastructure.
* `cloud_vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud VM cluster.
* `db_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB system.

## Import

Import is not supported for this resource.

