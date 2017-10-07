# oci\_core\_console_history

**API:**
[ConsoleHistory Reference][d1aa5bf0]

  [d1aa5bf0]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/ConsoleHistory/ "ConsoleHistoryReference"

Gets the history data for a specific console.

## Example Usage

```
data "oci_core_console_history_data" "s" {
      console_history_id = "ichid"
      length = 1
      offset = 1
    }
```

## Argument Reference

The following arguments are supported:

* `console_history_id` - (Required) The OCID of the console history.
* `offset` - (Optional) Offset of the snapshot data to retrieve.
* `length` - (Optional) Length of the snapshot data to retrieve.

## Attributes Reference

The following attributes are exported:

* `availability_domain` - The Availability Domain of an instance.
* `compartment_id` - The OCID of the compartment.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `id` - The OCID of the console history metadata object.
* `instance_id` - The OCID of the instance this console history was fetched from.
* `state` - The current state of the console history. Allowed values are: [REQUESTED, GETTING-HISTORY, SUCCEEDED, FAILED].
* `time_created` - The date and time the database was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z`.
