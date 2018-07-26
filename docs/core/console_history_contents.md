# oci_core_console_history_data

## ConsoleHistoryContent Singular DataSource

### ConsoleHistoryContent Reference

The following attributes are exported:

* `data` - The console history data.


### Get Operation
Gets the actual console history data (not the metadata).
See [CaptureConsoleHistory](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/ConsoleHistory/CaptureConsoleHistory)
for details about using the console history operations.


The following arguments are supported:

* `console_history_id` - (Required) The OCID of the console history.
* `length` - (Optional) Length of the snapshot data to retrieve.
* `offset` - (Optional) Offset of the snapshot data to retrieve.


### Example Usage

```hcl
data "oci_core_console_history_content" "test_console_history_content" {
	#Required
	console_history_id = "${oci_core_console_history.test_console_history.id}"

	#Optional
	length = "${var.console_history_content_length}"
	offset = "${var.console_history_content_offset}"
}
```
