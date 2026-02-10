---
subcategory: "Database Migration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_migration_script"
sidebar_current: "docs-oci-datasource-database_migration-script"
description: |-
  Provides details about a specific Script in Oracle Cloud Infrastructure Database Migration service
---

# Data Source: oci_database_migration_script
This data source provides details about a specific Script resource in Oracle Cloud Infrastructure Database Migration service.

Download DMS script.

## Example Usage

```hcl
data "oci_database_migration_script" "test_script" {
	#Required
	script_id = oci_database_migration_script.test_script.id
}
```

## Argument Reference

The following arguments are supported:

* `script_id` - (Required) The ID of the script to download.


## Attributes Reference

The following attributes are exported:


