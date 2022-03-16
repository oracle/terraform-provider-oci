---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_databases_asm_properties"
sidebar_current: "docs-oci-datasource-database_management-managed_databases_asm_properties"
description: |-
  Provides the list of Managed Databases Asm Properties in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_databases_asm_properties
This data source provides the list of Managed Databases Asm Properties in Oracle Cloud Infrastructure Database Management service.

Gets the list of ASM properties for the specified managedDatabaseId.

## Example Usage

```hcl
data "oci_database_management_managed_databases_asm_properties" "test_managed_databases_asm_properties" {
	#Required
	managed_database_id = oci_database_management_managed_database.test_managed_database.id

	#Optional
	name = var.managed_databases_asm_property_name
}
```

## Argument Reference

The following arguments are supported:

* `managed_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
* `name` - (Optional) A filter to return only resources that match the entire name.


## Attributes Reference

The following attributes are exported:

* `asm_property_collection` - The list of asm_property_collection.

### ManagedDatabasesAsmProperty Reference

The following attributes are exported:

* `items` - An array of AsmPropertySummary resources.
	* `disk_group` - The name of the disk group.

