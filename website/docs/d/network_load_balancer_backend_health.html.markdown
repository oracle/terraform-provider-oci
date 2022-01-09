---
subcategory: "Network Load Balancer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_load_balancer_backend_health"
sidebar_current: "docs-oci-datasource-network_load_balancer-backend_health"
description: |-
  Provides details about a specific Backend Health in Oracle Cloud Infrastructure Network Load Balancer service
---

# Data Source: oci_network_load_balancer_backend_health
This data source provides details about a specific Backend Health resource in Oracle Cloud Infrastructure Network Load Balancer service.

Retrieves the current health status of the specified backend server.

## Example Usage

```hcl
data "oci_network_load_balancer_backend_health" "test_backend_health" {
	#Required
	backend_name = oci_network_load_balancer_backend.test_backend.name
	backend_set_name = oci_network_load_balancer_backend_set.test_backend_set.name
	network_load_balancer_id = oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id
}
```

## Argument Reference

The following arguments are supported:

* `backend_name` - (Required) The name of the backend server to retrieve health status for. If the backend was created with an explicitly specified name, that name should be used here. If the backend was created without explicitly specifying the name, but was created using ipAddress, this is specified as <ipAddress>:<port>. If the backend was created without explicitly specifying the name, but was created using targetId, this is specified as <targetId>:<port>.  Example: `10.0.0.3:8080` or `ocid1.privateip..oc1.<var>&lt;unique_ID&gt;</var>:8080` 
* `backend_set_name` - (Required) The name of the backend set associated with the backend server for which to retrieve the health status.  Example: `example_backend_set` 
* `network_load_balancer_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network load balancer to update.


## Attributes Reference

The following attributes are exported:

* `health_check_results` - A list of the most recent health check results returned for the specified backend server. 
	* `health_check_status` - The result of the most recent health check. 
	* `timestamp` - The date and time the data was retrieved, in the format defined by RFC3339.  Example: `2020-05-01T18:28:11+00:00` 
* `status` - The general health status of the specified backend server.
	*   **OK:**  All health check probes return `OK`
	*   **WARNING:** At least one of the health check probes does not return `OK`
	*   **CRITICAL:** None of the health check probes return `OK`. *
	*   **UNKNOWN:** One of the health checks probes return `UNKNOWN`,
	*   or the system is unable to retrieve metrics at this time. 

