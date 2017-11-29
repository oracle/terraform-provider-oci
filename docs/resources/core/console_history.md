# oci\_core\_console_history

[ConsoleHistory Reference][03f57b12]

  [03f57b12]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/ConsoleHistory/ "ConsoleHistoryReference"

 Provides a console history resource.

## Example Usage

 ```
    resource "oci_core_console_history" "t" {
			instance_id = "instance_id"
    }
 ```

## Argument Reference

 The following arguments are supported:

 * `instance_id` - (Required) The OCID of the console history.

## Attributes Reference

 The following attributes are exported:

 * `availability_domain` - The Availability Domain of an instance.
 * `compartment_id` - The OCID of the compartment.
 * `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
 * `id` - The OCID of the console history metadata object.
 * `instance_id` - The OCID of the instance this console history was fetched from.
 * `state` - The current state of the console history. Allowed values are: [REQUESTED, GETTING-HISTORY, SUCCEEDED, FAILED]
 * `time_created` - The date and time the history was created, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z`.
