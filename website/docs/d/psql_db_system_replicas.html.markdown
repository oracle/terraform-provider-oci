---
subcategory: "Psql"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_psql_db_system_replicas"
sidebar_current: "docs-oci-datasource-psql-db_system_replicas"
description: |-
  Provides the list of Db System Replicas in Oracle Cloud Infrastructure Psql service
---

# Data Source: oci_psql_db_system_replicas
This data source provides the list of Db System Replicas in Oracle Cloud Infrastructure Psql service.

Returns a list of replica database systems.


## Example Usage

```hcl
data "oci_psql_db_system_replicas" "test_db_system_replicas" {
	#Required
	db_system_id = oci_psql_db_system.test_db_system.id
}
```

## Argument Reference

The following arguments are supported:

* `db_system_id` - (Required) A unique identifier for the database system.


## Attributes Reference

The following attributes are exported:

* `db_system_replica_collection` - The list of db_system_replica_collection.

### DbSystemReplica Reference

The following attributes are exported:

* `items` - List of replica database systems.
	* `id` - A unique identifier for the replica database system.
	* `region` - Region name of the replica database system region. Example: `us-phoenix-1`
