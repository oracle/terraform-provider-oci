---
subcategory: "Psql"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_psql_db_system_primary_db_instance"
sidebar_current: "docs-oci-datasource-psql-db_system_primary_db_instance"
description: |-
  Provides details about a specific Db System Primary Db Instance in Oracle Cloud Infrastructure Psql service
---

# Data Source: oci_psql_db_system_primary_db_instance
This data source provides details about a specific Db System Primary Db Instance resource in Oracle Cloud Infrastructure Psql service.

Gets the primary database instance node details.

## Example Usage

```hcl
data "oci_psql_db_system_primary_db_instance" "test_db_system_primary_db_instance" {
	#Required
	db_system_id = oci_psql_db_system.test_db_system.id
}
```

## Argument Reference

The following arguments are supported:

* `db_system_id` - (Required) A unique identifier for the database system.


## Attributes Reference

The following attributes are exported:

* `db_instance_id` - A unique identifier for the primary database instance node.

