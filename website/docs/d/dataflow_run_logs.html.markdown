---
subcategory: "Data Flow"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dataflow_run_logs"
sidebar_current: "docs-oci-datasource-dataflow-run_logs"
description: |-
  Provides the list of Run Logs in Oracle Cloud Infrastructure Data Flow service
---

# Data Source: oci_dataflow_run_logs
This data source provides the list of Run Logs in Oracle Cloud Infrastructure Data Flow service.

Retrieves summaries of the run's logs.


## Example Usage

```hcl
data "oci_dataflow_run_logs" "test_run_logs" {
	#Required
	run_id = oci_dataflow_run.test_run.id
}
```

## Argument Reference

The following arguments are supported:

* `run_id` - (Required) The unique ID for the run 


## Attributes Reference

The following attributes are exported:

* `run_logs` - The list of run_logs.

### RunLog Reference

The following attributes are exported:


