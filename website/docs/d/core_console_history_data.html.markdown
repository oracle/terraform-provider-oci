---
layout: "oci"
page_title: "OCI: oci_core_console_history_data"
sidebar_current: "docs-oci-datasource-core-console_history_data"
description: |-
Provides details about a specific ConsoleHistory
---

# Data Source: oci_core_console_history_data
The ConsoleHistory data source provides details about a specific ConsoleHistory

Gets the history data for a specific console.

## Example Usage

```hcl
data "oci_core_console_history_data" "test_console_history_data" {
	#Required
	console_history_id = "${oci_core_console_history.test_console_history.id}"

	#Optional
	length = 10240
	offset = 0
}
```

## Argument Reference

The following arguments are supported:

* `console_history_id` - (Required) The OCID of the console history.
* `length` - (Optional) Length of the snapshot data to retrieve. Cannot be less than 10240.
* `offset` - (Optional) Offset of the snapshot data to retrieve.


## Attributes Reference

The following attributes are exported:

* `data` - The console history data.

