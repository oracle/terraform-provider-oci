---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_database_character_sets"
sidebar_current: "docs-oci-datasource-database-autonomous_database_character_sets"
description: |-
  Provides the list of Autonomous Database Character Sets in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_autonomous_database_character_sets
This data source provides the list of Autonomous Database Character Sets in Oracle Cloud Infrastructure Database service.

Gets a list of supported character sets.

## Example Usage

```hcl
data "oci_database_autonomous_database_character_sets" "test_autonomous_database_character_sets" {

	#Optional
	character_set_type = var.autonomous_database_character_set_character_set_type
	is_dedicated = var.autonomous_database_character_set_is_dedicated
	is_shared = var.autonomous_database_character_set_is_shared
}
```

## Argument Reference

The following arguments are supported:

* `character_set_type` - (Optional) Specifies whether this request pertains to database character sets or national character sets. 
* `is_dedicated` - (Optional) Specifies if the request is for an Autonomous Database Dedicated instance. The default request is for an Autonomous Database Dedicated instance.
* `is_shared` - (Optional) Specifies whether this request is for Autonomous Database on Shared infrastructure. By default, this request will be for Autonomous Database on Dedicated Exadata Infrastructure.

:
## Attributes Reference

The following attributes are exported:

* `autonomous_database_character_sets` - The list of autonomous_database_character_sets.

### AutonomousDatabaseCharacterSet Reference

The following attributes are exported:

* `name` - A valid Oracle character set.

