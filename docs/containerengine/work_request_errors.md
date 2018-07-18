
# oci_containerengine_work_request_errors

## WorkRequestError DataSource

Gets a list of work_request_errors.

### List Operation
Get the errors of a work request.
The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `work_request_id` - (Required) The OCID of the work request.


The following attributes are exported:

* `work_request_errors` - The list of work_request_errors.

### Example Usage

```hcl
data "oci_containerengine_work_request_errors" "test_work_request_errors" {
	#Required
	compartment_id = "${var.compartment_id}"
	work_request_id = "${oci_containerengine_work_request.test_work_request.id}"
}
```
### WorkRequestError Reference

The following attributes are exported:

* `code` - A short error code that defines the error, meant for programmatic parsing. See [API Errors](https://docs.us-phoenix-1.oraclecloud.com/Content/API/References/apierrors.htm).
* `message` - A human-readable error string.
* `timestamp` - The date and time the error occurred.
