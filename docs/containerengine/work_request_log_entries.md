
# oci_containerengine_work_request_log_entries

## WorkRequestLogEntry DataSource

Gets a list of work_request_log_entries.

### List Operation
Get the logs of a work request.
The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `work_request_id` - (Required) The OCID of the work request.


The following attributes are exported:

* `work_request_log_entries` - The list of work_request_log_entries.

### Example Usage

```hcl
data "oci_containerengine_work_request_log_entries" "test_work_request_log_entries" {
	#Required
	compartment_id = "${var.compartment_id}"
	work_request_id = "${oci_containerengine_work_request.test_work_request.id}"
}
```
### WorkRequestLogEntry Reference

The following attributes are exported:

* `message` - The description of an action that occurred.
* `timestamp` - The date and time the log entry occurred.
