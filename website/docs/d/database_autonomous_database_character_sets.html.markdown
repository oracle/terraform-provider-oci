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
}
```

## Argument Reference

The following arguments are supported:



## Attributes Reference

The following attributes are exported:

* `autonomous_database_character_sets` - The list of autonomous_database_character_sets.

### AutonomousDatabaseCharacterSet Reference

The following attributes are exported:

* `name` - A valid Oracle character set.

