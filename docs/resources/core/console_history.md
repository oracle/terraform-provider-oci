 # baremetal\_core\_console_history

 Provides a console history resource.

 ## Example Usage

 ```
    resource "baremetal_core_console_history" "t" {
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
 * `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
 * `id` - The OCID of the console history metadata object.
 * `instance_id` - The OCID of the instance this console history was fetched from.
 * `state` - The current state of the console history.
 * `time_created` - The date and time the database was created.
