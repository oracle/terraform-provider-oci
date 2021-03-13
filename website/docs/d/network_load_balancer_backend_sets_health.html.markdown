---
subcategory: "Network Load Balancer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_load_balancer_backend_sets_health"
sidebar_current: "docs-oci-datasource-network_load_balancer-backend_sets_health"
description: |-
  Provides details about a specific Backend Sets Health in Oracle Cloud Infrastructure Network Load Balancer service
---

# Data Source: oci_network_load_balancer_backend_sets_health
This data source provides details about a specific Backend Sets Health resource in Oracle Cloud Infrastructure Network Load Balancer service.

Retrieves the health status for the specified backend set.

## Example Usage

```hcl
data "oci_network_load_balancer_backend_sets_health" "test_backend_sets_health" {
	#Required
	backend_set_name = oci_network_load_balancer_backend_set.test_backend_set.name
	network_load_balancer_id = oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id
}
```

## Argument Reference

The following arguments are supported:

* `backend_set_name` - (Required) The name of the backend set for which to retrieve the health status.  Example: `example_backend_set` 
* `network_load_balancer_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network load balancer to update.


## Attributes Reference

The following attributes are exported:

* `critical_state_backend_names` - A list of backend servers that are currently in the `CRITICAL` health state. The list identifies each backend server by IP address and port.  Example: `10.0.0.4:8080` 
* `status` - Overall health status of the backend set.
	*  **OK:** All backend servers in the backend set return a status of `OK`.
	*  **WARNING:** Half or more of the backend servers in a backend set return a status of `OK` and at least one backend server returns a status of `WARNING`, `CRITICAL`, or `UNKNOWN`.
	*  **CRITICAL:** Fewer than half of the backend servers in a backend set return a status of `OK`.
	*  **UNKNOWN:** If no probes have yet been sent to the backends, or the system is unable to retrieve metrics from the backends. 
* `total_backend_count` - The total number of backend servers in this backend set.  Example: `7` 
* `unknown_state_backend_names` - A list of backend servers that are currently in the `UNKNOWN` health state. The list identifies each backend server by IP address and port.  Example: `10.0.0.5:8080` 
* `warning_state_backend_names` - A list of backend servers that are currently in the `WARNING` health state. The list identifies each backend server by IP address or OCID and port.  Example: `10.0.0.3:8080` or `ocid1.privateip..oc1.<var>&lt;unique_ID&gt;</var>:8080` 

