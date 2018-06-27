# oci_load_balancer_backend_health

## BackendHealth Singular DataSource

### BackendHealth Reference

The following attributes are exported:

* `health_check_results` - A list of the most recent health check results returned for the specified backend server. 
	* `health_check_status` - The result of the most recent health check. 
	* `source_ip_address` - The IP address of the health check status report provider. This identifier helps you differentiate same-subnet (private) load balancers that report health check status.  Example: `10.0.0.7` 
	* `subnet_id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the subnet hosting the load balancer that reported this health check status. 
	* `timestamp` - The date and time the data was retrieved, in the format defined by RFC3339.  Example: `2017-06-02T18:28:11+00:00` 
* `status` - The general health status of the specified backend server as reported by the primary and standby load balancers.  *   **OK:** Both health checks returned `OK`.  *   **WARNING:** One health check returned `OK` and one did not.  *   **CRITICAL:** Neither health check returned `OK`.  *   **UNKNOWN:** One or both health checks returned `UNKNOWN`, or the system was unable to retrieve metrics at this time. 



### Get Operation
Gets the current health status of the specified backend server.

The following arguments are supported:

* `backend_name` - (Required) The IP address and port of the backend server to retrieve the health status for.  Example: `10.0.0.3:8080` 
* `backend_set_name` - (Required) The name of the backend set associated with the backend server to retrieve the health status for.  Example: `example_backend_set` 
* `load_balancer_id` - (Required) The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the load balancer associated with the backend server health status to be retrieved.


### Example Usage

```hcl
data "oci_load_balancer_backend_health" "test_backend_health" {
	#Required
	backend_name = "${var.backend_health_backend_name}"
	backend_set_name = "${var.backend_health_backend_set_name}"
	load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
}
```
