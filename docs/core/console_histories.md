# oci\_core\_console_history

## ConsoleHistory Resource

### ConsoleHistory Reference

The following attributes are exported:

* `availability_domain` - The Availability Domain of an instance.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - The OCID of the compartment.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My console history metadata` 
* `id` - The OCID of the console history metadata object.
* `instance_id` - The OCID of the instance this console history was fetched from.
* `state` - The current state of the console history.
* `time_created` - The date and time the history was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 



### Create Operation
Captures the most recent serial console data (up to a megabyte) for the
specified instance.

The `CaptureConsoleHistory` operation works with the other console history operations
as described below.

1. Use `CaptureConsoleHistory` to request the capture of up to a megabyte of the
most recent console history. This call returns a `ConsoleHistory`
object. The object will have a state of REQUESTED.
2. Wait for the capture operation to succeed by polling `GetConsoleHistory` with
the identifier of the console history metadata. The state of the
`ConsoleHistory` object will go from REQUESTED to GETTING-HISTORY and
then SUCCEEDED (or FAILED).
3. Use `GetConsoleHistoryContent` to get the actual console history data (not the
metadata).
4. Optionally, use `DeleteConsoleHistory` to delete the console history metadata
and the console history data.


The following arguments are supported:

* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable. 
* `instance_id` - (Required) The OCID of the instance to get the console history from.


### Update Operation
Updates the specified console history metadata.

The following arguments support updates:
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```
resource "oci_core_console_history" "test_console_history" {
	#Required
	instance_id = "${oci_core_instance.test_instance.id}"

	#Optional
	display_name = "${var.console_history_display_name}"
}
```

# oci\_core\_console_histories

## ConsoleHistoryData DataSource

Gets the history data for a specific console.

### Get Operation
Gets the actual console history data (not the metadata).

The following arguments are supported:

* `console_history_id` - (Required) The OCID of the console history.
* `length` - (Optional) Length of the snapshot data to retrieve. Cannot be less than 10240.
* `offset` - (Optional) Offset of the snapshot data to retrieve.

The following attributes are exported:

* `data` - The console history data.

### Example Usage

```
data "oci_core_console_history_data" "test_console_history_data" {
	#Required
	console_history_id = "${oci_core_console_history.test_console_history.id}"

	#Optional
	length = 10240
	offset = 0
}
```