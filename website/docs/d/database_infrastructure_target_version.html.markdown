---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_infrastructure_target_version"
sidebar_current: "docs-oci-datasource-database-infrastructure_target_version"
description: |-
  Provides details about a specific Infrastructure Target Version in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_infrastructure_target_version
This data source provides details about a specific Infrastructure Target Version resource in Oracle Cloud Infrastructure Database service.

Gets details of the Exadata Infrastructure target system software versions that can be applied to the specified infrastructure resource for maintenance updates.
Applies to Exadata Cloud@Customer and Exadata Cloud instances only.


## Example Usage

```hcl
data "oci_database_infrastructure_target_version" "test_infrastructure_target_version" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	target_resource_id = oci_database_target_resource.test_target_resource.id
	target_resource_type = var.infrastructure_target_version_target_resource_type
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `target_resource_id` - (Optional) The target resource ID.
* `target_resource_type` - (Optional) The type of the target resource.


## Attributes Reference

The following attributes are exported:

* `target_db_version_history_entry` - The history entry of the target system software version for the database server patching operation.
* `target_resource_id` - The OCID of the target Exadata Infrastructure resource that will receive the maintenance update.
* `target_resource_type` - The resource type of the target Exadata infrastructure resource that will receive the system software update.
* `target_storage_version_history_entry` - The history entry of the target storage cell system software version for the storage cell patching operation.

