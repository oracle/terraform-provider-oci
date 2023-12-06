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

Gets the primary DbInstance details.

## Example Usage

```hcl
data "oci_psql_db_system_primary_db_instance" "test_db_system_primary_db_instance" {
	#Required
	db_system_id = oci_psql_db_system.test_db_system.id
}
```

## Argument Reference

The following arguments are supported:

* `db_system_id` - (Required) unique DbSystem identifier


## Attributes Reference

The following attributes are exported:

* `db_instance_id` - Unique identifier of the DbInstance.

