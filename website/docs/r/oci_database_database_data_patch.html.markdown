---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_data_patch"
sidebar_current: "docs-oci-resource-database-data-patch"
description: |-
  Provides the Database Data Patch resource in Oracle Cloud Infrastructure Database service
---

# oci_database_data_patch

The `oci_database_data_patch` resource provides the Database Data Patch resource in Oracle Cloud Infrastructure Database service.

## Example Usage

```hcl
resource "oci_database_data_patch" "example" {
  database_id = oci_database_database.example.id
  action = "APPLY"

  data_patch_options {
    should_skip_closed_pdbs = true
  }

  pluggable_databases = [oci_database_pluggable_database.example.id]
}
```

## Argument Reference

The following arguments are supported:

* `database_id` - (Required) The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `action` - (Required) The action to perform on the Data Patch.
* `data_patch_options` - (Optional) The Data Patch options.
    * `should_skip_closed_pdbs` - (Optional) Indicates whether to skip the data patch on closed PDBs.
* `pluggable_databases` - (Optional) The list of Pluggable Databases to be patched.

## Attributes Reference

The following attributes are exported:

The `oci_database_data_patch` resource does not export any additional attributes beyond the arguments.

## Import

The `oci_database_data_patch` resource is not importable.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/timeouts) for certain operations.
* `create` - (Defaults to 1 hour), when creating the Database Data Patch.
