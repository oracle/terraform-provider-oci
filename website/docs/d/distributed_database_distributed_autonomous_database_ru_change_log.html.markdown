---
subcategory: "Distributed Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_distributed_database_distributed_autonomous_database_ru_change_log"
sidebar_current: "docs-oci-datasource-distributed_database-distributed_autonomous_database_ru_change_log"
description: |-
  Provides details about a specific Distributed Autonomous Database Ru Change Log in Oracle Cloud Infrastructure Distributed Database service
---

# Data Source: oci_distributed_database_distributed_autonomous_database_ru_change_log
This data source provides details about a specific Distributed Autonomous Database Ru Change Log resource in Oracle Cloud Infrastructure Distributed Database service.

Operation to retrieve move RU invocation history and status for the Globally distributed autonomous database.
Results are ordered by created timestamp in descending order.


## Example Usage

```hcl
data "oci_distributed_database_distributed_autonomous_database_ru_change_log" "test_distributed_autonomous_database_ru_change_log" {
	#Required
	distributed_autonomous_database_id = oci_distributed_database_distributed_autonomous_database.test_distributed_autonomous_database.id

	#Optional
	source_shard_name = var.distributed_autonomous_database_ru_change_log_source_shard_name
}
```

## Argument Reference

The following arguments are supported:

* `distributed_autonomous_database_id` - (Required) Globally distributed autonomous database identifier.
* `source_shard_name` - (Optional) The source shard name to filter RU change log entries. 


## Attributes Reference

The following attributes are exported:

* `ru_change_log` - Ordered list of move RU invocations (latest first).
	* `ru_id` - Replication unit identifier associated with the invocation.
	* `source_shard_name` - Source shard for move RU invocation.
	* `status` - Current or terminal status of the move RU invocation.
	* `target_shard_name` - Target shard for move RU invocation.
	* `time_created` - Entry creation timestamp in UTC.
	* `time_updated` - Entry update timestamp in UTC.
	* `triggered_by` - Actor who triggered the move RU invocation.

