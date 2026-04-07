---
subcategory: "Distributed Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_distributed_database_distributed_database_raft_metric"
sidebar_current: "docs-oci-datasource-distributed_database-distributed_database_raft_metric"
description: |-
  Provides details about a specific Distributed Database Raft Metric in Oracle Cloud Infrastructure Distributed Database service
---

# Data Source: oci_distributed_database_distributed_database_raft_metric
This data source provides details about a specific Distributed Database Raft Metric resource in Oracle Cloud Infrastructure Distributed Database service.

Operation to retrieve RAFT metrics for the Globally distributed database. If the Globally distributed database is not
RAFT based then empty response is returned from the API.


## Example Usage

```hcl
data "oci_distributed_database_distributed_database_raft_metric" "test_distributed_database_raft_metric" {
	#Required
	distributed_database_id = oci_distributed_database_distributed_database.test_distributed_database.id
}
```

## Argument Reference

The following arguments are supported:

* `distributed_database_id` - (Required) Globally distributed database identifier


## Attributes Reference

The following attributes are exported:

* `config_tasks` - Details of in-progress configuration tasks.
* `raft_metrics` - Raft metrics for the Globally distributed database.

